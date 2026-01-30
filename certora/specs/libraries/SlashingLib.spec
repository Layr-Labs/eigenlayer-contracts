import "../optimizations.spec";

methods {
    function SlashingLib.mulWad(uint256 x, uint256 y) internal returns (uint256) => cvlMulDiv(x, y, WAD());
    function SlashingLib.divWad(uint256 x, uint256 y) internal returns (uint256) => cvlMulDiv(x, WAD(), y);
    function SlashingLib.mulWadRoundUp(uint256 x, uint256 y) internal returns (uint256) => cvlMulDivDirectional(x, y, WAD(), Math.Rounding.Up);
}
