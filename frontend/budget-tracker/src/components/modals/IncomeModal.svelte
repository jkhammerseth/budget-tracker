<script>
    import Modal from './Modal.svelte';
    import { onMount } from "svelte";
    import { activeModal } from "../../stores/activeModal.js";
    import { fromISOString, formatDate } from '../../utility/functions';
    import { Calendar } from 'lucide-svelte';
    import { getIconComponent } from '../../utility/icons'
    import { fetchIncome } from '../../routes/api/fetchIncome';
    import { deleteIncome } from '../../routes/api/deleteIncome'
    import { updateIncome } from '../../routes/api/updateIncome'

    export let income = null;
  
    let editMode = false;
    let updatedIncome = {};
    let confirmationDelete = false;

    const incomeCategories = ["Salary", "Vipps"]

    function toggleEditMode() {
        editMode = !editMode;
        if (editMode) {
            updatedIncome = { 
            ...income,
            Category: income?.Category || "",
            Date: income?.Date ? income.Date.split('T')[0] : "",
            Amount: income?.Amount ? income.Amount.toString().split('.')[0] : "",
            };
        }
    }

    async function handleDelete() {
      try {
        await deleteIncome(income.ID);
        console.log('Income deleted successfully');
        await fetchIncome();
        confirmationDelete = false;
        closeModal();
      } catch (error) {
        console.error('Error deleting income:', error.message);
      }
    }
  
    // Function to close the modal by resetting the activeModal store
    function closeModal() {
      activeModal.set(null);  // Close the modal
    }
  
    $: if (!$activeModal) {
      income = null;  
    }

    async function handleUpdate(event) {
    event.preventDefault();

    // Prepare the payload
    const payload = {
        ID: income.ID, 
        Name: updatedIncome.Name,
        Amount: parseFloat(updatedIncome.Amount),
        Category: updatedIncome.Category,
        Date: updatedIncome.Date
            ? new Date(updatedIncome.Date).toISOString() // Convert to ISO string
            : null,
    };

    console.log('Updating expense with data:', payload);

    try {
        await updateIncome(payload);
        console.log('Income updated successfully');
        await fetchIncome();
        editMode = false;
    } catch (error) {
        console.error('Error updating income:', error.message);
    }
}



  </script>
  
