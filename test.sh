#!/bin/bash

source .env

operator1address=$(cast w a --private-key $OPERATOR_1)
operator2address=$(cast w a --private-key $OPERATOR_2)
operator3address=$(cast w a --private-key $OPERATOR_3)
operator4address=$(cast w a --private-key $OPERATOR_4)

# cast send --rpc-url $RPC_URL --private-key $OPERATOR_1 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "registerAsOperator((address,address,uint32), string)" "($operator1address,0x0000000000000000000000000000000000000000,10)" "https://operator.com/123"
# cast send --rpc-url $RPC_URL --private-key $OPERATOR_2 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "registerAsOperator((address,address,uint32), string)" "($operator2address,0x0000000000000000000000000000000000000000,0)" "https://123qwe.com/opop"
# cast send --rpc-url $RPC_URL --private-key $OPERATOR_3 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "registerAsOperator((address,address,uint32), string)" "($operator3address,0x0000000000000000000000000000000000000000,0)" "https://geezlouise.com/13"
# cast send --rpc-url $RPC_URL --private-key $OPERATOR_4 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "registerAsOperator((address,address,uint32), string)" "($operator4address,0x0000000000000000000000000000000000000000,0)" "https://urmama.com/rich_dank"

# cast send --rpc-url $RPC_URL --private-key $OPERATOR_1 0x78697cd4EE4BdE0514fE8a7C61E8cB1A152B5d78 "registerOperatorWithCoordinator(bytes memory, bytes calldata)" "" ""
# cast send --rpc-url $RPC_URL --private-key $OPERATOR_2 0x78697cd4EE4BdE0514fE8a7C61E8cB1A152B5d78 "registerOperatorWithCoordinator(bytes memory, bytes calldata)" "" ""
# cast send --rpc-url $RPC_URL --private-key $OPERATOR_3 0x78697cd4EE4BdE0514fE8a7C61E8cB1A152B5d78 "registerOperatorWithCoordinator(bytes memory, bytes calldata)" "" ""
# cast send --rpc-url $RPC_URL --private-key $OPERATOR_4 0x78697cd4EE4BdE0514fE8a7C61E8cB1A152B5d78 "registerOperatorWithCoordinator(bytes memory, bytes calldata)" "" ""

cast send --rpc-url $RPC_URL --private-key $OPERATOR_1 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "updateOperatorMetadataURI(string calldata)" "https://operator-metadata.s3.amazonaws.com/operator_1.json"
cast send --rpc-url $RPC_URL --private-key $OPERATOR_2 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "updateOperatorMetadataURI(string calldata)" "https://operator-metadata.s3.amazonaws.com/operator_2.json"
cast send --rpc-url $RPC_URL --private-key $OPERATOR_3 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "updateOperatorMetadataURI(string calldata)" "https://operator-metadata.s3.amazonaws.com/operator_3.json"
cast send --rpc-url $RPC_URL --private-key $OPERATOR_4 0x45b4c4DAE69393f62e1d14C5fe375792DF4E6332 "updateOperatorMetadataURI(string calldata)" "https://operator-metadata.s3.amazonaws.com/operator_4.json"