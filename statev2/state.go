package statev2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/hermeznetwork/hermez-core/encoding"
	"github.com/hermeznetwork/hermez-core/log"
	"github.com/hermeznetwork/hermez-core/merkletree"
	"github.com/hermeznetwork/hermez-core/statev2/runtime"
	"github.com/hermeznetwork/hermez-core/statev2/runtime/executor/pb"
	"github.com/hermeznetwork/hermez-core/statev2/runtime/fakevm"
	"github.com/hermeznetwork/hermez-core/statev2/runtime/instrumentation"
	"github.com/hermeznetwork/hermez-core/statev2/runtime/instrumentation/tracers"
	"github.com/holiman/uint256"
	"github.com/jackc/pgx/v4"
)

const (
	// TxTransferGas used for TXs that do not create a contract
	TxTransferGas uint64 = 21000
	// TxSmartContractCreationGas used for TXs that create a contract
	TxSmartContractCreationGas uint64 = 53000
	// Size of the memory in bytes reserved by the zkEVM
	zkEVMReservedMemorySize int  = 128
	two                     uint = 2
)

var (
	// ErrInvalidBatchHeader indicates the batch header is invalid
	ErrInvalidBatchHeader = errors.New("invalid batch header")
	// ErrStateNotSynchronized indicates the state database may be empty
	ErrStateNotSynchronized = errors.New("state not synchronized")
	// ErrNotFound indicates an object has not been found for the search criteria used
	ErrNotFound = errors.New("object not found")
	// ErrNilDBTransaction indicates the db transaction has not been properly initialized
	ErrNilDBTransaction = errors.New("database transaction not properly initialized")
	// ErrAlreadyInitializedDBTransaction indicates the db transaction was already initialized
	ErrAlreadyInitializedDBTransaction = errors.New("database transaction already initialized")
	// ErrNotEnoughIntrinsicGas indicates the gas is not enough to cover the intrinsic gas cost
	ErrNotEnoughIntrinsicGas = fmt.Errorf("not enough gas supplied for intrinsic gas costs")
	// ErrParsingExecutorTrace indicates an error occurred while parsing the executor trace
	ErrParsingExecutorTrace = fmt.Errorf("error while parsing executor trace")
	// ErrInvalidBatchNumber indicates the provided batch number is not the latest in db
	ErrInvalidBatchNumber = errors.New("provided batch number is not latest")
	// ErrNoIncludedTransactions
	ErrNoIncludedTransactions = errors.New("none of the batch transactions was processed")
)

var (
	// ZeroHash is the hash 0x0000000000000000000000000000000000000000000000000000000000000000
	ZeroHash = common.Hash{}
	// ZeroAddress is the address 0x0000000000000000000000000000000000000000
	ZeroAddress = common.Address{}
)

// State is a implementation of the state
type State struct {
	cfg Config
	*PostgresStorage
	executorClient pb.ExecutorServiceClient
	tree           *merkletree.StateTree
}

// NewState creates a new State
func NewState(cfg Config, storage *PostgresStorage, executorClient pb.ExecutorServiceClient, stateTree *merkletree.StateTree) *State {
	return &State{
		cfg:             cfg,
		PostgresStorage: storage,
		executorClient:  executorClient,
		tree:            stateTree,
	}
}

// BeginDBTransaction starts a state transaction
func (s *State) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := s.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Commit commits a state transaction
func (s *State) CommitStateTransaction(ctx context.Context, dbTx pgx.Tx) error {
	err := dbTx.Commit(ctx)
	return err
}

// Rollback rollbacks a state transaction
func (s *State) RollbackStateTransaction(ctx context.Context, dbTx pgx.Tx) error {
	err := dbTx.Rollback(ctx)
	return err
}

// GetBalance from a given address
func (s *State) GetBalance(ctx context.Context, address common.Address, blockNumber uint64, dbTx pgx.Tx) (*big.Int, error) {
	l2Block, err := s.GetL2BlockByNumber(ctx, blockNumber, dbTx)
	if err != nil {
		return nil, err
	}

	return s.tree.GetBalance(ctx, address, l2Block.Root().Bytes())
}

