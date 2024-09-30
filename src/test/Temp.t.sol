// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../test/EigenLayerTestHelper.t.sol";

contract Temp is EigenLayerTestHelper {

    bytes input = hex'334043960000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000a7bcc4248181e18708ea9358b142ed55b4f2ffa70000000000000000000000005accc90436492f24e6af278569691e2c942a676d000000000000000000000000a7bcc4248181e18708ea9358b142ed55b4f2ffa70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000138dc6400000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000acb55c530acdb2849e6d4f36992cd8c9d50ed8f7000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000618b3d09613a1555e000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000ec53bf9167f50cdeb3ae105f56099aaab9061f830000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001';

    uint256 mainnetForkBlock = 19_280_000;

    function test_Temp() public {
        cheats.createSelectFork(cheats.rpcUrl("mainnet"));

        bool isPending = DelegationManager(0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A).pendingWithdrawals(0x7084a0527e4da801c9659975543094ac9c5f4cf281b16e9578db3d92d09c0d50);
        emit log_named_string("isPending", isPending ? "true" : "false");

        cheats.startPrank(0xa7bCC4248181E18708EA9358B142Ed55B4f2fFa7);

        IERC20[][] memory tokens = new IERC20[][](1);
        tokens[0] = new IERC20[](1);
        tokens[0][0] = IERC20(0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83);

        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = IStrategy(0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7);

        uint[] memory shares = new uint[](1);
        shares[0] = 112460460062991799646;

        IDelegationManager.Withdrawal[] memory withdrawals = new IDelegationManager.Withdrawal[](1);
        withdrawals[0] = IDelegationManager.Withdrawal({
            staker: 0xa7bCC4248181E18708EA9358B142Ed55B4f2fFa7,
            delegatedTo: 0x5ACCC90436492F24E6aF278569691e2c942A676d,
            withdrawer: 0xa7bCC4248181E18708EA9358B142Ed55B4f2fFa7,
            nonce: 1,
            startBlock: 20503652,
            strategies: strats,
            shares: shares
        });

        bool[] memory asTokens = new bool[](1);
        asTokens[0] = true;

        // DelegationManager(0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A).completeQueuedWithdrawals({
        //     withdrawals: withdrawals,
        //     tokens: tokens,
        //     middlewareTimesIndexes: new uint[](1),
        //     receiveAsTokens: asTokens
        // });

        bytes memory data = abi.encodeWithSelector(IDelegationManager.completeQueuedWithdrawals.selector,
            withdrawals,
            tokens,
            new uint[](1),
            asTokens
        );

        emit log_named_bytes("data", data);
        emit log_named_bytes("input", input);

        (bool succ, bytes memory ret) = 0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A.call(input);
        require(succ);
    }
}