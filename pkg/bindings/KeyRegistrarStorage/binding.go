// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KeyRegistrarStorage

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// KeyRegistrarStorageMetaData contains all meta data concerning the KeyRegistrarStorage contract.
var KeyRegistrarStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configureOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encodeBN254KeyData\",\"inputs\":[{\"name\":\"g1Point\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"g2Point\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getBN254Key\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"g1Point\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"g2Point\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBN254KeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"keyData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getECDSAAddress\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getECDSAKey\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getECDSAKeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"keyAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyHash\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromSigningKey\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"keyData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCurveType\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isKeyGloballyRegistered\",\"inputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRegistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AggregateBN254KeyUpdated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"newAggregateKey\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeyDeregistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeyRegistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"curveType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetConfigured\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"curveType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIKeyRegistrarTypes.CurveType\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ConfigurationAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurveType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeyFormat\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKeypair\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyNotFound\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetNotConfigured\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorStillSlashable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPubkey\",\"inputs\":[]}]",
}

// KeyRegistrarStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyRegistrarStorageMetaData.ABI instead.
var KeyRegistrarStorageABI = KeyRegistrarStorageMetaData.ABI

// KeyRegistrarStorage is an auto generated Go binding around an Ethereum contract.
type KeyRegistrarStorage struct {
	KeyRegistrarStorageCaller     // Read-only binding to the contract
	KeyRegistrarStorageTransactor // Write-only binding to the contract
	KeyRegistrarStorageFilterer   // Log filterer for contract events
}

// KeyRegistrarStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyRegistrarStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistrarStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyRegistrarStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistrarStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyRegistrarStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistrarStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyRegistrarStorageSession struct {
	Contract     *KeyRegistrarStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KeyRegistrarStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyRegistrarStorageCallerSession struct {
	Contract *KeyRegistrarStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// KeyRegistrarStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyRegistrarStorageTransactorSession struct {
	Contract     *KeyRegistrarStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// KeyRegistrarStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyRegistrarStorageRaw struct {
	Contract *KeyRegistrarStorage // Generic contract binding to access the raw methods on
}

// KeyRegistrarStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyRegistrarStorageCallerRaw struct {
	Contract *KeyRegistrarStorageCaller // Generic read-only contract binding to access the raw methods on
}

// KeyRegistrarStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyRegistrarStorageTransactorRaw struct {
	Contract *KeyRegistrarStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyRegistrarStorage creates a new instance of KeyRegistrarStorage, bound to a specific deployed contract.
func NewKeyRegistrarStorage(address common.Address, backend bind.ContractBackend) (*KeyRegistrarStorage, error) {
	contract, err := bindKeyRegistrarStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorage{KeyRegistrarStorageCaller: KeyRegistrarStorageCaller{contract: contract}, KeyRegistrarStorageTransactor: KeyRegistrarStorageTransactor{contract: contract}, KeyRegistrarStorageFilterer: KeyRegistrarStorageFilterer{contract: contract}}, nil
}

// NewKeyRegistrarStorageCaller creates a new read-only instance of KeyRegistrarStorage, bound to a specific deployed contract.
func NewKeyRegistrarStorageCaller(address common.Address, caller bind.ContractCaller) (*KeyRegistrarStorageCaller, error) {
	contract, err := bindKeyRegistrarStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorageCaller{contract: contract}, nil
}

// NewKeyRegistrarStorageTransactor creates a new write-only instance of KeyRegistrarStorage, bound to a specific deployed contract.
func NewKeyRegistrarStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyRegistrarStorageTransactor, error) {
	contract, err := bindKeyRegistrarStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorageTransactor{contract: contract}, nil
}

// NewKeyRegistrarStorageFilterer creates a new log filterer instance of KeyRegistrarStorage, bound to a specific deployed contract.
func NewKeyRegistrarStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyRegistrarStorageFilterer, error) {
	contract, err := bindKeyRegistrarStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorageFilterer{contract: contract}, nil
}

