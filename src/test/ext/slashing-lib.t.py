import sys
import json
from decimal import Decimal

WAD = Decimal(1e18)

class SlashingLib:
    def mul_wad(self, x: int, y: int) -> int:
        return int((Decimal(x) * Decimal(y) / WAD))

    def div_wad(self, x: int, y: int) -> int:
        return int((Decimal(x) * WAD / Decimal(y)))

    def mul_wad_round_up(self, x: int, y: int) -> int:
        result = Decimal(x) * Decimal(y) / WAD
        return int(result.to_integral_value(rounding="ROUND_UP"))

    def scale_shares_for_queued_withdrawal(self, shares_to_withdraw: int, beacon_chain_scaling_factor: int, operator_magnitude: int) -> int:
        return int(Decimal(shares_to_withdraw) * WAD / Decimal(beacon_chain_scaling_factor)) * WAD / Decimal(operator_magnitude)

class Dispatcher:
    def __init__(self):
        self.slashing_lib = SlashingLib()
        self.args = json.loads(sys.argv[1])

    def test_scale_shares_for_queued_withdrawal(self):
        shares_to_withdraw = self.args['sharesToWithdraw']
        beacon_chain_scaling_factor = self.args['beaconChainScalingFactor']
        operator_magnitude = self.args['operatorMagnitude']
        result = self.slashing_lib.scale_shares_for_queued_withdrawal(shares_to_withdraw, beacon_chain_scaling_factor, operator_magnitude)
        
        # Encode the result as a 32-byte hex string.
        # Equivilent to `abi.encode(uint256(result))`.
        encoded_result = int(result).to_bytes(32, byteorder='big')
        
        # Foundry FFI return values must be printed to stdout.
        print(f'0x{encoded_result.hex()}') 

    def dispatch(self):
        if self.args['method'] == 'scaleSharesForQueuedWithdrawal':
            self.test_scale_shares_for_queued_withdrawal()

Dispatcher().dispatch()