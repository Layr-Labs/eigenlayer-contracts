// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyBase

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

// StrategyBaseMetaData contains all meta data concerning the StrategyBase contract.
var StrategyBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161142a38038061142a83398101604081905261002f91610150565b806001600160a01b038116610057576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03908116608052821660a052610072610079565b505061018a565b600054610100900460ff16156100e55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614610136576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461014d57600080fd5b50565b6000806040838503121561016357600080fd5b825161016e81610138565b602084015190925061017f81610138565b809150509250929050565b60805160a0516112516101d96000396000818161017b0152818161040b015281816107b201526108510152600081816102310152818161031501528181610551015261098501526112516000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063886f1195116100ad578063ce7c2ac211610071578063ce7c2ac2146102a1578063d9caed12146102b4578063e3dae51c146102c7578063f3e73875146102da578063fabc1cbc146102ed57600080fd5b8063886f11951461022c5780638c871019146102535780638f6a624014610266578063ab5921e114610279578063c4d66de81461028e57600080fd5b8063553ca5f8116100f4578063553ca5f8146101c7578063595c6a67146101da5780635ac86ab7146101e25780635c975abb146102115780637a8b26371461021957600080fd5b8063136439dd146101315780632495a5991461014657806339b70e38146101765780633a98ef391461019d57806347e7ef24146101b4575b600080fd5b61014461013f366004610f83565b610300565b005b603254610159906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101597f000000000000000000000000000000000000000000000000000000000000000081565b6101a660335481565b60405190815260200161016d565b6101a66101c2366004610fb4565b6103d7565b6101a66101d5366004610fe0565b610528565b61014461053c565b6102016101f0366004611013565b6001805460ff9092161b9081161490565b604051901515815260200161016d565b6001546101a6565b6101a6610227366004610f83565b6105ee565b6101597f000000000000000000000000000000000000000000000000000000000000000081565b6101a6610261366004610f83565b610639565b6101a6610274366004610fe0565b610644565b610281610652565b60405161016d9190611054565b61014461029c366004610fe0565b610672565b6101a66102af366004610fe0565b61078a565b6101446102c2366004611087565b61081f565b6101a66102d5366004610f83565b61093f565b6101a66102e8366004610f83565b610978565b6101446102fb366004610f83565b610983565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610364573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038891906110c8565b6103a557604051631d77d47760e21b815260040160405180910390fd5b60015481811681146103ca5760405163c61dca5d60e01b815260040160405180910390fd5b6103d382610a9b565b5050565b6001805460009182918116036104005760405163840a48d560e01b815260040160405180910390fd5b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610449576040516348da714f60e01b815260040160405180910390fd5b6104538484610ad8565b60335460006104646103e883611100565b905060006103e8610473610b06565b61047d9190611100565b9050600061048b8783611113565b9050806104988489611126565b6104a2919061113d565b9550856000036104c557604051630c392ed360e11b815260040160405180910390fd5b6104cf8685611100565b60338190556f4b3b4ca85a86c47a098a223fffffffff101561050457604051632f14e8a360e11b815260040160405180910390fd5b61051d826103e86033546105189190611100565b610b78565b505050505092915050565b60006105366102278361078a565b92915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156105a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c491906110c8565b6105e157604051631d77d47760e21b815260040160405180910390fd5b6105ec600019610a9b565b565b6000806103e86033546106019190611100565b905060006103e8610610610b06565b61061a9190611100565b9050816106278583611126565b610631919061113d565b949350505050565b60006105368261093f565b60006105366102e88361078a565b60606040518060800160405280604d81526020016111cf604d9139905090565b600054610100900460ff16158080156106925750600054600160ff909116105b806106ac5750303b1580156106ac575060005460ff166001145b6107145760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610737576000805461ff0019166101001790555b61074082610bc4565b80156103d3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b60405163fe243a1760e01b81526001600160a01b0382811660048301523060248301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063fe243a1790604401602060405180830381865afa1580156107fb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610536919061115f565b600180546002908116036108465760405163840a48d560e01b815260040160405180910390fd5b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461088f576040516348da714f60e01b815260040160405180910390fd5b61089a848484610d14565b603354808311156108be57604051630b469df360e41b815260040160405180910390fd5b60006108cc6103e883611100565b905060006103e86108db610b06565b6108e59190611100565b90506000826108f48784611126565b6108fe919061113d565b905061090a8685611113565b60335561092a61091a8284611113565b6103e86033546105189190611100565b610935888883610d47565b5050505050505050565b6000806103e86033546109529190611100565b905060006103e8610961610b06565b61096b9190611100565b9050806106278386611126565b6000610536826105ee565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a059190611178565b6001600160a01b0316336001600160a01b031614610a365760405163794821ff60e01b815260040160405180910390fd5b60015480198219811614610a5d5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6032546001600160a01b038381169116146103d357604051630312abdd60e61b815260040160405180910390fd5b6032546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa158015610b4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b73919061115f565b905090565b7fd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be881610bac84670de0b6b3a7640000611126565b610bb6919061113d565b60405190815260200161077e565b600054610100900460ff16610c2f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840161070b565b603280546001600160a01b0319166001600160a01b038316179055610c546000610a9b565b7f1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507603260009054906101000a90046001600160a01b0316826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ced9190611195565b604080516001600160a01b03909316835260ff90911660208301520160405180910390a150565b6032546001600160a01b03838116911614610d4257604051630312abdd60e61b815260040160405180910390fd5b505050565b604080516001600160a01b03858116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610d429290851691869185918591859190600090610de09084908490610e60565b9050805160001480610e01575080806020019051810190610e0191906110c8565b610d425760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161070b565b6060610631848460008585600080866001600160a01b03168587604051610e8791906111b2565b60006040518083038185875af1925050503d8060008114610ec4576040519150601f19603f3d011682016040523d82523d6000602084013e610ec9565b606091505b5091509150610eda87838387610ee5565b979650505050505050565b60608315610f54578251600003610f4d576001600160a01b0385163b610f4d5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161070b565b5081610631565b6106318383815115610f695781518083602001fd5b8060405162461bcd60e51b815260040161070b9190611054565b600060208284031215610f9557600080fd5b5035919050565b6001600160a01b0381168114610fb157600080fd5b50565b60008060408385031215610fc757600080fd5b8235610fd281610f9c565b946020939093013593505050565b600060208284031215610ff257600080fd5b8135610ffd81610f9c565b9392505050565b60ff81168114610fb157600080fd5b60006020828403121561102557600080fd5b8135610ffd81611004565b60005b8381101561104b578181015183820152602001611033565b50506000910152565b6020815260008251806020840152611073816040850160208701611030565b601f01601f19169190910160400192915050565b60008060006060848603121561109c57600080fd5b83356110a781610f9c565b925060208401356110b781610f9c565b929592945050506040919091013590565b6000602082840312156110da57600080fd5b81518015158114610ffd57600080fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610536576105366110ea565b81810381811115610536576105366110ea565b8082028115828204841417610536576105366110ea565b60008261115a57634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561117157600080fd5b5051919050565b60006020828403121561118a57600080fd5b8151610ffd81610f9c565b6000602082840312156111a757600080fd5b8151610ffd81611004565b600082516111c4818460208701611030565b919091019291505056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a2646970667358221220a29fb2f5a3c2a2c655a4f4396b4ab1275b69e6f553f45d033004e51e1a7f207064736f6c634300081b0033",
}

// StrategyBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyBaseMetaData.ABI instead.
var StrategyBaseABI = StrategyBaseMetaData.ABI

// StrategyBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyBaseMetaData.Bin instead.
var StrategyBaseBin = StrategyBaseMetaData.Bin

// DeployStrategyBase deploys a new Ethereum contract, binding an instance of StrategyBase to it.
func DeployStrategyBase(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *StrategyBase, error) {
	parsed, err := StrategyBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyBaseBin), backend, _strategyManager, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyBase{StrategyBaseCaller: StrategyBaseCaller{contract: contract}, StrategyBaseTransactor: StrategyBaseTransactor{contract: contract}, StrategyBaseFilterer: StrategyBaseFilterer{contract: contract}}, nil
}

// StrategyBase is an auto generated Go binding around an Ethereum contract.
type StrategyBase struct {
	StrategyBaseCaller     // Read-only binding to the contract
	StrategyBaseTransactor // Write-only binding to the contract
	StrategyBaseFilterer   // Log filterer for contract events
}

// StrategyBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyBaseSession struct {
	Contract     *StrategyBase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyBaseCallerSession struct {
	Contract *StrategyBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StrategyBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyBaseTransactorSession struct {
	Contract     *StrategyBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StrategyBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyBaseRaw struct {
	Contract *StrategyBase // Generic contract binding to access the raw methods on
}

// StrategyBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyBaseCallerRaw struct {
	Contract *StrategyBaseCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyBaseTransactorRaw struct {
	Contract *StrategyBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyBase creates a new instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBase(address common.Address, backend bind.ContractBackend) (*StrategyBase, error) {
	contract, err := bindStrategyBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyBase{StrategyBaseCaller: StrategyBaseCaller{contract: contract}, StrategyBaseTransactor: StrategyBaseTransactor{contract: contract}, StrategyBaseFilterer: StrategyBaseFilterer{contract: contract}}, nil
}

// NewStrategyBaseCaller creates a new read-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseCaller(address common.Address, caller bind.ContractCaller) (*StrategyBaseCaller, error) {
	contract, err := bindStrategyBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseCaller{contract: contract}, nil
}

// NewStrategyBaseTransactor creates a new write-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyBaseTransactor, error) {
	contract, err := bindStrategyBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTransactor{contract: contract}, nil
}

// NewStrategyBaseFilterer creates a new log filterer instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyBaseFilterer, error) {
	contract, err := bindStrategyBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseFilterer{contract: contract}, nil
}

