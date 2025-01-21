<script>
  import { categories } from '../../stores/categories'; // Import the categories store
  import { getIconComponent } from '../../utility/icons';
  import { onMount } from 'svelte';
  import { derived } from 'svelte/store';
  import { filteredExpenses } from '../../stores/filteredExpenses';
  import { getCategories } from '../api/fetchCategories';
  import { formatAmount } from '../../utility/functions';
  import SummaryCard from '../../components/SummaryCard.svelte';

  // Fetch categories on mount
  onMount(getCategories);

  // Derive total spent by category
  const expensesByCategory = derived(filteredExpenses, ($filteredExpenses) => {
    const categoryTotals = new Map();

    $filteredExpenses.forEach(({ category, amount }) => {
      if (categoryTotals.has(category.id)) {
        categoryTotals.set(category.id, categoryTotals.get(category.id) + amount);
      } else {
        categoryTotals.set(category.id, amount);
      }
    });

    return categoryTotals;
  });

  // Calculate remaining budget
  function calculateRemaining(category, totalSpent) {
    return category.budget ? category.budget - totalSpent : 'Set a budget';
  }

  const totalExpenses = derived(filteredExpenses, $filteredExpenses => {
    return $filteredExpenses.reduce((total, expense) => total + expense.amount, 0);
  });

   // Sample summary data
   const summaryData = {
    totalBudget: 5000,
    totalSpent: $totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses),
    overBudgetCategories: 2,
    nearLimitCategories: 1,
  };
</script>


<div class="category-container">
<div class="summary-cards">
  <SummaryCard title="Total Budget" value={`$${summaryData.totalBudget}`}/>
  <SummaryCard title="Total Spent" value={`${summaryData.totalSpent}`}/>
  <SummaryCard title="Over Budget" value={summaryData.overBudgetCategories}/>
  <SummaryCard title="Near Limit" value={summaryData.nearLimitCategories}/>
</div>


  {#each $categories as category}
    <div class="category-card">
      <!-- Top Row: Icon and Name -->
      <div class="top-row">
        <div class="category-icon">
          {#if getIconComponent(category.name)}
            <svelte:component this={getIconComponent(category.name)} size="24" />
          {:else}
            📁 <!-- Fallback icon -->
          {/if}
        </div>
        <div class="category-name">{category.name}</div>
      </div>

      <!-- Budget Input -->
      <div class="budget-section">
        <label for="budget-{category.id}">Set Budget:</label>
        <input
          id="budget-{category.id}"
          type="number"
          class="budget-input"
          bind:value={category.budget}
          placeholder="Enter amount"
        />
      </div>

      <!-- Spent -->
      {#if $expensesByCategory.has(category.id)}
        <div class="category-sum">
          Spent: {formatAmount($expensesByCategory.get(category.id))}
        </div>
      {:else}
        <div class="category-sum">Spent: 0 kr</div>
      {/if}

      <!-- Progress Bar -->
      {#if category.budget}
        <div class="progress-bar-container">
          <div
            class="progress-bar"
            style="width: {Math.min(
              ($expensesByCategory.get(category.id) || 0) / category.budget * 100,
              100
            )}%;"
          ></div>
        </div>
      {/if}

      <!-- Remaining -->
      <div class="remaining-amount">
        Remaining: {category.budget ? formatAmount(calculateRemaining(category, $expensesByCategory.get(category.id) || 0)) : 'Set a budget'}
      </div>
    </div>
  {/each}
</div>

<style>
  .category-container {
    display: flex;
    flex-direction: column;
    gap: 16px;
    max-width: 75%;
    margin: 4rem auto 0;
    margin-left: 2rem;
  }

  .category-card {
    display: flex;
    flex-direction: column;
    border-radius: 8px;
    padding: 16px;
    background-color: #fff;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }

  .summary-cards {
    display: grid;
    background-color: #fff;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 20px;
  }

  .top-row {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 16px;
  }

  .category-icon {
    font-size: 24px;
    padding: 10px;
    border-radius: 50%;
    background-color: #f1f1f1;
  }

  .category-name {
    font-size: 16px;
    font-weight: bold;
    color: #333;
  }

  .budget-section {
    margin-bottom: 16px;
  }

  .budget-section label {
    display: block;
    font-size: 14px;
    color: #666;
    margin-bottom: 4px;
  }

  .budget-input {
    width: 100%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 8px;
    font-size: 14px;
  }

  .category-sum {
    font-size: 14px;
    margin: 8px 0;
    color: #555;
  }

  .progress-bar-container {
    width: 100%;
    background-color: #eee;
    height: 8px;
    border-radius: 4px;
    overflow: hidden;
    margin: 8px 0;
  }

  .progress-bar {
    height: 100%;
    background-color: var(--primary-button-color);
    transition: width 0.3s ease;
  }

  .remaining-amount {
    font-size: 14px;
    font-weight: bold;
    color: #333;
    margin-top: 8px;
  }
</style>
