// // SPDX-License-Identifier: BUSL-1.1
// pragma solidity ^0.8.12;

// import "./../../utils/Releasoor.s.sol";

// contract v0_4_2_ is Releasoor {

//     using TxBuilder for *;
//     using AddressUtils for *;

//     function deploy(Addresses memory addrs) internal override {
//         // If you're deploying contracts, do that here
//     }

//     function queueUpgrade(Addresses memory addrs) internal override returns (Tx[] memory executorTxns, uint eta) {
//         // If you're queueing an upgrade via the timelock, you can
//         // define and encode those transactions here
//     }

//     function executeUpgrade(Addresses memory addrs) internal override {
//         // Whether you are using the timelock or just making transactions
//         // from the ops multisig, you can define/encode those transactions here
//     }
// }