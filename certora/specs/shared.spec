function pay_and_havoc(address receiver, env e) {
    if (e.msg.sender == receiver) {
        utils.havocAll();
        return;
    }

    uint oldBalanceSender = nativeBalances[e.msg.sender];
    uint oldBalanceRecipient = nativeBalances[receiver];
    utils.havocAll();
    require nativeBalances[e.msg.sender] == oldBalanceSender - e.msg.value;
    require nativeBalances[receiver] == oldBalanceRecipient + e.msg.value;
}