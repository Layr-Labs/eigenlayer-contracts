// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OperatorSetManager

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

// IOperatorSetManagerOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type IOperatorSetManagerOperatorSet struct {
	Avs common.Address
	Id  [4]byte
}

// IOperatorSetManagerSlashingMagnitudeParameters is an auto generated low-level Go binding around an user-defined struct.
type IOperatorSetManagerSlashingMagnitudeParameters struct {
	Strategy            common.Address
	TotalMagnitude      uint64
	OperatorSets        []IOperatorSetManagerOperatorSet
	SlashableMagnitudes []uint64
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// OperatorSetManagerMetaData contains all meta data concerning the OperatorSetManager contract.
var OperatorSetManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSlashableBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockMagnitudeUpdatesAtEpoch\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSlashingParameters\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashingMagnitudeParameters\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorSetManager.SlashingMagnitudeParameters[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"totalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structIOperatorSetManager.OperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"slashableMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"effectEpoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"SlashableMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"slashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectEpoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"totalMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectEpoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161172638038061172683398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516116956100916000396000818160a90152610ab301526116956000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633f76c6c714610051578063516b98831461007c578063b1344271146100a4578063fa88909b146100e3575b600080fd5b61006461005f366004611194565b6100f8565b60405161ffff90911681526020015b60405180910390f35b61008f61008a3660046111f8565b6104fe565b60405163ffffffff9091168152602001610073565b6100cb7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610073565b6100f66100f13660046112a4565b610a9c565b005b6000610102610d53565b63ffffffff168263ffffffff1611156101745760405162461bcd60e51b815260206004820152602960248201527f45706f6368206973206d6f7265207468616e20322065706f63687320696e207460448201526868652066757475726560b81b60648201526084015b60405180910390fd5b6001600160a01b038086166000908152600260209081526040808320938716835292905220545b80156102d8576001600160a01b03808716600090815260026020908152604080832093881683529290522063ffffffff8416906101d96001846112f3565b815481106101e9576101e961130a565b60009182526020909120015463ffffffff161180159061026857506001600160a01b03808716600090815260026020908152604080832093881683529290522063ffffffff84169061023c6001846112f3565b8154811061024c5761024c61130a565b600091825260209091200154600160201b900463ffffffff1610155b156102c6576001600160a01b03808716600090815260026020908152604080832093881683529290522061029d6001836112f3565b815481106102ad576102ad61130a565b60009182526020909120015463ffffffff1692506102d8565b806102d081611320565b91505061019b565b506001600160a01b0380861660009081526020818152604080832093871683529290529081208161031661031136899003890189611354565b610d6d565b81526020810191909152604001600090812080549092505b80156103bf5763ffffffff8516836103476001846112f3565b815481106103575761035761130a565b60009182526020909120015463ffffffff16116103ad578261037a6001836112f3565b8154811061038a5761038a61130a565b600091825260209091200154600160201b90046001600160401b031691506103bf565b806103b781611320565b91505061032e565b506001600160a01b0380881660009081526001602090815260408083209389168352929052908120545b80156104c6576001600160a01b03808a166000908152600160208181526040808420948c16845293905291902063ffffffff88169161042890846112f3565b815481106104385761043861130a565b60009182526020909120015463ffffffff16116104b4576001600160a01b03808a166000908152600160208181526040808420948c1684529390529190209061048190836112f3565b815481106104915761049161130a565b600091825260209091200154600160201b90046001600160401b031691506104c6565b806104be81611320565b9150506103e9565b50806001600160c01b0316612710836001600160c01b03166104e891906113be565b6104f291906113dd565b98975050505050505050565b6000610508610d53565b905060005b83811015610a9357368585838181106105285761052861130a565b905060200281019061053a91906113ff565b9050600061054b602083018361141f565b905061055a6060830183611443565b90506105696040840184611493565b9050146106195760405162461bcd60e51b815260206004820152606c60248201527f4f70657261746f725365744d616e616765722e7570646174654f70657261746f60448201527f72536574536c617368696e67506172616d65746572733a206f70657261746f7260648201527f5365747320616e6420736c61736861626c654d61676e697475646573206c656e60848201526b0cee8d040dad2e6dac2e8c6d60a31b60a482015260c40161016b565b60408051606081018252600080825260208083018290528284018290526001600160a01b038c811683526001825284832090861683529052919091205480156106ea576001600160a01b03808b16600090815260016020818152604080842094881684529390529190209061068e90836112f3565b8154811061069e5761069e61130a565b600091825260209182902060408051606081018252919092015463ffffffff811682526001600160401b03600160201b8204811694830194909452600160601b90049092169082015291505b60005b6106fa6040860186611493565b90508110156107c1576000806107698d8a8861071960408c018c611493565b888181106107295761072961130a565b9050604002018b806060019061073f9190611443565b8981811061074f5761074f61130a565b905060200201602081019061076491906114dc565b610dfe565b91509150818560400181815161077f9190611505565b6001600160401b031690525060408501805182919061079f90839061152d565b6001600160401b03169052508291506107b9905081611558565b9150506106ed565b506107d260408501602086016114dc565b6001600160401b031682604001516001600160401b031611156108925760405162461bcd60e51b815260206004820152606660248201527f4f70657261746f725365744d616e616765722e7570646174654f70657261746f60448201527f72536574536c617368696e67506172616d65746572733a20746f74616c416c6c60648201527f6f63617465644d61676e6974756465206578636565647320746f74616c4d61676084820152656e697475646560d01b60a482015260c40161016b565b6108a260408501602086016114dc565b6001600160401b03166020830152815163ffffffff87811691161415610975576001600160a01b03808b166000908152600160208181526040808420948816845293905291902083916108f590846112f3565b815481106109055761090561130a565b6000918252602091829020835191018054928401516040909401516001600160401b03908116600160601b0267ffffffffffffffff60601b1991909516600160201b026bffffffffffffffffffffffff1990941663ffffffff909316929092179290921716919091179055610a0b565b63ffffffff86811683526001600160a01b038b811660009081526001602081815260408084209489168452938152838320805492830181558352918290208651910180549287015193870151919094166bffffffffffffffffffffffff1990921691909117600160201b6001600160401b03938416021767ffffffffffffffff60601b1916600160601b92909116919091021790555b7f802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd68a84610a3e60408801602089016114dc565b604080516001600160a01b0394851681529390921660208401526001600160401b03169082015263ffffffff8816606082015260800160405180910390a1505050508080610a8b90611558565b91505061050d565b50949350505050565b6000610aa66110c0565b9050336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b545760405162461bcd60e51b815260206004820152604560248201527f4f70657261746f725365744d616e616765722e6c6f636b5374616b655570646160448201527f746573417445706f63683a2043616c6c6572206973206e6f742074686520736c60648201526430b9b432b960d91b608482015260a40161016b565b6001600160a01b038084166000908152600260209081526040808320938616835292905220548015610cd8576001600160a01b03808516600090815260026020908152604080832093871683529290522063ffffffff831690610bb86001846112f3565b81548110610bc857610bc861130a565b600091825260209091200154600160201b900463ffffffff161415610bed5750505050565b6001600160a01b03808516600090815260026020908152604080832093871683529290522063ffffffff831690610c256001846112f3565b81548110610c3557610c3561130a565b600091825260209091200154610c5990600160201b900463ffffffff166001611573565b63ffffffff161415610cd8576001600160a01b0380851660009081526002602090815260408083209387168352929052208290610c976001846112f3565b81548110610ca757610ca761130a565b9060005260206000200160000160046101000a81548163ffffffff021916908363ffffffff16021790555050505050565b506001600160a01b039283166000908152600260209081526040808320949095168252928352838120845180860190955263ffffffff92831680865285850190815281546001810183559183529390912093519301805492518216600160201b0267ffffffffffffffff199093169390911692909217179055565b6000610d5d6110c0565b610d68906003611573565b905090565b604080516001808252818301909252600091829190816020015b6040805180820190915260008082526020820152815260200190600190039081610d875790505090508281600081518110610dc457610dc461130a565b602002602001018190525080604051602001610de09190611592565b60405160208183030381529060405280519060200120915050919050565b60008080610e1461031136879003870187611354565b60408051808201909152600080825260208201529091506001600160a01b03808a16600090815260208181526040808320938b168352928152828220858352905220548015610ee0576001600160a01b03808b16600090815260208181526040808320938c168352928152828220868352905220610e936001836112f3565b81548110610ea357610ea361130a565b60009182526020918290206040805180820190915291015463ffffffff81168252600160201b90046001600160401b031691810182905290955091505b8593508863ffffffff16826000015163ffffffff161415610f79576001600160a01b03808b16600090815260208181526040808320938c1683529281528282208683529052208490610f336001846112f3565b81548110610f4357610f4361130a565b9060005260206000200160000160046101000a8154816001600160401b0302191690836001600160401b03160217905550611074565b60405180604001604052808a63ffffffff168152602001856001600160401b031681525091506000808b6001600160a01b03166001600160a01b031681526020019081526020016000206000896001600160a01b03166001600160a01b031681526020019081526020016000206000848152602001908152602001600020829080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a8154816001600160401b0302191690836001600160401b0316021790555050505b7fa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a904626542995068a8989878d6040516110ab9594939291906115f4565b60405180910390a15050509550959350505050565b6000610d68426000635fc630408210156111425760405162461bcd60e51b815260206004820152603d60248201527f45706f63685574696c732e67657445706f636846726f6d54696d657374616d7060448201527f3a2074696d657374616d70206973206265666f72652067656e65736973000000606482015260840161016b565b62093a80611154635fc63040846112f3565b61115e91906113dd565b92915050565b6001600160a01b038116811461117957600080fd5b50565b60006040828403121561118e57600080fd5b50919050565b60008060008060a085870312156111aa57600080fd5b84356111b581611164565b93506111c4866020870161117c565b925060608501356111d481611164565b9150608085013563ffffffff811681146111ed57600080fd5b939692955090935050565b6000806000806060858703121561120e57600080fd5b843561121981611164565b935060208501356001600160401b038082111561123557600080fd5b818701915087601f83011261124957600080fd5b81358181111561125857600080fd5b8860208260051b850101111561126d57600080fd5b60208301955080945050604087013591508082111561128b57600080fd5b506112988782880161117c565b91505092959194509250565b600080604083850312156112b757600080fd5b82356112c281611164565b915060208301356112d281611164565b809150509250929050565b634e487b7160e01b600052601160045260246000fd5b600082821015611305576113056112dd565b500390565b634e487b7160e01b600052603260045260246000fd5b60008161132f5761132f6112dd565b506000190190565b80356001600160e01b03198116811461134f57600080fd5b919050565b60006040828403121561136657600080fd5b604051604081018181106001600160401b038211171561139657634e487b7160e01b600052604160045260246000fd5b60405282356113a481611164565b81526113b260208401611337565b60208201529392505050565b60008160001904831182151516156113d8576113d86112dd565b500290565b6000826113fa57634e487b7160e01b600052601260045260246000fd5b500490565b60008235607e1983360301811261141557600080fd5b9190910192915050565b60006020828403121561143157600080fd5b813561143c81611164565b9392505050565b6000808335601e1984360301811261145a57600080fd5b8301803591506001600160401b0382111561147457600080fd5b6020019150600581901b360382131561148c57600080fd5b9250929050565b6000808335601e198436030181126114aa57600080fd5b8301803591506001600160401b038211156114c457600080fd5b6020019150600681901b360382131561148c57600080fd5b6000602082840312156114ee57600080fd5b81356001600160401b038116811461143c57600080fd5b60006001600160401b0383811690831681811015611525576115256112dd565b039392505050565b60006001600160401b0380831681851680830382111561154f5761154f6112dd565b01949350505050565b600060001982141561156c5761156c6112dd565b5060010190565b600063ffffffff80831681851680830382111561154f5761154f6112dd565b602080825282518282018190526000919060409081850190868401855b828110156115e757815180516001600160a01b031685528601516001600160e01b0319168685015292840192908501906001016115af565b5091979650505050505050565b6001600160a01b038681168252858116602083015260c0820190853561161981611164565b1660408301526001600160e01b031961163460208701611337565b1660608301526001600160401b038416608083015263ffffffff831660a0830152969550505050505056fea2646970667358221220feecdf45e3744870acc8b2ffa5f5c333d0dd27893f61310ae55cf0b9720083fb64736f6c634300080c0033",
}

// OperatorSetManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorSetManagerMetaData.ABI instead.
var OperatorSetManagerABI = OperatorSetManagerMetaData.ABI

// OperatorSetManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OperatorSetManagerMetaData.Bin instead.
var OperatorSetManagerBin = OperatorSetManagerMetaData.Bin

// DeployOperatorSetManager deploys a new Ethereum contract, binding an instance of OperatorSetManager to it.
func DeployOperatorSetManager(auth *bind.TransactOpts, backend bind.ContractBackend, _slasher common.Address) (common.Address, *types.Transaction, *OperatorSetManager, error) {
	parsed, err := OperatorSetManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OperatorSetManagerBin), backend, _slasher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OperatorSetManager{OperatorSetManagerCaller: OperatorSetManagerCaller{contract: contract}, OperatorSetManagerTransactor: OperatorSetManagerTransactor{contract: contract}, OperatorSetManagerFilterer: OperatorSetManagerFilterer{contract: contract}}, nil
}

