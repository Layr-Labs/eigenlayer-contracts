// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// NOTE: Mainnet uses the OpenZeppelin v4.9.0 contracts, but this imports the 4.7.1 version. This will be changed after an upgrade.

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../interfaces/IStrategyManager.sol";
import "../strategies/StrategyBase.sol";
import "../interfaces/IEigen.sol";

/**
 * @title Eigen Strategy implementation of `IStrategy` interface, designed to be inherited from by more complex strategies.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @dev Note that this EigenStrategy contract is designed to be compatible with both bEIGEN and EIGEN tokens. It functions exactly the same
 * as the `StrategyBase` contract if bEIGEN were the underlying token, but also allows for depositing and withdrawing EIGEN tokens. This is
 * achieved by unwrapping EIGEN into bEIGEN upon deposit, and wrapping bEIGEN into EIGEN upon withdrawal. Deposits and withdrawals with bEIGEN
 * does not perform and wrapping or unwrapping.
 * @notice This contract functions similarly to an ERC4626 vault, only without issuing a token.
 * To mitigate against the common "inflation attack" vector, we have chosen to use the 'virtual shares' mitigation route,
 * similar to [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/ERC4626.sol).
 * We acknowledge that this mitigation has the known downside of the virtual shares causing some losses to users, which are pronounced
 * particularly in the case of the share exchange rate changing significantly, either positively or negatively.
 * For a fairly thorough discussion of this issue and our chosen mitigation strategy, we recommend reading through
 * [this thread](https://github.com/OpenZeppelin/openzeppelin-contracts/issues/3706) on the OpenZeppelin repo.
 * We specifically use a share offset of `SHARES_OFFSET` and a balance offset of `BALANCE_OFFSET`.
 */
contract EigenStrategy is StrategyBase {
    using SafeERC20 for IERC20;

    /**
     * @notice EIGEN can be deposited into this strategy, where it is unwrapped into bEIGEN and staked in
     * this strategy contract. EIGEN can also be withdrawn by withdrawing bEIGEN from this strategy, and
     * then wrapping it back into EIGEN.
     */
    IEigen public EIGEN;

    /// @notice Since this contract is designed to be initializable, the constructor simply sets `strategyManager`, the only immutable variable.
    constructor(
        IStrategyManager _strategyManager,
        IPauserRegistry _pauserRegistry,
        string memory _version
    ) StrategyBase(_strategyManager, _pauserRegistry, _version) {}

    function initialize(IEigen _EIGEN, IERC20 _bEIGEN) public virtual initializer {
        EIGEN = _EIGEN;
        _initializeStrategyBase(_bEIGEN);
    }

    /**
     * @notice This function hook is called in EigenStrategy.deposit() and is overridden here to
     * allow for depositing of either EIGEN or bEIGEN tokens. If token is bEIGEN aka the underlyingToken,
     * then the contract functions exactly the same as the StrategyBase contract and the deposit is calculated into shares.
     * If token is EIGEN, then the EIGEN is first 1-1 unwrapped into bEIGEN and the deposit shares are calculated as normal.
     * @param token token to be deposited, can be either EIGEN or bEIGEN. If EIGEN, then is unwrapped into bEIGEN
     * @param amount deposit amount
     */
    function _beforeDeposit(IERC20 token, uint256 amount) internal virtual override {
        require(token == underlyingToken || token == EIGEN, OnlyUnderlyingToken());

        if (token == EIGEN) {
            // unwrap EIGEN into bEIGEN assuming a 1-1 unwrapping amount
            // the strategy will then hold `amount` of bEIGEN
            EIGEN.unwrap(amount);
        }
    }

    /**
     * @notice This function hook is called in EigenStrategy.withdraw() before withdrawn shares are calculated and is
     * overridden here to allow for withdrawing shares either into EIGEN or bEIGEN tokens. If wrapping bEIGEN into EIGEN is needed,
     * it is performed in _afterWithdrawal(). This hook just checks the token parameter is either EIGEN or bEIGEN.
     * @param token token to be withdrawn, can be either EIGEN or bEIGEN. If EIGEN, then bEIGEN is wrapped into EIGEN
     */
    function _beforeWithdrawal(
        address, /*recipient*/
        IERC20 token,
        uint256 /*amountShares*/
    ) internal virtual override {
        require(token == underlyingToken || token == EIGEN, OnlyUnderlyingToken());
    }

    /**
     * @notice This function hook is called in EigenStrategy.withdraw() after withdrawn shares are calculated and is
     * overridden here to allow for withdrawing shares either into EIGEN or bEIGEN tokens. If token is bEIGEN aka the underlyingToken,
     * then the contract functions exactly the same as the StrategyBase contract and transfers out bEIGEN to the recipient.
     * If token is EIGEN, then bEIGEN is first 1-1 wrapped into EIGEN and the strategy transfers out the EIGEN to the recipient.
     * @param recipient recipient of the withdrawal
     * @param token token to be withdrawn, can be either EIGEN or bEIGEN. If EIGEN, then bEIGEN is wrapped into EIGEN
     * @param amountToSend amount of tokens to transfer
     */
    function _afterWithdrawal(address recipient, IERC20 token, uint256 amountToSend) internal virtual override {
        if (token == EIGEN) {
            // wrap bEIGEN into EIGEN assuming a 1-1 wrapping amount
            // the strategy will then hold `amountToSend` of EIGEN
            underlyingToken.approve(address(token), amountToSend);
            EIGEN.wrap(amountToSend);
        }

        // Whether the withdrawal specified EIGEN or bEIGEN, the strategy
        // holds the correct balance and can transfer to the recipient here
        token.safeTransfer(recipient, amountToSend);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
