// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {CrosschainDeployLib} from "../CrosschainDeployLib.sol";
import {CoreContractsDeployer} from "../CoreContractsDeployer.sol";

/// Purpose: precompute and register the DelegationManager implementation address without deploying its bytecode.
contract DeployImplementations is CoreContractsDeployer {
    using Env for *;

    address internal constant EXPECTED_MAINNET_IMPLEMENTATION = 0x6a8BEd4062C895130E2d09bA442D3eCEAd5Df6c2;

    function _runAsEOA() internal virtual override {
        _validatePinnedCreationCode();
        address implementation = CrosschainDeployLib.computeCrosschainAddress(
            Env.protocolCouncilMultisig(), keccak256(_delegationManagerInitCode()), type(DelegationManager).name
        );
        if (keccak256(bytes(Env.env())) == keccak256("mainnet")) {
            require(implementation == EXPECTED_MAINNET_IMPLEMENTATION, "unexpected mainnet implementation");
        }
        deployImpl({name: type(DelegationManager).name, deployedTo: implementation});
    }

    function testScript() public virtual {
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        runAsEOA();

        assertEq(
            address(Env.impl.delegationManager()).code.length,
            0,
            "implementation must remain undeployed until execution"
        );
    }

    /// @dev Uses the pinned creation code rather than `type(DelegationManager).creationCode` so the
    /// CREATE2 address is identical on every machine, regardless of the build's compiler metadata hash.
    function _delegationManagerInitCode() internal view returns (bytes memory) {
        return abi.encodePacked(
            CrosschainDeployLib.DELEGATION_MANAGER_CREATION_CODE_V1_13_1,
            abi.encode(
                Env.proxy.strategyManager(),
                Env.proxy.eigenPodManager(),
                Env.proxy.allocationManager(),
                Env.impl.pauserRegistry(),
                Env.proxy.permissionController(),
                Env.MIN_WITHDRAWAL_DELAY(),
                Env.deployVersion()
            )
        );
    }

    /// @dev The compiler appends a 53-byte CBOR metadata blob to the end of the bytecode: an ipfs hash
    /// of the build environment plus the solc version, ending with its own length (0x0033 = 51, + 2).
    /// It is never executed and is the only part of the bytecode that varies across build machines.
    uint256 internal constant METADATA_TAIL_LENGTH = 53;

    /// @dev Ensures the pinned creation code matches this build's compiled DelegationManager,
    /// byte-for-byte except the metadata tail. Guards against the pinned bytes going stale
    /// if the contract source changes.
    function _validatePinnedCreationCode() internal pure {
        bytes memory pinned = CrosschainDeployLib.DELEGATION_MANAGER_CREATION_CODE_V1_13_1;
        bytes memory compiled = type(DelegationManager).creationCode;
        uint256 executableLen = compiled.length - METADATA_TAIL_LENGTH;
        require(
            pinned.length == compiled.length
                && _hashPrefix(pinned, executableLen) == _hashPrefix(compiled, executableLen),
            "pinned creation code drift"
        );
    }

    /// @dev Returns keccak256 of the first `len` bytes of `code`.
    function _hashPrefix(
        bytes memory code,
        uint256 len
    ) private pure returns (bytes32 hash) {
        assembly {
            hash := keccak256(add(code, 32), len)
        }
    }
}
