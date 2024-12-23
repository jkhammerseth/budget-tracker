<script>
  import { onMount } from "svelte";
  import { formatAmount, fromISOString } from "../../utility/functions";
  import { derived } from "svelte/store";
  import { filteredExpenses } from "../../stores/filteredExpenses";
  import { filteredIncome } from "../../stores/filteredIncome";
  import { selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate, selectedDay } from '../../stores/selectionMode';
  import { fetchIncome } from "../api/fetchIncome";
  import { FetchExpenses } from "../api/fetchExpenses";
  import { faDollarSign, faMoneyBillTrendUp, faCar, faHome, faBolt, faFileShield, faHeartbeat, faFilm, faTshirt, faBoxOpen, faCalendar } from '@fortawesome/free-solid-svg-icons';
  import { FontAwesomeIcon } from '@fortawesome/svelte-fontawesome';

  onMount(() => {
        FetchExpenses();
        fetchIncome();
    });


  let monthlyData = [
      { name: "Jul", income: 52000, expenses: 38000 },
      { name: "Aug", income: 54000, expenses: 42000 },
      { name: "Sep", income: 51000, expenses: 37000 },
      { name: "Oct", income: 55000, expenses: 41000 },
      { name: "Nov", income: 53000, expenses: 39000 },
      { name: "Dec", income: 56000, expenses: 43000 }
  ];

  let upcomingPayments = [
      { name: "Rent", amount: "9000 kr", due: "Dec 31", status: "Not Paid" },
      { name: "Internet Bill", amount: "599 kr", due: "Dec 28", status: "Upcoming" },
      { name: "Power Bill", amount: "1200 kr", due: "Dec 30", status: "Upcoming" }
  ];

  


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
                  <p>+2.5% from last month</p>
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
                <p>+2.5% from last month</p>
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
              <p>+2.5% from last period</p>
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
          <h2>Upcoming Payments</h2>
          {#each upcomingPayments as payment}
              <div class="payment">
                  <div>
                      <p>{payment.name}</p>
                      <p>{payment.due}</p>
                  </div>
                  <div>{payment.amount}</div>
              </div>
          {/each}
      </div>
  </div>
</div>


<style>

  .dashboard {
      padding: 16px;
      max-width: 75%;
      margin-top: 2rem;
  }

  .amounts {
    font-weight: 800;
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

  .payment {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 16px;
  }

  .payment:last-child {
      margin-bottom: 0;
  }
</style>