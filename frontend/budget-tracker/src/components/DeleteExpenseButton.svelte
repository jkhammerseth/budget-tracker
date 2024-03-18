<script>
    import DeleteExpenseModal from './DeleteExpenseModal.svelte';
    import FaRegTrashAlt from 'svelte-icons/fa/FaRegTrashAlt.svelte'

    export let expense;

    let showModal = false;
</script>

<button title='Delete' class='icon-button' text='Delete' on:click={() => (showModal = true)}>
    <span class="icon"><FaRegTrashAlt/></span>
</button>

<DeleteExpenseModal expense={expense} bind:showModal/>

<style>
    .icon-button {
    background: none;
    border: none;
    padding: 0;
    cursor: pointer;
    outline: inherit;
    display: inline-flex;
    align-items: center;
    justify-content: center;
  }

  .icon {
    width: 16px;
    height: 16px;
    color: #333;
  }

  .icon:hover {
    color: #e74c3c;
  }

  .icon:active {
    transform: scale(0.9);
  }
</style>