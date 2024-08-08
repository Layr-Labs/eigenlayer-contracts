import "../contracts/libraries/Merkle.sol";
import "../contracts/interfaces/IStrategy.sol";
import "../contracts/libraries/BytesLib.sol";
import "../contracts/interfaces/IAVSDirectory.sol";
import "../contracts/interfaces/IDelegationManager.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";


contract MerklizeScript is Script, Test {
    function run() public {
        
        IDelegationManager delegation = IDelegationManager("0x");
        IAVSDirectory avsDirectory = IAVSDirectory("0x");

        register2048OperatorsEl(delegation);
        register2048OperatorsAVS(avsDirectory);



    }

    function register2048OperatorsEl(IDelegationManager delegation) public {
        // for (uint i = 0; i < 2048; i++) {
        //     delegation.setIsOperator(address(i), true);
        // }
    }

    function register2048OperatorsAVS(IAVSDirectory avsDirectory) public {
        uint32[] memory operatorSetIds = new uint32[](2048);
        for (uint32 i = 0; i < 2048; i++) {
            operatorSetIds[i] = i;
        }
        avsDirectory.createOperatorSets(operatorSetIds);

        
    }
}