// bindKeyRegistrarStorage binds a generic wrapper to an already deployed contract.
func bindKeyRegistrarStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyRegistrarStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyRegistrarStorage *KeyRegistrarStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyRegistrarStorage.Contract.KeyRegistrarStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyRegistrarStorage *KeyRegistrarStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.KeyRegistrarStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyRegistrarStorage *KeyRegistrarStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.KeyRegistrarStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyRegistrarStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) AllocationManager() (common.Address, error) {
	return _KeyRegistrarStorage.Contract.AllocationManager(&_KeyRegistrarStorage.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) AllocationManager() (common.Address, error) {
	return _KeyRegistrarStorage.Contract.AllocationManager(&_KeyRegistrarStorage.CallOpts)
}

// EncodeBN254KeyData is a free data retrieval call binding the contract method 0x50435add.
//
// Solidity: function encodeBN254KeyData((uint256,uint256) g1Point, (uint256[2],uint256[2]) g2Point) pure returns(bytes)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) EncodeBN254KeyData(opts *bind.CallOpts, g1Point BN254G1Point, g2Point BN254G2Point) ([]byte, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "encodeBN254KeyData", g1Point, g2Point)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeBN254KeyData is a free data retrieval call binding the contract method 0x50435add.
//
// Solidity: function encodeBN254KeyData((uint256,uint256) g1Point, (uint256[2],uint256[2]) g2Point) pure returns(bytes)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) EncodeBN254KeyData(g1Point BN254G1Point, g2Point BN254G2Point) ([]byte, error) {
	return _KeyRegistrarStorage.Contract.EncodeBN254KeyData(&_KeyRegistrarStorage.CallOpts, g1Point, g2Point)
}

// EncodeBN254KeyData is a free data retrieval call binding the contract method 0x50435add.
//
// Solidity: function encodeBN254KeyData((uint256,uint256) g1Point, (uint256[2],uint256[2]) g2Point) pure returns(bytes)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) EncodeBN254KeyData(g1Point BN254G1Point, g2Point BN254G2Point) ([]byte, error) {
	return _KeyRegistrarStorage.Contract.EncodeBN254KeyData(&_KeyRegistrarStorage.CallOpts, g1Point, g2Point)
}

// GetBN254Key is a free data retrieval call binding the contract method 0x9a43e3fb.
//
// Solidity: function getBN254Key((address,uint32) operatorSet, address operator) view returns((uint256,uint256) g1Point, (uint256[2],uint256[2]) g2Point)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetBN254Key(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (struct {
	G1Point BN254G1Point
	G2Point BN254G2Point
}, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getBN254Key", operatorSet, operator)

	outstruct := new(struct {
		G1Point BN254G1Point
		G2Point BN254G2Point
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.G1Point = *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	outstruct.G2Point = *abi.ConvertType(out[1], new(BN254G2Point)).(*BN254G2Point)

	return *outstruct, err

}

// GetBN254Key is a free data retrieval call binding the contract method 0x9a43e3fb.
//
// Solidity: function getBN254Key((address,uint32) operatorSet, address operator) view returns((uint256,uint256) g1Point, (uint256[2],uint256[2]) g2Point)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetBN254Key(operatorSet OperatorSet, operator common.Address) (struct {
	G1Point BN254G1Point
	G2Point BN254G2Point
}, error) {
	return _KeyRegistrarStorage.Contract.GetBN254Key(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetBN254Key is a free data retrieval call binding the contract method 0x9a43e3fb.
//
// Solidity: function getBN254Key((address,uint32) operatorSet, address operator) view returns((uint256,uint256) g1Point, (uint256[2],uint256[2]) g2Point)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetBN254Key(operatorSet OperatorSet, operator common.Address) (struct {
	G1Point BN254G1Point
	G2Point BN254G2Point
}, error) {
	return _KeyRegistrarStorage.Contract.GetBN254Key(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetBN254KeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x7690e395.
//
// Solidity: function getBN254KeyRegistrationMessageHash(address operator, (address,uint32) operatorSet, bytes keyData) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetBN254KeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, keyData []byte) ([32]byte, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getBN254KeyRegistrationMessageHash", operator, operatorSet, keyData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBN254KeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x7690e395.
//
// Solidity: function getBN254KeyRegistrationMessageHash(address operator, (address,uint32) operatorSet, bytes keyData) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetBN254KeyRegistrationMessageHash(operator common.Address, operatorSet OperatorSet, keyData []byte) ([32]byte, error) {
	return _KeyRegistrarStorage.Contract.GetBN254KeyRegistrationMessageHash(&_KeyRegistrarStorage.CallOpts, operator, operatorSet, keyData)
}

// GetBN254KeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x7690e395.
//
// Solidity: function getBN254KeyRegistrationMessageHash(address operator, (address,uint32) operatorSet, bytes keyData) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetBN254KeyRegistrationMessageHash(operator common.Address, operatorSet OperatorSet, keyData []byte) ([32]byte, error) {
	return _KeyRegistrarStorage.Contract.GetBN254KeyRegistrationMessageHash(&_KeyRegistrarStorage.CallOpts, operator, operatorSet, keyData)
}

