// optimizing summaries

methods {
    function Math.mulDiv(uint256 x, uint256 y, uint256 denominator) internal returns (uint256) => cvlMulDiv(x, y, denominator);
    function Math.mulDiv(uint256 x, uint256 y, uint256 denominator, Math.Rounding rounding) internal returns (uint256) => cvlMulDivDirectional(x, y, denominator, rounding);
    function MathUpgradeable.average(uint256 a, uint256 b) internal returns (uint256) => cvlAverage(a, b);
}

function cvlMulDiv(uint256 x, uint256 y, uint256 denominator) returns uint256 {
    require denominator != 0;
    return require_uint256(x*y/denominator);
}

function cvlMulDivUp(uint256 x, uint256 y, uint256 denominator) returns uint256 {
    require denominator != 0;
    return require_uint256((x*y + denominator - 1)/denominator);
}

function cvlMulDivDirectional(uint256 x, uint256 y, uint256 denominator, Math.Rounding rounding) returns uint256 {
    if (rounding == Math.Rounding.Up) {
        return cvlMulDivUp(x, y, denominator);
    } else {
        return cvlMulDiv(x, y, denominator);
    }
}

function cvlAverage(uint256 a, uint256 b) returns uint256 {
    return require_uint256((a+b)/2);
}

