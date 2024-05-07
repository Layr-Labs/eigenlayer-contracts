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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigen\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_EIGEN\",\"type\":\"address\",\"internalType\":\"contractIEigen\"},{\"name\":\"_bEIGEN\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001bd238038062001bd2833981016040819052620000349162000116565b6001600160a01b038116608052806200004c62000054565b505062000148565b600054610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000114576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012957600080fd5b81516001600160a01b03811681146200014157600080fd5b9392505050565b608051611a5962000179600039600081816101af015281816105ac01528181610a340152610aff0152611a596000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c80637a8b2637116100c3578063ce7c2ac21161007c578063ce7c2ac2146102da578063d9caed12146102ed578063e3dae51c14610300578063f3e7387514610313578063fabc1cbc14610326578063fdc371ce1461033957600080fd5b80637a8b263714610260578063886f1195146102735780638c8710191461028c5780638f6a62401461029f578063ab5921e1146102b2578063c0c53b8b146102c757600080fd5b806347e7ef241161011557806347e7ef24146101e8578063485cc955146101fb578063553ca5f81461020e578063595c6a67146102215780635ac86ab7146102295780635c975abb1461025857600080fd5b806310d67a2f14610152578063136439dd146101675780632495a5991461017a57806339b70e38146101aa5780633a98ef39146101d1575b600080fd5b610165610160366004611653565b61034c565b005b610165610175366004611670565b610408565b60325461018d906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61018d7f000000000000000000000000000000000000000000000000000000000000000081565b6101da60335481565b6040519081526020016101a1565b6101da6101f6366004611689565b61054c565b6101656102093660046116b5565b6106f0565b6101da61021c366004611653565b6107be565b6101656107d2565b6102486102373660046116ee565b6001805460ff9092161b9081161490565b60405190151581526020016101a1565b6001546101da565b6101da61026e366004611670565b61089e565b60005461018d906201000090046001600160a01b031681565b6101da61029a366004611670565b6108e9565b6101da6102ad366004611653565b6108f4565b6102ba610902565b6040516101a1919061173d565b6101656102d5366004611770565b610922565b6101da6102e8366004611653565b610a0c565b6101656102fb3660046117bb565b610aa1565b6101da61030e366004611670565b610c6a565b6101da610321366004611670565b610ca3565b610165610334366004611670565b610cae565b60645461018d906001600160a01b031681565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561039f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c391906117fc565b6001600160a01b0316336001600160a01b0316146103fc5760405162461bcd60e51b81526004016103f390611819565b60405180910390fd5b61040581610e0a565b50565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610455573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104799190611863565b6104955760405162461bcd60e51b81526004016103f390611885565b6001548181161461050e5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016103f3565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600180546000918291811614156105a15760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b60448201526064016103f3565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106195760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e6167657260448201526064016103f3565b6106238484610f0f565b60335460006106346103e8836118e3565b905060006103e8610643611023565b61064d91906118e3565b9050600061065b87836118fb565b9050806106688489611912565b6106729190611931565b9550856106d85760405162461bcd60e51b815260206004820152602e60248201527f5374726174656779426173652e6465706f7369743a206e65775368617265732060448201526d63616e6e6f74206265207a65726f60901b60648201526084016103f3565b6106e286856118e3565b603355505050505092915050565b600054610100900460ff16158080156107105750600054600160ff909116105b8061072a5750303b15801561072a575060005460ff166001145b6107465760405162461bcd60e51b81526004016103f390611953565b6000805460ff191660011790558015610769576000805461ff0019166101001790555b6107738383611095565b80156107b9576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60006107cc61026e83610a0c565b92915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa15801561081f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108439190611863565b61085f5760405162461bcd60e51b81526004016103f390611885565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000806103e86033546108b191906118e3565b905060006103e86108c0611023565b6108ca91906118e3565b9050816108d78583611912565b6108e19190611931565b949350505050565b60006107cc82610c6a565b60006107cc61032183610a0c565b60606040518060800160405280604d81526020016119d7604d9139905090565b600054610100900460ff16158080156109425750600054600160ff909116105b8061095c5750303b15801561095c575060005460ff166001145b6109785760405162461bcd60e51b81526004016103f390611953565b6000805460ff19166001179055801561099b576000805461ff0019166101001790555b606480546001600160a01b0319166001600160a01b0386161790556109c08383611095565b8015610a06576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b604051633d3f06c960e11b81526001600160a01b0382811660048301523060248301526000917f000000000000000000000000000000000000000000000000000000000000000090911690637a7e0d9290604401602060405180830381865afa158015610a7d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cc91906119a1565b6001805460029081161415610af45760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b60448201526064016103f3565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b6c5760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e6167657260448201526064016103f3565b610b77848484611126565b60335480831115610c065760405162461bcd60e51b815260206004820152604d60248201527f5374726174656779426173652e77697468647261773a20616d6f756e7453686160448201527f726573206d757374206265206c657373207468616e206f7220657175616c207460648201526c6f20746f74616c53686172657360981b608482015260a4016103f3565b6000610c146103e8836118e3565b905060006103e8610c23611023565b610c2d91906118e3565b9050600082610c3c8784611912565b610c469190611931565b9050610c5286856118fb565b603355610c608888836111c1565b5050505050505050565b6000806103e8603354610c7d91906118e3565b905060006103e8610c8c611023565b610c9691906118e3565b9050806108d78386611912565b60006107cc8261089e565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2591906117fc565b6001600160a01b0316336001600160a01b031614610d555760405162461bcd60e51b81526004016103f390611819565b600154198119600154191614610dd35760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016103f3565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610541565b6001600160a01b038116610e985760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016103f3565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b6032546001600160a01b0383811691161480610f3857506064546001600160a01b038381169116145b610faa5760405162461bcd60e51b815260206004820152603760248201527f456967656e53747261746567792e6465706f7369743a2043616e206f6e6c792060448201527f6465706f7369742062454947454e206f7220454947454e00000000000000000060648201526084016103f3565b6064546001600160a01b038381169116141561101f57606454604051636f074d1f60e11b8152600481018390526001600160a01b039091169063de0e9a3e90602401600060405180830381600087803b15801561100657600080fd5b505af115801561101a573d6000803e3d6000fd5b505050505b5050565b6032546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa15801561106c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109091906119a1565b905090565b600054610100900460ff166111005760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016103f3565b603280546001600160a01b0319166001600160a01b03841617905561101f8160006112c2565b6032546001600160a01b038381169116148061114f57506064546001600160a01b038381169116145b6107b95760405162461bcd60e51b815260206004820152603960248201527f456967656e53747261746567792e77697468647261773a2043616e206f6e6c7960448201527f2077697468647261772062454947454e206f7220454947454e0000000000000060648201526084016103f3565b6064546001600160a01b03838116911614156112ae5760325460405163095ea7b360e01b81526001600160a01b038481166004830152602482018490529091169063095ea7b3906044016020604051808303816000875af115801561122a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124e9190611863565b50606454604051630ea598cb60e41b8152600481018390526001600160a01b039091169063ea598cb090602401600060405180830381600087803b15801561129557600080fd5b505af11580156112a9573d6000803e3d6000fd5b505050505b6107b96001600160a01b03831684836113ae565b6000546201000090046001600160a01b03161580156112e957506001600160a01b03821615155b61136b5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016103f3565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261101f82610e0a565b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526107b99286929160009161143e9185169084906114bb565b8051909150156107b9578080602001905181019061145c9190611863565b6107b95760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016103f3565b60606114ca84846000856114d4565b90505b9392505050565b6060824710156115355760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016103f3565b6001600160a01b0385163b61158c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103f3565b600080866001600160a01b031685876040516115a891906119ba565b60006040518083038185875af1925050503d80600081146115e5576040519150601f19603f3d011682016040523d82523d6000602084013e6115ea565b606091505b50915091506115fa828286611605565b979650505050505050565b606083156116145750816114cd565b8251156116245782518084602001fd5b8160405162461bcd60e51b81526004016103f3919061173d565b6001600160a01b038116811461040557600080fd5b60006020828403121561166557600080fd5b81356114cd8161163e565b60006020828403121561168257600080fd5b5035919050565b6000806040838503121561169c57600080fd5b82356116a78161163e565b946020939093013593505050565b600080604083850312156116c857600080fd5b82356116d38161163e565b915060208301356116e38161163e565b809150509250929050565b60006020828403121561170057600080fd5b813560ff811681146114cd57600080fd5b60005b8381101561172c578181015183820152602001611714565b83811115610a065750506000910152565b602081526000825180602084015261175c816040850160208701611711565b601f01601f19169190910160400192915050565b60008060006060848603121561178557600080fd5b83356117908161163e565b925060208401356117a08161163e565b915060408401356117b08161163e565b809150509250925092565b6000806000606084860312156117d057600080fd5b83356117db8161163e565b925060208401356117eb8161163e565b929592945050506040919091013590565b60006020828403121561180e57600080fd5b81516114cd8161163e565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561187557600080fd5b815180151581146114cd57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600082198211156118f6576118f66118cd565b500190565b60008282101561190d5761190d6118cd565b500390565b600081600019048311821515161561192c5761192c6118cd565b500290565b60008261194e57634e487b7160e01b600052601260045260246000fd5b500490565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6000602082840312156119b357600080fd5b5051919050565b600082516119cc818460208701611711565b919091019291505056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a264697066735822122091303774325b6a8843b54c6eadd6cb6b215ab5b7b80d3603da29a7c9dd7709c764736f6c634300080c0033",
}

// EigenStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenStrategyMetaData.ABI instead.
var EigenStrategyABI = EigenStrategyMetaData.ABI

// EigenStrategyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EigenStrategyMetaData.Bin instead.
var EigenStrategyBin = EigenStrategyMetaData.Bin

// DeployEigenStrategy deploys a new Ethereum contract, binding an instance of EigenStrategy to it.
func DeployEigenStrategy(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address) (common.Address, *types.Transaction, *EigenStrategy, error) {
	parsed, err := EigenStrategyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EigenStrategyBin), backend, _strategyManager)
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
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_EigenStrategy *EigenStrategyTransactor) Initialize(opts *bind.TransactOpts, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "initialize", _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_EigenStrategy *EigenStrategySession) Initialize(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize(&_EigenStrategy.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) Initialize(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize(&_EigenStrategy.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _EIGEN, address _bEIGEN, address _pauserRegistry) returns()
func (_EigenStrategy *EigenStrategyTransactor) Initialize0(opts *bind.TransactOpts, _EIGEN common.Address, _bEIGEN common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "initialize0", _EIGEN, _bEIGEN, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _EIGEN, address _bEIGEN, address _pauserRegistry) returns()
func (_EigenStrategy *EigenStrategySession) Initialize0(_EIGEN common.Address, _bEIGEN common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize0(&_EigenStrategy.TransactOpts, _EIGEN, _bEIGEN, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _EIGEN, address _bEIGEN, address _pauserRegistry) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) Initialize0(_EIGEN common.Address, _bEIGEN common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Initialize0(&_EigenStrategy.TransactOpts, _EIGEN, _bEIGEN, _pauserRegistry)
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

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_EigenStrategy *EigenStrategyTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_EigenStrategy *EigenStrategySession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.SetPauserRegistry(&_EigenStrategy.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _EigenStrategy.Contract.SetPauserRegistry(&_EigenStrategy.TransactOpts, newPauserRegistry)
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
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_EigenStrategy *EigenStrategyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_EigenStrategy *EigenStrategySession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Withdraw(&_EigenStrategy.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_EigenStrategy *EigenStrategyTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _EigenStrategy.Contract.Withdraw(&_EigenStrategy.TransactOpts, recipient, token, amountShares)
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

// EigenStrategyPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the EigenStrategy contract.
type EigenStrategyPauserRegistrySetIterator struct {
	Event *EigenStrategyPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *EigenStrategyPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenStrategyPauserRegistrySet)
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
		it.Event = new(EigenStrategyPauserRegistrySet)
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
func (it *EigenStrategyPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenStrategyPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenStrategyPauserRegistrySet represents a PauserRegistrySet event raised by the EigenStrategy contract.
type EigenStrategyPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_EigenStrategy *EigenStrategyFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*EigenStrategyPauserRegistrySetIterator, error) {

	logs, sub, err := _EigenStrategy.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &EigenStrategyPauserRegistrySetIterator{contract: _EigenStrategy.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_EigenStrategy *EigenStrategyFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *EigenStrategyPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _EigenStrategy.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenStrategyPauserRegistrySet)
				if err := _EigenStrategy.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_EigenStrategy *EigenStrategyFilterer) ParsePauserRegistrySet(log types.Log) (*EigenStrategyPauserRegistrySet, error) {
	event := new(EigenStrategyPauserRegistrySet)
	if err := _EigenStrategy.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
