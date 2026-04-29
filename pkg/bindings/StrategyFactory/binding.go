// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyFactory

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

// IDurationVaultStrategyTypesVaultConfig is an auto generated low-level Go binding around an user-defined struct.
type IDurationVaultStrategyTypesVaultConfig struct {
	UnderlyingToken             common.Address
	VaultAdmin                  common.Address
	Arbitrator                  common.Address
	Duration                    uint32
	MaxPerDeposit               *big.Int
	StakeCap                    *big.Int
	MetadataURI                 string
	OperatorSet                 OperatorSet
	OperatorSetRegistrationData []byte
	DelegationApprover          common.Address
	OperatorMetadataURI         string
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// StrategyFactoryMetaData contains all meta data concerning the StrategyFactory contract.
var StrategyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_strategyBeacon\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"},{\"name\":\"_durationVaultBeacon\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"},{\"name\":\"_EIGEN\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_bEIGEN\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bEIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blacklistTokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployDurationVaultStrategy\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIDurationVaultStrategyTypes.VaultConfig\",\"components\":[{\"name\":\"underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"vaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stakeCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorSetRegistrationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"newVault\",\"type\":\"address\",\"internalType\":\"contractIDurationVaultStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployNewStrategy\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"newStrategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployedStrategies\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"durationVaultBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"durationVaultsByToken\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDurationVaultStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDurationVaults\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIDurationVaultStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBlacklisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistStrategies\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DurationVaultDeployed\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIDurationVaultStrategy\"},{\"name\":\"underlyingToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"vaultAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"duration\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"maxPerDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stakeCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"operatorSetAVS\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategySetForToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBlacklisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyBlacklisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlacklistedToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProhibitedDurationVaultToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyExists\",\"inputs\":[]}]",
	Bin: "0x610140604052348015610010575f5ffd5b506040516120d03803806120d083398101604081905261002f9161016d565b846001600160a01b038116610057576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805286811660a05284811660c05283811660e0528281166101005281166101205261008c610097565b5050505050506101f0565b603354610100900460ff16156101035760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60335460ff90811614610154576033805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461016a575f5ffd5b50565b5f5f5f5f5f5f60c08789031215610182575f5ffd5b865161018d81610156565b602088015190965061019e81610156565b60408801519095506101af81610156565b60608801519094506101c081610156565b60808801519093506101d181610156565b60a08801519092506101e281610156565b809150509295509295509295565b60805160a05160c05160e0516101005161012051611e4c6102845f395f81816101db0152610c4e01525f81816103960152610c1101525f81816102c90152610ca201525f818161033601526107c101525f81816101b40152818161064c015281816108740152818161098401528181610dd60152610e6c01525f818161029101528181610eb801526110c50152611e4c5ff3fe608060405234801561000f575f5ffd5b5060043610610153575f3560e01c80638da5cb5b116100bf578063f2fde38b11610079578063f2fde38b14610358578063fabc1cbc1461036b578063fb5153031461037e578063fdc371ce14610391578063fe38b32d146103b8578063fe575a87146103cb575f5ffd5b80638da5cb5b146102b35780639e4f2dcd146102c4578063a5dae840146102eb578063b768ebc91461030b578063cd6dc6871461031e578063f0062d9a14610331575f5ffd5b8063595c6a6711610110578063595c6a67146102255780635ac86ab71461022d5780635c975abb146102605780636b9b622914610271578063715018a614610284578063886f11951461028c575f5ffd5b8063136439dd1461015757806323103c411461016c5780633101d5311461017f57806339b70e38146101af5780633f4da4c6146101d6578063581dfd65146101fd575b5f5ffd5b61016a61016536600461123c565b6103ed565b005b61016a61017a36600461129b565b610427565b61019261018d3660046112fe565b6106b5565b6040516001600160a01b0390911681526020015b60405180910390f35b6101927f000000000000000000000000000000000000000000000000000000000000000081565b6101927f000000000000000000000000000000000000000000000000000000000000000081565b61019261020b366004611328565b60016020525f90815260409020546001600160a01b031681565b61016a6106e9565b61025061023b36600461134a565b609954600160ff9092169190911b9081161490565b60405190151581526020016101a6565b6099546040519081526020016101a6565b61019261027f366004611328565b6106fd565b61016a6108e1565b6101927f000000000000000000000000000000000000000000000000000000000000000081565b6066546001600160a01b0316610192565b6101927f000000000000000000000000000000000000000000000000000000000000000081565b6102fe6102f9366004611328565b6108f2565b6040516101a6919061136a565b61016a61031936600461129b565b610965565b61016a61032c3660046112fe565b6109ec565b6101927f000000000000000000000000000000000000000000000000000000000000000081565b61016a610366366004611328565b610b0e565b61016a61037936600461123c565b610b87565b61019261038c3660046113b5565b610bf4565b6101927f000000000000000000000000000000000000000000000000000000000000000081565b61016a6103c636600461129b565b610e4d565b6102506103d9366004611328565b60026020525f908152604090205460ff1681565b6103f5610ea3565b609954818116811461041a5760405163c61dca5d60e01b815260040160405180910390fd5b61042382610f46565b5050565b61042f610f83565b5f8167ffffffffffffffff811115610449576104496113ed565b604051908082528060200260200182016040528015610472578160200160208202803683370190505b5090505f805b8381101561062b5760025f86868481811061049557610495611401565b90506020020160208101906104aa9190611328565b6001600160a01b0316815260208101919091526040015f205460ff16156104e45760405163f53de75f60e01b815260040160405180910390fd5b600160025f8787858181106104fb576104fb611401565b90506020020160208101906105109190611328565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b385858381811061056a5761056a611401565b905060200201602081019061057f9190611328565b6040516001600160a01b03909116815260200160405180910390a15f60015f8787858181106105b0576105b0611401565b90506020020160208101906105c59190611328565b6001600160a01b03908116825260208201929092526040015f2054169050801561062257808484815181106105fc576105fc611401565b6001600160a01b03909216602092830291909101909101528261061e81611415565b9350505b50600101610478565b5080825280156106af576040516316bb16b760e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b5d8b5b89061068190859060040161136a565b5f604051808303815f87803b158015610698575f5ffd5b505af11580156106aa573d5f5f3e3d5ffd5b505050505b50505050565b6003602052815f5260405f2081815481106106ce575f80fd5b5f918252602090912001546001600160a01b03169150829050565b6106f1610ea3565b6106fb5f19610f46565b565b5f5f61070881610fdd565b6001600160a01b0383165f9081526002602052604090205460ff16156107415760405163091867bd60e11b815260040160405180910390fd5b6001600160a01b038381165f9081526001602052604090205416156107795760405163c45546f760e01b815260040160405180910390fd5b604080516001600160a01b03851660248083019190915282518083039091018152604490910182526020810180516001600160e01b031663189acdbd60e31b17905290515f917f0000000000000000000000000000000000000000000000000000000000000000916107ea9061122f565b6107f5929190611439565b604051809103905ff08015801561080e573d5f5f3e3d5ffd5b50905061081b8482611008565b6040805160018082528183019092525f916020808301908036833701905050905081815f8151811061084f5761084f611401565b6001600160a01b039283166020918202929092010152604051632ef047f960e11b81527f000000000000000000000000000000000000000000000000000000000000000090911690635de08ff2906108ab90849060040161136a565b5f604051808303815f87803b1580156108c2575f5ffd5b505af11580156108d4573d5f5f3e3d5ffd5b5093979650505050505050565b6108e9610f83565b6106fb5f611072565b6001600160a01b0381165f9081526003602090815260409182902080548351818402810184019094528084526060939283018282801561095957602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161093b575b50505050509050919050565b61096d610f83565b604051632ef047f960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690635de08ff2906109bb908590859060040161147d565b5f604051808303815f87803b1580156109d2575f5ffd5b505af11580156109e4573d5f5f3e3d5ffd5b505050505050565b603354610100900460ff1615808015610a0c5750603354600160ff909116105b80610a265750303b158015610a26575060335460ff166001145b610a8e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6033805460ff191660011790558015610ab1576033805461ff0019166101001790555b610aba83611072565b610ac382610f46565b8015610b09576033805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610b16610f83565b6001600160a01b038116610b7b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a85565b610b8481611072565b50565b610b8f6110c3565b60995480198219811614610bb65760405163c61dca5d60e01b815260040160405180910390fd5b609982905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f5f610bff81610fdd565b5f610c0d6020850185611328565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614158015610c8357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031614155b610ca05760405163638ca9f360e11b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063c2cca26d60e01b85604051602401610cda9190611577565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051610d179061122f565b610d22929190611439565b604051809103905ff080158015610d3b573d5f5f3e3d5ffd5b506001600160a01b038083165f9081526003602090815260408220805460018101825590835291200180546001600160a01b03191691831691909117905592506040805160018082528183019092525f9181602001602082028036833701905050905083815f81518110610db157610db1611401565b6001600160a01b039283166020918202929092010152604051632ef047f960e11b81527f000000000000000000000000000000000000000000000000000000000000000090911690635de08ff290610e0d90849060040161136a565b5f604051808303815f87803b158015610e24575f5ffd5b505af1158015610e36573d5f5f3e3d5ffd5b50505050610e45848387611174565b505050919050565b610e55610f83565b6040516316bb16b760e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b5d8b5b8906109bb908590859060040161147d565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610f05573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f2991906116b3565b6106fb57604051631d77d47760e21b815260040160405180910390fd5b609981905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6066546001600160a01b031633146106fb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a85565b609954600160ff83161b90811603610b845760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038281165f8181526001602090815260409182902080546001600160a01b031916948616948517905581519283528201929092527f6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f910160405180910390a15050565b606680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561111f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061114391906116d2565b6001600160a01b0316336001600160a01b0316146106fb5760405163794821ff60e01b815260040160405180910390fd5b6111846040820160208301611328565b6001600160a01b0316826001600160a01b0316846001600160a01b03167f6c9e2cbc6fcd0f5b21ee2edb38f3421d7538cd98b8ff00803f818793455685048460600160208101906111d591906116ed565b608086013560a08701356111ec60c0890189611706565b6111fd6101008b0160e08c01611328565b61120f6101208c016101008d016116ed565b6040516112229796959493929190611749565b60405180910390a4505050565b6106788061179f83390190565b5f6020828403121561124c575f5ffd5b5035919050565b5f5f83601f840112611263575f5ffd5b50813567ffffffffffffffff81111561127a575f5ffd5b6020830191508360208260051b8501011115611294575f5ffd5b9250929050565b5f5f602083850312156112ac575f5ffd5b823567ffffffffffffffff8111156112c2575f5ffd5b6112ce85828601611253565b90969095509350505050565b6001600160a01b0381168114610b84575f5ffd5b80356112f9816112da565b919050565b5f5f6040838503121561130f575f5ffd5b823561131a816112da565b946020939093013593505050565b5f60208284031215611338575f5ffd5b8135611343816112da565b9392505050565b5f6020828403121561135a575f5ffd5b813560ff81168114611343575f5ffd5b602080825282518282018190525f918401906040840190835b818110156113aa5783516001600160a01b0316835260209384019390920191600101611383565b509095945050505050565b5f602082840312156113c5575f5ffd5b813567ffffffffffffffff8111156113db575f5ffd5b82016101808185031215611343575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f6001820161143257634e487b7160e01b5f52601160045260245ffd5b5060010190565b60018060a01b0383168152604060208201525f82518060408401528060208501606085015e5f606082850101526060601f19601f8301168401019150509392505050565b602080825281018290525f8360408301825b858110156114bf5782356114a2816112da565b6001600160a01b031682526020928301929091019060010161148f565b5095945050505050565b803563ffffffff811681146112f9575f5ffd5b5f5f8335601e198436030181126114f1575f5ffd5b830160208101925035905067ffffffffffffffff811115611510575f5ffd5b803603821315611294575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b8035611551816112da565b6001600160a01b0316825263ffffffff61156d602083016114c9565b1660208301525050565b602081526115986020820161158b846112ee565b6001600160a01b03169052565b5f6115a5602084016112ee565b6001600160a01b0381166040840152506115c1604084016112ee565b6001600160a01b0381166060840152506115dd606084016114c9565b63ffffffff8116608084015250608083013560a08381019190915283013560c080840191909152611610908401846114dc565b61018060e08501526116276101a08501828461151e565b91505061163b610100840160e08601611546565b6116496101208501856114dc565b848303601f190161014086015261166183828461151e565b9250505061167261014085016112ee565b6001600160a01b038116610160850152506116916101608501856114dc565b848303601f19016101808601526116a983828461151e565b9695505050505050565b5f602082840312156116c3575f5ffd5b81518015158114611343575f5ffd5b5f602082840312156116e2575f5ffd5b8151611343816112da565b5f602082840312156116fd575f5ffd5b611343826114c9565b5f5f8335601e1984360301811261171b575f5ffd5b83018035915067ffffffffffffffff821115611735575f5ffd5b602001915036819003821315611294575f5ffd5b63ffffffff8816815286602082015285604082015260c060608201525f61177460c08301868861151e565b6001600160a01b039490941660808301525063ffffffff9190911660a0909101529594505050505056fe6080604052604051610678380380610678833981016040819052610022916103ed565b61002d82825f610034565b5050610513565b61003d836100f1565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e905f90a25f8251118061007c5750805b156100ec576100ea836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100e491906104af565b83610273565b505b505050565b6001600160a01b0381163b61015b5760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101cd816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561019a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101be91906104af565b6001600160a01b03163b151590565b6102325760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610152565b7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b0392909216919091179055565b606061029883836040518060600160405280602781526020016106516027913961029f565b9392505050565b60605f5f856001600160a01b0316856040516102bb91906104c8565b5f60405180830381855af49150503d805f81146102f3576040519150601f19603f3d011682016040523d82523d5f602084013e6102f8565b606091505b50909250905061030a86838387610314565b9695505050505050565b606083156103825782515f0361037b576001600160a01b0385163b61037b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610152565b508161038c565b61038c8383610394565b949350505050565b8151156103a45781518083602001fd5b8060405162461bcd60e51b815260040161015291906104de565b80516001600160a01b03811681146103d4575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f604083850312156103fe575f5ffd5b610407836103be565b60208401519092506001600160401b03811115610422575f5ffd5b8301601f81018513610432575f5ffd5b80516001600160401b0381111561044b5761044b6103d9565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610479576104796103d9565b604052818152828201602001871015610490575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f602082840312156104bf575f5ffd5b610298826103be565b5f82518060208501845e5f920191825250919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b610131806105205f395ff3fe608060405236601057600e6013565b005b600e5b601f601b6021565b60b3565b565b5f60527fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015608c573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019060ae919060d0565b905090565b365f5f375f5f365f845af43d5f5f3e80801560cc573d5ff35b3d5ffd5b5f6020828403121560df575f5ffd5b81516001600160a01b038116811460f4575f5ffd5b939250505056fea26469706673582212207b94aecb5f7696f142cc1acc4dea07c18e8f9d41766baf4a3c001a89adb8a8a764736f6c634300081e0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220cd4f3da8c52ee50ca272e75c791a323e41b21010fc73e5b052056ea8ab31a3b964736f6c634300081e0033",
}

// StrategyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyFactoryMetaData.ABI instead.
var StrategyFactoryABI = StrategyFactoryMetaData.ABI

// StrategyFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyFactoryMetaData.Bin instead.
var StrategyFactoryBin = StrategyFactoryMetaData.Bin

// DeployStrategyFactory deploys a new Ethereum contract, binding an instance of StrategyFactory to it.
func DeployStrategyFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address, _strategyBeacon common.Address, _durationVaultBeacon common.Address, _EIGEN common.Address, _bEIGEN common.Address) (common.Address, *types.Transaction, *StrategyFactory, error) {
	parsed, err := StrategyFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyFactoryBin), backend, _strategyManager, _pauserRegistry, _strategyBeacon, _durationVaultBeacon, _EIGEN, _bEIGEN)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyFactory{StrategyFactoryCaller: StrategyFactoryCaller{contract: contract}, StrategyFactoryTransactor: StrategyFactoryTransactor{contract: contract}, StrategyFactoryFilterer: StrategyFactoryFilterer{contract: contract}}, nil
}

// StrategyFactory is an auto generated Go binding around an Ethereum contract.
type StrategyFactory struct {
	StrategyFactoryCaller     // Read-only binding to the contract
	StrategyFactoryTransactor // Write-only binding to the contract
	StrategyFactoryFilterer   // Log filterer for contract events
}

// StrategyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyFactorySession struct {
	Contract     *StrategyFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyFactoryCallerSession struct {
	Contract *StrategyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StrategyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyFactoryTransactorSession struct {
	Contract     *StrategyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StrategyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyFactoryRaw struct {
	Contract *StrategyFactory // Generic contract binding to access the raw methods on
}

// StrategyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyFactoryCallerRaw struct {
	Contract *StrategyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyFactoryTransactorRaw struct {
	Contract *StrategyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyFactory creates a new instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactory(address common.Address, backend bind.ContractBackend) (*StrategyFactory, error) {
	contract, err := bindStrategyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyFactory{StrategyFactoryCaller: StrategyFactoryCaller{contract: contract}, StrategyFactoryTransactor: StrategyFactoryTransactor{contract: contract}, StrategyFactoryFilterer: StrategyFactoryFilterer{contract: contract}}, nil
}

// NewStrategyFactoryCaller creates a new read-only instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactoryCaller(address common.Address, caller bind.ContractCaller) (*StrategyFactoryCaller, error) {
	contract, err := bindStrategyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryCaller{contract: contract}, nil
}

// NewStrategyFactoryTransactor creates a new write-only instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyFactoryTransactor, error) {
	contract, err := bindStrategyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryTransactor{contract: contract}, nil
}

// NewStrategyFactoryFilterer creates a new log filterer instance of StrategyFactory, bound to a specific deployed contract.
func NewStrategyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyFactoryFilterer, error) {
	contract, err := bindStrategyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryFilterer{contract: contract}, nil
}

