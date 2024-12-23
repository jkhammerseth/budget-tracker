<script>
    import ExpensesTable from "../../components/ExpensesTable.svelte";
    import { filteredExpenses } from "../../stores/filteredExpenses";
    import { derived } from 'svelte/store';
    import { onMount } from 'svelte';
    import { FetchExpenses } from "../api/fetchExpenses";

    onMount(() => {
        FetchExpenses();
    });

    const uniqueCategories = derived(filteredExpenses, $filteredExpenses => {
        const categories = new Set();
        $filteredExpenses.forEach(expense => {
        categories.add(expense.Category);
        });
        return Array.from(categories);
    });
</script>

<style>
    :root {
        --max-width: 1200px;
        --table-padding: 10px;
    }

    /* Container for the whole dashboard */
    .dashboard-container {
        display: flex;
        flex-wrap: wrap;
        gap: 20px;
        padding: 40px;
        margin: 0 auto;
        max-width: var(--max-width);
        flex-direction: column; /* For mobile responsiveness */
    }

    /* Graph and other content section */
    .graph-and-stuff {
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        gap: 20px;
    }

    /* Budget Table Container */
    .budget-table-container {
        width: 75%; /* Ensure it takes full width on small screens */
        background-color: transparent;
    }


    /* Responsive adjustments for mobile */
    @media (max-width: 768px) {
        .dashboard-container {
            padding: 20px; /* Less padding on small screens */
        }

        /* Stack graph and table vertically on smaller screens */
        .graph-and-stuff {
            width: 100%;
        }

        .budget-table-container {
            margin-top: 20px; /* Add space between elements */
        }
    }
</style>

<div class="dashboard-container">
    <div class="graph-and-stuff">
        <!-- Any graphs or other content can go here -->
    </div>
    <div class="budget-table-container">
        <ExpensesTable />
    </div>
</div>
