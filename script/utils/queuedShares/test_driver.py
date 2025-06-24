#!/usr/bin/env python3
"""
Foundry Test Driver for SlashingLib Analysis

DEPENDENCIES:
Run this command to install all required packages:
    pip install pandas numpy matplotlib seaborn

Or individually:
    pip install pandas>=1.5.0
    pip install numpy>=1.20.0
    pip install matplotlib>=3.5.0
    pip install seaborn>=0.11.0
"""

# Standard imports
import sys
import subprocess
from pathlib import Path
from decimal import Decimal, getcontext
import json
import argparse
from typing import Dict, List, Tuple, Optional
import os

# Third-party imports (will fail with clear error if not installed)
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns

# Set high precision for exact calculations
getcontext().prec = 50

class SlashingTestDriver:
    """Driver class for running and analyzing SlashingLib tests"""
    
    def __init__(self, output_dir: str = "./output"):
        self.WAD = Decimal('1000000000000000000')  # 1e18
        self.results = {}
        
        # Keep output directory relative to current working directory
        self.output_dir = Path(output_dir)
        self.output_dir.mkdir(exist_ok=True)
        print(f"📁 Output directory: {self.output_dir.absolute()}")
        
    def run_foundry_test(self, config: Dict) -> str:
        """Run Foundry test with specified configuration"""
        
        output_file = config.get('output_file', str(self.output_dir / 'slashing_analysis.csv'))
        fuzz_runs = config.get('fuzz_runs', 100)
        fuzz_seed = config.get('fuzz_seed', None)
        test_name = config.get('test_name', 'testFuzz_scaleForBurning')
        
        print(f"🚀 Running Foundry test with {fuzz_runs} runs...")
        print(f"📄 Output file: {output_file}")
        
        # Build command
        cmd = [
            'forge', 'test',
            '--match-test', test_name,
            '--fuzz-runs', str(fuzz_runs),
            '-v'
        ]
        
        if fuzz_seed is not None:
            cmd.extend(['--fuzz-seed', str(fuzz_seed)])
        
        # Convert to absolute path so Foundry can find it regardless of working directory
        absolute_output_file = str(Path(output_file).resolve())
        
        # Set environment variables
        env = {
            **subprocess.os.environ,
            'CSV_OUTPUT_FILE': absolute_output_file,  # Use absolute path
            **config.get('env', {})
        }
        
        # Run Foundry from current directory (not repo root)
        script_dir = Path(__file__).parent
        
        try:
            print(f"🔧 Command: {' '.join(cmd)}")
            print(f"🌍 CSV_OUTPUT_FILE: {absolute_output_file}")
            print(f"📁 Working directory: {script_dir}")
            
            result = subprocess.run(
                cmd,
                capture_output=True,
                text=True,
                timeout=300,
                env=env,
                cwd=script_dir  # Run from script directory, not repo root
            )
            
            # Check if test actually ran fuzz cases
            if "fuzz runs" in result.stdout or "runs" in result.stdout:
                print("✅ Fuzz test executed")
            else:
                print("⚠️ Warning: Fuzz test may not have generated data")
                print("STDOUT:", result.stdout[-500:])  # Last 500 chars
            
            if result.returncode != 0:
                print("❌ Foundry test failed!")
                print("STDERR:", result.stderr)
                return None
            
            # Verify file exists and has data
            if not Path(output_file).exists():
                print(f"❌ Output file not created: {output_file}")
                return None
            
            # Check if file has actual data (more than just header)
            with open(output_file, 'r') as f:
                lines = f.readlines()
                if len(lines) <= 1:
                    print(f"⚠️ Warning: CSV file only has {len(lines)} lines (header only)")
                    print("This usually means:")
                    print("  1. vm.assume() conditions are too restrictive")
                    print("  2. All fuzz inputs were rejected")
                    print("  3. Test didn't actually run")
                    return None
            
            print(f"✅ Foundry test completed - {len(lines)-1} data rows generated")
            return output_file
            
        except subprocess.TimeoutExpired:
            print("⏰ Test timed out")
            return None
        except Exception as e:
            print(f"💥 Error running test: {e}")
            return None
    
    def read_csv_data(self, csv_file: str) -> Optional[pd.DataFrame]:
        """Read and validate CSV data from Foundry test with proper data types"""
        
        if not Path(csv_file).exists():
            print(f"❌ CSV file not found: {csv_file}")
            return None
        
        try:
            # Read CSV with explicit data types
            df = pd.read_csv(csv_file, dtype={
                'scaled_shares': 'float64',
                'prev_max_mag': 'float64', 
                'new_max_mag': 'float64',
                'slashed_amount': 'float64'
            })
            
            print(f"📊 Loaded {len(df)} rows from {csv_file}")
            print(f"📋 Columns: {list(df.columns)}")
            
            # Validate required columns
            required_cols = ['scaled_shares', 'prev_max_mag', 'new_max_mag', 'slashed_amount']
            missing_cols = [col for col in required_cols if col not in df.columns]
            
            if missing_cols:
                print(f"❌ Missing required columns: {missing_cols}")
                return None
            
            # Check for any non-numeric data
            for col in required_cols:
                if df[col].dtype == 'object':  # String type
                    print(f"⚠️ Warning: Column '{col}' contains non-numeric data")
                    # Try to convert to numeric, replacing errors with NaN
                    df[col] = pd.to_numeric(df[col], errors='coerce')
                    
                    # Drop rows with NaN values
                    nan_count = df[col].isna().sum()
                    if nan_count > 0:
                        print(f"🧹 Dropping {nan_count} rows with invalid data in column '{col}'")
                        df = df.dropna(subset=[col])
            
            # Verify we still have data
            if len(df) == 0:
                print("❌ No valid data rows after cleaning")
                return None
            
            print(f"✅ Data validation complete - {len(df)} valid rows")
            print(f"📊 Data types: {df.dtypes.to_dict()}")
            
            return df
            
        except Exception as e:
            print(f"💥 Error reading CSV: {e}")
            
            # Debug: Show first few lines of the file
            try:
                with open(csv_file, 'r') as f:
                    lines = f.readlines()[:5]
                    print("📄 First few lines of CSV:")
                    for i, line in enumerate(lines):
                        print(f"  {i}: {line.strip()}")
            except:
                pass
            
            return None
    
    def calculate_theoretical_results(self, df: pd.DataFrame) -> pd.DataFrame:
        """Add theoretical calculations and error analysis to the dataframe"""
        
        print("🧮 Calculating theoretical results...")
        
        # Calculate theoretical slashed amount using high precision
        theoretical_results = []
        errors = []
        relative_errors = []
        
        for _, row in df.iterrows():
            scaled_shares = Decimal(str(row['scaled_shares']))
            prev_max_mag = Decimal(str(row['prev_max_mag']))
            new_max_mag = Decimal(str(row['new_max_mag']))
            solidity_result = Decimal(str(row['slashed_amount']))
            
            # Theoretical calculation: scaledShares * (prevMaxMag - newMaxMag) / WAD
            magnitude_diff = prev_max_mag - new_max_mag
            theoretical_result = (scaled_shares * magnitude_diff) / self.WAD
            
            # Calculate error
            error = abs(solidity_result - theoretical_result)
            relative_error = (error / theoretical_result * 100) if theoretical_result > 0 else 0
            
            theoretical_results.append(float(theoretical_result))
            errors.append(float(error))
            relative_errors.append(float(relative_error))
        
        # Add calculated columns
        df['theoretical_result'] = theoretical_results
        df['error'] = errors
        df['relative_error_percent'] = relative_errors
        
        # Add additional analysis columns
        df['magnitude_ratio'] = df['new_max_mag'] / df['prev_max_mag']
        df['slash_percentage'] = ((df['prev_max_mag'] - df['new_max_mag']) / df['prev_max_mag'] * 100)
        
        print("✅ Theoretical calculations completed")
        return df
    
    def analyze_precision(self, df: pd.DataFrame) -> Dict:
        """Perform comprehensive precision analysis with better error handling"""
        
        print("🔍 Performing precision analysis...")
        
        analysis = {
            'total_cases': len(df),
            'data_types': df.dtypes.to_dict(),
        }
        
        # Safe error analysis
        if 'error' in df.columns:
            error_series = pd.to_numeric(df['error'], errors='coerce').dropna()
            if len(error_series) > 0:
                analysis.update({
                    'cases_with_error': len(error_series[error_series > 0]),
                    'cases_with_large_error': len(error_series[error_series > 1]),
                    'max_error': float(error_series.max()),
                    'mean_error': float(error_series.mean()),
                    'median_error': float(error_series.median()),
                    'std_error': float(error_series.std()),
                })
            else:
                analysis['error_note'] = 'No valid error data found'
        
        # Safe relative error analysis
        if 'relative_error_percent' in df.columns:
            rel_error_series = pd.to_numeric(df['relative_error_percent'], errors='coerce').dropna()
            if len(rel_error_series) > 0:
                analysis.update({
                    'max_relative_error': float(rel_error_series.max()),
                    'mean_relative_error': float(rel_error_series.mean()),
                })
        
        # Error distribution by slash percentage ranges (with better handling)
        if 'slash_percentage' in df.columns and 'error' in df.columns:
            try:
                # Fix the pandas warning by adding observed=True
                df['slash_range'] = pd.cut(df['slash_percentage'], 
                                          bins=[0, 10, 25, 50, 75, 90, 100], 
                                          labels=['0-10%', '10-25%', '25-50%', '50-75%', '75-90%', '90-100%'])
                
                error_by_slash = df.groupby('slash_range', observed=True)['error'].agg(['count', 'mean', 'max']).round(6)
                analysis['error_by_slash_range'] = error_by_slash.to_dict()
            except Exception as e:
                analysis['error_by_slash_note'] = f'Could not analyze by slash range: {e}'
        
        print("✅ Precision analysis completed")
        return analysis
    
    def generate_visualizations(self, df: pd.DataFrame, output_subdir: str = 'plots'):
        """Generate comprehensive visualizations with better error handling"""
        
        print("📊 Generating visualizations...")
        
        if len(df) == 0:
            print("❌ Cannot generate visualizations: DataFrame is empty")
            return
        
        if len(df) < 10:
            print(f"⚠️ Warning: Only {len(df)} data points - visualizations may not be meaningful")
        
        # Create plots subdirectory within output directory
        plot_dir = self.output_dir / output_subdir
        plot_dir.mkdir(exist_ok=True)
        
        # Set style with error handling
        try:
            plt.style.use('seaborn-v0_8')
            sns.set_palette("husl")
        except:
            try:
                plt.style.use('seaborn')  # Fallback for older versions
            except:
                plt.style.use('default')  # Final fallback
        
        # Verify data types before plotting
        numeric_cols = ['theoretical_result', 'slashed_amount', 'error', 'relative_error_percent', 
                       'magnitude_ratio', 'slash_percentage']
        
        for col in numeric_cols:
            if col in df.columns and df[col].dtype == 'object':
                print(f"🔧 Converting column '{col}' to numeric")
                df[col] = pd.to_numeric(df[col], errors='coerce')
        
        try:
            # Figure 1: Error Distribution
            fig, ((ax1, ax2), (ax3, ax4)) = plt.subplots(2, 2, figsize=(15, 12))
            
            # Error histogram
            if 'error' in df.columns and df['error'].notna().any():
                ax1.hist(df['error'].dropna(), bins=min(50, len(df)//10 + 1), alpha=0.7, edgecolor='black')
                ax1.set_title('Error Distribution')
                ax1.set_xlabel('Absolute Error')
                ax1.set_ylabel('Frequency')
                ax1.axvline(x=1, color='red', linestyle='--', label='Error = 1')
                ax1.legend()
            else:
                ax1.text(0.5, 0.5, 'No error data', ha='center', va='center', transform=ax1.transAxes)
                ax1.set_title('Error Distribution (No Data)')
            
            # Error vs Slash Percentage
            if 'slash_percentage' in df.columns and 'error' in df.columns:
                valid_data = df[['slash_percentage', 'error']].dropna()
                if len(valid_data) > 0:
                    ax2.scatter(valid_data['slash_percentage'], valid_data['error'], alpha=0.6, s=20)
                    ax2.set_title('Error vs Slash Percentage')
                    ax2.set_xlabel('Slash Percentage (%)')
                    ax2.set_ylabel('Absolute Error')
                    ax2.axhline(y=1, color='red', linestyle='--', alpha=0.7)
                else:
                    ax2.text(0.5, 0.5, 'No valid data', ha='center', va='center', transform=ax2.transAxes)
            else:
                ax2.text(0.5, 0.5, 'Missing data', ha='center', va='center', transform=ax2.transAxes)
            
            # Relative Error Distribution
            if 'relative_error_percent' in df.columns:
                valid_data = df['relative_error_percent'].dropna()
                if len(valid_data) > 0:
                    ax3.hist(valid_data, bins=min(50, len(valid_data)//10 + 1), alpha=0.7, edgecolor='black')
                    ax3.set_title('Relative Error Distribution')
                    ax3.set_xlabel('Relative Error (%)')
                    ax3.set_ylabel('Frequency')
                else:
                    ax3.text(0.5, 0.5, 'No relative error data', ha='center', va='center', transform=ax3.transAxes)
            
            # Theoretical vs Actual Results
            if 'theoretical_result' in df.columns and 'slashed_amount' in df.columns:
                valid_data = df[['theoretical_result', 'slashed_amount']].dropna()
                if len(valid_data) > 0:
                    # Convert to numeric and handle the max calculation safely
                    theoretical = pd.to_numeric(valid_data['theoretical_result'], errors='coerce')
                    actual = pd.to_numeric(valid_data['slashed_amount'], errors='coerce')
                    
                    # Remove any remaining NaN values
                    mask = theoretical.notna() & actual.notna()
                    theoretical = theoretical[mask]
                    actual = actual[mask]
                    
                    if len(theoretical) > 0 and len(actual) > 0:
                        ax4.scatter(theoretical, actual, alpha=0.6, s=20)
                        max_val = max(theoretical.max(), actual.max())
                        ax4.plot([0, max_val], [0, max_val], 'r--', alpha=0.7, label='Perfect Match')
                        ax4.set_title('Theoretical vs Actual Results')
                        ax4.set_xlabel('Theoretical Result')
                        ax4.set_ylabel('Solidity Result')
                        ax4.legend()
                    else:
                        ax4.text(0.5, 0.5, 'No valid numeric data', ha='center', va='center', transform=ax4.transAxes)
                else:
                    ax4.text(0.5, 0.5, 'No comparison data', ha='center', va='center', transform=ax4.transAxes)
            
            plt.tight_layout()
            plt.savefig(plot_dir / 'precision_analysis.png', dpi=300, bbox_inches='tight')
            plt.close()
            
            print(f"✅ Basic visualizations saved to {plot_dir}/")
            
            # Skip complex visualizations if we have data issues
            if len(df) < 20:
                print("⏭️ Skipping complex visualizations due to insufficient data")
                return
            
            # Additional plots can go here...
            
        except Exception as e:
            print(f"⚠️ Error generating visualizations: {e}")
            print("📊 Generating basic summary instead...")
            
            # Create a simple text summary plot
            fig, ax = plt.subplots(1, 1, figsize=(10, 6))
            summary_text = f"""
            Data Summary:
            Total rows: {len(df)}
            Columns: {', '.join(df.columns)}
            
            Data types:
            {df.dtypes.to_string()}
            
            Sample data:
            {df.head().to_string()}
            """
            ax.text(0.1, 0.5, summary_text, transform=ax.transAxes, fontfamily='monospace', fontsize=8)
            ax.set_title('Data Summary (Visualization Failed)')
            ax.axis('off')
            plt.savefig(plot_dir / 'data_summary.png', dpi=300, bbox_inches='tight')
            plt.close()
    
    def save_enhanced_csv(self, df: pd.DataFrame, output_file: str):
        """Save enhanced CSV with all calculations"""
        df.to_csv(output_file, index=False)
        print(f"💾 Enhanced CSV saved to {output_file}")
    
    def save_analysis_report(self, analysis: Dict, output_file: str):
        """Save analysis report as JSON"""
        with open(output_file, 'w') as f:
            json.dump(analysis, f, indent=2, default=str)
        print(f"📋 Analysis report saved to {output_file}")
    
    def run_complete_analysis(self, config: Dict):
        """Run complete analysis pipeline"""
        
        print("🎯 Starting complete SlashingLib analysis...")
        print("=" * 50)
        
        # Change to repo root for consistency
        original_cwd = Path.cwd()
        repo_root = Path(__file__).parent
        while not (repo_root / 'foundry.toml').exists() and repo_root.parent != repo_root:
            repo_root = repo_root.parent
        
        try:
            os.chdir(repo_root)
            print(f"📁 Changed working directory to: {repo_root}")
            
            # Step 1: Run Foundry test
            csv_file = self.run_foundry_test(config)
            if not csv_file:
                print("❌ Failed to run Foundry test")
                return None
            
            # Step 2: Read CSV data
            df = self.read_csv_data(csv_file)
            if df is None:
                print("❌ Failed to read CSV data")
                return None
            
            # Step 3: Calculate theoretical results
            df = self.calculate_theoretical_results(df)
            
            # Step 4: Perform precision analysis
            analysis = self.analyze_precision(df)
            
            # Step 5: Generate visualizations
            self.generate_visualizations(df, config.get('plot_subdir', 'plots'))
            
            # Step 6: Save enhanced data and reports
            enhanced_csv = config.get('enhanced_csv', str(self.output_dir / 'enhanced_slashing_analysis.csv'))
            report_file = config.get('report_file', str(self.output_dir / 'analysis_report.json'))
            
            self.save_enhanced_csv(df, enhanced_csv)
            self.save_analysis_report(analysis, report_file)
            
            # Step 7: Print summary
            self._print_summary(analysis)
            
            print("🎉 Complete analysis finished!")
            return df, analysis
            
        finally:
            # Always restore original working directory
            os.chdir(original_cwd)
    
    def _print_summary(self, analysis: Dict):
        """Print analysis summary to console"""
        
        print("\n" + "=" * 50)
        print("📊 ANALYSIS SUMMARY")
        print("=" * 50)
        print(f"Total test cases: {analysis['total_cases']}")
        print(f"Cases with errors: {analysis['cases_with_error']} ({analysis['cases_with_error']/analysis['total_cases']*100:.2f}%)")
        print(f"Cases with large errors (>1): {analysis['cases_with_large_error']}")
        print(f"Maximum error: {analysis['max_error']:.6f}")
        print(f"Average error: {analysis['mean_error']:.6f}")
        print(f"Maximum relative error: {analysis['max_relative_error']:.6f}%")
        
        if 'problematic_cases' in analysis:
            print(f"\n🚨 Found {analysis['problematic_cases']['count']} problematic cases")
            print("Examples of problematic cases:")
            for i, case in enumerate(analysis['problematic_cases']['examples']):
                print(f"  {i+1}. Shares: {case['scaled_shares']}, Prev: {case['prev_max_mag']}, New: {case['new_max_mag']}, Error: {case['error']}")

def main():
    """Main function with command line interface"""
    
    parser = argparse.ArgumentParser(description='Drive SlashingLib Foundry tests and perform analysis')
    parser.add_argument('--fuzz-runs', type=int, default=500, help='Number of fuzz runs')
    parser.add_argument('--fuzz-seed', type=int, help='Fuzz seed for reproducibility')
    parser.add_argument('--output-dir', default='output', help='Output directory relative to repo root')
    parser.add_argument('--output-csv', default='slashing_test_output.csv', help='Output CSV filename (within output dir)')
    parser.add_argument('--enhanced-csv', default='enhanced_slashing_analysis.csv', help='Enhanced CSV filename (within output dir)')
    parser.add_argument('--report', default='analysis_report.json', help='Analysis report filename (within output dir)')
    parser.add_argument('--plot-subdir', default='plots', help='Subdirectory for plots (within output dir)')
    parser.add_argument('--test-name', default='testFuzz_scaleForBurning', help='Foundry test name')
    
    args = parser.parse_args()
    
    # Create driver with output directory
    driver = SlashingTestDriver(output_dir=args.output_dir)
    
    # Configuration - all paths now relative to output directory
    config = {
        'fuzz_runs': args.fuzz_runs,
        'fuzz_seed': args.fuzz_seed,
        'output_file': str(Path(args.output_dir) / args.output_csv),
        'enhanced_csv': str(driver.output_dir / args.enhanced_csv),
        'report_file': str(driver.output_dir / args.report),
        'plot_subdir': args.plot_subdir,
        'test_name': args.test_name,
    }
    
    # Run analysis
    result = driver.run_complete_analysis(config)
    if result is not None:
        df, analysis = result
    else:
        print("❌ Analysis failed")
        sys.exit(1)

if __name__ == "__main__":
    main()