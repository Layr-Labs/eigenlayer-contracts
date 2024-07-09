// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Slasher

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

// SlasherMetaData contains all meta data concerning the Slasher contract.
var SlasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_operatorSetManager\",\"type\":\"address\",\"internalType\":\"contractIOperatorSetManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedRate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOperatorSetManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareScalingFactor\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shareScalingFactorAtEpoch\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashingEpochHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashingUpdates\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashingRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"scalingFactor\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOperatorSetManager.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"slashingRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620019873803806200198783398101604081905262000034916200006b565b6001600160a01b0392831660805290821660a0521660c052620000bf565b6001600160a01b03811681146200006857600080fd5b50565b6000806000606084860312156200008157600080fd5b83516200008e8162000052565b6020850151909350620000a18162000052565b6040850151909250620000b48162000052565b809150509250925092565b60805160a05160c051611891620000f66000396000818161030b015261070701526000610332015260006101f301526118916000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80635c975abb116100ad578063c78d4bcd11610071578063c78d4bcd14610306578063df5cf7231461032d578063e49a1e8414610354578063f2fde38b14610367578063fabc1cbc1461037a57600080fd5b80635c975abb146102b6578063715018a6146102c7578063886f1195146102cf5780638da5cb5b146102e2578063b59d8fdb146102f357600080fd5b80634279a7e6116100f45780634279a7e61461022d5780634a1def9a14610255578063595c6a67146102685780635ab112d6146102705780635ac86ab71461028357600080fd5b80630b1b781e1461013157806310d67a2f1461019b578063136439dd146101b0578063334f00d6146101c357806339b70e38146101ee575b600080fd5b61017661013f3660046112e4565b60976020908152600093845260408085208252928452828420905282529020546001600160401b0380821691600160401b90041682565b604080516001600160401b039384168152929091166020830152015b60405180910390f35b6101ae6101a936600461132b565b61038d565b005b6101ae6101be366004611348565b610449565b6101d66101d1366004611361565b610588565b6040516001600160401b039091168152602001610192565b6102157f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610192565b61024061023b36600461139a565b6105d8565b60405163ffffffff9091168152602001610192565b6101ae6102633660046113f1565b61062e565b6101ae610a13565b61024061027e366004611361565b610ada565b6102a66102913660046114fa565b606654600160ff9092169190911b9081161490565b6040519015158152602001610192565b606654604051908152602001610192565b6101ae610b7f565b606554610215906001600160a01b031681565b6033546001600160a01b0316610215565b6101d661030136600461151d565b610b93565b6102157f000000000000000000000000000000000000000000000000000000000000000081565b6102157f000000000000000000000000000000000000000000000000000000000000000081565b6101d66103623660046112e4565b610bfa565b6101ae61037536600461132b565b610c6e565b6101ae610388366004611348565b610ce4565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104049190611576565b6001600160a01b0316336001600160a01b03161461043d5760405162461bcd60e51b815260040161043490611593565b60405180910390fd5b61044681610e40565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610491573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b591906115dd565b6104d15760405162461bcd60e51b8152600401610434906115ff565b6066548181161461054a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610434565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6001600160a01b0380831660009081526098602090815260408083209385168352929052908120546001600160401b0316806105cf57670de0b6b3a76400009150506105d2565b90505b92915050565b6099602052826000526040600020602052816000526040600020818154811061060057600080fd5b906000526020600020906008918282040191900660040292509250509054906101000a900463ffffffff1681565b8063ffffffff16600010801561064b57506127108163ffffffff16105b6106b05760405162461bcd60e51b815260206004820152603060248201527f536c61736865722e5f736c617368526571756573746564426970733a20696e7660448201526f0c2d8d2c840c4d2e0e6a8dea6d8c2e6d60831b6064820152608401610434565b604080518082019091523381526001600160e01b03198416602082015260006106d7610f37565b905060005b8451811015610a0a5760008582815181106106f9576106f9611647565b6020026020010151905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633f76c6c78a8785886040518563ffffffff1660e01b8152600401610757949392919061167f565b602060405180830381865afa158015610774573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079891906116bc565b6107ac9061ffff1663ffffffff88166116f6565b905060006107ba8a84610588565b905060006107c88284610f47565b9050609960008c6001600160a01b03166001600160a01b031681526020019081526020016000206000856001600160a01b03166001600160a01b031681526020019081526020016000208690806001815401808255809150506001900390600052602060002090600891828204019190066004029091909190916101000a81548163ffffffff021916908363ffffffff1602179055506040518060400160405280846001600160401b03168152602001826001600160401b0316815250609760008d6001600160a01b03166001600160a01b031681526020019081526020016000206000866001600160a01b03166001600160a01b0316815260200190815260200160002060008863ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a8154816001600160401b0302191690836001600160401b0316021790555060208201518160000160086101000a8154816001600160401b0302191690836001600160401b0316021790555090505080609860008d6001600160a01b03166001600160a01b031681526020019081526020016000206000866001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b031602179055507f471fe23f2a18902ad4f5859f431c6cc59256d682c861ee3405719f2faa09f937868c868a8c886040516109ed96959493929190611725565b60405180910390a15050505080610a039061177b565b90506106dc565b50505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610a5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7f91906115dd565b610a9b5760405162461bcd60e51b8152600401610434906115ff565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6001600160a01b03808316600090815260996020908152604080832093851683529290529081205480610b115760009150506105d2565b6001600160a01b038085166000908152609960209081526040808320938716835292905220610b41600183611796565b81548110610b5157610b51611647565b90600052602060002090600891828204019190066004029054906101000a900463ffffffff169150506105d2565b610b8761109c565b610b9160006110f6565b565b600080600080610ba4888887611148565b915091508015610bee576001600160a01b038089166000908152609760209081526040808320938b16835292815282822063ffffffff86168352905220546001600160401b031692505b50909695505050505050565b6000670de0b6b3a76400008180610c12878787611148565b915091508015610c63576001600160a01b038781166000908152609760209081526040808320938a16835292815282822063ffffffff8616835290522054600160401b90046001600160401b031692505b509095945050505050565b610c7661109c565b6001600160a01b038116610cdb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610434565b610446816110f6565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5b9190611576565b6001600160a01b0316336001600160a01b031614610d8b5760405162461bcd60e51b815260040161043490611593565b606654198119606654191614610e095760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610434565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161057d565b6001600160a01b038116610ece5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610434565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6000610f4242611213565b905090565b60006001600160401b038216610f955760405162461bcd60e51b815260206004820152601360248201527263616e6e6f7420736c61736820666f7220302560681b6044820152606401610434565b6305f5e1006001600160401b0383161115610ffe5760405162461bcd60e51b815260206004820152602360248201527f63616e6e6f7420736c617368206d6f7265207468616e2031303025206174206f6044820152626e636560e81b6064820152608401610434565b60006001600160401b0383166305f5e100148061105b57506001600160401b03808416908516611042670de0b6b3a76400006bffffffffffffffffffffffff6117ad565b61104e906000196117e2565b61105891906117e2565b10155b1561106e57506001600160401b036105cf565b61107c836305f5e1006117f6565b61108a6305f5e100866116f6565b611094919061181e565b949350505050565b6033546001600160a01b03163314610b915760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610434565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0380841660009081526099602090815260408083209386168352929052908120548190819081905b8015611206576001600160a01b038089166000908152609960209081526040808320938b168352929052206111ad600183611796565b815481106111bd576111bd611647565b6000918252602090912060088204015460079091166004026101000a900463ffffffff9081169350861683116111f65760019150611206565b6111ff81611844565b9050611177565b5090969095509350505050565b6000635fc6304082101561128f5760405162461bcd60e51b815260206004820152603d60248201527f45706f63685574696c732e67657445706f636846726f6d54696d657374616d7060448201527f3a2074696d657374616d70206973206265666f72652067656e657369730000006064820152608401610434565b62093a806112a1635fc6304084611796565b6105d291906117e2565b6001600160a01b038116811461044657600080fd5b80356112cb816112ab565b919050565b803563ffffffff811681146112cb57600080fd5b6000806000606084860312156112f957600080fd5b8335611304816112ab565b92506020840135611314816112ab565b9150611322604085016112d0565b90509250925092565b60006020828403121561133d57600080fd5b81356105cf816112ab565b60006020828403121561135a57600080fd5b5035919050565b6000806040838503121561137457600080fd5b823561137f816112ab565b9150602083013561138f816112ab565b809150509250929050565b6000806000606084860312156113af57600080fd5b83356113ba816112ab565b925060208401356113ca816112ab565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561140757600080fd5b8435611412816112ab565b93506020858101356001600160e01b03198116811461143057600080fd5b935060408601356001600160401b038082111561144c57600080fd5b818801915088601f83011261146057600080fd5b813581811115611472576114726113db565b8060051b604051601f19603f83011681018181108582111715611497576114976113db565b60405291825284820192508381018501918b8311156114b557600080fd5b938501935b828510156114da576114cb856112c0565b845293850193928501926114ba565b8097505050505050506114ef606086016112d0565b905092959194509250565b60006020828403121561150c57600080fd5b813560ff811681146105cf57600080fd5b60008060008084860360a081121561153457600080fd5b853561153f816112ab565b9450602086013561154f816112ab565b93506040603f198201121561156357600080fd5b506040850191506114ef608086016112d0565b60006020828403121561158857600080fd5b81516105cf816112ab565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156115ef57600080fd5b815180151581146105cf57600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b80516001600160a01b031682526020908101516001600160e01b031916910152565b6001600160a01b03858116825260a082019061169e602084018761165d565b80851660608401525063ffffffff8316608083015295945050505050565b6000602082840312156116ce57600080fd5b815161ffff811681146105cf57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160401b038083168185168183048111821515161561171c5761171c6116e0565b02949350505050565b63ffffffff87811682526001600160a01b0387811660208401528616604083015260e0820190611758606084018761165d565b80851660a0840152506001600160401b03831660c0830152979650505050505050565b600060001982141561178f5761178f6116e0565b5060010190565b6000828210156117a8576117a86116e0565b500390565b60008160001904831182151516156117c7576117c76116e0565b500290565b634e487b7160e01b600052601260045260246000fd5b6000826117f1576117f16117cc565b500490565b60006001600160401b0383811690831681811015611816576118166116e0565b039392505050565b60006001600160401b0380841680611838576118386117cc565b92169190910492915050565b600081611853576118536116e0565b50600019019056fea26469706673582212209c03e3539ed77eb2eba2b44f15246d6e76ba113ae6b557086a5b8fa7b919590964736f6c634300080c0033",
}

// SlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use SlasherMetaData.ABI instead.
var SlasherABI = SlasherMetaData.ABI

// SlasherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlasherMetaData.Bin instead.
var SlasherBin = SlasherMetaData.Bin

// DeploySlasher deploys a new Ethereum contract, binding an instance of Slasher to it.
func DeploySlasher(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _delegationManager common.Address, _operatorSetManager common.Address) (common.Address, *types.Transaction, *Slasher, error) {
	parsed, err := SlasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlasherBin), backend, _strategyManager, _delegationManager, _operatorSetManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Slasher{SlasherCaller: SlasherCaller{contract: contract}, SlasherTransactor: SlasherTransactor{contract: contract}, SlasherFilterer: SlasherFilterer{contract: contract}}, nil
}

// Slasher is an auto generated Go binding around an Ethereum contract.
type Slasher struct {
	SlasherCaller     // Read-only binding to the contract
	SlasherTransactor // Write-only binding to the contract
	SlasherFilterer   // Log filterer for contract events
}

// SlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlasherSession struct {
	Contract     *Slasher          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlasherCallerSession struct {
	Contract *SlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlasherTransactorSession struct {
	Contract     *SlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlasherRaw struct {
	Contract *Slasher // Generic contract binding to access the raw methods on
}

// SlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlasherCallerRaw struct {
	Contract *SlasherCaller // Generic read-only contract binding to access the raw methods on
}

// SlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlasherTransactorRaw struct {
	Contract *SlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlasher creates a new instance of Slasher, bound to a specific deployed contract.
func NewSlasher(address common.Address, backend bind.ContractBackend) (*Slasher, error) {
	contract, err := bindSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Slasher{SlasherCaller: SlasherCaller{contract: contract}, SlasherTransactor: SlasherTransactor{contract: contract}, SlasherFilterer: SlasherFilterer{contract: contract}}, nil
}

// NewSlasherCaller creates a new read-only instance of Slasher, bound to a specific deployed contract.
func NewSlasherCaller(address common.Address, caller bind.ContractCaller) (*SlasherCaller, error) {
	contract, err := bindSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherCaller{contract: contract}, nil
}

// NewSlasherTransactor creates a new write-only instance of Slasher, bound to a specific deployed contract.
func NewSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*SlasherTransactor, error) {
	contract, err := bindSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherTransactor{contract: contract}, nil
}

