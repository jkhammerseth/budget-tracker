<script>
    import { onMount, onDestroy } from 'svelte';
    import { Chart, registerables } from 'chart.js';
    import { filteredExpenses } from '../../stores/filteredExpenses'; // Import your store for filtered expenses
    import ChartDataLabels from 'chartjs-plugin-datalabels'; // Plugin for displaying labels
    import ZoomPlugin from 'chartjs-plugin-zoom'; // Plugin for zooming
    import { formatAmount } from '../../utility/functions';

    let chart;

    onMount(() => {
        // Register all necessary Chart.js components and plugins
        Chart.register(...registerables, ChartDataLabels, ZoomPlugin);

        // Prepare data for the bar chart
        const dataByMonth = Array(12).fill(0);

        filteredExpenses.subscribe((expenses) => {
            expenses.forEach((expense) => {
                const month = new Date(expense.payment_date).getMonth();
                dataByMonth[month] += expense.amount;
            });
        });

        // Get the canvas context
        const ctx = document.getElementById('monthlyExpensesChart').getContext('2d');

        // Create the bar chart
        chart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
                datasets: [
                    {
                        label: 'Monthly Expenses',
                        data: dataByMonth,
                        backgroundColor: 'rgba(75, 192, 192, 0.9)',
                        borderColor: 'rgba(75, 192, 192, 1)',
                        borderWidth: 2,
                        borderSkipped: false,
                    },
                ],
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    tooltip: {
                        backgroundColor: 'rgba(75, 192, 192, 0.9)',
                        titleColor: '#fff',
                        bodyColor: '#fff',
                        borderColor: 'rgba(0,0,0,0.1)',
                        borderWidth: 1,
                        padding: 10,
                    },
                    datalabels: {
                        color: '#333',
                        anchor: 'end',
                        align: 'top',
                        formatter: (value) => `${formatAmount(value)}`,
                        font: {
                            family: 'Inter, Arial, sans-serif',
                            size: 12,
                            weight: 'bold',
                        },
                    },
                    title: {
                        display: false,
                        text: 'Monthly Expenses',
                        font: {
                            family: 'Inter, Arial, sans-serif',
                            size: 18,
                            weight: 'bold',
                        },
                        color: '#333',
                    },
                    legend: {
                        display: false, // Optional: Disable legend for a cleaner look
                    },
                    zoom: {
                        pan: {
                            enabled: true,
                            mode: 'x',
                        },
                        zoom: {
                            wheel: {
                                enabled: true,
                            },
                            mode: 'x',
                        },
                    },
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        grid: {
                            color: 'rgba(200, 200, 200, 0.2)', // Subtle grid lines
                        },
                        ticks: {
                            color: '#666', // Subtle Y-axis labels
                            font: {
                                family: 'Inter, Arial, sans-serif',
                                size: 12,
                            },
                        },
                    },
                    x: {
                        grid: {
                            display: false, // No vertical grid lines
                        },
                        ticks: {
                            color: '#666', // Subtle X-axis labels
                            font: {
                                family: 'Inter, Arial, sans-serif',
                                size: 12,
                            },
                        },
                    },
                },
                animations: {
                    tension: {
                        duration: 1000,
                        easing: 'easeInOutQuart',
                        from: 0.3,
                        to: 0,
                    },
                },
            },
        });
    });

    // Cleanup the chart on component destroy
    onDestroy(() => {
        if (chart) {
            chart.destroy();
        }
    });
</script>

<div style="width: 100%; max-width: 1000px; margin: auto;">
    <canvas id="monthlyExpensesChart" width="800" height="600"></canvas>
</div>



