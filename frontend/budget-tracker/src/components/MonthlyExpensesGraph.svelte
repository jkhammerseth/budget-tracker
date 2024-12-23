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
          const month = new Date(expense.PaymentDate).getMonth();
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
        box-shadow: var(--component-box-shadow);
        font-family: var(--font-family);
        margin-top: 40px;
        height: auto;
    }

  .chart-container {
        margin: auto;
        width: 70rem;
        height: 23rem;
        max-width: 100%;
        max-height: 100%;
        }

</style>


<div class="chart-container">
    <canvas id="expensesChart"></canvas>
  </div>
  