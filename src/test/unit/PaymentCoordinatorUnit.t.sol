// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/core/PaymentCoordinator.sol";
import "../../contracts/interfaces/IPaymentCoordinator.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "../mocks/ERC20Mock.sol";



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
        
        dummyToken = new ERC20Mock()
    }

    function testInitialize() public {
        require(paymentCoordinator.rootPublisher() == address(this), "rootPublisher should be set");
        require(paymentCoordinator.merkleRootActivationDelay() == 7200, "merkleRootActivationDelay should be set");
        require(paymentCoordinator.eigenLayerShareBIPs() == 1000, "eigenLayerShareBIPs should be set");
        require(paymentCoordinator.owner() == address(this), "owner should be set");
    }

    function testMakePayment() external {
        PaymentCoordinator.Payment memory payment;
        payment.token = dummyToken;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 100;
        amounts[1] = 200;
        payment.amounts = amounts;


        paymentCoordinator.makePayment(payment);

        require(paymentCoordinator.cumulativeEigenLayerTokeEarnings(dummyToken) == 30, "cumulativeEigenLayerTokeEarnings should be set");
        require(dummyToken.balance(address(paymentCoordinator)) == 300, "incorrect token balance");
    }



}