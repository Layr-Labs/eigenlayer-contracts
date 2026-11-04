// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ComputeRegistry

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

// IComputeRegistryTypesTOSSignature is an auto generated low-level Go binding around an user-defined struct.
type IComputeRegistryTypesTOSSignature struct {
	Signer    common.Address
	TosHash   [32]byte
	Signature []byte
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ComputeRegistryMetaData contains all meta data concerning the ComputeRegistry contract.
var ComputeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_releaseManager\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_crossChainRegistry\",\"type\":\"address\",\"internalType\":\"contractICrossChainRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_tosHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CROSS_CHAIN_REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICrossChainRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KEY_REGISTRAR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_EXPIRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELEASE_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOS_AGREEMENT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOS_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateTOSAgreementDigest\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromCompute\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetTosSignature\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIComputeRegistryTypes.TOSSignature\",\"components\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tosHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSetKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForCompute\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetDeregistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetRegistered\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tosHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurveTypeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTOSSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveGenerationReservation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x610160604052348015610010575f5ffd5b50604051611a36380380611a3683398101604081905261002f916101aa565b6001600160a01b0380881660805280871660a05280861660c05280851660e052610100839052831661012052808061006681610081565b610140525061007590506100c7565b50505050505050610323565b5f5f829050601f815111156100b4578260405163305a27a960e01b81526004016100ab91906102c8565b60405180910390fd5b80516100bf826102fd565b179392505050565b5f54610100900460ff161561012e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100ab565b5f5460ff9081161461017d575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610193575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f5f60e0888a0312156101c0575f5ffd5b87516101cb8161017f565b60208901519097506101dc8161017f565b60408901519096506101ed8161017f565b60608901519095506101fe8161017f565b608089015190945061020f8161017f565b60a089015160c08a015191945092506001600160401b03811115610231575f5ffd5b88015f601f82018b13610242575f5ffd5b81516001600160401b0381111561025b5761025b610196565b604051601f8201601f19908116603f011681016001600160401b038111828210171561028957610289610196565b6040528181528382016020018d10156102a0575f5ffd5b8160208501602083015e5f602083830101528092508094505050505092959891949750929550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561031d575f198160200360031b1b821691505b50919050565b60805160a05160c05160e05161010051610120516101405161167c6103ba5f395f81816108230152610d2101525f81816101960152610bd201525f818161020d0152818161071d015281816107b1015261088201525f818161023401526105c001525f81816102bc015261050201525f81816101570152818161041c015261093b01525f818161025b0152610668015261167c5ff3fe608060405234801561000f575f5ffd5b50600436106100f0575f3560e01c806389fec15f11610093578063b967169011610063578063b96716901461027d578063c4a1ca0514610285578063e6414b48146102b7578063f698da25146102de575f5ffd5b806389fec15f146101f55780638df643c7146102085780639b2508441461022f578063b39d254f14610256575f5ffd5b80634657e26a116100ce5780634657e26a14610191578063536b2353146101b857806354fd4d50146101cd5780637a1bb660146101e2575f5ffd5b8063130d6165146100f45780631de02dbb1461011d57806331232bc914610152575b5f5ffd5b6101076101023660046110d4565b6102e6565b604051610114919061111c565b60405180910390f35b6101447f21081bb395a13368a22c061a9d5d4bf6ddd7a9ae84a22a2a48c19d3f2b58a92381565b604051908152602001610114565b6101797f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610114565b6101797f000000000000000000000000000000000000000000000000000000000000000081565b6101cb6101c6366004611181565b6103db565b005b6101d561081c565b604051610114919061120d565b6101446101f036600461121f565b61084c565b6101cb6102033660046110d4565b6108fa565b6101447f000000000000000000000000000000000000000000000000000000000000000081565b6101797f000000000000000000000000000000000000000000000000000000000000000081565b6101797f000000000000000000000000000000000000000000000000000000000000000081565b6101445f1981565b6102a7610293366004611251565b60016020525f908152604090205460ff1681565b6040519015158152602001610114565b6101797f000000000000000000000000000000000000000000000000000000000000000081565b610144610a72565b60408051606080820183525f80835260208301529181019190915260025f61030d84610b2b565b815260208082019290925260409081015f20815160608101835281546001600160a01b03168152600182015493810193909352600281018054919284019161035490611268565b80601f016020809104026020016040519081016040528092919081815260200182805461038090611268565b80156103cb5780601f106103a2576101008083540402835291602001916103cb565b820191905f5260205f20905b8154815290600101906020018083116103ae57829003601f168201915b5050505050815250509050919050565b81516103e681610b94565b6104035760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815283906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc758906104519084906004016112a0565b602060405180830381865afa15801561046c573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061049091906112c6565b6104ad57604051631fb1705560e21b815260040160405180910390fd5b5f6104b785610b2b565b5f8181526001602052604090205490915060ff16156104e957604051630a81ab1560e11b815260040160405180910390fd5b604051631f3ff92360e21b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690637cffe48c906105379089906004016112a0565b602060405180830381865afa158015610552573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061057691906112e5565b90505f81600281111561058b5761058b611303565b036105a957604051633104b8e760e01b815260040160405180910390fd5b604051631b59006f60e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906336b200de906105f59089906004016112a0565b602060405180830381865afa158015610610573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061063491906112c6565b6106515760405163d0147d2d60e01b815260040160405180910390fd5b604051631a61dd7160e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d30eeb889061069d9089906004016112a0565b5f60405180830381865afa1580156106b7573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106de9190810190611322565b50506106f6336106ee883361084c565b875f19610c3e565b5f828152600160208181526040808420805460ff19168417905580516060810182523381527f00000000000000000000000000000000000000000000000000000000000000008184019081528183018b815288875260029485905292909520815181546001600160a01b0319166001600160a01b03909116178155945193850193909355519192919082019061078c908261150f565b50506040805188516001600160a01b031681526020808a015163ffffffff16908201527f0000000000000000000000000000000000000000000000000000000000000000925033910160405180910390207fab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc88860405161080c919061120d565b60405180910390a4505050505050565b60606108477f0000000000000000000000000000000000000000000000000000000000000000610c96565b905090565b8151602080840151604080517f21081bb395a13368a22c061a9d5d4bf6ddd7a9ae84a22a2a48c19d3f2b58a923938101939093527f0000000000000000000000000000000000000000000000000000000000000000908301526001600160a01b03928316606083015263ffffffff16608082015290821660a08201525f1960c08201525f906108f39060e00160405160208183030381529060405280519060200120610cd3565b9392505050565b805161090581610b94565b6109225760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b815282906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc758906109709084906004016112a0565b602060405180830381865afa15801561098b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109af91906112c6565b6109cc57604051631fb1705560e21b815260040160405180910390fd5b5f6109d684610b2b565b5f8181526001602052604090205490915060ff16610a0757604051631d171d6360e11b815260040160405180910390fd5b5f818152600160209081526040808320805460ff19169055805187516001600160a01b031681528783015163ffffffff1692810192909252805191829003018120917f1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c4598891a250505050565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b6020909101525f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea610adf610d19565b805160209182012060408051928301949094529281019190915260608101919091524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f815f0151826020015163ffffffff16604051602001610b7692919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b604051602081830303815290604052610b8e906115ca565b92915050565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af1158015610c1a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b8e91906112c6565b42811015610c5f57604051630819bdcd60e01b815260040160405180910390fd5b610c736001600160a01b0385168484610d8e565b610c9057604051638baa579f60e01b815260040160405180910390fd5b50505050565b60605f610ca283610dec565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f610cdc610a72565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60605f610d457f0000000000000000000000000000000000000000000000000000000000000000610c96565b9050805f81518110610d5957610d596115ed565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b5f5f5f610d9b8585610e13565b90925090505f816004811115610db357610db3611303565b148015610dd15750856001600160a01b0316826001600160a01b0316145b80610de25750610de2868686610e55565b9695505050505050565b5f60ff8216601f811115610b8e57604051632cd44ac360e21b815260040160405180910390fd5b5f5f8251604103610e47576020830151604084015160608501515f1a610e3b87828585610f3c565b94509450505050610e4e565b505f905060025b9250929050565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401610e7d929190611601565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051610ebb9190611619565b5f60405180830381855afa9150503d805f8114610ef3576040519150601f19603f3d011682016040523d82523d5f602084013e610ef8565b606091505b5091509150818015610f0c57506020815110155b8015610de257508051630b135d3f60e11b90610f31908301602090810190840161162f565b149695505050505050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610f7157505f90506003610ff0565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610fc2573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610fea575f60019250925050610ff0565b91505f90505b94509492505050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff8111828210171561103057611030610ff9565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561105f5761105f610ff9565b604052919050565b80356001600160a01b038116811461107d575f5ffd5b919050565b63ffffffff81168114611093575f5ffd5b50565b5f604082840312156110a6575f5ffd5b6110ae61100d565b90506110b982611067565b815260208201356110c981611082565b602082015292915050565b5f604082840312156110e4575f5ffd5b6108f38383611096565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6020815260018060a01b038251166020820152602082015160408201525f604083015160608084015261115260808401826110ee565b949350505050565b5f67ffffffffffffffff82111561117357611173610ff9565b50601f01601f191660200190565b5f5f60608385031215611192575f5ffd5b61119c8484611096565b9150604083013567ffffffffffffffff8111156111b7575f5ffd5b8301601f810185136111c7575f5ffd5b80356111da6111d58261115a565b611036565b8181528660208385010111156111ee575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b602081525f6108f360208301846110ee565b5f5f60608385031215611230575f5ffd5b61123a8484611096565b915061124860408401611067565b90509250929050565b5f60208284031215611261575f5ffd5b5035919050565b600181811c9082168061127c57607f821691505b60208210810361129a57634e487b7160e01b5f52602260045260245ffd5b50919050565b81516001600160a01b0316815260209182015163ffffffff169181019190915260400190565b5f602082840312156112d6575f5ffd5b815180151581146108f3575f5ffd5b5f602082840312156112f5575f5ffd5b8151600381106108f3575f5ffd5b634e487b7160e01b5f52602160045260245ffd5b805161107d81611082565b5f5f60408385031215611333575f5ffd5b82519150602083015167ffffffffffffffff811115611350575f5ffd5b830160408186031215611361575f5ffd5b61136961100d565b815167ffffffffffffffff81111561137f575f5ffd5b8201601f8101871361138f575f5ffd5b805167ffffffffffffffff8111156113a9576113a9610ff9565b8060051b6113b960208201611036565b9182526020818401810192908101908a8411156113d4575f5ffd5b6020850192505b8383101561149f57825167ffffffffffffffff8111156113f9575f5ffd5b85016040818d03601f1901121561140e575f5ffd5b61141661100d565b60208201518152604082015167ffffffffffffffff811115611436575f5ffd5b6020818401019250508c601f83011261144d575f5ffd5b815161145b6111d58261115a565b8181528e602083860101111561146f575f5ffd5b8160208501602083015e5f60208383010152806020840152505080845250506020820191506020830192506113db565b8552506114b29250505060208301611317565b602082015280925050509250929050565b601f82111561150a57805f5260205f20601f840160051c810160208510156114e85750805b601f840160051c820191505b81811015611507575f81556001016114f4565b50505b505050565b815167ffffffffffffffff81111561152957611529610ff9565b61153d816115378454611268565b846114c3565b6020601f82116001811461156f575f83156115585750848201515b5f19600385901b1c1916600184901b178455611507565b5f84815260208120601f198516915b8281101561159e578785015182556020948501946001909201910161157e565b50848210156115bb57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b8051602080830151919081101561129a575f1960209190910360031b1b16919050565b634e487b7160e01b5f52603260045260245ffd5b828152604060208201525f61115260408301846110ee565b5f82518060208501845e5f920191825250919050565b5f6020828403121561163f575f5ffd5b505191905056fea264697066735822122081bed0519528ea29183a4a321299f0504c747b0333e4643408403bab5e2553e164736f6c634300081b0033",
}

// ComputeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ComputeRegistryMetaData.ABI instead.
var ComputeRegistryABI = ComputeRegistryMetaData.ABI

// ComputeRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ComputeRegistryMetaData.Bin instead.
var ComputeRegistryBin = ComputeRegistryMetaData.Bin

// DeployComputeRegistry deploys a new Ethereum contract, binding an instance of ComputeRegistry to it.
func DeployComputeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _releaseManager common.Address, _allocationManager common.Address, _keyRegistrar common.Address, _crossChainRegistry common.Address, _permissionController common.Address, _tosHash [32]byte, _version string) (common.Address, *types.Transaction, *ComputeRegistry, error) {
	parsed, err := ComputeRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ComputeRegistryBin), backend, _releaseManager, _allocationManager, _keyRegistrar, _crossChainRegistry, _permissionController, _tosHash, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ComputeRegistry{ComputeRegistryCaller: ComputeRegistryCaller{contract: contract}, ComputeRegistryTransactor: ComputeRegistryTransactor{contract: contract}, ComputeRegistryFilterer: ComputeRegistryFilterer{contract: contract}}, nil
}

// ComputeRegistry is an auto generated Go binding around an Ethereum contract.
type ComputeRegistry struct {
	ComputeRegistryCaller     // Read-only binding to the contract
	ComputeRegistryTransactor // Write-only binding to the contract
	ComputeRegistryFilterer   // Log filterer for contract events
}

// ComputeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComputeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComputeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComputeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComputeRegistrySession struct {
	Contract     *ComputeRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComputeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComputeRegistryCallerSession struct {
	Contract *ComputeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ComputeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComputeRegistryTransactorSession struct {
	Contract     *ComputeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ComputeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComputeRegistryRaw struct {
	Contract *ComputeRegistry // Generic contract binding to access the raw methods on
}

// ComputeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComputeRegistryCallerRaw struct {
	Contract *ComputeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ComputeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComputeRegistryTransactorRaw struct {
	Contract *ComputeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComputeRegistry creates a new instance of ComputeRegistry, bound to a specific deployed contract.
func NewComputeRegistry(address common.Address, backend bind.ContractBackend) (*ComputeRegistry, error) {
	contract, err := bindComputeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistry{ComputeRegistryCaller: ComputeRegistryCaller{contract: contract}, ComputeRegistryTransactor: ComputeRegistryTransactor{contract: contract}, ComputeRegistryFilterer: ComputeRegistryFilterer{contract: contract}}, nil
}

// NewComputeRegistryCaller creates a new read-only instance of ComputeRegistry, bound to a specific deployed contract.
func NewComputeRegistryCaller(address common.Address, caller bind.ContractCaller) (*ComputeRegistryCaller, error) {
	contract, err := bindComputeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryCaller{contract: contract}, nil
}

// NewComputeRegistryTransactor creates a new write-only instance of ComputeRegistry, bound to a specific deployed contract.
func NewComputeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ComputeRegistryTransactor, error) {
	contract, err := bindComputeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryTransactor{contract: contract}, nil
}

// NewComputeRegistryFilterer creates a new log filterer instance of ComputeRegistry, bound to a specific deployed contract.
func NewComputeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ComputeRegistryFilterer, error) {
	contract, err := bindComputeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryFilterer{contract: contract}, nil
}

// bindComputeRegistry binds a generic wrapper to an already deployed contract.
func bindComputeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComputeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeRegistry *ComputeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeRegistry.Contract.ComputeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeRegistry *ComputeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.ComputeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeRegistry *ComputeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.ComputeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeRegistry *ComputeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeRegistry *ComputeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeRegistry *ComputeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_ComputeRegistry *ComputeRegistryCaller) ALLOCATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "ALLOCATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_ComputeRegistry *ComputeRegistrySession) ALLOCATIONMANAGER() (common.Address, error) {
	return _ComputeRegistry.Contract.ALLOCATIONMANAGER(&_ComputeRegistry.CallOpts)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_ComputeRegistry *ComputeRegistryCallerSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _ComputeRegistry.Contract.ALLOCATIONMANAGER(&_ComputeRegistry.CallOpts)
}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_ComputeRegistry *ComputeRegistryCaller) CROSSCHAINREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "CROSS_CHAIN_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_ComputeRegistry *ComputeRegistrySession) CROSSCHAINREGISTRY() (common.Address, error) {
	return _ComputeRegistry.Contract.CROSSCHAINREGISTRY(&_ComputeRegistry.CallOpts)
}

