<script>
    import IncomeTable from "../../components/IncomeTable.svelte";
    import MonthlyIncomeGraph from "../../components/MonthlyIncomeGraph.svelte";
</script>

<style>
    .dashboard-container {
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        gap: 20px;
        padding: 40px;
        max-width: 75%;
    }

    .graph-and-stuff {
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        gap: 20px;
    }


    .bar-chart-container,
    .budget-table-container {
        background-color: transparent;   
    }

   
</style>


<div class="dashboard-container">
    <div class="graph-and-stuff">
    <div class="bar-chart-container">
        <MonthlyIncomeGraph/>
    </div>


</div>
    <div class="budget-table-container">
        <IncomeTable />
    </div>

</div>