// GetCode from a given address
func (s *State) GetCode(ctx context.Context, address common.Address, blockNumber uint64, dbTx pgx.Tx) ([]byte, error) {
	l2Block, err := s.GetL2BlockByNumber(ctx, blockNumber, dbTx)
	if err != nil {
		return nil, err
	}

	return s.tree.GetCode(ctx, address, l2Block.Root().Bytes())
}

// GetNonce returns the nonce of the given account at the given block number
func (s *State) GetNonce(ctx context.Context, address common.Address, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	l2Block, err := s.GetL2BlockByNumber(ctx, blockNumber, dbTx)
	if err != nil {
		return 0, err
	}

	nonce, err := s.tree.GetNonce(ctx, address, l2Block.Root().Bytes())

	return nonce.Uint64(), err
}

// GetStorageAt from a given address
func (s *State) GetStorageAt(ctx context.Context, address common.Address, position *big.Int, blockNumber uint64, dbTx pgx.Tx) (*big.Int, error) {
	l2Block, err := s.GetL2BlockByNumber(ctx, blockNumber, dbTx)
	if err != nil {
		return nil, err
	}

	return s.tree.GetStorageAt(ctx, address, position, l2Block.Root().Bytes())
}

// EstimateGas for a transaction
func (s *State) EstimateGas(transaction *types.Transaction, senderAddress common.Address) (uint64, error) {
	// TODO: implement
	return 0, nil
}

// ProcessBatch is used by the Trusted Sequencer to add transactions to the batch
func (s *State) ProcessBatch(ctx context.Context, batchNumber uint64, txs []types.Transaction, dbTx pgx.Tx) (*ProcessBatchResponse, error) {
	lastBatches, err := s.PostgresStorage.GetLastNBatches(ctx, two, dbTx)
	if err != nil {
		return nil, err
	}

	// Get latest batch from the database to get GER and Timestamp
	lastBatch := lastBatches[0]
	// Get batch before latest to get state root and local exit root
	previousBatch := lastBatches[1]

	// Check provided batch number is the latest in db
	if lastBatch.BatchNumber != batchNumber {
		return nil, ErrInvalidBatchNumber
	}

	batchL2Data, err := EncodeTransactions(txs)
	if err != nil {
		return nil, err
	}

	// Create Batch
	processBatchRequest := &pb.ProcessBatchRequest{
		BatchNum:             lastBatch.BatchNumber,
		Coinbase:             lastBatch.Coinbase.String(),
		BatchL2Data:          batchL2Data,
		OldStateRoot:         previousBatch.OldStateRoot.Bytes(),
		GlobalExitRoot:       lastBatch.GlobalExitRootNum.Bytes(),
		OldLocalExitRoot:     previousBatch.OldLocalExitRoot.Bytes(),
		EthTimestamp:         uint64(lastBatch.Timestamp.Unix()),
		UpdateMerkleTree:     true,
		GenerateExecuteTrace: false,
		GenerateCallTrace:    false,
	}

	// Send Batch to the Executor
	processBatchResponse, err := s.executorClient.ProcessBatch(ctx, processBatchRequest)
	return convertToProcessBatchResponse(txs, processBatchResponse), err
}

// StoreTransactions is used by the Trusted Sequencer to add processed transactions into the data base
func (s *State) StoreTransactions(ctx context.Context, batchNumber uint64, processedTxs []*ProcessTransactionResponse, dbTx pgx.Tx) error {
	foundPosition := -1

	batch, err := s.PostgresStorage.GetBatchByNumber(ctx, batchNumber, dbTx)
	if err != nil {
		return err
	}

	lastL2Block, err := s.GetLastL2Block(ctx, dbTx)
	if err != nil {
		return err
	}

	// Look for the transaction that matches latest state root in data base
	// in case we already have l2blocks for that batch
	// to just store new transactions
	if lastL2Block.Header().Number.Uint64() == batchNumber {
		stateRoot := lastL2Block.Header().Root

		for i, processedTx := range processedTxs {
			if processedTx.StateRoot == stateRoot {
				foundPosition = i
				break
			}
		}
	}

	foundPosition++

	for i := foundPosition; i < len(processedTxs); i++ {
		processedTx := processedTxs[i]

		lastL2Block, err := s.GetLastL2Block(ctx, dbTx)
		if err != nil {
			return err
		}

		header := &types.Header{
			Number:     new(big.Int).SetUint64(lastL2Block.Number().Uint64() + 1),
			ParentHash: lastL2Block.Hash(),
			Coinbase:   batch.Coinbase,
			Root:       processedTx.StateRoot,
		}

		transactions := []*types.Transaction{}
		transactions = append(transactions, &processedTx.Tx)

		// Create block to be able to calculate its hash
		block := types.NewBlock(header, transactions, []*types.Header{}, []*types.Receipt{}, &trie.StackTrie{})
		block.ReceivedAt = batch.Timestamp

		receipt := generateReceipt(block, processedTx)
		receipts := []*types.Receipt{}
		receipts = append(receipts, receipt)

		// Store L2 block and its transaction
		err = s.PostgresStorage.AddL2Block(ctx, batchNumber, block, receipts, dbTx)
		if err != nil {
			return err
		}
	}
	return nil
}

