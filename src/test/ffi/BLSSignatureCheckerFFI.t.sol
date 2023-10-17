// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./FFIBase.sol";

contract BLSSignatureCheckerFFITests is FFIBase {
    using BN254 for BN254.G1Point;

    BLSSignatureChecker blsSignatureChecker;
    uint256 setQuorumBitmap;

    function setUp() virtual public {
        _deployMockEigenLayerAndAVS();

        blsSignatureChecker = new BLSSignatureChecker(registryCoordinator);
    }

    function xtestSingleBLSSignatureChecker() public {
        uint64 pseudoRandomNumber = 1;
        uint64 numOperators = 2;
        uint64 numNonSigners = 1;
        uint64 numQuorums = 1;
        uint256 quorumBitmap = 0;
        
        _setNonSignerPrivKeys(numNonSigners, pseudoRandomNumber);

        (
            bytes32 msgHash, 
            bytes memory quorumNumbers, 
            uint32 referenceBlockNumber, 
            BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = _getRandomNonSignerStakeAndSignatures(
            pseudoRandomNumber, 
            numOperators, 
            numNonSigners, 
            numQuorums,
            quorumBitmap
        );

        (
            BLSSignatureChecker.QuorumStakeTotals memory quorumStakeTotals,
            /* bytes32 signatoryRecordHash */
        ) = blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );

        assertTrue(quorumStakeTotals.signedStakeForQuorum[0] > 0);
    }

    function xtestFuzzBLSSignatureChecker(
        uint64 pseudoRandomNumber, 
        uint64 numOperators, 
        uint64 numNonSigners, 
        uint64 numQuorums
    ) public {
        vm.assume(numOperators > 0 && numOperators <= 200);
        vm.assume(numQuorums > 0 && numQuorums <= 192);
        vm.assume(numOperators > numNonSigners);
        vm.assume(numOperators > numQuorums);

        _setNonSignerPrivKeys(numNonSigners, pseudoRandomNumber);

        (
            bytes32 msgHash, 
            bytes memory quorumNumbers, 
            uint32 referenceBlockNumber, 
            BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = _getRandomNonSignerStakeAndSignatures(
            pseudoRandomNumber, 
            numOperators, 
            numNonSigners, 
            numQuorums,
            setQuorumBitmap
        );

        (
            BLSSignatureChecker.QuorumStakeTotals memory quorumStakeTotals,
            /* bytes32 signatoryRecordHash */
        ) = blsSignatureChecker.checkSignatures(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );

        assertTrue(quorumStakeTotals.signedStakeForQuorum[0] > 0);
    }

}