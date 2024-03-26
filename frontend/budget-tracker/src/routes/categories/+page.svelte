<script>
    import ExpensesByCategoryBox from "../../components/ExpensesByCategoryBox.svelte";
</script>

<div class="dashboard-container"></div>
<div class="expense-by-category-container">
    <ExpensesByCategoryBox />
</div>

<style>
    .dashboard-container {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        grid-template-rows: auto auto;
        gap: 20px;
        padding: 20px;
        margin: 0 auto;
        max-width: 1200px;
    }
    .expense-by-category-container {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 20px;
        padding: 20px;
        margin: 0 auto;
        max-width: var(--max-width);
    }
</style>
