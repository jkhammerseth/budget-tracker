<script>
    import { writable } from 'svelte/store';
    import AddExpenseModal from './modals/AddExpenseModal.svelte';  // Import the modal component
    import AddIncomeModal from './modals/AddIncomeModal.svelte';
    import { activeModal } from "../stores/activeModal";  // Import the activeModal store

    const isFabMenuOpen = writable(false);

    function toggleFabMenu() {
        isFabMenuOpen.update((open) => !open);
    }

    function openAddExpenseModal() {
        activeModal.set("addExpense");  // Trigger the modal to open
    }

    function openAddIncomeModal() {
      activeModal.set("addIncome");
    }


    function handleAddOther() {
        console.log('Add Other action triggered');
    }
</script>

<div class="fab-container">
    <!-- FAB Menu -->
    {#if $isFabMenuOpen}
        <div class="fab-menu">
            <button class="fab-menu-item" on:click={openAddIncomeModal}>Add Income</button>
            <button class="fab-menu-item" on:click={openAddExpenseModal}>Add Expense</button>
            <button class="fab-menu-item" on:click={handleAddOther}>Other</button>
        </div>
    {/if}

    <!-- Main FAB -->
    <button class="fab-button" on:click={toggleFabMenu} aria-label="Toggle FAB Menu">
        <span class="fab-icon" aria-hidden="true">+</span>
    </button>
</div>

<!-- Conditionally render the modal when the activeModal is set to 'addExpense' -->
{#if $activeModal === 'addExpense'}
    <AddExpenseModal />
{/if}

{#if $activeModal === 'addIncome'}
    <AddIncomeModal />
{/if}
  
  <style>
    .fab-container {
      position: fixed;
      margin-right: 25rem;
      bottom: 16px;
      right: 16px;
      display: flex;
      flex-direction: column;
      align-items: center;
      z-index: 1000;
    }
  
    .fab-button {
      background-color: #007bff;
      color: white;
      border: none;
      border-radius: 50%;
      width: 56px;
      height: 56px;
      font-size: 24px;
      display: flex;
      align-items: center;
      justify-content: center;
      box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.2);
      cursor: pointer;
      transition: transform 0.2s ease, box-shadow 0.2s ease;
    }
  
    .fab-button:hover {
      transform: scale(1.1);
      box-shadow: 0px 6px 8px rgba(0, 0, 0, 0.3);
    }
  
    .fab-icon {
      display: inline-block;
      transition: transform 0.2s ease;
    }
  
    /* Rotate '+' to '×' when menu is open */
    :global(.fab-container .fab-button.open .fab-icon) {
      transform: rotate(45deg);
    }
  
    .fab-menu {
      position: absolute;
      bottom: 64px; /* Positioned above the FAB button */
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 8px;
    }
  
    .fab-menu-item {
      background-color: white;
      color: #007bff;
      border: 1px solid #007bff;
      border-radius: 8px;
      padding: 8px 12px;
      font-size: 14px;
      cursor: pointer;
      transition: background-color 0.2s ease, color 0.2s ease;
      box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
    }
  
    .fab-menu-item:hover {
      background-color: #007bff;
      color: white;
    }
  </style>
  