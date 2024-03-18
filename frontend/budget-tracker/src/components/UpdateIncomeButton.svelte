<script>
    import UpdateIncomeModal from './UpdateIncomeModal.svelte';
    import FaRegEdit from 'svelte-icons/fa/FaRegEdit.svelte'
    
    export let income;

    let showModal = false;
</script>

<button title='Edit' class='icon-button' text='Edit' on:click={() => (showModal = true)}>
    <span class="icon"><FaRegEdit></FaRegEdit></span>
</button>

<UpdateIncomeModal income={income} date={income.Date} bind:showModal/>

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
    color: #3A87F2; 
  }

  .icon:active {
    transform: scale(0.9);
  }
</style>