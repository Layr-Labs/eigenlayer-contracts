// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DelayedWithdrawalRouter

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

// IDelayedWithdrawalRouterDelayedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelayedWithdrawalRouterDelayedWithdrawal struct {
	Amount       *big.Int
	BlockCreated uint32
}

// IDelayedWithdrawalRouterUserDelayedWithdrawals is an auto generated low-level Go binding around an user-defined struct.
type IDelayedWithdrawalRouterUserDelayedWithdrawals struct {
	DelayedWithdrawalsCompleted *big.Int
	DelayedWithdrawals          []IDelayedWithdrawalRouterDelayedWithdrawal
}

// DelayedWithdrawalRouterMetaData contains all meta data concerning the DelayedWithdrawalRouter contract.
var DelayedWithdrawalRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canClaimDelayedWithdrawal\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelayedWithdrawals\",\"inputs\":[{\"name\":\"maxNumberOfDelayedWithdrawalsToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDelayedWithdrawals\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxNumberOfDelayedWithdrawalsToClaim\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDelayedWithdrawal\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClaimableUserDelayedWithdrawals\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserDelayedWithdrawals\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userDelayedWithdrawalByIndex\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userWithdrawals\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelayedWithdrawalRouter.UserDelayedWithdrawals\",\"components\":[{\"name\":\"delayedWithdrawalsCompleted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delayedWithdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelayedWithdrawalRouter.DelayedWithdrawal[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint224\",\"internalType\":\"uint224\"},{\"name\":\"blockCreated\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userWithdrawalsLength\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelayedWithdrawalCreated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelayedWithdrawalsClaimed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountClaimed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"delayedWithdrawalsCompleted\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalDelayBlocksSet\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001f0e38038062001f0e8339810160408190526200003491620001a8565b6001600160a01b038116620000cb5760405162461bcd60e51b815260206004820152604c60248201527f44656c617965645769746864726177616c526f757465722e636f6e737472756360448201527f746f723a205f656967656e506f644d616e616765722063616e6e6f742062652060648201526b7a65726f206164647265737360a01b608482015260a4015b60405180910390fd5b6001600160a01b038116608052620000e2620000e9565b50620001da565b600054610100900460ff1615620001535760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401620000c2565b60005460ff9081161015620001a6576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b600060208284031215620001bb57600080fd5b81516001600160a01b0381168114620001d357600080fd5b9392505050565b608051611d11620001fd600039600081816101fa0152610c000152611d116000f3fe60806040526004361061014b5760003560e01c806385594e58116100b6578063e4f4f8871161006f578063e4f4f887146103cc578063e5db06c014610405578063eb990c5914610425578063ecb7cb1b14610445578063f2fde38b14610472578063fabc1cbc1461049257600080fd5b806385594e5814610317578063886f1195146103445780638da5cb5b14610364578063c0db354c14610382578063ca661c0414610395578063d44e1b76146103ac57600080fd5b806350f73e7c1161010857806350f73e7c14610254578063595c6a67146102785780635ac86ab71461028d5780635c975abb146102cd578063715018a6146102e257806375608896146102f757600080fd5b806310d67a2f14610150578063136439dd146101725780631f39d87f146101925780633e1de008146101c85780634665bcda146101e85780634d50f9a414610234575b600080fd5b34801561015c57600080fd5b5061017061016b36600461196d565b6104b2565b005b34801561017e57600080fd5b5061017061018d366004611991565b61056e565b34801561019e57600080fd5b506101b26101ad36600461196d565b6106ad565b6040516101bf91906119c8565b60405180910390f35b3480156101d457600080fd5b506101b26101e336600461196d565b6108a8565b3480156101f457600080fd5b5061021c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101bf565b34801561024057600080fd5b5061017061024f366004611991565b6109ee565b34801561026057600080fd5b5061026a60c95481565b6040519081526020016101bf565b34801561028457600080fd5b506101706109ff565b34801561029957600080fd5b506102bd6102a8366004611a15565b609854600160ff9092169190911b9081161490565b60405190151581526020016101bf565b3480156102d957600080fd5b5060985461026a565b3480156102ee57600080fd5b50610170610ac6565b34801561030357600080fd5b506102bd610312366004611a38565b610ada565b34801561032357600080fd5b50610337610332366004611a38565b610b5d565b6040516101bf9190611a64565b34801561035057600080fd5b5060975461021c906001600160a01b031681565b34801561037057600080fd5b506033546001600160a01b031661021c565b610170610390366004611a72565b610bdd565b3480156103a157600080fd5b5061026a62034bc081565b3480156103b857600080fd5b506101706103c7366004611991565b610e9d565b3480156103d857600080fd5b5061026a6103e736600461196d565b6001600160a01b0316600090815260ca602052604090206001015490565b34801561041157600080fd5b50610170610420366004611a38565b610f31565b34801561043157600080fd5b50610170610440366004611aab565b610fc6565b34801561045157600080fd5b5061046561046036600461196d565b6110ee565b6040516101bf9190611af1565b34801561047e57600080fd5b5061017061048d36600461196d565b6111a8565b34801561049e57600080fd5b506101706104ad366004611991565b61121e565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610505573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105299190611b47565b6001600160a01b0316336001600160a01b0316146105625760405162461bcd60e51b815260040161055990611b64565b60405180910390fd5b61056b8161137a565b50565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156105b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105da9190611bae565b6105f65760405162461bcd60e51b815260040161055990611bd0565b6098548181161461066f5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610559565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6001600160a01b038116600090815260ca6020526040812080546001909101546060926106da8383611c2e565b90508060005b82811015610786576001600160a01b038716600090815260ca6020526040812060010161070d8388611c45565b8154811061071d5761071d611c5d565b6000918252602091829020604080518082019091529101546001600160e01b0381168252600160e01b900463ffffffff1691810182905260c95490925061076391611c45565b4310156107735781925050610786565b508061077e81611c73565b9150506106e0565b508060008167ffffffffffffffff8111156107a3576107a3611c8e565b6040519080825280602002602001820160405280156107e857816020015b60408051808201909152600080825260208201528152602001906001900390816107c15790505b509050811561089d5760005b8281101561089b576001600160a01b038916600090815260ca602052604090206001016108218289611c45565b8154811061083157610831611c5d565b6000918252602091829020604080518082019091529101546001600160e01b0381168252600160e01b900463ffffffff1691810191909152825183908390811061087d5761087d611c5d565b6020026020010181905250808061089390611c73565b9150506107f4565b505b979650505050505050565b6001600160a01b038116600090815260ca6020526040812080546001909101546060926108d58383611c2e565b905060008167ffffffffffffffff8111156108f2576108f2611c8e565b60405190808252806020026020018201604052801561093757816020015b60408051808201909152600080825260208201528152602001906001900390816109105790505b50905060005b828110156109e4576001600160a01b038716600090815260ca6020526040902060010161096a8287611c45565b8154811061097a5761097a611c5d565b6000918252602091829020604080518082019091529101546001600160e01b0381168252600160e01b900463ffffffff169181019190915282518390839081106109c6576109c6611c5d565b602002602001018190525080806109dc90611c73565b91505061093d565b5095945050505050565b6109f6611471565b61056b816114cb565b60975460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610a47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6b9190611bae565b610a875760405162461bcd60e51b815260040161055990611bd0565b600019609881905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b610ace611471565b610ad86000611593565b565b6001600160a01b038216600090815260ca60205260408120548210801590610b54575060c9546001600160a01b038416600090815260ca60205260409020600101805484908110610b2d57610b2d611c5d565b600091825260209091200154610b509190600160e01b900463ffffffff16611c45565b4310155b90505b92915050565b60408051808201909152600080825260208201526001600160a01b038316600090815260ca60205260409020600101805483908110610b9e57610b9e611c5d565b6000918252602091829020604080518082019091529101546001600160e01b0381168252600160e01b900463ffffffff16918101919091529392505050565b60405163a38406a360e01b81526001600160a01b038084166004830152839133917f0000000000000000000000000000000000000000000000000000000000000000169063a38406a390602401602060405180830381865afa158015610c47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6b9190611b47565b6001600160a01b031614610ce75760405162461bcd60e51b815260206004820152603d60248201527f44656c617965645769746864726177616c526f757465722e6f6e6c794569676560448201527f6e506f643a206e6f7420706f644f776e6572277320456967656e506f640000006064820152608401610559565b60985460009060019081161415610d105760405162461bcd60e51b815260040161055990611ca4565b6001600160a01b038316610da65760405162461bcd60e51b815260206004820152605160248201527f44656c617965645769746864726177616c526f757465722e637265617465446560448201527f6c617965645769746864726177616c3a20726563697069656e742063616e6e6f60648201527074206265207a65726f206164647265737360781b608482015260a401610559565b346001600160e01b03811615610e96576040805180820182526001600160e01b03808416825263ffffffff43811660208085019182526001600160a01b038a16600081815260ca8352968720600190810180548083018255818a5293892088519551909616600160e01b029490961693909317939091019290925593525490917fb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f5991889188918691610e5791611c2e565b604080516001600160a01b0395861681529490931660208501526001600160e01b039091169183019190915260608201526080015b60405180910390a1505b5050505050565b60026065541415610ef05760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610559565b600260655560985460009060019081161415610f1e5760405162461bcd60e51b815260040161055990611ca4565b610f2833836115e5565b50506001606555565b60026065541415610f845760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610559565b600260655560985460009060019081161415610fb25760405162461bcd60e51b815260040161055990611ca4565b610fbc83836115e5565b5050600160655550565b600054610100900460ff1615808015610fe65750600054600160ff909116105b806110005750303b158015611000575060005460ff166001145b6110635760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610559565b6000805460ff191660011790558015611086576000805461ff0019166101001790555b61108f85611593565b6110998484611750565b6110a2826114cb565b8015610e96576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b6040805180820190915260008152606060208201526001600160a01b038216600090815260ca6020908152604080832081518083018352815481526001820180548451818702810187019095528085529195929486810194939192919084015b8282101561119a57600084815260209081902060408051808201909152908401546001600160e01b0381168252600160e01b900463ffffffff168183015282526001909201910161114e565b505050915250909392505050565b6111b0611471565b6001600160a01b0381166112155760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610559565b61056b81611593565b609760009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611271573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112959190611b47565b6001600160a01b0316336001600160a01b0316146112c55760405162461bcd60e51b815260040161055990611b64565b6098541981196098541916146113435760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610559565b609881905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016106a2565b6001600160a01b0381166114085760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610559565b609754604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1609780546001600160a01b0319166001600160a01b0392909216919091179055565b6033546001600160a01b03163314610ad85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610559565b62034bc08111156115525760405162461bcd60e51b815260206004820152604560248201527f44656c617965645769746864726177616c526f757465722e5f7365745769746860448201527f64726177616c44656c6179426c6f636b733a206e657756616c756520746f6f206064820152646c6172676560d81b608482015260a401610559565b60c95460408051918252602082018390527f4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e910160405180910390a160c955565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038216600090815260ca602052604081208054600190910154825b848110801561161e57508161161c8285611c45565b105b156116cb576001600160a01b038616600090815260ca602052604081206001016116488386611c45565b8154811061165857611658611c5d565b6000918252602091829020604080518082019091529101546001600160e01b0381168252600160e01b900463ffffffff1691810182905260c95490925061169e91611c45565b4310156116ab57506116cb565b80516116c0906001600160e01b031686611c45565b945050600101611607565b6116d58184611c45565b6001600160a01b038716600090815260ca602052604090205583156116fe576116fe868561183a565b7f6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943868561172b8487611c45565b604080516001600160a01b039094168452602084019290925290820152606001610e8c565b6097546001600160a01b031615801561177157506001600160a01b03821615155b6117f35760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610559565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26118368261137a565b5050565b8047101561188a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610559565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146118d7576040519150601f19603f3d011682016040523d82523d6000602084013e6118dc565b606091505b50509050806119535760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610559565b505050565b6001600160a01b038116811461056b57600080fd5b60006020828403121561197f57600080fd5b813561198a81611958565b9392505050565b6000602082840312156119a357600080fd5b5035919050565b80516001600160e01b0316825260209081015163ffffffff16910152565b602080825282518282018190526000919060409081850190868401855b82811015611a08576119f88483516119aa565b92840192908501906001016119e5565b5091979650505050505050565b600060208284031215611a2757600080fd5b813560ff8116811461198a57600080fd5b60008060408385031215611a4b57600080fd5b8235611a5681611958565b946020939093013593505050565b60408101610b5782846119aa565b60008060408385031215611a8557600080fd5b8235611a9081611958565b91506020830135611aa081611958565b809150509250929050565b60008060008060808587031215611ac157600080fd5b8435611acc81611958565b93506020850135611adc81611958565b93969395505050506040820135916060013590565b602080825282518282015282810151604080840181905281516060850181905260009392830191849160808701905b8084101561089b57611b338286516119aa565b938501936001939093019290820190611b20565b600060208284031215611b5957600080fd5b815161198a81611958565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215611bc057600080fd5b8151801515811461198a57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600082821015611c4057611c40611c18565b500390565b60008219821115611c5857611c58611c18565b500190565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611c8757611c87611c18565b5060010190565b634e487b7160e01b600052604160045260246000fd5b60208082526019908201527f5061757361626c653a20696e646578206973207061757365640000000000000060408201526060019056fea264697066735822122000d7cf28f4403941c299573f87dac2e0ac298f40aace10d232ce8c2e9d0894a564736f6c634300080c0033",
}

// DelayedWithdrawalRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayedWithdrawalRouterMetaData.ABI instead.
var DelayedWithdrawalRouterABI = DelayedWithdrawalRouterMetaData.ABI

// DelayedWithdrawalRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelayedWithdrawalRouterMetaData.Bin instead.
var DelayedWithdrawalRouterBin = DelayedWithdrawalRouterMetaData.Bin

// DeployDelayedWithdrawalRouter deploys a new Ethereum contract, binding an instance of DelayedWithdrawalRouter to it.
func DeployDelayedWithdrawalRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _eigenPodManager common.Address) (common.Address, *types.Transaction, *DelayedWithdrawalRouter, error) {
	parsed, err := DelayedWithdrawalRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelayedWithdrawalRouterBin), backend, _eigenPodManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelayedWithdrawalRouter{DelayedWithdrawalRouterCaller: DelayedWithdrawalRouterCaller{contract: contract}, DelayedWithdrawalRouterTransactor: DelayedWithdrawalRouterTransactor{contract: contract}, DelayedWithdrawalRouterFilterer: DelayedWithdrawalRouterFilterer{contract: contract}}, nil
}

// DelayedWithdrawalRouter is an auto generated Go binding around an Ethereum contract.
type DelayedWithdrawalRouter struct {
	DelayedWithdrawalRouterCaller     // Read-only binding to the contract
	DelayedWithdrawalRouterTransactor // Write-only binding to the contract
	DelayedWithdrawalRouterFilterer   // Log filterer for contract events
}

// DelayedWithdrawalRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWithdrawalRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWithdrawalRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayedWithdrawalRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWithdrawalRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedWithdrawalRouterSession struct {
	Contract     *DelayedWithdrawalRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DelayedWithdrawalRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedWithdrawalRouterCallerSession struct {
	Contract *DelayedWithdrawalRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// DelayedWithdrawalRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedWithdrawalRouterTransactorSession struct {
	Contract     *DelayedWithdrawalRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// DelayedWithdrawalRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedWithdrawalRouterRaw struct {
	Contract *DelayedWithdrawalRouter // Generic contract binding to access the raw methods on
}

// DelayedWithdrawalRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterCallerRaw struct {
	Contract *DelayedWithdrawalRouterCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedWithdrawalRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedWithdrawalRouterTransactorRaw struct {
	Contract *DelayedWithdrawalRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedWithdrawalRouter creates a new instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouter(address common.Address, backend bind.ContractBackend) (*DelayedWithdrawalRouter, error) {
	contract, err := bindDelayedWithdrawalRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouter{DelayedWithdrawalRouterCaller: DelayedWithdrawalRouterCaller{contract: contract}, DelayedWithdrawalRouterTransactor: DelayedWithdrawalRouterTransactor{contract: contract}, DelayedWithdrawalRouterFilterer: DelayedWithdrawalRouterFilterer{contract: contract}}, nil
}

// NewDelayedWithdrawalRouterCaller creates a new read-only instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouterCaller(address common.Address, caller bind.ContractCaller) (*DelayedWithdrawalRouterCaller, error) {
	contract, err := bindDelayedWithdrawalRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterCaller{contract: contract}, nil
}

