// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/ClonesUpgradeable.sol";
import "../interfaces/ISlashEscrow.sol";
import "../interfaces/ISlashEscrowFactory.sol";

contract SlashEscrow is ISlashEscrow {
    using OperatorSetLib for *;
    using SafeERC20 for IERC20;

    /// @inheritdoc ISlashEscrow
    function burnOrRedistributeUnderlyingTokens(
        ISlashEscrowFactory slashEscrowFactory,
        ISlashEscrow slashEscrowImplementation,
        OperatorSet calldata operatorSet,
        uint256 slashId,
        address recipient,
        IStrategy strategy
    ) external virtual {
        // Assert that the deployment parameters are valid by validating against the address of this proxy.
        require(
            verifyDeploymentParameters(slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId),
            InvalidDeploymentParameters()
        );

        // Assert that the caller is the slash escrow factory.
        require(msg.sender == address(slashEscrowFactory), OnlySlashEscrowFactory());

        // Burn or redistribute the underlying tokens.
        IERC20 underlyingToken = strategy.underlyingToken();
        underlyingToken.safeTransfer(recipient, underlyingToken.balanceOf(address(this)));
    }

    /// @inheritdoc ISlashEscrow
    function verifyDeploymentParameters(
        ISlashEscrowFactory slashEscrowFactory,
        ISlashEscrow slashEscrowImplementation,
        OperatorSet calldata operatorSet,
        uint256 slashId
    ) public view virtual returns (bool) {
        return ClonesUpgradeable.predictDeterministicAddress(
            address(slashEscrowImplementation),
            keccak256(abi.encodePacked(operatorSet.key(), slashId)),
            address(slashEscrowFactory)
        ) == address(this);
    }
}
