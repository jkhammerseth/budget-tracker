<script>
    import AddExpense from "../../components/AddExpense.svelte";
    import Calendar from "../../components/Calendar.svelte";
    import ExpensesTable from "../../components/ExpensesTable.svelte";
    import MonthlyExpensesGraph from "../../components/MonthlyExpensesGraph.svelte";
    import SharedExpense from "../../components/SharedExpense.svelte";
    import { onMount } from 'svelte';
 

</script>

<style>
    .dashboard-container {
    display: grid;
    grid-template-columns: 2fr 1fr; /* 2 columns where first part is double the size of the second */
    grid-template-rows: auto; /* Defines the rows to be automatically sized based on content */
    gap: 20px;
    padding: 40px;
    margin: 0 auto;
    max-width: 1400px;
}


    .add-expense-container,
    .calendar-container,
    .bar-chart-container,
    .budget-table-container {
        background: #FFFFFF;
        padding: 20px;
        background-color: transparent;
        
    }

    .bar-chart-container {
        grid-column: 1 / 3;
        grid-row: 1;
    }

    .calendar-container {
        grid-column: 3 / 4;
        grid-row: 1;
    }


    .add-expense-container {
        grid-column: 3 / 4;
        grid-row: 2;
        text-align: right;
        align-items: center;
    }

    .shared-expense-container {
        grid-column: 3 / 4;
        grid-row: 3;
        text-align: right;
        align-items: center;
    }



    .budget-table-container {
        grid-column: 1 / 3;
        grid-row: 2 / 4;
        padding-top: 0%;
    }

    @media (max-width: 1400px) {
        .dashboard-container {
            grid-template-columns: 1fr;
            grid-template-rows: auto auto auto auto;
        }

        .calendar-container {
            grid-column: 1 / 2;
            grid-row: 1;
        }

        .bar-chart-container {
            grid-column: 1 / 2;
            grid-row: 2;
        }

        .add-expense-container {
            grid-column: 1 / 2;
            grid-row: 3;
            text-align: center;
        }

        .budget-table-container {
            grid-column: 1 / 2;
            grid-row: 4;
        }
    }
</style>

<div class="dashboard-container">
    <div class="bar-chart-container">
        <MonthlyExpensesGraph/>
    </div>
   
    <div class="add-expense-container">
        <AddExpense />
    </div>
    <div class="shared-expense-container">
        <SharedExpense />
    </div>
 
    <div class="budget-table-container">
        <ExpensesTable />
    </div>
    <div class="calendar-container">
        <Calendar />
    </div>

</div>
