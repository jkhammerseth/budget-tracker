<script>
    import UpdateExpenseModal from '../modals/UpdateExpenseModal.svelte';
    import FaRegEdit from 'svelte-icons/fa/FaRegEdit.svelte'
    
    export let expense;
    

    let showModal = false;

    function fromISOString(isoString) {
    const date = new Date(isoString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0'); 
    const year = date.getFullYear().toString().substr(-2); 

    // Formatting the date as DD.MM.YY
    return `${day}.${month}.${year}`;
    }
</script>

<button title='Edit' class='icon-button' text='Edit' on:click={() => (showModal = true)}>
    <span class="icon"><FaRegEdit></FaRegEdit></span>
</button>

<UpdateExpenseModal expense={expense} date={fromISOString(expense.Date)} bind:showModal/>

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