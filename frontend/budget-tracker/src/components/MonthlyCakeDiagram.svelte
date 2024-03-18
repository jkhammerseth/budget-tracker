<script>
    import { onMount } from 'svelte';
    import { Chart, registerables } from 'chart.js';
    import { expenses } from '../stores/expenses'; // Adjust the import path as needed
  
    Chart.register(...registerables);
  
    let chart;
  
    onMount(() => {
      expenses.subscribe(($expenses) => {
        if (chart) chart.destroy(); // Destroy the existing chart before creating a new one
  
        const ctx = document.getElementById('expensesChart').getContext('2d');
  
        // Aggregate expenses by category
        const expensesByCategory = {};
        $expenses.forEach(expense => {
          if (expensesByCategory[expense.Category]) {
            expensesByCategory[expense.Category] += expense.Amount;
          } else {
            expensesByCategory[expense.Category] = expense.Amount;
          }
        });
  
        const categories = Object.keys(expensesByCategory);
        const amounts = categories.map(category => expensesByCategory[category]);
  
        chart = new Chart(ctx, {
          type: 'doughnut',
          data: {
            labels: categories,
            datasets: [{
              label: 'Expenses by Category',
              data: amounts,
              backgroundColor: [
                'rgba(255, 99, 132, 0.2)',
                'rgba(54, 162, 235, 0.2)',
                'rgba(255, 206, 86, 0.2)',
                'rgba(75, 192, 192, 0.2)',
                'rgba(153, 102, 255, 0.2)',
                'rgba(255, 159, 64, 0.2)'
              ],
              borderColor: [
                'rgba(255, 99, 132, 1)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(255, 159, 64, 1)'
              ],
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            plugins: {
              legend: {
                position: 'right',
              },
              title: {
                display: true,
                text: 'Expenses by Category'
              }
            }
          }
        });
      });
    });
  </script>
  
  <style>
    canvas {
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      border-radius: 10px;
      background: #f8f9fa;
      border-style: solid;
      border-color: black;
      border-width: 1px;

      
    }
  </style>
  
  <canvas id="expensesChart"></canvas>
  