// bindStrategyFactory binds a generic wrapper to an already deployed contract.
func bindStrategyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyFactory *StrategyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyFactory.Contract.StrategyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyFactory *StrategyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.Contract.StrategyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyFactory *StrategyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyFactory.Contract.StrategyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyFactory *StrategyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyFactory *StrategyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyFactory *StrategyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyFactory.Contract.contract.Transact(opts, method, params...)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) EIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_StrategyFactory *StrategyFactorySession) EIGEN() (common.Address, error) {
	return _StrategyFactory.Contract.EIGEN(&_StrategyFactory.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) EIGEN() (common.Address, error) {
	return _StrategyFactory.Contract.EIGEN(&_StrategyFactory.CallOpts)
}

// BEIGEN is a free data retrieval call binding the contract method 0x3f4da4c6.
//
// Solidity: function bEIGEN() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) BEIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "bEIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BEIGEN is a free data retrieval call binding the contract method 0x3f4da4c6.
//
// Solidity: function bEIGEN() view returns(address)
func (_StrategyFactory *StrategyFactorySession) BEIGEN() (common.Address, error) {
	return _StrategyFactory.Contract.BEIGEN(&_StrategyFactory.CallOpts)
}

// BEIGEN is a free data retrieval call binding the contract method 0x3f4da4c6.
//
// Solidity: function bEIGEN() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) BEIGEN() (common.Address, error) {
	return _StrategyFactory.Contract.BEIGEN(&_StrategyFactory.CallOpts)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) DeployedStrategies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "deployedStrategies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactory *StrategyFactorySession) DeployedStrategies(arg0 common.Address) (common.Address, error) {
	return _StrategyFactory.Contract.DeployedStrategies(&_StrategyFactory.CallOpts, arg0)
}

// DeployedStrategies is a free data retrieval call binding the contract method 0x581dfd65.
//
// Solidity: function deployedStrategies(address ) view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) DeployedStrategies(arg0 common.Address) (common.Address, error) {
	return _StrategyFactory.Contract.DeployedStrategies(&_StrategyFactory.CallOpts, arg0)
}

// DurationVaultBeacon is a free data retrieval call binding the contract method 0x9e4f2dcd.
//
// Solidity: function durationVaultBeacon() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) DurationVaultBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "durationVaultBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DurationVaultBeacon is a free data retrieval call binding the contract method 0x9e4f2dcd.
//
// Solidity: function durationVaultBeacon() view returns(address)
func (_StrategyFactory *StrategyFactorySession) DurationVaultBeacon() (common.Address, error) {
	return _StrategyFactory.Contract.DurationVaultBeacon(&_StrategyFactory.CallOpts)
}