// NewDelayedWithdrawalRouterTransactor creates a new write-only instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedWithdrawalRouterTransactor, error) {
	contract, err := bindDelayedWithdrawalRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterTransactor{contract: contract}, nil
}

// NewDelayedWithdrawalRouterFilterer creates a new log filterer instance of DelayedWithdrawalRouter, bound to a specific deployed contract.
func NewDelayedWithdrawalRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayedWithdrawalRouterFilterer, error) {
	contract, err := bindDelayedWithdrawalRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterFilterer{contract: contract}, nil
}

// bindDelayedWithdrawalRouter binds a generic wrapper to an already deployed contract.
func bindDelayedWithdrawalRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelayedWithdrawalRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedWithdrawalRouter.Contract.DelayedWithdrawalRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.DelayedWithdrawalRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.DelayedWithdrawalRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedWithdrawalRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.contract.Transact(opts, method, params...)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.MAXWITHDRAWALDELAYBLOCKS(&_DelayedWithdrawalRouter.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.MAXWITHDRAWALDELAYBLOCKS(&_DelayedWithdrawalRouter.CallOpts)
}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) CanClaimDelayedWithdrawal(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "canClaimDelayedWithdrawal", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) CanClaimDelayedWithdrawal(user common.Address, index *big.Int) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.CanClaimDelayedWithdrawal(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// CanClaimDelayedWithdrawal is a free data retrieval call binding the contract method 0x75608896.
//
// Solidity: function canClaimDelayedWithdrawal(address user, uint256 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) CanClaimDelayedWithdrawal(user common.Address, index *big.Int) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.CanClaimDelayedWithdrawal(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) EigenPodManager() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.EigenPodManager(&_DelayedWithdrawalRouter.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) EigenPodManager() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.EigenPodManager(&_DelayedWithdrawalRouter.CallOpts)
}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) GetClaimableUserDelayedWithdrawals(opts *bind.CallOpts, user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "getClaimableUserDelayedWithdrawals", user)

	if err != nil {
		return *new([]IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelayedWithdrawalRouterDelayedWithdrawal)).(*[]IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) GetClaimableUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetClaimableUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// GetClaimableUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x1f39d87f.