// CloseBatch is used by the Trusted Sequencer to close batch
func (s *State) CloseBatch(ctx context.Context, batchNum uint64, stateRoot, localExitRoot common.Hash, dbTx pgx.Tx) error {
	// TODO: implement
	return nil
}

// ProcessAndStoreClosedBatch is used by the Synchronizer to a add closed batch into the data base
func (s *State) ProcessAndStoreClosedBatch(ctx context.Context, batch Batch, dbTx pgx.Tx) error {
	processBatchResponse, err := s.ProcessBatch(ctx, batch.BatchNumber, batch.Transactions, dbTx)
	if err != nil {
		return err
	}

	err = s.StoreTransactions(ctx, batch.BatchNumber, processBatchResponse.Responses, dbTx)
	if err != nil {
		return err
	}

	i := len(processBatchResponse.Responses) - 1

	for i >= 0 {
		if !processBatchResponse.Responses[i].UnprocessedTransaction {
			break
		}
	}

	if i < 0 {
		return ErrNoIncludedTransactions
	}

	return s.CloseBatch(ctx, batch.BatchNumber, processBatchResponse.Responses[i].StateRoot, batch.OldLocalExitRoot, dbTx)
}

// GetLastTrustedBatchNumber get last trusted batch number
func (s *State) GetLastTrustedBatchNumber(ctx context.Context) (uint64, error) {
	// TODO: implement
	return 0, nil
}

// GetLastBatch gets latest batch (closed or not) on the data base
func (s *State) GetLastBatch(ctx context.Context, dbTx pgx.Tx) (*Batch, error) {
	batches, err := s.PostgresStorage.GetLastNBatches(ctx, 1, dbTx)
	if err != nil {
		return nil, err
	}
	return batches[0], nil
}

// DebugTransaction re-executes a tx to generate its trace
func (s *State) DebugTransaction(ctx context.Context, transactionHash common.Hash, tracer string) (*runtime.ExecutionResult, error) {
	// TODO: Implement
	return new(runtime.ExecutionResult), nil
}

