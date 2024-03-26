<script>
  import { onMount } from 'svelte';
  import { Chart, registerables } from 'chart.js';
  import { filteredExpenses } from '../stores/filteredExpenses';

  let chart;

  // This function setups the chart, intended to be executed on the client side
  const setupChart = ($filteredExpenses) => {
      const ctx = document.getElementById('expensesChart').getContext('2d');
      
      // Sum expenses by month
      const dataByMonth = new Array(12).fill(0);
      $filteredExpenses.forEach(expense => {
          const month = new Date(expense.Date).getMonth();
          dataByMonth[month] += expense.Amount;
      });

      chart = new Chart(ctx, {
          type: 'bar',
          data: {
              labels: ['Januar', 'Februar', 'Mars', 'April', 'Mai', 'Juni', 'Juli', 'August', 'September', 'October', 'November', 'Desember'],
              datasets: [{
                  label: 'Monthly Expenses',
                  data: dataByMonth,
                  backgroundColor: 'rgba(219, 38, 0, 0.8)',
                  borderColor: 'rgba(219, 162, 235, 1)',
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
                      text: 'Monthly Expenses Overview'
                  }
              }
          }
      });
  };

  onMount(() => {
      // Only register Chart.js and setup chart on the client side
      if (typeof window !== 'undefined') {
          Chart.register(...registerables);

          // Use a subscription to filteredExpenses store
          const unsubscribe = filteredExpenses.subscribe(($filteredExpenses) => {
              if (chart) {
                  chart.destroy(); // Destroy the existing chart before creating a new one
              }
              setupChart($filteredExpenses);
          });

          // Cleanup the subscription and chart instance when component is destroyed
          return () => {
              unsubscribe();
              if (chart) {
                  chart.destroy();
              }
          };
      }
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
        max-width: 45rem; /* or any other max-width */
        margin: auto;
        height: 20rem;
        }

</style>


<div class="chart-container">
    <canvas id="expensesChart"></canvas>
  </div>
  