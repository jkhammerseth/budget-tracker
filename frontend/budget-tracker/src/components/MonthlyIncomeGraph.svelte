<script>
    import { onMount } from 'svelte';
    import { Chart, registerables } from 'chart.js';
    import { filteredIncome } from "../stores/filteredIncome";

    

    let key = new Date().getTime();

    Chart.register(...registerables);

    let chart;

    onMount(() => {
      filteredIncome.subscribe(($filteredIncome) => {
        if (chart) chart.destroy(); 

        const ctx = document.getElementById('incomeChart').getContext('2d');

        const dataByMonth = new Array(12).fill(0);
        $filteredIncome.forEach(income => {
          const month = new Date(income.Date).getMonth(); // Ensure ISO date is correctly parsed
          dataByMonth[month] += income.Amount; // Accumulate amounts
        });

        chart = new Chart(ctx, {
          type: 'bar', 
          data: {
            labels: ['Januar', 'Februar', 'Mars', 'April', 'Mai', 'Juni', 'Juli', 'August', 'September', 'October', 'November', 'December'],
            datasets: [{
              label: 'Monthly Income',
              data: dataByMonth,
              backgroundColor: 'rgba(54, 162, 235, 0.8)',
              borderColor: 'black',
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            plugins: {
              legend: {
                position: 'top',
              },
              title: {
                display: true,
                text: 'Monthly Income Overview'
              }
            }
          }
        });
      });
    });
</script>

<style>
   canvas {
        border-radius: var(--component-border-radius);

        background: var(--component-bg-color);
        border-style: solid;
        border-color: var(--component-border-color);
        border-width: 1px;
        border-radius: var(--component-border-radius);
        font-family: var(--font-family);
        margin-top: 40px;
        height: auto;
    }

  .chart-container {
        width: 100%;
        max-width: 60rem;
        margin: auto;
        height: 24rem;
        }
</style>

<div class="chart-container">
    <canvas id="incomeChart" height="213" key={key}></canvas>
</div>
