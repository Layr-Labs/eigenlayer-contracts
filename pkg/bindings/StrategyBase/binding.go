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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b506040516115c03803806115c083398101604081905261002e916101b6565b80826001600160a01b038116610057576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b031660805261006c8161008d565b60a052506001600160a01b03831660c0526100856100d3565b5050506102e7565b5f5f829050601f815111156100c0578260405163305a27a960e01b81526004016100b7919061028c565b60405180910390fd5b80516100cb826102c1565b179392505050565b5f54610100900460ff161561013a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100b7565b5f5460ff90811614610189575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461019f575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f606084860312156101c8575f5ffd5b83516101d38161018b565b60208501519093506101e48161018b565b60408501519092506001600160401b038111156101ff575f5ffd5b8401601f8101861361020f575f5ffd5b80516001600160401b03811115610228576102286101a2565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610256576102566101a2565b60405281815282820160200188101561026d575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102e1575f198160200360031b1b821691505b50919050565b60805160a05160c05161128561033b5f395f818161018001528181610415015281816107d9015261087601525f61053501525f818161024b015281816103220152818161058601526109a401526112855ff3fe608060405234801561000f575f5ffd5b5060043610610132575f3560e01c80637a8b2637116100b4578063c4d66de811610079578063c4d66de81461029b578063ce7c2ac2146102ae578063d9caed12146102c1578063e3dae51c146102d4578063f3e73875146102e7578063fabc1cbc146102fa575f5ffd5b80637a8b263714610233578063886f1195146102465780638c8710191461026d5780638f6a624014610280578063ab5921e114610293575f5ffd5b806354fd4d50116100fa57806354fd4d50146101cc578063553ca5f8146101e1578063595c6a67146101f45780635ac86ab7146101fc5780635c975abb1461022b575f5ffd5b8063136439dd146101365780632495a5991461014b57806339b70e381461017b5780633a98ef39146101a257806347e7ef24146101b9575b5f5ffd5b610149610144366004610ffa565b61030d565b005b60325461015e906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61015e7f000000000000000000000000000000000000000000000000000000000000000081565b6101ab60335481565b604051908152602001610172565b6101ab6101c7366004611028565b6103e2565b6101d461052e565b6040516101729190611052565b6101ab6101ef366004611087565b61055e565b610149610571565b61021b61020a3660046110b7565b6001805460ff9092161b9081161490565b6040519015158152602001610172565b6001546101ab565b6101ab610241366004610ffa565b610620565b61015e7f000000000000000000000000000000000000000000000000000000000000000081565b6101ab61027b366004610ffa565b610669565b6101ab61028e366004611087565b610673565b6101d4610680565b6101496102a9366004611087565b6106a0565b6101ab6102bc366004611087565b6107b2565b6101496102cf3660046110d2565b610844565b6101ab6102e2366004610ffa565b610961565b6101ab6102f5366004610ffa565b610998565b610149610308366004610ffa565b6109a2565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa15801561036f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103939190611110565b6103b057604051631d77d47760e21b815260040160405180910390fd5b60015481811681146103d55760405163c61dca5d60e01b815260040160405180910390fd5b6103de82610ab8565b5050565b600180545f91829181160361040a5760405163840a48d560e01b815260040160405180910390fd5b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610453576040516348da714f60e01b815260040160405180910390fd5b61045d8484610af5565b6033545f61046d6103e883611143565b90505f6103e861047b610b23565b6104859190611143565b90505f6104928783611156565b90508061049f8489611169565b6104a99190611180565b9550855f036104cb57604051630c392ed360e11b815260040160405180910390fd5b6104d58685611143565b60338190556f4b3b4ca85a86c47a098a223fffffffff101561050a57604051632f14e8a360e11b815260040160405180910390fd5b610523826103e860335461051e9190611143565b610b8d565b505050505092915050565b60606105597f0000000000000000000000000000000000000000000000000000000000000000610bd9565b905090565b5f61056b610241836107b2565b92915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156105d3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105f79190611110565b61061457604051631d77d47760e21b815260040160405180910390fd5b61061e5f19610ab8565b565b5f5f6103e86033546106329190611143565b90505f6103e8610640610b23565b61064a9190611143565b9050816106578583611169565b6106619190611180565b949350505050565b5f61056b82610961565b5f61056b6102f5836107b2565b60606040518060800160405280604d8152602001611203604d9139905090565b5f54610100900460ff16158080156106be57505f54600160ff909116105b806106d75750303b1580156106d757505f5460ff166001145b61073f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015610760575f805461ff0019166101001790555b61076982610c16565b80156103de575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b60405163fe243a1760e01b81526001600160a01b0382811660048301523060248301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063fe243a1790604401602060405180830381865afa158015610820573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056b919061119f565b6001805460029081160361086b5760405163840a48d560e01b815260040160405180910390fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108b4576040516348da714f60e01b815260040160405180910390fd5b6108bf848484610d61565b603354808311156108e357604051630b469df360e41b815260040160405180910390fd5b5f6108f06103e883611143565b90505f6103e86108fe610b23565b6109089190611143565b90505f826109168784611169565b6109209190611180565b905061092c8685611156565b60335561094c61093c8284611156565b6103e860335461051e9190611143565b610957888883610d94565b5050505050505050565b5f5f6103e86033546109739190611143565b90505f6103e8610981610b23565b61098b9190611143565b9050806106578386611169565b5f61056b82610620565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a2291906111b6565b6001600160a01b0316336001600160a01b031614610a535760405163794821ff60e01b815260040160405180910390fd5b60015480198219811614610a7a5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6032546001600160a01b038381169116146103de57604051630312abdd60e61b815260040160405180910390fd5b6032546040516370a0823160e01b81523060048201525f916001600160a01b0316906370a0823190602401602060405180830381865afa158015610b69573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610559919061119f565b7fd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be881610bc184670de0b6b3a7640000611169565b610bcb9190611180565b6040519081526020016107a6565b60605f610be583610da8565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f54610100900460ff16610c805760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610736565b603280546001600160a01b0319166001600160a01b038316179055610ca45f610ab8565b7f1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af55750760325f9054906101000a90046001600160a01b0316826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d16573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d3a91906111d1565b604080516001600160a01b03909316835260ff90911660208301520160405180910390a150565b6032546001600160a01b03838116911614610d8f57604051630312abdd60e61b815260040160405180910390fd5b505050565b610d8f6001600160a01b0383168483610dcf565b5f60ff8216601f81111561056b57604051632cd44ac360e21b815260040160405180910390fd5b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610d8f928692915f91610e5e918516908490610edd565b905080515f1480610e7e575080806020019051810190610e7e9190611110565b610d8f5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610736565b606061066184845f85855f5f866001600160a01b03168587604051610f0291906111ec565b5f6040518083038185875af1925050503d805f8114610f3c576040519150601f19603f3d011682016040523d82523d5f602084013e610f41565b606091505b5091509150610f5287838387610f5d565b979650505050505050565b60608315610fcb5782515f03610fc4576001600160a01b0385163b610fc45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610736565b5081610661565b6106618383815115610fe05781518083602001fd5b8060405162461bcd60e51b81526004016107369190611052565b5f6020828403121561100a575f5ffd5b5035919050565b6001600160a01b0381168114611025575f5ffd5b50565b5f5f60408385031215611039575f5ffd5b823561104481611011565b946020939093013593505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215611097575f5ffd5b81356110a281611011565b9392505050565b60ff81168114611025575f5ffd5b5f602082840312156110c7575f5ffd5b81356110a2816110a9565b5f5f5f606084860312156110e4575f5ffd5b83356110ef81611011565b925060208401356110ff81611011565b929592945050506040919091013590565b5f60208284031215611120575f5ffd5b815180151581146110a2575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561056b5761056b61112f565b8181038181111561056b5761056b61112f565b808202811582820484141761056b5761056b61112f565b5f8261119a57634e487b7160e01b5f52601260045260245ffd5b500490565b5f602082840312156111af575f5ffd5b5051919050565b5f602082840312156111c6575f5ffd5b81516110a281611011565b5f602082840312156111e1575f5ffd5b81516110a2816110a9565b5f82518060208501845e5f92019182525091905056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a2646970667358221220a862a526b97a17c779b349552714351b121d9dacb17bddd32bebb7a478528aa164736f6c634300081b0033",
}

// StrategyBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyBaseMetaData.ABI instead.
var StrategyBaseABI = StrategyBaseMetaData.ABI

// StrategyBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyBaseMetaData.Bin instead.
var StrategyBaseBin = StrategyBaseMetaData.Bin

// DeployStrategyBase deploys a new Ethereum contract, binding an instance of StrategyBase to it.
func DeployStrategyBase(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address, _version string) (common.Address, *types.Transaction, *StrategyBase, error) {
	parsed, err := StrategyBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyBaseBin), backend, _strategyManager, _pauserRegistry, _version)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyBase *StrategyBaseCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyBase *StrategyBaseSession) Version() (string, error) {
	return _StrategyBase.Contract.Version(&_StrategyBase.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StrategyBase *StrategyBaseCallerSession) Version() (string, error) {
	return _StrategyBase.Contract.Version(&_StrategyBase.CallOpts)
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