// GetECDSAAddress is a free data retrieval call binding the contract method 0x3b32a7bd.
//
// Solidity: function getECDSAAddress((address,uint32) operatorSet, address operator) view returns(address)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetECDSAAddress(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getECDSAAddress", operatorSet, operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetECDSAAddress is a free data retrieval call binding the contract method 0x3b32a7bd.
//
// Solidity: function getECDSAAddress((address,uint32) operatorSet, address operator) view returns(address)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetECDSAAddress(operatorSet OperatorSet, operator common.Address) (common.Address, error) {
	return _KeyRegistrarStorage.Contract.GetECDSAAddress(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetECDSAAddress is a free data retrieval call binding the contract method 0x3b32a7bd.
//
// Solidity: function getECDSAAddress((address,uint32) operatorSet, address operator) view returns(address)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetECDSAAddress(operatorSet OperatorSet, operator common.Address) (common.Address, error) {
	return _KeyRegistrarStorage.Contract.GetECDSAAddress(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetECDSAKey is a free data retrieval call binding the contract method 0xaa165c30.
//
// Solidity: function getECDSAKey((address,uint32) operatorSet, address operator) view returns(bytes)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetECDSAKey(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) ([]byte, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getECDSAKey", operatorSet, operator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetECDSAKey is a free data retrieval call binding the contract method 0xaa165c30.
//
// Solidity: function getECDSAKey((address,uint32) operatorSet, address operator) view returns(bytes)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetECDSAKey(operatorSet OperatorSet, operator common.Address) ([]byte, error) {
	return _KeyRegistrarStorage.Contract.GetECDSAKey(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetECDSAKey is a free data retrieval call binding the contract method 0xaa165c30.
//
// Solidity: function getECDSAKey((address,uint32) operatorSet, address operator) view returns(bytes)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetECDSAKey(operatorSet OperatorSet, operator common.Address) ([]byte, error) {
	return _KeyRegistrarStorage.Contract.GetECDSAKey(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetECDSAKeyRegistrationMessageHash is a free data retrieval call binding the contract method 0xd9f12db2.
//
// Solidity: function getECDSAKeyRegistrationMessageHash(address operator, (address,uint32) operatorSet, address keyAddress) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetECDSAKeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, keyAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getECDSAKeyRegistrationMessageHash", operator, operatorSet, keyAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetECDSAKeyRegistrationMessageHash is a free data retrieval call binding the contract method 0xd9f12db2.
//
// Solidity: function getECDSAKeyRegistrationMessageHash(address operator, (address,uint32) operatorSet, address keyAddress) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetECDSAKeyRegistrationMessageHash(operator common.Address, operatorSet OperatorSet, keyAddress common.Address) ([32]byte, error) {
	return _KeyRegistrarStorage.Contract.GetECDSAKeyRegistrationMessageHash(&_KeyRegistrarStorage.CallOpts, operator, operatorSet, keyAddress)
}

// GetECDSAKeyRegistrationMessageHash is a free data retrieval call binding the contract method 0xd9f12db2.
//
// Solidity: function getECDSAKeyRegistrationMessageHash(address operator, (address,uint32) operatorSet, address keyAddress) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetECDSAKeyRegistrationMessageHash(operator common.Address, operatorSet OperatorSet, keyAddress common.Address) ([32]byte, error) {
	return _KeyRegistrarStorage.Contract.GetECDSAKeyRegistrationMessageHash(&_KeyRegistrarStorage.CallOpts, operator, operatorSet, keyAddress)
}

// GetKeyHash is a free data retrieval call binding the contract method 0xea194e2e.
//
// Solidity: function getKeyHash((address,uint32) operatorSet, address operator) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetKeyHash(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getKeyHash", operatorSet, operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeyHash is a free data retrieval call binding the contract method 0xea194e2e.
//
// Solidity: function getKeyHash((address,uint32) operatorSet, address operator) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetKeyHash(operatorSet OperatorSet, operator common.Address) ([32]byte, error) {
	return _KeyRegistrarStorage.Contract.GetKeyHash(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetKeyHash is a free data retrieval call binding the contract method 0xea194e2e.
//
// Solidity: function getKeyHash((address,uint32) operatorSet, address operator) view returns(bytes32)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetKeyHash(operatorSet OperatorSet, operator common.Address) ([32]byte, error) {
	return _KeyRegistrarStorage.Contract.GetKeyHash(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// GetOperatorFromSigningKey is a free data retrieval call binding the contract method 0x8256909c.
//
// Solidity: function getOperatorFromSigningKey((address,uint32) operatorSet, bytes keyData) view returns(address, bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetOperatorFromSigningKey(opts *bind.CallOpts, operatorSet OperatorSet, keyData []byte) (common.Address, bool, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getOperatorFromSigningKey", operatorSet, keyData)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetOperatorFromSigningKey is a free data retrieval call binding the contract method 0x8256909c.
//
// Solidity: function getOperatorFromSigningKey((address,uint32) operatorSet, bytes keyData) view returns(address, bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetOperatorFromSigningKey(operatorSet OperatorSet, keyData []byte) (common.Address, bool, error) {
	return _KeyRegistrarStorage.Contract.GetOperatorFromSigningKey(&_KeyRegistrarStorage.CallOpts, operatorSet, keyData)
}

// GetOperatorFromSigningKey is a free data retrieval call binding the contract method 0x8256909c.
//
// Solidity: function getOperatorFromSigningKey((address,uint32) operatorSet, bytes keyData) view returns(address, bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetOperatorFromSigningKey(operatorSet OperatorSet, keyData []byte) (common.Address, bool, error) {
	return _KeyRegistrarStorage.Contract.GetOperatorFromSigningKey(&_KeyRegistrarStorage.CallOpts, operatorSet, keyData)
}

// GetOperatorSetCurveType is a free data retrieval call binding the contract method 0x7cffe48c.
//
// Solidity: function getOperatorSetCurveType((address,uint32) operatorSet) view returns(uint8)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) GetOperatorSetCurveType(opts *bind.CallOpts, operatorSet OperatorSet) (uint8, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "getOperatorSetCurveType", operatorSet)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorSetCurveType is a free data retrieval call binding the contract method 0x7cffe48c.
//
// Solidity: function getOperatorSetCurveType((address,uint32) operatorSet) view returns(uint8)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) GetOperatorSetCurveType(operatorSet OperatorSet) (uint8, error) {
	return _KeyRegistrarStorage.Contract.GetOperatorSetCurveType(&_KeyRegistrarStorage.CallOpts, operatorSet)
}

// GetOperatorSetCurveType is a free data retrieval call binding the contract method 0x7cffe48c.
//
// Solidity: function getOperatorSetCurveType((address,uint32) operatorSet) view returns(uint8)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) GetOperatorSetCurveType(operatorSet OperatorSet) (uint8, error) {
	return _KeyRegistrarStorage.Contract.GetOperatorSetCurveType(&_KeyRegistrarStorage.CallOpts, operatorSet)
}

