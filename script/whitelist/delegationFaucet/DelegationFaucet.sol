// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IDelegationFaucet.sol";
import "script/whitelist/delegationFaucet/DelegationFaucetStaker.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "../ERC20PresetMinterPauser.sol";

/**
 * @title DelegationFaucet for M2
 * @author Layr Labs, Inc.
 * @notice Faucet contract setup with a dummy ERC20 token and strategy for M2 testnet release.
 * Each operator address gets a staker contract assigned to them deterministically.
 * This contract assumes minting role of the ERC20 token and that the ERC20 token has a whitelisted strategy.
 */
contract DelegationFaucet is IDelegationFaucet, Ownable {
    IStrategyManager immutable strategyManager;
    ERC20PresetMinterPauser immutable stakeToken;
    IStrategy immutable stakeStrategy;
    IDelegationManager immutable delegation;

    uint256 public constant DEFAULT_AMOUNT = 100e18;

    constructor(
        IStrategyManager _strategyManager,
        IDelegationManager _delegation,
        ERC20PresetMinterPauser _token,
        IStrategy _strategy
    ) {
        strategyManager = _strategyManager;
        delegation = _delegation;
        stakeToken = _token;
        stakeStrategy = _strategy;
    }

    /**
     * Deploys a DelegationFaucetStaker contract if not already deployed for operator. DelegationFaucetStaker gets minted _depositAmount or
     * DEFAULT_AMOUNT if _depositAmount is 0. Then DelegationFaucetStaker contract deposits into the strategy and delegates to operator.
     * @param _operator The operator to delegate to
     * @param _approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param _approverSalt A unique single use value tied to an individual signature.
     * @param _depositAmount The amount to deposit into the strategy
     */
    function mintDepositAndDelegate(
        address _operator,
        IDelegationManager.SignatureWithExpiry memory _approverSignatureAndExpiry,
        bytes32 _approverSalt,
        uint256 _depositAmount
    ) public onlyOwner {
        // Operator must be registered
        require(delegation.isOperator(_operator), "DelegationFaucet: Operator not registered");
        address staker = getStaker(_operator);
        // Set deposit amount
        if (_depositAmount == 0) {
            _depositAmount = DEFAULT_AMOUNT;
        }

        // Deploy staker address if not already deployed, staker constructor will approve the StrategyManager to spend the stakeToken
        if (!Address.isContract(staker)) {
            Create2.deploy(
                0,
                bytes32(uint256(uint160(_operator))),
                abi.encodePacked(type(DelegationFaucetStaker).creationCode, abi.encode(strategyManager, stakeToken))
            );
        }

        // mint stakeToken to staker
        stakeToken.mint(staker, _depositAmount);
        // deposit into stakeToken strategy, which will increase delegated shares to operator if already delegated
        _depositIntoStrategy(staker, stakeStrategy, stakeToken, _depositAmount);
        // delegateTo operator if not delegated
        if (!delegation.isDelegated(staker)) {
            delegateTo(_operator, _approverSignatureAndExpiry, _approverSalt);
        }
    }

    /**
     * Calls staker contract to deposit into designated strategy, mints staked token if stakeToken and stakeStrategy
     * are specified.
     * @param _staker address of staker contract for operator
     * @param _strategy StakeToken strategy contract
     * @param _token StakeToken
     * @param _amount amount to get minted and to deposit
     */
    function depositIntoStrategy(
        address _staker,
        IStrategy _strategy,
        IERC20 _token,
        uint256 _amount
    ) public onlyOwner returns (bytes memory) {
        // mint stakeToken to staker
        if (_token == stakeToken && _strategy == stakeStrategy) {
            stakeToken.mint(_staker, _amount);
        }
        return _depositIntoStrategy(_staker, _strategy, _token, _amount);
    }

    /**
     * Call staker to delegate to operator
     * @param _operator operator to get staker address from and delegate to
     * @param _approverSignatureAndExpiry Verifies the operator approves of this delegation
     * @param _approverSalt A unique single use value tied to an individual signature.
     */
    function delegateTo(
        address _operator,
        IDelegationManager.SignatureWithExpiry memory _approverSignatureAndExpiry,
        bytes32 _approverSalt
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IDelegationManager.delegateTo.selector,
            _operator,
            _approverSignatureAndExpiry,
            _approverSalt
        );
        return DelegationFaucetStaker(getStaker(_operator)).callAddress(address(delegation), data);
    }

    /**
     * Call queueWithdrawal through staker contract
     */
    function queueWithdrawal(
        address staker,
         IDelegationManager.QueuedWithdrawalParams[] calldata queuedWithdrawalParams
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IDelegationManager.queueWithdrawals.selector, 
            queuedWithdrawalParams
        );
        return DelegationFaucetStaker(staker).callAddress(address(delegation), data);
    }

    /**
     * Call completeQueuedWithdrawal through staker contract
     */
    function completeQueuedWithdrawal(
        address staker,
        IDelegationManager.Withdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IDelegationManager.completeQueuedWithdrawal.selector,
            queuedWithdrawal,
            tokens,
            middlewareTimesIndex,
            receiveAsTokens
        );
        return DelegationFaucetStaker(staker).callAddress(address(delegation), data);
    }

    /**
     * Transfers tokens from staker contract to designated address
     * @param staker staker contract to transfer from
     * @param token ERC20 token
     * @param to the to address
     * @param amount transfer amount
     */
    function transfer(
        address staker,
        address token,
        address to,
        uint256 amount
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(IERC20.transfer.selector, to, amount);
        return DelegationFaucetStaker(staker).callAddress(token, data);
    }

    function callAddress(address to, bytes memory data) public payable onlyOwner returns (bytes memory) {
        (bool ok, bytes memory res) = payable(to).call{value: msg.value}(data);
        if (!ok) {
            revert(string(res));
        }
        return res;
    }

    /**
     * @notice Returns the deterministic staker contract address for the operator
     * @param _operator The operator to get the staker contract address for
     */
    function getStaker(address _operator) public view returns (address) {
        return
            Create2.computeAddress(
                bytes32(uint256(uint160(_operator))), //salt
                keccak256(
                    abi.encodePacked(type(DelegationFaucetStaker).creationCode, abi.encode(strategyManager, stakeToken))
                )
            );
    }

    /**
     * @notice Internal function to deposit into a strategy, has same function signature as StrategyManager.depositIntoStrategy
     */
    function _depositIntoStrategy(
        address _staker,
        IStrategy _strategy,
        IERC20 _token,
        uint256 _amount
    ) internal returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IStrategyManager.depositIntoStrategy.selector,
            _strategy,
            _token,
            _amount
        );
        return DelegationFaucetStaker(_staker).callAddress(address(strategyManager), data);
    }
}
