// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PauserRegistry

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

// PauserRegistryMetaData contains all meta data concerning the PauserRegistry contract.
var PauserRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_pausers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_unpauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPauser\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setIsPauser\",\"inputs\":[{\"name\":\"newPauser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"canPause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnpauser\",\"inputs\":[{\"name\":\"newUnpauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PauserStatusChanged\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"canPause\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnpauserChanged\",\"inputs\":[{\"name\":\"previousUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610b4d380380610b4d83398181016040528101906100319190610440565b5f5f90505b82518110156100775761006a8382815181106100555761005461049a565b5b6020026020010151600161008e60201b60201c565b8080600101915050610036565b506100878161018360201b60201c565b505061053e565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036100f3576040517f7363217600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b915282826040516101779291906104f0565b60405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036101e8576040517f7363217600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e889260015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260405161023a929190610517565b60405180910390a18060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102e08261029a565b810181811067ffffffffffffffff821117156102ff576102fe6102aa565b5b80604052505050565b5f610311610285565b905061031d82826102d7565b919050565b5f67ffffffffffffffff82111561033c5761033b6102aa565b5b602082029050602081019050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61037a82610351565b9050919050565b61038a81610370565b8114610394575f5ffd5b50565b5f815190506103a581610381565b92915050565b5f6103bd6103b884610322565b610308565b905080838252602082019050602084028301858111156103e0576103df61034d565b5b835b8181101561040957806103f58882610397565b8452602084019350506020810190506103e2565b5050509392505050565b5f82601f83011261042757610426610296565b5b81516104378482602086016103ab565b91505092915050565b5f5f604083850312156104565761045561028e565b5b5f83015167ffffffffffffffff81111561047357610472610292565b5b61047f85828601610413565b925050602061049085828601610397565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b6104d081610370565b82525050565b5f8115159050919050565b6104ea816104d6565b82525050565b5f6040820190506105035f8301856104c7565b61051060208301846104e1565b9392505050565b5f60408201905061052a5f8301856104c7565b61053760208301846104c7565b9392505050565b6106028061054b5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806346fbf68e1461004e578063856852061461007e578063ce5484281461009a578063eab66d7a146100b6575b5f5ffd5b61006860048036038101906100639190610490565b6100d4565b60405161007591906104d5565b60405180910390f35b61009860048036038101906100939190610518565b6100f0565b005b6100b460048036038101906100af9190610490565b610184565b005b6100be610216565b6040516100cb9190610565565b60405180910390f35b5f602052805f5260405f205f915054906101000a900460ff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610176576040517f794821ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610180828261023b565b5050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461020a576040517f794821ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61021381610330565b50565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036102a0576040517f7363217600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152828260405161032492919061057e565b60405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610395576040517f7363217600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e889260015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826040516103e79291906105a5565b60405180910390a18060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61045f82610436565b9050919050565b61046f81610455565b8114610479575f5ffd5b50565b5f8135905061048a81610466565b92915050565b5f602082840312156104a5576104a4610432565b5b5f6104b28482850161047c565b91505092915050565b5f8115159050919050565b6104cf816104bb565b82525050565b5f6020820190506104e85f8301846104c6565b92915050565b6104f7816104bb565b8114610501575f5ffd5b50565b5f81359050610512816104ee565b92915050565b5f5f6040838503121561052e5761052d610432565b5b5f61053b8582860161047c565b925050602061054c85828601610504565b9150509250929050565b61055f81610455565b82525050565b5f6020820190506105785f830184610556565b92915050565b5f6040820190506105915f830185610556565b61059e60208301846104c6565b9392505050565b5f6040820190506105b85f830185610556565b6105c56020830184610556565b939250505056fea26469706673582212208eb48f096bf58e6151bd2afec5787523e61c66bb95654d16b49aeed459f3377e64736f6c634300081b0033",
}

// PauserRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PauserRegistryMetaData.ABI instead.
var PauserRegistryABI = PauserRegistryMetaData.ABI

// PauserRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PauserRegistryMetaData.Bin instead.
var PauserRegistryBin = PauserRegistryMetaData.Bin

// DeployPauserRegistry deploys a new Ethereum contract, binding an instance of PauserRegistry to it.
func DeployPauserRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _pausers []common.Address, _unpauser common.Address) (common.Address, *types.Transaction, *PauserRegistry, error) {
	parsed, err := PauserRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PauserRegistryBin), backend, _pausers, _unpauser)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PauserRegistry{PauserRegistryCaller: PauserRegistryCaller{contract: contract}, PauserRegistryTransactor: PauserRegistryTransactor{contract: contract}, PauserRegistryFilterer: PauserRegistryFilterer{contract: contract}}, nil
}

// PauserRegistry is an auto generated Go binding around an Ethereum contract.
type PauserRegistry struct {
	PauserRegistryCaller     // Read-only binding to the contract
	PauserRegistryTransactor // Write-only binding to the contract
	PauserRegistryFilterer   // Log filterer for contract events
}

// PauserRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauserRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauserRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauserRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauserRegistrySession struct {
	Contract     *PauserRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauserRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauserRegistryCallerSession struct {
	Contract *PauserRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PauserRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauserRegistryTransactorSession struct {
	Contract     *PauserRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PauserRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauserRegistryRaw struct {
	Contract *PauserRegistry // Generic contract binding to access the raw methods on
}

// PauserRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauserRegistryCallerRaw struct {
	Contract *PauserRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PauserRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauserRegistryTransactorRaw struct {
	Contract *PauserRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauserRegistry creates a new instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistry(address common.Address, backend bind.ContractBackend) (*PauserRegistry, error) {
	contract, err := bindPauserRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauserRegistry{PauserRegistryCaller: PauserRegistryCaller{contract: contract}, PauserRegistryTransactor: PauserRegistryTransactor{contract: contract}, PauserRegistryFilterer: PauserRegistryFilterer{contract: contract}}, nil
}

// NewPauserRegistryCaller creates a new read-only instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryCaller(address common.Address, caller bind.ContractCaller) (*PauserRegistryCaller, error) {
	contract, err := bindPauserRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryCaller{contract: contract}, nil
}

// NewPauserRegistryTransactor creates a new write-only instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PauserRegistryTransactor, error) {
	contract, err := bindPauserRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryTransactor{contract: contract}, nil
}

// NewPauserRegistryFilterer creates a new log filterer instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PauserRegistryFilterer, error) {
	contract, err := bindPauserRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryFilterer{contract: contract}, nil
}

