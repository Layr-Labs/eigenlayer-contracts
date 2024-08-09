import "../contracts/libraries/Merkle.sol";
import "../contracts/interfaces/IStrategy.sol";
import "../contracts/libraries/BytesLib.sol";
import "../contracts/interfaces/IAVSDirectory.sol";
import "../contracts/interfaces/IDelegationManager.sol";
import "../contracts/core/StakeRootCompendium.sol";
import "../contracts/core/AVSDirectory.sol";
import "../test/mocks/StrategyManagerMock.sol";
import "../test/mocks/SlasherMock.sol";
import "../test/mocks/EigenPodManagerMock.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";


contract MerklizeScript is Script, Test {
    string internal constant TEST_MNEMONIC = "hundred february vast fluid produce radar notice ridge armed glare panther balance";

    address avs = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    address delegationManagerAddress = 0x5FbDB2315678afecb367f032d93F642f64180aa3;
    address avsDirectoryAddress = 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512;
    address stakeRootCompendiumAddress = 0x5FbDB2315678afecb367f032d93F642f64180aa3;

    IDelegationManager public delegationManager;
    IAVSDirectory public avsDirectory;
    IStakeRootCompendium public stakeRootCompendium;
    function run() public {

        
        delegationManager = IDelegationManager(delegationManagerAddress);
        
        avsDirectory = new AVSDirectory(delegationManager);
        stakeRootCompendium =  new StakeRootCompendium(delegationManager, avsDirectory);


        uint32[] memory operatorSetIds = new uint32[](2048);
        for (uint256 i = 0; i < 2048; ++i) {
            operatorSetIds[i] = uint32(i);
        }

        vm.startBroadcast(avs);
        avsDirectory.createOperatorSets(operatorSetIds);
        avsDirectory.becomeOperatorSetAVS();
        vm.stopBroadcast();

        registerOperators(uint256(2048), operatorSetIds, avs);

    }

    function registerOperators(uint256 numOperators, uint32[] memory operatorSetIds, address avs) internal {
        for (uint256 i = 0; i < numOperators; ++i) {
            (address operator, uint256 privateKey) = deriveRememberKey(TEST_MNEMONIC, uint32(i));
            // get signature
            ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _getOperatorSignature(
                privateKey,
                operator,
                address(0),
                bytes32(0), // salt
                type(uint256).max //expiry
            );

            // Transfer strategy token to staker
            vm.startBroadcast(operator);
            IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
                __deprecated_earningsReceiver: address(operator),
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            delegationManager.registerAsOperator(
                operatorDetails,
                ""
            );
            vm.stopBroadcast();


            vm.startBroadcast(avs);
            avsDirectory.registerOperatorToAVS(operator, operatorSignature);
            avsDirectory.registerOperatorToOperatorSets(operator, operatorSetIds, operatorSignature);
            vm.stopBroadcast();
        }
    }
     function _getOperatorSignature(
        uint256 _operatorPrivateKey,
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) internal view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        operatorSignature.expiry = expiry;
        operatorSignature.salt = salt;
        {
            bytes32 digestHash = avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(_operatorPrivateKey, digestHash);
            operatorSignature.signature = abi.encodePacked(r, s, v);
        }
        return operatorSignature;
    }

}