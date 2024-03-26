<script>
    import AddCommentButton from "./ui/AddCommentButton.svelte";
    import UpdateExpenseButton from "./ui/UpdateExpenseButton.svelte";
    import DeleteExpenseButton from "./ui/DeleteExpenseButton.svelte";


    export let show = false;
    export let expense;
</script>

{#if show}
<div>
  <div class="expense-action-menu">
  <ul>
    <li><AddCommentButton entry={expense}/></li>
    <li><UpdateExpenseButton {expense}/></li>
    <li><DeleteExpenseButton {expense}/></li>
  </ul>
  </div>
</div>
{/if}


<style>

</style>