// IsKeyGloballyRegistered is a free data retrieval call binding the contract method 0xdab42d7e.
//
// Solidity: function isKeyGloballyRegistered(bytes32 keyHash) view returns(bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) IsKeyGloballyRegistered(opts *bind.CallOpts, keyHash [32]byte) (bool, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "isKeyGloballyRegistered", keyHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeyGloballyRegistered is a free data retrieval call binding the contract method 0xdab42d7e.
//
// Solidity: function isKeyGloballyRegistered(bytes32 keyHash) view returns(bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) IsKeyGloballyRegistered(keyHash [32]byte) (bool, error) {
	return _KeyRegistrarStorage.Contract.IsKeyGloballyRegistered(&_KeyRegistrarStorage.CallOpts, keyHash)
}

// IsKeyGloballyRegistered is a free data retrieval call binding the contract method 0xdab42d7e.
//
// Solidity: function isKeyGloballyRegistered(bytes32 keyHash) view returns(bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) IsKeyGloballyRegistered(keyHash [32]byte) (bool, error) {
	return _KeyRegistrarStorage.Contract.IsKeyGloballyRegistered(&_KeyRegistrarStorage.CallOpts, keyHash)
}

// IsRegistered is a free data retrieval call binding the contract method 0xbd30a0b9.
//
// Solidity: function isRegistered((address,uint32) operatorSet, address operator) view returns(bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) IsRegistered(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (bool, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "isRegistered", operatorSet, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xbd30a0b9.
//
// Solidity: function isRegistered((address,uint32) operatorSet, address operator) view returns(bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) IsRegistered(operatorSet OperatorSet, operator common.Address) (bool, error) {
	return _KeyRegistrarStorage.Contract.IsRegistered(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// IsRegistered is a free data retrieval call binding the contract method 0xbd30a0b9.
//
// Solidity: function isRegistered((address,uint32) operatorSet, address operator) view returns(bool)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) IsRegistered(operatorSet OperatorSet, operator common.Address) (bool, error) {
	return _KeyRegistrarStorage.Contract.IsRegistered(&_KeyRegistrarStorage.CallOpts, operatorSet, operator)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KeyRegistrarStorage *KeyRegistrarStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KeyRegistrarStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) Version() (string, error) {
	return _KeyRegistrarStorage.Contract.Version(&_KeyRegistrarStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KeyRegistrarStorage *KeyRegistrarStorageCallerSession) Version() (string, error) {
	return _KeyRegistrarStorage.Contract.Version(&_KeyRegistrarStorage.CallOpts)
}

// ConfigureOperatorSet is a paid mutator transaction binding the contract method 0xea0d8149.
//
// Solidity: function configureOperatorSet((address,uint32) operatorSet, uint8 curveType) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactor) ConfigureOperatorSet(opts *bind.TransactOpts, operatorSet OperatorSet, curveType uint8) (*types.Transaction, error) {
	return _KeyRegistrarStorage.contract.Transact(opts, "configureOperatorSet", operatorSet, curveType)
}