// OperatorSetManager is an auto generated Go binding around an Ethereum contract.
type OperatorSetManager struct {
	OperatorSetManagerCaller     // Read-only binding to the contract
	OperatorSetManagerTransactor // Write-only binding to the contract
	OperatorSetManagerFilterer   // Log filterer for contract events
}

// OperatorSetManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorSetManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSetManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorSetManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSetManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorSetManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSetManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSetManagerSession struct {
	Contract     *OperatorSetManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperatorSetManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorSetManagerCallerSession struct {
	Contract *OperatorSetManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OperatorSetManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorSetManagerTransactorSession struct {
	Contract     *OperatorSetManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OperatorSetManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorSetManagerRaw struct {
	Contract *OperatorSetManager // Generic contract binding to access the raw methods on
}

// OperatorSetManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorSetManagerCallerRaw struct {
	Contract *OperatorSetManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorSetManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorSetManagerTransactorRaw struct {
	Contract *OperatorSetManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorSetManager creates a new instance of OperatorSetManager, bound to a specific deployed contract.
func NewOperatorSetManager(address common.Address, backend bind.ContractBackend) (*OperatorSetManager, error) {
	contract, err := bindOperatorSetManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorSetManager{OperatorSetManagerCaller: OperatorSetManagerCaller{contract: contract}, OperatorSetManagerTransactor: OperatorSetManagerTransactor{contract: contract}, OperatorSetManagerFilterer: OperatorSetManagerFilterer{contract: contract}}, nil
}

// NewOperatorSetManagerCaller creates a new read-only instance of OperatorSetManager, bound to a specific deployed contract.
func NewOperatorSetManagerCaller(address common.Address, caller bind.ContractCaller) (*OperatorSetManagerCaller, error) {
	contract, err := bindOperatorSetManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorSetManagerCaller{contract: contract}, nil
}

// NewOperatorSetManagerTransactor creates a new write-only instance of OperatorSetManager, bound to a specific deployed contract.
func NewOperatorSetManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorSetManagerTransactor, error) {
	contract, err := bindOperatorSetManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorSetManagerTransactor{contract: contract}, nil
}

// NewOperatorSetManagerFilterer creates a new log filterer instance of OperatorSetManager, bound to a specific deployed contract.
func NewOperatorSetManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorSetManagerFilterer, error) {
	contract, err := bindOperatorSetManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorSetManagerFilterer{contract: contract}, nil
}

// bindOperatorSetManager binds a generic wrapper to an already deployed contract.
func bindOperatorSetManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorSetManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorSetManager *OperatorSetManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorSetManager.Contract.OperatorSetManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorSetManager *OperatorSetManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.OperatorSetManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorSetManager *OperatorSetManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.OperatorSetManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorSetManager *OperatorSetManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorSetManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorSetManager *OperatorSetManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorSetManager *OperatorSetManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.contract.Transact(opts, method, params...)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x3f76c6c7.
//
// Solidity: function getSlashableBips(address operator, (address,bytes4) operatorSet, address strategy, uint32 epoch) view returns(uint16 slashableBips)
func (_OperatorSetManager *OperatorSetManagerCaller) GetSlashableBips(opts *bind.CallOpts, operator common.Address, operatorSet IOperatorSetManagerOperatorSet, strategy common.Address, epoch uint32) (uint16, error) {
	var out []interface{}
	err := _OperatorSetManager.contract.Call(opts, &out, "getSlashableBips", operator, operatorSet, strategy, epoch)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetSlashableBips is a free data retrieval call binding the contract method 0x3f76c6c7.
//
// Solidity: function getSlashableBips(address operator, (address,bytes4) operatorSet, address strategy, uint32 epoch) view returns(uint16 slashableBips)
func (_OperatorSetManager *OperatorSetManagerSession) GetSlashableBips(operator common.Address, operatorSet IOperatorSetManagerOperatorSet, strategy common.Address, epoch uint32) (uint16, error) {
	return _OperatorSetManager.Contract.GetSlashableBips(&_OperatorSetManager.CallOpts, operator, operatorSet, strategy, epoch)
}

// GetSlashableBips is a free data retrieval call binding the contract method 0x3f76c6c7.
//
// Solidity: function getSlashableBips(address operator, (address,bytes4) operatorSet, address strategy, uint32 epoch) view returns(uint16 slashableBips)
func (_OperatorSetManager *OperatorSetManagerCallerSession) GetSlashableBips(operator common.Address, operatorSet IOperatorSetManagerOperatorSet, strategy common.Address, epoch uint32) (uint16, error) {
	return _OperatorSetManager.Contract.GetSlashableBips(&_OperatorSetManager.CallOpts, operator, operatorSet, strategy, epoch)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_OperatorSetManager *OperatorSetManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorSetManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_OperatorSetManager *OperatorSetManagerSession) Slasher() (common.Address, error) {
	return _OperatorSetManager.Contract.Slasher(&_OperatorSetManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_OperatorSetManager *OperatorSetManagerCallerSession) Slasher() (common.Address, error) {
	return _OperatorSetManager.Contract.Slasher(&_OperatorSetManager.CallOpts)
}

// LockMagnitudeUpdatesAtEpoch is a paid mutator transaction binding the contract method 0xfa88909b.
//
// Solidity: function lockMagnitudeUpdatesAtEpoch(address operator, address strategy) returns()
func (_OperatorSetManager *OperatorSetManagerTransactor) LockMagnitudeUpdatesAtEpoch(opts *bind.TransactOpts, operator common.Address, strategy common.Address) (*types.Transaction, error) {
	return _OperatorSetManager.contract.Transact(opts, "lockMagnitudeUpdatesAtEpoch", operator, strategy)
}

// LockMagnitudeUpdatesAtEpoch is a paid mutator transaction binding the contract method 0xfa88909b.
//
// Solidity: function lockMagnitudeUpdatesAtEpoch(address operator, address strategy) returns()
func (_OperatorSetManager *OperatorSetManagerSession) LockMagnitudeUpdatesAtEpoch(operator common.Address, strategy common.Address) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.LockMagnitudeUpdatesAtEpoch(&_OperatorSetManager.TransactOpts, operator, strategy)
}

// LockMagnitudeUpdatesAtEpoch is a paid mutator transaction binding the contract method 0xfa88909b.
//
// Solidity: function lockMagnitudeUpdatesAtEpoch(address operator, address strategy) returns()
func (_OperatorSetManager *OperatorSetManagerTransactorSession) LockMagnitudeUpdatesAtEpoch(operator common.Address, strategy common.Address) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.LockMagnitudeUpdatesAtEpoch(&_OperatorSetManager.TransactOpts, operator, strategy)
}

// UpdateSlashingParameters is a paid mutator transaction binding the contract method 0x516b9883.
//
// Solidity: function updateSlashingParameters(address operator, (address,uint64,(address,bytes4)[],uint64[])[] slashingMagnitudeParameters, (bytes,uint256) allocatorSignature) returns(uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerTransactor) UpdateSlashingParameters(opts *bind.TransactOpts, operator common.Address, slashingMagnitudeParameters []IOperatorSetManagerSlashingMagnitudeParameters, allocatorSignature ISignatureUtilsSignatureWithExpiry) (*types.Transaction, error) {
	return _OperatorSetManager.contract.Transact(opts, "updateSlashingParameters", operator, slashingMagnitudeParameters, allocatorSignature)
}

