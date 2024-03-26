<script>
    import AddExpenseContainer from "../../components/AddExpenseContainer.svelte";
    import Calendar from "../../components/Calendar.svelte";
    import Clock from "../../components/Clock.svelte";
    import ExpensesTable from "../../components/ExpensesTable.svelte";
    import MonthlyExpensesGraph from "../../components/MonthlyExpensesGraph.svelte"
    import TotalExpenses from "../../components/expense/TotalExpenses.svelte";

</script>

<style>
    .dashboard-container {
        display: flex;
        flex-wrap: wrap;
        gap: 20px;
        padding: 40px;
        margin: 0 auto;

        
        max-width: var(--max-width);
    }

    .graph-and-stuff {
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        gap: 20px;
    }

    .info {
        display: flex;
        width: 43rem;
        flex-direction: row;
        gap: 20px;
    }

    .add-expense-container,
    .calendar-container,
    .clock-container,
    .bar-chart-container,
    .budget-table-container {
        background-color: transparent;   
    }

    .bar-chart-container {
        width: 43rem;
        
    }

  
</style>

<div class="dashboard-container">
    <div class="graph-and-stuff">
    <div class="bar-chart-container">
        <MonthlyExpensesGraph/>
    </div>
    <div class="info">
    <div class="clock-container">
        <TotalExpenses />
    </div>
    <div class="total-expenses-container">
        <TotalExpenses />
    </div>
    <div class="total-expenses-container">
        <TotalExpenses />
    </div>
</div>
</div>
    <div class="calendar-container">
        <Calendar />
    </div> 
    <div class="budget-table-container">
        <ExpensesTable />
    </div>
    <div class="add-expense-container">
        <AddExpenseContainer />
    </div>

</div>
