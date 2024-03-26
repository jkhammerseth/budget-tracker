<script>
    import { fetchIncome } from "../../routes/api/fetchIncome";

    export let income;

    async function updateStatusofIncome(status) {
      try {
        const response = await fetch(`http://localhost:8080/api/users/incomes/${income.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            Received: status,
          }),
          credentials: 'include',
        });
  
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.message || 'Failed to update the status of income');
        }
        console.log('income status updated successfully');
        await fetchIncome();
      } catch (error) {
        console.error('Error updating the status of the income:', error.message);
      }
    }
    function handleStatusClick() {
        const newStatus = income.Received === true ? false : true;
        updateStatusofIncome(newStatus);
    }

</script>

{#if income.Received === false}
    <button class="income-status-button income-status-button--not-paid" on:click={handleStatusClick}>Not Received</button>
{:else}
    <button class="income-status-button income-status-button--paid" on:click={handleStatusClick}>Received</button>
{/if}

<style>
    .income-status-button {
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 5px;
        border-radius: 10px;
        border: none;
        cursor: pointer;
    }

    .income-status-button--paid {
        background-color: transparent;
        color: var(--success);
    }

    .income-status-button--not-paid {
        background-color: transparent;
        color: var(--danger);
    }
</style>