// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import "src/contracts/interfaces/IEigenPod.sol";
import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/integration/deprecatedInterfaces/mainnet/IStrategyManager.sol";

import "script/utils/ExistingDeploymentParser.sol";
import "src/test/inspector/PrintUtils.t.sol";

interface ITargetDeployer {
    function whitelistedStrategies() external view returns (StrategyBase[] memory);
}

struct HeldAsset {
    StrategyBase strategy;
    uint shares;
    uint tokens;
}

contract Target is Test, PrintUtils {

    EigenPodManager immutable eigenPodManager;
    StrategyManager immutable strategyManager;
    DelegationManager immutable delegationManager;

    StrategyBase[] whitelistedStrategies;
    bool initialized;

    string public name;

    IEigenPod public pod;

    HeldAsset[] assets;

    IDelegationManager.Withdrawal[] queuedWithdrawals;
    IStrategyManager.DeprecatedStruct_QueuedWithdrawal[] queuedWithdrawals_M1;

    constructor() {
        ExistingDeploymentParser parser = ExistingDeploymentParser(msg.sender);
        delegationManager = parser.delegationManager();
        eigenPodManager = parser.eigenPodManager();
        strategyManager = parser.strategyManager();
    }

    modifier action(string memory _action) {
        _logAction(name, _action);
        _;
        _updateAssets();
        _log("");
    }

    function init(string memory _name, StrategyBase[] memory _whitelistedStrategies) public {
        require(!initialized, "init: twice");
        initialized = true;

        name = _name;

        whitelistedStrategies = _whitelistedStrategies;
        pod = eigenPodManager.ownerToPod(address(this));
        _updateAssets();
    }

    /// @dev Deposit all held assets into EigenLayer
    /// NOTE: native ETH currently not supported
    function depositIntoEigenlayer() public action("depositIntoEigenLayer") {
        for (uint i = 0; i < assets.length; i++) {
            HeldAsset memory asset = assets[i];
            IERC20 underlyingToken = asset.strategy.underlyingToken();

            underlyingToken.approve(address(strategyManager), asset.tokens);
            strategyManager.depositIntoStrategy(asset.strategy, underlyingToken, asset.tokens);
        }
    }

    function queueWithdrawals_M1() public action("queueWithdrawals_M1") {
        for (uint i = 0; i < assets.length; i++) {
            HeldAsset memory asset = assets[i];
            IERC20 underlyingToken = asset.strategy.underlyingToken();

            // Get the index of the strategy in stakerStrategyList
            uint index = type(uint).max;
            uint length = strategyManager.stakerStrategyListLength(address(this));
            for (uint j = 0; j < length; j++) {
                if (
                    address(strategyManager.stakerStrategyList(address(this), j)) == 
                    address(asset.strategy)
                ) {
                    index = j;
                    break;
                }
            }

            if (index == type(uint).max) {
                revert("Asset not found");
            }

            uint[] memory indices = new uint[](1);
            IStrategy[] memory strategies = new IStrategy[](1);
            uint[] memory shares = new uint[](1);

            indices[0] = index;
            strategies[0] = IStrategy(address(asset.strategy));
            shares[0] = asset.shares;

            queuedWithdrawals_M1.push(IStrategyManager.DeprecatedStruct_QueuedWithdrawal({
                strategies: strategies,
                shares: shares,
                staker: address(this),
                withdrawerAndNonce: IStrategyManager.DeprecatedStruct_WithdrawerAndNonce({
                    withdrawer: address(this),
                    nonce: uint96(IStrategyManager_DeprecatedM1(address(strategyManager)).numWithdrawalsQueued(address(this)))
                }),
                withdrawalStartBlock: uint32(block.number),
                delegatedAddress: address(0)
            }));

            IStrategyManager_DeprecatedM1(address(strategyManager)).queueWithdrawal({
                strategyIndexes: indices,
                strategies: strategies,
                shares: shares,
                withdrawer: address(this),
                undelegateIfPossible: false
            });
        }
    }

    function migrateM1Withdrawals() public action("migrateM1Withdrawals") {
        for (uint i = 0; i < queuedWithdrawals_M1.length; i++) {
            IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory qw = queuedWithdrawals_M1[i];

            // Calculate equivalent M2 withdrawal
            queuedWithdrawals.push(IDelegationManager.Withdrawal({
                staker: qw.staker,
                delegatedTo: qw.delegatedAddress,
                withdrawer: qw.withdrawerAndNonce.withdrawer,
                nonce: delegationManager.cumulativeWithdrawalsQueued(qw.staker),
                startBlock: qw.withdrawalStartBlock,
                strategies: qw.strategies,
                shares: qw.shares
            }));

            IStrategyManager.DeprecatedStruct_QueuedWithdrawal[] memory qws
                = new IStrategyManager.DeprecatedStruct_QueuedWithdrawal[](1);  
            qws[0] = qw;

            delegationManager.migrateQueuedWithdrawals(qws);
        }
    }

    function logBasicInfo() public {
        _logHeader(name, address(this));

        address _pod = address(pod);

        _logEth("ETH balance", address(this).balance);

        // Basic info - EigenPod
        _log("Has EigenPod", _pod != address(0));
        if (_pod != address(0)) {
            _log("- pod", _pod);
        }

        // LST strategy shares and underlying tokens
        _logSection("Held Assets");
        if (assets.length == 0) {
            _log("No held assets.");
        } else {
            for (uint i = 0; i < assets.length; i++) {
                HeldAsset memory asset = assets[i];

                _log(_stratName(asset.strategy));
                _log("- shares", asset.shares);
                _log("- tokens", asset.tokens);
            }
        }

        // Queued withdrawals we know to exist (because we created them)

        if (queuedWithdrawals.length == 0 && queuedWithdrawals_M1.length == 0) {
            // _log("No known queued withdrawals");
        } else {
            
            if (queuedWithdrawals_M1.length != 0) {
                _logSection("Queued Withdrawals (M1)");
                for (uint i = 0; i < queuedWithdrawals_M1.length; i++) {
                    IStrategyManager.DeprecatedStruct_QueuedWithdrawal memory qw = queuedWithdrawals_M1[i];
    
                    bytes32 withdrawalRoot = strategyManager.calculateWithdrawalRoot(qw);
                    _log("- withdrawalRoot", withdrawalRoot);
                    _log("- pending", strategyManager.withdrawalRootPending(withdrawalRoot));
                    _log("- strategy", _stratName(qw.strategies[0]));
                    _log("- shares", qw.shares[0]);
                    _log("- staker", qw.staker);
                    _log("- withdrawer", qw.withdrawerAndNonce.withdrawer);
                    _log("- nonce", qw.withdrawerAndNonce.nonce);
                    _log("- startBlock", qw.withdrawalStartBlock);
                }
            } 

            if (queuedWithdrawals.length != 0) {
                _logSection("Queued Withdrawals (M2)");
                for (uint i = 0; i < queuedWithdrawals.length; i++) {
                    IDelegationManager.Withdrawal memory qw = queuedWithdrawals[i];
    
                    bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(qw);
                    _log("- withdrawalRoot", withdrawalRoot);
                    _log("- pending", delegationManager.pendingWithdrawals(withdrawalRoot));
                    _log("- strategy", _stratName(qw.strategies[0]));
                    _log("- shares", qw.shares[0]);
                    _log("- staker", qw.staker);
                    _log("- withdrawer", qw.withdrawer);
                    _log("- nonce", qw.nonce);
                    _log("- startBlock", qw.startBlock);
                }
            }
        }
    }

    function _updateAssets() internal {
        delete assets;

        for (uint i = 0; i < whitelistedStrategies.length; i++) {
            StrategyBase strat = whitelistedStrategies[i];
            IERC20 underlying = strat.underlyingToken();

            uint shares = strategyManager.stakerStrategyShares(address(this), strat);
            uint tokens = underlying.balanceOf(address(this));

            if (shares != 0 || tokens != 0) {
                assets.push(HeldAsset({
                    strategy: strat,
                    shares: shares,
                    tokens: tokens
                }));
            }
        }
    }

    function _stratName(IStrategy strat) internal view returns (string memory) {
        return string.concat(
            "Strategy (",
            IERC20Metadata(address(strat.underlyingToken())).symbol(),
            ")"
        );
    }
}