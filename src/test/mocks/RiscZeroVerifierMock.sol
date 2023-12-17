// Copyright 2023 RISC Zero, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.9;

import "../../contracts/interfaces/IRiscZeroVerifier.sol";


contract RiscZeroVerifierMock is IRiscZeroVerifier {
    /// @notice verify that the given receipt is a valid Groth16 RISC Zero recursion receipt.
    /// @return true if the receipt passes the verification checks.
    function verify(Receipt calldata receipt) external view returns (bool){
        return true;
    }

    /// @notice verifies that the given seal is a valid Groth16 RISC Zero proof of execution over the
    ///     given image ID, post-state digest, and journal. Asserts that the input hash
    //      is all-zeros (i.e. no committed input) and the exit code is (Halted, 0).
    /// @return true if the receipt passes the verification checks.
    function verify(bytes calldata seal, bytes32 imageId, bytes32 postStateDigest, bytes32 journalHash)
        external
        view
        returns (bool){
            return true;
        }
}