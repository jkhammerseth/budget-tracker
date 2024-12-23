<script>
    import { filteredExpenses } from '../stores/filteredExpenses';
    import { derived } from 'svelte/store';

    const categoryColors = {
        Food: 'rgba(255, 99, 132, 1)',
        Transportation: 'rgba(54, 162, 235, 1)',
        Housing: 'rgba(255, 206, 86, 1)',
        Utilities: 'rgba(75, 192, 192, 1)',
        Insurance: 'rgba(153, 102, 255, 1)',
        Healthcare: 'rgba(255, 159, 64, 1)',
        Entertainment: '#CDDC39',
        Clothing: '#FF5722',
        Miscellaneous: '#607D8B',
    };

    const maxRadius = 150;
    const svgWidth = 1200;
    const svgHeight = 400;
    const padding = 0;

    const expensesByCategory = derived(filteredExpenses, $filteredExpenses => {
        const totals = $filteredExpenses.reduce((acc, { Category, Amount }) => {
            acc[Category] = (acc[Category] || 0) + Amount;
            return acc;
        }, {});

        const maxTotal = Math.max(...Object.values(totals), 1);

        let positions = generateNonOverlappingPositions(totals, maxTotal);

        return positions.map(({ category, amount, radius, x, y, color }) => ({
            category, amount, radius, color, x, y
        }));
    });

    function generateNonOverlappingPositions(totals, maxTotal) {
        let positions = [];
        let attempts = 0;
        const maxAttempts = 5000;

        Object.entries(totals).forEach(([category, amount]) => {
            let radius = Math.sqrt((amount / maxTotal) * (maxRadius ** 2));
            radius = Math.min(radius, svgWidth / 4, svgHeight / 4); // Constrain radius
            let color = getCategoryColor(category);
            let position;
            let found = false;

            while (attempts < maxAttempts) {
                position = {
                    x: radius + Math.random() * (svgWidth - 2 * radius),
                    y: radius + Math.random() * (svgHeight - 2 * radius),
                    radius, category, amount, color
                };
                if (!isOverlapping(position, positions)) {
                    found = true;
                    break;
                }
                attempts++;
            }

            if (found) {
                positions.push(position);
            } else {
                console.warn(`Failed to place category: ${category} after ${attempts} attempts.`);
            }
        });

        return positions;
    }

    function isOverlapping(newPosition, positions) {
        return positions.some(pos => {
            let dx = newPosition.x - pos.x;
            let dy = newPosition.y - pos.y;
            let distance = Math.sqrt(dx * dx + dy * dy);
            return distance < newPosition.radius + pos.radius + padding - 5; // Reduce padding overlap slightly
        });
    }

    function getCategoryColor(category) {
        return categoryColors[category] || '#757575';
    }
</script>

<svg width="{svgWidth}" height="{svgHeight}">
    {#each $expensesByCategory as { category, amount, radius, color, x, y }}
        <g>
            <circle cx="{x}" cy="{y}" r="{radius}" fill="{color}">
                <title>{category}: {amount + ' kr'}</title>
            </circle>
            <text
                x="{x}"
                y="{y}"
                dy="0.3em"
                text-anchor="middle"
                fill="white"
                font-size="{Math.max(10, radius / 5)}"
                style="pointer-events: none;"
            >
                {category}
            </text>
        </g>
    {/each}
</svg>
