// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./FFIBase.sol";
import "./util/BLSSigCheckerExperimental.sol";

//memory_limit = 1073741824
//gas_limit = "18446744073709551615"
contract OperatorCapAnalysisFFI is FFIBase {
    using BN254 for BN254.G1Point;

    BLSSigCheckerExperimental blsSignatureChecker;

    function xtestLoopedScalarMulComparison() public {
        for(uint64 i = 1; i < 193; i++) {
            _compareScalarMuls(
                1, 
                2, 
                1, 
                i,
                (1 << i) - 1
            );
        }
    }
    
    function testSingleScalarMulComparison() public {
        uint64 pseudoRandomNumber = 1;
        uint64 numOperators = 100;
        uint64 numNonSigners = 99;
        uint64 numQuorums = 2;
        uint256 quorumBitmap = (1 << numQuorums) - 1;

        _compareScalarMuls(
            pseudoRandomNumber, 
            numOperators, 
            numNonSigners, 
            numQuorums,
            quorumBitmap
        );
    }

    function _compareScalarMuls(
        uint64 pseudoRandomNumber,
        uint64 numOperators, 
        uint64 numNonSigners, 
        uint64 numQuorums,
        uint256 setQuorumBitmap
    ) internal returns (uint256) {
        _deployMockEigenLayerAndAVS();
        blsSignatureChecker = new BLSSigCheckerExperimental(registryCoordinator);

        vm.pauseGasMetering();
        _setNonSignerPrivKeys(numNonSigners, pseudoRandomNumber);
        vm.resumeGasMetering();

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
        
        uint256 gasBeforeReg = gasleft();
        blsSignatureChecker.checkSignatures_reg(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
        uint256 gasAfterReg = gasleft();
        uint256 regCost = gasBeforeReg - gasAfterReg;

        uint256 gasBeforeTiny = gasleft();
        blsSignatureChecker.checkSignatures_tiny(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
        uint256 gasAfterTiny = gasleft();
        uint256 tinyCost = gasBeforeTiny - gasAfterTiny;

        emit log_named_uint("Operators", numOperators);
        emit log_named_uint("NonSigners", numNonSigners);
        emit log_named_uint("Quorums", numQuorums);
        emit log_named_uint("scalar_mul", regCost);
        emit log_named_uint("scalar_mul_tiny", tinyCost);
        
        if(tinyCost < regCost){
            emit log_named_uint("-", regCost - tinyCost);
            return(tinyCost);
        } else {
            emit log_named_uint("+", tinyCost - regCost);
            return(regCost);
        }
    }

}