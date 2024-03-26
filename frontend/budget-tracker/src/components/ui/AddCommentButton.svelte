<script>
    import FaRegComment from 'svelte-icons/fa/FaRegComment.svelte'
    import FaRegCommentDots from 'svelte-icons/fa/FaRegCommentDots.svelte'
    import CommentModal from '../modals/CommentModal.svelte';

    export let entry;

    let showModal = false;

    const isExpense = () => entry && 'Comment' in entry;

    const isIncome = () => entry && 'Comment' in entry;
</script>

{#if isExpense() && entry.Comment}
    <button class='icon-button' text='Show Comments' 
    on:click={() => (showModal = true)} 
    on:mouseover={() => console.log('it has commeth')} 
    on:mouseout={() => console.log('it has goneth')}
    on:focus={() => console.log('it has focuth')}
    on:blur={() => console.log('it has blurreth')}>
        <span class="icon"><FaRegCommentDots></FaRegCommentDots></span>
        <span>Add comment</span>
    </button>
{:else}
    <button title='Add Comment' class='icon-button' text='Add Comment' on:click={() => (showModal = true)}>
        <span class="icon"><FaRegComment></FaRegComment></span>
    </button>
{/if}





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