// CROSSCHAINREGISTRY is a free data retrieval call binding the contract method 0x9b250844.
//
// Solidity: function CROSS_CHAIN_REGISTRY() view returns(address)
func (_ComputeRegistry *ComputeRegistryCallerSession) CROSSCHAINREGISTRY() (common.Address, error) {
	return _ComputeRegistry.Contract.CROSSCHAINREGISTRY(&_ComputeRegistry.CallOpts)
}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_ComputeRegistry *ComputeRegistryCaller) KEYREGISTRAR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "KEY_REGISTRAR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_ComputeRegistry *ComputeRegistrySession) KEYREGISTRAR() (common.Address, error) {
	return _ComputeRegistry.Contract.KEYREGISTRAR(&_ComputeRegistry.CallOpts)
}

// KEYREGISTRAR is a free data retrieval call binding the contract method 0xe6414b48.
//
// Solidity: function KEY_REGISTRAR() view returns(address)
func (_ComputeRegistry *ComputeRegistryCallerSession) KEYREGISTRAR() (common.Address, error) {
	return _ComputeRegistry.Contract.KEYREGISTRAR(&_ComputeRegistry.CallOpts)
}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_ComputeRegistry *ComputeRegistryCaller) MAXEXPIRY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "MAX_EXPIRY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_ComputeRegistry *ComputeRegistrySession) MAXEXPIRY() (*big.Int, error) {
	return _ComputeRegistry.Contract.MAXEXPIRY(&_ComputeRegistry.CallOpts)
}

