// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyBaseTVLLimits

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

// StrategyBaseTVLLimitsMetaData contains all meta data concerning the StrategyBaseTVLLimits contract.
var StrategyBaseTVLLimitsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTVLLimits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_maxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxTotalDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxPerDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTotalDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setTVLLimits\",\"inputs\":[{\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newMaxTotalDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPerDepositUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxTotalDepositsUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516115f73803806115f783398101604081905261002e9161014f565b8181806001600160a01b038116610058576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03908116608052821660a05261007361007c565b50505050610187565b5f54610100900460ff16156100e75760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610136575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461014c575f5ffd5b50565b5f5f60408385031215610160575f5ffd5b825161016b81610138565b602084015190925061017c81610138565b809150509250929050565b60805160a05161142c6101cb5f395f81816101ce015281816103f50152818161077d01526107ff01525f81816102960152818161099a0152610b02015261142c5ff3fe608060405234801561000f575f5ffd5b506004361061016d575f3560e01c80637a8b2637116100d9578063c4d66de811610093578063df6fadc11161006e578063df6fadc11461033f578063e3dae51c1461035a578063f3e738751461036d578063fabc1cbc14610380575f5ffd5b8063c4d66de814610306578063ce7c2ac214610319578063d9caed121461032c575f5ffd5b80637a8b26371461027e578063886f1195146102915780638c871019146102b85780638f6a6240146102cb578063a6ab36f2146102de578063ab5921e1146102f1575f5ffd5b806347e7ef241161012a57806347e7ef2414610210578063553ca5f814610223578063595c6a67146102365780635ac86ab71461023e5780635c975abb1461026d57806361b01b5d14610275575f5ffd5b806311c70c9d14610171578063136439dd146101865780632495a5991461019957806339b70e38146101c95780633a98ef39146101f057806343fe08b014610207575b5f5ffd5b61018461017f366004611100565b610393565b005b610184610194366004611120565b6103a9565b6032546101ac906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101ac7f000000000000000000000000000000000000000000000000000000000000000081565b6101f960335481565b6040519081526020016101c0565b6101f960645481565b6101f961021e36600461114b565b6103df565b6101f9610231366004611175565b61050e565b610184610521565b61025d61024c3660046111a5565b6001805460ff9092161b9081161490565b60405190151581526020016101c0565b6001546101f9565b6101f960655481565b6101f961028c366004611120565b610535565b6101ac7f000000000000000000000000000000000000000000000000000000000000000081565b6101f96102c6366004611120565b61057e565b6101f96102d9366004611175565b610588565b6101846102ec3660046111c0565b610595565b6102f9610670565b6040516101c091906111f6565b610184610314366004611175565b610690565b6101f9610327366004611175565b610756565b6101f961033a36600461122b565b6107e8565b606454606554604080519283526020830191909152016101c0565b6101f9610368366004611120565b6108ea565b6101f961037b366004611120565b610921565b61018461038e366004611120565b61092b565b61039b610998565b6103a58282610a49565b5050565b6103b1610aed565b60015481811681146103d65760405163c61dca5d60e01b815260040160405180910390fd5b6103a582610b90565b5f5f6103ea81610bcd565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610433576040516348da714f60e01b815260040160405180910390fd5b61043d8484610c03565b6033545f61044d6103e88361127d565b90505f6103e861045b610c5a565b610465919061127d565b90505f6104728783611290565b90508061047f84896112a3565b61048991906112ba565b9550855f036104ab57604051630c392ed360e11b815260040160405180910390fd5b6104b5868561127d565b60338190556f4b3b4ca85a86c47a098a223fffffffff10156104ea57604051632f14e8a360e11b815260040160405180910390fd5b610503826103e86033546104fe919061127d565b610cc9565b505050505092915050565b5f61051b61028c83610756565b92915050565b610529610aed565b6105335f19610b90565b565b5f5f6103e8603354610547919061127d565b90505f6103e8610555610c5a565b61055f919061127d565b90508161056c85836112a3565b61057691906112ba565b949350505050565b5f61051b826108ea565b5f61051b61037b83610756565b5f54610100900460ff16158080156105b357505f54600160ff909116105b806105cc5750303b1580156105cc57505f5460ff166001145b6105f15760405162461bcd60e51b81526004016105e8906112d9565b60405180910390fd5b5f805460ff191660011790558015610612575f805461ff0019166101001790555b61061c8484610a49565b61062582610d15565b801561066a575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60606040518060800160405280604d81526020016113aa604d9139905090565b5f54610100900460ff16158080156106ae57505f54600160ff909116105b806106c75750303b1580156106c757505f5460ff166001145b6106e35760405162461bcd60e51b81526004016105e8906112d9565b5f805460ff191660011790558015610704575f805461ff0019166101001790555b61070d82610d15565b80156103a5575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15050565b60405163fe243a1760e01b81526001600160a01b0382811660048301523060248301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063fe243a1790604401602060405180830381865afa1580156107c4573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061051b9190611327565b5f60016107f481610bcd565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461083d576040516348da714f60e01b815260040160405180910390fd5b610848858585610e60565b6033548084111561086c57604051630b469df360e41b815260040160405180910390fd5b5f6108796103e88361127d565b90505f6103e8610887610c5a565b610891919061127d565b90508161089e87836112a3565b6108a891906112ba565b94506108b48684611290565b6033556108d46108c48683611290565b6103e86033546104fe919061127d565b6108df888887610e93565b505050509392505050565b5f5f6103e86033546108fc919061127d565b90505f6103e861090a610c5a565b610914919061127d565b90508061056c83866112a3565b5f61051b82610535565b610933610998565b6001548019821981161461095a5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109f4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a18919061133e565b6001600160a01b0316336001600160a01b0316146105335760405163794821ff60e01b815260040160405180910390fd5b60645460408051918252602082018490527ff97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5910160405180910390a160655460408051918252602082018390527f6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452910160405180910390a180821115610ae25760405163052b07b760e21b815260040160405180910390fd5b606491909155606555565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610b4f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b739190611359565b61053357604051631d77d47760e21b815260040160405180910390fd5b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b610be2816001805460ff9092161b9081161490565b15610c005760405163840a48d560e01b815260040160405180910390fd5b50565b606454811115610c265760405163052b07b760e21b815260040160405180910390fd5b606554610c31610c5a565b1115610c505760405163d86bae6760e01b815260040160405180910390fd5b6103a58282610ea7565b6032546040516370a0823160e01b81523060048201525f916001600160a01b0316906370a0823190602401602060405180830381865afa158015610ca0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cc49190611327565b905090565b7fd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be881610cfd84670de0b6b3a76400006112a3565b610d0791906112ba565b60405190815260200161074a565b5f54610100900460ff16610d7f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016105e8565b603280546001600160a01b0319166001600160a01b038316179055610da35f610b90565b7f1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af55750760325f9054906101000a90046001600160a01b0316826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e15573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e399190611378565b604080516001600160a01b03909316835260ff90911660208301520160405180910390a150565b6032546001600160a01b03838116911614610e8e57604051630312abdd60e61b815260040160405180910390fd5b505050565b610e8e6001600160a01b0383168483610ed5565b6032546001600160a01b038381169116146103a557604051630312abdd60e61b815260040160405180910390fd5b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610e8e928692915f91610f64918516908490610fe3565b905080515f1480610f84575080806020019051810190610f849190611359565b610e8e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016105e8565b606061057684845f85855f5f866001600160a01b031685876040516110089190611393565b5f6040518083038185875af1925050503d805f8114611042576040519150601f19603f3d011682016040523d82523d5f602084013e611047565b606091505b509150915061105887838387611063565b979650505050505050565b606083156110d15782515f036110ca576001600160a01b0385163b6110ca5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016105e8565b5081610576565b61057683838151156110e65781518083602001fd5b8060405162461bcd60e51b81526004016105e891906111f6565b5f5f60408385031215611111575f5ffd5b50508035926020909101359150565b5f60208284031215611130575f5ffd5b5035919050565b6001600160a01b0381168114610c00575f5ffd5b5f5f6040838503121561115c575f5ffd5b823561116781611137565b946020939093013593505050565b5f60208284031215611185575f5ffd5b813561119081611137565b9392505050565b60ff81168114610c00575f5ffd5b5f602082840312156111b5575f5ffd5b813561119081611197565b5f5f5f606084860312156111d2575f5ffd5b833592506020840135915060408401356111eb81611137565b809150509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f6060848603121561123d575f5ffd5b833561124881611137565b9250602084013561125881611137565b929592945050506040919091013590565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561051b5761051b611269565b8181038181111561051b5761051b611269565b808202811582820484141761051b5761051b611269565b5f826112d457634e487b7160e01b5f52601260045260245ffd5b500490565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b5f60208284031215611337575f5ffd5b5051919050565b5f6020828403121561134e575f5ffd5b815161119081611137565b5f60208284031215611369575f5ffd5b81518015158114611190575f5ffd5b5f60208284031215611388575f5ffd5b815161119081611197565b5f82518060208501845e5f92019182525091905056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a2646970667358221220d3746c383368871009e44f7ae962ddb31f64bef2f69a89f96760599304322c1c64736f6c634300081e0033",
}

// StrategyBaseTVLLimitsABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyBaseTVLLimitsMetaData.ABI instead.
var StrategyBaseTVLLimitsABI = StrategyBaseTVLLimitsMetaData.ABI

// StrategyBaseTVLLimitsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyBaseTVLLimitsMetaData.Bin instead.
var StrategyBaseTVLLimitsBin = StrategyBaseTVLLimitsMetaData.Bin

// DeployStrategyBaseTVLLimits deploys a new Ethereum contract, binding an instance of StrategyBaseTVLLimits to it.
func DeployStrategyBaseTVLLimits(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *StrategyBaseTVLLimits, error) {
	parsed, err := StrategyBaseTVLLimitsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyBaseTVLLimitsBin), backend, _strategyManager, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyBaseTVLLimits{StrategyBaseTVLLimitsCaller: StrategyBaseTVLLimitsCaller{contract: contract}, StrategyBaseTVLLimitsTransactor: StrategyBaseTVLLimitsTransactor{contract: contract}, StrategyBaseTVLLimitsFilterer: StrategyBaseTVLLimitsFilterer{contract: contract}}, nil
}

// StrategyBaseTVLLimits is an auto generated Go binding around an Ethereum contract.
type StrategyBaseTVLLimits struct {
	StrategyBaseTVLLimitsCaller     // Read-only binding to the contract
	StrategyBaseTVLLimitsTransactor // Write-only binding to the contract
	StrategyBaseTVLLimitsFilterer   // Log filterer for contract events
}

// StrategyBaseTVLLimitsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTVLLimitsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTVLLimitsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyBaseTVLLimitsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTVLLimitsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyBaseTVLLimitsSession struct {
	Contract     *StrategyBaseTVLLimits // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StrategyBaseTVLLimitsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyBaseTVLLimitsCallerSession struct {
	Contract *StrategyBaseTVLLimitsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// StrategyBaseTVLLimitsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyBaseTVLLimitsTransactorSession struct {
	Contract     *StrategyBaseTVLLimitsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// StrategyBaseTVLLimitsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsRaw struct {
	Contract *StrategyBaseTVLLimits // Generic contract binding to access the raw methods on
}

// StrategyBaseTVLLimitsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsCallerRaw struct {
	Contract *StrategyBaseTVLLimitsCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyBaseTVLLimitsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsTransactorRaw struct {
	Contract *StrategyBaseTVLLimitsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyBaseTVLLimits creates a new instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimits(address common.Address, backend bind.ContractBackend) (*StrategyBaseTVLLimits, error) {
	contract, err := bindStrategyBaseTVLLimits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimits{StrategyBaseTVLLimitsCaller: StrategyBaseTVLLimitsCaller{contract: contract}, StrategyBaseTVLLimitsTransactor: StrategyBaseTVLLimitsTransactor{contract: contract}, StrategyBaseTVLLimitsFilterer: StrategyBaseTVLLimitsFilterer{contract: contract}}, nil
}

// NewStrategyBaseTVLLimitsCaller creates a new read-only instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimitsCaller(address common.Address, caller bind.ContractCaller) (*StrategyBaseTVLLimitsCaller, error) {
	contract, err := bindStrategyBaseTVLLimits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsCaller{contract: contract}, nil
}

// NewStrategyBaseTVLLimitsTransactor creates a new write-only instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimitsTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyBaseTVLLimitsTransactor, error) {
	contract, err := bindStrategyBaseTVLLimits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsTransactor{contract: contract}, nil
}