<!-- Modal Content Structure -->
{#if $activeModal === 'income' && income}
    <Modal onClose={closeModal} modalId="income" title={income.Name}>
        <div class="modal-content">

            <!-- Delete Confirmation -->
            {#if confirmationDelete}
                <div class="confirmation-dialog">
                    <div class="warning-icon">⚠️</div>
                    <p class="confirmation-text">Are you sure you want to delete this income?</p>
                    <div class="button-group">
                        <button on:click={() => (confirmationDelete = false)} class="btn btn-secondary">Cancel</button>
                        <button on:click={handleDelete} class="btn btn-danger">Delete</button>
                    </div>
                </div>

                        <!-- Edit Form -->
                    {:else if editMode}
                    <form on:submit={handleUpdate} class="edit-form">
                        <div class="form-group">
                            <label for="name">Name</label>
                            <input id="name" type="text" bind:value={updatedIncome.Name} required />
                        </div>
                        <div class="form-group">
                            <label for="amount">Amount</label>
                            <div class="input-with-prefix">
                                <span class="currency-prefix">kr</span>
                                <input id="amount" type="float" bind:value={updatedIncome.Amount} required />
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="category">Category</label>
                            <select id="category" bind:value={updatedIncome.Category}>
                                <option value="">Select category</option>
                                {#each incomeCategories as cat}
                                    <option value={cat}>{cat}</option>
                                {/each}
                            </select>
                        </div>
                        <div class="form-group">
                            <label for="date">Date</label>
                            <div class="input-with-icon">
                                <Calendar size="16" class="input-icon" />
                                <input id="date" type="date" bind:value={updatedIncome.Date} required />
                            </div>
                        </div>
                        <div class="button-group">
                            <button type="button" on:click={toggleEditMode} class="btn btn-secondary">Cancel</button>
                            <button type="submit" class="btn btn-primary">Save Changes</button>
                        </div>
                    </form>

                <!-- View Mode -->
                {:else}
                <div class="expense-details">
                    <div class="detail-item">
                        <span class="detail-label">Amount</span>
                        <span class="detail-value">{income.Amount} kr</span>
                    </div>
                    <div class="detail-item">
                        <span class="detail-label">Category</span>
                        <div class="detail-value with-icon">
                        {#if getIconComponent(income.Category)}
                            <span class="category-icon">
                            <svelte:component this={getIconComponent(income.Category)} size="16" />
                            </span>
                        {:else}
                            📁
                        {/if}
                            <span class="detail-value">{income.Category}</span>
                    </div>
                    </div>
                    <div class="detail-item">
                        <span class="detail-label">Date</span>
                        <span class="detail-value with-icon">
                            <Calendar size="16" />
                            {formatDate(income.Date)}
                        </span>
                    </div>
                    <div class="detail-item">
                        <span class="detail-label">Comment</span>
                        <span class="detail-value">{income.Comment}</span>

                    </div>
                </div>
                <div class="button-group">
                    <button on:click={() => (confirmationDelete = true)} class="btn btn-danger">Delete</button>
                    <button on:click={toggleEditMode} class="btn btn-primary">Edit</button>
                </div>
                {/if}
            </div>
    </Modal>
{/if}
  
<style>
    .modal-content {
        padding: 0;
        border-radius: 12px;
        overflow: hidden;
        max-width: 500px;
        width: 100%;
    }

    .confirmation-dialog {
        padding: 2rem;
        text-align: center;
    }

    .warning-icon {
        font-size: 2.5rem;
        margin-bottom: 1rem;
    }

    .confirmation-text {
        color: #4a5568;
        margin-bottom: 1.5rem;
        font-size: 1.1rem;
    }

    .edit-form {
        padding: 1.5rem;
    }

    .form-group {
        margin-bottom: 1.25rem;
    }

    .form-group label {
        display: block;
        margin-bottom: 0.5rem;
        color: #4a5568;
        font-size: 0.9rem;
        font-weight: 500;
    }

    .input-with-prefix,
    .input-with-icon {
        position: relative;
        display: flex;
        align-items: center;
    }

    .currency-prefix {
        position: absolute;
        left: 1rem;
        color: #666;
    }

    input, select {
        width: 100%;
        padding: 0.75rem 1rem;
        border: 1px solid #e2e8f0;
        border-radius: 8px;
        font-size: 1rem;
        transition: border-color 0.2s, box-shadow 0.2s;
    }

    input:focus, select:focus {
        outline: none;
        border-color: #3A87F2;
        box-shadow: 0 0 0 3px rgba(58, 135, 242, 0.1);
    }

    .input-with-prefix input {
        padding-left: 3rem;
    }

    .input-with-icon input {
        padding-left: 2.5rem;
    }

    .expense-details {
        padding: 1.5rem;
    }

    .detail-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0.75rem 0;
        border-bottom: 1px solid #e9ecef;
    }

    .detail-item:last-child {
        border-bottom: none;
    }

    .detail-label {
        color: #666;
        font-size: 0.9rem;
    }

    .detail-value {
        font-weight: 500;
        color: #1a1a1a;
    }

    .detail-value.with-icon {
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .comment {
        font-style: italic;
        color: #666;
    }

    .button-group {
        display: flex;
        gap: 1rem;
        padding-bottom: 1rem;
        border-top: 1px solid #e9ecef;
    }

    .btn {
        flex: 1;
        padding: 0.75rem 1.5rem;
        border: none;
        border-radius: 8px;
        font-size: 0.95rem;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.2s;
    }

    .btn-primary {
        background-color: #3A87F2;
        color: white;
    }

    .btn-primary:hover {
        background-color: #2970d6;
    }

    .btn-secondary {
        background-color: #e2e8f0;
        color: #4a5568;
    }

    .btn-secondary:hover {
        background-color: #cbd5e0;
    }

    .btn-danger {
        background-color: #f12626;
        color: white;
    }

    .btn-danger:hover {
        background-color: #d91f1f;
    }
</style>