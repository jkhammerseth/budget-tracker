<script>
import { filteredIncome } from '../../stores/filteredIncome';
import { derived } from 'svelte/store';
import { onMount } from 'svelte';
  import { fetchIncome } from "../../routes/api/fetchIncome";

onMount(() =>{
    fetchIncome();
});

const totalIncome = derived(filteredIncome, $filteredIncome => {
    return $filteredIncome.reduce((total, income) => total + income.Amount, 0);
});

    // Function to format the amount in a more readable way
function formatAmount(amount) {
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'NOK' }).format(amount); // Adjust currency code as needed
}
</script>

<div class="total-income-container">
    <span class="title">Income</span>
    <span class="income">{$totalIncome >= 0 ? formatAmount($totalIncome) : '-' + formatAmount(-$totalIncome)}</span>
</div>

<style>
    .total-income-container {
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

    .income {
        font-size: 1.2rem;
    }
</style>