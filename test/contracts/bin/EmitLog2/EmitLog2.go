// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EmitLog2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EmitLog2MetaData contains all meta data concerning the EmitLog2 contract.
var EmitLog2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"LogA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"LogAB\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"LogABC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"LogABCD\",\"type\":\"event\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516000805160206101d683398151915290600090a160405160019060008051602061019683398151915290600090a260405160029060019060008051602061021683398151915290600090a36003600260016000805160206101f683398151915260405160405180910390a46003600260016000805160206101b683398151915260046040516100a591815260200190565b60405180910390a46002600360046000805160206101b683398151915260016040516100d391815260200190565b60405180910390a46001600260036000805160206101f683398151915260405160405180910390a460405160019060029060008051602061021683398151915290600090a360405160019060008051602061019683398151915290600090a26040516000805160206101d683398151915290600090a1603f806101576000396000f3fe6080604052600080fdfea2646970667358221220311d93a0faf5fe46d5a196e7b40f1aefa02f70a6e32b7a5dcb3b6ab001543eb264736f6c634300080c0033977224b24e70d33f3be87246a29c5636cfc8dd6853e175b54af01ff493ffac62e5562b12d9276c5c987df08afff7b1946f2d869236866ea2285c7e2e95685a645e7df75d54e493185612379c616118a4c9ac802de621b010c96f74d22df4b30a966018f1afaee50c6bcf5eb4ae089eeb650bd1deb473395d69dd307ef2e689b7bb6e4da744abea70325874159d52c1ad3e57babfae7c329a948e7dcb274deb09",
}

// EmitLog2ABI is the input ABI used to generate the binding from.
// Deprecated: Use EmitLog2MetaData.ABI instead.
var EmitLog2ABI = EmitLog2MetaData.ABI

// EmitLog2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EmitLog2MetaData.Bin instead.
var EmitLog2Bin = EmitLog2MetaData.Bin

// DeployEmitLog2 deploys a new Ethereum contract, binding an instance of EmitLog2 to it.
func DeployEmitLog2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EmitLog2, error) {
	parsed, err := EmitLog2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EmitLog2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EmitLog2{EmitLog2Caller: EmitLog2Caller{contract: contract}, EmitLog2Transactor: EmitLog2Transactor{contract: contract}, EmitLog2Filterer: EmitLog2Filterer{contract: contract}}, nil
}

// EmitLog2 is an auto generated Go binding around an Ethereum contract.
type EmitLog2 struct {
	EmitLog2Caller     // Read-only binding to the contract
	EmitLog2Transactor // Write-only binding to the contract
	EmitLog2Filterer   // Log filterer for contract events
}