//
// Solidity: function getClaimableUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) GetClaimableUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetClaimableUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) GetUserDelayedWithdrawals(opts *bind.CallOpts, user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "getUserDelayedWithdrawals", user)

	if err != nil {
		return *new([]IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelayedWithdrawalRouterDelayedWithdrawal)).(*[]IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) GetUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// GetUserDelayedWithdrawals is a free data retrieval call binding the contract method 0x3e1de008.
//
// Solidity: function getUserDelayedWithdrawals(address user) view returns((uint224,uint32)[])
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) GetUserDelayedWithdrawals(user common.Address) ([]IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.GetUserDelayedWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Owner() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.Owner(&_DelayedWithdrawalRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) Owner() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.Owner(&_DelayedWithdrawalRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Paused(index uint8) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.Paused(&_DelayedWithdrawalRouter.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) Paused(index uint8) (bool, error) {
	return _DelayedWithdrawalRouter.Contract.Paused(&_DelayedWithdrawalRouter.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Paused0() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.Paused0(&_DelayedWithdrawalRouter.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) Paused0() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.Paused0(&_DelayedWithdrawalRouter.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) PauserRegistry() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.PauserRegistry(&_DelayedWithdrawalRouter.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) PauserRegistry() (common.Address, error) {
	return _DelayedWithdrawalRouter.Contract.PauserRegistry(&_DelayedWithdrawalRouter.CallOpts)
}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) UserDelayedWithdrawalByIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "userDelayedWithdrawalByIndex", user, index)

	if err != nil {
		return *new(IDelayedWithdrawalRouterDelayedWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelayedWithdrawalRouterDelayedWithdrawal)).(*IDelayedWithdrawalRouterDelayedWithdrawal)

	return out0, err

}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) UserDelayedWithdrawalByIndex(user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.UserDelayedWithdrawalByIndex(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// UserDelayedWithdrawalByIndex is a free data retrieval call binding the contract method 0x85594e58.
//
// Solidity: function userDelayedWithdrawalByIndex(address user, uint256 index) view returns((uint224,uint32))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) UserDelayedWithdrawalByIndex(user common.Address, index *big.Int) (IDelayedWithdrawalRouterDelayedWithdrawal, error) {
	return _DelayedWithdrawalRouter.Contract.UserDelayedWithdrawalByIndex(&_DelayedWithdrawalRouter.CallOpts, user, index)
}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) UserWithdrawals(opts *bind.CallOpts, user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "userWithdrawals", user)

	if err != nil {
		return *new(IDelayedWithdrawalRouterUserDelayedWithdrawals), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelayedWithdrawalRouterUserDelayedWithdrawals)).(*IDelayedWithdrawalRouterUserDelayedWithdrawals)

	return out0, err

}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) UserWithdrawals(user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawals is a free data retrieval call binding the contract method 0xecb7cb1b.
//
// Solidity: function userWithdrawals(address user) view returns((uint256,(uint224,uint32)[]))
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) UserWithdrawals(user common.Address) (IDelayedWithdrawalRouterUserDelayedWithdrawals, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawals(&_DelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) UserWithdrawalsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "userWithdrawalsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) UserWithdrawalsLength(user common.Address) (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawalsLength(&_DelayedWithdrawalRouter.CallOpts, user)
}

