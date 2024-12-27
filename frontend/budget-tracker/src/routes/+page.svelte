<script>
  import { Bell, Settings, User, TrendingUp, Wallet, Calendar, ArrowRight } from 'lucide-svelte';
  import { user } from '../stores/user';
  import { goto } from '$app/navigation';
  import { filteredExpenses } from '../stores/filteredExpenses';
  import { filteredIncome } from '../stores/filteredIncome';
  import { formatAmount, fromISOString } from '../utility/functions';
  import { onMount } from 'svelte';
  import { derived } from "svelte/store";
  import { fetchIncome } from "../routes/api/fetchIncome";
  import { FetchExpenses } from "../routes/api/fetchExpenses"
  import { expenses } from '../stores/expenses';

  onMount(() => {
        FetchExpenses();
        fetchIncome();
    });


  const totalIncome = derived(filteredIncome, $filteredIncome => {
    return $filteredIncome.reduce((total, income) => total + income.Amount, 0);
  });

  const totalExpenses = derived(filteredExpenses, $filteredExpenses => {
    return $filteredExpenses.reduce((total, expense) => total + expense.Amount, 0);
  });

  const totalBalance = derived(
  [totalIncome, totalExpenses],
  ([$totalIncome, $totalExpenses]) => {
    return $totalIncome - $totalExpenses;
  }
  );

  const today = new Date();
  const twoWeeksFromToday = new Date();
  twoWeeksFromToday.setDate(today.getDate() + 14);

  const upcomingExpenses = derived(expenses, $expenses => {
    return $expenses.filter(expense => {
      const expenseDate = new Date(expense.PaymentDate);
      return expenseDate >= today && expenseDate <= twoWeeksFromToday;
    });
  });

  // Calculate total upcoming expenses
  const totalUpcomingExpenses = derived(upcomingExpenses, $upcomingExpenses => {
    return $upcomingExpenses.reduce((acc, expense) => acc + Number(expense.Amount), 0); 
  });

  let todayDate = new Date().toLocaleDateString('en-US', { 
    weekday: 'long', 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  });
</script>

<div class="landing-page">
  <!-- Header -->
  <div class="header">
    <div class="header-title">Welcome back, { $user.firstName }</div>
   
  </div>

  <!-- Date and Quick Actions -->
  <div class="quick-actions">
    <div class="date">{todayDate}</div>
  </div>

  <!-- Financial Snapshot -->
  <div class="financial-snapshot">
    <h2>Quick Overview</h2>
    <div class="snapshot-grid">
      <div class="snapshot-item available-balance">
        <div class="snapshot-title">
          <span>Available Balance</span>
          <Wallet class="icon" />
        </div>
        <div class="snapshot-value">{formatAmount($totalBalance)}</div>
      </div>
      <div class="snapshot-item monthly-savings">
        <div class="snapshot-title">
          <span>Monthly Savings</span>
          <TrendingUp class="icon" />
        </div>
        <div class="snapshot-value">+4,500 kr</div>
      </div>
      <div class="snapshot-item upcoming-bills">
        <div class="snapshot-title">
          <span>Upcoming Bills</span>
          <Calendar class="icon" />
        </div>
        <div class="snapshot-value">10,799 kr</div>
      </div>
    </div>
  </div>

  <!-- Quick Links -->
  <div class="quick-links">
    <div class="quick-link">
      <h3>Dashboard</h3>
      <p>View your complete financial overview</p>
      <ArrowRight class="icon" />
    </div>
    <!-- svelte-ignore missing-declaration -->
    <div class="quick-link">
      <h3>Track Expenses</h3>
      <p>Add and manage your expenses</p>
      <ArrowRight class="icon" />
    </div>
    <div class="quick-link">
      <h3>Manage Income</h3>
      <p>Track your income sources</p>
      <ArrowRight class="icon" />
    </div>
    <div class="quick-link">
      <h3>Calendar</h3>
      <p>View scheduled transactions</p>
      <ArrowRight class="icon" />
    </div>
  </div>

  <!-- Important Updates -->
  <div class="updates">
    <div class="update">
      <h2>Upcoming Bills</h2>
      {#if $upcomingExpenses.length > 0}
      {#each $upcomingExpenses as expense}
      <div class="update-item">
        <div>
          <p>{expense.Name}</p>
          <p>Due {fromISOString(expense.PaymentDate)}</p>
        </div>
        <p>{formatAmount(expense.Amount)}</p>
      </div>
      {/each}
      {:else}
        <p>No upcoming payments in the next two weeks.</p>
        {/if}
      
      
    </div>

    <div class="update">
      <h2>Recent Activity</h2>
      <div class="update-item">
        <div class="activity-item">
          <div class="activity-icon">🛒</div>
          <div>
            <p>Grocery Store</p>
            <p>Today</p>
          </div>
        </div>
        <p class="amount">-890 kr</p>
      </div>
      <div class="update-item">
        <div class="activity-item">
          <div class="activity-icon">💰</div>
          <div>
            <p>Salary Deposit</p>
            <p>Yesterday</p>
          </div>
        </div>
        <p class="amount">+32,500 kr</p>
      </div>
    </div>
  </div>
</div>

<style>
  .landing-page {
    padding: 20px;
    min-height: 100vh;
    margin-top: 2rem;
    max-width: 75%;
  }
  
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }

  .header-title {
    font-size: 24px;
    font-weight: bold;
  }


  .quick-actions {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 40px;
  }

  .financial-snapshot {
    background-color: white;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    padding: 20px;
    margin-bottom: 40px;
  }

  .snapshot-grid {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    gap: 20px;
  }

  .snapshot-item {
    padding: 20px;
    border-radius: 8px;
    color: #333;
  }

  .available-balance {
    background-color: #e1f3ff;
  }

  .monthly-savings {
    background-color: #d6ffd6;
  }

  .upcoming-bills {
    background-color: #f0e0ff;
  }

  .snapshot-title {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .snapshot-value {
    font-size: 24px;
    font-weight: bold;
    margin-top: 10px;
  }

  .quick-links {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 20px;
    margin-bottom: 40px;
  }

  .quick-link {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    cursor: pointer;
    transition: box-shadow 0.2s ease;
  }

  .quick-link:hover {
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.2);
  }

  .updates {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
  }

  .update {
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }

  .update-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 8px;
  }

  .activity-item {
    display: flex;
    gap: 10px;
    align-items: center;
  }

  .activity-icon {
    padding: 10px;
    border-radius: 50%;
    background-color: #f1f1f1;
  }

  .amount {
    font-weight: bold;
    color: #ff4d4d;
  }

  .update-item:nth-child(odd) {
    background-color: #f9f9f9;
  }

  .update-item:nth-child(even) {
    background-color: #e9f7e9;
  }
</style>