// NewStrategyBaseTVLLimitsFilterer creates a new log filterer instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimitsFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyBaseTVLLimitsFilterer, error) {
	contract, err := bindStrategyBaseTVLLimits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsFilterer{contract: contract}, nil
}

// bindStrategyBaseTVLLimits binds a generic wrapper to an already deployed contract.
func bindStrategyBaseTVLLimits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyBaseTVLLimitsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBaseTVLLimits.Contract.StrategyBaseTVLLimitsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyBaseTVLLimitsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyBaseTVLLimitsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBaseTVLLimits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Explanation() (string, error) {
	return _StrategyBaseTVLLimits.Contract.Explanation(&_StrategyBaseTVLLimits.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Explanation() (string, error) {
	return _StrategyBaseTVLLimits.Contract.Explanation(&_StrategyBaseTVLLimits.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) GetTVLLimits(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "getTVLLimits")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.GetTVLLimits(&_StrategyBaseTVLLimits.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.GetTVLLimits(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) MaxPerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "maxPerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) MaxPerDeposit() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxPerDeposit(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) MaxPerDeposit() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxPerDeposit(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) MaxTotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "maxTotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) MaxTotalDeposits() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxTotalDeposits(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) MaxTotalDeposits() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxTotalDeposits(&_StrategyBaseTVLLimits.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Paused(index uint8) (bool, error) {
	return _StrategyBaseTVLLimits.Contract.Paused(&_StrategyBaseTVLLimits.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Paused(index uint8) (bool, error) {
	return _StrategyBaseTVLLimits.Contract.Paused(&_StrategyBaseTVLLimits.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Paused0() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Paused0(&_StrategyBaseTVLLimits.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Paused0() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Paused0(&_StrategyBaseTVLLimits.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) PauserRegistry() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.PauserRegistry(&_StrategyBaseTVLLimits.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) PauserRegistry() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.PauserRegistry(&_StrategyBaseTVLLimits.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Shares(&_StrategyBaseTVLLimits.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Shares(&_StrategyBaseTVLLimits.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlying(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlying(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) StrategyManager() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyManager(&_StrategyBaseTVLLimits.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) StrategyManager() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyManager(&_StrategyBaseTVLLimits.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) TotalShares() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.TotalShares(&_StrategyBaseTVLLimits.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) TotalShares() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.TotalShares(&_StrategyBaseTVLLimits.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToShares(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToShares(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToSharesView(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToSharesView(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToken(&_StrategyBaseTVLLimits.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToken(&_StrategyBaseTVLLimits.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Deposit(&_StrategyBaseTVLLimits.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Deposit(&_StrategyBaseTVLLimits.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Initialize(opts *bind.TransactOpts, _maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "initialize", _maxPerDeposit, _maxTotalDeposits, _underlyingToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize(&_StrategyBaseTVLLimits.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize(&_StrategyBaseTVLLimits.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Initialize0(opts *bind.TransactOpts, _underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "initialize0", _underlyingToken)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Initialize0(_underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize0(&_StrategyBaseTVLLimits.TransactOpts, _underlyingToken)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _underlyingToken) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Initialize0(_underlyingToken common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize0(&_StrategyBaseTVLLimits.TransactOpts, _underlyingToken)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Pause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Pause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.PauseAll(&_StrategyBaseTVLLimits.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.PauseAll(&_StrategyBaseTVLLimits.TransactOpts)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) SetTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "setTVLLimits", newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.SetTVLLimits(&_StrategyBaseTVLLimits.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.SetTVLLimits(&_StrategyBaseTVLLimits.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Unpause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Unpause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlying(&_StrategyBaseTVLLimits.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlying(&_StrategyBaseTVLLimits.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Withdraw(&_StrategyBaseTVLLimits.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns(uint256 amountOut)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Withdraw(&_StrategyBaseTVLLimits.TransactOpts, recipient, token, amountShares)
}

// StrategyBaseTVLLimitsExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsExchangeRateEmittedIterator struct {
	Event *StrategyBaseTVLLimitsExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsExchangeRateEmitted)
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
		it.Event = new(StrategyBaseTVLLimitsExchangeRateEmitted)
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
func (it *StrategyBaseTVLLimitsExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsExchangeRateEmitted represents a ExchangeRateEmitted event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsExchangeRateEmittedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsExchangeRateEmittedIterator{contract: _StrategyBaseTVLLimits.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsExchangeRateEmitted)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseExchangeRateEmitted(log types.Log) (*StrategyBaseTVLLimitsExchangeRateEmitted, error) {
	event := new(StrategyBaseTVLLimitsExchangeRateEmitted)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsInitializedIterator struct {
	Event *StrategyBaseTVLLimitsInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsInitialized)
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
		it.Event = new(StrategyBaseTVLLimitsInitialized)
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
func (it *StrategyBaseTVLLimitsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsInitialized represents a Initialized event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsInitializedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsInitializedIterator{contract: _StrategyBaseTVLLimits.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsInitialized)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseInitialized(log types.Log) (*StrategyBaseTVLLimitsInitialized, error) {
	event := new(StrategyBaseTVLLimitsInitialized)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator is returned from FilterMaxPerDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxPerDepositUpdated events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator struct {
	Event *StrategyBaseTVLLimitsMaxPerDepositUpdated // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
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
		it.Event = new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
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
func (it *StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsMaxPerDepositUpdated represents a MaxPerDepositUpdated event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxPerDepositUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxPerDepositUpdated is a free log retrieval operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterMaxPerDepositUpdated(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator{contract: _StrategyBaseTVLLimits.contract, event: "MaxPerDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPerDepositUpdated is a free log subscription operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchMaxPerDepositUpdated(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsMaxPerDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
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

// ParseMaxPerDepositUpdated is a log parse operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseMaxPerDepositUpdated(log types.Log) (*StrategyBaseTVLLimitsMaxPerDepositUpdated, error) {
	event := new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator is returned from FilterMaxTotalDepositsUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalDepositsUpdated events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator struct {
	Event *StrategyBaseTVLLimitsMaxTotalDepositsUpdated // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
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
		it.Event = new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
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
func (it *StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsMaxTotalDepositsUpdated represents a MaxTotalDepositsUpdated event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxTotalDepositsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalDepositsUpdated is a free log retrieval operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterMaxTotalDepositsUpdated(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator{contract: _StrategyBaseTVLLimits.contract, event: "MaxTotalDepositsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalDepositsUpdated is a free log subscription operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchMaxTotalDepositsUpdated(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsMaxTotalDepositsUpdated) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
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

// ParseMaxTotalDepositsUpdated is a log parse operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseMaxTotalDepositsUpdated(log types.Log) (*StrategyBaseTVLLimitsMaxTotalDepositsUpdated, error) {
	event := new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsPausedIterator struct {
	Event *StrategyBaseTVLLimitsPaused // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsPaused)
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
		it.Event = new(StrategyBaseTVLLimitsPaused)
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
func (it *StrategyBaseTVLLimitsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsPaused represents a Paused event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBaseTVLLimitsPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsPausedIterator{contract: _StrategyBaseTVLLimits.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsPaused)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParsePaused(log types.Log) (*StrategyBaseTVLLimitsPaused, error) {
	event := new(StrategyBaseTVLLimitsPaused)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsStrategyTokenSetIterator struct {
	Event *StrategyBaseTVLLimitsStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsStrategyTokenSet)
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
		it.Event = new(StrategyBaseTVLLimitsStrategyTokenSet)
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
func (it *StrategyBaseTVLLimitsStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsStrategyTokenSet represents a StrategyTokenSet event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsStrategyTokenSetIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsStrategyTokenSetIterator{contract: _StrategyBaseTVLLimits.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsStrategyTokenSet)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseStrategyTokenSet(log types.Log) (*StrategyBaseTVLLimitsStrategyTokenSet, error) {
	event := new(StrategyBaseTVLLimitsStrategyTokenSet)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsUnpausedIterator struct {
	Event *StrategyBaseTVLLimitsUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsUnpaused)
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
		it.Event = new(StrategyBaseTVLLimitsUnpaused)
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
func (it *StrategyBaseTVLLimitsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsUnpaused represents a Unpaused event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBaseTVLLimitsUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsUnpausedIterator{contract: _StrategyBaseTVLLimits.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsUnpaused)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseUnpaused(log types.Log) (*StrategyBaseTVLLimitsUnpaused, error) {
	event := new(StrategyBaseTVLLimitsUnpaused)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
