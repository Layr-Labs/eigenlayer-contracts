// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AVSDirectory

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// AVSDirectoryMetaData contains all meta data concerning the AVSDirectory contract.
var AVSDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectoryTypes.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isSpent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSMigratedToOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectoryTypes.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMigratedToOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToEigenLayer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]}]",
	Bin: "0x61010060405234801561001157600080fd5b506040516117c13803806117c18339810160408190526100309161020a565b81816001600160a01b038116610059576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b039081166080521660a0524660c052610077610089565b60e052610082610133565b5050610244565b600060c051461461012c5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b5060e05190565b600054610100900460ff161561019f5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146101f0576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461020757600080fd5b50565b6000806040838503121561021d57600080fd5b8251610228816101f2565b6020840151909250610239816101f2565b809150509250929050565b60805160a05160c05160e0516115286102996000396000610e4201526000610d82015260008181610348015261064f015260008181610224015281816103e1015281816104b80152610ae701526115286000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063a364f4da116100b8578063dce974b91161007c578063dce974b91461031c578063df5cf72314610343578063ec76f4421461036a578063f2fde38b1461039e578063f698da25146103b1578063fabc1cbc146103b957600080fd5b8063a364f4da14610295578063a98fb355146102a8578063c825fe68146102bb578063cd6dc687146102e2578063d79aceab146102f557600080fd5b80635c975abb1161010a5780635c975abb14610205578063715018a614610217578063886f11951461021f5780638da5cb5b1461025e5780639926ee7d1461026f578063a1060c881461028257600080fd5b8063136439dd14610147578063374823b51461015c57806349075da31461019f578063595c6a67146101da5780635ac86ab7146101e2575b600080fd5b61015a6101553660046110ba565b6103cc565b005b61018a61016a3660046110e8565b609960209081526000928352604080842090915290825290205460ff1681565b60405190151581526020015b60405180910390f35b6101cd6101ad366004611114565b609860209081526000928352604080842090915290825290205460ff1681565b6040516101969190611163565b61015a6104a3565b61018a6101f036600461118b565b606654600160ff9092169190911b9081161490565b6066545b604051908152602001610196565b61015a610555565b6102467f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610196565b6033546001600160a01b0316610246565b61015a61027d366004611225565b610567565b61020961029036600461131a565b61078c565b61015a6102a3366004611360565b61080c565b61015a6102b636600461137d565b6108f4565b6102097f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92981565b61015a6102f03660046110e8565b61093b565b6102097fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd81565b6102097f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49581565b6102467f000000000000000000000000000000000000000000000000000000000000000081565b61015a6103783660046110ba565b33600090815260996020908152604080832093835292905220805460ff19166001179055565b61015a6103ac366004611360565b610a5d565b610209610ad6565b61015a6103c73660046110ba565b610ae5565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610430573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045491906113f1565b61047157604051631d77d47760e21b815260040160405180910390fd5b60665481811681146104965760405163c61dca5d60e01b815260040160405180910390fd5b61049f82610bf6565b5050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610507573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061052b91906113f1565b61054857604051631d77d47760e21b815260040160405180910390fd5b610553600019610bf6565b565b61055d610c33565b6105536000610c8d565b6066546000906001908116036105905760405163840a48d560e01b815260040160405180910390fd5b60013360009081526098602090815260408083206001600160a01b038816845290915290205460ff1660018111156105ca576105ca61114d565b036105e857604051631aa528bb60e11b815260040160405180910390fd5b6001600160a01b038316600090815260996020908152604080832085830151845290915290205460ff161561063057604051630d4c4c9160e21b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015610696573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ba91906113f1565b6106d757604051639f88c8af60e01b815260040160405180910390fd5b6106fb836106ef85338660200151876040015161078c565b84516040860151610cdf565b6001600160a01b038316600081815260996020908152604080832086830151845282528083208054600160ff19918216811790925533808652609885528386208787529094529382902080549094168117909355519092917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b419161077f9190611163565b60405180910390a3505050565b604080517fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd60208201526001600160a01b038087169282019290925290841660608201526080810183905260a081018290526000906108039060c00160405160208183030381529060405280519060200120610d37565b95945050505050565b6066546000906001908116036108355760405163840a48d560e01b815260040160405180910390fd5b60013360009081526098602090815260408083206001600160a01b038716845290915290205460ff16600181111561086f5761086f61114d565b1461088d576040516352df45c960e01b815260040160405180910390fd5b3360008181526098602090815260408083206001600160a01b0387168085529252808320805460ff191690555190917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41916108e89190611163565b60405180910390a35050565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713838360405161092f929190611413565b60405180910390a25050565b600054610100900460ff161580801561095b5750600054600160ff909116105b806109755750303b158015610975575060005460ff166001145b6109dd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610a00576000805461ff0019166101001790555b610a0982610bf6565b610a1283610c8d565b8015610a58576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610a65610c33565b6001600160a01b038116610aca5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109d4565b610ad381610c8d565b50565b6000610ae0610d7e565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b679190611442565b6001600160a01b0316336001600160a01b031614610b985760405163794821ff60e01b815260040160405180910390fd5b60665480198219811614610bbf5760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161092f565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6033546001600160a01b031633146105535760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109d4565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b42811015610d0057604051630819bdcd60e01b815260040160405180910390fd5b610d146001600160a01b0385168484610e64565b610d3157604051638baa579f60e01b815260040160405180910390fd5b50505050565b6000610d41610d7e565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60007f00000000000000000000000000000000000000000000000000000000000000004614610e3f5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b507f000000000000000000000000000000000000000000000000000000000000000090565b6000806000610e738585610ec5565b90925090506000816004811115610e8c57610e8c61114d565b148015610eaa5750856001600160a01b0316826001600160a01b0316145b80610ebb5750610ebb868686610f0a565b9695505050505050565b6000808251604103610efb5760208301516040840151606085015160001a610eef87828585610ff6565b94509450505050610f03565b506000905060025b9250929050565b6000806000856001600160a01b0316631626ba7e60e01b8686604051602401610f34929190611483565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051610f7291906114bd565b600060405180830381855afa9150503d8060008114610fad576040519150601f19603f3d011682016040523d82523d6000602084013e610fb2565b606091505b5091509150818015610fc657506020815110155b8015610ebb57508051630b135d3f60e11b90610feb90830160209081019084016114d9565b149695505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561102d57506000905060036110b1565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611081573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166110aa576000600192509250506110b1565b9150600090505b94509492505050565b6000602082840312156110cc57600080fd5b5035919050565b6001600160a01b0381168114610ad357600080fd5b600080604083850312156110fb57600080fd5b8235611106816110d3565b946020939093013593505050565b6000806040838503121561112757600080fd5b8235611132816110d3565b91506020830135611142816110d3565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b602081016002831061118557634e487b7160e01b600052602160045260246000fd5b91905290565b60006020828403121561119d57600080fd5b813560ff811681146111ae57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156111ee576111ee6111b5565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561121d5761121d6111b5565b604052919050565b6000806040838503121561123857600080fd5b8235611243816110d3565b9150602083013567ffffffffffffffff81111561125f57600080fd5b83016060818603121561127157600080fd5b6112796111cb565b813567ffffffffffffffff81111561129057600080fd5b8201601f810187136112a157600080fd5b803567ffffffffffffffff8111156112bb576112bb6111b5565b6112ce601f8201601f19166020016111f4565b8181528860208385010111156112e357600080fd5b8160208401602083013760006020928201830152835283810135908301525060409182013591810191909152919491935090915050565b6000806000806080858703121561133057600080fd5b843561133b816110d3565b9350602085013561134b816110d3565b93969395505050506040820135916060013590565b60006020828403121561137257600080fd5b81356111ae816110d3565b6000806020838503121561139057600080fd5b823567ffffffffffffffff8111156113a757600080fd5b8301601f810185136113b857600080fd5b803567ffffffffffffffff8111156113cf57600080fd5b8560208284010111156113e157600080fd5b6020919091019590945092505050565b60006020828403121561140357600080fd5b815180151581146111ae57600080fd5b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561145457600080fd5b81516111ae816110d3565b60005b8381101561147a578181015183820152602001611462565b50506000910152565b82815260406020820152600082518060408401526114a881606085016020870161145f565b601f01601f1916919091016060019392505050565b600082516114cf81846020870161145f565b9190910192915050565b6000602082840312156114eb57600080fd5b505191905056fea2646970667358221220fb2464241996dd1fd1c6a880ea9de67364f839547eaf07b816ff98a29573387a64736f6c634300081b0033",
}

// AVSDirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSDirectoryMetaData.ABI instead.
var AVSDirectoryABI = AVSDirectoryMetaData.ABI

// AVSDirectoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AVSDirectoryMetaData.Bin instead.
var AVSDirectoryBin = AVSDirectoryMetaData.Bin

// DeployAVSDirectory deploys a new Ethereum contract, binding an instance of AVSDirectory to it.
func DeployAVSDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *AVSDirectory, error) {
	parsed, err := AVSDirectoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AVSDirectoryBin), backend, _delegation, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AVSDirectory{AVSDirectoryCaller: AVSDirectoryCaller{contract: contract}, AVSDirectoryTransactor: AVSDirectoryTransactor{contract: contract}, AVSDirectoryFilterer: AVSDirectoryFilterer{contract: contract}}, nil
}

// AVSDirectory is an auto generated Go binding around an Ethereum contract.
type AVSDirectory struct {
	AVSDirectoryCaller     // Read-only binding to the contract
	AVSDirectoryTransactor // Write-only binding to the contract
	AVSDirectoryFilterer   // Log filterer for contract events
}

// AVSDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSDirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSDirectorySession struct {
	Contract     *AVSDirectory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AVSDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSDirectoryCallerSession struct {
	Contract *AVSDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AVSDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSDirectoryTransactorSession struct {
	Contract     *AVSDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AVSDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSDirectoryRaw struct {
	Contract *AVSDirectory // Generic contract binding to access the raw methods on
}

// AVSDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSDirectoryCallerRaw struct {
	Contract *AVSDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// AVSDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSDirectoryTransactorRaw struct {
	Contract *AVSDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSDirectory creates a new instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectory(address common.Address, backend bind.ContractBackend) (*AVSDirectory, error) {
	contract, err := bindAVSDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSDirectory{AVSDirectoryCaller: AVSDirectoryCaller{contract: contract}, AVSDirectoryTransactor: AVSDirectoryTransactor{contract: contract}, AVSDirectoryFilterer: AVSDirectoryFilterer{contract: contract}}, nil
}

// NewAVSDirectoryCaller creates a new read-only instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectoryCaller(address common.Address, caller bind.ContractCaller) (*AVSDirectoryCaller, error) {
	contract, err := bindAVSDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryCaller{contract: contract}, nil
}

// NewAVSDirectoryTransactor creates a new write-only instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSDirectoryTransactor, error) {
	contract, err := bindAVSDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryTransactor{contract: contract}, nil
}

