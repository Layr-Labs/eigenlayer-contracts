methods {
    // ERC20 standard
    function _.name()                                           external => NONDET; // can we use PER_CALLEE_CONSTANT?
    function _.symbol()                                         external => NONDET; // can we use PER_CALLEE_CONSTANT?
    function _.decimals()                                       external => PER_CALLEE_CONSTANT;
    function _.totalSupply()                                    external => totalSupplyCVL(calledContract) expect uint256;
    function _.balanceOf(address a)                             external => balanceOfCVL(calledContract, a) expect uint256;
    function _.allowance(address a, address b)                  external => allowanceCVL(calledContract, a, b) expect uint256;
    function _.approve(address a, uint256 x)                    external with (env e) => approveCVL(calledContract, e.msg.sender, a, x) expect bool;
    function _.transfer(address a, uint256 x)                   external with (env e) => transferCVL(calledContract, e.msg.sender, a, x) expect bool;
    function _.transferFrom(address a, address b, uint256 x)    external with (env e) => transferFromCVL(calledContract, e.msg.sender, a, b, x) expect bool;

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

function totalSupplyCVL(address token) returns uint256 {
    return totalSupplyByToken[token];
}

function balanceOfCVL(address token, address a) returns uint256 {
    return balanceByToken[token][a];
}

function allowanceCVL(address token, address a, address b) returns uint256 {
    return allowanceByToken[token][a][b];
}

function approveCVL(address token, address approver, address spender, uint256 amount) returns bool {
    // should be randomly reverting xxx
    bool nondetSuccess;
    if (!nondetSuccess) return false;

    allowanceByToken[token][approver][spender] = amount;
    return true;
}

function transferFromCVL(address token, address spender, address from, address to, uint256 amount) returns bool {
    // should be randomly reverting xxx
    bool nondetSuccess;
    if (!nondetSuccess) return false;

    if (allowanceByToken[token][from][spender] < amount) return false;
    allowanceByToken[token][from][spender] = assert_uint256(allowanceByToken[token][from][spender] - amount);
    return transferCVL(token, from, to, amount);
}

function transferCVL(address token, address from, address to, uint256 amount) returns bool {
    // should be randomly reverting xxx
    bool nondetSuccess;
    if (!nondetSuccess) return false;

    if(balanceByToken[token][from] < amount) return false;
    balanceByToken[token][from] = assert_uint256(balanceByToken[token][from] - amount);
    balanceByToken[token][to] = require_uint256(balanceByToken[token][to] + amount);  // We neglect overflows.
    return true;
}