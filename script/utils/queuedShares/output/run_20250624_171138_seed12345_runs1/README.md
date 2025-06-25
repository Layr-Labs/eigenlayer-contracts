# SlashingLib Analysis Run

## Run Configuration
- **Run ID**: 20250624_171138_seed12345_runs1
- **Timestamp**: 2025-06-24 17:11:39
- **Fuzz Runs**: 1
- **Fuzz Seed**: 12345
- **Test**: testFuzz_scaleForBurning

## Results Summary
- **Total test cases**: 1
- **Cases with errors**: 1 (100.00%)
- **Cases with large errors (>1)**: 1
- **Maximum error**: 2781.706018
- **Average error**: 2781.706018
- **Maximum relative error**: 100.000000%

## Files Generated
- `slashing_test_output.csv` - Raw test data from Foundry
- `enhanced_slashing_analysis.csv` - Enhanced data with theoretical calculations
- `analysis_report.json` - Detailed analysis in JSON format
- `run_metadata.json` - Run configuration and metadata
- `plots/precision_analysis.png` - Visualization of precision analysis

## Reproducibility
To reproduce this exact run:
```bash
python3 test_driver.py --fuzz-runs 1 --fuzz-seed 12345 --no-unique
```
