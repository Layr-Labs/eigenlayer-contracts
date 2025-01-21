// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PermissionController

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

// PermissionControllerMetaData contains all meta data concerning the PermissionController contract.
var PermissionControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCall\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointeePermissions\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointees\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminAdded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdminAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotHaveZeroAdmins\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b5061001e61002360201b60201c565b6101b3565b5f60019054906101000a900460ff1615610072576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006990610161565b60405180910390fd5b60ff80165f5f9054906101000a900460ff1660ff16146100df5760ff5f5f6101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516100d6919061019a565b60405180910390a15b565b5f82825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e6974695f8201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b5f61014b6027836100e1565b9150610156826100f1565b604082019050919050565b5f6020820190508181035f8301526101788161013f565b9050919050565b5f60ff82169050919050565b6101948161017f565b82525050565b5f6020820190506101ad5f83018461018b565b92915050565b611a8c806101c05f395ff3fe608060405234801561000f575f5ffd5b50600436106100cd575f3560e01c8063910067451161008a578063ad8aca7711610064578063ad8aca771461021e578063df595cb81461024e578063eb5a4e871461027e578063fddbdefd1461029a576100cd565b806391006745146101a2578063950d806e146101d2578063ad5f2210146101ee576100cd565b806306641201146100d1578063268959e5146100ed5780634f906cf914610109578063628806ef146101255780636bddfa1f14610141578063882a3b3814610171575b5f5ffd5b6100eb60048036038101906100e691906115f5565b6102ca565b005b61010760048036038101906101029190611659565b6104c7565b005b610123600480360381019061011e9190611659565b61062a565b005b61013f600480360381019061013a9190611697565b61074a565b005b61015b60048036038101906101569190611697565b61083e565b6040516101689190611779565b60405180910390f35b61018b60048036038101906101869190611659565b61088e565b604051610199929190611850565b60405180910390f35b6101bc60048036038101906101b79190611659565b610aa0565b6040516101c9919061189f565b60405180910390f35b6101ec60048036038101906101e791906115f5565b610b80565b005b61020860048036038101906102039190611697565b610d7e565b6040516102159190611779565b60405180910390f35b61023860048036038101906102339190611659565b610ec2565b604051610245919061189f565b60405180910390f35b610268600480360381019061026391906115f5565b610f1c565b604051610275919061189f565b60405180910390f35b61029860048036038101906102939190611659565b610fce565b005b6102b460048036038101906102af91906118b8565b61113b565b6040516102c19190611779565b60405180910390f35b836102d58133610aa0565b61030b576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f61035685856111ab565b90506103a981836004015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206111ea90919063ffffffff16565b6103df576040517f262118cd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61043081836004015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206111ff90919063ffffffff16565b5061045686836005015f8481526020019081526020015f2061121490919063ffffffff16565b508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c687876040516104b6929190611926565b60405180910390a350505050505050565b816104d28133610aa0565b610508576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002019050600161055682611241565b1161058d576040517f8674495800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105a0838261121490919063ffffffff16565b6105d6576040517fe2db036000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff167fdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce8460405161061c919061194d565b60405180910390a250505050565b816106358133610aa0565b61066b576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0190506106c0838261121490919063ffffffff16565b6106f6576040517fbed8295f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff167fd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac78460405161073c919061194d565b60405180910390a250505050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905061079f33825f0161121490919063ffffffff16565b6107d5576040517fbed8295f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107eb338260020161125490919063ffffffff16565b508173ffffffffffffffffffffffffffffffffffffffff167fbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff9733604051610832919061194d565b60405180910390a25050565b606061088760015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f01611281565b9050919050565b6060805f60015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f610919826112a0565b90505f8167ffffffffffffffff81111561093657610935611966565b5b6040519080825280602002602001820160405280156109645781602001602082028036833780820191505090505b5090505f8267ffffffffffffffff81111561098257610981611966565b5b6040519080825280602002602001820160405280156109b05781602001602082028036833780820191505090505b5090505f5f90505b83811015610a8e576109db6109d682876112b390919063ffffffff16565b6112c8565b8483815181106109ee576109ed611993565b5b60200260200101848481518110610a0857610a07611993565b5b60200260200101827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050508060010190506109b8565b50818195509550505050509250929050565b5f5f610aea60015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600201611241565b03610b25578173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16149050610b7a565b610b778260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206002016112ee90919063ffffffff16565b90505b92915050565b83610b8b8133610aa0565b610bc1576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f610c0c85856111ab565b9050610c5f81836004015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206111ea90919063ffffffff16565b15610c96576040517fad8efeb700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ce781836004015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2061131b90919063ffffffff16565b50610d0d86836005015f8481526020019081526020015f2061125490919063ffffffff16565b508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc01698787604051610d6d929190611926565b60405180910390a350505050505050565b60605f610dc960015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600201611241565b03610e72575f600167ffffffffffffffff811115610dea57610de9611966565b5b604051908082528060200260200182016040528015610e185781602001602082028036833780820191505090505b50905082815f81518110610e2f57610e2e611993565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080915050610ebd565b610eba60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600201611281565b90505b919050565b5f610f148260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f016112ee90919063ffffffff16565b905092915050565b5f610f278585610aa0565b80610fc45750610fc3610f3a84846111ab565b60015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206111ea90919063ffffffff16565b5b9050949350505050565b81610fd98133610aa0565b61100f576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905061106583826002016112ee90919063ffffffff16565b1561109c576040517f980b072800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110b183825f0161125490919063ffffffff16565b6110e7576040517f3357dbc600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff167fb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c8460405161112d919061194d565b60405180910390a250505050565b60605f61114884846111ab565b90506111a160015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206005015f8381526020019081526020015f20611281565b9150509392505050565b5f5f60608473ffffffffffffffffffffffffffffffffffffffff16901b90505f60408460e01c63ffffffff16901b90508082175f1b9250505092915050565b5f6111f7835f0183611330565b905092915050565b5f61120c835f0183611350565b905092915050565b5f611239835f018373ffffffffffffffffffffffffffffffffffffffff165f1b611350565b905092915050565b5f61124d825f0161144c565b9050919050565b5f611279835f018373ffffffffffffffffffffffffffffffffffffffff165f1b61145b565b905092915050565b60605f61128f835f016114c2565b905060608190508092505050919050565b5f6112ac825f0161144c565b9050919050565b5f6112c0835f018361151b565b905092915050565b5f5f5f6060845f1c901c90505f6040855f1c901c60e01b90508181935093505050915091565b5f611313835f018373ffffffffffffffffffffffffffffffffffffffff165f1b611330565b905092915050565b5f611328835f018361145b565b905092915050565b5f5f836001015f8481526020019081526020015f20541415905092915050565b5f5f836001015f8481526020019081526020015f205490505f8114611441575f60018261137d91906119f6565b90505f6001865f018054905061139391906119f6565b90508181146113f9575f865f0182815481106113b2576113b1611993565b5b905f5260205f200154905080875f0184815481106113d3576113d2611993565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f0180548061140c5761140b611a29565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611446565b5f9150505b92915050565b5f815f01805490509050919050565b5f6114668383611330565b6114b857825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f2081905550600190506114bc565b5f90505b92915050565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561150f57602002820191905f5260205f20905b8154815260200190600101908083116114fb575b50505050509050919050565b5f825f01828154811061153157611530611993565b5b905f5260205f200154905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61156f82611546565b9050919050565b61157f81611565565b8114611589575f5ffd5b50565b5f8135905061159a81611576565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6115d4816115a0565b81146115de575f5ffd5b50565b5f813590506115ef816115cb565b92915050565b5f5f5f5f6080858703121561160d5761160c611542565b5b5f61161a8782880161158c565b945050602061162b8782880161158c565b935050604061163c8782880161158c565b925050606061164d878288016115e1565b91505092959194509250565b5f5f6040838503121561166f5761166e611542565b5b5f61167c8582860161158c565b925050602061168d8582860161158c565b9150509250929050565b5f602082840312156116ac576116ab611542565b5b5f6116b98482850161158c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6116f481611565565b82525050565b5f61170583836116eb565b60208301905092915050565b5f602082019050919050565b5f611727826116c2565b61173181856116cc565b935061173c836116dc565b805f5b8381101561176c57815161175388826116fa565b975061175e83611711565b92505060018101905061173f565b5085935050505092915050565b5f6020820190508181035f830152611791818461171d565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6117cb816115a0565b82525050565b5f6117dc83836117c2565b60208301905092915050565b5f602082019050919050565b5f6117fe82611799565b61180881856117a3565b9350611813836117b3565b805f5b8381101561184357815161182a88826117d1565b9750611835836117e8565b925050600181019050611816565b5085935050505092915050565b5f6040820190508181035f830152611868818561171d565b9050818103602083015261187c81846117f4565b90509392505050565b5f8115159050919050565b61189981611885565b82525050565b5f6020820190506118b25f830184611890565b92915050565b5f5f5f606084860312156118cf576118ce611542565b5b5f6118dc8682870161158c565b93505060206118ed8682870161158c565b92505060406118fe868287016115e1565b9150509250925092565b61191181611565565b82525050565b611920816115a0565b82525050565b5f6040820190506119395f830185611908565b6119466020830184611917565b9392505050565b5f6020820190506119605f830184611908565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611a00826119c0565b9150611a0b836119c0565b9250828203905081811115611a2357611a226119c9565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212204d34614385038f370eaea1f99b8449ed39bdab47f6b52761b1136c9f7bb7b1d664736f6c634300081b0033",
}

// PermissionControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionControllerMetaData.ABI instead.
var PermissionControllerABI = PermissionControllerMetaData.ABI

// PermissionControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PermissionControllerMetaData.Bin instead.
var PermissionControllerBin = PermissionControllerMetaData.Bin

// DeployPermissionController deploys a new Ethereum contract, binding an instance of PermissionController to it.
func DeployPermissionController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PermissionController, error) {
	parsed, err := PermissionControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PermissionControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermissionController{PermissionControllerCaller: PermissionControllerCaller{contract: contract}, PermissionControllerTransactor: PermissionControllerTransactor{contract: contract}, PermissionControllerFilterer: PermissionControllerFilterer{contract: contract}}, nil
}

// PermissionController is an auto generated Go binding around an Ethereum contract.
type PermissionController struct {
	PermissionControllerCaller     // Read-only binding to the contract
	PermissionControllerTransactor // Write-only binding to the contract
	PermissionControllerFilterer   // Log filterer for contract events
}

// PermissionControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionControllerSession struct {
	Contract     *PermissionController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PermissionControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionControllerCallerSession struct {
	Contract *PermissionControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PermissionControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionControllerTransactorSession struct {
	Contract     *PermissionControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PermissionControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionControllerRaw struct {
	Contract *PermissionController // Generic contract binding to access the raw methods on
}

// PermissionControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionControllerCallerRaw struct {
	Contract *PermissionControllerCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionControllerTransactorRaw struct {
	Contract *PermissionControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionController creates a new instance of PermissionController, bound to a specific deployed contract.
func NewPermissionController(address common.Address, backend bind.ContractBackend) (*PermissionController, error) {
	contract, err := bindPermissionController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionController{PermissionControllerCaller: PermissionControllerCaller{contract: contract}, PermissionControllerTransactor: PermissionControllerTransactor{contract: contract}, PermissionControllerFilterer: PermissionControllerFilterer{contract: contract}}, nil
}

// NewPermissionControllerCaller creates a new read-only instance of PermissionController, bound to a specific deployed contract.
func NewPermissionControllerCaller(address common.Address, caller bind.ContractCaller) (*PermissionControllerCaller, error) {
	contract, err := bindPermissionController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerCaller{contract: contract}, nil
}

// NewPermissionControllerTransactor creates a new write-only instance of PermissionController, bound to a specific deployed contract.
func NewPermissionControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionControllerTransactor, error) {
	contract, err := bindPermissionController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerTransactor{contract: contract}, nil
}

// NewPermissionControllerFilterer creates a new log filterer instance of PermissionController, bound to a specific deployed contract.
func NewPermissionControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionControllerFilterer, error) {
	contract, err := bindPermissionController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerFilterer{contract: contract}, nil
}

// bindPermissionController binds a generic wrapper to an already deployed contract.
func bindPermissionController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionController *PermissionControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionController.Contract.PermissionControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionController *PermissionControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionController.Contract.PermissionControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionController *PermissionControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionController.Contract.PermissionControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionController *PermissionControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionController *PermissionControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionController *PermissionControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionController.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_PermissionController *PermissionControllerCaller) CanCall(opts *bind.CallOpts, account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "canCall", account, caller, target, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_PermissionController *PermissionControllerSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _PermissionController.Contract.CanCall(&_PermissionController.CallOpts, account, caller, target, selector)
}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_PermissionController *PermissionControllerCallerSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _PermissionController.Contract.CanCall(&_PermissionController.CallOpts, account, caller, target, selector)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCaller) GetAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetAdmins(&_PermissionController.CallOpts, account)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCallerSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetAdmins(&_PermissionController.CallOpts, account)
}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_PermissionController *PermissionControllerCaller) GetAppointeePermissions(opts *bind.CallOpts, account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getAppointeePermissions", account, appointee)

	if err != nil {
		return *new([]common.Address), *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][4]byte)).(*[][4]byte)

	return out0, out1, err

}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_PermissionController *PermissionControllerSession) GetAppointeePermissions(account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	return _PermissionController.Contract.GetAppointeePermissions(&_PermissionController.CallOpts, account, appointee)
}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_PermissionController *PermissionControllerCallerSession) GetAppointeePermissions(account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	return _PermissionController.Contract.GetAppointeePermissions(&_PermissionController.CallOpts, account, appointee)
}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_PermissionController *PermissionControllerCaller) GetAppointees(opts *bind.CallOpts, account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getAppointees", account, target, selector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_PermissionController *PermissionControllerSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	return _PermissionController.Contract.GetAppointees(&_PermissionController.CallOpts, account, target, selector)
}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_PermissionController *PermissionControllerCallerSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	return _PermissionController.Contract.GetAppointees(&_PermissionController.CallOpts, account, target, selector)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCaller) GetPendingAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "getPendingAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetPendingAdmins(&_PermissionController.CallOpts, account)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionController *PermissionControllerCallerSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionController.Contract.GetPendingAdmins(&_PermissionController.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionController *PermissionControllerCaller) IsAdmin(opts *bind.CallOpts, account common.Address, caller common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "isAdmin", account, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionController *PermissionControllerSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _PermissionController.Contract.IsAdmin(&_PermissionController.CallOpts, account, caller)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionController *PermissionControllerCallerSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _PermissionController.Contract.IsAdmin(&_PermissionController.CallOpts, account, caller)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionController *PermissionControllerCaller) IsPendingAdmin(opts *bind.CallOpts, account common.Address, pendingAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionController.contract.Call(opts, &out, "isPendingAdmin", account, pendingAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionController *PermissionControllerSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _PermissionController.Contract.IsPendingAdmin(&_PermissionController.CallOpts, account, pendingAdmin)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionController *PermissionControllerCallerSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _PermissionController.Contract.IsPendingAdmin(&_PermissionController.CallOpts, account, pendingAdmin)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionController *PermissionControllerTransactor) AcceptAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "acceptAdmin", account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionController *PermissionControllerSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AcceptAdmin(&_PermissionController.TransactOpts, account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionController *PermissionControllerTransactorSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AcceptAdmin(&_PermissionController.TransactOpts, account)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactor) AddPendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "addPendingAdmin", account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AddPendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactorSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.AddPendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "removeAdmin", account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactorSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactor) RemoveAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "removeAppointee", account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactorSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.RemoveAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactor) RemovePendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "removePendingAdmin", account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemovePendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionController *PermissionControllerTransactorSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionController.Contract.RemovePendingAdmin(&_PermissionController.TransactOpts, account, admin)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactor) SetAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.contract.Transact(opts, "setAppointee", account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.SetAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionController *PermissionControllerTransactorSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionController.Contract.SetAppointee(&_PermissionController.TransactOpts, account, appointee, target, selector)
}

// PermissionControllerAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the PermissionController contract.
type PermissionControllerAdminRemovedIterator struct {
	Event *PermissionControllerAdminRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAdminRemoved)
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
		it.Event = new(PermissionControllerAdminRemoved)
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
func (it *PermissionControllerAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAdminRemoved represents a AdminRemoved event raised by the PermissionController contract.
type PermissionControllerAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAdminRemovedIterator{contract: _PermissionController.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAdminRemoved)
				if err := _PermissionController.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParseAdminRemoved(log types.Log) (*PermissionControllerAdminRemoved, error) {
	event := new(PermissionControllerAdminRemoved)
	if err := _PermissionController.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the PermissionController contract.
type PermissionControllerAdminSetIterator struct {
	Event *PermissionControllerAdminSet // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAdminSet)
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
		it.Event = new(PermissionControllerAdminSet)
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
func (it *PermissionControllerAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAdminSet represents a AdminSet event raised by the PermissionController contract.
type PermissionControllerAdminSet struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterAdminSet(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerAdminSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAdminSetIterator{contract: _PermissionController.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *PermissionControllerAdminSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAdminSet)
				if err := _PermissionController.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// ParseAdminSet is a log parse operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParseAdminSet(log types.Log) (*PermissionControllerAdminSet, error) {
	event := new(PermissionControllerAdminSet)
	if err := _PermissionController.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerAppointeeRemovedIterator is returned from FilterAppointeeRemoved and is used to iterate over the raw logs and unpacked data for AppointeeRemoved events raised by the PermissionController contract.
type PermissionControllerAppointeeRemovedIterator struct {
	Event *PermissionControllerAppointeeRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAppointeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAppointeeRemoved)
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
		it.Event = new(PermissionControllerAppointeeRemoved)
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
func (it *PermissionControllerAppointeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAppointeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAppointeeRemoved represents a AppointeeRemoved event raised by the PermissionController contract.
type PermissionControllerAppointeeRemoved struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeRemoved is a free log retrieval operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) FilterAppointeeRemoved(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*PermissionControllerAppointeeRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAppointeeRemovedIterator{contract: _PermissionController.contract, event: "AppointeeRemoved", logs: logs, sub: sub}, nil
}

// WatchAppointeeRemoved is a free log subscription operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) WatchAppointeeRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerAppointeeRemoved, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAppointeeRemoved)
				if err := _PermissionController.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
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

// ParseAppointeeRemoved is a log parse operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) ParseAppointeeRemoved(log types.Log) (*PermissionControllerAppointeeRemoved, error) {
	event := new(PermissionControllerAppointeeRemoved)
	if err := _PermissionController.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerAppointeeSetIterator is returned from FilterAppointeeSet and is used to iterate over the raw logs and unpacked data for AppointeeSet events raised by the PermissionController contract.
type PermissionControllerAppointeeSetIterator struct {
	Event *PermissionControllerAppointeeSet // Event containing the contract specifics and raw log

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
func (it *PermissionControllerAppointeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerAppointeeSet)
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
		it.Event = new(PermissionControllerAppointeeSet)
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
func (it *PermissionControllerAppointeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerAppointeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerAppointeeSet represents a AppointeeSet event raised by the PermissionController contract.
type PermissionControllerAppointeeSet struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeSet is a free log retrieval operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) FilterAppointeeSet(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*PermissionControllerAppointeeSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerAppointeeSetIterator{contract: _PermissionController.contract, event: "AppointeeSet", logs: logs, sub: sub}, nil
}

// WatchAppointeeSet is a free log subscription operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) WatchAppointeeSet(opts *bind.WatchOpts, sink chan<- *PermissionControllerAppointeeSet, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerAppointeeSet)
				if err := _PermissionController.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
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

// ParseAppointeeSet is a log parse operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionController *PermissionControllerFilterer) ParseAppointeeSet(log types.Log) (*PermissionControllerAppointeeSet, error) {
	event := new(PermissionControllerAppointeeSet)
	if err := _PermissionController.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PermissionController contract.
type PermissionControllerInitializedIterator struct {
	Event *PermissionControllerInitialized // Event containing the contract specifics and raw log

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
func (it *PermissionControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerInitialized)
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
		it.Event = new(PermissionControllerInitialized)
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
func (it *PermissionControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerInitialized represents a Initialized event raised by the PermissionController contract.
type PermissionControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionController *PermissionControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PermissionControllerInitializedIterator, error) {

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermissionControllerInitializedIterator{contract: _PermissionController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionController *PermissionControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PermissionControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerInitialized)
				if err := _PermissionController.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PermissionController *PermissionControllerFilterer) ParseInitialized(log types.Log) (*PermissionControllerInitialized, error) {
	event := new(PermissionControllerInitialized)
	if err := _PermissionController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerPendingAdminAddedIterator is returned from FilterPendingAdminAdded and is used to iterate over the raw logs and unpacked data for PendingAdminAdded events raised by the PermissionController contract.
type PermissionControllerPendingAdminAddedIterator struct {
	Event *PermissionControllerPendingAdminAdded // Event containing the contract specifics and raw log

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
func (it *PermissionControllerPendingAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerPendingAdminAdded)
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
		it.Event = new(PermissionControllerPendingAdminAdded)
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
func (it *PermissionControllerPendingAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerPendingAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerPendingAdminAdded represents a PendingAdminAdded event raised by the PermissionController contract.
type PermissionControllerPendingAdminAdded struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminAdded is a free log retrieval operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterPendingAdminAdded(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerPendingAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerPendingAdminAddedIterator{contract: _PermissionController.contract, event: "PendingAdminAdded", logs: logs, sub: sub}, nil
}

// WatchPendingAdminAdded is a free log subscription operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchPendingAdminAdded(opts *bind.WatchOpts, sink chan<- *PermissionControllerPendingAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerPendingAdminAdded)
				if err := _PermissionController.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
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

// ParsePendingAdminAdded is a log parse operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParsePendingAdminAdded(log types.Log) (*PermissionControllerPendingAdminAdded, error) {
	event := new(PermissionControllerPendingAdminAdded)
	if err := _PermissionController.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerPendingAdminRemovedIterator is returned from FilterPendingAdminRemoved and is used to iterate over the raw logs and unpacked data for PendingAdminRemoved events raised by the PermissionController contract.
type PermissionControllerPendingAdminRemovedIterator struct {
	Event *PermissionControllerPendingAdminRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerPendingAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerPendingAdminRemoved)
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
		it.Event = new(PermissionControllerPendingAdminRemoved)
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
func (it *PermissionControllerPendingAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerPendingAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerPendingAdminRemoved represents a PendingAdminRemoved event raised by the PermissionController contract.
type PermissionControllerPendingAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminRemoved is a free log retrieval operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) FilterPendingAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerPendingAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.FilterLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerPendingAdminRemovedIterator{contract: _PermissionController.contract, event: "PendingAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingAdminRemoved is a free log subscription operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) WatchPendingAdminRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerPendingAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionController.contract.WatchLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerPendingAdminRemoved)
				if err := _PermissionController.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
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

// ParsePendingAdminRemoved is a log parse operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionController *PermissionControllerFilterer) ParsePendingAdminRemoved(log types.Log) (*PermissionControllerPendingAdminRemoved, error) {
	event := new(PermissionControllerPendingAdminRemoved)
	if err := _PermissionController.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
