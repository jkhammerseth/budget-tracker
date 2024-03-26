<script>
    import { FetchExpenses } from "../../routes/api/fetchExpenses";
    import { onMount, onDestroy } from "svelte";

    export let showModal = false;
    export let expense;

    let modalContent;

    function close() {
        showModal = false;
    }

    function handleKeydown(event) {
        if (event.key === "Escape") {
            close();
        }
    }

    async function handleDelete(expenseId) {
    try {
      const response = await fetch(`http://localhost:8080/api/users/expenses/${expenseId}`, {
        method: 'DELETE',
        credentials: 'include',
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to delete expense');
      }

      console.log('Expense deleted successfully');
      await FetchExpenses();
      close();
    } catch (error) {
      console.error('Error deleting expense:', error.message);
      // TODO: Update UI to show error message to the user here
    }
  }

    onMount(() => {
        window.addEventListener("keydown", handleKeydown);
    });

    onDestroy(() => {
        window.removeEventListener("keydown", handleKeydown);
    });
</script>

{#if showModal}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
    <div role="dialog" class="modal-backdrop" on:click={close} on:keyup={close} aria-label="Close modal">
            <!-- svelte-ignore a11y-no-static-element-interactions -->
    <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
    <div bind:this={modalContent} class="modal-content" on:click|stopPropagation tabindex="0">
            <h2>Delete Expense</h2>
            <p>Are you sure you want to delete this expense?</p>
            <div class="modal-buttons">
                <div class="button-group">
                    <button class="delete-button" type="submit" on:click={handleDelete(expense.ID), close}>Delete</button>
                    <button class="cancel-button" type="button" on:click={close}>Cancel</button>
                </div>
            </div>
        </div>
    </div>
{/if}

<style>


    .modal-backdrop {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.4);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000; 
    }

    .modal-content {
        background-color: #f8f9fa;
        border: 1px solid black;
        border-radius: 10px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        max-width: 400px;
        padding: 20px;
        width: calc(100% - 40px);
        box-sizing: border-box;
        display: flex;
        flex-direction: column;
        gap: 20px;
    }

    button {
        width: 100%;
        padding: 10px;
        margin-top: 5px;
        border: 1px solid #ccc;
        border-radius: 5px;
        box-sizing: border-box;
    }

    .button-group {
        display: flex;
        justify-content: space-between;
        flex-direction: row;
        gap: 10px;
    }

    .cancel-button {
        background-color: #3A87F2;
        color: white;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    .cancel-button:hover {
        background-color: #0056b3;
    }

    .delete-button {
        background-color: #f12626;
        color: white;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    .delete-button:hover {
        background-color: #b30000;
    }

    @media (max-width: 600px) {
        .modal-content {
            width: calc(100% - 20px);
            padding: 10px;
        }
    }
</style>