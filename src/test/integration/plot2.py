import json
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker

# Load data
with open('iterations2.json', 'r') as f:
    data = json.load(f)

x_values = []
percent_drops = []

# Set up the figure before we draw vertical lines
plt.figure(figsize=(10, 6))

# Collect percent drops (skip increases) and draw red lines for increases
for i in range(1, len(data)):
    if data[i] <= data[i - 1]:
        drop = (data[i - 1] - data[i]) / data[i - 1] * 100
        percent_drops.append(drop)
        x_values.append(i)
    else:
        # DSF increased — draw vertical red line and annotate
        plt.axvline(x=i, color='red', linestyle='--', linewidth=2)

        prev_dsf = f'{data[i - 1]:.2e}'
        new_dsf = f'{data[i]:.2e}'
        label = f'{prev_dsf} → {new_dsf}'

        plt.text(
            x=i + 2,
            y=50*0.75,  # temporary value, updated after max_drop is known
            s=label,
            color='red',
            rotation=-90,
            va='center',
            ha='left',
            fontsize=8,
            bbox=dict(facecolor='white', edgecolor='red', boxstyle='round,pad=0.3'),
        )

        # Annotate iteration number at bottom
        plt.text(
            x=i,
            y=-0.005,  # a bit below the X-axis; tweak if needed
            s=f'i = {i}',
            color='red',
            rotation=0,
            va='top',
            ha='center',
            fontsize=7,
            fontweight='bold',
            transform=plt.gca().get_xaxis_transform()
        )

# Plot the valid percent drops
plt.plot(x_values, percent_drops, marker='o', linestyle='-', markersize=1)

# Axis labels and title
plt.xlabel('iterations')
plt.ylabel('% decrease in DSF')
plt.title('Percent Decrease in DSF per Iteration')
plt.grid(True)

# Set Y-axis limit to max % decrease
max_drop = max(percent_drops)
plt.ylim(0, max_drop * 1.05)  # 5% padding

# Optional: highlight a known cliff
cliff_index = 1003
dsf_value = data[cliff_index]
dsf_label = f'{dsf_value:.2e}'

plt.axvline(x=cliff_index, color='red', linestyle='--', label=f'Cliff @ {cliff_index}')
plt.text(
    x=cliff_index + 2,
    y=max_drop * 0.75,
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
    x=i,
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

# Save plot
plt.tight_layout()
plt.savefig('dsf_percent_deltas_plot3.png', dpi=300)