// DurationVaultBeacon is a free data retrieval call binding the contract method 0x9e4f2dcd.
//
// Solidity: function durationVaultBeacon() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) DurationVaultBeacon() (common.Address, error) {
	return _StrategyFactory.Contract.DurationVaultBeacon(&_StrategyFactory.CallOpts)
}

// DurationVaultsByToken is a free data retrieval call binding the contract method 0x3101d531.
//
// Solidity: function durationVaultsByToken(address , uint256 ) view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) DurationVaultsByToken(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "durationVaultsByToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DurationVaultsByToken is a free data retrieval call binding the contract method 0x3101d531.
//
// Solidity: function durationVaultsByToken(address , uint256 ) view returns(address)
func (_StrategyFactory *StrategyFactorySession) DurationVaultsByToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyFactory.Contract.DurationVaultsByToken(&_StrategyFactory.CallOpts, arg0, arg1)
}

// DurationVaultsByToken is a free data retrieval call binding the contract method 0x3101d531.
//
// Solidity: function durationVaultsByToken(address , uint256 ) view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) DurationVaultsByToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyFactory.Contract.DurationVaultsByToken(&_StrategyFactory.CallOpts, arg0, arg1)
}

// GetDurationVaults is a free data retrieval call binding the contract method 0xa5dae840.
//
// Solidity: function getDurationVaults(address token) view returns(address[])
func (_StrategyFactory *StrategyFactoryCaller) GetDurationVaults(opts *bind.CallOpts, token common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "getDurationVaults", token)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDurationVaults is a free data retrieval call binding the contract method 0xa5dae840.
//
// Solidity: function getDurationVaults(address token) view returns(address[])
func (_StrategyFactory *StrategyFactorySession) GetDurationVaults(token common.Address) ([]common.Address, error) {
	return _StrategyFactory.Contract.GetDurationVaults(&_StrategyFactory.CallOpts, token)
}

// GetDurationVaults is a free data retrieval call binding the contract method 0xa5dae840.
//
// Solidity: function getDurationVaults(address token) view returns(address[])
func (_StrategyFactory *StrategyFactoryCallerSession) GetDurationVaults(token common.Address) ([]common.Address, error) {
	return _StrategyFactory.Contract.GetDurationVaults(&_StrategyFactory.CallOpts, token)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactory *StrategyFactoryCaller) IsBlacklisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "isBlacklisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactory *StrategyFactorySession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _StrategyFactory.Contract.IsBlacklisted(&_StrategyFactory.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address ) view returns(bool)
