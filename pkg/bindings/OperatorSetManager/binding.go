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
	Bin: "0x60a060405234801561001057600080fd5b50604051611b52380380611b5283398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051611ac16100916000396000818160a901526111a10152611ac16000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633f76c6c714610051578063516b98831461007c578063b1344271146100a4578063fa88909b146100e3575b600080fd5b61006461005f3660046115c0565b6100f8565b60405161ffff90911681526020015b60405180910390f35b61008f61008a366004611624565b6104fe565b60405163ffffffff9091168152602001610073565b6100cb7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610073565b6100f66100f13660046116d0565b61118a565b005b6000610102611441565b63ffffffff168263ffffffff1611156101745760405162461bcd60e51b815260206004820152602960248201527f45706f6368206973206d6f7265207468616e20322065706f63687320696e207460448201526868652066757475726560b81b60648201526084015b60405180910390fd5b6001600160a01b038086166000908152600260209081526040808320938716835292905220545b80156102d8576001600160a01b03808716600090815260026020908152604080832093881683529290522063ffffffff8416906101d960018461171f565b815481106101e9576101e9611736565b60009182526020909120015463ffffffff161180159061026857506001600160a01b03808716600090815260026020908152604080832093881683529290522063ffffffff84169061023c60018461171f565b8154811061024c5761024c611736565b600091825260209091200154600160201b900463ffffffff1610155b156102c6576001600160a01b03808716600090815260026020908152604080832093881683529290522061029d60018361171f565b815481106102ad576102ad611736565b60009182526020909120015463ffffffff1692506102d8565b806102d08161174c565b91505061019b565b506001600160a01b0380861660009081526020818152604080832093871683529290529081208161031661031136899003890189611780565b61145b565b81526020810191909152604001600090812080549092505b80156103bf5763ffffffff85168361034760018461171f565b8154811061035757610357611736565b60009182526020909120015463ffffffff16116103ad578261037a60018361171f565b8154811061038a5761038a611736565b600091825260209091200154600160201b90046001600160401b031691506103bf565b806103b78161174c565b91505061032e565b506001600160a01b0380881660009081526001602090815260408083209389168352929052908120545b80156104c6576001600160a01b03808a166000908152600160208181526040808420948c16845293905291902063ffffffff881691610428908461171f565b8154811061043857610438611736565b60009182526020909120015463ffffffff16116104b4576001600160a01b03808a166000908152600160208181526040808420948c16845293905291902090610481908361171f565b8154811061049157610491611736565b600091825260209091200154600160201b90046001600160401b031691506104c6565b806104be8161174c565b9150506103e9565b50806001600160c01b0316612710836001600160c01b03166104e891906117ea565b6104f29190611809565b98975050505050505050565b6000610508611441565b905060005b838110156111815784848281811061052757610527611736565b9050602002810190610539919061182b565b61054790606081019061184b565b905085858381811061055b5761055b611736565b905060200281019061056d919061182b565b61057b90604081019061189b565b90501461062b5760405162461bcd60e51b815260206004820152606c60248201527f4f70657261746f725365744d616e616765722e7570646174654f70657261746f60448201527f72536574536c617368696e67506172616d65746572733a206f70657261746f7260648201527f5365747320616e6420736c61736861626c654d61676e697475646573206c656e60848201526b0cee8d040dad2e6dac2e8c6d60a31b60a482015260c40161016b565b60408051606081018252600080825260208201819052918101919091526001600160a01b03871660009081526001602052604081208188888681811061067357610673611736565b9050602002810190610685919061182b565b6106939060208101906118e4565b6001600160a01b0316815260208101919091526040016000205490508015610780576001600160a01b0388166000908152600160205260408120908888868181106106e0576106e0611736565b90506020028101906106f2919061182b565b6107009060208101906118e4565b6001600160a01b03168152602081019190915260400160002061072460018361171f565b8154811061073457610734611736565b600091825260209182902060408051606081018252919092015463ffffffff811682526001600160401b03600160201b8204811694830194909452600160601b90049092169082015291505b60005b87878581811061079557610795611736565b90506020028101906107a7919061182b565b6107b590604081019061189b565b9050811015610d8a57600061081b8989878181106107d5576107d5611736565b90506020028101906107e7919061182b565b6107f590604081019061189b565b8481811061080557610805611736565b9050604002018036038101906103119190611780565b60408051808201909152600080825260208201529091506001600160a01b038b166000908152602081905260408120818c8c8a81811061085d5761085d611736565b905060200281019061086f919061182b565b61087d9060208101906118e4565b6001600160a01b031681526020808201929092526040908101600090812086825290925290205490508015610973576001600160a01b038c166000908152602081905260408120908c8c8a8181106108d7576108d7611736565b90506020028101906108e9919061182b565b6108f79060208101906118e4565b6001600160a01b0316815260208082019290925260409081016000908120868252909252902061092860018361171f565b8154811061093857610938611736565b60009182526020918290206040805180820190915291015463ffffffff81168252600160201b90046001600160401b03169181019190915291505b8160200151866040018181516109899190611908565b6001600160401b03169052508a8a888181106109a7576109a7611736565b90506020028101906109b9919061182b565b6109c790606081019061184b565b858181106109d7576109d7611736565b90506020020160208101906109ec9190611930565b866040018181516109fd9190611959565b6001600160401b0316905250815163ffffffff89811691161415610b34578a8a88818110610a2d57610a2d611736565b9050602002810190610a3f919061182b565b610a4d90606081019061184b565b85818110610a5d57610a5d611736565b9050602002016020810190610a729190611930565b6001600160a01b038d166000908152602081905260408120908d8d8b818110610a9d57610a9d611736565b9050602002810190610aaf919061182b565b610abd9060208101906118e4565b6001600160a01b03168152602080820192909252604090810160009081208782529092529020610aee60018461171f565b81548110610afe57610afe611736565b9060005260206000200160000160046101000a8154816001600160401b0302191690836001600160401b03160217905550610c67565b60405180604001604052808963ffffffff1681526020018c8c8a818110610b5d57610b5d611736565b9050602002810190610b6f919061182b565b610b7d90606081019061184b565b87818110610b8d57610b8d611736565b9050602002016020810190610ba29190611930565b6001600160401b031690526001600160a01b038d1660009081526020819052604081209193508c8c8a818110610bda57610bda611736565b9050602002810190610bec919061182b565b610bfa9060208101906118e4565b6001600160a01b03168152602080820192909252604090810160009081208682528352908120805460018101825590825290829020845191018054928501516001600160401b0316600160201b026001600160601b031990931663ffffffff909216919091179190911790555b7fa051327ef1123f482ec636fa78d997e135019b2bcfa0fab32a904626542995068c8c8c8a818110610c9b57610c9b611736565b9050602002810190610cad919061182b565b610cbb9060208101906118e4565b8d8d8b818110610ccd57610ccd611736565b9050602002810190610cdf919061182b565b610ced90604081019061189b565b88818110610cfd57610cfd611736565b9050604002018e8e8c818110610d1557610d15611736565b9050602002810190610d27919061182b565b610d3590606081019061184b565b89818110610d4557610d45611736565b9050602002016020810190610d5a9190611930565b8c604051610d6c959493929190611984565b60405180910390a15050508080610d82906119ef565b915050610783565b50868684818110610d9d57610d9d611736565b9050602002810190610daf919061182b565b610dc0906040810190602001611930565b6001600160401b031682604001516001600160401b03161115610e805760405162461bcd60e51b815260206004820152606660248201527f4f70657261746f725365744d616e616765722e7570646174654f70657261746f60448201527f72536574536c617368696e67506172616d65746572733a20746f74616c416c6c60648201527f6f63617465644d61676e6974756465206578636565647320746f74616c4d61676084820152656e697475646560d01b60a482015260c40161016b565b868684818110610e9257610e92611736565b9050602002810190610ea4919061182b565b610eb5906040810190602001611930565b6001600160401b0390811660208401526040830180519091169052815163ffffffff85811691161415610fcd576001600160a01b03881660009081526001602052604081208391898987818110610f0e57610f0e611736565b9050602002810190610f20919061182b565b610f2e9060208101906118e4565b6001600160a01b031681526020810191909152604001600020610f5260018461171f565b81548110610f6257610f62611736565b6000918252602091829020835191018054928401516040909401516001600160401b03908116600160601b0267ffffffffffffffff60601b1991909516600160201b026001600160601b031990941663ffffffff9093169290921792909217169190911790556110a5565b63ffffffff841682526001600160a01b03881660009081526001602052604081209088888681811061100157611001611736565b9050602002810190611013919061182b565b6110219060208101906118e4565b6001600160a01b031681526020808201929092526040908101600090812080546001810182559082529083902085519101805493860151928601516001600160401b03908116600160601b0267ffffffffffffffff60601b1991909416600160201b026001600160601b031990951663ffffffff9093169290921793909317161790555b7f802e045376358152b85ba2107735ff6f465df424b0fcbcb4690c83951d73ebd6888888868181106110d9576110d9611736565b90506020028101906110eb919061182b565b6110f99060208101906118e4565b89898781811061110b5761110b611736565b905060200281019061111d919061182b565b61112e906040810190602001611930565b604080516001600160a01b0394851681529390921660208401526001600160401b03169082015263ffffffff8616606082015260800160405180910390a150508080611179906119ef565b91505061050d565b50949350505050565b60006111946114ec565b9050336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146112425760405162461bcd60e51b815260206004820152604560248201527f4f70657261746f725365744d616e616765722e6c6f636b5374616b655570646160448201527f746573417445706f63683a2043616c6c6572206973206e6f742074686520736c60648201526430b9b432b960d91b608482015260a40161016b565b6001600160a01b0380841660009081526002602090815260408083209386168352929052205480156113c6576001600160a01b03808516600090815260026020908152604080832093871683529290522063ffffffff8316906112a660018461171f565b815481106112b6576112b6611736565b600091825260209091200154600160201b900463ffffffff1614156112db5750505050565b6001600160a01b03808516600090815260026020908152604080832093871683529290522063ffffffff83169061131360018461171f565b8154811061132357611323611736565b60009182526020909120015461134790600160201b900463ffffffff166001611a0a565b63ffffffff1614156113c6576001600160a01b038085166000908152600260209081526040808320938716835292905220829061138560018461171f565b8154811061139557611395611736565b9060005260206000200160000160046101000a81548163ffffffff021916908363ffffffff16021790555050505050565b506001600160a01b039283166000908152600260209081526040808320949095168252928352838120845180860190955263ffffffff92831680865285850190815281546001810183559183529390912093519301805492518216600160201b0267ffffffffffffffff199093169390911692909217179055565b600061144b6114ec565b611456906003611a0a565b905090565b604080516001808252818301909252600091829190816020015b604080518082019091526000808252602082015281526020019060019003908161147557905050905082816000815181106114b2576114b2611736565b6020026020010181905250806040516020016114ce9190611a29565b60405160208183030381529060405280519060200120915050919050565b6000611456426000635fc6304082101561156e5760405162461bcd60e51b815260206004820152603d60248201527f45706f63685574696c732e67657445706f636846726f6d54696d657374616d7060448201527f3a2074696d657374616d70206973206265666f72652067656e65736973000000606482015260840161016b565b62093a80611580635fc630408461171f565b61158a9190611809565b92915050565b6001600160a01b03811681146115a557600080fd5b50565b6000604082840312156115ba57600080fd5b50919050565b60008060008060a085870312156115d657600080fd5b84356115e181611590565b93506115f086602087016115a8565b9250606085013561160081611590565b9150608085013563ffffffff8116811461161957600080fd5b939692955090935050565b6000806000806060858703121561163a57600080fd5b843561164581611590565b935060208501356001600160401b038082111561166157600080fd5b818701915087601f83011261167557600080fd5b81358181111561168457600080fd5b8860208260051b850101111561169957600080fd5b6020830195508094505060408701359150808211156116b757600080fd5b506116c4878288016115a8565b91505092959194509250565b600080604083850312156116e357600080fd5b82356116ee81611590565b915060208301356116fe81611590565b809150509250929050565b634e487b7160e01b600052601160045260246000fd5b60008282101561173157611731611709565b500390565b634e487b7160e01b600052603260045260246000fd5b60008161175b5761175b611709565b506000190190565b80356001600160e01b03198116811461177b57600080fd5b919050565b60006040828403121561179257600080fd5b604051604081018181106001600160401b03821117156117c257634e487b7160e01b600052604160045260246000fd5b60405282356117d081611590565b81526117de60208401611763565b60208201529392505050565b600081600019048311821515161561180457611804611709565b500290565b60008261182657634e487b7160e01b600052601260045260246000fd5b500490565b60008235607e1983360301811261184157600080fd5b9190910192915050565b6000808335601e1984360301811261186257600080fd5b8301803591506001600160401b0382111561187c57600080fd5b6020019150600581901b360382131561189457600080fd5b9250929050565b6000808335601e198436030181126118b257600080fd5b8301803591506001600160401b038211156118cc57600080fd5b6020019150600681901b360382131561189457600080fd5b6000602082840312156118f657600080fd5b813561190181611590565b9392505050565b60006001600160401b038381169083168181101561192857611928611709565b039392505050565b60006020828403121561194257600080fd5b81356001600160401b038116811461190157600080fd5b60006001600160401b0380831681851680830382111561197b5761197b611709565b01949350505050565b6001600160a01b038681168252858116602083015260c082019085356119a981611590565b1660408301526001600160e01b03196119c460208701611763565b1660608301526001600160401b038416608083015263ffffffff831660a08301529695505050505050565b6000600019821415611a0357611a03611709565b5060010190565b600063ffffffff80831681851680830382111561197b5761197b611709565b602080825282518282018190526000919060409081850190868401855b82811015611a7e57815180516001600160a01b031685528601516001600160e01b031916868501529284019290850190600101611a46565b509197965050505050505056fea264697066735822122041982517f779c033cdd098e2b2ee7c4118aedfb87dc1ab1b765532e2696951fc64736f6c634300080c0033",
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
