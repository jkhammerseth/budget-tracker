<script>
import { FetchExpenses } from "../../routes/api/fetchExpenses";
import { filteredExpenses } from '../../stores/filteredExpenses';
import { derived } from 'svelte/store';
import { onMount } from 'svelte';

onMount(() =>{
    FetchExpenses();
});

const totalExpenses = derived(filteredExpenses, $filteredExpenses => {
    return $filteredExpenses.reduce((total, expense) => total + expense.Amount, 0);
});

    // Function to format the amount in a more readable way
function formatAmount(amount) {
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'NOK' }).format(amount); // Adjust currency code as needed
}
</script>

<div class="total-expenses-container">
    <span class="title">Expenses</span>
    <span class="expense">{$totalExpenses >= 0 ? formatAmount($totalExpenses) : '-' + formatAmount(-$totalExpenses)}</span>
</div>

<style>
    .total-expenses-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        border: black 1px solid;
        color: var(--primary-button-text-color);
        background-color: var(--primary-button-color);
        padding: 20px;
        border-radius: 10px;
        font-family: 'Roboto', sans-serif;
        font-size: 0.5rem;
        text-align: left;
        height: 2rem;
    }

    .title {
        font-size: 1.5rem;
        font-weight: bold;
    }

    .expense {
        font-size: 1.2rem;
    }
</style>