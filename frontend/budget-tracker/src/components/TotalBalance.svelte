<script>
  import { onMount } from 'svelte';
    import { filteredExpenses } from '../stores/filteredExpenses';
    import { filteredIncome } from '../stores/filteredIncome'; // Ensure correct import path
    import { derived } from 'svelte/store';
  import { fetchIncome } from '../routes/api/fetchIncome';
  import { FetchExpenses } from '../routes/api/fetchExpenses';

    onMount(() =>{
        fetchIncome();
        FetchExpenses();
    });

    // Calculating the total balance by subtracting total expenses from total income
    const totalBalance = derived([filteredExpenses, filteredIncome], ([$filteredExpenses, $filteredIncome]) => {
        const totalExpenses = $filteredExpenses.reduce((total, expense) => total + expense.Amount, 0);
        const totalIncome = $filteredIncome.reduce((total, income) => total + income.Amount, 0);
        return totalIncome - totalExpenses;
    });

    const totalIncome = derived(filteredIncome, $filteredIncome => {
        return $filteredIncome.reduce((total, income) => total + income.Amount, 0);
    });

    const totalExpenses = derived(filteredExpenses, $filteredExpenses => {
        return $filteredExpenses.reduce((total, expense) => total + expense.Amount, 0);
    });

    // Function to format the amount in a more readable way
    function formatAmount(amount) {
        return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'NOK' }).format(amount); // Adjust currency code as needed
    }
</script>
<div class="asset-box">
    <h1>Assets</h1>
    <hr />
<div class="container">
    <div class="total-balance-container">
        <h2>Total Balance</h2>
        <!-- Use the formatAmount function to format the displayed total balance -->
        <h1>{$totalBalance >= 0 ? formatAmount($totalBalance) : '-' + formatAmount(-$totalBalance)}</h1>
    </div>
    <div class="total-income-container">
        <h2>Total Income</h2>
        <h1>{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</h1>
    </div>
    <div class="total-expenses-container">
        <h2>Total Expenses</h2>
        <h1>{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</h1>
    </div>
</div>
</div>

<style>
    .asset-box {
        
        border-radius: 20px;
        padding: 20px;
        max-width: 400px;
        font-family: 'Roboto', sans-serif;
    }

    .container {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        grid-template-rows: auto auto;
        max-width: full-width;
        font-family: 'Roboto', sans-serif;
        padding: 20px;
        gap: 20px;
        border-radius: 20px;
        margin: 40px auto;
    }

    .total-balance-container, .total-income-container, .total-expenses-container {
        border: black 1px solid;
        background-color: #f8f9fa;
        padding: 20px;
        border-radius: 10px;
        font-family: 'Roboto', sans-serif;

    }

    .total-balance-container {
        grid-column: 2 / 3;
        grid-row: 1;
        text-align: left;
        margin: 20px 0;
    }

    .total-income-container {
        grid-column: 1 / 2;
        grid-row: 1;
        text-align: left;
        margin: 20px 0;
    }

    .total-expenses-container {
        grid-column: 1 / 2;
        grid-row: 2;
        text-align: left;
        margin: 20px 0;
    }

    h2 {
        margin: 0;
        color: #666;
        text-align: left;
    }

    h1 {
        margin: 10px 0 0;
        color: #333;
        font-size: 2em;
        text-align: left;
    }
</style>
