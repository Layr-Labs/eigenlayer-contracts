// Helpers to make Points-to-analysis work
using ShortStringsUpgradeableHarness as ShortStringsUpgradeableHarness;

methods {
    function ShortStringsUpgradeable.toString(ShortStringsUpgradeableHarness.ShortString sstr) internal returns (string memory) => cvlToString(sstr);
    function ShortStringsUpgradeableHarness.toStringHarness(ShortStringsUpgradeableHarness.ShortString sstr) external returns (string memory) envfree;
}

function cvlToString(ShortStringsUpgradeableHarness.ShortString sstr) returns string {
    return ShortStringsUpgradeableHarness.toStringHarness(sstr);
}