// ConfigureOperatorSet is a paid mutator transaction binding the contract method 0xea0d8149.
//
// Solidity: function configureOperatorSet((address,uint32) operatorSet, uint8 curveType) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) ConfigureOperatorSet(operatorSet OperatorSet, curveType uint8) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.ConfigureOperatorSet(&_KeyRegistrarStorage.TransactOpts, operatorSet, curveType)
}

// ConfigureOperatorSet is a paid mutator transaction binding the contract method 0xea0d8149.
//
// Solidity: function configureOperatorSet((address,uint32) operatorSet, uint8 curveType) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactorSession) ConfigureOperatorSet(operatorSet OperatorSet, curveType uint8) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.ConfigureOperatorSet(&_KeyRegistrarStorage.TransactOpts, operatorSet, curveType)
}

// DeregisterKey is a paid mutator transaction binding the contract method 0x87ab86f4.
//
// Solidity: function deregisterKey(address operator, (address,uint32) operatorSet) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactor) DeregisterKey(opts *bind.TransactOpts, operator common.Address, operatorSet OperatorSet) (*types.Transaction, error) {
	return _KeyRegistrarStorage.contract.Transact(opts, "deregisterKey", operator, operatorSet)
}

// DeregisterKey is a paid mutator transaction binding the contract method 0x87ab86f4.
//
// Solidity: function deregisterKey(address operator, (address,uint32) operatorSet) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) DeregisterKey(operator common.Address, operatorSet OperatorSet) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.DeregisterKey(&_KeyRegistrarStorage.TransactOpts, operator, operatorSet)
}

