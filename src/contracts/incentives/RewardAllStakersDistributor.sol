// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../interfaces/IEigen.sol";
import "./ProgrammaticIncentivesConfig.sol";

/**
 * @notice Version of the IRewardsCoordinator interface with struct names/definitions
 * modified to work smoothly with the LinearVector library.
 */
interface IRewardsCoordinator_VectorModification {
    struct RewardsSubmission {
        VectorEntry[] strategiesAndMultipliers;
        IERC20 token;
        uint256 amount;
        uint32 startTimestamp;
        uint32 duration;
    }

    function createRewardsForAllEarners(
        RewardsSubmission[] calldata rewardsSubmissions
    ) external;
}

interface IIncentivesDistributor {
    function distributeIncentives() external;
    event IncentivesDistributed(uint256 amountEIGEN);
}

/**
 * @notice Programmatically distributes EIGEN incentives via minting new EIGEN tokens and calling
 * the RewardsCoordinator.createRewardsForAllEarners(...) function
 * @dev Reads from a single, fixes streamID and vectorIndex of the ProgrammaticIncentivesConfig and
 * mints tokens via the TokenInflationNexus contract
 */
contract RewardAllStakersDistributor is IIncentivesDistributor {
    using SafeERC20 for IERC20;

    IProgrammaticIncentivesConfig public immutable programmaticIncentivesConfig;
    address public immutable rewardsCoordinator;
    IERC20 public immutable bEIGEN;
    IERC20 public immutable EIGEN;
    uint256 public immutable vectorIndex;

    constructor(
        IProgrammaticIncentivesConfig _programmaticIncentivesConfig,
        address _rewardsCoordinator,
        IERC20 _bEIGEN,
        IERC20 _EIGEN,
        uint256 _vectorIndex
    ) {
        programmaticIncentivesConfig = _programmaticIncentivesConfig;
        rewardsCoordinator = _rewardsCoordinator;
        bEIGEN = _bEIGEN;
        EIGEN = _EIGEN;
        vectorIndex = _vectorIndex;
    }

    function distributeIncentives() external {
        // 0) mint new tokens
        programmaticIncentivesConfig.claimForSubstream({substreamRecipient: address(this)});

        // 1) check how many tokens were minted (also accounts for transfers in)
        uint256 tokenAmount = bEIGEN.balanceOf(address(this));

        // 2) approve the bEIGEN token for transfer so it can be wrapped
        bEIGEN.safeApprove(address(EIGEN), tokenAmount);

        // 3) wrap the bEIGEN token to receive EIGEN
        IEigen(address(EIGEN)).wrap(tokenAmount);

        // 4) Set the proper allowance on the coordinator
        EIGEN.safeApprove(address(rewardsCoordinator), tokenAmount);

        // 5) Call the reward coordinator's ForAll API
        IRewardsCoordinator_VectorModification.RewardsSubmission[] memory rewardsSubmission =
            new IRewardsCoordinator_VectorModification.RewardsSubmission[](1);
        rewardsSubmission[0] = IRewardsCoordinator_VectorModification.RewardsSubmission({
            strategiesAndMultipliers: programmaticIncentivesConfig.vector(vectorIndex),
            token: EIGEN,
            amount: tokenAmount,
            // round up to timescale (i.e. start of *next* 'TIMESCALE' interval, in UTC time)
            startTimestamp: uint32(((block.timestamp / TIMESCALE) + 1) * TIMESCALE),
            duration: uint32(TIMESCALE)
        });
        IRewardsCoordinator_VectorModification(rewardsCoordinator).createRewardsForAllEarners(rewardsSubmission);

        emit IncentivesDistributed(tokenAmount);
    }
}























