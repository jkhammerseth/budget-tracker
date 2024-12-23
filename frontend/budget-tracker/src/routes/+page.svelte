<script>
  import { onMount } from 'svelte';
  import { checkAuthStatus } from '../routes/api/auth';
  import { user } from '../stores/user';
  import ExpensesByCategoryBox from '../components/ExpensesByCategoryBox.svelte';
  import UpcomingExpenses from '../components/UpcomingExpenses.svelte';
  import CalendarBig from '../components/CalendarBig.svelte';


  onMount(() => {
    checkAuthStatus();
  });
</script>

{#if $user.loggedIn}
  <div class="dashboard-container">
    <div class="welcome-header">
      <div>
        <h1 class="greeting">Welcome back, {$user.firstName}</h1>
      </div>
    </div>
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


<style>
  .dashboard-container {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    padding: 1.5rem;
    margin: 2.5rem auto;
    max-width: var(--max-width);
  }

  .welcome-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
  }

  .greeting {
    font-size: 1.875rem;
    font-weight: bold;
  }

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
</style>
