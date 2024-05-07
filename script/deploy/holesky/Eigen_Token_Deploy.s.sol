// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../../../src/contracts/token/Eigen.sol";
import "../../../src/contracts/token/BackingEigen.sol";
import "../../../src/test/mocks/EmptyContract.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

/**
 * @notice Script used for the first deployment of Eigen token on Holesky
 * forge script script/deploy/holesky/Eigen_Token_Deploy.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/Eigen_Token_Deploy.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 */
contract Eigen_Token_Deploy is Script, Test {
    address operationsMultisig = 0xfaEF7338b7490b9E272d80A1a39f4657cAf2b97d;
    EmptyContract public emptyContract = EmptyContract(0x9690d52B1Ce155DB2ec5eCbF5a262ccCc7B3A6D2);

    ProxyAdmin public tokenProxyAdmin;
    Eigen public EIGENImpl;
    Eigen public EIGEN;
    BackingEigen public bEIGENImpl;
    BackingEigen public bEIGEN;

    uint256 constant TOTAL_SUPPLY = 1673646668284660000000000000;

    function run() external virtual {
        vm.startBroadcast();

        _deployToken();

        vm.stopBroadcast();

        _verifyDeployment();

        emit log_string("====Deployed Contracts====");

        emit log_named_address("ProxyAdmin", address(tokenProxyAdmin));
        emit log_named_address("EIGEN", address(EIGEN));
        emit log_named_address("bEIGEN", address(bEIGEN));
        emit log_named_address("EIGENImpl", address(EIGENImpl));
        emit log_named_address("bEIGENImpl", address(bEIGENImpl));
    }

    function _deployToken() internal {
        // Deploy ProxyAdmin, later set admins for all proxies to be executorMultisig
        tokenProxyAdmin = new ProxyAdmin();

        EIGEN = Eigen(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(tokenProxyAdmin), ""))
        ); 

        bEIGEN = BackingEigen(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(tokenProxyAdmin), ""))
        );

        // deploy impls
        EIGENImpl = new Eigen(IERC20(address(bEIGEN)));
        bEIGENImpl = new BackingEigen(IERC20(address(EIGEN)));

        address[] memory minters = new address[](1);
        minters[0] = msg.sender;

        uint256[] memory mintingAllowances = new uint256[](1);
        mintingAllowances[0] = TOTAL_SUPPLY;

        uint256[] memory mintAllowedAfters = new uint256[](1);

        // upgrade and initialize proxies
        tokenProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(EIGEN))), 
            address(EIGENImpl), 
            abi.encodeWithSelector(
                Eigen.initialize.selector, 
                msg.sender,
                minters,
                mintingAllowances,
                mintAllowedAfters
            )
        );

        EIGEN.mint();

        tokenProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(bEIGEN))), 
            address(bEIGENImpl),
            abi.encodeWithSelector(
                BackingEigen.initialize.selector,
                msg.sender
            )
        );

        tokenProxyAdmin.transferOwnership(operationsMultisig);
    }

    function _verifyDeployment() internal {
        require(EIGEN.totalSupply() == TOTAL_SUPPLY, "Eigen_Token_Deploy: total supply mismatch");
        require(bEIGEN.totalSupply() == TOTAL_SUPPLY, "Eigen_Token_Deploy: total supply mismatch");

        require(bEIGEN.balanceOf(address(EIGEN)) == TOTAL_SUPPLY, "Eigen_Token_Deploy: bEIGEN balance mismatch");
        require(EIGEN.balanceOf(msg.sender) == TOTAL_SUPPLY, "Eigen_Token_Deploy: EIGEN balance mismatch");

        require(EIGEN.owner() == msg.sender, "Eigen_Token_Deploy: EIGEN owner mismatch");
        require(bEIGEN.owner() == msg.sender, "Eigen_Token_Deploy: bEIGEN owner mismatch");

        require(tokenProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(EIGEN)))) == address(EIGENImpl), "Eigen_Token_Deploy: EIGEN implementation mismatch");
        require(tokenProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(bEIGEN)))) == address(bEIGENImpl), "Eigen_Token_Deploy: bEIGEN implementation mismatch");

        require(tokenProxyAdmin.owner() == operationsMultisig, "Eigen_Token_Deploy: ProxyAdmin owner mismatch");
    }
}
