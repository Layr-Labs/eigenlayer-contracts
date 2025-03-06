import "@openzeppelin-upgrades/contracts/utils/ShortStringsUpgradeable.sol";

contract ShortStringsUpgradeableHarness {
    using ShortStringsUpgradeable for ShortString;
    function toStringHarness(ShortString sstr) external pure returns (string memory) {
        // Get the length of the string
        uint256 len = sstr.byteLength();

        // Create a bytes array of the correct length
        bytes memory bytesArray = new bytes(len);

        // Copy each byte directly from the ShortString to bytesArray
        bytes32 raw = ShortString.unwrap(sstr);
        for (uint256 i = 0; i < len; i++) {
            bytesArray[i] = raw[i];
        }
        
        // Convert bytes to string
        return string(bytesArray);
    }
}