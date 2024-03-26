<script>
  import LoanCard from "../../components/LoanCard.svelte";
  import { onMount } from "svelte";
  import { fetchLoans } from "../api/fetchLoans";

  onMount(() => {
    fetchLoans();
  });

</script>

<div class="loans-container">
<LoanCard />
</div>

<style>
  .loans-container {
    display: flex;
    max-width: var(--max-width);
    justify-items: center;
  }
</style>