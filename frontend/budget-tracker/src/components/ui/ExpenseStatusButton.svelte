<script>
    import { FetchExpenses } from "../../routes/api/fetchExpenses";

    export let expense;

    async function updateStatusofExpense(status) {
      try {
        const response = await fetch(`http://localhost:8080/api/users/expenses/${expense.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            Paid: status,
          }),
          credentials: 'include',
        });
  
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.message || 'Failed to update the status of expense');
        }
        console.log('Expense status updated successfully');
        await FetchExpenses();
      } catch (error) {
        console.error('Error updating the status of the expense:', error.message);
      }
    }
 
    function handleStatusClick() {
        const newStatus = expense.Paid === true ? false : true;
        updateStatusofExpense(newStatus);
    }

</script>

{#if expense.Paid === false}
    <button class="expense-status-button expense-status-button--not-paid" on:click={handleStatusClick}>Not Paid</button>
{:else}
    <button class="expense-status-button expense-status-button--paid" on:click={handleStatusClick}>Paid</button>
{/if}

<style>
    .expense-status-button {
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 5px;
        border-radius: 10px;
        border: none;
        cursor: pointer;
    }

    .expense-status-button--paid {
        background-color: transparent;
        color: var(--success);
    }

    .expense-status-button--not-paid {
        background-color: transparent;
        color: var(--danger);
    }
</style>