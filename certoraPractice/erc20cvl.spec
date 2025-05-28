methods {
    // ERC20 standard
    function _.name()                                            external => NONDET; // can we use PER_CALLEE_CONSTANT?
    function _.symbol()                                          external => NONDET; // can we use PER_CALLEE_CONSTANT?
    function _.decimals()                                        external => PER_CALLEE_CONSTANT;
    function _.totalSupply()                                     external => totalSupplyCVL(calledContract) expect uint256;
    function _.balanceOf(address a)                              external => balanceOfCVL(calledContract, a) expect uint256;
    function _.allowance(address a, address b)                   external => allowanceCVL(calledContract, a, b) expect uint256;
    function _.approve(address a, uint256 x)                     external with (env e) => approveCVL(calledContract, e.msg.sender, a, x) expect bool;
    function _.transfer(address a, uint256 x)                    external with (env e) => transferCVL(calledContract, e.msg.sender, a, x) expect bool;
    function _.transferFrom(address a, address b, uint256 x)     external with (env e) => transferFromCVL(calledContract, e.msg.sender, a, b, x) expect bool;
    function _.safeTransferFrom(address a, address b, uint256 x) external with (env e) => safeTransferFromCVL(calledContract, e.msg.sender, a, b, x) expect void;
}


/// CVL simple implementations of IERC20:
/// token => totalSupply
ghost mapping(address => uint256) totalSupplyByToken;
/// token => account => balance
ghost mapping(address => mapping(address => uint256)) balanceByToken;
/// token => owner => spender => allowance
ghost mapping(address => mapping(address => mapping(address => uint256))) allowanceByToken;

// function tokenBalanceOf(address token, address account) returns uint256 {
//     return balanceByToken[token][account];
// }

function totalSupplyCVL(address t) returns uint256 {
    return totalSupplyByToken[t];
}

function balanceOfCVL(address t, address a) returns uint256 {
    return balanceByToken[t][a];
}

function allowanceCVL(address t, address a, address b) returns uint256 {
    return allowanceByToken[t][a][b];
}

function approveCVL(address t, address approver, address spender, uint256 amount) returns bool {
    // should be randomly reverting xxx
    bool nondetSuccess;
    if (!nondetSuccess) return false;

    allowanceByToken[t][approver][spender] = amount;
    return true;
}

function transferFromCVL(address t, address spender, address from, address to, uint256 amount) returns bool {
    // should be randomly reverting xxx
    bool nondetSuccess;
    if (!nondetSuccess) return false;

    if (allowanceByToken[t][from][spender] < amount) return false;
    allowanceByToken[t][from][spender] = assert_uint256(allowanceByToken[t][from][spender] - amount);
    return transferCVL(t, from, to, amount);
}

function safeTransferFromCVL(address t, address spender, address from, address to, uint256 amount) {
    bool nondetSuccess;
    if (!nondetSuccess) require false;
    transferFromCVL(t, spender, from, to, amount);
}

function transferCVL(address t, address from, address to, uint256 amount) returns bool {
    // should be randomly reverting xxx
    bool nondetSuccess;
    if (!nondetSuccess) return false;

    if(balanceByToken[t][from] < amount) return false;
    balanceByToken[t][from] = assert_uint256(balanceByToken[t][from] - amount);
    balanceByToken[t][to] = require_uint256(balanceByToken[t][to] + amount);  // We neglect overflows.
    return true;
}