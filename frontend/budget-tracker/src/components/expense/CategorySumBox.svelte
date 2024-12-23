<script>
    import { filteredExpenses } from '../../stores/filteredExpenses';
    import { derived } from 'svelte/store';
    import { onMount } from 'svelte';
    import { FetchExpenses } from "../../routes/api/fetchExpenses";
    import { formatAmount } from '../../utility/functions';

    export let category = '';

    const categoryColors = {
        Food: 'rgba(255, 99, 132, 1)',
        Transportation: 'rgba(54, 162, 235, 1)',
        Housing: 'rgba(255, 206, 86, 1)',
        Utilities: 'rgba(75, 192, 192, 1)',
        Insurance: 'rgba(153, 102, 255, 1)',
        Healthcare: 'rgba(255, 159, 64, 1)',
        Entertainment: '#CDDC39',
        Clothing: '#FF5722',
        Miscellaneous: '#607D8B',
    };

    onMount(FetchExpenses);

    const totalCategoryExpense = derived(filteredExpenses, $filteredExpenses => {
        const total = $filteredExpenses.reduce((acc, {Category, Amount}) => {
            if (Category === category) {
                return acc + Amount;
            }
            return acc;
        }, 0);
        return total;
    });

    $: categoryColor = getCategoryColor(category);

    function getCategoryColor(category) {
        return categoryColors[category] || '#757575';
    }
</script>

<style>
    .total-expenses-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        border: var(--component-border);
        color: var(--primary-button-text-color);
        box-shadow: var(--component-box-shadow);
        padding: 20px;
        border-radius: 10px;
        font-family: var(--font-family);
        font-size: 1rem;
        text-align: center;
        width: fit-content;
    }

    .title {
        font-size: 1.18rem;
        font-weight: bold;
    }

    .expense {
        font-size: 0.8rem;
    }
</style>

<div class="total-expenses-container" style="background-color: {categoryColor};">
    <span class="title">{category}</span>
    <span class="expense">{formatAmount($totalCategoryExpense)}</span>
</div>