// bindStrategyBase binds a generic wrapper to an already deployed contract.
func bindStrategyBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.StrategyBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseSession) Explanation() (string, error) {
	return _StrategyBase.Contract.Explanation(&_StrategyBase.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseCallerSession) Explanation() (string, error) {
	return _StrategyBase.Contract.Explanation(&_StrategyBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBase *StrategyBaseCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBase *StrategyBaseSession) Paused(index uint8) (bool, error) {
	return _StrategyBase.Contract.Paused(&_StrategyBase.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) Paused(index uint8) (bool, error) {
	return _StrategyBase.Contract.Paused(&_StrategyBase.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) Paused0() (*big.Int, error) {
	return _StrategyBase.Contract.Paused0(&_StrategyBase.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) Paused0() (*big.Int, error) {
	return _StrategyBase.Contract.Paused0(&_StrategyBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBase *StrategyBaseCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBase *StrategyBaseSession) PauserRegistry() (common.Address, error) {
	return _StrategyBase.Contract.PauserRegistry(&_StrategyBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) PauserRegistry() (common.Address, error) {
	return _StrategyBase.Contract.PauserRegistry(&_StrategyBase.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.Shares(&_StrategyBase.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.Shares(&_StrategyBase.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlying(&_StrategyBase.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlying(&_StrategyBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlyingView(&_StrategyBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlyingView(&_StrategyBase.CallOpts, amountShares)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseSession) StrategyManager() (common.Address, error) {
	return _StrategyBase.Contract.StrategyManager(&_StrategyBase.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) StrategyManager() (common.Address, error) {
	return _StrategyBase.Contract.StrategyManager(&_StrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) TotalShares() (*big.Int, error) {
	return _StrategyBase.Contract.TotalShares(&_StrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) TotalShares() (*big.Int, error) {
	return _StrategyBase.Contract.TotalShares(&_StrategyBase.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToShares(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToShares(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToSharesView(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToSharesView(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBase *StrategyBaseCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBase *StrategyBaseSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBase.Contract.UnderlyingToken(&_StrategyBase.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBase.Contract.UnderlyingToken(&_StrategyBase.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserUnderlyingView(&_StrategyBase.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserUnderlyingView(&_StrategyBase.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_StrategyBase *StrategyBaseTransactor) Initialize(opts *bind.TransactOpts, _underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "initialize", _underlyingToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_StrategyBase *StrategyBaseSession) Initialize(_underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _underlyingToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Initialize(_underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _underlyingToken)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Pause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Pause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBase *StrategyBaseTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBase *StrategyBaseSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBase.Contract.PauseAll(&_StrategyBase.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBase *StrategyBaseTransactorSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBase.Contract.PauseAll(&_StrategyBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Unpause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Unpause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserUnderlying(&_StrategyBase.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserUnderlying(&_StrategyBase.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, token, amountShares)
}

// StrategyBaseExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the StrategyBase contract.
type StrategyBaseExchangeRateEmittedIterator struct {
	Event *StrategyBaseExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *StrategyBaseExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseExchangeRateEmitted)
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
		it.Event = new(StrategyBaseExchangeRateEmitted)
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
func (it *StrategyBaseExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseExchangeRateEmitted represents a ExchangeRateEmitted event raised by the StrategyBase contract.
type StrategyBaseExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBase *StrategyBaseFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*StrategyBaseExchangeRateEmittedIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseExchangeRateEmittedIterator{contract: _StrategyBase.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBase *StrategyBaseFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *StrategyBaseExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseExchangeRateEmitted)
				if err := _StrategyBase.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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

// ParseExchangeRateEmitted is a log parse operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBase *StrategyBaseFilterer) ParseExchangeRateEmitted(log types.Log) (*StrategyBaseExchangeRateEmitted, error) {
	event := new(StrategyBaseExchangeRateEmitted)
	if err := _StrategyBase.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyBase contract.
type StrategyBaseInitializedIterator struct {
	Event *StrategyBaseInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseInitialized)
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
		it.Event = new(StrategyBaseInitialized)
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
func (it *StrategyBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseInitialized represents a Initialized event raised by the StrategyBase contract.
type StrategyBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyBase *StrategyBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyBaseInitializedIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseInitializedIterator{contract: _StrategyBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyBase *StrategyBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseInitialized)
				if err := _StrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StrategyBase *StrategyBaseFilterer) ParseInitialized(log types.Log) (*StrategyBaseInitialized, error) {
	event := new(StrategyBaseInitialized)
	if err := _StrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrategyBase contract.
type StrategyBasePausedIterator struct {
	Event *StrategyBasePaused // Event containing the contract specifics and raw log

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
func (it *StrategyBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBasePaused)
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
		it.Event = new(StrategyBasePaused)
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
func (it *StrategyBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBasePaused represents a Paused event raised by the StrategyBase contract.
type StrategyBasePaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBasePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBasePausedIterator{contract: _StrategyBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyBasePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBasePaused)
				if err := _StrategyBase.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StrategyBase *StrategyBaseFilterer) ParsePaused(log types.Log) (*StrategyBasePaused, error) {
	event := new(StrategyBasePaused)
	if err := _StrategyBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the StrategyBase contract.
type StrategyBaseStrategyTokenSetIterator struct {
	Event *StrategyBaseStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *StrategyBaseStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseStrategyTokenSet)
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
		it.Event = new(StrategyBaseStrategyTokenSet)
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
func (it *StrategyBaseStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseStrategyTokenSet represents a StrategyTokenSet event raised by the StrategyBase contract.
type StrategyBaseStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBase *StrategyBaseFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*StrategyBaseStrategyTokenSetIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseStrategyTokenSetIterator{contract: _StrategyBase.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBase *StrategyBaseFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *StrategyBaseStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseStrategyTokenSet)
				if err := _StrategyBase.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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

// ParseStrategyTokenSet is a log parse operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBase *StrategyBaseFilterer) ParseStrategyTokenSet(log types.Log) (*StrategyBaseStrategyTokenSet, error) {
	event := new(StrategyBaseStrategyTokenSet)
	if err := _StrategyBase.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrategyBase contract.
type StrategyBaseUnpausedIterator struct {
	Event *StrategyBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUnpaused)
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
		it.Event = new(StrategyBaseUnpaused)
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
func (it *StrategyBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUnpaused represents a Unpaused event raised by the StrategyBase contract.
type StrategyBaseUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBaseUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUnpausedIterator{contract: _StrategyBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyBaseUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUnpaused)
				if err := _StrategyBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StrategyBase *StrategyBaseFilterer) ParseUnpaused(log types.Log) (*StrategyBaseUnpaused, error) {
	event := new(StrategyBaseUnpaused)
	if err := _StrategyBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
