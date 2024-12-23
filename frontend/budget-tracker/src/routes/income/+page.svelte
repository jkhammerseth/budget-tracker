<script>
    import Calendar from "../../components/Calendar.svelte";
    import IncomeTable from "../../components/IncomeTable.svelte";
    import AddIncome from "../../components/AddIncome.svelte";
    import MonthlyIncomeGraph from "../../components/MonthlyIncomeGraph.svelte";
  import TotalIncome from "../../components/income/TotalIncome.svelte";
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

    .add-income-container,
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
        <MonthlyIncomeGraph/>
    </div>
    <div class="info">
</div>
</div>
    <div class="budget-table-container">
        <IncomeTable />
    </div>
    <div class="add-income-container">
        <AddIncome />
    </div>

</div>