// UpdateSlashingParameters is a paid mutator transaction binding the contract method 0x516b9883.
//
// Solidity: function updateSlashingParameters(address operator, (address,uint64,(address,bytes4)[],uint64[])[] slashingMagnitudeParameters, (bytes,uint256) allocatorSignature) returns(uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerSession) UpdateSlashingParameters(operator common.Address, slashingMagnitudeParameters []IOperatorSetManagerSlashingMagnitudeParameters, allocatorSignature ISignatureUtilsSignatureWithExpiry) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.UpdateSlashingParameters(&_OperatorSetManager.TransactOpts, operator, slashingMagnitudeParameters, allocatorSignature)
}

// UpdateSlashingParameters is a paid mutator transaction binding the contract method 0x516b9883.
//
// Solidity: function updateSlashingParameters(address operator, (address,uint64,(address,bytes4)[],uint64[])[] slashingMagnitudeParameters, (bytes,uint256) allocatorSignature) returns(uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerTransactorSession) UpdateSlashingParameters(operator common.Address, slashingMagnitudeParameters []IOperatorSetManagerSlashingMagnitudeParameters, allocatorSignature ISignatureUtilsSignatureWithExpiry) (*types.Transaction, error) {
	return _OperatorSetManager.Contract.UpdateSlashingParameters(&_OperatorSetManager.TransactOpts, operator, slashingMagnitudeParameters, allocatorSignature)
}

