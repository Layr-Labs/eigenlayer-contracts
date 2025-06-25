# SlashingLib Analysis Run

## Run Configuration
- **Run ID**: 20250624_171223_seed12345_runs500
- **Timestamp**: 2025-06-24 17:12:25
- **Fuzz Runs**: 500
- **Fuzz Seed**: 12345
- **Test**: testFuzz_scaleForBurning

## Results Summary
- **Total test cases**: 500
- **Cases with errors**: 500 (100.00%)
- **Cases with large errors (>1)**: 495
- **Maximum error**: 340056190538802240290778805243914223616.000000
- **Average error**: 3349870372353426682558070372228399104.000000
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
python3 test_driver.py --fuzz-runs 500 --fuzz-seed 12345 --no-unique
```
