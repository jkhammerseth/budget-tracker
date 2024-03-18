<script>
    import { onMount } from 'svelte';
    import { Chart, registerables } from 'chart.js';
    import { filteredIncome } from '../stores/filteredincome';

    

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
            labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
            datasets: [{
              label: 'Monthly Income',
              data: dataByMonth,
              backgroundColor: 'rgba(54, 162, 235, 0.8)',
              borderColor: 'rgba(54, 162, 235, 1)',
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

<canvas id="incomeChart" height="213" key={key}></canvas>