// bindPauserRegistry binds a generic wrapper to an already deployed contract.
func bindPauserRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PauserRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRegistry *PauserRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauserRegistry.Contract.PauserRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRegistry *PauserRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRegistry.Contract.PauserRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRegistry *PauserRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRegistry.Contract.PauserRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRegistry *PauserRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauserRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRegistry *PauserRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRegistry *PauserRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistryCaller) IsPauser(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PauserRegistry.contract.Call(opts, &out, "isPauser", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistrySession) IsPauser(arg0 common.Address) (bool, error) {
	return _PauserRegistry.Contract.IsPauser(&_PauserRegistry.CallOpts, arg0)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistryCallerSession) IsPauser(arg0 common.Address) (bool, error) {
	return _PauserRegistry.Contract.IsPauser(&_PauserRegistry.CallOpts, arg0)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistryCaller) Unpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PauserRegistry.contract.Call(opts, &out, "unpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistrySession) Unpauser() (common.Address, error) {
	return _PauserRegistry.Contract.Unpauser(&_PauserRegistry.CallOpts)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistryCallerSession) Unpauser() (common.Address, error) {
	return _PauserRegistry.Contract.Unpauser(&_PauserRegistry.CallOpts)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistryTransactor) SetIsPauser(opts *bind.TransactOpts, newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.contract.Transact(opts, "setIsPauser", newPauser, canPause)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistrySession) SetIsPauser(newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetIsPauser(&_PauserRegistry.TransactOpts, newPauser, canPause)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistryTransactorSession) SetIsPauser(newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetIsPauser(&_PauserRegistry.TransactOpts, newPauser, canPause)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistryTransactor) SetUnpauser(opts *bind.TransactOpts, newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.contract.Transact(opts, "setUnpauser", newUnpauser)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistrySession) SetUnpauser(newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetUnpauser(&_PauserRegistry.TransactOpts, newUnpauser)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistryTransactorSession) SetUnpauser(newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetUnpauser(&_PauserRegistry.TransactOpts, newUnpauser)
}

// PauserRegistryPauserStatusChangedIterator is returned from FilterPauserStatusChanged and is used to iterate over the raw logs and unpacked data for PauserStatusChanged events raised by the PauserRegistry contract.
type PauserRegistryPauserStatusChangedIterator struct {
	Event *PauserRegistryPauserStatusChanged // Event containing the contract specifics and raw log

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
func (it *PauserRegistryPauserStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRegistryPauserStatusChanged)
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
		it.Event = new(PauserRegistryPauserStatusChanged)
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
func (it *PauserRegistryPauserStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRegistryPauserStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRegistryPauserStatusChanged represents a PauserStatusChanged event raised by the PauserRegistry contract.
type PauserRegistryPauserStatusChanged struct {
	Pauser   common.Address
	CanPause bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauserStatusChanged is a free log retrieval operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) FilterPauserStatusChanged(opts *bind.FilterOpts) (*PauserRegistryPauserStatusChangedIterator, error) {

	logs, sub, err := _PauserRegistry.contract.FilterLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return &PauserRegistryPauserStatusChangedIterator{contract: _PauserRegistry.contract, event: "PauserStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPauserStatusChanged is a free log subscription operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) WatchPauserStatusChanged(opts *bind.WatchOpts, sink chan<- *PauserRegistryPauserStatusChanged) (event.Subscription, error) {

	logs, sub, err := _PauserRegistry.contract.WatchLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRegistryPauserStatusChanged)
				if err := _PauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
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

// ParsePauserStatusChanged is a log parse operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) ParsePauserStatusChanged(log types.Log) (*PauserRegistryPauserStatusChanged, error) {
	event := new(PauserRegistryPauserStatusChanged)
	if err := _PauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauserRegistryUnpauserChangedIterator is returned from FilterUnpauserChanged and is used to iterate over the raw logs and unpacked data for UnpauserChanged events raised by the PauserRegistry contract.
type PauserRegistryUnpauserChangedIterator struct {
	Event *PauserRegistryUnpauserChanged // Event containing the contract specifics and raw log

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
func (it *PauserRegistryUnpauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRegistryUnpauserChanged)
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
		it.Event = new(PauserRegistryUnpauserChanged)
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
func (it *PauserRegistryUnpauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRegistryUnpauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRegistryUnpauserChanged represents a UnpauserChanged event raised by the PauserRegistry contract.
type PauserRegistryUnpauserChanged struct {
	PreviousUnpauser common.Address
	NewUnpauser      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnpauserChanged is a free log retrieval operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) FilterUnpauserChanged(opts *bind.FilterOpts) (*PauserRegistryUnpauserChangedIterator, error) {

	logs, sub, err := _PauserRegistry.contract.FilterLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return &PauserRegistryUnpauserChangedIterator{contract: _PauserRegistry.contract, event: "UnpauserChanged", logs: logs, sub: sub}, nil
}

// WatchUnpauserChanged is a free log subscription operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) WatchUnpauserChanged(opts *bind.WatchOpts, sink chan<- *PauserRegistryUnpauserChanged) (event.Subscription, error) {

	logs, sub, err := _PauserRegistry.contract.WatchLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRegistryUnpauserChanged)
				if err := _PauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
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

// ParseUnpauserChanged is a log parse operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) ParseUnpauserChanged(log types.Log) (*PauserRegistryUnpauserChanged, error) {
	event := new(PauserRegistryUnpauserChanged)
	if err := _PauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