// OperatorSetManagerSlashableMagnitudeUpdatedIterator is returned from FilterSlashableMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for SlashableMagnitudeUpdated events raised by the OperatorSetManager contract.
type OperatorSetManagerSlashableMagnitudeUpdatedIterator struct {
	Event *OperatorSetManagerSlashableMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *OperatorSetManagerSlashableMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSetManagerSlashableMagnitudeUpdated)
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
		it.Event = new(OperatorSetManagerSlashableMagnitudeUpdated)
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
func (it *OperatorSetManagerSlashableMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSetManagerSlashableMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSetManagerSlashableMagnitudeUpdated represents a SlashableMagnitudeUpdated event raised by the OperatorSetManager contract.
type OperatorSetManagerSlashableMagnitudeUpdated struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        IOperatorSetManagerOperatorSet
	SlashableMagnitude uint64
	EffectEpoch        uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSlashableMagnitudeUpdated is a free log retrieval operation binding the contract event 0xa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a90462654299506.
//
// Solidity: event SlashableMagnitudeUpdated(address operator, address strategy, (address,bytes4) operatorSet, uint64 slashableMagnitude, uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerFilterer) FilterSlashableMagnitudeUpdated(opts *bind.FilterOpts) (*OperatorSetManagerSlashableMagnitudeUpdatedIterator, error) {

	logs, sub, err := _OperatorSetManager.contract.FilterLogs(opts, "SlashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &OperatorSetManagerSlashableMagnitudeUpdatedIterator{contract: _OperatorSetManager.contract, event: "SlashableMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashableMagnitudeUpdated is a free log subscription operation binding the contract event 0xa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a90462654299506.
//
// Solidity: event SlashableMagnitudeUpdated(address operator, address strategy, (address,bytes4) operatorSet, uint64 slashableMagnitude, uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerFilterer) WatchSlashableMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *OperatorSetManagerSlashableMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _OperatorSetManager.contract.WatchLogs(opts, "SlashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSetManagerSlashableMagnitudeUpdated)
				if err := _OperatorSetManager.contract.UnpackLog(event, "SlashableMagnitudeUpdated", log); err != nil {
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

// ParseSlashableMagnitudeUpdated is a log parse operation binding the contract event 0xa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a90462654299506.
//
// Solidity: event SlashableMagnitudeUpdated(address operator, address strategy, (address,bytes4) operatorSet, uint64 slashableMagnitude, uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerFilterer) ParseSlashableMagnitudeUpdated(log types.Log) (*OperatorSetManagerSlashableMagnitudeUpdated, error) {
	event := new(OperatorSetManagerSlashableMagnitudeUpdated)
	if err := _OperatorSetManager.contract.UnpackLog(event, "SlashableMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSetManagerTotalMagnitudeUpdatedIterator is returned from FilterTotalMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for TotalMagnitudeUpdated events raised by the OperatorSetManager contract.
type OperatorSetManagerTotalMagnitudeUpdatedIterator struct {
	Event *OperatorSetManagerTotalMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *OperatorSetManagerTotalMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSetManagerTotalMagnitudeUpdated)
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
		it.Event = new(OperatorSetManagerTotalMagnitudeUpdated)
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
func (it *OperatorSetManagerTotalMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSetManagerTotalMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSetManagerTotalMagnitudeUpdated represents a TotalMagnitudeUpdated event raised by the OperatorSetManager contract.
type OperatorSetManagerTotalMagnitudeUpdated struct {
	Operator       common.Address
	Strategy       common.Address
	TotalMagnitude uint64
	EffectEpoch    uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalMagnitudeUpdated is a free log retrieval operation binding the contract event 0x802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd6.
//
// Solidity: event TotalMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude, uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerFilterer) FilterTotalMagnitudeUpdated(opts *bind.FilterOpts) (*OperatorSetManagerTotalMagnitudeUpdatedIterator, error) {

	logs, sub, err := _OperatorSetManager.contract.FilterLogs(opts, "TotalMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &OperatorSetManagerTotalMagnitudeUpdatedIterator{contract: _OperatorSetManager.contract, event: "TotalMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalMagnitudeUpdated is a free log subscription operation binding the contract event 0x802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd6.
//
// Solidity: event TotalMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude, uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerFilterer) WatchTotalMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *OperatorSetManagerTotalMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _OperatorSetManager.contract.WatchLogs(opts, "TotalMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSetManagerTotalMagnitudeUpdated)
				if err := _OperatorSetManager.contract.UnpackLog(event, "TotalMagnitudeUpdated", log); err != nil {
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

// ParseTotalMagnitudeUpdated is a log parse operation binding the contract event 0x802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd6.
//
// Solidity: event TotalMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude, uint32 effectEpoch)
func (_OperatorSetManager *OperatorSetManagerFilterer) ParseTotalMagnitudeUpdated(log types.Log) (*OperatorSetManagerTotalMagnitudeUpdated, error) {
	event := new(OperatorSetManagerTotalMagnitudeUpdated)
	if err := _OperatorSetManager.contract.UnpackLog(event, "TotalMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
