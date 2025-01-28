methods
{
    function fromLittleEndianUint64(bytes32) external returns uint64 envfree;
    function toLittleEndianUint64(uint64 num) external returns bytes32 envfree;
    function getByteAt(bytes32, uint) external returns bytes1 envfree;
}

/// @title toLittleEndianUint64 is an inverse of fromLittleEndianUint64
// toLittleEndianUint64(fromLittleEndianUint64(x)) == x
rule transformationsAreInverse1()
{
    bytes32 input_LE;
    require input_LE << 64 == to_bytes32(0);    //the other parts are ignored
    uint64 res = fromLittleEndianUint64(input_LE);
    bytes32 inverse = toLittleEndianUint64(res);

    assert input_LE == inverse;
}

/// @title fromLittleEndianUint64 is an inverse of toLittleEndianUint64
// fromLittleEndianUint64(toLittleEndianUint64(x)) == x
rule transformationsAreInverse2()
{
    uint64 x;
    bytes32 inverse = toLittleEndianUint64(x);
    uint64 res = fromLittleEndianUint64(inverse);
    assert x == res;
}

/// @title fromLittleEndianUint64 returns the expected result
rule fromLittleEndianUint64_correctness()
{
    bytes32 input_LE;
    uint64 res = fromLittleEndianUint64(input_LE);
    bytes32 resAsBytes = to_bytes32(assert_uint256(res));
    uint i;
    require i < 8;
    bytes1 inputByte = getByteAt(input_LE, i);
    bytes1 outputByte = getByteAt(resAsBytes, assert_uint256(31 - i));
    
    assert inputByte == outputByte;
    satisfy inputByte == outputByte;
}


///////////////////   IN DEVELOPMENT / OBSOLETE    ////////


/*
// this currently times out but this is implied by fromLittleEndianUint64_correctness 
// this doesn't compile unless you set "disable_local_typechecking": true
rule fromLittleEndianUint64_isSurjective()
{
    assert forall uint64 res . exists bytes32 input . fromLittleEndianUint64(input) == res;
    //assert forall uint64 res . exists bytes32 input . fromLittleEndianUint64_CVL(input) == res;
}
*/

//just a copy of the contract method to CVL
function fromLittleEndianUint64_CVL(bytes32 lenum) returns uint64 
{
    uint256 inputAsUint;
    require to_bytes32(inputAsUint) == lenum;
    uint64 n = require_uint64(require_uint256(inputAsUint >> 192));
    return (n >> 56) | 
        ((0x00FF000000000000 & n) >> 40) | 
        ((0x0000FF0000000000 & n) >> 24) | 
        ((0x000000FF00000000 & n) >> 8)  | 
        ((0x00000000FF000000 & n) << 8)  | 
        ((0x0000000000FF0000 & n) << 24) | 
        ((0x000000000000FF00 & n) << 40) | 
        ((0x00000000000000FF & n) << 56);
}