// DeregisterKey is a paid mutator transaction binding the contract method 0x87ab86f4.
//
// Solidity: function deregisterKey(address operator, (address,uint32) operatorSet) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactorSession) DeregisterKey(operator common.Address, operatorSet OperatorSet) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.DeregisterKey(&_KeyRegistrarStorage.TransactOpts, operator, operatorSet)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xd40cda16.
//
// Solidity: function registerKey(address operator, (address,uint32) operatorSet, bytes pubkey, bytes signature) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactor) RegisterKey(opts *bind.TransactOpts, operator common.Address, operatorSet OperatorSet, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _KeyRegistrarStorage.contract.Transact(opts, "registerKey", operator, operatorSet, pubkey, signature)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xd40cda16.
//
// Solidity: function registerKey(address operator, (address,uint32) operatorSet, bytes pubkey, bytes signature) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageSession) RegisterKey(operator common.Address, operatorSet OperatorSet, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.RegisterKey(&_KeyRegistrarStorage.TransactOpts, operator, operatorSet, pubkey, signature)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xd40cda16.
//
// Solidity: function registerKey(address operator, (address,uint32) operatorSet, bytes pubkey, bytes signature) returns()
func (_KeyRegistrarStorage *KeyRegistrarStorageTransactorSession) RegisterKey(operator common.Address, operatorSet OperatorSet, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _KeyRegistrarStorage.Contract.RegisterKey(&_KeyRegistrarStorage.TransactOpts, operator, operatorSet, pubkey, signature)
}

