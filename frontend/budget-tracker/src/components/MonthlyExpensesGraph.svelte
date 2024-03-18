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
              labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
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
      border-radius: 5px;
      padding: 5px;
      background: #E8EFF1;
      border-style: solid;
      border-color: black;
      border-width: 1px;
      border-radius: 10px;
      font-family: 'Roboto', sans-serif;
      margin-top: 40px;
  }
</style>

<canvas id="expensesChart" height="213"></canvas>
