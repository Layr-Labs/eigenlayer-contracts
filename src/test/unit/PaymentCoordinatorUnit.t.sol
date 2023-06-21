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

    address rootPublisher = address(25);
    address initialOwner = address(26);

    mapping (address => bool) fuzzedAddressMapping;

    event PaymentReceived(address indexed receivedFrom, PaymentCoordinator.Payment payment);
    event NewMerkleRootPosted(PaymentCoordinator.MerkleRootPost merkleRootPost);

    modifier fuzzedAddress(address addr) virtual {
        cheats.assume(fuzzedAddressMapping[addr] == false);
        _;
    }


    function setUp() public {
        proxyAdmin = new ProxyAdmin();
        paymentCoordinatorImplementation = new PaymentCoordinator();
        paymentCoordinator = PaymentCoordinator(
                                address(new TransparentUpgradeableProxy(
                                    address(paymentCoordinatorImplementation), 
                                    address(proxyAdmin),
                                    abi.encodeWithSelector(
                                        paymentCoordinatorImplementation.initialize.selector,
                                        initialOwner,
                                        rootPublisher,
                                        eigenLayerShareBIPs,
                                        rootPostDelay
                                    )
                                )));
        
        dummyToken = new ERC20Mock();

        fuzzedAddressMapping[address(0)] = true;
        fuzzedAddressMapping[address(proxyAdmin)] = true;
    }

    function testInitialize() public view {
        require(paymentCoordinator.rootPublisher() == rootPublisher, "rootPublisher should be set");
        require(paymentCoordinator.merkleRootActivationDelayBlocks() == 7200, "merkleRootActivationDelay should be set");
        require(paymentCoordinator.eigenLayerShareBIPs() == 1000, "eigenLayerShareBIPs should be set");
        require(paymentCoordinator.owner() == initialOwner, "owner should be set");
    }

    function testMakePayment(uint256[] memory amounts) public {
        cheats.assume(amounts.length > 0);
        uint256 sum = _sum(amounts);
        PaymentCoordinator.Payment memory payment;
        payment.token = dummyToken;
        payment.amounts = amounts;

        uint256 balanceBefore = dummyToken.balanceOf(address(paymentCoordinator));
        uint256 cumumlativeEarningsBefore = paymentCoordinator.cumulativeEigenLayerTokenEarnings(dummyToken);
        cheats.expectEmit(true, true, true, true, address(paymentCoordinator));
        emit PaymentReceived(address(this), payment);
        paymentCoordinator.makePayment(payment);
        require(paymentCoordinator.cumulativeEigenLayerTokenEarnings(dummyToken) - cumumlativeEarningsBefore == sum * eigenLayerShareBIPs / MAX_BIPS, "cumulativeEigenLayerTokeEarnings should be set");
        require(dummyToken.balanceOf(address(paymentCoordinator)) - balanceBefore == sum, "incorrect token balance");
    }

    function testPostMerkleRootFromUnauthorizedAddress(address unauthorizedAddress) external fuzzedAddress(unauthorizedAddress) {
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

        uint256 numMerkleRootsBefore = paymentCoordinator.merkleRootPostsLength();

        cheats.startPrank(paymentCoordinator.rootPublisher());
        paymentCoordinator.postMerkleRoot(root, height, calculatedUpToBlockNumber);
        cheats.stopPrank();

        IPaymentCoordinator.MerkleRootPost memory post  = paymentCoordinator.merkleRootPosts(numMerkleRootsBefore);
        require(post.root == root, "root should be set");
        require(post.height == height, "height should be set");
        require(post.confirmedAtBlockNumber == block.number + rootPostDelay, "confirmedAtBlockNumber should be set");
        require(post.calculatedUpToBlockNumber == calculatedUpToBlockNumber, "calculatedUpToBlockNumber should be set");
    }

    function testWithdrawEigenLayerShare(uint256[] memory amounts, address recipient) external {
        cheats.assume(recipient != address(0));
        cheats.assume(recipient != address(paymentCoordinator));
        uint256 balanceBefore = dummyToken.balanceOf(recipient);
        testMakePayment(amounts);
        uint256 amountToClaim = paymentCoordinator.cumulativeEigenLayerTokenEarnings(dummyToken);
        cheats.startPrank(paymentCoordinator.owner());
        paymentCoordinator.withdrawEigenlayerShare(dummyToken, recipient);
        cheats.stopPrank();
        require(paymentCoordinator.cumulativeEigenLayerTokenEarnings(dummyToken) == 0, "cumulativeEigenLayerTokeEarnings not updated correctly");
        emit log_named_uint("amountToClaim", amountToClaim);
        emit log_named_uint("balanceBefore", balanceBefore);
        emit log_named_uint("balanceAfter", dummyToken.balanceOf(recipient));
        require(dummyToken.balanceOf(recipient) - balanceBefore == amountToClaim, "incorrect token balance");
    }

    function testNullifyMerkleRoot(bytes32 root, uint256 height, uint256 calculatedUpToBlockNumber) external {
        testPostMerkleRoot(root, height, calculatedUpToBlockNumber);
        cheats.startPrank(paymentCoordinator.owner());
        paymentCoordinator.nullifyMerkleRoot(paymentCoordinator.merkleRootPostsLength() - 1);
        cheats.stopPrank();
        IPaymentCoordinator.MerkleRootPost memory post = paymentCoordinator.merkleRootPosts(paymentCoordinator.merkleRootPostsLength() - 1);
        require(post.root == bytes32(0), "root should be set");

        PaymentCoordinator.MerkleLeaf memory leaf;
        leaf.amounts = new uint256[](1);
        leaf.tokens = new IERC20[](1);

        cheats.roll(block.number + rootPostDelay + 1);
        cheats.expectRevert(bytes("PaymentCoordinator.proveAndClaimEarnings: Merkle root is null"));
        paymentCoordinator.proveAndClaimEarnings(new bytes(0), 0, leaf, 0);
    }

    function testCallNullifyMerkleRootFromUnauthorizedAddress(address unauthorizedAddress) external fuzzedAddress(unauthorizedAddress) {
        cheats.assume(unauthorizedAddress != paymentCoordinator.owner());
        cheats.startPrank(unauthorizedAddress);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        paymentCoordinator.nullifyMerkleRoot(0);
        cheats.stopPrank();
    }

    function testMismatchLengthLeafEntries() external {
        PaymentCoordinator.MerkleLeaf memory leaf;
        leaf.amounts = new uint256[](1);
        leaf.tokens = new IERC20[](2);
        cheats.expectRevert(bytes("PaymentCoordinator.proveAndClaimEarnings: leaf amounts and tokens must be same length"));
        paymentCoordinator.proveAndClaimEarnings(new bytes(0), 0, leaf, 0);
    }

    function testClaimEarningsForRootTooEarly(bytes32 root, uint256 height, uint256 calculatedUpToBlockNumber) external {
        cheats.assume(root != bytes32(0));
        testPostMerkleRoot(root, height, calculatedUpToBlockNumber);
        paymentCoordinator.merkleRootPosts(paymentCoordinator.merkleRootPostsLength() - 1);

        PaymentCoordinator.MerkleLeaf memory leaf;
        leaf.amounts = new uint256[](1);
        leaf.tokens = new IERC20[](1);
        cheats.expectRevert(bytes("PaymentCoordinator.proveAndClaimEarnings: Merkle root not yet confirmed"));
        paymentCoordinator.proveAndClaimEarnings(new bytes(0), 0, leaf, 0);
    }

    function _sum(uint256[] memory numbers) internal view returns (uint256) {
        uint256 total = 0;
        
        for (uint i = 0; i < numbers.length; i++) {
            cheats.assume(numbers[i] < type(uint32).max);
            total += numbers[i];
        }
        return total;
    }



}