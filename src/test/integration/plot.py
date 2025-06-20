import json
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker

# Load data from JSON file
with open('iterations.json', 'r') as f:
    data = json.load(f)

# Compute percent decreases
percent_drops = [(data[i - 1] - data[i]) / data[i - 1] * 100 for i in range(1, len(data))]
x_values = list(range(1, len(data)))

# Plotting
plt.figure(figsize=(10, 6))
plt.plot(x_values, percent_drops, marker='o', linestyle='-', markersize=1)

# Labels and title
plt.xlabel('iterations')
plt.ylabel('% decrease in DSF')
plt.title('Percent Decrease in DSF per Iteration')
plt.grid(True)

cliff_index = 1000  # replace with actual index where plateau begins
dsf_value = data[cliff_index]

# Format DSF in scientific notation
dsf_label = f'{dsf_value:.2e}'

plt.axvline(x=cliff_index, color='red', linestyle='--', label=f'Cliff @ {cliff_index}')

# Add vertical red text along the line
plt.text(
    x=cliff_index + 2,     # slight offset to avoid overlapping the line
    y=max(percent_drops) * 0.9,  # place vertically near the top
    s=dsf_label,
    color='red',
    rotation=-90,
    va='center',
    ha='left',
    fontsize=10,
    bbox=dict(facecolor='white', edgecolor='red', boxstyle='round,pad=0.3')
)

# Annotate iteration number at bottom
plt.text(
    x=cliff_index,
    y=-0.005,  # a bit below the X-axis; tweak if needed
    s=f'i = {cliff_index}',
    color='red',
    rotation=0,
    va='top',
    ha='center',
    fontsize=7,
    fontweight='bold',
    transform=plt.gca().get_xaxis_transform()
)

# Save the plot
plt.tight_layout()
plt.savefig('dsf_percent_deltas_plot4.png', dpi=300)
