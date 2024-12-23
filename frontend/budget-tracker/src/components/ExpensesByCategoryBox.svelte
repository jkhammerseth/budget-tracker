<script>
  import { filteredExpenses } from '../stores/filteredExpenses';
  import { derived } from 'svelte/store';
  import { FetchExpenses } from '../routes/api/fetchExpenses';
  import { onMount } from 'svelte';
  import { formatAmount } from '../utility/functions'

  onMount(FetchExpenses);

  const categoryIcons = {
    Food: '🍔',
    Transportation: '🚗',
    Housing: '🏠',
    Utilities: '💡',
    Insurance: '📋',
    Healthcare: '💊',
    Entertainment: '🎉',
    Clothing: '👗',
    Miscellaneous: '📦',
  };

  function getCategoryIcon(category) {
    return categoryIcons[category] || '📁'; // Default icon if category not found
  }

  const expensesByCategory = derived(filteredExpenses, $filteredExpenses => {
    const categoryMap = new Map();

    // Aggregate totals for each category
    $filteredExpenses.forEach(({ Category, Amount }) => {
      if (categoryMap.has(Category)) {
        categoryMap.set(Category, categoryMap.get(Category) + Amount);
      } else {
        categoryMap.set(Category, Amount);
      }
    });

    // Convert map to array and sort by TotalAmount (descending order)
    return Array.from(categoryMap, ([Category, TotalAmount]) => ({
      Category,
      TotalAmount,
    })).sort((a, b) => b.TotalAmount - a.TotalAmount);
  });

  function openAddCategoryModal() {
    console.log('Open Add Category Modal');
    // Add modal functionality later
  }
</script>

<div class="category-container">
  <div class="container-header">
    <h2>Categories</h2>
    <button class="add-category-button" on:click={openAddCategoryModal}>+</button>
  </div>
  <div class="category-list">
    {#each $expensesByCategory as { Category, TotalAmount }}
    <div class="category-item">
      <span class="category-icon">{getCategoryIcon(Category)}</span>
      <div class="category-details">
        <span class="category-name">{Category}</span>
        <span class="category-amount">{formatAmount(TotalAmount)}</span>
      </div>
    </div>
    {/each}
  </div>
</div>

<style>
  .category-container {
    max-width: 600px;
    margin-top: 100px;
    padding: 20px;
    font-family: 'Roboto', sans-serif;
  }

  .container-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }

  h2 {
    font-size: 1.5rem;
    color: #333;
    font-weight: 700;
    margin: 0;
  }

  .add-category-button {
    font-size: 1.5rem;
    font-weight: bold;
    border: none;
    border-radius: 50%;
    width: 36px;
    height: 36px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  .category-list {
    max-height: calc(4 * 70px); /* Limit to 4 items (assuming each item is ~70px tall) */
    overflow-y: auto;
    scrollbar-width: none; /* For Firefox */
    -ms-overflow-style: none; /* For Internet Explorer and Edge */
  }

  .category-list::-webkit-scrollbar {
    display: none; /* For Chrome, Safari, and Opera */
  }

  .category-item {
    display: flex;
    align-items: center;
    border-radius: 8px;
    padding: 10px 15px;
    margin-bottom: 10px;
    background-color: #f9f9f9;
    transition: background-color 0.2s ease;
  }

  .category-item:hover {
    background-color: #f1f1f1;
  }

  .category-icon {
    font-size: 2rem;
    margin-right: 15px;
  }

  .category-details {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .category-name {
    font-size: 1rem;
    font-weight: bold;
    color: #333;
  }

  .category-amount {
    font-size: 0.85rem;
    color: #777;
  }
</style>