// MAXEXPIRY is a free data retrieval call binding the contract method 0xb9671690.
//
// Solidity: function MAX_EXPIRY() view returns(uint256)
func (_ComputeRegistry *ComputeRegistryCallerSession) MAXEXPIRY() (*big.Int, error) {
	return _ComputeRegistry.Contract.MAXEXPIRY(&_ComputeRegistry.CallOpts)
}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_ComputeRegistry *ComputeRegistryCaller) RELEASEMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "RELEASE_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_ComputeRegistry *ComputeRegistrySession) RELEASEMANAGER() (common.Address, error) {
	return _ComputeRegistry.Contract.RELEASEMANAGER(&_ComputeRegistry.CallOpts)
}

// RELEASEMANAGER is a free data retrieval call binding the contract method 0xb39d254f.
//
// Solidity: function RELEASE_MANAGER() view returns(address)
func (_ComputeRegistry *ComputeRegistryCallerSession) RELEASEMANAGER() (common.Address, error) {
	return _ComputeRegistry.Contract.RELEASEMANAGER(&_ComputeRegistry.CallOpts)
}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCaller) TOSAGREEMENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "TOS_AGREEMENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistrySession) TOSAGREEMENTTYPEHASH() ([32]byte, error) {
	return _ComputeRegistry.Contract.TOSAGREEMENTTYPEHASH(&_ComputeRegistry.CallOpts)
}

// TOSAGREEMENTTYPEHASH is a free data retrieval call binding the contract method 0x1de02dbb.
//
// Solidity: function TOS_AGREEMENT_TYPEHASH() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCallerSession) TOSAGREEMENTTYPEHASH() ([32]byte, error) {
	return _ComputeRegistry.Contract.TOSAGREEMENTTYPEHASH(&_ComputeRegistry.CallOpts)
}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCaller) TOSHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "TOS_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistrySession) TOSHASH() ([32]byte, error) {
	return _ComputeRegistry.Contract.TOSHASH(&_ComputeRegistry.CallOpts)
}

