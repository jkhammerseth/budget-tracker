<script>
    import ApexChart from 'svelte-apexcharts';
    import { filteredExpenses } from '../../stores/filteredExpenses'; // Import your store
    import { selectedMonth, selectedYear } from '../../stores/selectionMode';

    let chartOptions = {};
    let chartSeries = [];

    $: updateChartData();

    function updateChartData() {
        const daysInMonth = new Date($selectedYear, $selectedMonth + 1, 0).getDate(); // Get number of days in the month
        const dataByDay = Array(daysInMonth).fill(0);

        // Aggregate expenses by day
        $filteredExpenses.forEach((expense) => {
            const expenseDate = new Date(expense.payment_date);
            const expenseMonth = expenseDate.getMonth();
            const expenseYear = expenseDate.getFullYear();

            if (expenseMonth === $selectedMonth && expenseYear === $selectedYear) {
                const day = expenseDate.getDate() - 1; // Day (0-indexed for array)
                dataByDay[day] += expense.amount;
            }
        });

        // Update the chart data
        chartOptions = {
            chart: {
                type: 'bar',
                height: 350,
                toolbar: {
                    show: true,
                },
                zoom: {
                    enabled: true,
                },
            },
            title: {
                text: `Daily Expenses for ${$selectedMonth + 1}/${$selectedYear}`,
                align: 'center',
                style: {
                    fontFamily: 'Inter, Arial, sans-serif',
                    fontWeight: 'bold',
                },
            },
            xaxis: {
                categories: Array.from({ length: daysInMonth }, (_, i) => i + 1), // Days of the month
                title: {
                    text: 'Day of the Month',
                },
            },
            yaxis: {
                title: {
                    text: 'Amount Spent',
                },
            },
            tooltip: {
                y: {
                    formatter: (val) => `$${val.toFixed(2)}`,
                },
            },
            dataLabels: {
                enabled: false, // Optional: Disable for cleaner look
            },
            colors: ['#4bc0c0'],
        };

        chartSeries = [
            {
                name: 'Daily Expenses',
                data: dataByDay,
            },
        ];
    }
</script>

<div style="width: 100%; max-width: 1000px; margin: auto;">
    <ApexChart
        options={chartOptions}
        series={chartSeries}
        type="bar"
        height="350"
    />
</div>
