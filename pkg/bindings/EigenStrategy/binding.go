// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EigenStrategy

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

// EigenStrategyMetaData contains all meta data concerning the EigenStrategy contract.
var EigenStrategyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigen\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_EIGEN\",\"type\":\"address\",\"internalType\":\"contractIEigen\"},{\"name\":\"_bEIGEN\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161181238038061181283398101604081905261002e916101bc565b82828280826001600160a01b03811661005a576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b031660805261006f81610093565b60a052506001600160a01b03831660c0526100886100d9565b5050505050506102ed565b5f5f829050601f815111156100c6578260405163305a27a960e01b81526004016100bd9190610292565b60405180910390fd5b80516100d1826102c7565b179392505050565b5f54610100900460ff16156101405760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100bd565b5f5460ff9081161461018f575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101a5575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f606084860312156101ce575f5ffd5b83516101d981610191565b60208501519093506101ea81610191565b60408501519092506001600160401b03811115610205575f5ffd5b8401601f81018613610215575f5ffd5b80516001600160401b0381111561022e5761022e6101a8565b604051601f8201601f19908116603f011681016001600160401b038111828210171561025c5761025c6101a8565b604052818152828201602001881015610273575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156102e7575f198160200360031b1b821691505b50919050565b60805160a05160c0516114d861033a5f395f8181610196015281816103990152818161076101526107e301525f6105a401525f8181610274015281816109910152610eca01526114d85ff3fe608060405234801561000f575f5ffd5b5060043610610148575f3560e01c80637a8b2637116100bf578063ce7c2ac211610079578063ce7c2ac2146102d7578063d9caed12146102ea578063e3dae51c146102fd578063f3e7387514610310578063fabc1cbc14610323578063fdc371ce14610336575f5ffd5b80637a8b26371461025c578063886f11951461026f5780638c871019146102965780638f6a6240146102a9578063ab5921e1146102bc578063c4d66de8146102c4575f5ffd5b8063485cc95511610110578063485cc955146101e257806354fd4d50146101f5578063553ca5f81461020a578063595c6a671461021d5780635ac86ab7146102255780635c975abb14610254575f5ffd5b8063136439dd1461014c5780632495a5991461016157806339b70e38146101915780633a98ef39146101b857806347e7ef24146101cf575b5f5ffd5b61015f61015a3660046111cb565b610349565b005b603254610174906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101747f000000000000000000000000000000000000000000000000000000000000000081565b6101c160335481565b604051908152602001610188565b6101c16101dd3660046111f6565b610383565b61015f6101f0366004611220565b6104b2565b6101fd61059d565b6040516101889190611257565b6101c161021836600461128c565b6105cd565b61015f6105e0565b6102446102333660046112bc565b6001805460ff9092161b9081161490565b6040519015158152602001610188565b6001546101c1565b6101c161026a3660046111cb565b6105f4565b6101747f000000000000000000000000000000000000000000000000000000000000000081565b6101c16102a43660046111cb565b61063d565b6101c16102b736600461128c565b610647565b6101fd610654565b61015f6102d236600461128c565b610674565b6101c16102e536600461128c565b61073a565b6101c16102f83660046112d7565b6107cc565b6101c161030b3660046111cb565b6108ce565b6101c161031e3660046111cb565b610905565b61015f6103313660046111cb565b61090f565b606454610174906001600160a01b031681565b61035161097c565b60015481811681146103765760405163c61dca5d60e01b815260040160405180910390fd5b61037f82610a1f565b5050565b5f5f61038e81610a5c565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103d7576040516348da714f60e01b815260040160405180910390fd5b6103e18484610a92565b6033545f6103f16103e883611329565b90505f6103e86103ff610b4b565b6104099190611329565b90505f610416878361133c565b905080610423848961134f565b61042d9190611366565b9550855f0361044f57604051630c392ed360e11b815260040160405180910390fd5b6104598685611329565b60338190556f4b3b4ca85a86c47a098a223fffffffff101561048e57604051632f14e8a360e11b815260040160405180910390fd5b6104a7826103e86033546104a29190611329565b610bb5565b505050505092915050565b5f54610100900460ff16158080156104d057505f54600160ff909116105b806104e95750303b1580156104e957505f5460ff166001145b61050e5760405162461bcd60e51b815260040161050590611385565b60405180910390fd5b5f805460ff19166001179055801561052f575f805461ff0019166101001790555b606480546001600160a01b0319166001600160a01b03851617905561055382610c01565b8015610598575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60606105c87f0000000000000000000000000000000000000000000000000000000000000000610d4c565b905090565b5f6105da61026a8361073a565b92915050565b6105e861097c565b6105f25f19610a1f565b565b5f5f6103e86033546106069190611329565b90505f6103e8610614610b4b565b61061e9190611329565b90508161062b858361134f565b6106359190611366565b949350505050565b5f6105da826108ce565b5f6105da61031e8361073a565b60606040518060800160405280604d8152602001611456604d9139905090565b5f54610100900460ff161580801561069257505f54600160ff909116105b806106ab5750303b1580156106ab57505f5460ff166001145b6106c75760405162461bcd60e51b815260040161050590611385565b5f805460ff1916600117905580156106e8575f805461ff0019166101001790555b6106f182610c01565b801561037f575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b60405163fe243a1760e01b81526001600160a01b0382811660048301523060248301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063fe243a1790604401602060405180830381865afa1580156107a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105da91906113d3565b5f60016107d881610a5c565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610821576040516348da714f60e01b815260040160405180910390fd5b61082c858585610d89565b6033548084111561085057604051630b469df360e41b815260040160405180910390fd5b5f61085d6103e883611329565b90505f6103e861086b610b4b565b6108759190611329565b905081610882878361134f565b61088c9190611366565b9450610898868461133c565b6033556108b86108a8868361133c565b6103e86033546104a29190611329565b6108c3888887610dcf565b505050509392505050565b5f5f6103e86033546108e09190611329565b90505f6103e86108ee610b4b565b6108f89190611329565b90508061062b838661134f565b5f6105da826105f4565b610917610ec8565b6001548019821981161461093e5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156109de573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a0291906113ea565b6105f257604051631d77d47760e21b815260040160405180910390fd5b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b610a71816001805460ff9092161b9081161490565b15610a8f5760405163840a48d560e01b815260040160405180910390fd5b50565b6032546001600160a01b0383811691161480610abb57506064546001600160a01b038381169116145b610ad857604051630312abdd60e61b815260040160405180910390fd5b6064546001600160a01b039081169083160361037f57606454604051636f074d1f60e11b8152600481018390526001600160a01b039091169063de0e9a3e906024015f604051808303815f87803b158015610b31575f5ffd5b505af1158015610b43573d5f5f3e3d5ffd5b505050505050565b6032546040516370a0823160e01b81523060048201525f916001600160a01b0316906370a0823190602401602060405180830381865afa158015610b91573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105c891906113d3565b7fd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be881610be984670de0b6b3a764000061134f565b610bf39190611366565b60405190815260200161072e565b5f54610100900460ff16610c6b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610505565b603280546001600160a01b0319166001600160a01b038316179055610c8f5f610a1f565b7f1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af55750760325f9054906101000a90046001600160a01b0316826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d01573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d259190611409565b604080516001600160a01b03909316835260ff90911660208301520160405180910390a150565b60605f610d5883610f79565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b6032546001600160a01b0383811691161480610db257506064546001600160a01b038381169116145b61059857604051630312abdd60e61b815260040160405180910390fd5b6064546001600160a01b0390811690831603610eb45760325460405163095ea7b360e01b81526001600160a01b038481166004830152602482018490529091169063095ea7b3906044016020604051808303815f875af1158015610e35573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e5991906113ea565b50606454604051630ea598cb60e41b8152600481018390526001600160a01b039091169063ea598cb0906024015f604051808303815f87803b158015610e9d575f5ffd5b505af1158015610eaf573d5f5f3e3d5ffd5b505050505b6105986001600160a01b0383168483610fa0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f24573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f489190611424565b6001600160a01b0316336001600160a01b0316146105f25760405163794821ff60e01b815260040160405180910390fd5b5f60ff8216601f8111156105da57604051632cd44ac360e21b815260040160405180910390fd5b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610598928692915f9161102f9185169084906110ae565b905080515f148061104f57508080602001905181019061104f91906113ea565b6105985760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610505565b606061063584845f85855f5f866001600160a01b031685876040516110d3919061143f565b5f6040518083038185875af1925050503d805f811461110d576040519150601f19603f3d011682016040523d82523d5f602084013e611112565b606091505b50915091506111238783838761112e565b979650505050505050565b6060831561119c5782515f03611195576001600160a01b0385163b6111955760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610505565b5081610635565b61063583838151156111b15781518083602001fd5b8060405162461bcd60e51b81526004016105059190611257565b5f602082840312156111db575f5ffd5b5035919050565b6001600160a01b0381168114610a8f575f5ffd5b5f5f60408385031215611207575f5ffd5b8235611212816111e2565b946020939093013593505050565b5f5f60408385031215611231575f5ffd5b823561123c816111e2565b9150602083013561124c816111e2565b809150509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f6020828403121561129c575f5ffd5b81356112a7816111e2565b9392505050565b60ff81168114610a8f575f5ffd5b5f602082840312156112cc575f5ffd5b81356112a7816112ae565b5f5f5f606084860312156112e9575f5ffd5b83356112f4816111e2565b92506020840135611304816111e2565b929592945050506040919091013590565b634e487b7160e01b5f52601160045260245ffd5b808201808211156105da576105da611315565b818103818111156105da576105da611315565b80820281158282048414176105da576105da611315565b5f8261138057634e487b7160e01b5f52601260045260245ffd5b500490565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b5f602082840312156113e3575f5ffd5b5051919050565b5f602082840312156113fa575f5ffd5b815180151581146112a7575f5ffd5b5f60208284031215611419575f5ffd5b81516112a7816112ae565b5f60208284031215611434575f5ffd5b81516112a7816111e2565b5f82518060208501845e5f92019182525091905056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a2646970667358221220a8b204485e6b6badc79eb89b625dc24b994cc150c07650aa24fcbd653df9eab364736f6c634300081b0033",
}

// EigenStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenStrategyMetaData.ABI instead.
var EigenStrategyABI = EigenStrategyMetaData.ABI

// EigenStrategyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EigenStrategyMetaData.Bin instead.
var EigenStrategyBin = EigenStrategyMetaData.Bin

// DeployEigenStrategy deploys a new Ethereum contract, binding an instance of EigenStrategy to it.
func DeployEigenStrategy(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address, _version string) (common.Address, *types.Transaction, *EigenStrategy, error) {
	parsed, err := EigenStrategyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EigenStrategyBin), backend, _strategyManager, _pauserRegistry, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EigenStrategy{EigenStrategyCaller: EigenStrategyCaller{contract: contract}, EigenStrategyTransactor: EigenStrategyTransactor{contract: contract}, EigenStrategyFilterer: EigenStrategyFilterer{contract: contract}}, nil
}

// EigenStrategy is an auto generated Go binding around an Ethereum contract.
type EigenStrategy struct {
	EigenStrategyCaller     // Read-only binding to the contract
	EigenStrategyTransactor // Write-only binding to the contract
	EigenStrategyFilterer   // Log filterer for contract events
}

// EigenStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigenStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigenStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigenStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigenStrategySession struct {
	Contract     *EigenStrategy    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EigenStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigenStrategyCallerSession struct {
	Contract *EigenStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EigenStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigenStrategyTransactorSession struct {
	Contract     *EigenStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EigenStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigenStrategyRaw struct {
	Contract *EigenStrategy // Generic contract binding to access the raw methods on
}

// EigenStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigenStrategyCallerRaw struct {
	Contract *EigenStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// EigenStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigenStrategyTransactorRaw struct {
	Contract *EigenStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigenStrategy creates a new instance of EigenStrategy, bound to a specific deployed contract.
func NewEigenStrategy(address common.Address, backend bind.ContractBackend) (*EigenStrategy, error) {
	contract, err := bindEigenStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EigenStrategy{EigenStrategyCaller: EigenStrategyCaller{contract: contract}, EigenStrategyTransactor: EigenStrategyTransactor{contract: contract}, EigenStrategyFilterer: EigenStrategyFilterer{contract: contract}}, nil
}

// NewEigenStrategyCaller creates a new read-only instance of EigenStrategy, bound to a specific deployed contract.
func NewEigenStrategyCaller(address common.Address, caller bind.ContractCaller) (*EigenStrategyCaller, error) {
	contract, err := bindEigenStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigenStrategyCaller{contract: contract}, nil
}

// NewEigenStrategyTransactor creates a new write-only instance of EigenStrategy, bound to a specific deployed contract.
func NewEigenStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*EigenStrategyTransactor, error) {
	contract, err := bindEigenStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigenStrategyTransactor{contract: contract}, nil
}

// NewEigenStrategyFilterer creates a new log filterer instance of EigenStrategy, bound to a specific deployed contract.
func NewEigenStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*EigenStrategyFilterer, error) {
	contract, err := bindEigenStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigenStrategyFilterer{contract: contract}, nil
}

