// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractStrategyBase

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

// ContractStrategyBaseMetaData contains all meta data concerning the ContractStrategyBase contract.
var ContractStrategyBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"_strategyManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newShares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"explanation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_underlyingToken\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"sharesToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"sharesToUnderlyingView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyManager\",\"outputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"}],\"name\":\"underlyingToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"}],\"name\":\"underlyingToSharesView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userUnderlyingView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051620018b3380380620018b38339810160408190526100319161010c565b6001600160a01b03811660805261004661004c565b5061013c565b600054610100900460ff16156100b85760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101561010a576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561011e57600080fd5b81516001600160a01b038116811461013557600080fd5b9392505050565b6080516117466200016d6000396000818161019901528181610570015281816109c70152610a9201526117466000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80635c975abb116100b8578063ab5921e11161007c578063ab5921e11461029c578063ce7c2ac2146102b1578063d9caed12146102c4578063e3dae51c146102d7578063f3e73875146102ea578063fabc1cbc146102fd57600080fd5b80635c975abb146102425780637a8b26371461024a578063886f11951461025d5780638c871019146102765780638f6a62401461028957600080fd5b806347e7ef24116100ff57806347e7ef24146101d2578063485cc955146101e5578063553ca5f8146101f8578063595c6a671461020b5780635ac86ab71461021357600080fd5b806310d67a2f1461013c578063136439dd146101515780632495a5991461016457806339b70e38146101945780633a98ef39146101bb575b600080fd5b61014f61014a3660046113d5565b610310565b005b61014f61015f3660046113f2565b6103cc565b603254610177906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101777f000000000000000000000000000000000000000000000000000000000000000081565b6101c460335481565b60405190815260200161018b565b6101c46101e036600461140b565b610510565b61014f6101f3366004611437565b610726565b6101c46102063660046113d5565b61083b565b61014f61084f565b610232610221366004611470565b6001805460ff9092161b9081161490565b604051901515815260200161018b565b6001546101c4565b6101c46102583660046113f2565b61091b565b600054610177906201000090046001600160a01b031681565b6101c46102843660046113f2565b610966565b6101c46102973660046113d5565b610971565b6102a461097f565b60405161018b91906114c3565b6101c46102bf3660046113d5565b61099f565b61014f6102d23660046114f6565b610a34565b6101c46102e53660046113f2565b610c81565b6101c46102f83660046113f2565b610cba565b61014f61030b3660046113f2565b610cc5565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103879190611537565b6001600160a01b0316336001600160a01b0316146103c05760405162461bcd60e51b81526004016103b790611554565b60405180910390fd5b6103c981610e21565b50565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610419573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043d919061159e565b6104595760405162461bcd60e51b81526004016103b7906115c0565b600154818116146104d25760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016103b7565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600180546000918291811614156105655760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b60448201526064016103b7565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105dd5760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e6167657260448201526064016103b7565b6032546001600160a01b038581169116146106595760405162461bcd60e51b815260206004820152603660248201527f5374726174656779426173652e6465706f7369743a2043616e206f6e6c79206460448201527532b837b9b4ba103ab73232b9363cb4b733aa37b5b2b760511b60648201526084016103b7565b603354600061066a6103e88361161e565b905060006103e8610679610f2a565b610683919061161e565b905060006106918783611636565b90508061069e848961164d565b6106a8919061166c565b95508561070e5760405162461bcd60e51b815260206004820152602e60248201527f5374726174656779426173652e6465706f7369743a206e65775368617265732060448201526d63616e6e6f74206265207a65726f60901b60648201526084016103b7565b610718868561161e565b603355505050505092915050565b600054610100900460ff16158080156107465750600054600160ff909116105b806107605750303b158015610760575060005460ff166001145b6107c35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103b7565b6000805460ff1916600117905580156107e6576000805461ff0019166101001790555b6107f08383610f9c565b8015610836576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60006108496102588361099f565b92915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa15801561089c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c0919061159e565b6108dc5760405162461bcd60e51b81526004016103b7906115c0565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000806103e860335461092e919061161e565b905060006103e861093d610f2a565b610947919061161e565b905081610954858361164d565b61095e919061166c565b949350505050565b600061084982610c81565b60006108496102f88361099f565b60606040518060800160405280604d81526020016116c4604d9139905090565b604051633d3f06c960e11b81526001600160a01b0382811660048301523060248301526000917f000000000000000000000000000000000000000000000000000000000000000090911690637a7e0d9290604401602060405180830381865afa158015610a10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610849919061168e565b6001805460029081161415610a875760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b60448201526064016103b7565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610aff5760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e6167657260448201526064016103b7565b6032546001600160a01b03848116911614610b825760405162461bcd60e51b815260206004820152603b60248201527f5374726174656779426173652e77697468647261773a2043616e206f6e6c792060448201527f77697468647261772074686520737472617465677920746f6b656e000000000060648201526084016103b7565b60335480831115610c115760405162461bcd60e51b815260206004820152604d60248201527f5374726174656779426173652e77697468647261773a20616d6f756e7453686160448201527f726573206d757374206265206c657373207468616e206f7220657175616c207460648201526c6f20746f74616c53686172657360981b608482015260a4016103b7565b6000610c1f6103e88361161e565b905060006103e8610c2e610f2a565b610c38919061161e565b9050600082610c47878461164d565b610c51919061166c565b9050610c5d8685611636565b603355603254610c77906001600160a01b0316898361102d565b5050505050505050565b6000806103e8603354610c94919061161e565b905060006103e8610ca3610f2a565b610cad919061161e565b905080610954838661164d565b60006108498261091b565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3c9190611537565b6001600160a01b0316336001600160a01b031614610d6c5760405162461bcd60e51b81526004016103b790611554565b600154198119600154191614610dea5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016103b7565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610505565b6001600160a01b038116610eaf5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016103b7565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b5050565b6032546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa158015610f73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f97919061168e565b905090565b600054610100900460ff166110075760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016103b7565b603280546001600160a01b0319166001600160a01b038416179055610f2681600061107f565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b17905261083690849061116b565b6000546201000090046001600160a01b03161580156110a657506001600160a01b03821615155b6111285760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016103b7565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2610f2682610e21565b60006111c0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661123d9092919063ffffffff16565b80519091501561083657808060200190518101906111de919061159e565b6108365760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016103b7565b606061124c8484600085611256565b90505b9392505050565b6060824710156112b75760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016103b7565b6001600160a01b0385163b61130e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103b7565b600080866001600160a01b0316858760405161132a91906116a7565b60006040518083038185875af1925050503d8060008114611367576040519150601f19603f3d011682016040523d82523d6000602084013e61136c565b606091505b509150915061137c828286611387565b979650505050505050565b6060831561139657508161124f565b8251156113a65782518084602001fd5b8160405162461bcd60e51b81526004016103b791906114c3565b6001600160a01b03811681146103c957600080fd5b6000602082840312156113e757600080fd5b813561124f816113c0565b60006020828403121561140457600080fd5b5035919050565b6000806040838503121561141e57600080fd5b8235611429816113c0565b946020939093013593505050565b6000806040838503121561144a57600080fd5b8235611455816113c0565b91506020830135611465816113c0565b809150509250929050565b60006020828403121561148257600080fd5b813560ff8116811461124f57600080fd5b60005b838110156114ae578181015183820152602001611496565b838111156114bd576000848401525b50505050565b60208152600082518060208401526114e2816040850160208701611493565b601f01601f19169190910160400192915050565b60008060006060848603121561150b57600080fd5b8335611516816113c0565b92506020840135611526816113c0565b929592945050506040919091013590565b60006020828403121561154957600080fd5b815161124f816113c0565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156115b057600080fd5b8151801515811461124f57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b6000821982111561163157611631611608565b500190565b60008282101561164857611648611608565b500390565b600081600019048311821515161561166757611667611608565b500290565b60008261168957634e487b7160e01b600052601260045260246000fd5b500490565b6000602082840312156116a057600080fd5b5051919050565b600082516116b9818460208701611493565b919091019291505056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a26469706673582212204545e3e02ff394336095262550ba11b5af57f807bd01ce5c56f0395f4724df7264736f6c634300080c0033",
}

// ContractStrategyBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractStrategyBaseMetaData.ABI instead.
var ContractStrategyBaseABI = ContractStrategyBaseMetaData.ABI

// ContractStrategyBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractStrategyBaseMetaData.Bin instead.
var ContractStrategyBaseBin = ContractStrategyBaseMetaData.Bin

// DeployContractStrategyBase deploys a new Ethereum contract, binding an instance of ContractStrategyBase to it.
func DeployContractStrategyBase(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address) (common.Address, *types.Transaction, *ContractStrategyBase, error) {
	parsed, err := ContractStrategyBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractStrategyBaseBin), backend, _strategyManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractStrategyBase{ContractStrategyBaseCaller: ContractStrategyBaseCaller{contract: contract}, ContractStrategyBaseTransactor: ContractStrategyBaseTransactor{contract: contract}, ContractStrategyBaseFilterer: ContractStrategyBaseFilterer{contract: contract}}, nil
}

// ContractStrategyBase is an auto generated Go binding around an Ethereum contract.
type ContractStrategyBase struct {
	ContractStrategyBaseCaller     // Read-only binding to the contract
	ContractStrategyBaseTransactor // Write-only binding to the contract
	ContractStrategyBaseFilterer   // Log filterer for contract events
}

// ContractStrategyBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractStrategyBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractStrategyBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractStrategyBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractStrategyBaseSession struct {
	Contract     *ContractStrategyBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContractStrategyBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractStrategyBaseCallerSession struct {
	Contract *ContractStrategyBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContractStrategyBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractStrategyBaseTransactorSession struct {
	Contract     *ContractStrategyBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractStrategyBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractStrategyBaseRaw struct {
	Contract *ContractStrategyBase // Generic contract binding to access the raw methods on
}

// ContractStrategyBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractStrategyBaseCallerRaw struct {
	Contract *ContractStrategyBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ContractStrategyBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractStrategyBaseTransactorRaw struct {
	Contract *ContractStrategyBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractStrategyBase creates a new instance of ContractStrategyBase, bound to a specific deployed contract.
func NewContractStrategyBase(address common.Address, backend bind.ContractBackend) (*ContractStrategyBase, error) {
	contract, err := bindContractStrategyBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBase{ContractStrategyBaseCaller: ContractStrategyBaseCaller{contract: contract}, ContractStrategyBaseTransactor: ContractStrategyBaseTransactor{contract: contract}, ContractStrategyBaseFilterer: ContractStrategyBaseFilterer{contract: contract}}, nil
}

// NewContractStrategyBaseCaller creates a new read-only instance of ContractStrategyBase, bound to a specific deployed contract.
func NewContractStrategyBaseCaller(address common.Address, caller bind.ContractCaller) (*ContractStrategyBaseCaller, error) {
	contract, err := bindContractStrategyBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBaseCaller{contract: contract}, nil
}

// NewContractStrategyBaseTransactor creates a new write-only instance of ContractStrategyBase, bound to a specific deployed contract.
func NewContractStrategyBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractStrategyBaseTransactor, error) {
	contract, err := bindContractStrategyBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBaseTransactor{contract: contract}, nil
}

// NewContractStrategyBaseFilterer creates a new log filterer instance of ContractStrategyBase, bound to a specific deployed contract.
func NewContractStrategyBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractStrategyBaseFilterer, error) {
	contract, err := bindContractStrategyBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBaseFilterer{contract: contract}, nil
}

// bindContractStrategyBase binds a generic wrapper to an already deployed contract.
func bindContractStrategyBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractStrategyBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyBase *ContractStrategyBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyBase.Contract.ContractStrategyBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyBase *ContractStrategyBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.ContractStrategyBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyBase *ContractStrategyBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.ContractStrategyBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyBase *ContractStrategyBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyBase *ContractStrategyBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyBase *ContractStrategyBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_ContractStrategyBase *ContractStrategyBaseCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_ContractStrategyBase *ContractStrategyBaseSession) Explanation() (string, error) {
	return _ContractStrategyBase.Contract.Explanation(&_ContractStrategyBase.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) Explanation() (string, error) {
	return _ContractStrategyBase.Contract.Explanation(&_ContractStrategyBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyBase *ContractStrategyBaseCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyBase *ContractStrategyBaseSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyBase.Contract.Paused(&_ContractStrategyBase.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyBase.Contract.Paused(&_ContractStrategyBase.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) Paused0() (*big.Int, error) {
	return _ContractStrategyBase.Contract.Paused0(&_ContractStrategyBase.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) Paused0() (*big.Int, error) {
	return _ContractStrategyBase.Contract.Paused0(&_ContractStrategyBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyBase.Contract.PauserRegistry(&_ContractStrategyBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyBase.Contract.PauserRegistry(&_ContractStrategyBase.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) Shares(user common.Address) (*big.Int, error) {
	return _ContractStrategyBase.Contract.Shares(&_ContractStrategyBase.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _ContractStrategyBase.Contract.Shares(&_ContractStrategyBase.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.SharesToUnderlying(&_ContractStrategyBase.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.SharesToUnderlying(&_ContractStrategyBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.SharesToUnderlyingView(&_ContractStrategyBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.SharesToUnderlyingView(&_ContractStrategyBase.CallOpts, amountShares)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseSession) StrategyManager() (common.Address, error) {
	return _ContractStrategyBase.Contract.StrategyManager(&_ContractStrategyBase.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) StrategyManager() (common.Address, error) {
	return _ContractStrategyBase.Contract.StrategyManager(&_ContractStrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) TotalShares() (*big.Int, error) {
	return _ContractStrategyBase.Contract.TotalShares(&_ContractStrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) TotalShares() (*big.Int, error) {
	return _ContractStrategyBase.Contract.TotalShares(&_ContractStrategyBase.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.UnderlyingToShares(&_ContractStrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.UnderlyingToShares(&_ContractStrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.UnderlyingToSharesView(&_ContractStrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractStrategyBase.Contract.UnderlyingToSharesView(&_ContractStrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseSession) UnderlyingToken() (common.Address, error) {
	return _ContractStrategyBase.Contract.UnderlyingToken(&_ContractStrategyBase.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) UnderlyingToken() (common.Address, error) {
	return _ContractStrategyBase.Contract.UnderlyingToken(&_ContractStrategyBase.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyBase.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _ContractStrategyBase.Contract.UserUnderlyingView(&_ContractStrategyBase.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _ContractStrategyBase.Contract.UserUnderlyingView(&_ContractStrategyBase.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_ContractStrategyBase *ContractStrategyBaseTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_ContractStrategyBase *ContractStrategyBaseSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Deposit(&_ContractStrategyBase.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Deposit(&_ContractStrategyBase.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactor) Initialize(opts *bind.TransactOpts, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "initialize", _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_ContractStrategyBase *ContractStrategyBaseSession) Initialize(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Initialize(&_ContractStrategyBase.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) Initialize(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Initialize(&_ContractStrategyBase.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyBase *ContractStrategyBaseSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Pause(&_ContractStrategyBase.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Pause(&_ContractStrategyBase.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyBase *ContractStrategyBaseSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.PauseAll(&_ContractStrategyBase.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.PauseAll(&_ContractStrategyBase.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractStrategyBase *ContractStrategyBaseSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.SetPauserRegistry(&_ContractStrategyBase.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.SetPauserRegistry(&_ContractStrategyBase.TransactOpts, newPauserRegistry)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyBase *ContractStrategyBaseSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Unpause(&_ContractStrategyBase.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Unpause(&_ContractStrategyBase.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.UserUnderlying(&_ContractStrategyBase.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.UserUnderlying(&_ContractStrategyBase.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactor) Withdraw(opts *bind.TransactOpts, depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.contract.Transact(opts, "withdraw", depositor, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_ContractStrategyBase *ContractStrategyBaseSession) Withdraw(depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Withdraw(&_ContractStrategyBase.TransactOpts, depositor, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_ContractStrategyBase *ContractStrategyBaseTransactorSession) Withdraw(depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyBase.Contract.Withdraw(&_ContractStrategyBase.TransactOpts, depositor, token, amountShares)
}

// ContractStrategyBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractStrategyBase contract.
type ContractStrategyBaseInitializedIterator struct {
	Event *ContractStrategyBaseInitialized // Event containing the contract specifics and raw log

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
func (it *ContractStrategyBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyBaseInitialized)
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
		it.Event = new(ContractStrategyBaseInitialized)
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
func (it *ContractStrategyBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyBaseInitialized represents a Initialized event raised by the ContractStrategyBase contract.
type ContractStrategyBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractStrategyBaseInitializedIterator, error) {

	logs, sub, err := _ContractStrategyBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBaseInitializedIterator{contract: _ContractStrategyBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractStrategyBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyBaseInitialized)
				if err := _ContractStrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractStrategyBase *ContractStrategyBaseFilterer) ParseInitialized(log types.Log) (*ContractStrategyBaseInitialized, error) {
	event := new(ContractStrategyBaseInitialized)
	if err := _ContractStrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractStrategyBase contract.
type ContractStrategyBasePausedIterator struct {
	Event *ContractStrategyBasePaused // Event containing the contract specifics and raw log

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
func (it *ContractStrategyBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyBasePaused)
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
		it.Event = new(ContractStrategyBasePaused)
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
func (it *ContractStrategyBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyBasePaused represents a Paused event raised by the ContractStrategyBase contract.
type ContractStrategyBasePaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyBasePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyBase.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBasePausedIterator{contract: _ContractStrategyBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyBasePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyBase.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyBasePaused)
				if err := _ContractStrategyBase.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractStrategyBase *ContractStrategyBaseFilterer) ParsePaused(log types.Log) (*ContractStrategyBasePaused, error) {
	event := new(ContractStrategyBasePaused)
	if err := _ContractStrategyBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyBasePauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractStrategyBase contract.
type ContractStrategyBasePauserRegistrySetIterator struct {
	Event *ContractStrategyBasePauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractStrategyBasePauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyBasePauserRegistrySet)
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
		it.Event = new(ContractStrategyBasePauserRegistrySet)
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
func (it *ContractStrategyBasePauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyBasePauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyBasePauserRegistrySet represents a PauserRegistrySet event raised by the ContractStrategyBase contract.
type ContractStrategyBasePauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractStrategyBasePauserRegistrySetIterator, error) {

	logs, sub, err := _ContractStrategyBase.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBasePauserRegistrySetIterator{contract: _ContractStrategyBase.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractStrategyBasePauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyBase.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyBasePauserRegistrySet)
				if err := _ContractStrategyBase.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) ParsePauserRegistrySet(log types.Log) (*ContractStrategyBasePauserRegistrySet, error) {
	event := new(ContractStrategyBasePauserRegistrySet)
	if err := _ContractStrategyBase.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractStrategyBase contract.
type ContractStrategyBaseUnpausedIterator struct {
	Event *ContractStrategyBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractStrategyBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyBaseUnpaused)
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
		it.Event = new(ContractStrategyBaseUnpaused)
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
func (it *ContractStrategyBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyBaseUnpaused represents a Unpaused event raised by the ContractStrategyBase contract.
type ContractStrategyBaseUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyBaseUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyBase.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyBaseUnpausedIterator{contract: _ContractStrategyBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyBase *ContractStrategyBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyBaseUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyBase.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyBaseUnpaused)
				if err := _ContractStrategyBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractStrategyBase *ContractStrategyBaseFilterer) ParseUnpaused(log types.Log) (*ContractStrategyBaseUnpaused, error) {
	event := new(ContractStrategyBaseUnpaused)
	if err := _ContractStrategyBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
