<script>
import { onMount } from 'svelte';
import { filteredExpenses } from '../stores/filteredExpenses';
import { filteredIncome } from '../stores/filteredIncome';
import { derived } from 'svelte/store';
import { fetchIncome } from '../routes/api/fetchIncome';
import { FetchExpenses } from '../routes/api/fetchExpenses';
import { formatAmount } from '../utility/functions'

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
</script>



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


<style>

    .container {
        display: flex;
        flex-direction: row;
        max-width: full-width;
        font-family: var(--font-family);
        gap: 20px;
        border-radius: 20px;
        margin: 40px auto;
        border-radius: 20px;
        padding: 20px;
        max-width: 40rem;
        background-color: var(--component-bg-color);
        max-height: fit-content;
        border: 1px solid var(--component-border-color);
        color: var(--primary-button-text-color);
    }

    .total-balance-container, .total-income-container, .total-expenses-container {
        border: black 1px solid;
        background-color: var(--primary-button-color);
        padding: 20px;
        border-radius: 10px;
        font-family: 'Roboto', sans-serif;
        font-size: 0.5rem;
        text-align: left;

    }

    h2 {
        margin: 0;
        color: var(--primary-button-text-color);
        text-align: left;
    }

    h1 {
        margin: 10px 0 0;
        color: var(--primary-button-text-color);
        font-size: 2em;
        text-align: left;
    }
</style>
