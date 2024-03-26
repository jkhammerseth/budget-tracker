<script>
    import { fetchIncome } from '../../routes/api/fetchIncome';
    import { onMount, onDestroy } from 'svelte';
  
    export let showModal = false;
    export let income;
    export let date;
  
    let modalContent;

    const categories = [
        'Salary',
        'Investment',
        'Rental',
        'Interest',
        'Dividends',
        'Business',
        'Other'
    ];
  
    function close() {
        showModal = false;
    }
  
    function handleKeydown(event) {
        if (event.key === 'Escape') {
            close();
        }
    }

    function toISOFormat(datetoFormat) {
        const [year, month, day] = datetoFormat.split('-');
        return `${year}-${month}-${day}T00:00:00Z`;
    }
  
    async function updateIncome(event) {
      event.preventDefault();
      try {
        const response = await fetch(`http://localhost:8080/api/users/incomes/${income.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            name: income.Name,
            amount: income.Amount,
            category: income.Category,
            date: toISOFormat(income.Date),
          }),
          credentials: 'include',
        });
  
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.message || 'Failed to update income');
        }
        console.log('Income updated successfully');
        await fetchIncomee();
        close();
      } catch (error) {
        console.error('Error updating income:', error.message);
      }
    }
 
    onMount(() => {
      window.addEventListener('keydown', handleKeydown);
    });
  
    onDestroy(() => {
      window.removeEventListener('keydown', handleKeydown);
    });
  </script>

{#if showModal}
<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<div role="dialog" class="modal-backdrop" on:click={close} aria-label="Close modal">
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
    <div bind:this={modalContent} class="modal-content" on:click|stopPropagation tabindex="0">
        <h1>Edit Income</h1>
        <form on:submit={updateIncome}>
            <div class="form-control">
            <input type="text" id="name" name="name" required bind:value={income.Name}>
        </div>
        <div class="form-control">
            <input type="number" id="amount" name="amount" required bind:value={income.Amount}>
        </div>
        <div class="form-control">
            <select id="category" name="category" required bind:value={income.Category}>
            {#each categories as cat}
                <option value={cat}>{cat}</option>
            {/each}
            </select>
        </div>
        <div class="form-control">
            <input type="date" id="date" name="date" required bind:value={date}>
        </div>
            <div class="button-group">
                <button class="submit-button" type="submit">Save</button>
                <button class="cancel-button" type="button" on:click={close}>Cancel</button>
            </div>
        </form>
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
    .form-control {
        margin-bottom: 20px;
        width: 100%;
    }
    input, select, button {
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
    .submit-button {
        background-color: #3A87F2;
        color: white;
        cursor: pointer;
        transition: background-color 0.3s;
    }
    .submit-button:hover {
        background-color: #0056b3;
    }
    .cancel-button {
        background-color: #f12626;
        color: white;
        cursor: pointer;
        transition: background-color 0.3s;
    }
    .cancel-button:hover {
        background-color: #b30000;
    }

    @media (max-width: 600px) {
        .modal-content {
            width: calc(100% - 20px);
            padding: 10px;
        }
    }
</style>

