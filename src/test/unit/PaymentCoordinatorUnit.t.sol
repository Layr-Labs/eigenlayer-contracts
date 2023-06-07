// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/core/PaymentCoordinator.sol";
import "../../contracts/interfaces/IPaymentCoordinator.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";


import "forge-std/Test.sol";


contract PaymentCoordinatorTest is Test {
    PaymentCoordinator public paymentCoordinator;
    PaymentCoordinator public paymentCoordinatorImplementation;
    ProxyAdmin public proxyAdmin;

    function setUp() public {
        proxyAdmin = new ProxyAdmin();
        paymentCoordinatorImplementation = new PaymentCoordinator();
        paymentCoordinator = PaymentCoordinator(
                                address(new TransparentUpgradeableProxy(
                                    address(paymentCoordinatorImplementation), 
                                    address(proxyAdmin),
                                    abi.encodeWithSelector(
                                        paymentCoordinatorImplementation.initialize.selector,
                                        address(this),
                                        address(this),
                                        1000,
                                        7200
                                    )
                                )));
    }

    function testInitialize() public {
        require(paymentCoordinator.rootPublisher() == address(this), "rootPublisher should be set");
        require(paymentCoordinator.merkleRootActivationDelay() == 7200, "merkleRootActivationDelay should be set");
        require(paymentCoordinator.eigenLayerShareBIPs() == 1000, "eigenLayerShareBIPs should be set");
        require(paymentCoordinator.owner() == address(this), "owner should be set");
    }

}