func (_StrategyFactory *StrategyFactoryCallerSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _StrategyFactory.Contract.IsBlacklisted(&_StrategyFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyFactory *StrategyFactorySession) Owner() (common.Address, error) {
	return _StrategyFactory.Contract.Owner(&_StrategyFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) Owner() (common.Address, error) {
	return _StrategyFactory.Contract.Owner(&_StrategyFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyFactory *StrategyFactoryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyFactory *StrategyFactorySession) Paused(index uint8) (bool, error) {
	return _StrategyFactory.Contract.Paused(&_StrategyFactory.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyFactory *StrategyFactoryCallerSession) Paused(index uint8) (bool, error) {
	return _StrategyFactory.Contract.Paused(&_StrategyFactory.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyFactory *StrategyFactoryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyFactory *StrategyFactorySession) Paused0() (*big.Int, error) {
	return _StrategyFactory.Contract.Paused0(&_StrategyFactory.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyFactory *StrategyFactoryCallerSession) Paused0() (*big.Int, error) {
	return _StrategyFactory.Contract.Paused0(&_StrategyFactory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyFactory *StrategyFactorySession) PauserRegistry() (common.Address, error) {
	return _StrategyFactory.Contract.PauserRegistry(&_StrategyFactory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) PauserRegistry() (common.Address, error) {
	return _StrategyFactory.Contract.PauserRegistry(&_StrategyFactory.CallOpts)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) StrategyBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "strategyBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactory *StrategyFactorySession) StrategyBeacon() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyBeacon(&_StrategyFactory.CallOpts)
}

// StrategyBeacon is a free data retrieval call binding the contract method 0xf0062d9a.
//
// Solidity: function strategyBeacon() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) StrategyBeacon() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyBeacon(&_StrategyFactory.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyFactory *StrategyFactoryCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyFactory.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyFactory *StrategyFactorySession) StrategyManager() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyManager(&_StrategyFactory.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyFactory *StrategyFactoryCallerSession) StrategyManager() (common.Address, error) {
	return _StrategyFactory.Contract.StrategyManager(&_StrategyFactory.CallOpts)
}

// BlacklistTokens is a paid mutator transaction binding the contract method 0x23103c41.
//
// Solidity: function blacklistTokens(address[] tokens) returns()
func (_StrategyFactory *StrategyFactoryTransactor) BlacklistTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "blacklistTokens", tokens)
}

// BlacklistTokens is a paid mutator transaction binding the contract method 0x23103c41.
//
// Solidity: function blacklistTokens(address[] tokens) returns()
func (_StrategyFactory *StrategyFactorySession) BlacklistTokens(tokens []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.BlacklistTokens(&_StrategyFactory.TransactOpts, tokens)
}

// BlacklistTokens is a paid mutator transaction binding the contract method 0x23103c41.
//
// Solidity: function blacklistTokens(address[] tokens) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) BlacklistTokens(tokens []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.BlacklistTokens(&_StrategyFactory.TransactOpts, tokens)
}

// DeployDurationVaultStrategy is a paid mutator transaction binding the contract method 0xfb515303.
//
// Solidity: function deployDurationVaultStrategy((address,address,address,uint32,uint256,uint256,string,(address,uint32),bytes,address,string) config) returns(address newVault)
func (_StrategyFactory *StrategyFactoryTransactor) DeployDurationVaultStrategy(opts *bind.TransactOpts, config IDurationVaultStrategyTypesVaultConfig) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "deployDurationVaultStrategy", config)
}

// DeployDurationVaultStrategy is a paid mutator transaction binding the contract method 0xfb515303.
//
// Solidity: function deployDurationVaultStrategy((address,address,address,uint32,uint256,uint256,string,(address,uint32),bytes,address,string) config) returns(address newVault)
func (_StrategyFactory *StrategyFactorySession) DeployDurationVaultStrategy(config IDurationVaultStrategyTypesVaultConfig) (*types.Transaction, error) {
	return _StrategyFactory.Contract.DeployDurationVaultStrategy(&_StrategyFactory.TransactOpts, config)
}

// DeployDurationVaultStrategy is a paid mutator transaction binding the contract method 0xfb515303.
//
// Solidity: function deployDurationVaultStrategy((address,address,address,uint32,uint256,uint256,string,(address,uint32),bytes,address,string) config) returns(address newVault)
func (_StrategyFactory *StrategyFactoryTransactorSession) DeployDurationVaultStrategy(config IDurationVaultStrategyTypesVaultConfig) (*types.Transaction, error) {
	return _StrategyFactory.Contract.DeployDurationVaultStrategy(&_StrategyFactory.TransactOpts, config)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactory *StrategyFactoryTransactor) DeployNewStrategy(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "deployNewStrategy", token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactory *StrategyFactorySession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.DeployNewStrategy(&_StrategyFactory.TransactOpts, token)
}

// DeployNewStrategy is a paid mutator transaction binding the contract method 0x6b9b6229.
//
// Solidity: function deployNewStrategy(address token) returns(address newStrategy)
func (_StrategyFactory *StrategyFactoryTransactorSession) DeployNewStrategy(token common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.DeployNewStrategy(&_StrategyFactory.TransactOpts, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _initialOwner, uint256 _initialPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "initialize", _initialOwner, _initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _initialOwner, uint256 _initialPausedStatus) returns()
func (_StrategyFactory *StrategyFactorySession) Initialize(_initialOwner common.Address, _initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Initialize(&_StrategyFactory.TransactOpts, _initialOwner, _initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _initialOwner, uint256 _initialPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) Initialize(_initialOwner common.Address, _initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Initialize(&_StrategyFactory.TransactOpts, _initialOwner, _initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactorySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Pause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Pause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyFactory *StrategyFactoryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyFactory *StrategyFactorySession) PauseAll() (*types.Transaction, error) {
	return _StrategyFactory.Contract.PauseAll(&_StrategyFactory.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _StrategyFactory.Contract.PauseAll(&_StrategyFactory.TransactOpts)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactor) RemoveStrategiesFromWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "removeStrategiesFromWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactory *StrategyFactorySession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.RemoveStrategiesFromWhitelist(&_StrategyFactory.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromWhitelist is a paid mutator transaction binding the contract method 0xfe38b32d.
//
// Solidity: function removeStrategiesFromWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) RemoveStrategiesFromWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.RemoveStrategiesFromWhitelist(&_StrategyFactory.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyFactory *StrategyFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyFactory *StrategyFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyFactory.Contract.RenounceOwnership(&_StrategyFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategyFactory.Contract.RenounceOwnership(&_StrategyFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyFactory *StrategyFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyFactory *StrategyFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.TransferOwnership(&_StrategyFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.TransferOwnership(&_StrategyFactory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactorySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Unpause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyFactory.Contract.Unpause(&_StrategyFactory.TransactOpts, newPausedStatus)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactor) WhitelistStrategies(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.contract.Transact(opts, "whitelistStrategies", strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactory *StrategyFactorySession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.WhitelistStrategies(&_StrategyFactory.TransactOpts, strategiesToWhitelist)
}

// WhitelistStrategies is a paid mutator transaction binding the contract method 0xb768ebc9.
//
// Solidity: function whitelistStrategies(address[] strategiesToWhitelist) returns()
func (_StrategyFactory *StrategyFactoryTransactorSession) WhitelistStrategies(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyFactory.Contract.WhitelistStrategies(&_StrategyFactory.TransactOpts, strategiesToWhitelist)
}

// StrategyFactoryDurationVaultDeployedIterator is returned from FilterDurationVaultDeployed and is used to iterate over the raw logs and unpacked data for DurationVaultDeployed events raised by the StrategyFactory contract.
type StrategyFactoryDurationVaultDeployedIterator struct {
	Event *StrategyFactoryDurationVaultDeployed // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryDurationVaultDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryDurationVaultDeployed)
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
		it.Event = new(StrategyFactoryDurationVaultDeployed)
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
func (it *StrategyFactoryDurationVaultDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryDurationVaultDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryDurationVaultDeployed represents a DurationVaultDeployed event raised by the StrategyFactory contract.
type StrategyFactoryDurationVaultDeployed struct {
	Vault           common.Address
	UnderlyingToken common.Address
	VaultAdmin      common.Address
	Duration        uint32
	MaxPerDeposit   *big.Int
	StakeCap        *big.Int
	MetadataURI     string
	OperatorSetAVS  common.Address
	OperatorSetId   uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDurationVaultDeployed is a free log retrieval operation binding the contract event 0x6c9e2cbc6fcd0f5b21ee2edb38f3421d7538cd98b8ff00803f81879345568504.
//
// Solidity: event DurationVaultDeployed(address indexed vault, address indexed underlyingToken, address indexed vaultAdmin, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI, address operatorSetAVS, uint32 operatorSetId)
func (_StrategyFactory *StrategyFactoryFilterer) FilterDurationVaultDeployed(opts *bind.FilterOpts, vault []common.Address, underlyingToken []common.Address, vaultAdmin []common.Address) (*StrategyFactoryDurationVaultDeployedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}
	var vaultAdminRule []interface{}
	for _, vaultAdminItem := range vaultAdmin {
		vaultAdminRule = append(vaultAdminRule, vaultAdminItem)
	}

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "DurationVaultDeployed", vaultRule, underlyingTokenRule, vaultAdminRule)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryDurationVaultDeployedIterator{contract: _StrategyFactory.contract, event: "DurationVaultDeployed", logs: logs, sub: sub}, nil
}

// WatchDurationVaultDeployed is a free log subscription operation binding the contract event 0x6c9e2cbc6fcd0f5b21ee2edb38f3421d7538cd98b8ff00803f81879345568504.
//
// Solidity: event DurationVaultDeployed(address indexed vault, address indexed underlyingToken, address indexed vaultAdmin, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI, address operatorSetAVS, uint32 operatorSetId)
func (_StrategyFactory *StrategyFactoryFilterer) WatchDurationVaultDeployed(opts *bind.WatchOpts, sink chan<- *StrategyFactoryDurationVaultDeployed, vault []common.Address, underlyingToken []common.Address, vaultAdmin []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var underlyingTokenRule []interface{}
	for _, underlyingTokenItem := range underlyingToken {
		underlyingTokenRule = append(underlyingTokenRule, underlyingTokenItem)
	}
	var vaultAdminRule []interface{}
	for _, vaultAdminItem := range vaultAdmin {
		vaultAdminRule = append(vaultAdminRule, vaultAdminItem)
	}

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "DurationVaultDeployed", vaultRule, underlyingTokenRule, vaultAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryDurationVaultDeployed)
				if err := _StrategyFactory.contract.UnpackLog(event, "DurationVaultDeployed", log); err != nil {
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

// ParseDurationVaultDeployed is a log parse operation binding the contract event 0x6c9e2cbc6fcd0f5b21ee2edb38f3421d7538cd98b8ff00803f81879345568504.
//
// Solidity: event DurationVaultDeployed(address indexed vault, address indexed underlyingToken, address indexed vaultAdmin, uint32 duration, uint256 maxPerDeposit, uint256 stakeCap, string metadataURI, address operatorSetAVS, uint32 operatorSetId)
func (_StrategyFactory *StrategyFactoryFilterer) ParseDurationVaultDeployed(log types.Log) (*StrategyFactoryDurationVaultDeployed, error) {
	event := new(StrategyFactoryDurationVaultDeployed)
	if err := _StrategyFactory.contract.UnpackLog(event, "DurationVaultDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyFactory contract.
type StrategyFactoryInitializedIterator struct {
	Event *StrategyFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryInitialized)
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
		it.Event = new(StrategyFactoryInitialized)
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
func (it *StrategyFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryInitialized represents a Initialized event raised by the StrategyFactory contract.
type StrategyFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyFactory *StrategyFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyFactoryInitializedIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryInitializedIterator{contract: _StrategyFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyFactory *StrategyFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryInitialized)
				if err := _StrategyFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyFactory *StrategyFactoryFilterer) ParseInitialized(log types.Log) (*StrategyFactoryInitialized, error) {
	event := new(StrategyFactoryInitialized)
	if err := _StrategyFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrategyFactory contract.
type StrategyFactoryOwnershipTransferredIterator struct {
	Event *StrategyFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryOwnershipTransferred)
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
		it.Event = new(StrategyFactoryOwnershipTransferred)
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
func (it *StrategyFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the StrategyFactory contract.
type StrategyFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyFactory *StrategyFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrategyFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryOwnershipTransferredIterator{contract: _StrategyFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyFactory *StrategyFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrategyFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryOwnershipTransferred)
				if err := _StrategyFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategyFactory *StrategyFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*StrategyFactoryOwnershipTransferred, error) {
	event := new(StrategyFactoryOwnershipTransferred)
	if err := _StrategyFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrategyFactory contract.
type StrategyFactoryPausedIterator struct {
	Event *StrategyFactoryPaused // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryPaused)
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
		it.Event = new(StrategyFactoryPaused)
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
func (it *StrategyFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryPaused represents a Paused event raised by the StrategyFactory contract.
type StrategyFactoryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyFactoryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryPausedIterator{contract: _StrategyFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyFactoryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryPaused)
				if err := _StrategyFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) ParsePaused(log types.Log) (*StrategyFactoryPaused, error) {
	event := new(StrategyFactoryPaused)
	if err := _StrategyFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryStrategySetForTokenIterator is returned from FilterStrategySetForToken and is used to iterate over the raw logs and unpacked data for StrategySetForToken events raised by the StrategyFactory contract.
type StrategyFactoryStrategySetForTokenIterator struct {
	Event *StrategyFactoryStrategySetForToken // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryStrategySetForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryStrategySetForToken)
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
		it.Event = new(StrategyFactoryStrategySetForToken)
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
func (it *StrategyFactoryStrategySetForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryStrategySetForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryStrategySetForToken represents a StrategySetForToken event raised by the StrategyFactory contract.
type StrategyFactoryStrategySetForToken struct {
	Token    common.Address
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategySetForToken is a free log retrieval operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactory *StrategyFactoryFilterer) FilterStrategySetForToken(opts *bind.FilterOpts) (*StrategyFactoryStrategySetForTokenIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryStrategySetForTokenIterator{contract: _StrategyFactory.contract, event: "StrategySetForToken", logs: logs, sub: sub}, nil
}

// WatchStrategySetForToken is a free log subscription operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactory *StrategyFactoryFilterer) WatchStrategySetForToken(opts *bind.WatchOpts, sink chan<- *StrategyFactoryStrategySetForToken) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "StrategySetForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryStrategySetForToken)
				if err := _StrategyFactory.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
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

// ParseStrategySetForToken is a log parse operation binding the contract event 0x6852a55230ef089d785bce7ffbf757985de34026df90a87d7b4a6e56f95d251f.
//
// Solidity: event StrategySetForToken(address token, address strategy)
func (_StrategyFactory *StrategyFactoryFilterer) ParseStrategySetForToken(log types.Log) (*StrategyFactoryStrategySetForToken, error) {
	event := new(StrategyFactoryStrategySetForToken)
	if err := _StrategyFactory.contract.UnpackLog(event, "StrategySetForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryTokenBlacklistedIterator is returned from FilterTokenBlacklisted and is used to iterate over the raw logs and unpacked data for TokenBlacklisted events raised by the StrategyFactory contract.
type StrategyFactoryTokenBlacklistedIterator struct {
	Event *StrategyFactoryTokenBlacklisted // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryTokenBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryTokenBlacklisted)
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
		it.Event = new(StrategyFactoryTokenBlacklisted)
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
func (it *StrategyFactoryTokenBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryTokenBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryTokenBlacklisted represents a TokenBlacklisted event raised by the StrategyFactory contract.
type StrategyFactoryTokenBlacklisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenBlacklisted is a free log retrieval operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactory *StrategyFactoryFilterer) FilterTokenBlacklisted(opts *bind.FilterOpts) (*StrategyFactoryTokenBlacklistedIterator, error) {

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryTokenBlacklistedIterator{contract: _StrategyFactory.contract, event: "TokenBlacklisted", logs: logs, sub: sub}, nil
}

// WatchTokenBlacklisted is a free log subscription operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactory *StrategyFactoryFilterer) WatchTokenBlacklisted(opts *bind.WatchOpts, sink chan<- *StrategyFactoryTokenBlacklisted) (event.Subscription, error) {

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "TokenBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryTokenBlacklisted)
				if err := _StrategyFactory.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
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

// ParseTokenBlacklisted is a log parse operation binding the contract event 0x75519c51f39873ec0e27dd3bbc09549e4865a113f505393fb9eab5898f6418b3.
//
// Solidity: event TokenBlacklisted(address token)
func (_StrategyFactory *StrategyFactoryFilterer) ParseTokenBlacklisted(log types.Log) (*StrategyFactoryTokenBlacklisted, error) {
	event := new(StrategyFactoryTokenBlacklisted)
	if err := _StrategyFactory.contract.UnpackLog(event, "TokenBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrategyFactory contract.
type StrategyFactoryUnpausedIterator struct {
	Event *StrategyFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyFactoryUnpaused)
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
		it.Event = new(StrategyFactoryUnpaused)
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
func (it *StrategyFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyFactoryUnpaused represents a Unpaused event raised by the StrategyFactory contract.
type StrategyFactoryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyFactoryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyFactoryUnpausedIterator{contract: _StrategyFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyFactoryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyFactory.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyFactoryUnpaused)
				if err := _StrategyFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyFactory *StrategyFactoryFilterer) ParseUnpaused(log types.Log) (*StrategyFactoryUnpaused, error) {
	event := new(StrategyFactoryUnpaused)
	if err := _StrategyFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