// NewSlasherFilterer creates a new log filterer instance of Slasher, bound to a specific deployed contract.
func NewSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*SlasherFilterer, error) {
	contract, err := bindSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlasherFilterer{contract: contract}, nil
}

// bindSlasher binds a generic wrapper to an already deployed contract.
func bindSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slasher *SlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slasher.Contract.SlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slasher *SlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.Contract.SlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slasher *SlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slasher.Contract.SlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slasher *SlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slasher *SlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slasher *SlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slasher.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Slasher *SlasherCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Slasher *SlasherSession) Delegation() (common.Address, error) {
	return _Slasher.Contract.Delegation(&_Slasher.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_Slasher *SlasherCallerSession) Delegation() (common.Address, error) {
	return _Slasher.Contract.Delegation(&_Slasher.CallOpts)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0xb59d8fdb.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCaller) GetSlashedRate(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "getSlashedRate", operator, strategy, operatorSet, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSlashedRate is a free data retrieval call binding the contract method 0xb59d8fdb.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint64, error) {
	return _Slasher.Contract.GetSlashedRate(&_Slasher.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0xb59d8fdb.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,bytes4) operatorSet, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCallerSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet IOperatorSetManagerOperatorSet, epoch uint32) (uint64, error) {
	return _Slasher.Contract.GetSlashedRate(&_Slasher.CallOpts, operator, strategy, operatorSet, epoch)
}

// LastSlashed is a free data retrieval call binding the contract method 0x5ab112d6.
//
// Solidity: function lastSlashed(address operator, address strategy) view returns(uint32)
func (_Slasher *SlasherCaller) LastSlashed(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint32, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "lastSlashed", operator, strategy)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastSlashed is a free data retrieval call binding the contract method 0x5ab112d6.
//
// Solidity: function lastSlashed(address operator, address strategy) view returns(uint32)
func (_Slasher *SlasherSession) LastSlashed(operator common.Address, strategy common.Address) (uint32, error) {
	return _Slasher.Contract.LastSlashed(&_Slasher.CallOpts, operator, strategy)
}

// LastSlashed is a free data retrieval call binding the contract method 0x5ab112d6.
//
// Solidity: function lastSlashed(address operator, address strategy) view returns(uint32)
func (_Slasher *SlasherCallerSession) LastSlashed(operator common.Address, strategy common.Address) (uint32, error) {
	return _Slasher.Contract.LastSlashed(&_Slasher.CallOpts, operator, strategy)
}

// OperatorSetManager is a free data retrieval call binding the contract method 0xc78d4bcd.
//
// Solidity: function operatorSetManager() view returns(address)
func (_Slasher *SlasherCaller) OperatorSetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "operatorSetManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorSetManager is a free data retrieval call binding the contract method 0xc78d4bcd.
//
// Solidity: function operatorSetManager() view returns(address)
func (_Slasher *SlasherSession) OperatorSetManager() (common.Address, error) {
	return _Slasher.Contract.OperatorSetManager(&_Slasher.CallOpts)
}

// OperatorSetManager is a free data retrieval call binding the contract method 0xc78d4bcd.
//
// Solidity: function operatorSetManager() view returns(address)
func (_Slasher *SlasherCallerSession) OperatorSetManager() (common.Address, error) {
	return _Slasher.Contract.OperatorSetManager(&_Slasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Slasher *SlasherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Slasher *SlasherSession) Owner() (common.Address, error) {
	return _Slasher.Contract.Owner(&_Slasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Slasher *SlasherCallerSession) Owner() (common.Address, error) {
	return _Slasher.Contract.Owner(&_Slasher.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Slasher *SlasherCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Slasher *SlasherSession) Paused(index uint8) (bool, error) {
	return _Slasher.Contract.Paused(&_Slasher.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Slasher *SlasherCallerSession) Paused(index uint8) (bool, error) {
	return _Slasher.Contract.Paused(&_Slasher.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Slasher *SlasherCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Slasher *SlasherSession) Paused0() (*big.Int, error) {
	return _Slasher.Contract.Paused0(&_Slasher.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Slasher *SlasherCallerSession) Paused0() (*big.Int, error) {
	return _Slasher.Contract.Paused0(&_Slasher.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Slasher *SlasherCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Slasher *SlasherSession) PauserRegistry() (common.Address, error) {
	return _Slasher.Contract.PauserRegistry(&_Slasher.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Slasher *SlasherCallerSession) PauserRegistry() (common.Address, error) {
	return _Slasher.Contract.PauserRegistry(&_Slasher.CallOpts)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_Slasher *SlasherCaller) ShareScalingFactor(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "shareScalingFactor", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_Slasher *SlasherSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactor(&_Slasher.CallOpts, operator, strategy)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_Slasher *SlasherCallerSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactor(&_Slasher.CallOpts, operator, strategy)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCaller) ShareScalingFactorAtEpoch(opts *bind.CallOpts, operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "shareScalingFactorAtEpoch", operator, strategy, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactorAtEpoch(&_Slasher.CallOpts, operator, strategy, epoch)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_Slasher *SlasherCallerSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _Slasher.Contract.ShareScalingFactorAtEpoch(&_Slasher.CallOpts, operator, strategy, epoch)
}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_Slasher *SlasherCaller) SlashingEpochHistory(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "slashingEpochHistory", arg0, arg1, arg2)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_Slasher *SlasherSession) SlashingEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _Slasher.Contract.SlashingEpochHistory(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// SlashingEpochHistory is a free data retrieval call binding the contract method 0x4279a7e6.
//
// Solidity: function slashingEpochHistory(address , address , uint256 ) view returns(uint32)
func (_Slasher *SlasherCallerSession) SlashingEpochHistory(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (uint32, error) {
	return _Slasher.Contract.SlashingEpochHistory(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_Slasher *SlasherCaller) SlashingUpdates(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "slashingUpdates", arg0, arg1, arg2)

	outstruct := new(struct {
		SlashingRate  uint64
		ScalingFactor uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlashingRate = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ScalingFactor = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_Slasher *SlasherSession) SlashingUpdates(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _Slasher.Contract.SlashingUpdates(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// SlashingUpdates is a free data retrieval call binding the contract method 0x0b1b781e.
//
// Solidity: function slashingUpdates(address , address , uint32 ) view returns(uint64 slashingRate, uint64 scalingFactor)
func (_Slasher *SlasherCallerSession) SlashingUpdates(arg0 common.Address, arg1 common.Address, arg2 uint32) (struct {
	SlashingRate  uint64
	ScalingFactor uint64
}, error) {
	return _Slasher.Contract.SlashingUpdates(&_Slasher.CallOpts, arg0, arg1, arg2)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Slasher *SlasherCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Slasher *SlasherSession) StrategyManager() (common.Address, error) {
	return _Slasher.Contract.StrategyManager(&_Slasher.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Slasher *SlasherCallerSession) StrategyManager() (common.Address, error) {
	return _Slasher.Contract.StrategyManager(&_Slasher.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Pause(&_Slasher.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Pause(&_Slasher.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Slasher *SlasherTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Slasher *SlasherSession) PauseAll() (*types.Transaction, error) {
	return _Slasher.Contract.PauseAll(&_Slasher.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Slasher *SlasherTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Slasher.Contract.PauseAll(&_Slasher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Slasher *SlasherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Slasher *SlasherSession) RenounceOwnership() (*types.Transaction, error) {
	return _Slasher.Contract.RenounceOwnership(&_Slasher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Slasher *SlasherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Slasher.Contract.RenounceOwnership(&_Slasher.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Slasher *SlasherTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Slasher *SlasherSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.SetPauserRegistry(&_Slasher.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Slasher *SlasherTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.SetPauserRegistry(&_Slasher.TransactOpts, newPauserRegistry)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x4a1def9a.
//
// Solidity: function slashOperator(address operator, bytes4 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_Slasher *SlasherTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId [4]byte, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x4a1def9a.
//
// Solidity: function slashOperator(address operator, bytes4 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_Slasher *SlasherSession) SlashOperator(operator common.Address, operatorSetId [4]byte, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _Slasher.Contract.SlashOperator(&_Slasher.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x4a1def9a.
//
// Solidity: function slashOperator(address operator, bytes4 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_Slasher *SlasherTransactorSession) SlashOperator(operator common.Address, operatorSetId [4]byte, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _Slasher.Contract.SlashOperator(&_Slasher.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Slasher *SlasherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Slasher *SlasherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.TransferOwnership(&_Slasher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Slasher *SlasherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Slasher.Contract.TransferOwnership(&_Slasher.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Unpause(&_Slasher.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Slasher *SlasherTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Slasher.Contract.Unpause(&_Slasher.TransactOpts, newPausedStatus)
}

// SlasherInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Slasher contract.
type SlasherInitializedIterator struct {
	Event *SlasherInitialized // Event containing the contract specifics and raw log

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
func (it *SlasherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherInitialized)
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
		it.Event = new(SlasherInitialized)
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
func (it *SlasherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherInitialized represents a Initialized event raised by the Slasher contract.
type SlasherInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Slasher *SlasherFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlasherInitializedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlasherInitializedIterator{contract: _Slasher.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Slasher *SlasherFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlasherInitialized) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherInitialized)
				if err := _Slasher.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParseInitialized(log types.Log) (*SlasherInitialized, error) {
	event := new(SlasherInitialized)
	if err := _Slasher.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the Slasher contract.
type SlasherOperatorSlashedIterator struct {
	Event *SlasherOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *SlasherOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherOperatorSlashed)
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
		it.Event = new(SlasherOperatorSlashed)
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
func (it *SlasherOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherOperatorSlashed represents a OperatorSlashed event raised by the Slasher contract.
type SlasherOperatorSlashed struct {
	Epoch        uint32
	Operator     common.Address
	Strategy     common.Address
	OperatorSet  IOperatorSetManagerOperatorSet
	BipsToSlash  uint32
	SlashingRate uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x471fe23f2a18902ad4f5859f431c6cc59256d682c861ee3405719f2faa09f937.
//
// Solidity: event OperatorSlashed(uint32 epoch, address operator, address strategy, (address,bytes4) operatorSet, uint32 bipsToSlash, uint64 slashingRate)
func (_Slasher *SlasherFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*SlasherOperatorSlashedIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &SlasherOperatorSlashedIterator{contract: _Slasher.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x471fe23f2a18902ad4f5859f431c6cc59256d682c861ee3405719f2faa09f937.
//
// Solidity: event OperatorSlashed(uint32 epoch, address operator, address strategy, (address,bytes4) operatorSet, uint32 bipsToSlash, uint64 slashingRate)
func (_Slasher *SlasherFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *SlasherOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherOperatorSlashed)
				if err := _Slasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x471fe23f2a18902ad4f5859f431c6cc59256d682c861ee3405719f2faa09f937.
//
// Solidity: event OperatorSlashed(uint32 epoch, address operator, address strategy, (address,bytes4) operatorSet, uint32 bipsToSlash, uint64 slashingRate)
func (_Slasher *SlasherFilterer) ParseOperatorSlashed(log types.Log) (*SlasherOperatorSlashed, error) {
	event := new(SlasherOperatorSlashed)
	if err := _Slasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Slasher contract.
type SlasherOwnershipTransferredIterator struct {
	Event *SlasherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SlasherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherOwnershipTransferred)
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
		it.Event = new(SlasherOwnershipTransferred)
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
func (it *SlasherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherOwnershipTransferred represents a OwnershipTransferred event raised by the Slasher contract.
type SlasherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Slasher *SlasherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SlasherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SlasherOwnershipTransferredIterator{contract: _Slasher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Slasher *SlasherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SlasherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherOwnershipTransferred)
				if err := _Slasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParseOwnershipTransferred(log types.Log) (*SlasherOwnershipTransferred, error) {
	event := new(SlasherOwnershipTransferred)
	if err := _Slasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Slasher contract.
type SlasherPausedIterator struct {
	Event *SlasherPaused // Event containing the contract specifics and raw log

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
func (it *SlasherPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherPaused)
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
		it.Event = new(SlasherPaused)
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
func (it *SlasherPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherPaused represents a Paused event raised by the Slasher contract.
type SlasherPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*SlasherPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlasherPausedIterator{contract: _Slasher.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SlasherPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherPaused)
				if err := _Slasher.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParsePaused(log types.Log) (*SlasherPaused, error) {
	event := new(SlasherPaused)
	if err := _Slasher.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the Slasher contract.
type SlasherPauserRegistrySetIterator struct {
	Event *SlasherPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *SlasherPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherPauserRegistrySet)
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
		it.Event = new(SlasherPauserRegistrySet)
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
func (it *SlasherPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherPauserRegistrySet represents a PauserRegistrySet event raised by the Slasher contract.
type SlasherPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Slasher *SlasherFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*SlasherPauserRegistrySetIterator, error) {

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &SlasherPauserRegistrySetIterator{contract: _Slasher.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Slasher *SlasherFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *SlasherPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherPauserRegistrySet)
				if err := _Slasher.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParsePauserRegistrySet(log types.Log) (*SlasherPauserRegistrySet, error) {
	event := new(SlasherPauserRegistrySet)
	if err := _Slasher.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlasherUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Slasher contract.
type SlasherUnpausedIterator struct {
	Event *SlasherUnpaused // Event containing the contract specifics and raw log

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
func (it *SlasherUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlasherUnpaused)
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
		it.Event = new(SlasherUnpaused)
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
func (it *SlasherUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlasherUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlasherUnpaused represents a Unpaused event raised by the Slasher contract.
type SlasherUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*SlasherUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &SlasherUnpausedIterator{contract: _Slasher.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Slasher *SlasherFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SlasherUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Slasher.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlasherUnpaused)
				if err := _Slasher.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Slasher *SlasherFilterer) ParseUnpaused(log types.Log) (*SlasherUnpaused, error) {
	event := new(SlasherUnpaused)
	if err := _Slasher.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
