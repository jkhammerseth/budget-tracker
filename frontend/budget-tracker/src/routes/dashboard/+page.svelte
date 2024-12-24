<script>
  import { onMount } from "svelte";
  import { formatAmount, fromISOString } from "../../utility/functions";
  import { derived } from "svelte/store";
  import { filteredExpenses, expensesChange } from "../../stores/filteredExpenses";
  import { filteredIncome, incomeChange } from "../../stores/filteredIncome";
  import { selectionMode, selectedStartDate, selectedEndDate } from '../../stores/selectionMode';
  import { fetchIncome } from "../api/fetchIncome";
  import { FetchExpenses } from "../api/fetchExpenses";
  import { faDollarSign, faMoneyBillTrendUp, faCalendar } from '@fortawesome/free-solid-svg-icons';
  import { FontAwesomeIcon } from '@fortawesome/svelte-fontawesome';
  import { expenses } from '../../stores/expenses';

  onMount(() => {
        FetchExpenses();
        fetchIncome();
    });

  const today = new Date();
  const twoWeeksFromToday = new Date();
  twoWeeksFromToday.setDate(today.getDate() + 14);

  const upcomingExpenses = derived(expenses, $expenses => {
    return $expenses.filter(expense => {
      const expenseDate = new Date(expense.PaymentDate);
      return expenseDate >= today && expenseDate <= twoWeeksFromToday;
    });
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
</script>

<div class="dashboard">
  <div class="cards">
      <div class="card">
          <div class="card-header">
              <div>
                  {#if $selectionMode === 'month'}
                  <p>Monthly Balance</p>
                  <p class="amounts">{formatAmount($totalBalance)}</p>
                  
                  {:else}
                  {#if $selectionMode === 'range'}
                  <p>Balance from {fromISOString($selectedStartDate)} to {fromISOString($selectedEndDate)}</p>
                  <p class="amounts">{formatAmount($totalBalance)}</p>
                  <p>+2.5% from last period</p>
                    {:else}
                    <p>Total Balance</p>
                    <p class="amounts">{formatAmount($totalBalance)}</p>
                  {/if}
                  {/if}
              </div>
              <div class="total-label">
                <span class="total-icon"><FontAwesomeIcon icon={faDollarSign}/></span>
              </div>
          </div>
      </div>
      <div class="card">
          <div class="card-header">
            <div>
              {#if $selectionMode === 'month'}
                <p>Monthly Income</p>
                <p class="amounts">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</p>
                <p class="trend {$incomeChange >= 0 ? 'positive' : 'negative'}">
                  {$incomeChange >= 0 ? '+' : ''}{$incomeChange.toFixed(1)}% from last month
                </p> 
              {:else}
                {#if $selectionMode === 'range'}
                  <p>Income from {fromISOString($selectedStartDate)} to {fromISOString($selectedEndDate)}</p>
                  <p class="amounts">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</p>
                  <p>+2.5% from last period</p>
                {:else}
                  <p>Total Income</p>
                  <p class="amounts">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</p>
                {/if}
              {/if}
            </div>            
              <div class="income-label">
                <span class="income-icon"><FontAwesomeIcon icon={faMoneyBillTrendUp}/></span>
              </div>
          </div>
      </div>
      <div class="card">
        <div class="card-header">
          <div>
            {#if $selectionMode === 'month'}
              <p>Monthly Expenses</p>
              <p class="amounts">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</p>
              <p class="expense-trend {$expensesChange > 0 ? 'positive' : 'negative'}">
                {$expensesChange > 0 ? '+' : ''}{$expensesChange.toFixed(1)}% from last month
              </p>              
            {:else}
              {#if $selectionMode === 'range'}
                <p>Expenses from {fromISOString($selectedStartDate)} to {fromISOString($selectedEndDate)}</p>
                <p class="amounts">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</p>
                <p>+2.5% from last period</p> 
              {:else}
                <p>Total Expenses</p>
                <p class="amounts">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</p>
              {/if}
            {/if}
          </div>  
            <div class="expense-label">
              <span class="expense-icon"><FontAwesomeIcon icon={faMoneyBillTrendUp}/></span>
            </div>
        </div>
    </div>
  </div>

  <!-- Main Content -->
  <div class="main-content">
      <div class="chart">
          <h2>Income vs Expenses</h2>
          <div>Chart Placeholder</div>
      </div>
      <div class="upcoming-payments">
        <div class="upcoming-header">
        <h2>Upcoming Payments</h2>
        <span class="calendar-icon"><FontAwesomeIcon icon={faCalendar}/></span>
      </div>
        {#if $upcomingExpenses.length > 0}
            {#each $upcomingExpenses as payment}
            <div class="payment">
                <div>
                    <p class="payment-name">{payment.Name}</p>
                    <p>Due {fromISOString(payment.PaymentDate)}</p>
                </div>
                <div>
                    <p class="payment-amount">{formatAmount(payment.Amount)}</p>
                    <span>{payment.Paid ? "Paid" : "Upcoming"}</span>
                </div>
                
            </div>
            {/each}
            <div class="view">
              <a class="view-link" href="/expenses">View All</a>
            </div>
        {:else}
            <p>No upcoming payments this month.</p>
        {/if}
    </div>
    
  </div>
</div>


<style>

  .trend.positive {
    color: green;
  }

  .trend.negative {
    color: red;
  }

  .expense-trend.positive {
    color: red;
  }

  .expense-trend.negative {
    color: green;
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


  .dashboard {
      padding: 16px;
      max-width: 75%;
      margin-top: 2rem;
  }

  .amounts {
    font-weight: 700;
    font-size: 1.7rem;
  }

  .cards {
      display: flex;
      flex-direction: column;
      gap: 16px;
      margin-bottom: 24px;
  }

  .card {
      background: #fff;
      padding: 16px;
      border-radius: 8px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .card .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
  }

  .card p {
      margin: 8px 0;
  }

  .main-content {
      display: flex;
      flex-direction: column;
      gap: 16px;
  }

  .chart {
      background: #fff;
      padding: 16px;
      border-radius: 8px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .total-icon {
    height: 100px;
    width: 100px;
    color: #040404;
  }

  .total-label {
    padding: 16px;
    background-color: #4a86ff;
    border-radius: 20%;
  }

  .expense-icon {
    height: 100px;
    width: 100px;
    color: #040404;
  }

  .expense-label {
    padding: 16px;
    background-color: #f46b64;
    border-radius: 20%;
  }

  .income-icon {
    height: 100px;
    width: 100px;
    color: #228927;
  }

  .income-label {
    padding: 16px;
    background-color: #7ffa86;
    border-radius: 20%;
  }

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

  .payment {
      display: flex;
      justify-content: space-between;
      align-items: center;
      border-radius: 8px;
      padding: 2px 10px;
      margin-bottom: 5px;
      background-color: #f9f9f9;
      transition: background-color 0.2s ease;
  }
  .payment-amount,
  .payment-name {
    text-transform: capitalize;
    padding: 0;
    font-weight: 600;
    font-size: 1.2em;
  }

  .calendar-icon {
    padding: 16px;
  }

  .payment:last-child {
      margin-bottom: 0;
  }
</style>