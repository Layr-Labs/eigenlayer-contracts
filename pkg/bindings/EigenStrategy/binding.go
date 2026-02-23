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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigen\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beforeAddShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beforeRemoveShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_EIGEN\",\"type\":\"address\",\"internalType\":\"contractIEigen\"},{\"name\":\"_bEIGEN\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161167a38038061167a83398101604081905261002e9161014f565b8181806001600160a01b038116610058576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03908116608052821660a05261007361007c565b50505050610187565b5f54610100900460ff16156100e75760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610136575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461014c575f5ffd5b50565b5f5f60408385031215610160575f5ffd5b825161016b81610138565b602084015190925061017c81610138565b809150509250929050565b60805160a0516114a86101d25f395f81816101b40152818161036a015281816103f801528181610790015261081201525f818161027d015281816109c00152610ec101526114a85ff3fe608060405234801561000f575f5ffd5b5060043610610153575f3560e01c80637a8b2637116100bf578063ce7c2ac211610079578063ce7c2ac2146102ed578063d9caed1214610300578063e3dae51c14610313578063f3e7387514610326578063fabc1cbc14610339578063fdc371ce1461034c575f5ffd5b80637a8b263714610265578063886f1195146102785780638c8710191461029f5780638f6a6240146102b2578063ab5921e1146102c5578063c4d66de8146102da575f5ffd5b8063485cc95511610110578063485cc95514610200578063553ca5f814610213578063595c6a67146102265780635ac86ab71461022e5780635c975abb1461025d57806373e3c28014610157575f5ffd5b806303e3e6eb14610157578063136439dd1461016c5780632495a5991461017f57806339b70e38146101af5780633a98ef39146101d657806347e7ef24146101ed575b5f5ffd5b61016a6101653660046111af565b61035f565b005b61016a61017a3660046111d9565b6103ac565b603254610192906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101927f000000000000000000000000000000000000000000000000000000000000000081565b6101df60335481565b6040519081526020016101a6565b6101df6101fb3660046111af565b6103e2565b61016a61020e3660046111f0565b610511565b6101df610221366004611227565b6105fc565b61016a61060f565b61024d61023c366004611257565b6001805460ff9092161b9081161490565b60405190151581526020016101a6565b6001546101df565b6101df6102733660046111d9565b610623565b6101927f000000000000000000000000000000000000000000000000000000000000000081565b6101df6102ad3660046111d9565b61066c565b6101df6102c0366004611227565b610676565b6102cd610683565b6040516101a69190611272565b61016a6102e8366004611227565b6106a3565b6101df6102fb366004611227565b610769565b6101df61030e3660046112a7565b6107fb565b6101df6103213660046111d9565b6108fd565b6101df6103343660046111d9565b610934565b61016a6103473660046111d9565b61093e565b606454610192906001600160a01b031681565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103a8576040516348da714f60e01b815260040160405180910390fd5b5050565b6103b46109ab565b60015481811681146103d95760405163c61dca5d60e01b815260040160405180910390fd5b6103a882610a4e565b5f5f6103ed81610a8b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610436576040516348da714f60e01b815260040160405180910390fd5b6104408484610ac1565b6033545f6104506103e8836112f9565b90505f6103e861045e610b7a565b61046891906112f9565b90505f610475878361130c565b905080610482848961131f565b61048c9190611336565b9550855f036104ae57604051630c392ed360e11b815260040160405180910390fd5b6104b886856112f9565b60338190556f4b3b4ca85a86c47a098a223fffffffff10156104ed57604051632f14e8a360e11b815260040160405180910390fd5b610506826103e860335461050191906112f9565b610be9565b505050505092915050565b5f54610100900460ff161580801561052f57505f54600160ff909116105b806105485750303b15801561054857505f5460ff166001145b61056d5760405162461bcd60e51b815260040161056490611355565b60405180910390fd5b5f805460ff19166001179055801561058e575f805461ff0019166101001790555b606480546001600160a01b0319166001600160a01b0385161790556105b282610c35565b80156105f7575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b5f61060961027383610769565b92915050565b6106176109ab565b6106215f19610a4e565b565b5f5f6103e860335461063591906112f9565b90505f6103e8610643610b7a565b61064d91906112f9565b90508161065a858361131f565b6106649190611336565b949350505050565b5f610609826108fd565b5f61060961033483610769565b60606040518060800160405280604d8152602001611426604d9139905090565b5f54610100900460ff16158080156106c157505f54600160ff909116105b806106da5750303b1580156106da57505f5460ff166001145b6106f65760405162461bcd60e51b815260040161056490611355565b5f805460ff191660011790558015610717575f805461ff0019166101001790555b61072082610c35565b80156103a8575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b60405163fe243a1760e01b81526001600160a01b0382811660048301523060248301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063fe243a1790604401602060405180830381865afa1580156107d7573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061060991906113a3565b5f600161080781610a8b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610850576040516348da714f60e01b815260040160405180910390fd5b61085b858585610d80565b6033548084111561087f57604051630b469df360e41b815260040160405180910390fd5b5f61088c6103e8836112f9565b90505f6103e861089a610b7a565b6108a491906112f9565b9050816108b1878361131f565b6108bb9190611336565b94506108c7868461130c565b6033556108e76108d7868361130c565b6103e860335461050191906112f9565b6108f2888887610dc6565b505050509392505050565b5f5f6103e860335461090f91906112f9565b90505f6103e861091d610b7a565b61092791906112f9565b90508061065a838661131f565b5f61060982610623565b610946610ebf565b6001548019821981161461096d5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610a0d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a3191906113ba565b61062157604051631d77d47760e21b815260040160405180910390fd5b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b610aa0816001805460ff9092161b9081161490565b15610abe5760405163840a48d560e01b815260040160405180910390fd5b50565b6032546001600160a01b0383811691161480610aea57506064546001600160a01b038381169116145b610b0757604051630312abdd60e61b815260040160405180910390fd5b6064546001600160a01b03908116908316036103a857606454604051636f074d1f60e11b8152600481018390526001600160a01b039091169063de0e9a3e906024015f604051808303815f87803b158015610b60575f5ffd5b505af1158015610b72573d5f5f3e3d5ffd5b505050505050565b6032546040516370a0823160e01b81523060048201525f916001600160a01b0316906370a0823190602401602060405180830381865afa158015610bc0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be491906113a3565b905090565b7fd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be881610c1d84670de0b6b3a764000061131f565b610c279190611336565b60405190815260200161075d565b5f54610100900460ff16610c9f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610564565b603280546001600160a01b0319166001600160a01b038316179055610cc35f610a4e565b7f1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af55750760325f9054906101000a90046001600160a01b0316826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d35573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d5991906113d9565b604080516001600160a01b03909316835260ff90911660208301520160405180910390a150565b6032546001600160a01b0383811691161480610da957506064546001600160a01b038381169116145b6105f757604051630312abdd60e61b815260040160405180910390fd5b6064546001600160a01b0390811690831603610eab5760325460405163095ea7b360e01b81526001600160a01b038481166004830152602482018490529091169063095ea7b3906044016020604051808303815f875af1158015610e2c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e5091906113ba565b50606454604051630ea598cb60e41b8152600481018390526001600160a01b039091169063ea598cb0906024015f604051808303815f87803b158015610e94575f5ffd5b505af1158015610ea6573d5f5f3e3d5ffd5b505050505b6105f76001600160a01b0383168483610f70565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f1b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f3f91906113f4565b6001600160a01b0316336001600160a01b0316146106215760405163794821ff60e01b815260040160405180910390fd5b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526105f7928692915f91610fff91851690849061107e565b905080515f148061101f57508080602001905181019061101f91906113ba565b6105f75760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610564565b606061066484845f85855f5f866001600160a01b031685876040516110a3919061140f565b5f6040518083038185875af1925050503d805f81146110dd576040519150601f19603f3d011682016040523d82523d5f602084013e6110e2565b606091505b50915091506110f3878383876110fe565b979650505050505050565b6060831561116c5782515f03611165576001600160a01b0385163b6111655760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610564565b5081610664565b61066483838151156111815781518083602001fd5b8060405162461bcd60e51b81526004016105649190611272565b6001600160a01b0381168114610abe575f5ffd5b5f5f604083850312156111c0575f5ffd5b82356111cb8161119b565b946020939093013593505050565b5f602082840312156111e9575f5ffd5b5035919050565b5f5f60408385031215611201575f5ffd5b823561120c8161119b565b9150602083013561121c8161119b565b809150509250929050565b5f60208284031215611237575f5ffd5b81356112428161119b565b9392505050565b60ff81168114610abe575f5ffd5b5f60208284031215611267575f5ffd5b813561124281611249565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f606084860312156112b9575f5ffd5b83356112c48161119b565b925060208401356112d48161119b565b929592945050506040919091013590565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610609576106096112e5565b81810381811115610609576106096112e5565b8082028115828204841417610609576106096112e5565b5f8261135057634e487b7160e01b5f52601260045260245ffd5b500490565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b5f602082840312156113b3575f5ffd5b5051919050565b5f602082840312156113ca575f5ffd5b81518015158114611242575f5ffd5b5f602082840312156113e9575f5ffd5b815161124281611249565b5f60208284031215611404575f5ffd5b81516112428161119b565b5f82518060208501845e5f92019182525091905056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a26469706673582212201e029437f2e814142ddc4881ce7cd13c117417abfb9dfe1532a2bb4316affd0b64736f6c634300081e0033",
}

// EigenStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenStrategyMetaData.ABI instead.
var EigenStrategyABI = EigenStrategyMetaData.ABI

// EigenStrategyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EigenStrategyMetaData.Bin instead.
var EigenStrategyBin = EigenStrategyMetaData.Bin

// DeployEigenStrategy deploys a new Ethereum contract, binding an instance of EigenStrategy to it.
func DeployEigenStrategy(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *EigenStrategy, error) {
	parsed, err := EigenStrategyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EigenStrategyBin), backend, _strategyManager, _pauserRegistry)
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

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address , uint256 ) returns()
func (_EigenStrategy *EigenStrategyTransactor) BeforeAddShares(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "beforeAddShares", arg0, arg1)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address , uint256 ) returns()
func (_EigenStrategy *EigenStrategySession) BeforeAddShares(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.BeforeAddShares(&_EigenStrategy.TransactOpts, arg0, arg1)
}

// BeforeAddShares is a paid mutator transaction binding the contract method 0x73e3c280.
//
// Solidity: function beforeAddShares(address , uint256 ) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) BeforeAddShares(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.BeforeAddShares(&_EigenStrategy.TransactOpts, arg0, arg1)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address , uint256 ) returns()
func (_EigenStrategy *EigenStrategyTransactor) BeforeRemoveShares(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "beforeRemoveShares", arg0, arg1)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address , uint256 ) returns()
func (_EigenStrategy *EigenStrategySession) BeforeRemoveShares(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.BeforeRemoveShares(&_EigenStrategy.TransactOpts, arg0, arg1)
}

// BeforeRemoveShares is a paid mutator transaction binding the contract method 0x03e3e6eb.
//
// Solidity: function beforeRemoveShares(address , uint256 ) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) BeforeRemoveShares(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.BeforeRemoveShares(&_EigenStrategy.TransactOpts, arg0, arg1)
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
