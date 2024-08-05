import "../contracts/libraries/Merkle.sol";
import "../contracts/interfaces/IStrategy.sol";
import "../contracts/libraries/BytesLib.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";


// contract MerklizeTest is Test {
//     Vm cheats = Vm(HEVM_ADDRESS);

//     MerklizeTarget target;

//     function setUp() public {
//         target = new MerklizeTarget();
//     }

//     function test_merklizeOperatorSet(uint256 seed) public {
//         cheats.pauseGasMetering();
//         MerklizeTarget.OperatorLeaf[] memory operatorLeaves = _createOperatorLeafArray(seed, 200, 10);
//         cheats.resumeGasMetering();
//         uint256 gas = gasleft();
//         target.merklizeOperatorSet(operatorLeaves);
//         emit log_named_uint("Gas used: ", gas - gasleft());
//     }

//     function _createOperatorLeafArray(uint256 seed, uint256 numOperators, uint256 numStrategies) internal pure returns (MerklizeTarget.OperatorLeaf[] memory) {
//         MerklizeTarget.OperatorLeaf[] memory operatorLeaves = new MerklizeTarget.OperatorLeaf[](numOperators);
//         for (uint256 i = 0; i < numOperators; i++) {
//             operatorLeaves[i].operator = address(uint160(uint256(keccak256(abi.encodePacked(seed, i)))));
//             operatorLeaves[i].shares = new uint32[](numStrategies);
//             for (uint256 j = 0; j < numStrategies; j++) {
//                 operatorLeaves[i].shares[j] = uint32(uint256(keccak256(abi.encodePacked(seed, i, j))));
//             }
//         }
//         return operatorLeaves;
//     }
// }

contract MerklizeScript is Script, Test {
    function run() public {
        MerklizeTarget target = MerklizeTarget(0xCA8c8688914e0F7096c920146cd0Ad85cD7Ae8b9);
        
        OperatorLeafLib.OperatorLeaf[] memory operatorLeaves = _createOperatorLeafArray(0, 200, 10);
        vm.startBroadcast();
        // bytes32 root = target.merklizeOperatorSetAndCalcStakes(operatorLeaves); // 1812699
        bytes32 root = target.merklizeOperatorSetAndCalcStakes(operatorLeaves); // 1812699
        // bytes32 root = target.hashOperatorSet2(operatorLeaves); // 779500

        vm.stopBroadcast();
        emit log_named_bytes32("Root: ", root);
    }

    function _createOperatorLeafArray(uint256 seed, uint256 numOperators, uint256 numStrategies) internal pure returns (OperatorLeafLib.OperatorLeaf[] memory) {
        OperatorLeafLib.OperatorLeaf[] memory operatorLeaves = new OperatorLeafLib.OperatorLeaf[](numOperators);
        for (uint256 i = 0; i < numOperators; i++) {
            operatorLeaves[i].operator = address(uint160(uint256(keccak256(abi.encodePacked(seed, i)))));
            operatorLeaves[i].shares = 700_000 ether; ///uint64(uint256(keccak256(abi.encodePacked(seed, i, j))));
        }
        return operatorLeaves;
    }
}


library OperatorLeafLib {
//     using BytesLib for bytes;

    struct OperatorLeaf {
        address operator;
        uint256 shares;
    }
}


contract MerklizeTarget {
    event log(uint256 value);

    mapping(uint256 => uint256) public slots;

    constructor () {
        for (uint256 i = 0; i < 1000; i++) {
            slots[i] = 1;
        }
    }

    function merklizeOperatorSet(OperatorLeafLib.OperatorLeaf[] calldata leaves) public returns (bytes32) {
        bytes32[] memory operatorLeaves = new bytes32[](leaves.length);
        for (uint256 i = 0; i < leaves.length; i++) {
            operatorLeaves[i] = keccak256(abi.encodePacked(leaves[i].operator, leaves[i].shares));
        }
        return Merkle.merkleizeKeccak256(operatorLeaves);
    }

    function hashOperatorSet2(OperatorLeafLib.OperatorLeaf[] calldata leaves) public returns (bytes32) {
        bytes32 root = keccak256(abi.encode(leaves));
        return root;
    }

    function merklizeOperatorSetAndCalcStakes(OperatorLeafLib.OperatorLeaf[] calldata leaves) public returns (bytes32) {
        // bytes32 root = keccak256(abi.encode(leaves));

        for (uint256 i = 0; i < leaves.length; i++) {
            if ((leaves[i].shares + slots[i] + i) % 2 == 0) {
                slots[i] = i;
            }
        }

        return keccak256(abi.encode(leaves));
    }
}