// UserWithdrawalsLength is a free data retrieval call binding the contract method 0xe4f4f887.
//
// Solidity: function userWithdrawalsLength(address user) view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) UserWithdrawalsLength(user common.Address) (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.UserWithdrawalsLength(&_DelayedWithdrawalRouter.CallOpts, user)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCaller) WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWithdrawalRouter.contract.Call(opts, &out, "withdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.WithdrawalDelayBlocks(&_DelayedWithdrawalRouter.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterCallerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _DelayedWithdrawalRouter.Contract.WithdrawalDelayBlocks(&_DelayedWithdrawalRouter.CallOpts)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) ClaimDelayedWithdrawals(opts *bind.TransactOpts, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "claimDelayedWithdrawals", maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) ClaimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals(&_DelayedWithdrawalRouter.TransactOpts, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0xd44e1b76.
//
// Solidity: function claimDelayedWithdrawals(uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) ClaimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals(&_DelayedWithdrawalRouter.TransactOpts, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) ClaimDelayedWithdrawals0(opts *bind.TransactOpts, recipient common.Address, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "claimDelayedWithdrawals0", recipient, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) ClaimDelayedWithdrawals0(recipient common.Address, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals0(&_DelayedWithdrawalRouter.TransactOpts, recipient, maxNumberOfDelayedWithdrawalsToClaim)
}

// ClaimDelayedWithdrawals0 is a paid mutator transaction binding the contract method 0xe5db06c0.
//
// Solidity: function claimDelayedWithdrawals(address recipient, uint256 maxNumberOfDelayedWithdrawalsToClaim) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) ClaimDelayedWithdrawals0(recipient common.Address, maxNumberOfDelayedWithdrawalsToClaim *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.ClaimDelayedWithdrawals0(&_DelayedWithdrawalRouter.TransactOpts, recipient, maxNumberOfDelayedWithdrawalsToClaim)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) CreateDelayedWithdrawal(opts *bind.TransactOpts, podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "createDelayedWithdrawal", podOwner, recipient)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) CreateDelayedWithdrawal(podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.CreateDelayedWithdrawal(&_DelayedWithdrawalRouter.TransactOpts, podOwner, recipient)
}

// CreateDelayedWithdrawal is a paid mutator transaction binding the contract method 0xc0db354c.
//
// Solidity: function createDelayedWithdrawal(address podOwner, address recipient) payable returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) CreateDelayedWithdrawal(podOwner common.Address, recipient common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.CreateDelayedWithdrawal(&_DelayedWithdrawalRouter.TransactOpts, podOwner, recipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address initOwner, address _pauserRegistry, uint256 initPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) Initialize(opts *bind.TransactOpts, initOwner common.Address, _pauserRegistry common.Address, initPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "initialize", initOwner, _pauserRegistry, initPausedStatus, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address initOwner, address _pauserRegistry, uint256 initPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Initialize(initOwner common.Address, _pauserRegistry common.Address, initPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Initialize(&_DelayedWithdrawalRouter.TransactOpts, initOwner, _pauserRegistry, initPausedStatus, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address initOwner, address _pauserRegistry, uint256 initPausedStatus, uint256 _withdrawalDelayBlocks) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) Initialize(initOwner common.Address, _pauserRegistry common.Address, initPausedStatus *big.Int, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Initialize(&_DelayedWithdrawalRouter.TransactOpts, initOwner, _pauserRegistry, initPausedStatus, _withdrawalDelayBlocks)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Pause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Pause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) PauseAll() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.PauseAll(&_DelayedWithdrawalRouter.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) PauseAll() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.PauseAll(&_DelayedWithdrawalRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.RenounceOwnership(&_DelayedWithdrawalRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.RenounceOwnership(&_DelayedWithdrawalRouter.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetPauserRegistry(&_DelayedWithdrawalRouter.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetPauserRegistry(&_DelayedWithdrawalRouter.TransactOpts, newPauserRegistry)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) SetWithdrawalDelayBlocks(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "setWithdrawalDelayBlocks", newValue)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) SetWithdrawalDelayBlocks(newValue *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetWithdrawalDelayBlocks(&_DelayedWithdrawalRouter.TransactOpts, newValue)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 newValue) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) SetWithdrawalDelayBlocks(newValue *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.SetWithdrawalDelayBlocks(&_DelayedWithdrawalRouter.TransactOpts, newValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.TransferOwnership(&_DelayedWithdrawalRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.TransferOwnership(&_DelayedWithdrawalRouter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Unpause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelayedWithdrawalRouter.Contract.Unpause(&_DelayedWithdrawalRouter.TransactOpts, newPausedStatus)
}

// DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator is returned from FilterDelayedWithdrawalCreated and is used to iterate over the raw logs and unpacked data for DelayedWithdrawalCreated events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator struct {
	Event *DelayedWithdrawalRouterDelayedWithdrawalCreated // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
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
		it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
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
func (it *DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterDelayedWithdrawalCreated represents a DelayedWithdrawalCreated event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalCreated struct {
	PodOwner  common.Address
	Recipient common.Address
	Amount    *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedWithdrawalCreated is a free log retrieval operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterDelayedWithdrawalCreated(opts *bind.FilterOpts) (*DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "DelayedWithdrawalCreated")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterDelayedWithdrawalCreatedIterator{contract: _DelayedWithdrawalRouter.contract, event: "DelayedWithdrawalCreated", logs: logs, sub: sub}, nil
}

// WatchDelayedWithdrawalCreated is a free log subscription operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchDelayedWithdrawalCreated(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterDelayedWithdrawalCreated) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "DelayedWithdrawalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalCreated", log); err != nil {
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

// ParseDelayedWithdrawalCreated is a log parse operation binding the contract event 0xb8f1b14c7caf74150801dcc9bc18d575cbeaf5b421943497e409df92c92e0f59.
//
// Solidity: event DelayedWithdrawalCreated(address podOwner, address recipient, uint256 amount, uint256 index)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseDelayedWithdrawalCreated(log types.Log) (*DelayedWithdrawalRouterDelayedWithdrawalCreated, error) {
	event := new(DelayedWithdrawalRouterDelayedWithdrawalCreated)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator is returned from FilterDelayedWithdrawalsClaimed and is used to iterate over the raw logs and unpacked data for DelayedWithdrawalsClaimed events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator struct {
	Event *DelayedWithdrawalRouterDelayedWithdrawalsClaimed // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
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
		it.Event = new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
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
func (it *DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterDelayedWithdrawalsClaimed represents a DelayedWithdrawalsClaimed event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterDelayedWithdrawalsClaimed struct {
	Recipient                   common.Address
	AmountClaimed               *big.Int
	DelayedWithdrawalsCompleted *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDelayedWithdrawalsClaimed is a free log retrieval operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterDelayedWithdrawalsClaimed(opts *bind.FilterOpts) (*DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "DelayedWithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterDelayedWithdrawalsClaimedIterator{contract: _DelayedWithdrawalRouter.contract, event: "DelayedWithdrawalsClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedWithdrawalsClaimed is a free log subscription operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchDelayedWithdrawalsClaimed(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterDelayedWithdrawalsClaimed) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "DelayedWithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalsClaimed", log); err != nil {
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

// ParseDelayedWithdrawalsClaimed is a log parse operation binding the contract event 0x6b7151500bd0b5cc211bcc47b3029831b769004df4549e8e1c9a69da05bb0943.
//
// Solidity: event DelayedWithdrawalsClaimed(address recipient, uint256 amountClaimed, uint256 delayedWithdrawalsCompleted)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseDelayedWithdrawalsClaimed(log types.Log) (*DelayedWithdrawalRouterDelayedWithdrawalsClaimed, error) {
	event := new(DelayedWithdrawalRouterDelayedWithdrawalsClaimed)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "DelayedWithdrawalsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterInitializedIterator struct {
	Event *DelayedWithdrawalRouterInitialized // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterInitialized)
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
		it.Event = new(DelayedWithdrawalRouterInitialized)
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
func (it *DelayedWithdrawalRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterInitialized represents a Initialized event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*DelayedWithdrawalRouterInitializedIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterInitializedIterator{contract: _DelayedWithdrawalRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterInitialized)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseInitialized(log types.Log) (*DelayedWithdrawalRouterInitialized, error) {
	event := new(DelayedWithdrawalRouterInitialized)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterOwnershipTransferredIterator struct {
	Event *DelayedWithdrawalRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterOwnershipTransferred)
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
		it.Event = new(DelayedWithdrawalRouterOwnershipTransferred)
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
func (it *DelayedWithdrawalRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterOwnershipTransferred represents a OwnershipTransferred event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DelayedWithdrawalRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterOwnershipTransferredIterator{contract: _DelayedWithdrawalRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterOwnershipTransferred)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseOwnershipTransferred(log types.Log) (*DelayedWithdrawalRouterOwnershipTransferred, error) {
	event := new(DelayedWithdrawalRouterOwnershipTransferred)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPausedIterator struct {
	Event *DelayedWithdrawalRouterPaused // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterPaused)
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
		it.Event = new(DelayedWithdrawalRouterPaused)
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
func (it *DelayedWithdrawalRouterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterPaused represents a Paused event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*DelayedWithdrawalRouterPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterPausedIterator{contract: _DelayedWithdrawalRouter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterPaused)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParsePaused(log types.Log) (*DelayedWithdrawalRouterPaused, error) {
	event := new(DelayedWithdrawalRouterPaused)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPauserRegistrySetIterator struct {
	Event *DelayedWithdrawalRouterPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterPauserRegistrySet)
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
		it.Event = new(DelayedWithdrawalRouterPauserRegistrySet)
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
func (it *DelayedWithdrawalRouterPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterPauserRegistrySet represents a PauserRegistrySet event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*DelayedWithdrawalRouterPauserRegistrySetIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterPauserRegistrySetIterator{contract: _DelayedWithdrawalRouter.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterPauserRegistrySet)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParsePauserRegistrySet(log types.Log) (*DelayedWithdrawalRouterPauserRegistrySet, error) {
	event := new(DelayedWithdrawalRouterPauserRegistrySet)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterUnpausedIterator struct {
	Event *DelayedWithdrawalRouterUnpaused // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterUnpaused)
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
		it.Event = new(DelayedWithdrawalRouterUnpaused)
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
func (it *DelayedWithdrawalRouterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterUnpaused represents a Unpaused event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*DelayedWithdrawalRouterUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterUnpausedIterator{contract: _DelayedWithdrawalRouter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterUnpaused)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseUnpaused(log types.Log) (*DelayedWithdrawalRouterUnpaused, error) {
	event := new(DelayedWithdrawalRouterUnpaused)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator is returned from FilterWithdrawalDelayBlocksSet and is used to iterate over the raw logs and unpacked data for WithdrawalDelayBlocksSet events raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator struct {
	Event *DelayedWithdrawalRouterWithdrawalDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
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
		it.Event = new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
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
func (it *DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWithdrawalRouterWithdrawalDelayBlocksSet represents a WithdrawalDelayBlocksSet event raised by the DelayedWithdrawalRouter contract.
type DelayedWithdrawalRouterWithdrawalDelayBlocksSet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayBlocksSet is a free log retrieval operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) FilterWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.FilterLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &DelayedWithdrawalRouterWithdrawalDelayBlocksSetIterator{contract: _DelayedWithdrawalRouter.contract, event: "WithdrawalDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayBlocksSet is a free log subscription operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) WatchWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *DelayedWithdrawalRouterWithdrawalDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _DelayedWithdrawalRouter.contract.WatchLogs(opts, "WithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
				if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
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

// ParseWithdrawalDelayBlocksSet is a log parse operation binding the contract event 0x4ffb00400574147429ee377a5633386321e66d45d8b14676014b5fa393e61e9e.
//
// Solidity: event WithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelayedWithdrawalRouter *DelayedWithdrawalRouterFilterer) ParseWithdrawalDelayBlocksSet(log types.Log) (*DelayedWithdrawalRouterWithdrawalDelayBlocksSet, error) {
	event := new(DelayedWithdrawalRouterWithdrawalDelayBlocksSet)
	if err := _DelayedWithdrawalRouter.contract.UnpackLog(event, "WithdrawalDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
