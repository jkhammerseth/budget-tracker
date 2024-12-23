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

  const upcomingExpenses = derived(expenses, $expenses => {
    return $expenses.filter(expense => {
      const expenseDate = new Date(expense.PaymentDate); 
      return expenseDate >= today && 
             expense.Paid !== true &&
             expenseDate >= firstDayOfMonth && 
             expenseDate <= lastDayOfMonth;
    });
  });

  // Calculate total upcoming expenses
  const totalUpcomingExpenses = derived(upcomingExpenses, $upcomingExpenses => {
    return $upcomingExpenses.reduce((acc, expense) => acc + Number(expense.Amount), 0); 
  });
</script>

<style>
.container {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  background-color: var(--component-bg-color);
  box-shadow: var(--component-box-shadow);
  border-radius: var(--component-border-radius);
  border: var(--component-border);
  font-family: var(--font-family);
}

.greeting {
  font-size: 1.4em;
  color: var(--text-color);
  margin-bottom: 20px;
}

.overview {
  font-size: 1em;
  color: var(--text-color-light);
  margin-bottom: 40px;
}

h3 {
  color: var(--text-color);
  margin: 20px 0 10px;
}

p {
  font-size: 1.1em;
  color: var(--text-color);
  font-weight: bold;
  margin-bottom: 20px;
}

ul {
  margin-top: 20px;
}

li {
  background-color: white;
  margin-bottom: 10px;
  padding: 15px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-left: 4px solid var(--primary-color);
  transition: transform 0.2s ease-in-out;
}

li:hover {
  transform: translateY(-2px);
}

li strong {
  color: var(--text-color);
  font-size: 1.1em;
}

.amount {
  color: var(--error-color);
  font-weight: bold;
}

.date {
  color: var(--text-color-light);
}

li:last-child {
  margin-bottom: 0;
}

/* Scrollable list styles */
ul {
  max-height: 300px; /* Adjust based on your preference */
  overflow-y: auto;
}

/* Scrollbar styles */
ul::-webkit-scrollbar {
  width: 6px;
}

ul::-webkit-scrollbar-thumb {
  background: var(--primary-color);
  border-radius: 10px;
}

ul::-webkit-scrollbar-track {
  background: var(--background-color);
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
            <span class="date"> - Due on {new Date(expense.PaymentDate).toLocaleDateString()}</span>
          </div>
        </li>
      {/each}
    {/if}
  </ul>
</div>