// NewAVSDirectoryFilterer creates a new log filterer instance of AVSDirectory, bound to a specific deployed contract.
func NewAVSDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSDirectoryFilterer, error) {
	contract, err := bindAVSDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryFilterer{contract: contract}, nil
}

// bindAVSDirectory binds a generic wrapper to an already deployed contract.
func bindAVSDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSDirectoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectory *AVSDirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectory.Contract.AVSDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectory *AVSDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.Contract.AVSDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectory *AVSDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectory.Contract.AVSDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSDirectory *AVSDirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSDirectory *AVSDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSDirectory *AVSDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSDirectory.Contract.contract.Transact(opts, method, params...)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) OPERATORSETFORCEDEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _AVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_AVSDirectory.CallOpts)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_AVSDirectory *AVSDirectoryCaller) AvsOperatorStatus(opts *bind.CallOpts, avs common.Address, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "avsOperatorStatus", avs, operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_AVSDirectory *AVSDirectorySession) AvsOperatorStatus(avs common.Address, operator common.Address) (uint8, error) {
	return _AVSDirectory.Contract.AvsOperatorStatus(&_AVSDirectory.CallOpts, avs, operator)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_AVSDirectory *AVSDirectoryCallerSession) AvsOperatorStatus(avs common.Address, operator common.Address) (uint8, error) {
	return _AVSDirectory.Contract.AvsOperatorStatus(&_AVSDirectory.CallOpts, avs, operator)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _AVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectory *AVSDirectoryCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectory *AVSDirectorySession) Delegation() (common.Address, error) {
	return _AVSDirectory.Contract.Delegation(&_AVSDirectory.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_AVSDirectory *AVSDirectoryCallerSession) Delegation() (common.Address, error) {
	return _AVSDirectory.Contract.Delegation(&_AVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectory *AVSDirectorySession) DomainSeparator() ([32]byte, error) {
	return _AVSDirectory.Contract.DomainSeparator(&_AVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AVSDirectory *AVSDirectoryCallerSession) DomainSeparator() ([32]byte, error) {
	return _AVSDirectory.Contract.DomainSeparator(&_AVSDirectory.CallOpts)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_AVSDirectory *AVSDirectoryCaller) OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "operatorSaltIsSpent", operator, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_AVSDirectory *AVSDirectorySession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _AVSDirectory.Contract.OperatorSaltIsSpent(&_AVSDirectory.CallOpts, operator, salt)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_AVSDirectory *AVSDirectoryCallerSession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _AVSDirectory.Contract.OperatorSaltIsSpent(&_AVSDirectory.CallOpts, operator, salt)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSDirectory *AVSDirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSDirectory *AVSDirectorySession) Owner() (common.Address, error) {
	return _AVSDirectory.Contract.Owner(&_AVSDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSDirectory *AVSDirectoryCallerSession) Owner() (common.Address, error) {
	return _AVSDirectory.Contract.Owner(&_AVSDirectory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AVSDirectory *AVSDirectoryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AVSDirectory *AVSDirectorySession) Paused(index uint8) (bool, error) {
	return _AVSDirectory.Contract.Paused(&_AVSDirectory.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_AVSDirectory *AVSDirectoryCallerSession) Paused(index uint8) (bool, error) {
	return _AVSDirectory.Contract.Paused(&_AVSDirectory.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AVSDirectory *AVSDirectoryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AVSDirectory *AVSDirectorySession) Paused0() (*big.Int, error) {
	return _AVSDirectory.Contract.Paused0(&_AVSDirectory.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_AVSDirectory *AVSDirectoryCallerSession) Paused0() (*big.Int, error) {
	return _AVSDirectory.Contract.Paused0(&_AVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AVSDirectory *AVSDirectoryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSDirectory.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AVSDirectory *AVSDirectorySession) PauserRegistry() (common.Address, error) {
	return _AVSDirectory.Contract.PauserRegistry(&_AVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_AVSDirectory *AVSDirectoryCallerSession) PauserRegistry() (common.Address, error) {
	return _AVSDirectory.Contract.PauserRegistry(&_AVSDirectory.CallOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectory *AVSDirectoryTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectory *AVSDirectorySession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CancelSalt(&_AVSDirectory.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _AVSDirectory.Contract.CancelSalt(&_AVSDirectory.TransactOpts, salt)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_AVSDirectory *AVSDirectoryTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_AVSDirectory *AVSDirectorySession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.DeregisterOperatorFromAVS(&_AVSDirectory.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.DeregisterOperatorFromAVS(&_AVSDirectory.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_AVSDirectory *AVSDirectorySession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Initialize(&_AVSDirectory.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Initialize(&_AVSDirectory.TransactOpts, initialOwner, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectorySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Pause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Pause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AVSDirectory *AVSDirectoryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AVSDirectory *AVSDirectorySession) PauseAll() (*types.Transaction, error) {
	return _AVSDirectory.Contract.PauseAll(&_AVSDirectory.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _AVSDirectory.Contract.PauseAll(&_AVSDirectory.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectorySession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.RegisterOperatorToAVS(&_AVSDirectory.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSDirectory.Contract.RegisterOperatorToAVS(&_AVSDirectory.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSDirectory *AVSDirectoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSDirectory *AVSDirectorySession) RenounceOwnership() (*types.Transaction, error) {
	return _AVSDirectory.Contract.RenounceOwnership(&_AVSDirectory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AVSDirectory.Contract.RenounceOwnership(&_AVSDirectory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSDirectory *AVSDirectoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSDirectory *AVSDirectorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.TransferOwnership(&_AVSDirectory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AVSDirectory.Contract.TransferOwnership(&_AVSDirectory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectorySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Unpause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _AVSDirectory.Contract.Unpause(&_AVSDirectory.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectory *AVSDirectoryTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _AVSDirectory.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectory *AVSDirectorySession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectory.Contract.UpdateAVSMetadataURI(&_AVSDirectory.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSDirectory *AVSDirectoryTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSDirectory.Contract.UpdateAVSMetadataURI(&_AVSDirectory.TransactOpts, metadataURI)
}

// AVSDirectoryAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the AVSDirectory contract.
type AVSDirectoryAVSMetadataURIUpdatedIterator struct {
	Event *AVSDirectoryAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryAVSMetadataURIUpdated)
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
		it.Event = new(AVSDirectoryAVSMetadataURIUpdated)
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
func (it *AVSDirectoryAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the AVSDirectory contract.
type AVSDirectoryAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectory *AVSDirectoryFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*AVSDirectoryAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryAVSMetadataURIUpdatedIterator{contract: _AVSDirectory.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectory *AVSDirectoryFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryAVSMetadataURIUpdated)
				if err := _AVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_AVSDirectory *AVSDirectoryFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*AVSDirectoryAVSMetadataURIUpdated, error) {
	event := new(AVSDirectoryAVSMetadataURIUpdated)
	if err := _AVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryAVSMigratedToOperatorSetsIterator is returned from FilterAVSMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for AVSMigratedToOperatorSets events raised by the AVSDirectory contract.
type AVSDirectoryAVSMigratedToOperatorSetsIterator struct {
	Event *AVSDirectoryAVSMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryAVSMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryAVSMigratedToOperatorSets)
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
		it.Event = new(AVSDirectoryAVSMigratedToOperatorSets)
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
func (it *AVSDirectoryAVSMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryAVSMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryAVSMigratedToOperatorSets represents a AVSMigratedToOperatorSets event raised by the AVSDirectory contract.
type AVSDirectoryAVSMigratedToOperatorSets struct {
	Avs common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAVSMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectory *AVSDirectoryFilterer) FilterAVSMigratedToOperatorSets(opts *bind.FilterOpts, avs []common.Address) (*AVSDirectoryAVSMigratedToOperatorSetsIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryAVSMigratedToOperatorSetsIterator{contract: _AVSDirectory.contract, event: "AVSMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchAVSMigratedToOperatorSets is a free log subscription operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectory *AVSDirectoryFilterer) WatchAVSMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *AVSDirectoryAVSMigratedToOperatorSets, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "AVSMigratedToOperatorSets", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryAVSMigratedToOperatorSets)
				if err := _AVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
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

// ParseAVSMigratedToOperatorSets is a log parse operation binding the contract event 0x702b0c1f6cb1cf511aaa81f72bc05a215bb3497632d72c690c822b044ab494bf.
//
// Solidity: event AVSMigratedToOperatorSets(address indexed avs)
func (_AVSDirectory *AVSDirectoryFilterer) ParseAVSMigratedToOperatorSets(log types.Log) (*AVSDirectoryAVSMigratedToOperatorSets, error) {
	event := new(AVSDirectoryAVSMigratedToOperatorSets)
	if err := _AVSDirectory.contract.UnpackLog(event, "AVSMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AVSDirectory contract.
type AVSDirectoryInitializedIterator struct {
	Event *AVSDirectoryInitialized // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryInitialized)
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
		it.Event = new(AVSDirectoryInitialized)
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
func (it *AVSDirectoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryInitialized represents a Initialized event raised by the AVSDirectory contract.
type AVSDirectoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSDirectory *AVSDirectoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AVSDirectoryInitializedIterator, error) {

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryInitializedIterator{contract: _AVSDirectory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSDirectory *AVSDirectoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AVSDirectoryInitialized) (event.Subscription, error) {

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryInitialized)
				if err := _AVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseInitialized(log types.Log) (*AVSDirectoryInitialized, error) {
	event := new(AVSDirectoryInitialized)
	if err := _AVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the AVSDirectory contract.
type AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *AVSDirectoryOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
func (it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the AVSDirectory contract.
type AVSDirectoryOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator{contract: _AVSDirectory.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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

// ParseOperatorAVSRegistrationStatusUpdated is a log parse operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*AVSDirectoryOperatorAVSRegistrationStatusUpdated, error) {
	event := new(AVSDirectoryOperatorAVSRegistrationStatusUpdated)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOperatorMigratedToOperatorSetsIterator is returned from FilterOperatorMigratedToOperatorSets and is used to iterate over the raw logs and unpacked data for OperatorMigratedToOperatorSets events raised by the AVSDirectory contract.
type AVSDirectoryOperatorMigratedToOperatorSetsIterator struct {
	Event *AVSDirectoryOperatorMigratedToOperatorSets // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOperatorMigratedToOperatorSetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOperatorMigratedToOperatorSets)
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
		it.Event = new(AVSDirectoryOperatorMigratedToOperatorSets)
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
func (it *AVSDirectoryOperatorMigratedToOperatorSetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOperatorMigratedToOperatorSetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOperatorMigratedToOperatorSets represents a OperatorMigratedToOperatorSets event raised by the AVSDirectory contract.
type AVSDirectoryOperatorMigratedToOperatorSets struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorMigratedToOperatorSets is a free log retrieval operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOperatorMigratedToOperatorSets(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*AVSDirectoryOperatorMigratedToOperatorSetsIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOperatorMigratedToOperatorSetsIterator{contract: _AVSDirectory.contract, event: "OperatorMigratedToOperatorSets", logs: logs, sub: sub}, nil
}

// WatchOperatorMigratedToOperatorSets is a free log subscription operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOperatorMigratedToOperatorSets(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOperatorMigratedToOperatorSets, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OperatorMigratedToOperatorSets", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOperatorMigratedToOperatorSets)
				if err := _AVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
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

// ParseOperatorMigratedToOperatorSets is a log parse operation binding the contract event 0x54f33cfdd1ca703d795986b986fd47d742eab1904ecd2a5fdb8d6595e5904a01.
//
// Solidity: event OperatorMigratedToOperatorSets(address indexed operator, address indexed avs, uint32[] operatorSetIds)
func (_AVSDirectory *AVSDirectoryFilterer) ParseOperatorMigratedToOperatorSets(log types.Log) (*AVSDirectoryOperatorMigratedToOperatorSets, error) {
	event := new(AVSDirectoryOperatorMigratedToOperatorSets)
	if err := _AVSDirectory.contract.UnpackLog(event, "OperatorMigratedToOperatorSets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AVSDirectory contract.
type AVSDirectoryOwnershipTransferredIterator struct {
	Event *AVSDirectoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryOwnershipTransferred)
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
		it.Event = new(AVSDirectoryOwnershipTransferred)
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
func (it *AVSDirectoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryOwnershipTransferred represents a OwnershipTransferred event raised by the AVSDirectory contract.
type AVSDirectoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AVSDirectory *AVSDirectoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AVSDirectoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryOwnershipTransferredIterator{contract: _AVSDirectory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AVSDirectory *AVSDirectoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AVSDirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryOwnershipTransferred)
				if err := _AVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseOwnershipTransferred(log types.Log) (*AVSDirectoryOwnershipTransferred, error) {
	event := new(AVSDirectoryOwnershipTransferred)
	if err := _AVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AVSDirectory contract.
type AVSDirectoryPausedIterator struct {
	Event *AVSDirectoryPaused // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryPaused)
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
		it.Event = new(AVSDirectoryPaused)
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
func (it *AVSDirectoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryPaused represents a Paused event raised by the AVSDirectory contract.
type AVSDirectoryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*AVSDirectoryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryPausedIterator{contract: _AVSDirectory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AVSDirectoryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryPaused)
				if err := _AVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParsePaused(log types.Log) (*AVSDirectoryPaused, error) {
	event := new(AVSDirectoryPaused)
	if err := _AVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSDirectoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AVSDirectory contract.
type AVSDirectoryUnpausedIterator struct {
	Event *AVSDirectoryUnpaused // Event containing the contract specifics and raw log

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
func (it *AVSDirectoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSDirectoryUnpaused)
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
		it.Event = new(AVSDirectoryUnpaused)
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
func (it *AVSDirectoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSDirectoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSDirectoryUnpaused represents a Unpaused event raised by the AVSDirectory contract.
type AVSDirectoryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*AVSDirectoryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &AVSDirectoryUnpausedIterator{contract: _AVSDirectory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_AVSDirectory *AVSDirectoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AVSDirectoryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AVSDirectory.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSDirectoryUnpaused)
				if err := _AVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AVSDirectory *AVSDirectoryFilterer) ParseUnpaused(log types.Log) (*AVSDirectoryUnpaused, error) {
	event := new(AVSDirectoryUnpaused)
	if err := _AVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