// KeyRegistrarStorageAggregateBN254KeyUpdatedIterator is returned from FilterAggregateBN254KeyUpdated and is used to iterate over the raw logs and unpacked data for AggregateBN254KeyUpdated events raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageAggregateBN254KeyUpdatedIterator struct {
	Event *KeyRegistrarStorageAggregateBN254KeyUpdated // Event containing the contract specifics and raw log

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
func (it *KeyRegistrarStorageAggregateBN254KeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrarStorageAggregateBN254KeyUpdated)
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
		it.Event = new(KeyRegistrarStorageAggregateBN254KeyUpdated)
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
func (it *KeyRegistrarStorageAggregateBN254KeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrarStorageAggregateBN254KeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrarStorageAggregateBN254KeyUpdated represents a AggregateBN254KeyUpdated event raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageAggregateBN254KeyUpdated struct {
	OperatorSet     OperatorSet
	NewAggregateKey BN254G1Point
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAggregateBN254KeyUpdated is a free log retrieval operation binding the contract event 0xdfa2f59e55747ba641fbdff4eb78577de8789d605920d5be4a74ee3a6470d1d1.
//
// Solidity: event AggregateBN254KeyUpdated((address,uint32) operatorSet, (uint256,uint256) newAggregateKey)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) FilterAggregateBN254KeyUpdated(opts *bind.FilterOpts) (*KeyRegistrarStorageAggregateBN254KeyUpdatedIterator, error) {

	logs, sub, err := _KeyRegistrarStorage.contract.FilterLogs(opts, "AggregateBN254KeyUpdated")
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorageAggregateBN254KeyUpdatedIterator{contract: _KeyRegistrarStorage.contract, event: "AggregateBN254KeyUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregateBN254KeyUpdated is a free log subscription operation binding the contract event 0xdfa2f59e55747ba641fbdff4eb78577de8789d605920d5be4a74ee3a6470d1d1.
//
// Solidity: event AggregateBN254KeyUpdated((address,uint32) operatorSet, (uint256,uint256) newAggregateKey)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) WatchAggregateBN254KeyUpdated(opts *bind.WatchOpts, sink chan<- *KeyRegistrarStorageAggregateBN254KeyUpdated) (event.Subscription, error) {

	logs, sub, err := _KeyRegistrarStorage.contract.WatchLogs(opts, "AggregateBN254KeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrarStorageAggregateBN254KeyUpdated)
				if err := _KeyRegistrarStorage.contract.UnpackLog(event, "AggregateBN254KeyUpdated", log); err != nil {
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

// ParseAggregateBN254KeyUpdated is a log parse operation binding the contract event 0xdfa2f59e55747ba641fbdff4eb78577de8789d605920d5be4a74ee3a6470d1d1.
//
// Solidity: event AggregateBN254KeyUpdated((address,uint32) operatorSet, (uint256,uint256) newAggregateKey)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) ParseAggregateBN254KeyUpdated(log types.Log) (*KeyRegistrarStorageAggregateBN254KeyUpdated, error) {
	event := new(KeyRegistrarStorageAggregateBN254KeyUpdated)
	if err := _KeyRegistrarStorage.contract.UnpackLog(event, "AggregateBN254KeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrarStorageKeyDeregisteredIterator is returned from FilterKeyDeregistered and is used to iterate over the raw logs and unpacked data for KeyDeregistered events raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageKeyDeregisteredIterator struct {
	Event *KeyRegistrarStorageKeyDeregistered // Event containing the contract specifics and raw log

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
func (it *KeyRegistrarStorageKeyDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrarStorageKeyDeregistered)
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
		it.Event = new(KeyRegistrarStorageKeyDeregistered)
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
func (it *KeyRegistrarStorageKeyDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrarStorageKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrarStorageKeyDeregistered represents a KeyDeregistered event raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageKeyDeregistered struct {
	OperatorSet OperatorSet
	Operator    common.Address
	CurveType   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKeyDeregistered is a free log retrieval operation binding the contract event 0x28d3c3cee49478ec6fd219cfd685cd15cd01d95cabf69b4b7b57f9eaa3eb6442.
//
// Solidity: event KeyDeregistered((address,uint32) operatorSet, address indexed operator, uint8 curveType)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) FilterKeyDeregistered(opts *bind.FilterOpts, operator []common.Address) (*KeyRegistrarStorageKeyDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KeyRegistrarStorage.contract.FilterLogs(opts, "KeyDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorageKeyDeregisteredIterator{contract: _KeyRegistrarStorage.contract, event: "KeyDeregistered", logs: logs, sub: sub}, nil
}

// WatchKeyDeregistered is a free log subscription operation binding the contract event 0x28d3c3cee49478ec6fd219cfd685cd15cd01d95cabf69b4b7b57f9eaa3eb6442.
//
// Solidity: event KeyDeregistered((address,uint32) operatorSet, address indexed operator, uint8 curveType)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) WatchKeyDeregistered(opts *bind.WatchOpts, sink chan<- *KeyRegistrarStorageKeyDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KeyRegistrarStorage.contract.WatchLogs(opts, "KeyDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrarStorageKeyDeregistered)
				if err := _KeyRegistrarStorage.contract.UnpackLog(event, "KeyDeregistered", log); err != nil {
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

// ParseKeyDeregistered is a log parse operation binding the contract event 0x28d3c3cee49478ec6fd219cfd685cd15cd01d95cabf69b4b7b57f9eaa3eb6442.
//
// Solidity: event KeyDeregistered((address,uint32) operatorSet, address indexed operator, uint8 curveType)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) ParseKeyDeregistered(log types.Log) (*KeyRegistrarStorageKeyDeregistered, error) {
	event := new(KeyRegistrarStorageKeyDeregistered)
	if err := _KeyRegistrarStorage.contract.UnpackLog(event, "KeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrarStorageKeyRegisteredIterator is returned from FilterKeyRegistered and is used to iterate over the raw logs and unpacked data for KeyRegistered events raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageKeyRegisteredIterator struct {
	Event *KeyRegistrarStorageKeyRegistered // Event containing the contract specifics and raw log

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
func (it *KeyRegistrarStorageKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrarStorageKeyRegistered)
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
		it.Event = new(KeyRegistrarStorageKeyRegistered)
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
func (it *KeyRegistrarStorageKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrarStorageKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrarStorageKeyRegistered represents a KeyRegistered event raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageKeyRegistered struct {
	OperatorSet OperatorSet
	Operator    common.Address
	CurveType   uint8
	Pubkey      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKeyRegistered is a free log retrieval operation binding the contract event 0x1201ce0c5e577111bce91e907fd99cb183da5edc1e3fb650ca40769e4e9176dd.
//
// Solidity: event KeyRegistered((address,uint32) operatorSet, address indexed operator, uint8 curveType, bytes pubkey)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) FilterKeyRegistered(opts *bind.FilterOpts, operator []common.Address) (*KeyRegistrarStorageKeyRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KeyRegistrarStorage.contract.FilterLogs(opts, "KeyRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorageKeyRegisteredIterator{contract: _KeyRegistrarStorage.contract, event: "KeyRegistered", logs: logs, sub: sub}, nil
}

// WatchKeyRegistered is a free log subscription operation binding the contract event 0x1201ce0c5e577111bce91e907fd99cb183da5edc1e3fb650ca40769e4e9176dd.
//
// Solidity: event KeyRegistered((address,uint32) operatorSet, address indexed operator, uint8 curveType, bytes pubkey)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) WatchKeyRegistered(opts *bind.WatchOpts, sink chan<- *KeyRegistrarStorageKeyRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KeyRegistrarStorage.contract.WatchLogs(opts, "KeyRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrarStorageKeyRegistered)
				if err := _KeyRegistrarStorage.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
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

// ParseKeyRegistered is a log parse operation binding the contract event 0x1201ce0c5e577111bce91e907fd99cb183da5edc1e3fb650ca40769e4e9176dd.
//
// Solidity: event KeyRegistered((address,uint32) operatorSet, address indexed operator, uint8 curveType, bytes pubkey)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) ParseKeyRegistered(log types.Log) (*KeyRegistrarStorageKeyRegistered, error) {
	event := new(KeyRegistrarStorageKeyRegistered)
	if err := _KeyRegistrarStorage.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistrarStorageOperatorSetConfiguredIterator is returned from FilterOperatorSetConfigured and is used to iterate over the raw logs and unpacked data for OperatorSetConfigured events raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageOperatorSetConfiguredIterator struct {
	Event *KeyRegistrarStorageOperatorSetConfigured // Event containing the contract specifics and raw log

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
func (it *KeyRegistrarStorageOperatorSetConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistrarStorageOperatorSetConfigured)
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
		it.Event = new(KeyRegistrarStorageOperatorSetConfigured)
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
func (it *KeyRegistrarStorageOperatorSetConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistrarStorageOperatorSetConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistrarStorageOperatorSetConfigured represents a OperatorSetConfigured event raised by the KeyRegistrarStorage contract.
type KeyRegistrarStorageOperatorSetConfigured struct {
	OperatorSet OperatorSet
	CurveType   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetConfigured is a free log retrieval operation binding the contract event 0xb2266cb118e57095fcdbedb24dabd9fc9f5127e2dbedf62ce6ee71696fb8b6e7.
//
// Solidity: event OperatorSetConfigured((address,uint32) operatorSet, uint8 curveType)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) FilterOperatorSetConfigured(opts *bind.FilterOpts) (*KeyRegistrarStorageOperatorSetConfiguredIterator, error) {

	logs, sub, err := _KeyRegistrarStorage.contract.FilterLogs(opts, "OperatorSetConfigured")
	if err != nil {
		return nil, err
	}
	return &KeyRegistrarStorageOperatorSetConfiguredIterator{contract: _KeyRegistrarStorage.contract, event: "OperatorSetConfigured", logs: logs, sub: sub}, nil
}

// WatchOperatorSetConfigured is a free log subscription operation binding the contract event 0xb2266cb118e57095fcdbedb24dabd9fc9f5127e2dbedf62ce6ee71696fb8b6e7.
//
// Solidity: event OperatorSetConfigured((address,uint32) operatorSet, uint8 curveType)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) WatchOperatorSetConfigured(opts *bind.WatchOpts, sink chan<- *KeyRegistrarStorageOperatorSetConfigured) (event.Subscription, error) {

	logs, sub, err := _KeyRegistrarStorage.contract.WatchLogs(opts, "OperatorSetConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistrarStorageOperatorSetConfigured)
				if err := _KeyRegistrarStorage.contract.UnpackLog(event, "OperatorSetConfigured", log); err != nil {
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

// ParseOperatorSetConfigured is a log parse operation binding the contract event 0xb2266cb118e57095fcdbedb24dabd9fc9f5127e2dbedf62ce6ee71696fb8b6e7.
//
// Solidity: event OperatorSetConfigured((address,uint32) operatorSet, uint8 curveType)
func (_KeyRegistrarStorage *KeyRegistrarStorageFilterer) ParseOperatorSetConfigured(log types.Log) (*KeyRegistrarStorageOperatorSetConfigured, error) {
	event := new(KeyRegistrarStorageOperatorSetConfigured)
	if err := _KeyRegistrarStorage.contract.UnpackLog(event, "OperatorSetConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
