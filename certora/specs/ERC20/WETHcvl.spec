using DummyWeth as weth; // we are limited by the fact that we cannot do transfers from CVL
using Utilities as utils;

methods {
    // Utilities
    function Utilities.justRevert() external envfree;

    // WETH
    function _.deposit() external with (env e) => wethDeposit(calledContract, e.msg.sender, e.msg.value) expect void;
    function _.withdraw(uint256 amount) external with (env e) => wethWithdraw(calledContract, e.msg.sender, amount) expect void;
}

function wethDeposit(address target, address caller, uint256 value) {
    // should be reverting if target != weth. Instead, we will use a contract to revert
    if (target != weth) {
        utils.justRevert(); // check this works xxx
    } else {
        // money will be transferred because of the payability of deposit
        env e2;
        require e2.msg.sender == caller;
        require e2.msg.value == value;
        weth.deposit(e2);
    }
}

function wethWithdraw(address target, address caller, uint256 amount) {
    // should be reverting if target != weth. Instead, we will use a contract to revert
    if (target != weth) {
        utils.justRevert(); // check this works xxx
    } else {
        env e2;
        require e2.msg.sender == caller;
        weth.withdraw(e2, amount);
    }
}
