import json
import glob
import os
import argparse
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from collections import defaultdict

def main():
    parser = argparse.ArgumentParser(description='Analyze ATSE scalability results')
    parser.add_argument('--input', required=True, help='Directory containing JSONL result files')
    parser.add_argument('--output', required=True, help='Directory to save graphs and reports')
    args = parser.parse_args()

    # Ensure output dirs exist
    os.makedirs(os.path.join(args.output, 'graphs'), exist_ok=True)
    os.makedirs(os.path.join(args.output, 'reports'), exist_ok=True)

    print(f"Loading results from {args.input}...")
    
    raw_data = []
    files = glob.glob(os.path.join(args.input, '*.jsonl'))
    
    if not files:
        print("No result files found!")
        return

    for fpath in files:
        with open(fpath, 'r') as f:
            for line in f:
                try:
                    data = json.loads(line)
                    if 'test_meta' in data:
                        raw_data.append(data)
                except json.JSONDecodeError:
                    continue

    print(f"Loaded {len(raw_data)} metrics entries.")
    
    df = process_data(raw_data)
    generate_graphs(df, args.output)
    generate_report(df, args.output)

def process_data(raw_data):
    records = []
    for item in raw_data:
        meta = item['test_meta']
        # Convert memory to MB if needed (though our metric is already MB)
        # Convert duration to ms
        
        record = {
            'size_category': meta['size_category'],
            'file_count': int(meta['file_count']),
            'language': meta['language'],
            'query_type': meta['query_type'],
            'iteration': int(meta['iteration']),
            'outcome': meta['outcome'],
            
            # Metrics
            'duration_ms': float(item.get('execution_duration_ms', 0)),
            'memory_peak_mb': float(item.get('memory_peak_mb', 0)),
            'files_processed': int(item.get('files_processed', 0)),
            'throughput': 0.0
        }
        
        # Calculate throughput (files/sec)
        if record['duration_ms'] > 0:
            record['throughput'] = record['files_processed'] / (record['duration_ms'] / 1000.0)
            
        records.append(record)
        
    return pd.DataFrame(records)

def generate_graphs(df, output_dir):
    print("Generating graphs...")
    graph_dir = os.path.join(output_dir, 'graphs')
    
    # Filter for successful runs only for performance graphs
    success_df = df[df['outcome'] == 'success']
    
    # 1. Execution Time vs File Count (Log-Log)
    plt.figure(figsize=(10, 6))
    
    stats = success_df.groupby('file_count').agg({
        'duration_ms': ['mean', 'median', lambda x: np.percentile(x, 95)]
    }).reset_index()
    stats.columns = ['file_count', 'mean', 'median', 'p95']
    
    plt.loglog(stats['file_count'], stats['mean'], 'o-', label='Mean Time')
    plt.loglog(stats['file_count'], stats['p95'], 's--', label='P95 Time', alpha=0.5)
    
    plt.xlabel('File Count')
    plt.ylabel('Execution Time (ms)')
    plt.title('ATSE Search: Time vs Database Size')
    plt.legend()
    plt.grid(True, which="both", ls="-", alpha=0.2)
    plt.savefig(os.path.join(graph_dir, 'time_vs_files.png'))
    plt.close()
    
    # 2. Memory vs File Count
    plt.figure(figsize=(10, 6))
    
    mem_stats = success_df.groupby('file_count')['memory_peak_mb'].mean().reset_index()
    
    plt.loglog(mem_stats['file_count'], mem_stats['memory_peak_mb'], 'o-', color='green')
    
    plt.xlabel('File Count')
    plt.ylabel('Peak Memory (MB)')
    plt.title('ATSE Search: Memory Usage vs Database Size')
    plt.grid(True, which="both", ls="-", alpha=0.2)
    plt.savefig(os.path.join(graph_dir, 'memory_vs_files.png'))
    plt.close()
    
    # 3. Throughput vs File Count
    plt.figure(figsize=(10, 6))
    
    tp_stats = success_df.groupby('file_count')['throughput'].mean().reset_index()
    
    plt.semilogx(tp_stats['file_count'], tp_stats['throughput'], 'o-', color='orange')
    
    plt.xlabel('File Count')
    plt.ylabel('Throughput (files/sec)')
    plt.title('ATSE Search Throughput Stability')
    plt.grid(True, which="both", ls="-", alpha=0.2)
    plt.savefig(os.path.join(graph_dir, 'throughput_vs_files.png'))
    plt.close()
    
    # 4. Success Rate
    plt.figure(figsize=(10, 6))
    
    success_rates = df.groupby('file_count').apply(
        lambda x: (x['outcome'] == 'success').mean() * 100
    ).reset_index(name='rate')
    
    plt.semilogx(success_rates['file_count'], success_rates['rate'], 'o-', color='purple')
    
    plt.xlabel('File Count')
    plt.ylabel('Success Rate (%)')
    plt.title('Search Success Rate at Scale')
    plt.ylim(0, 105)
    plt.grid(True)
    plt.savefig(os.path.join(graph_dir, 'success_rate_vs_files.png'))
    plt.close()

def generate_report(df, output_dir):
    print("Generating report...")
    success_df = df[df['outcome'] == 'success']
    
    report_path = os.path.join(output_dir, 'reports', 'scalability_report.md')
    
    with open(report_path, 'w') as f:
        f.write("# ATSE Scalability Benchmark Report\n\n")
        
        # Summary Stats
        f.write("## Executive Summary\n\n")
        f.write(f"- **Total Runs**: {len(df)}\n")
        f.write(f"- **Max File Count**: {df['file_count'].max()}\n")
        f.write(f"- **Overall Success Rate**: {(df['outcome']=='success').mean()*100:.1f}%\n\n")
        
        # Performance Table
        f.write("## Performance by Scale\n\n")
        f.write("| Size | Files | Mean Time (ms) | P95 Time (ms) | Peak Mem (MB) | Throughput (files/s) |\n")
        f.write("|---|---|---|---|---|---|\n")
        
        stats = success_df.groupby(['size_category', 'file_count']).agg({
            'duration_ms': ['mean', lambda x: np.percentile(x, 95)],
            'memory_peak_mb': 'mean',
            'throughput': 'mean'
        }).sort_values(('size_category'), key=lambda x: x.map({'tiny':0, 'small':1, 'medium':2, 'large':3, 'huge':4, 'massive':5}))
        
        for idx, row in stats.iterrows():
            size_cat, count = idx
            mean_time = row[('duration_ms', 'mean')]
            p95_time = row[('duration_ms', '<lambda_0>')]
            mem = row[('memory_peak_mb', 'mean')]
            tp = row[('throughput', 'mean')]
            
            f.write(f"| {size_cat} | {count} | {mean_time:.0f} | {p95_time:.0f} | {mem:.1f} | {tp:.0f} |\n")
            
        f.write("\n")
        
        # Failed Runs
        failures = df[df['outcome'] != 'success']
        if not failures.empty:
            f.write("## Failures & Timeouts\n\n")
            f.write("| Size | Query | Outcome |\n")
            f.write("|---|---|---|\n")
            for _, row in failures.iterrows():
                f.write(f"| {row['size_category']} | {row['query_type']} | {row['outcome']} |\n")
        
    print(f"Report saved to {report_path}")

if __name__ == "__main__":
    main()
