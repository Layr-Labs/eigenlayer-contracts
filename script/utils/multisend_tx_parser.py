import sys
import binascii

def parse_multisend_transactions(tx_bytes: bytes):
    i = 0
    txs = []

    while i < len(tx_bytes):
        if i + 1 + 20 + 32 + 32 > len(tx_bytes):
            raise ValueError(f"Not enough data for header at offset {i}")

        operation = tx_bytes[i]
        if operation != 0:
            raise ValueError(f"Invalid operation at offset {i}: expected 0, got {operation}")
        i += 1

        to = tx_bytes[i:i+20]
        i += 20

        value = int.from_bytes(tx_bytes[i:i+32], byteorder='big')
        i += 32

        data_length = int.from_bytes(tx_bytes[i:i+32], byteorder='big')
        i += 32

        if i + data_length > len(tx_bytes):
            raise ValueError(f"Not enough data for transaction data at offset {i}")

        data = tx_bytes[i:i+data_length]
        i += data_length

        txs.append({
            'operation': operation,
            'to': '0x' + to.hex(),
            'value': value,
            'data_length': data_length,
            'data': '0x' + data.hex()
        })

    return txs


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(f"Usage: {sys.argv[0]} <0x-prefixed hex string>")
        sys.exit(1)

    hex_input = sys.argv[1]
    if hex_input.startswith("0x"):
        hex_input = hex_input[2:]

    try:
        tx_bytes = binascii.unhexlify(hex_input)
    except binascii.Error:
        print("Invalid hex string.")
        sys.exit(1)

    try:
        transactions = parse_multisend_transactions(tx_bytes)
    except Exception as e:
        print(f"Error parsing transactions: {e}")
        sys.exit(1)

    for idx, tx in enumerate(transactions):
        print(f"Transaction {idx + 1}:")
        print(f"  operation: {tx['operation']}")
        print(f"  to: {tx['to']}")
        print(f"  value: {tx['value']}")
        print(f"  data_length: {tx['data_length']}")
        print(f"  data: {tx['data']}")