// TOSHASH is a free data retrieval call binding the contract method 0x8df643c7.
//
// Solidity: function TOS_HASH() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCallerSession) TOSHASH() ([32]byte, error) {
	return _ComputeRegistry.Contract.TOSHASH(&_ComputeRegistry.CallOpts)
}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCaller) CalculateTOSAgreementDigest(opts *bind.CallOpts, operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "calculateTOSAgreementDigest", operatorSet, signer)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_ComputeRegistry *ComputeRegistrySession) CalculateTOSAgreementDigest(operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	return _ComputeRegistry.Contract.CalculateTOSAgreementDigest(&_ComputeRegistry.CallOpts, operatorSet, signer)
}

// CalculateTOSAgreementDigest is a free data retrieval call binding the contract method 0x7a1bb660.
//
// Solidity: function calculateTOSAgreementDigest((address,uint32) operatorSet, address signer) view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCallerSession) CalculateTOSAgreementDigest(operatorSet OperatorSet, signer common.Address) ([32]byte, error) {
	return _ComputeRegistry.Contract.CalculateTOSAgreementDigest(&_ComputeRegistry.CallOpts, operatorSet, signer)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistrySession) DomainSeparator() ([32]byte, error) {
	return _ComputeRegistry.Contract.DomainSeparator(&_ComputeRegistry.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ComputeRegistry *ComputeRegistryCallerSession) DomainSeparator() ([32]byte, error) {
	return _ComputeRegistry.Contract.DomainSeparator(&_ComputeRegistry.CallOpts)
}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_ComputeRegistry *ComputeRegistryCaller) GetOperatorSetTosSignature(opts *bind.CallOpts, operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "getOperatorSetTosSignature", operatorSet)

	if err != nil {
		return *new(IComputeRegistryTypesTOSSignature), err
	}

	out0 := *abi.ConvertType(out[0], new(IComputeRegistryTypesTOSSignature)).(*IComputeRegistryTypesTOSSignature)

	return out0, err

}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_ComputeRegistry *ComputeRegistrySession) GetOperatorSetTosSignature(operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	return _ComputeRegistry.Contract.GetOperatorSetTosSignature(&_ComputeRegistry.CallOpts, operatorSet)
}

// GetOperatorSetTosSignature is a free data retrieval call binding the contract method 0x130d6165.
//
// Solidity: function getOperatorSetTosSignature((address,uint32) operatorSet) view returns((address,bytes32,bytes))
func (_ComputeRegistry *ComputeRegistryCallerSession) GetOperatorSetTosSignature(operatorSet OperatorSet) (IComputeRegistryTypesTOSSignature, error) {
	return _ComputeRegistry.Contract.GetOperatorSetTosSignature(&_ComputeRegistry.CallOpts, operatorSet)
}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_ComputeRegistry *ComputeRegistryCaller) IsOperatorSetRegistered(opts *bind.CallOpts, operatorSetKey [32]byte) (bool, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "isOperatorSetRegistered", operatorSetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_ComputeRegistry *ComputeRegistrySession) IsOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _ComputeRegistry.Contract.IsOperatorSetRegistered(&_ComputeRegistry.CallOpts, operatorSetKey)
}

