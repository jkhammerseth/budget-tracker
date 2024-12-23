<script>
  import TotalBalance from "../../components/TotalBalance.svelte";
  import TransactionTable from "../../components/TransactionTable.svelte";
</script>

<div class="container">
<div class="total-balance">
  <TotalBalance />
</div>
<div class="transaction-table">
    <TransactionTable />
  </div>
</div>

<style>
 .container {
    display: grid;
    grid-template-columns: 2fr auto;
    grid-template-rows: auto;
    gap: 20px;
    padding: 40px;
    margin: 0 auto;
    max-width: var(--max-width);
  }
  
  .transaction-table {
      grid-column: 1 / 3;
      grid-row: 2;
  }

  .total-balance {
      grid-column: 2 / 3;
      grid-row: 1;
  }
    

</style>