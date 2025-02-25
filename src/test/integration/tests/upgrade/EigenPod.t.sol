// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_EigenPod_Base is UpgradeTest {
    User staker;
    IStrategy[] strategies;
    uint256[] tokenBalances;
    uint256[] shares;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_ETH);
        _configUserTypes(DEFAULT);

        /// 0. Create a staker with underlying assets
        (staker, strategies, tokenBalances) = _newRandomStaker();
        shares = _calculateExpectedShares(strategies, tokenBalances);

        ///  1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
    }
}

contract Integration_Upgrade_EigenPod_FullSlash is Integration_Upgrade_EigenPod_Base {
    uint40[] validators;
    uint64 slashedGwei;

    function _init() internal override {
        super._init();

        /// 2. Fully slash the staker
        validators = staker.getActiveValidators();
        slashedGwei = beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Full);
        beaconChain.advanceEpoch_NoRewards(); // Withdraw slashed validators to pod

        /// 3. Upgrade contracts
        _upgradeEigenLayerContracts();
    }

    function testFuzz_deposit_fullSlash_upgrade_delegate(uint24 _rand) public rand(_rand) {
        /// 4. Delegate to operator
        (User operator,,) = _newRandomOperator();
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);
    }

    function testFuzz_deposit_fullSlash_upgrade_deposit_delegate(uint24 _rand) public rand(_rand) {
        // 5. Start a new validator & verify withdrawal credentials
        cheats.deal(address(staker), 32 ether);
        tokenBalances[0] = tokenBalances[0] + 32 ether;
        (uint40[] memory newValidators, uint64 addedBeaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(newValidators);
        check_VerifyWC_State(staker, newValidators, addedBeaconBalanceGwei);
        shares = _calculateExpectedShares(strategies, tokenBalances);

        // 6. Delegate to operator
        (User operator,,) = _newRandomOperator();
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);
    }
}