// IsOperatorSetRegistered is a free data retrieval call binding the contract method 0xc4a1ca05.
//
// Solidity: function isOperatorSetRegistered(bytes32 operatorSetKey) view returns(bool isRegistered)
func (_ComputeRegistry *ComputeRegistryCallerSession) IsOperatorSetRegistered(operatorSetKey [32]byte) (bool, error) {
	return _ComputeRegistry.Contract.IsOperatorSetRegistered(&_ComputeRegistry.CallOpts, operatorSetKey)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeRegistry *ComputeRegistryCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeRegistry *ComputeRegistrySession) PermissionController() (common.Address, error) {
	return _ComputeRegistry.Contract.PermissionController(&_ComputeRegistry.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeRegistry *ComputeRegistryCallerSession) PermissionController() (common.Address, error) {
	return _ComputeRegistry.Contract.PermissionController(&_ComputeRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeRegistry *ComputeRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ComputeRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeRegistry *ComputeRegistrySession) Version() (string, error) {
	return _ComputeRegistry.Contract.Version(&_ComputeRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeRegistry *ComputeRegistryCallerSession) Version() (string, error) {
	return _ComputeRegistry.Contract.Version(&_ComputeRegistry.CallOpts)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_ComputeRegistry *ComputeRegistryTransactor) DeregisterFromCompute(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _ComputeRegistry.contract.Transact(opts, "deregisterFromCompute", operatorSet)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_ComputeRegistry *ComputeRegistrySession) DeregisterFromCompute(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.DeregisterFromCompute(&_ComputeRegistry.TransactOpts, operatorSet)
}

// DeregisterFromCompute is a paid mutator transaction binding the contract method 0x89fec15f.
//
// Solidity: function deregisterFromCompute((address,uint32) operatorSet) returns()
func (_ComputeRegistry *ComputeRegistryTransactorSession) DeregisterFromCompute(operatorSet OperatorSet) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.DeregisterFromCompute(&_ComputeRegistry.TransactOpts, operatorSet)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_ComputeRegistry *ComputeRegistryTransactor) RegisterForCompute(opts *bind.TransactOpts, operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _ComputeRegistry.contract.Transact(opts, "registerForCompute", operatorSet, signature)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_ComputeRegistry *ComputeRegistrySession) RegisterForCompute(operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.RegisterForCompute(&_ComputeRegistry.TransactOpts, operatorSet, signature)
}

// RegisterForCompute is a paid mutator transaction binding the contract method 0x536b2353.
//
// Solidity: function registerForCompute((address,uint32) operatorSet, bytes signature) returns()
func (_ComputeRegistry *ComputeRegistryTransactorSession) RegisterForCompute(operatorSet OperatorSet, signature []byte) (*types.Transaction, error) {
	return _ComputeRegistry.Contract.RegisterForCompute(&_ComputeRegistry.TransactOpts, operatorSet, signature)
}

// ComputeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ComputeRegistry contract.
type ComputeRegistryInitializedIterator struct {
	Event *ComputeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ComputeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeRegistryInitialized)
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
		it.Event = new(ComputeRegistryInitialized)
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
func (it *ComputeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeRegistryInitialized represents a Initialized event raised by the ComputeRegistry contract.
type ComputeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ComputeRegistry *ComputeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ComputeRegistryInitializedIterator, error) {

	logs, sub, err := _ComputeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryInitializedIterator{contract: _ComputeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ComputeRegistry *ComputeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ComputeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ComputeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeRegistryInitialized)
				if err := _ComputeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ComputeRegistry *ComputeRegistryFilterer) ParseInitialized(log types.Log) (*ComputeRegistryInitialized, error) {
	event := new(ComputeRegistryInitialized)
	if err := _ComputeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeRegistryOperatorSetDeregisteredIterator is returned from FilterOperatorSetDeregistered and is used to iterate over the raw logs and unpacked data for OperatorSetDeregistered events raised by the ComputeRegistry contract.
type ComputeRegistryOperatorSetDeregisteredIterator struct {
	Event *ComputeRegistryOperatorSetDeregistered // Event containing the contract specifics and raw log

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
func (it *ComputeRegistryOperatorSetDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeRegistryOperatorSetDeregistered)
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
		it.Event = new(ComputeRegistryOperatorSetDeregistered)
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
func (it *ComputeRegistryOperatorSetDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeRegistryOperatorSetDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeRegistryOperatorSetDeregistered represents a OperatorSetDeregistered event raised by the ComputeRegistry contract.
type ComputeRegistryOperatorSetDeregistered struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetDeregistered is a free log retrieval operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_ComputeRegistry *ComputeRegistryFilterer) FilterOperatorSetDeregistered(opts *bind.FilterOpts, operatorSet []OperatorSet) (*ComputeRegistryOperatorSetDeregisteredIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ComputeRegistry.contract.FilterLogs(opts, "OperatorSetDeregistered", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryOperatorSetDeregisteredIterator{contract: _ComputeRegistry.contract, event: "OperatorSetDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorSetDeregistered is a free log subscription operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_ComputeRegistry *ComputeRegistryFilterer) WatchOperatorSetDeregistered(opts *bind.WatchOpts, sink chan<- *ComputeRegistryOperatorSetDeregistered, operatorSet []OperatorSet) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ComputeRegistry.contract.WatchLogs(opts, "OperatorSetDeregistered", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeRegistryOperatorSetDeregistered)
				if err := _ComputeRegistry.contract.UnpackLog(event, "OperatorSetDeregistered", log); err != nil {
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

// ParseOperatorSetDeregistered is a log parse operation binding the contract event 0x1d9dd94cbbd6d9d2bc86468a197493d2b960bfc33eb09ec2c9f1731e60c45988.
//
// Solidity: event OperatorSetDeregistered((address,uint32) indexed operatorSet)
func (_ComputeRegistry *ComputeRegistryFilterer) ParseOperatorSetDeregistered(log types.Log) (*ComputeRegistryOperatorSetDeregistered, error) {
	event := new(ComputeRegistryOperatorSetDeregistered)
	if err := _ComputeRegistry.contract.UnpackLog(event, "OperatorSetDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeRegistryOperatorSetRegisteredIterator is returned from FilterOperatorSetRegistered and is used to iterate over the raw logs and unpacked data for OperatorSetRegistered events raised by the ComputeRegistry contract.
type ComputeRegistryOperatorSetRegisteredIterator struct {
	Event *ComputeRegistryOperatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *ComputeRegistryOperatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeRegistryOperatorSetRegistered)
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
		it.Event = new(ComputeRegistryOperatorSetRegistered)
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
func (it *ComputeRegistryOperatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeRegistryOperatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeRegistryOperatorSetRegistered represents a OperatorSetRegistered event raised by the ComputeRegistry contract.
type ComputeRegistryOperatorSetRegistered struct {
	OperatorSet OperatorSet
	Signer      common.Address
	TosHash     [32]byte
	Signature   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetRegistered is a free log retrieval operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_ComputeRegistry *ComputeRegistryFilterer) FilterOperatorSetRegistered(opts *bind.FilterOpts, operatorSet []OperatorSet, signer []common.Address, tosHash [][32]byte) (*ComputeRegistryOperatorSetRegisteredIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var tosHashRule []interface{}
	for _, tosHashItem := range tosHash {
		tosHashRule = append(tosHashRule, tosHashItem)
	}

	logs, sub, err := _ComputeRegistry.contract.FilterLogs(opts, "OperatorSetRegistered", operatorSetRule, signerRule, tosHashRule)
	if err != nil {
		return nil, err
	}
	return &ComputeRegistryOperatorSetRegisteredIterator{contract: _ComputeRegistry.contract, event: "OperatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorSetRegistered is a free log subscription operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_ComputeRegistry *ComputeRegistryFilterer) WatchOperatorSetRegistered(opts *bind.WatchOpts, sink chan<- *ComputeRegistryOperatorSetRegistered, operatorSet []OperatorSet, signer []common.Address, tosHash [][32]byte) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var tosHashRule []interface{}
	for _, tosHashItem := range tosHash {
		tosHashRule = append(tosHashRule, tosHashItem)
	}

	logs, sub, err := _ComputeRegistry.contract.WatchLogs(opts, "OperatorSetRegistered", operatorSetRule, signerRule, tosHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeRegistryOperatorSetRegistered)
				if err := _ComputeRegistry.contract.UnpackLog(event, "OperatorSetRegistered", log); err != nil {
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

// ParseOperatorSetRegistered is a log parse operation binding the contract event 0xab9d9ee579ad2bc5e08fba9e2b3d59225f1400f1a5830745283697e35ce2dcc8.
//
// Solidity: event OperatorSetRegistered((address,uint32) indexed operatorSet, address indexed signer, bytes32 indexed tosHash, bytes signature)
func (_ComputeRegistry *ComputeRegistryFilterer) ParseOperatorSetRegistered(log types.Log) (*ComputeRegistryOperatorSetRegistered, error) {
	event := new(ComputeRegistryOperatorSetRegistered)
	if err := _ComputeRegistry.contract.UnpackLog(event, "OperatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
