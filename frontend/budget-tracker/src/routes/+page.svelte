<script>
  import { onMount } from 'svelte';
  import { checkAuthStatus } from '../routes/api/auth';
  import { user } from '../stores/user';
  import ExpensesByCategoryBox from '../components/ExpensesByCategoryBox.svelte';
  import UpcomingExpenses from '../components/UpcomingExpenses.svelte';
  import Calendar from '../components/Calendar.svelte';
  import TotalBalance from '../components/TotalBalance.svelte';
  import CalendarBig from '../components/CalendarBig.svelte';

  onMount(() => {
    checkAuthStatus();
  });
</script>

<style>

  .welcome {
    max-width: 400px;
    padding: 50px;
    background: white;
    border-radius: 8px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    gap: 20px;
    text-align: center;
    margin: 200px auto;
    align-items: center;
  }
  .title {
    font-size: 2em;
    margin-bottom: 20px;
  }

  .description {
    font-size: 1.2em;
    margin-bottom: 20px;
  }

  .button {
    padding: 10px 20px;
    margin-right: 10px;
    font-size: 1em;
    background-color: #3A87F2;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    width: 150px;
  }

  .dashboard-container {
    display: grid;
    grid-template-columns: 3fr 1fr 1fr; /* Adjust based on your content and preference */
    grid-template-rows: auto;
    gap: 20px;
    padding: 20px;
    margin: 40px auto;
    max-width: var(--max-width);
  }

  .greeting-container, .balance, .expense-by-category-container, .upcoming-expenses-container {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
  }

  .greeting-container {
    grid-column: 1 / 4;
    font-size: 1.5em;
    color: #333;

  }

  .expense-by-category-container{
    grid-column: 3 / 6;
    grid-row: 2;
  }
  
  .upcoming-expenses-container {
    grid-column: 2 / 3;
    grid-row: 2;

  }

  .calendar-container {
    grid-column: 5 / 6;
    grid-row: 1;
  }

  .balance {
    grid-column: 1/3;
    grid-row: 2;
    background-color: #E8EFF1;
    border: black 1px solid;
    margin-top: 20px;
  }

  /* Add responsive adjustments */
  @media (max-width: 768px) {
    .dashboard-container {
      grid-template-columns: 1fr;
      grid-template-rows: auto;
    }

    .expense-by-category-container, .upcoming-expenses-container, .calendar-container {
      grid-column: 1 / -1;
    }
  }
</style>

<div class="dashboard-container">
  {#if $user.loggedIn}
  <div class="greeting-container">
      
      <div class="upcoming-expenses-container">
        <UpcomingExpenses />
      </div>
    
  </div>
    <div class="balance">
      <CalendarBig />
    </div>
    <div class="calendar-container">
      <Calendar />
    </div>
    <div class="expense-by-category-container">
      <ExpensesByCategoryBox />
    </div>
    
  {:else}
  <div class="welcome">
    <div class="title">Welcome to Budget Tracker</div>
    <p class="description">
      Take control of your finances with our easy-to-use budgeting tool. Track your expenses, set budgets, and achieve your financial goals.
    </p>
    <button class="button" on:click={() => (window.location.href = '/login')}>Log In</button>
    <button class="button" on:click={() => (window.location.href = '/signup')}>Sign Up</button>
  </div>
  {/if}
</div>
