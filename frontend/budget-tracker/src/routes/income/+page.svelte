<script>
    import IncomeTable from "../../components/IncomeTable.svelte";
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
</style>


<div class="dashboard-container">
        <IncomeTable />
</div>

