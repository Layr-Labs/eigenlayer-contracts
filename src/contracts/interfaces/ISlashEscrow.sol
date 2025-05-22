// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/ISlashEscrowFactory.sol";
import "../libraries/OperatorSetLib.sol";
import "../interfaces/IStrategy.sol";

interface ISlashEscrow {
    /// @notice Burns or redistributes the underlying tokens of the strategies.
    /// @param slashEscrowFactory The factory contract that created the slash escrow.
    /// @param slashEscrowImplementation The implementation contract that was used to create the slash escrow.
    /// @param operatorSet The operator set that was used to create the slash escrow.
    /// @param slashId The slash ID that was used to create the slash escrow.
    /// @param recipient The recipient of the underlying tokens.
    /// @param strategy The strategy that was used to create the slash escrow.
    function burnOrRedistributeUnderlyingTokens(
        ISlashEscrowFactory slashEscrowFactory,
        ISlashEscrow slashEscrowImplementation,
        OperatorSet calldata operatorSet,
        uint256 slashId,
        address recipient,
        IStrategy strategy
    ) external;

    /// @notice Verifies the deployment parameters of the slash escrow.
    /// @param slashEscrowFactory The factory contract that created the slash escrow.
    /// @param slashEscrowImplementation The implementation contract that was used to create the slash escrow.
    /// @param operatorSet The operator set that was used to create the slash escrow.
    /// @param slashId The slash ID that was used to create the slash escrow.
    /// @return True if the provided parameters create this contract's address, false otherwise.
    function verifyDeploymentParameters(
        ISlashEscrowFactory slashEscrowFactory,
        ISlashEscrow slashEscrowImplementation,
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) external view returns (bool);
}