// bindEigenStrategy binds a generic wrapper to an already deployed contract.
func bindEigenStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigenStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenStrategy *EigenStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenStrategy.Contract.EigenStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenStrategy *EigenStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenStrategy.Contract.EigenStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenStrategy *EigenStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenStrategy.Contract.EigenStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenStrategy *EigenStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenStrategy *EigenStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenStrategy *EigenStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenStrategy.Contract.contract.Transact(opts, method, params...)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_EigenStrategy *EigenStrategyCaller) EIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_EigenStrategy *EigenStrategySession) EIGEN() (common.Address, error) {
	return _EigenStrategy.Contract.EIGEN(&_EigenStrategy.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_EigenStrategy *EigenStrategyCallerSession) EIGEN() (common.Address, error) {
	return _EigenStrategy.Contract.EIGEN(&_EigenStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_EigenStrategy *EigenStrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_EigenStrategy *EigenStrategySession) Explanation() (string, error) {
	return _EigenStrategy.Contract.Explanation(&_EigenStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_EigenStrategy *EigenStrategyCallerSession) Explanation() (string, error) {
	return _EigenStrategy.Contract.Explanation(&_EigenStrategy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EigenStrategy *EigenStrategyCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EigenStrategy *EigenStrategySession) Paused(index uint8) (bool, error) {
	return _EigenStrategy.Contract.Paused(&_EigenStrategy.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EigenStrategy *EigenStrategyCallerSession) Paused(index uint8) (bool, error) {
	return _EigenStrategy.Contract.Paused(&_EigenStrategy.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EigenStrategy *EigenStrategySession) Paused0() (*big.Int, error) {
	return _EigenStrategy.Contract.Paused0(&_EigenStrategy.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) Paused0() (*big.Int, error) {
	return _EigenStrategy.Contract.Paused0(&_EigenStrategy.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EigenStrategy *EigenStrategyCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EigenStrategy *EigenStrategySession) PauserRegistry() (common.Address, error) {
	return _EigenStrategy.Contract.PauserRegistry(&_EigenStrategy.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EigenStrategy *EigenStrategyCallerSession) PauserRegistry() (common.Address, error) {
	return _EigenStrategy.Contract.PauserRegistry(&_EigenStrategy.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_EigenStrategy *EigenStrategySession) Shares(user common.Address) (*big.Int, error) {
	return _EigenStrategy.Contract.Shares(&_EigenStrategy.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _EigenStrategy.Contract.Shares(&_EigenStrategy.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_EigenStrategy *EigenStrategySession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.SharesToUnderlying(&_EigenStrategy.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.SharesToUnderlying(&_EigenStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_EigenStrategy *EigenStrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.SharesToUnderlyingView(&_EigenStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.SharesToUnderlyingView(&_EigenStrategy.CallOpts, amountShares)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_EigenStrategy *EigenStrategyCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_EigenStrategy *EigenStrategySession) StrategyManager() (common.Address, error) {
	return _EigenStrategy.Contract.StrategyManager(&_EigenStrategy.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_EigenStrategy *EigenStrategyCallerSession) StrategyManager() (common.Address, error) {
	return _EigenStrategy.Contract.StrategyManager(&_EigenStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_EigenStrategy *EigenStrategySession) TotalShares() (*big.Int, error) {
	return _EigenStrategy.Contract.TotalShares(&_EigenStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) TotalShares() (*big.Int, error) {
	return _EigenStrategy.Contract.TotalShares(&_EigenStrategy.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_EigenStrategy *EigenStrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.UnderlyingToShares(&_EigenStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.UnderlyingToShares(&_EigenStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_EigenStrategy *EigenStrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.UnderlyingToSharesView(&_EigenStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _EigenStrategy.Contract.UnderlyingToSharesView(&_EigenStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_EigenStrategy *EigenStrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_EigenStrategy *EigenStrategySession) UnderlyingToken() (common.Address, error) {
	return _EigenStrategy.Contract.UnderlyingToken(&_EigenStrategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_EigenStrategy *EigenStrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _EigenStrategy.Contract.UnderlyingToken(&_EigenStrategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_EigenStrategy *EigenStrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_EigenStrategy *EigenStrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _EigenStrategy.Contract.UserUnderlyingView(&_EigenStrategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_EigenStrategy *EigenStrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _EigenStrategy.Contract.UserUnderlyingView(&_EigenStrategy.CallOpts, user)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EigenStrategy *EigenStrategyCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EigenStrategy.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EigenStrategy *EigenStrategySession) Version() (string, error) {
	return _EigenStrategy.Contract.Version(&_EigenStrategy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EigenStrategy *EigenStrategyCallerSession) Version() (string, error) {
	return _EigenStrategy.Contract.Version(&_EigenStrategy.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_EigenStrategy *EigenStrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_EigenStrategy *EigenStrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Deposit(&_EigenStrategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_EigenStrategy *EigenStrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Deposit(&_EigenStrategy.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _EIGEN, address _bEIGEN) returns()
func (_EigenStrategy *EigenStrategyTransactor) Initialize(opts *bind.TransactOpts, _EIGEN common.Address, _bEIGEN common.Address) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "initialize", _EIGEN, _bEIGEN)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _EIGEN, address _bEIGEN) returns()
func (_EigenStrategy *EigenStrategySession) Initialize(_EIGEN common.Address, _bEIGEN common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize(&_EigenStrategy.TransactOpts, _EIGEN, _bEIGEN)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _EIGEN, address _bEIGEN) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) Initialize(_EIGEN common.Address, _bEIGEN common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize(&_EigenStrategy.TransactOpts, _EIGEN, _bEIGEN)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_EigenStrategy *EigenStrategyTransactor) Initialize0(opts *bind.TransactOpts, _underlyingToken common.Address) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "initialize0", _underlyingToken)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_EigenStrategy *EigenStrategySession) Initialize0(_underlyingToken common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize0(&_EigenStrategy.TransactOpts, _underlyingToken)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) Initialize0(_underlyingToken common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize0(&_EigenStrategy.TransactOpts, _underlyingToken)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EigenStrategy *EigenStrategyTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EigenStrategy *EigenStrategySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Pause(&_EigenStrategy.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Pause(&_EigenStrategy.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EigenStrategy *EigenStrategyTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EigenStrategy *EigenStrategySession) PauseAll() (*types.Transaction, error) {
	return _EigenStrategy.Contract.PauseAll(&_EigenStrategy.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EigenStrategy *EigenStrategyTransactorSession) PauseAll() (*types.Transaction, error) {
	return _EigenStrategy.Contract.PauseAll(&_EigenStrategy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EigenStrategy *EigenStrategyTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EigenStrategy *EigenStrategySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Unpause(&_EigenStrategy.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Unpause(&_EigenStrategy.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_EigenStrategy *EigenStrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_EigenStrategy *EigenStrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.UserUnderlying(&_EigenStrategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_EigenStrategy *EigenStrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.UserUnderlying(&_EigenStrategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_EigenStrategy *EigenStrategyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_EigenStrategy *EigenStrategySession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Withdraw(&_EigenStrategy.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_EigenStrategy *EigenStrategyTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Withdraw(&_EigenStrategy.TransactOpts, recipient, token, amountShares)
}

// EigenStrategyExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the EigenStrategy contract.
type EigenStrategyExchangeRateEmittedIterator struct {
	Event *EigenStrategyExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *EigenStrategyExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenStrategyExchangeRateEmitted)
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
		it.Event = new(EigenStrategyExchangeRateEmitted)
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
func (it *EigenStrategyExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenStrategyExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenStrategyExchangeRateEmitted represents a ExchangeRateEmitted event raised by the EigenStrategy contract.
type EigenStrategyExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_EigenStrategy *EigenStrategyFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*EigenStrategyExchangeRateEmittedIterator, error) {

	logs, sub, err := _EigenStrategy.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &EigenStrategyExchangeRateEmittedIterator{contract: _EigenStrategy.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_EigenStrategy *EigenStrategyFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *EigenStrategyExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _EigenStrategy.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenStrategyExchangeRateEmitted)
				if err := _EigenStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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
func (_EigenStrategy *EigenStrategyFilterer) ParseExchangeRateEmitted(log types.Log) (*EigenStrategyExchangeRateEmitted, error) {
	event := new(EigenStrategyExchangeRateEmitted)
	if err := _EigenStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenStrategyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EigenStrategy contract.
type EigenStrategyInitializedIterator struct {
	Event *EigenStrategyInitialized // Event containing the contract specifics and raw log

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
func (it *EigenStrategyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenStrategyInitialized)
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
		it.Event = new(EigenStrategyInitialized)
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
func (it *EigenStrategyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenStrategyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenStrategyInitialized represents a Initialized event raised by the EigenStrategy contract.
type EigenStrategyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EigenStrategy *EigenStrategyFilterer) FilterInitialized(opts *bind.FilterOpts) (*EigenStrategyInitializedIterator, error) {

	logs, sub, err := _EigenStrategy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EigenStrategyInitializedIterator{contract: _EigenStrategy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EigenStrategy *EigenStrategyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EigenStrategyInitialized) (event.Subscription, error) {

	logs, sub, err := _EigenStrategy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenStrategyInitialized)
				if err := _EigenStrategy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EigenStrategy *EigenStrategyFilterer) ParseInitialized(log types.Log) (*EigenStrategyInitialized, error) {
	event := new(EigenStrategyInitialized)
	if err := _EigenStrategy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenStrategyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EigenStrategy contract.
type EigenStrategyPausedIterator struct {
	Event *EigenStrategyPaused // Event containing the contract specifics and raw log

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
func (it *EigenStrategyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenStrategyPaused)
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
		it.Event = new(EigenStrategyPaused)
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
func (it *EigenStrategyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenStrategyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenStrategyPaused represents a Paused event raised by the EigenStrategy contract.
type EigenStrategyPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EigenStrategy *EigenStrategyFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*EigenStrategyPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenStrategy.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EigenStrategyPausedIterator{contract: _EigenStrategy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EigenStrategy *EigenStrategyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EigenStrategyPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenStrategy.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenStrategyPaused)
				if err := _EigenStrategy.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EigenStrategy *EigenStrategyFilterer) ParsePaused(log types.Log) (*EigenStrategyPaused, error) {
	event := new(EigenStrategyPaused)
	if err := _EigenStrategy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenStrategyStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the EigenStrategy contract.
type EigenStrategyStrategyTokenSetIterator struct {
	Event *EigenStrategyStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *EigenStrategyStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenStrategyStrategyTokenSet)
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
		it.Event = new(EigenStrategyStrategyTokenSet)
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
func (it *EigenStrategyStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenStrategyStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenStrategyStrategyTokenSet represents a StrategyTokenSet event raised by the EigenStrategy contract.
type EigenStrategyStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_EigenStrategy *EigenStrategyFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*EigenStrategyStrategyTokenSetIterator, error) {

	logs, sub, err := _EigenStrategy.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &EigenStrategyStrategyTokenSetIterator{contract: _EigenStrategy.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_EigenStrategy *EigenStrategyFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *EigenStrategyStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _EigenStrategy.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenStrategyStrategyTokenSet)
				if err := _EigenStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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
func (_EigenStrategy *EigenStrategyFilterer) ParseStrategyTokenSet(log types.Log) (*EigenStrategyStrategyTokenSet, error) {
	event := new(EigenStrategyStrategyTokenSet)
	if err := _EigenStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenStrategyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EigenStrategy contract.
type EigenStrategyUnpausedIterator struct {
	Event *EigenStrategyUnpaused // Event containing the contract specifics and raw log

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
func (it *EigenStrategyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenStrategyUnpaused)
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
		it.Event = new(EigenStrategyUnpaused)
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
func (it *EigenStrategyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenStrategyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenStrategyUnpaused represents a Unpaused event raised by the EigenStrategy contract.
type EigenStrategyUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EigenStrategy *EigenStrategyFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*EigenStrategyUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenStrategy.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EigenStrategyUnpausedIterator{contract: _EigenStrategy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EigenStrategy *EigenStrategyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EigenStrategyUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EigenStrategy.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenStrategyUnpaused)
				if err := _EigenStrategy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EigenStrategy *EigenStrategyFilterer) ParseUnpaused(log types.Log) (*EigenStrategyUnpaused, error) {
	event := new(EigenStrategyUnpaused)
	if err := _EigenStrategy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
