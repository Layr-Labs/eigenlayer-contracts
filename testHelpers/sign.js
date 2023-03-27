const Web3 = require('web3');
const web3 = new Web3()

main()

async function main() {
    const signature = await web3.eth.accounts.sign("0xfb56e77dd50bc6258879ea2401037cbb4227cf266191e54c9d0f7fa6fbd1e873", "0xe88d9d864d5d731226020c5d2f02b62a4ce2a4534a39c225d32d3db795f83319");
    console.log(signature)
}

function toEthSignedMessageHash(messageHex) {
    const messageBuffer = Buffer.from(messageHex.substring(2), 'hex');
    const prefix = Buffer.from(`\u0019Ethereum Signed Message:\n${messageBuffer.length}`);
    return web3.utils.sha3(Buffer.concat([prefix, messageBuffer]));
}

0xd014256b124b583eb14192505340f3c785c31c3412c7daa358072c0b81b85aa5
0xf4c6a5d675588f311a9061cf6cbf6eacd95a15f42c9d111a9f65767dfe2e6ff8
