<script>
  import { onMount } from "svelte";
  import { formatAmount, fromISOString } from "../../utility/functions";
  import { derived } from "svelte/store";
  import { filteredExpenses, expensesChange } from "../../stores/filteredExpenses";
  import { filteredIncome, incomeChange } from "../../stores/filteredIncome";
  import { selectionMode, selectedStartDate, selectedEndDate } from '../../stores/selectionMode';
  import { fetchIncome } from "../api/fetchIncome";
  import { FetchExpenses } from "../api/fetchExpenses";
  import { WalletMinimal, TrendingUp, TrendingDown } from 'lucide-svelte';
  import { getCategories } from '../../routes/api/fetchCategories.js'

  onMount(() => {
        FetchExpenses();
        fetchIncome();
        getCategories();
    });
  
  const totalIncome = derived(filteredIncome, $filteredIncome => {
    return $filteredIncome.reduce((total, income) => total + income.Amount, 0);
  });

  const totalExpenses = derived(filteredExpenses, $filteredExpenses => {
    return $filteredExpenses.reduce((total, expense) => total + expense.amount, 0);
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
                {#if $selectionMode === 'year'}
                  <p>Yearly Balance</p>
                  <div class="chart">Chart Placeholder</div>
                  <p class="amounts">{formatAmount($totalBalance)}</p>
                  
                  {:else}
                  {#if $selectionMode === 'month'}
                  <p>Monthly Balance</p>
                  <div class="chart">Chart Placeholder</div>
                  <p class="amounts">{formatAmount($totalBalance)}</p>
                  
                  {:else}
                  {#if $selectionMode === 'range'}
                  <p>Balance from {fromISOString($selectedStartDate)} to {fromISOString($selectedEndDate)}</p>
                  <div class="chart">Chart Placeholder</div>
                  <p class="amounts">{formatAmount($totalBalance)}</p>
                  <p>+2.5% from last period</p>
                    {:else}
                    <p>Total Balance</p>
                    <div class="chart">Chart Placeholder</div>
                    <p class="amounts">{formatAmount($totalBalance)}</p>
                  {/if}
                  {/if}
                  {/if}
              </div>
              <div class="total-label">
                <span class="total-icon"><WalletMinimal size="24" /></span>
              </div>
          </div>
      </div>

      <div class="card income">
          <div class="card-header">
            <div>
              {#if $selectionMode === 'year'}
                <p>Yearly Income</p>
                <div class="chart">Chart Placeholder</div>
                <p class="amounts">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</p>
                <p class="trend {$incomeChange >= 0 ? 'positive' : 'negative'}">
                  {$incomeChange >= 0 ? '+' : ''}{$incomeChange.toFixed(1)}% from last year
                </p> 
              {:else}
              {#if $selectionMode === 'month'}
                <p>Monthly Income</p>
                <div class="chart">Chart Placeholder</div>
                <p class="amounts">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</p>
                <p class="trend {$incomeChange >= 0 ? 'positive' : 'negative'}">
                  {$incomeChange >= 0 ? '+' : ''}{$incomeChange.toFixed(1)}% from last month
                </p> 
              {:else}
                {#if $selectionMode === 'range'}
                  <p>Income from {fromISOString($selectedStartDate)} to {fromISOString($selectedEndDate)}</p>
                  <div class="chart">Chart Placeholder</div>
                  <p class="amounts">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</p>
                  <p>+2.5% from last period</p>
                {:else}
                  <p>Total Income</p>
                  <div class="chart">Chart Placeholder</div>
                  <p class="amounts">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</p>
                {/if}
              {/if}
              {/if}
            </div>            
              <div class="income-label">
                <span class="income-icon"><TrendingUp size="24" /> </span>
              </div>
          </div>
      </div>

      <div class="card">
        <div class="card-header">
          <div>
            {#if $selectionMode === 'year'}
              <p>Yearly Expenses</p>
              <div class="chart">Chart Placeholder</div>
              <p class="amounts">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</p>
              <p class="expense-trend {$expensesChange > 0 ? 'positive' : 'negative'}">
                {$expensesChange > 0 ? '+' : ''}{$expensesChange.toFixed(1)}% from last year
              </p>              
            {:else}
            {#if $selectionMode === 'month'}
              <p>Monthly Expenses</p>
              <div class="chart">Chart Placeholder</div>
              <p class="amounts">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</p>
              <p class="expense-trend {$expensesChange > 0 ? 'positive' : 'negative'}">
                {$expensesChange > 0 ? '+' : ''}{$expensesChange.toFixed(1)}% from last month
              </p>              
            {:else}
              {#if $selectionMode === 'range'}
                <p>Expenses from {fromISOString($selectedStartDate)} to {fromISOString($selectedEndDate)}</p>
                <div class="chart">Chart Placeholder</div>
                <p class="amounts">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</p>
                <p>+2.5% from last period</p> 
              {:else}
                <p>Total Expenses</p>
                <div class="chart">Chart Placeholder</div>
                <p class="amounts">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</p>
              {/if}
            {/if}
            {/if}
          </div>  
            <div class="expense-label">
              <span class="expense-icon"><TrendingDown size="24" /></span>
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
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      margin-bottom: 20px;
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
      gap: 40px;
  }

  .chart {
      background: #fff;
      padding: 16px;
      border-radius: 8px;
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



.income {
  background-color: #d6ffd6;
}

</style>