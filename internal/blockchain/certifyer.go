// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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
	_ = abi.ConvertType
)

// CertifyerMetaData contains all meta data concerning the Certifyer contract.
var CertifyerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"certificates\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerCertificate\",\"inputs\":[{\"name\":\"datahash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateCertificate\",\"inputs\":[{\"name\":\"datahash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CertificateCreated\",\"inputs\":[{\"name\":\"datahash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// CertifyerABI is the input ABI used to generate the binding from.
// Deprecated: Use CertifyerMetaData.ABI instead.
var CertifyerABI = CertifyerMetaData.ABI

// Certifyer is an auto generated Go binding around an Ethereum contract.
type Certifyer struct {
	CertifyerCaller     // Read-only binding to the contract
	CertifyerTransactor // Write-only binding to the contract
	CertifyerFilterer   // Log filterer for contract events
}

// CertifyerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertifyerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertifyerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertifyerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertifyerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertifyerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertifyerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertifyerSession struct {
	Contract     *Certifyer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertifyerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertifyerCallerSession struct {
	Contract *CertifyerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CertifyerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertifyerTransactorSession struct {
	Contract     *CertifyerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CertifyerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertifyerRaw struct {
	Contract *Certifyer // Generic contract binding to access the raw methods on
}

// CertifyerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertifyerCallerRaw struct {
	Contract *CertifyerCaller // Generic read-only contract binding to access the raw methods on
}

// CertifyerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertifyerTransactorRaw struct {
	Contract *CertifyerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertifyer creates a new instance of Certifyer, bound to a specific deployed contract.
func NewCertifyer(address common.Address, backend bind.ContractBackend) (*Certifyer, error) {
	contract, err := bindCertifyer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Certifyer{CertifyerCaller: CertifyerCaller{contract: contract}, CertifyerTransactor: CertifyerTransactor{contract: contract}, CertifyerFilterer: CertifyerFilterer{contract: contract}}, nil
}

// NewCertifyerCaller creates a new read-only instance of Certifyer, bound to a specific deployed contract.
func NewCertifyerCaller(address common.Address, caller bind.ContractCaller) (*CertifyerCaller, error) {
	contract, err := bindCertifyer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertifyerCaller{contract: contract}, nil
}

// NewCertifyerTransactor creates a new write-only instance of Certifyer, bound to a specific deployed contract.
func NewCertifyerTransactor(address common.Address, transactor bind.ContractTransactor) (*CertifyerTransactor, error) {
	contract, err := bindCertifyer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertifyerTransactor{contract: contract}, nil
}

// NewCertifyerFilterer creates a new log filterer instance of Certifyer, bound to a specific deployed contract.
func NewCertifyerFilterer(address common.Address, filterer bind.ContractFilterer) (*CertifyerFilterer, error) {
	contract, err := bindCertifyer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertifyerFilterer{contract: contract}, nil
}

// bindCertifyer binds a generic wrapper to an already deployed contract.
func bindCertifyer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertifyerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certifyer *CertifyerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certifyer.Contract.CertifyerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certifyer *CertifyerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certifyer.Contract.CertifyerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certifyer *CertifyerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certifyer.Contract.CertifyerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certifyer *CertifyerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certifyer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certifyer *CertifyerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certifyer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certifyer *CertifyerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certifyer.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Certifyer *CertifyerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Certifyer.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Certifyer *CertifyerSession) Admin() (common.Address, error) {
	return _Certifyer.Contract.Admin(&_Certifyer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Certifyer *CertifyerCallerSession) Admin() (common.Address, error) {
	return _Certifyer.Contract.Admin(&_Certifyer.CallOpts)
}

// Certificates is a free data retrieval call binding the contract method 0x742f0688.
//
// Solidity: function certificates(bytes32 ) view returns(bool)
func (_Certifyer *CertifyerCaller) Certificates(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Certifyer.contract.Call(opts, &out, "certificates", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Certificates is a free data retrieval call binding the contract method 0x742f0688.
//
// Solidity: function certificates(bytes32 ) view returns(bool)
func (_Certifyer *CertifyerSession) Certificates(arg0 [32]byte) (bool, error) {
	return _Certifyer.Contract.Certificates(&_Certifyer.CallOpts, arg0)
}

// Certificates is a free data retrieval call binding the contract method 0x742f0688.
//
// Solidity: function certificates(bytes32 ) view returns(bool)
func (_Certifyer *CertifyerCallerSession) Certificates(arg0 [32]byte) (bool, error) {
	return _Certifyer.Contract.Certificates(&_Certifyer.CallOpts, arg0)
}

// ValidateCertificate is a free data retrieval call binding the contract method 0xe4f50ff9.
//
// Solidity: function validateCertificate(bytes32 datahash) view returns(bool)
func (_Certifyer *CertifyerCaller) ValidateCertificate(opts *bind.CallOpts, datahash [32]byte) (bool, error) {
	var out []interface{}
	err := _Certifyer.contract.Call(opts, &out, "validateCertificate", datahash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateCertificate is a free data retrieval call binding the contract method 0xe4f50ff9.
//
// Solidity: function validateCertificate(bytes32 datahash) view returns(bool)
func (_Certifyer *CertifyerSession) ValidateCertificate(datahash [32]byte) (bool, error) {
	return _Certifyer.Contract.ValidateCertificate(&_Certifyer.CallOpts, datahash)
}

// ValidateCertificate is a free data retrieval call binding the contract method 0xe4f50ff9.
//
// Solidity: function validateCertificate(bytes32 datahash) view returns(bool)
func (_Certifyer *CertifyerCallerSession) ValidateCertificate(datahash [32]byte) (bool, error) {
	return _Certifyer.Contract.ValidateCertificate(&_Certifyer.CallOpts, datahash)
}

// RegisterCertificate is a paid mutator transaction binding the contract method 0xf101dce8.
//
// Solidity: function registerCertificate(bytes32 datahash) returns()
func (_Certifyer *CertifyerTransactor) RegisterCertificate(opts *bind.TransactOpts, datahash [32]byte) (*types.Transaction, error) {
	return _Certifyer.contract.Transact(opts, "registerCertificate", datahash)
}

// RegisterCertificate is a paid mutator transaction binding the contract method 0xf101dce8.
//
// Solidity: function registerCertificate(bytes32 datahash) returns()
func (_Certifyer *CertifyerSession) RegisterCertificate(datahash [32]byte) (*types.Transaction, error) {
	return _Certifyer.Contract.RegisterCertificate(&_Certifyer.TransactOpts, datahash)
}

// RegisterCertificate is a paid mutator transaction binding the contract method 0xf101dce8.
//
// Solidity: function registerCertificate(bytes32 datahash) returns()
func (_Certifyer *CertifyerTransactorSession) RegisterCertificate(datahash [32]byte) (*types.Transaction, error) {
	return _Certifyer.Contract.RegisterCertificate(&_Certifyer.TransactOpts, datahash)
}

// CertifyerCertificateCreatedIterator is returned from FilterCertificateCreated and is used to iterate over the raw logs and unpacked data for CertificateCreated events raised by the Certifyer contract.
type CertifyerCertificateCreatedIterator struct {
	Event *CertifyerCertificateCreated // Event containing the contract specifics and raw log

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
func (it *CertifyerCertificateCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertifyerCertificateCreated)
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
		it.Event = new(CertifyerCertificateCreated)
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
func (it *CertifyerCertificateCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertifyerCertificateCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertifyerCertificateCreated represents a CertificateCreated event raised by the Certifyer contract.
type CertifyerCertificateCreated struct {
	Datahash  [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCertificateCreated is a free log retrieval operation binding the contract event 0xbd5a05444054f4f4d80536c2d1b53f08c81a842c88390bbddbe576f68730bbfa.
//
// Solidity: event CertificateCreated(bytes32 indexed datahash, uint256 timestamp)
func (_Certifyer *CertifyerFilterer) FilterCertificateCreated(opts *bind.FilterOpts, datahash [][32]byte) (*CertifyerCertificateCreatedIterator, error) {

	var datahashRule []interface{}
	for _, datahashItem := range datahash {
		datahashRule = append(datahashRule, datahashItem)
	}

	logs, sub, err := _Certifyer.contract.FilterLogs(opts, "CertificateCreated", datahashRule)
	if err != nil {
		return nil, err
	}
	return &CertifyerCertificateCreatedIterator{contract: _Certifyer.contract, event: "CertificateCreated", logs: logs, sub: sub}, nil
}

// WatchCertificateCreated is a free log subscription operation binding the contract event 0xbd5a05444054f4f4d80536c2d1b53f08c81a842c88390bbddbe576f68730bbfa.
//
// Solidity: event CertificateCreated(bytes32 indexed datahash, uint256 timestamp)
func (_Certifyer *CertifyerFilterer) WatchCertificateCreated(opts *bind.WatchOpts, sink chan<- *CertifyerCertificateCreated, datahash [][32]byte) (event.Subscription, error) {

	var datahashRule []interface{}
	for _, datahashItem := range datahash {
		datahashRule = append(datahashRule, datahashItem)
	}

	logs, sub, err := _Certifyer.contract.WatchLogs(opts, "CertificateCreated", datahashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertifyerCertificateCreated)
				if err := _Certifyer.contract.UnpackLog(event, "CertificateCreated", log); err != nil {
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

// ParseCertificateCreated is a log parse operation binding the contract event 0xbd5a05444054f4f4d80536c2d1b53f08c81a842c88390bbddbe576f68730bbfa.
//
// Solidity: event CertificateCreated(bytes32 indexed datahash, uint256 timestamp)
func (_Certifyer *CertifyerFilterer) ParseCertificateCreated(log types.Log) (*CertifyerCertificateCreated, error) {
	event := new(CertifyerCertificateCreated)
	if err := _Certifyer.contract.UnpackLog(event, "CertificateCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
