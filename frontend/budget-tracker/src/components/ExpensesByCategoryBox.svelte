<script>
  import { categories } from '../stores/categories';
  import { derived } from 'svelte/store';
  import { filteredExpenses } from '../stores/filteredExpenses';
  import { onMount } from 'svelte';
  import { FetchExpenses } from '../routes/api/fetchExpenses';
  import { formatAmount } from '../utility/functions';
  import { getIconComponent } from '../utility/icons'
  import { checkAuthStatus } from '../routes/api/auth';
  import { activeModal } from "../stores/activeModal";
  import CategoryModal from "../components/modals/CategoryModal.svelte"
  import { getCategoryById } from "../routes/api/fetchCategories"

  let selectedCategory = null;

  onMount(() => {
    checkAuthStatus();
  });
    

  const expensesByCategory = derived(filteredExpenses, ($filteredExpenses) => {
    const categoryMap = new Map();

    // Aggregate totals for each category
    $filteredExpenses.forEach(({ category, amount }) => {
      const categoryId = category.id;
      const categoryName = category.name;
      const categoryIcon = category.icon;
      
      if (categoryMap.has(categoryId)) {
        const existing = categoryMap.get(categoryId);
        categoryMap.set(categoryId, {
          ...existing,
          totalAmount: existing.totalAmount + amount
        });
      } else {
        categoryMap.set(categoryId, {
          id: categoryId,
          name: categoryName,
          icon: categoryIcon,
          totalAmount: amount
        });
      }
    });

    // Convert map to array and sort by totalAmount
    return Array.from(categoryMap.values())
      .sort((a, b) => b.totalAmount - a.totalAmount);
  });

  function openAddCategoryModal() {
    console.log('Open Add Category Modal');
    // Add modal functionality later
  }

  async function openCategoryModal(category) {
    try {
        const categoryData = await getCategoryById(category);
        selectedCategory = categoryData;
        activeModal.set("category");
    } catch (error) {
        console.error("Error fetching category data:", error);
        alert("Failed to load category details.");
    }
}

</script>

<div class="category-container">
  <div class="container-header">
    <h2>Categories</h2>
    <button class="add-category-button" on:click={openAddCategoryModal}>+</button>
  </div>
  <div class="category-list">
    {#each $expensesByCategory as category}
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <!-- svelte-ignore a11y-no-static-element-interactions -->
      <div class="category-item" on:click={openCategoryModal(category)}>
        <span class="category-icon">
          {#if getIconComponent(category.name)}
            <svelte:component this={getIconComponent(category.name)} size="24" />
          {:else}
            📁 <!-- Fallback icon -->
          {/if}
        </span>
        <div class="category-details">
          <span class="category-name">{category.name}</span>
          <span class="category-amount">{formatAmount(category.totalAmount)}</span>
        </div>
      </div>
    {/each}
  </div>
</div>

{#if $activeModal === 'category' && selectedCategory}
  <CategoryModal category={selectedCategory} />
{/if}

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
    max-height: calc(4 * 70px);
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
    font-size: 1.4rem;
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