// EmitLog2Caller is an auto generated read-only Go binding around an Ethereum contract.
type EmitLog2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitLog2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EmitLog2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitLog2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmitLog2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitLog2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmitLog2Session struct {
	Contract     *EmitLog2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmitLog2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmitLog2CallerSession struct {
	Contract *EmitLog2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EmitLog2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmitLog2TransactorSession struct {
	Contract     *EmitLog2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EmitLog2Raw is an auto generated low-level Go binding around an Ethereum contract.
type EmitLog2Raw struct {
	Contract *EmitLog2 // Generic contract binding to access the raw methods on
}

// EmitLog2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmitLog2CallerRaw struct {
	Contract *EmitLog2Caller // Generic read-only contract binding to access the raw methods on
}

// EmitLog2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmitLog2TransactorRaw struct {
	Contract *EmitLog2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEmitLog2 creates a new instance of EmitLog2, bound to a specific deployed contract.
func NewEmitLog2(address common.Address, backend bind.ContractBackend) (*EmitLog2, error) {
	contract, err := bindEmitLog2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmitLog2{EmitLog2Caller: EmitLog2Caller{contract: contract}, EmitLog2Transactor: EmitLog2Transactor{contract: contract}, EmitLog2Filterer: EmitLog2Filterer{contract: contract}}, nil
}

// NewEmitLog2Caller creates a new read-only instance of EmitLog2, bound to a specific deployed contract.
func NewEmitLog2Caller(address common.Address, caller bind.ContractCaller) (*EmitLog2Caller, error) {
	contract, err := bindEmitLog2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmitLog2Caller{contract: contract}, nil
}

// NewEmitLog2Transactor creates a new write-only instance of EmitLog2, bound to a specific deployed contract.
func NewEmitLog2Transactor(address common.Address, transactor bind.ContractTransactor) (*EmitLog2Transactor, error) {
	contract, err := bindEmitLog2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmitLog2Transactor{contract: contract}, nil
}

// NewEmitLog2Filterer creates a new log filterer instance of EmitLog2, bound to a specific deployed contract.
func NewEmitLog2Filterer(address common.Address, filterer bind.ContractFilterer) (*EmitLog2Filterer, error) {
	contract, err := bindEmitLog2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmitLog2Filterer{contract: contract}, nil
}

// bindEmitLog2 binds a generic wrapper to an already deployed contract.
func bindEmitLog2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmitLog2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmitLog2 *EmitLog2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmitLog2.Contract.EmitLog2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmitLog2 *EmitLog2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmitLog2.Contract.EmitLog2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmitLog2 *EmitLog2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmitLog2.Contract.EmitLog2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmitLog2 *EmitLog2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmitLog2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmitLog2 *EmitLog2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmitLog2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmitLog2 *EmitLog2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmitLog2.Contract.contract.Transact(opts, method, params...)
}

// EmitLog2LogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the EmitLog2 contract.
type EmitLog2LogIterator struct {
	Event *EmitLog2Log // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitLog2LogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitLog2Log)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitLog2Log)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitLog2LogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitLog2LogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitLog2Log represents a Log event raised by the EmitLog2 contract.
type EmitLog2Log struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x5e7df75d54e493185612379c616118a4c9ac802de621b010c96f74d22df4b30a.
//
// Solidity: event Log()
func (_EmitLog2 *EmitLog2Filterer) FilterLog(opts *bind.FilterOpts) (*EmitLog2LogIterator, error) {

	logs, sub, err := _EmitLog2.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &EmitLog2LogIterator{contract: _EmitLog2.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x5e7df75d54e493185612379c616118a4c9ac802de621b010c96f74d22df4b30a.
//
// Solidity: event Log()
func (_EmitLog2 *EmitLog2Filterer) WatchLog(opts *bind.WatchOpts, sink chan<- *EmitLog2Log) (event.Subscription, error) {

	logs, sub, err := _EmitLog2.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitLog2Log)
				if err := _EmitLog2.contract.UnpackLog(event, "Log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLog is a log parse operation binding the contract event 0x5e7df75d54e493185612379c616118a4c9ac802de621b010c96f74d22df4b30a.
//
// Solidity: event Log()
func (_EmitLog2 *EmitLog2Filterer) ParseLog(log types.Log) (*EmitLog2Log, error) {
	event := new(EmitLog2Log)
	if err := _EmitLog2.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmitLog2LogAIterator is returned from FilterLogA and is used to iterate over the raw logs and unpacked data for LogA events raised by the EmitLog2 contract.
type EmitLog2LogAIterator struct {
	Event *EmitLog2LogA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitLog2LogAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitLog2LogA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitLog2LogA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitLog2LogAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitLog2LogAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitLog2LogA represents a LogA event raised by the EmitLog2 contract.
type EmitLog2LogA struct {
	A   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogA is a free log retrieval operation binding the contract event 0x977224b24e70d33f3be87246a29c5636cfc8dd6853e175b54af01ff493ffac62.
//
// Solidity: event LogA(uint256 indexed a)
func (_EmitLog2 *EmitLog2Filterer) FilterLogA(opts *bind.FilterOpts, a []*big.Int) (*EmitLog2LogAIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _EmitLog2.contract.FilterLogs(opts, "LogA", aRule)
	if err != nil {
		return nil, err
	}
	return &EmitLog2LogAIterator{contract: _EmitLog2.contract, event: "LogA", logs: logs, sub: sub}, nil
}

// WatchLogA is a free log subscription operation binding the contract event 0x977224b24e70d33f3be87246a29c5636cfc8dd6853e175b54af01ff493ffac62.
//
// Solidity: event LogA(uint256 indexed a)
func (_EmitLog2 *EmitLog2Filterer) WatchLogA(opts *bind.WatchOpts, sink chan<- *EmitLog2LogA, a []*big.Int) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _EmitLog2.contract.WatchLogs(opts, "LogA", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitLog2LogA)
				if err := _EmitLog2.contract.UnpackLog(event, "LogA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogA is a log parse operation binding the contract event 0x977224b24e70d33f3be87246a29c5636cfc8dd6853e175b54af01ff493ffac62.
//
// Solidity: event LogA(uint256 indexed a)
func (_EmitLog2 *EmitLog2Filterer) ParseLogA(log types.Log) (*EmitLog2LogA, error) {
	event := new(EmitLog2LogA)
	if err := _EmitLog2.contract.UnpackLog(event, "LogA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmitLog2LogABIterator is returned from FilterLogAB and is used to iterate over the raw logs and unpacked data for LogAB events raised by the EmitLog2 contract.
type EmitLog2LogABIterator struct {
	Event *EmitLog2LogAB // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitLog2LogABIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitLog2LogAB)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitLog2LogAB)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitLog2LogABIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitLog2LogABIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitLog2LogAB represents a LogAB event raised by the EmitLog2 contract.
type EmitLog2LogAB struct {
	A   *big.Int
	B   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogAB is a free log retrieval operation binding the contract event 0xbb6e4da744abea70325874159d52c1ad3e57babfae7c329a948e7dcb274deb09.
//
// Solidity: event LogAB(uint256 indexed a, uint256 indexed b)
func (_EmitLog2 *EmitLog2Filterer) FilterLogAB(opts *bind.FilterOpts, a []*big.Int, b []*big.Int) (*EmitLog2LogABIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var bRule []interface{}
	for _, bItem := range b {
		bRule = append(bRule, bItem)
	}

	logs, sub, err := _EmitLog2.contract.FilterLogs(opts, "LogAB", aRule, bRule)
	if err != nil {
		return nil, err
	}
	return &EmitLog2LogABIterator{contract: _EmitLog2.contract, event: "LogAB", logs: logs, sub: sub}, nil
}

// WatchLogAB is a free log subscription operation binding the contract event 0xbb6e4da744abea70325874159d52c1ad3e57babfae7c329a948e7dcb274deb09.
//
// Solidity: event LogAB(uint256 indexed a, uint256 indexed b)
func (_EmitLog2 *EmitLog2Filterer) WatchLogAB(opts *bind.WatchOpts, sink chan<- *EmitLog2LogAB, a []*big.Int, b []*big.Int) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var bRule []interface{}
	for _, bItem := range b {
		bRule = append(bRule, bItem)
	}

	logs, sub, err := _EmitLog2.contract.WatchLogs(opts, "LogAB", aRule, bRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitLog2LogAB)
				if err := _EmitLog2.contract.UnpackLog(event, "LogAB", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAB is a log parse operation binding the contract event 0xbb6e4da744abea70325874159d52c1ad3e57babfae7c329a948e7dcb274deb09.
//
// Solidity: event LogAB(uint256 indexed a, uint256 indexed b)
func (_EmitLog2 *EmitLog2Filterer) ParseLogAB(log types.Log) (*EmitLog2LogAB, error) {
	event := new(EmitLog2LogAB)
	if err := _EmitLog2.contract.UnpackLog(event, "LogAB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmitLog2LogABCIterator is returned from FilterLogABC and is used to iterate over the raw logs and unpacked data for LogABC events raised by the EmitLog2 contract.
type EmitLog2LogABCIterator struct {
	Event *EmitLog2LogABC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitLog2LogABCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitLog2LogABC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitLog2LogABC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitLog2LogABCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitLog2LogABCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitLog2LogABC represents a LogABC event raised by the EmitLog2 contract.
type EmitLog2LogABC struct {
	A   *big.Int
	B   *big.Int
	C   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogABC is a free log retrieval operation binding the contract event 0x966018f1afaee50c6bcf5eb4ae089eeb650bd1deb473395d69dd307ef2e689b7.
//
// Solidity: event LogABC(uint256 indexed a, uint256 indexed b, uint256 indexed c)
func (_EmitLog2 *EmitLog2Filterer) FilterLogABC(opts *bind.FilterOpts, a []*big.Int, b []*big.Int, c []*big.Int) (*EmitLog2LogABCIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var bRule []interface{}
	for _, bItem := range b {
		bRule = append(bRule, bItem)
	}
	var cRule []interface{}
	for _, cItem := range c {
		cRule = append(cRule, cItem)
	}

	logs, sub, err := _EmitLog2.contract.FilterLogs(opts, "LogABC", aRule, bRule, cRule)
	if err != nil {
		return nil, err
	}
	return &EmitLog2LogABCIterator{contract: _EmitLog2.contract, event: "LogABC", logs: logs, sub: sub}, nil
}

// WatchLogABC is a free log subscription operation binding the contract event 0x966018f1afaee50c6bcf5eb4ae089eeb650bd1deb473395d69dd307ef2e689b7.
//
// Solidity: event LogABC(uint256 indexed a, uint256 indexed b, uint256 indexed c)
func (_EmitLog2 *EmitLog2Filterer) WatchLogABC(opts *bind.WatchOpts, sink chan<- *EmitLog2LogABC, a []*big.Int, b []*big.Int, c []*big.Int) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var bRule []interface{}
	for _, bItem := range b {
		bRule = append(bRule, bItem)
	}
	var cRule []interface{}
	for _, cItem := range c {
		cRule = append(cRule, cItem)
	}

	logs, sub, err := _EmitLog2.contract.WatchLogs(opts, "LogABC", aRule, bRule, cRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitLog2LogABC)
				if err := _EmitLog2.contract.UnpackLog(event, "LogABC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogABC is a log parse operation binding the contract event 0x966018f1afaee50c6bcf5eb4ae089eeb650bd1deb473395d69dd307ef2e689b7.
//
// Solidity: event LogABC(uint256 indexed a, uint256 indexed b, uint256 indexed c)
func (_EmitLog2 *EmitLog2Filterer) ParseLogABC(log types.Log) (*EmitLog2LogABC, error) {
	event := new(EmitLog2LogABC)
	if err := _EmitLog2.contract.UnpackLog(event, "LogABC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmitLog2LogABCDIterator is returned from FilterLogABCD and is used to iterate over the raw logs and unpacked data for LogABCD events raised by the EmitLog2 contract.
type EmitLog2LogABCDIterator struct {
	Event *EmitLog2LogABCD // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitLog2LogABCDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitLog2LogABCD)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitLog2LogABCD)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitLog2LogABCDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitLog2LogABCDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitLog2LogABCD represents a LogABCD event raised by the EmitLog2 contract.
type EmitLog2LogABCD struct {
	A   *big.Int
	B   *big.Int
	C   *big.Int
	D   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogABCD is a free log retrieval operation binding the contract event 0xe5562b12d9276c5c987df08afff7b1946f2d869236866ea2285c7e2e95685a64.
//
// Solidity: event LogABCD(uint256 indexed a, uint256 indexed b, uint256 indexed c, uint256 d)
func (_EmitLog2 *EmitLog2Filterer) FilterLogABCD(opts *bind.FilterOpts, a []*big.Int, b []*big.Int, c []*big.Int) (*EmitLog2LogABCDIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var bRule []interface{}
	for _, bItem := range b {
		bRule = append(bRule, bItem)
	}
	var cRule []interface{}
	for _, cItem := range c {
		cRule = append(cRule, cItem)
	}

	logs, sub, err := _EmitLog2.contract.FilterLogs(opts, "LogABCD", aRule, bRule, cRule)
	if err != nil {
		return nil, err
	}
	return &EmitLog2LogABCDIterator{contract: _EmitLog2.contract, event: "LogABCD", logs: logs, sub: sub}, nil
}

// WatchLogABCD is a free log subscription operation binding the contract event 0xe5562b12d9276c5c987df08afff7b1946f2d869236866ea2285c7e2e95685a64.
//
// Solidity: event LogABCD(uint256 indexed a, uint256 indexed b, uint256 indexed c, uint256 d)
func (_EmitLog2 *EmitLog2Filterer) WatchLogABCD(opts *bind.WatchOpts, sink chan<- *EmitLog2LogABCD, a []*big.Int, b []*big.Int, c []*big.Int) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}
	var bRule []interface{}
	for _, bItem := range b {
		bRule = append(bRule, bItem)
	}
	var cRule []interface{}
	for _, cItem := range c {
		cRule = append(cRule, cItem)
	}

	logs, sub, err := _EmitLog2.contract.WatchLogs(opts, "LogABCD", aRule, bRule, cRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitLog2LogABCD)
				if err := _EmitLog2.contract.UnpackLog(event, "LogABCD", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogABCD is a log parse operation binding the contract event 0xe5562b12d9276c5c987df08afff7b1946f2d869236866ea2285c7e2e95685a64.
//
// Solidity: event LogABCD(uint256 indexed a, uint256 indexed b, uint256 indexed c, uint256 d)
func (_EmitLog2 *EmitLog2Filterer) ParseLogABCD(log types.Log) (*EmitLog2LogABCD, error) {
	event := new(EmitLog2LogABCD)
	if err := _EmitLog2.contract.UnpackLog(event, "LogABCD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
