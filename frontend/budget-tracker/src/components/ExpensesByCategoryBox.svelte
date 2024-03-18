<script>
  import { filteredExpenses } from '../stores/filteredExpenses';
  import { derived } from 'svelte/store';
  import { FetchExpenses } from '../routes/api/fetchExpenses';
  import { onMount } from 'svelte';

  onMount(FetchExpenses);

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


  function getCategoryColor(category) {
    return categoryColors[category] || '#757575'; // Default color if category not found
  }

  let totalExpenses = derived(filteredExpenses, $filteredExpenses =>
    $filteredExpenses.reduce((total, expense) => total + expense.Amount, 0)
  );

  const expensesByCategory = derived([filteredExpenses, totalExpenses], ([$filteredExpenses, $totalExpenses]) => {
    const categoryMap = new Map();

    $filteredExpenses.forEach(({ Category, Amount }) => {
      if (categoryMap.has(Category)) {
        categoryMap.set(Category, categoryMap.get(Category) + Amount);
      } else {
        categoryMap.set(Category, Amount);
      }
    });

    return Array.from(categoryMap, ([Category, TotalAmount]) => ({
      Category,
      TotalAmount,
      Percentage: $totalExpenses ? (TotalAmount / $totalExpenses * 100) : 0,
    }));
  });

  function formatAmount(amount) {
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'NOK' }).format(amount);
  }
</script>

<style>
  .category-container {
    max-width: 600px;
    border: black 1px solid;
    background-color: #E8EFF1;
    padding: 20px;
    border-radius: 10px;
    font-family: 'Roboto', sans-serif;
  }

  .category-expense {
    display: flex;
    justify-content: space-between;
    background-color: #e0e0e0;
    border-radius: 5px;
    margin-bottom: 10px;
    overflow: hidden;
  }

  .progress-bar {
    background-color: #4CAF50;
    height: 40px;
    border-radius: 5px 0 0 5px;
    text-align: right;
    padding-right: 10px;
    color: #E8EFF1;
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }

  .total-expenses-box {
    font-family: 'Roboto', sans-serif;
    display: flex;
    flex-direction: column;
  }

  .total-expenses-box h1,
  .total-expenses-box h2,
  .category-container h2 {
    margin: 0;
    padding: 5px 0;
    font-size: 1.5rem;
    color: #333;
    font-weight: 700;
  }

  .category-container h2 {
    margin-bottom: 20px;
  }

  .category-expense span {
    padding: 10px;
  }

  .category-expense span:nth-child(2) {
    flex: 1;
    text-align: center;
  }

  .category-expense span:nth-child(3) {
    padding-right: 10px;
  }

  .category-details {
    display: flex;
    width: 100%;
    justify-content: space-between;
    padding: 10px;
    background-color: #f8f9fa;
  }

  .progress-bar-container {
    width: 100%;
    background-color: #e0e0e0;
    border-radius: 5px;
    overflow: hidden;
    margin-top: 5px;
  }

  .progress-bar {
    height: 10px;
    text-align: right;
    color: white;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    padding-right: 10px;
  }

  .progress-bar-container {
    width: 100%;
    background-color: #e0e0e0; 
    border-radius: 5px;
    overflow: hidden;
    margin-top: 5px;
  }

</style>



<div class="category-container">
  <div class="total-expenses-box">
    <h1>Total Expenses:</h1>
    <h2>{formatAmount($totalExpenses)}</h2>
  </div>
  <hr>
  <h2>Expenses by Category</h2>
  {#each $expensesByCategory as { Category, TotalAmount, Percentage }}
  <div class="progress-bar-container">
    <div class="progress-bar" style="width: {Percentage}%; background-color: {getCategoryColor(Category)};">
    </div>
  </div>
    <div class="category-expense">
      <div class="category-details">
        <span>{Category}</span>
        <span>{formatAmount(TotalAmount)}</span>
        <span>{Percentage.toFixed(2)}%</span>
      </div>
    </div> 
  {/each}
</div>