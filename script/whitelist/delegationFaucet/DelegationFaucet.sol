// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IWhitelister.sol";
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
    IDelegationManager delegation;

    uint256 public constant DEFAULT_AMOUNT = 100e18;

    //TODO: Deploy ERC20PresetMinterPauser and a corresponding StrategyBase for it
    //TODO: Transfer ownership of Whitelister to multisig after deployment
    //TODO: Give mint/admin/pauser permssions of whitelistToken to Whitelister and multisig after deployment
    //TODO: Give up mint/admin/pauser permssions of whitelistToken for deployer
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
     * Deploys a Staker contract if not already deployed for operator. Staker gets minted _depositAmount or
     * DEFAULT_AMOUNT if _depositAmount is 0. Then Staker contract deposits into the strategy and delegates to operator.
     * @param _operator The operator to delegate to
     * @param _depositAmount The amount of stakeToken to mint to the staker and deposit into the strategy
     */
    function mintDepositAndDelegate(address _operator, uint256 _depositAmount) public onlyOwner {
        address staker = getStaker(_operator);
        // Set deposit amount
        if (_depositAmount == 0) {
            _depositAmount = DEFAULT_AMOUNT;
        }
        stakeToken.mint(staker, _depositAmount);

        // Deploy staker address if not already deployed, staker constructor will approve the StrategyManager to spend the stakeToken
        if (!Address.isContract(staker)) {
            Create2.deploy(
                0,
                bytes32(uint256(uint160(_operator))),
                abi.encodePacked(type(DelegationFaucetStaker).creationCode, abi.encode(strategyManager, stakeToken))
            );
        }

        // deposit into stakeToken strategy
        depositIntoStrategy(staker, stakeStrategy, stakeToken, _depositAmount);

        // delegate to operator
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegateTo(_operator, signatureWithExpiry, bytes32(0));
    }

    function getStaker(address operator) public view returns (address) {
        return
            Create2.computeAddress(
                bytes32(uint256(uint160(operator))), //salt
                keccak256(
                    abi.encodePacked(type(DelegationFaucetStaker).creationCode, abi.encode(strategyManager, stakeToken))
                )
            );
    }

    function depositIntoStrategy(
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IStrategyManager.depositIntoStrategy.selector,
            strategy,
            token,
            amount
        );

        return Staker(staker).callAddress(address(strategyManager), data);
    }

    function delegateTo(
        address operator,
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 approverSalt
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IDelegationManager.delegateTo.selector,
            operator,
            approverSignatureAndExpiry,
            approverSalt
        );

        return Staker(getStaker(operator)).callAddress(address(delegation), data);
    }

    function queueWithdrawal(
        address staker,
        uint256[] calldata strategyIndexes,
        IStrategy[] calldata strategies,
        uint256[] calldata shares,
        address withdrawer
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IStrategyManager.queueWithdrawal.selector,
            strategyIndexes,
            strategies,
            shares,
            withdrawer
        );
        return Staker(staker).callAddress(address(strategyManager), data);
    }

    function completeQueuedWithdrawal(
        address staker,
        IStrategyManager.QueuedWithdrawal calldata queuedWithdrawal,
        IERC20[] calldata tokens,
        uint256 middlewareTimesIndex,
        bool receiveAsTokens
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(
            IStrategyManager.completeQueuedWithdrawal.selector,
            queuedWithdrawal,
            tokens,
            middlewareTimesIndex,
            receiveAsTokens
        );

        return Staker(staker).callAddress(address(strategyManager), data);
    }

    function transfer(
        address staker,
        address token,
        address to,
        uint256 amount
    ) public onlyOwner returns (bytes memory) {
        bytes memory data = abi.encodeWithSelector(IERC20.transfer.selector, to, amount);

        return Staker(staker).callAddress(token, data);
    }

    function callAddress(address to, bytes memory data) public payable onlyOwner returns (bytes memory) {
        (bool ok, bytes memory res) = payable(to).call{value: msg.value}(data);
        if (!ok) {
            revert(string(res));
        }
        return res;
    }
}
