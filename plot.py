import matplotlib.pyplot as plt
import numpy as np

# Data for each algorithm with 8 threads
data_8_threads = {
    'Algorithm': [
        'Parallel Linear Search', 'Parallel Binary Search',
        'Parallel Hash Search', 'Parallel Tree Search'
    ],
    'Small Dataset': {
        'Processing Time (s)': [0.000073375, 0.000194959, 0.000194959, 0.000084083],
        'Search Time (ns)': [62500, 54000, 21042, 387458],
        'Processed Elements': [100, 100, 61, 100]
    },
    'Large Dataset': {
        'Processing Time (s)': [0.954187, 0.894470, 4.745985, 1.143710],
        'Search Time (ns)': [0, 0, 72611700, 570754200],
        'Processed Elements': [10000000, 10000000, 6322042, 10000000]
    }
}

# X-axis positions
x = np.arange(len(data_8_threads['Algorithm']))  # X-axis positions
width = 0.35  # Width of the bars

# Set up the figure for two separate subplots
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(14, 6))

# Plot 1: Processing Time for Small and Large Dataset
ax1.bar(x - width / 2, data_8_threads['Small Dataset']['Processing Time (s)'], width, color='skyblue', label='Small Dataset (100 elements)')
ax1.bar(x + width / 2, data_8_threads['Large Dataset']['Processing Time (s)'], width, color='salmon', label='Large Dataset (10M elements)')
ax1.set_yscale('log')  # Apply logarithmic scale to y-axis
ax1.set_xlabel('Algorithm')
ax1.set_ylabel('Processing Time (s)')
ax1.set_title('Processing Time Comparison')
ax1.set_xticks(x)
ax1.set_xticklabels(data_8_threads['Algorithm'], rotation=45, ha='right')
ax1.legend()

# Plot 2: Search Time for Small and Large Dataset
ax2.bar(x - width / 2, data_8_threads['Small Dataset']['Search Time (ns)'], width, color='skyblue', label='Small Dataset (100 elements)')
ax2.bar(x + width / 2, data_8_threads['Large Dataset']['Search Time (ns)'], width, color='salmon', label='Large Dataset (10M elements)')
ax2.set_yscale('log')  # Apply logarithmic scale to y-axis
ax2.set_xlabel('Algorithm')
ax2.set_ylabel('Search Time (ns)')
ax2.set_title('Search Time Comparison')
ax2.set_xticks(x)
ax2.set_xticklabels(data_8_threads['Algorithm'], rotation=45, ha='right')
ax2.legend()

plt.tight_layout()
plt.show()
