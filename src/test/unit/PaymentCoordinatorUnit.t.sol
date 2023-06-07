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
    uint256 rootPostDelay = 7200;

    event PaymentReceived(address indexed receivedFrom, PaymentCoordinator.Payment payment);
    event NewMerkleRootPosted(PaymentCoordinator.MerkleRootPost merkleRootPost);


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
                                        rootPostDelay
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

    function testMakePayment(uint256[] memory amounts) public {
        cheats.assume(amounts.length > 0);
        uint256 sum = _sum(amounts);
        PaymentCoordinator.Payment memory payment;
        payment.token = dummyToken;
        payment.amounts = amounts;

        uint256 balanceBefore = dummyToken.balanceOf(address(paymentCoordinator));
        uint256 cumumlativeEarningsBefore = paymentCoordinator.cumulativeEigenLayerTokeEarnings(dummyToken);
        cheats.expectEmit(address(paymentCoordinator));
        emit PaymentReceived(address(this), payment);
        paymentCoordinator.makePayment(payment);
        require(paymentCoordinator.cumulativeEigenLayerTokeEarnings(dummyToken) - cumumlativeEarningsBefore == sum * eigenLayerShareBIPs / MAX_BIPS, "cumulativeEigenLayerTokeEarnings should be set");
        require(dummyToken.balanceOf(address(paymentCoordinator)) - balanceBefore == sum, "incorrect token balance");
    }

    function testPostMerkleRootFromUnauthorizedAddress(address unauthorizedAddress) external {
        cheats.assume(unauthorizedAddress != paymentCoordinator.rootPublisher());
        cheats.startPrank(unauthorizedAddress);
        cheats.expectRevert(bytes("PaymentCoordinator: Only rootPublisher"));
        paymentCoordinator.postMerkleRoot(bytes32(0), 0, 0);
        cheats.stopPrank();
    }

    function testPostMerkleRoot(bytes32 root, uint256 height, uint256 calculatedUpToBlockNumber) public {
        PaymentCoordinator.MerkleRootPost memory merkleRootPost;
        merkleRootPost.root = root;
        merkleRootPost.height = height;
        merkleRootPost.calculatedUpToBlockNumber = calculatedUpToBlockNumber;

        uint256 numMerkleRootsBefore = paymentCoordinator.merkleRootsLength();

        cheats.startPrank(paymentCoordinator.rootPublisher());
        paymentCoordinator.postMerkleRoot(root, height, calculatedUpToBlockNumber);
        cheats.stopPrank();

        (bytes32 rootPost, uint256 heightPost, uint256 confirmedAtBlockNumber, uint256 calculatedUpToBlockNumberPost) = paymentCoordinator.merkleRoots(numMerkleRootsBefore);
        require(rootPost == root, "root should be set");
        require(heightPost == height, "height should be set");
        require(confirmedAtBlockNumber == block.number + rootPostDelay, "confirmedAtBlockNumber should be set");
        require(calculatedUpToBlockNumberPost == calculatedUpToBlockNumber, "calculatedUpToBlockNumber should be set");
    }

    function testWithdrawEigenLayerShare(uint256[] memory amounts, address recipient) external {
        cheats.assume(recipient != address(0));
        testMakePayment(amounts);
        uint256 amountToClaim = paymentCoordinator.cumulativeEigenLayerTokeEarnings(dummyToken);
        cheats.startPrank(paymentCoordinator.rootPublisher());
        paymentCoordinator.withdrawEigenlayerShare(dummyToken, recipient);
        cheats.stopPrank();
        require(paymentCoordinator.cumulativeEigenLayerTokeEarnings(dummyToken) == 0, "cumulativeEigenLayerTokeEarnings not updated correctly");
        require(dummyToken.balanceOf(recipient) == amountToClaim, "incorrect token balance");

        emit log_named_uint("balance", dummyToken.balanceOf(recipient));
    }

    function testNullifyMerkleRoot(bytes32 root, uint256 height, uint256 calculatedUpToBlockNumber) external {
        testPostMerkleRoot(root, height, calculatedUpToBlockNumber);
        paymentCoordinator.nullifyMerkleRoot(paymentCoordinator.merkleRootsLength() - 1);
        (bytes32 rootPost, , , ) = paymentCoordinator.merkleRoots(paymentCoordinator.merkleRootsLength() - 1);
        require(rootPost == bytes32(0), "root should be set");

        PaymentCoordinator.MerkleLeaf memory leaf;
        leaf.amounts = new uint256[](1);
        leaf.tokens = new IERC20[](1);

        cheats.roll(block.number + rootPostDelay + 1);
        cheats.expectRevert(bytes("PaymentCoordinator.proveAndClaimEarnings: Merkle root is null"));
        paymentCoordinator.proveAndClaimEarnings(new bytes(0), 0, leaf);
    }

    function testMismatchLengthLeafEntries() external {
        PaymentCoordinator.MerkleLeaf memory leaf;
        leaf.amounts = new uint256[](1);
        leaf.tokens = new IERC20[](2);
        cheats.expectRevert(bytes("PaymentCoordinator.proveAndClaimEarnings: leaf amounts and tokens must be same length"));
        paymentCoordinator.proveAndClaimEarnings(new bytes(0), 0, leaf);
    }

    function testClaimEarningsForRootTooEarly(bytes32 root, uint256 height, uint256 calculatedUpToBlockNumber) external {
        cheats.assume(root != bytes32(0));
        testPostMerkleRoot(root, height, calculatedUpToBlockNumber);
        (bytes32 rootPost, , , ) = paymentCoordinator.merkleRoots(paymentCoordinator.merkleRootsLength() - 1);

        PaymentCoordinator.MerkleLeaf memory leaf;
        leaf.amounts = new uint256[](1);
        leaf.tokens = new IERC20[](1);
        cheats.expectRevert(bytes("PaymentCoordinator.proveAndClaimEarnings: Merkle root not yet confirmed"));
        paymentCoordinator.proveAndClaimEarnings(new bytes(0), 0, leaf);
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