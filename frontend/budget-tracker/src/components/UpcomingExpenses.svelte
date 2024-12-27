<script>
  import { onMount } from "svelte";
  import { formatAmount, formatExpenseAmount, fromISOString, daysFromNow } from "../utility/functions";
  import { derived } from "svelte/store";
  import { FetchExpenses } from "../routes/api/fetchExpenses";
  import { expenses } from '../stores/expenses'
  import { Calendar } from 'lucide-svelte';


  onMount(FetchExpenses);

  const today = new Date();
  const twoWeeksFromToday = new Date();
  twoWeeksFromToday.setDate(today.getDate() + 14);

  const upcomingExpenses = derived(expenses, $expenses => {
    return $expenses.filter(expense => {
      const expenseDate = new Date(expense.PaymentDate);
      return expenseDate >= today && expenseDate <= twoWeeksFromToday;
    });
  });

</script>

<div class="upcoming-payments">
  <div class="upcoming-header">
  <h2>Upcoming Payments</h2>
  <span class="calendar-icon"><Calendar size="24" /></span>
</div>
  {#if $upcomingExpenses.length > 0}
      {#each $upcomingExpenses as payment}
      <div class="payment">
          <div>
              <p class="payment-name">{payment.Name}</p>
              <p>Due {daysFromNow(payment.PaymentDate)}</p>
          </div>
          <div>
              <p class="payment-amount">{formatExpenseAmount(payment.Amount)}</p>
              <span>{payment.Paid ? "Paid" : "Upcoming"}</span>
          </div>
          
      </div>
      {/each}
      <div class="view">
        <a class="view-link" href="/expenses" aria-label="View all expenses">View All</a>
      </div>
  {:else}
  <p>No upcoming payments in the next two weeks.</p>
  {/if}
</div>

<style>
  .upcoming-payments {
      background: #fff;
      padding: 16px;
      border-radius: 8px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .upcoming-header {
    display: flex;
    justify-content: space-between;
  }

  .upcoming-header h2 {
  margin: 0;
  }

  .upcoming-header .calendar-icon {
    display: flex;
    align-items: center;
  }

  .payment {
    display: flex;
    justify-content: space-between; /* Space between name and amount */
    align-items: center; /* Align items vertically */
    border-radius: 8px;
    padding: 8px 16px; /* Add more padding for better spacing */
    margin-bottom: 8px; /* Spacing between payment items */
    background-color: #f9f9f9;
    transition: background-color 0.2s ease;
  }

  .payment:hover {
    background-color: #e8f0fe; /* Subtle blue on hover */
  }

  .payment div {
    display: flex;
    flex-direction: column; /* Align text within each section */
  }

  .payment-name {
    text-transform: capitalize;
    font-weight: 600;
    font-size: 1.2rem; /* Adjust size */
    margin: 0;
    margin-top: 1rem;
  }

  .payment-amount {
    font-weight: 700;
    font-size: 1.2rem;
    text-align: right; /* Align amount text to the right */
    margin: 0;
  }

  .calendar-icon {
    padding: 16px;
  }

  .payment:hover {
    background-color: #f1f1f1;
    cursor: pointer;
  }

  .payment:last-child {
      margin-bottom: 0;
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
</style>