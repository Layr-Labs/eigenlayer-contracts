// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/core/PaymentCoordinator.sol";
import "../../contracts/interfaces/IPaymentCoordinator.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "../mocks/ERC20Mock.sol";



import "forge-std/Test.sol";


contract PaymentCoordinatorTest is Test {
    Vm cheats = Vm(HEVM_ADDRESS);
    PaymentCoordinator public paymentCoordinator;
    PaymentCoordinator public paymentCoordinatorImplementation;
    ProxyAdmin public proxyAdmin;
    ERC20Mock public dummyToken;
    uint256 MAX_BIPS = 10000;
    uint256 eigenLayerShareBIPs = 1000;

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
                                        eigenLayerShareBIPs,
                                        7200
                                    )
                                )));
        
        dummyToken = new ERC20Mock();
    }

    function testInitialize() public {
        require(paymentCoordinator.rootPublisher() == address(this), "rootPublisher should be set");
        require(paymentCoordinator.merkleRootActivationDelay() == 7200, "merkleRootActivationDelay should be set");
        require(paymentCoordinator.eigenLayerShareBIPs() == 1000, "eigenLayerShareBIPs should be set");
        require(paymentCoordinator.owner() == address(this), "owner should be set");
    }

    function testMakePayment(uint256[] memory amounts) external {
        cheats.assume(amounts.length > 0);
        uint256 sum = _sum(amounts);
        PaymentCoordinator.Payment memory payment;
        payment.token = dummyToken;
        payment.amounts = amounts;


        paymentCoordinator.makePayment(payment);
        require(paymentCoordinator.cumulativeEigenLayerTokeEarnings(dummyToken) == sum * eigenLayerShareBIPs / MAX_BIPS, "cumulativeEigenLayerTokeEarnings should be set");
        require(dummyToken.balanceOf(address(paymentCoordinator)) == sum, "incorrect token balance");
    }

    function _sum(uint[] memory numbers) internal returns (uint) {
        uint total = 0;
        
        for (uint i = 0; i < numbers.length; i++) {
            cheats.assume(numbers[i] < type(uint32).max && numbers[i] > 0);
            total += numbers[i];
        }
        
        return total;
    }



}