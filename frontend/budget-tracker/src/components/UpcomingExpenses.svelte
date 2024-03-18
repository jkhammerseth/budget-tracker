<script>
  import { onMount } from 'svelte';
  import { expenses } from '../stores/expenses';
  import { FetchExpenses } from '../routes/api/fetchExpenses';
  import { derived } from 'svelte/store';
  import { user } from '../stores/user';

  onMount(FetchExpenses);

  const today = new Date();
  const firstDayOfMonth = new Date(today.getFullYear(), today.getMonth(), 1);
  const lastDayOfMonth = new Date(today.getFullYear(), today.getMonth() + 1, 0);

  // Use derived store to automatically update upcomingExpenses based on expenses store
  const upcomingExpenses = derived(expenses, $expenses => {
    return $expenses.filter(expense => {
      const expenseDate = new Date(expense.Date); // Assuming the property is 'Date'
      return expenseDate >= today && 
             expense.Paid !== true &&
             expenseDate >= firstDayOfMonth && 
             expenseDate <= lastDayOfMonth;
    });
  });

  // Calculate total upcoming expenses
  const totalUpcomingExpenses = derived(upcomingExpenses, $upcomingExpenses => {
    return $upcomingExpenses.reduce((acc, expense) => acc + Number(expense.Amount), 0); // Assuming the property is 'Amount'
  });
</script>

<style>
  .container {
    width: full-width;
    padding: 20px;
    background-color: #E8EFF1;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    font-family: 'Roboto', sans-serif;
    border: 1px solid black;
  }

  .greeting {
    font-size: 1.2em;
    color: #333;
    margin-bottom: 30px;
    
  }

  h3 {
    color: #333;
    margin-top: 20px;
    margin-bottom: 10px;
  }

  .overview {
    font-size: 0.95em;
    color: #666;
    margin-bottom: 80px;

  }

  p {
    font-size: 1em;
    color: #333;
    font-weight: bold;
  }

  ul {
    list-style: none;
    padding: 0;
    margin-top: 10px;
  }

  li {
    background-color: white;
    margin-bottom: 10px;
    padding: 15px;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  li strong {
    color: #333;
    font-weight: 500;
  }

  .amount {
    color: #e74c3c;
    font-weight: bold;
  }

  .date {
    color: #7f8c8d;
  }
</style>

<div class="container">
  <div class="greeting">
    Welcome back, {$user.firstName}!
  </div>
  <div class="overview">
    Here is an overview of your finances for the month.
  </div>
  <h3>Scheduled Transactions This Month</h3>
  <p>Total Upcoming Expenses: {$totalUpcomingExpenses} kr</p>
  <ul>
    {#if $upcomingExpenses.length === 0}
      <li>No more expenses this month!</li>
    {:else}
      {#each $upcomingExpenses as expense}
        <li>
          <strong>{expense.Name}</strong>
          <div>
            <span class="amount">{expense.Amount} kr</span>
            <span class="date"> - Due on {new Date(expense.Date).toLocaleDateString()}</span>
          </div>
        </li>
      {/each}
    {/if}
  </ul>
</div>
