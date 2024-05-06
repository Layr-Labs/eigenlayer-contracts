// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of Eigen token strategy contracts on Holesky
 * forge script script/deploy/holesky/Eigen_Strategy_Deploy.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/Eigen_Strategy_Deploy.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 */
contract Eigen_Strategy_Deploy is ExistingDeploymentParser {

    function run() external virtual {
        _parseDeployedContracts("script/configs/holesky/Holesky_current_deployment.config.json");

        vm.startBroadcast();

        _deployStrategy();

        _setTransferRestrictions();

        vm.stopBroadcast();

        _verifyDeployment();

        emit log_string("====Deployed Contracts====");

        emit log_named_address("EigenStrategy", address(eigenStrategy));
        emit log_named_address("EigenStrategyImpl", address(eigenStrategyImpl));
    }

    function _deployStrategy() internal {
        eigenStrategyImpl = new EigenStrategy(strategyManager);

        eigenStrategy = EigenStrategy(
            address(
                new TransparentUpgradeableProxy(
                    address(eigenStrategyImpl), 
                    address(eigenLayerProxyAdmin), 
                    abi.encodeWithSelector(
                        EigenStrategy.initialize.selector,
                        EIGEN,
                        bEIGEN,
                        eigenLayerPauserReg
                    )
                )
            )
        );
    }

    function _setTransferRestrictions() internal {
        EIGEN.setAllowedFrom(address(eigenStrategy), true);
        EIGEN.setAllowedTo(address(eigenStrategy), true);
        bEIGEN.setAllowedFrom(address(eigenStrategy), true);
        bEIGEN.setAllowedTo(address(eigenStrategy), true);
    }

    function _verifyDeployment() internal {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = eigenStrategy;
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        thirdPartyTransfersForbiddenValues[0] = true;

        vm.prank(executorMultisig);
        strategyManager.addStrategiesToDepositWhitelist(strategies, thirdPartyTransfersForbiddenValues);

        vm.startPrank(msg.sender);
        EIGEN.approve(address(strategyManager), type(uint256).max);
        strategyManager.depositIntoStrategy(eigenStrategy, EIGEN, 1 ether);
    }
}
