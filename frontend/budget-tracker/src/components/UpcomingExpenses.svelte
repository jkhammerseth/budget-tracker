<script>
  import { onMount } from "svelte";
  import { formatAmount, formatExpenseAmount, fromISOString, daysFromNow } from "../utility/functions";
  import { derived } from "svelte/store";
  import { FetchExpenses } from "../routes/api/fetchExpenses";
  import { expenses } from '../stores/expenses'
  import { Calendar } from 'lucide-svelte';
  import ExpenseModal from './modals/ExpenseModal.svelte';
  import { activeModal } from '../stores/activeModal';
    import { getIconComponent } from '../utility/icons'


  onMount(FetchExpenses);

  let selectedExpense = null;

  const today = new Date();
  const twoWeeksFromToday = new Date();
  twoWeeksFromToday.setDate(today.getDate() + 14);

  const upcomingExpenses = derived(expenses, $expenses => {
    return $expenses.filter(expense => {
      const expenseDate = new Date(expense.payment_date);
      return expenseDate >= today && expenseDate <= twoWeeksFromToday;
    });
  });

  function openExpenseModal(expense) {
    
    selectedExpense = expense;
    activeModal.set('expense'); // Make sure 'expense' is set to trigger the modal
  }
// on:click={() => openExpenseModal(payment)}
</script>

<div class="upcoming-payments">
  <div class="section-header">
    <h2>Upcoming Payments</h2>
    <span class="icon"><Calendar size="24" /></span>
  </div>
  <div class="payments-list">
    {#if $upcomingExpenses.length > 0}
      {#each $upcomingExpenses as payment}
      <div class="payment">
        <div class="payment-item">
          <div class="category-icon">
            {#if getIconComponent(payment.category.name)}
              <svelte:component this={getIconComponent(payment.category.name)} size="24" />
            {:else}
              📁 <!-- Fallback icon -->
            {/if}
          </div>
          <div class="payment-details">
            <p class="payment-name">{payment.name}</p>
            <p class="payment-date">Due {daysFromNow(payment.payment_date)}</p>
          </div>
        </div>
        <p class="payment-amount">{formatExpenseAmount(payment.amount)}</p>
      </div>
    {/each}
      <div class="view">
        <a class="view-link" href="/expenses" aria-label="View all upcoming payments">View All</a>
      </div>
    {:else}
      <p>No upcoming payments in the next two weeks.</p>
    {/if}
  </div>
</div>


{#if $activeModal === 'expense' && selectedExpense}
<ExpenseModal expense={selectedExpense} />
{/if}

<style>
  .upcoming-payments {
    background: #fff;
    padding: 16px;
    border-radius: 8px;
  }

  .upcoming-payments h2 {
    margin-bottom: 16px;
  }

  .payments-list {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .payment {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px;
    border-radius: 8px;
    background-color: #f9f9f9;
    transition: background-color 0.2s ease;
  }

  .payment:hover {
    background-color: #e8f0fe;
  }

  .payment-item {
    display: flex;
    gap: 10px;
    align-items: center;
    flex: 1;
  }

  .payment-details {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .payment-name {
    font-weight: 600;
    margin: 0;
    text-transform: capitalize;
  }

  .payment-date {
    font-size: 0.9rem;
    color: #666;
    margin: 4px 0 0;
  }

  .payment-amount {
    font-weight: 700;
    font-size: 1.2rem;
    align-items: flex-end;
  }

  .category-icon {
    padding: 10px;
    border-radius: 50%;
    background-color: #f1f1f1;
  }

  .view {
    display: flex;
    justify-content: center;
  }

  .view-link {
    text-decoration: none;
    font-size: 1.1rem;
    padding: 16px;
  }

  .view-link:hover {
    color: #4a86ff;
    text-decoration: none;
  }

  .section-header {
    display: flex;
    justify-content: space-between;
  }

  .section-header h2 {
    margin: 0;
  }

  .section-header .icon {
    display: flex;
    align-items: center;
  }

  .icon {
    padding: 16px;
  }
</style>