func (s *State) ParseTheTraceUsingTheTracer(env *fakevm.FakeEVM, trace instrumentation.ExecutorTrace, jsTracer tracers.Tracer) (json.RawMessage, error) {
	var previousDepth int
	var previousOpcode string
	var stateRoot []byte

	contextGas, ok := new(big.Int).SetString(trace.Context.Gas, encoding.Base10)
	if !ok {
		log.Debugf("error while parsing contextGas")
		return nil, ErrParsingExecutorTrace
	}
	value, ok := new(big.Int).SetString(trace.Context.Value, encoding.Base10)
	if !ok {
		log.Debugf("error while parsing value")
		return nil, ErrParsingExecutorTrace
	}

	jsTracer.CaptureTxStart(contextGas.Uint64())
	jsTracer.CaptureStart(env, common.HexToAddress(trace.Context.From), common.HexToAddress(trace.Context.To), trace.Context.Type == "CREATE", common.Hex2Bytes(strings.TrimLeft(trace.Context.Input, "0x")), contextGas.Uint64(), value)

	stack := fakevm.Newstack()
	memory := fakevm.NewMemory()

	bigStateRoot, ok := new(big.Int).SetString(trace.Context.OldStateRoot, 0)
	if !ok {
		log.Debugf("error while parsing context oldStateRoot")
		return nil, ErrParsingExecutorTrace
	}
	stateRoot = bigStateRoot.Bytes()
	env.StateDB.SetStateRoot(stateRoot)

	for i, step := range trace.Steps {
		gas, ok := new(big.Int).SetString(step.Gas, encoding.Base10)
		if !ok {
			log.Debugf("error while parsing step gas")
			return nil, ErrParsingExecutorTrace
		}

		gasCost, ok := new(big.Int).SetString(step.GasCost, encoding.Base10)
		if !ok {
			log.Debugf("error while parsing step gasCost")
			return nil, ErrParsingExecutorTrace
		}

		value, ok := new(big.Int).SetString(step.Contract.Value, encoding.Base10)
		if !ok {
			log.Debugf("error while parsing step value")
			return nil, ErrParsingExecutorTrace
		}

		op, ok := new(big.Int).SetString(step.Op, 0)
		if !ok {
			log.Debugf("error while parsing step op")
			return nil, ErrParsingExecutorTrace
		}

		scope := &fakevm.ScopeContext{
			Contract: vm.NewContract(fakevm.NewAccount(common.HexToAddress(step.Contract.Caller)), fakevm.NewAccount(common.HexToAddress(step.Contract.Address)), value, gas.Uint64()),
			Memory:   memory,
			Stack:    stack,
		}

		codeAddr := common.HexToAddress(step.Contract.Address)
		scope.Contract.CodeAddr = &codeAddr

		opcode := vm.OpCode(op.Uint64()).String()

		if previousOpcode == "CALL" && step.Pc != 0 {
			jsTracer.CaptureExit(common.Hex2Bytes(step.ReturnData), gasCost.Uint64(), fmt.Errorf(step.Error))
		}

		if opcode != "CALL" || trace.Steps[i+1].Pc == 0 {
			if step.Error != "" {
				err := fmt.Errorf(step.Error)
				jsTracer.CaptureFault(step.Pc, vm.OpCode(op.Uint64()), gas.Uint64(), gasCost.Uint64(), scope, step.Depth, err)
			} else {
				jsTracer.CaptureState(step.Pc, vm.OpCode(op.Uint64()), gas.Uint64(), gasCost.Uint64(), scope, common.Hex2Bytes(strings.TrimLeft(step.ReturnData, "0x")), step.Depth, nil)
			}
		}

		if opcode == "CREATE" || opcode == "CREATE2" || opcode == "CALL" || opcode == "CALLCODE" || opcode == "DELEGATECALL" || opcode == "STATICCALL" || opcode == "SELFDESTRUCT" {
			jsTracer.CaptureEnter(vm.OpCode(op.Uint64()), common.HexToAddress(step.Contract.Caller), common.HexToAddress(step.Contract.Address), common.Hex2Bytes(strings.TrimLeft(step.Contract.Input, "0x")), gas.Uint64(), value)
			if step.OpCode == "SELFDESTRUCT" {
				jsTracer.CaptureExit(common.Hex2Bytes(step.ReturnData), gasCost.Uint64(), fmt.Errorf(step.Error))
			}
		}

		// Set Memory
		if len(step.Memory) > 0 {
			memory.Resize(uint64(fakevm.MemoryItemSize*len(step.Memory) + zkEVMReservedMemorySize))
			for offset, memoryContent := range step.Memory {
				memory.Set(uint64((offset*fakevm.MemoryItemSize)+zkEVMReservedMemorySize), uint64(fakevm.MemoryItemSize), common.Hex2Bytes(memoryContent))
			}
		} else {
			memory = fakevm.NewMemory()
		}

		// Set Stack
		stack = fakevm.Newstack()
		for _, stackContent := range step.Stack {
			valueBigInt, ok := new(big.Int).SetString(stackContent, 0)
			if !ok {
				log.Debugf("error while parsing stack valueBigInt")
				return nil, ErrParsingExecutorTrace
			}
			value, _ := uint256.FromBig(valueBigInt)
			stack.Push(value)
		}

		// Returning from a call or create
		if previousDepth > step.Depth {
			jsTracer.CaptureExit(common.Hex2Bytes(step.ReturnData), gasCost.Uint64(), fmt.Errorf(step.Error))
		}

		// Set StateRoot
		bigStateRoot, ok := new(big.Int).SetString(step.StateRoot, 0)
		if !ok {
			log.Debugf("error while parsing step stateRoot")
			return nil, ErrParsingExecutorTrace
		}

		stateRoot = bigStateRoot.Bytes()
		env.StateDB.SetStateRoot(stateRoot)
		previousDepth = step.Depth
		previousOpcode = step.OpCode
	}

	gasUsed, ok := new(big.Int).SetString(trace.Context.GasUsed, encoding.Base10)
	if !ok {
		log.Debugf("error while parsing gasUsed")
		return nil, ErrParsingExecutorTrace
	}

	jsTracer.CaptureTxEnd(gasUsed.Uint64())
	jsTracer.CaptureEnd(common.Hex2Bytes(trace.Context.Output), gasUsed.Uint64(), time.Duration(trace.Context.Time), nil)

	return jsTracer.GetResult()
}

