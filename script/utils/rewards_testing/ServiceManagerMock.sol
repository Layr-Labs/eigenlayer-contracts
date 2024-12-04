// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ISignatureUtils} from "src/contracts/interfaces/ISignatureUtils.sol";
import {IAVSDirectory} from "src/contracts/interfaces/IAVSDirectory.sol";
import {IRewardsCoordinator} from "src/contracts/interfaces/IRewardsCoordinator.sol";

/**
 * @title Minimal implementation of a ServiceManager-type contract.
 * This contract can be inherited from or simply used as a point-of-reference.
 * @author Layr Labs, Inc.
 */
contract ServiceManagerMock is Ownable {

    IAVSDirectory internal immutable _avsDirectory;
    IRewardsCoordinator internal immutable _rewardsCoordinator;

    address[] public restakeableStrategies;
    mapping(address => address[]) public operatorRestakedStrategies;

    /// @notice Sets the (immutable) `_registryCoordinator` address
    constructor(
        address initialOwner,
        IAVSDirectory __avsDirectory,
        IRewardsCoordinator ___rewardsCoordinator
    ) {
        _avsDirectory = __avsDirectory;
        _rewardsCoordinator = ___rewardsCoordinator;
        _transferOwnership(initialOwner);
    }

    /**
     * @notice Updates the metadata URI for the AVS
     * @param _metadataURI is the metadata URI for the AVS
     * @dev only callable by the owner
     */
    function updateAVSMetadataURI(string memory _metadataURI) public virtual onlyOwner {
        _avsDirectory.updateAVSMetadataURI(_metadataURI);
    }

    /**
     * @notice Creates a new range payment on behalf of an AVS, to be split amongst the
     * set of stakers delegated to operators who are registered to the `avs`.
     * Note that the owner calling this function must have approved the tokens to be transferred to the ServiceManager
     * and of course has the required balances.
     * @param rewardsSubmissions The range payments being created
     * @dev Expected to be called by the ServiceManager of the AVS on behalf of which the payment is being made
     * @dev The duration of the `rangePayment` cannot exceed `rewardsCoordinator.MAX_PAYMENT_DURATION()`
     * @dev The tokens are sent to the `PaymentCoordinator` contract
     * @dev Strategies must be in ascending order of addresses to check for duplicates
     * @dev This function will revert if the `rangePayment` is malformed,
     * e.g. if the `strategies` and `weights` arrays are of non-equal lengths
     */
    function createAVSRewardsSubmission(
        IRewardsCoordinator.RewardsSubmission[] calldata rewardsSubmissions
    ) public virtual onlyOwner {
        for (uint256 i = 0; i < rewardsSubmissions.length; ++i) {
            // transfer token to ServiceManager and approve PaymentCoordinator to transfer again
            // in createAVSRewardsSubmission() call
            rewardsSubmissions[i].token.transferFrom(msg.sender, address(this), rewardsSubmissions[i].amount);
            uint256 allowance = rewardsSubmissions[i].token.allowance(address(this), address(_rewardsCoordinator));
            rewardsSubmissions[i].token.approve(address(_rewardsCoordinator), rewardsSubmissions[i].amount + allowance);
        }

        _rewardsCoordinator.createAVSRewardsSubmission(rewardsSubmissions);
    }

    function createOperatorDirectedAVSRewardsSubmission(
        IRewardsCoordinator.OperatorDirectedRewardsSubmission[] calldata rewardsSubmissions
    ) public virtual onlyOwner {
        for (uint256 i = 0; i < rewardsSubmissions.length; ++i) {
            uint256 amount = 0;
            for (uint256 j = 0; j < rewardsSubmissions[i].operatorRewards.length; ++j) {
                amount += rewardsSubmissions[i].operatorRewards[j].amount;
            }
            rewardsSubmissions[i].token.transferFrom(msg.sender, address(this), amount);
            uint256 allowance = rewardsSubmissions[i].token.allowance(address(this), address(_rewardsCoordinator));
            rewardsSubmissions[i].token.approve(address(_rewardsCoordinator), amount + allowance);
        }

        _rewardsCoordinator.createOperatorDirectedAVSRewardsSubmission(address(this), rewardsSubmissions);
    }

    /**
     * @notice Forwards a call to EigenLayer's AVSDirectory contract to confirm operator registration with the AVS
     * @param operator The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public virtual {
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    /**
     * @notice Forwards a call to EigenLayer's AVSDirectory contract to confirm operator deregistration from the AVS
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator) public virtual {
        _avsDirectory.deregisterOperatorFromAVS(operator);
    }

    /**
     * @notice Returns the list of strategies that the AVS supports for restaking
     * @dev This function is intended to be called off-chain
     * @dev No guarantee is made on uniqueness of each element in the returned array. 
     *      The off-chain service should do that validation separately
     */
    function getRestakeableStrategies() external view returns (address[] memory) {
        return restakeableStrategies;
    }

    /**
     * @notice Forwards a call to Eigenlayer's RewardsCoordinator contract to set the address of the entity that can call `processClaim` on behalf of this contract.
     * @param claimer The address of the entity that can call `processClaim` on behalf of the earner
     * @dev Only callabe by the owner.
     */
    function setClaimerFor(address claimer) public virtual onlyOwner {
        _rewardsCoordinator.setClaimerFor(claimer);
    }

    /**
     * @notice Returns the list of strategies that the operator has potentially restaked on the AVS
     * @param operator The address of the operator to get restaked strategies for
     * @dev This function is intended to be called off-chain
     * @dev No guarantee is made on whether the operator has shares for a strategy in a quorum or uniqueness 
     *      of each element in the returned array. The off-chain service should do that validation separately
     */
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory) {
        return operatorRestakedStrategies[operator];    
    }

    function setRestakeableStrategies(address[] memory strategies) external {
        delete restakeableStrategies;
        for (uint256 i = 0; i < strategies.length; ++i) {
            restakeableStrategies.push(strategies[i]);
        }
    }

    function setOperatorRestakedStrategies(address operator, address[] memory strategies) external {
        delete operatorRestakedStrategies[operator];
        for (uint256 i = 0; i < strategies.length; ++i) {
            operatorRestakedStrategies[operator].push(strategies[i]);
        }
    }

    /// @notice Returns the EigenLayer AVSDirectory contract.
    function avsDirectory() external view returns (address) {
        return address(_avsDirectory);
    }

    function rewardsCoordinator() external view returns (address) {
        return address(_rewardsCoordinator);
    }
    
    // storage gap for upgradeability
    // slither-disable-next-line shadowing-state
    uint256[48] private __GAP;
}