func (s *State) ProcessUnsignedTransaction(ctx context.Context, tx *types.Transaction, senderAddress, sequencerAddress common.Address, blockNumber uint64, dbTx pgx.Tx) *runtime.ExecutionResult {
	panic("not implemented yet")
}

// AddBatchNumberInForcedBatch updates the forced_batch table with the batchNumber.
func (s *State) AddBatchNumberInForcedBatch(ctx context.Context, forceBatchNumber, batchNumber uint64, dbTx pgx.Tx) error {
	return s.PostgresStorage.AddBatchNumberInForcedBatch(ctx, forceBatchNumber, batchNumber, dbTx)
}

// GetTree returns State inner tree
func (s *State) GetTree() *merkletree.StateTree {
	return s.tree
}

// SetGenesis populates state with genesis information
func (s *State) SetGenesis(ctx context.Context, genesis Genesis, dbTx pgx.Tx) error {
	var (
		root    common.Hash
		newRoot []byte
		err     error
	)

	if genesis.Balances != nil {
		for address, balance := range genesis.Balances {
			newRoot, _, err = s.tree.SetBalance(ctx, address, balance, newRoot)
			if err != nil {
				return err
			}
		}
	}

	if genesis.SmartContracts != nil {
		for address, sc := range genesis.SmartContracts {
			newRoot, _, err = s.tree.SetCode(ctx, address, sc, newRoot)
			if err != nil {
				return err
			}
		}
	}

	if len(genesis.Storage) > 0 {
		for address, storage := range genesis.Storage {
			for key, value := range storage {
				newRoot, _, err = s.tree.SetStorageAt(ctx, address, key, value, newRoot)
				if err != nil {
					return err
				}
			}
		}
	}

	if genesis.Nonces != nil {
		for address, nonce := range genesis.Nonces {
			newRoot, _, err = s.tree.SetNonce(ctx, address, nonce, newRoot)
			if err != nil {
				return err
			}
		}
	}

	root.SetBytes(newRoot)

	receivedAt := time.Now()

	// Store Genesis Batch
	batch := Batch{
		BatchNumber:       0,
		Coinbase:          ZeroAddress,
		BatchL2Data:       nil,
		OldStateRoot:      ZeroHash,
		GlobalExitRootNum: big.NewInt(0),
		OldLocalExitRoot:  ZeroHash,
		Timestamp:         receivedAt,
		Transactions:      []types.Transaction{},
		GlobalExitRoot:    ZeroHash,
	}

	err = s.PostgresStorage.StoreBatchHeader(ctx, batch, dbTx)
	if err != nil {
		return err
	}

	// Store L2 Genesis Block
	header := &types.Header{
		Number:     big.NewInt(0),
		ParentHash: ZeroHash,
		Coinbase:   ZeroAddress,
		Root:       root,
	}
	block := types.NewBlock(header, []*types.Transaction{}, []*types.Header{}, []*types.Receipt{}, &trie.StackTrie{})
	block.ReceivedAt = receivedAt

	return s.PostgresStorage.AddL2Block(ctx, batch.BatchNumber, block, []*types.Receipt{}, dbTx)
}
