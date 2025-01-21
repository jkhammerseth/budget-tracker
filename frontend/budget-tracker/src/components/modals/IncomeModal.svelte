<script>
    import Modal from './Modal.svelte';
    import { onMount } from "svelte";
    import { activeModal } from "../../stores/activeModal.js";
    import { fromISOString, formatDate } from '../../utility/functions';
    import { Calendar } from 'lucide-svelte';
    import { getIconComponent } from '../../utility/icons'
    import { fetchIncome } from '../../routes/api/fetchIncome';
    import { deleteIncome } from '../../routes/api/deleteIncome'

    export let income = null;
  
    let confirmationDelete = false;

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
                            <svelte:component this={getIconComponent(income.Category.name)} size="16" />
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
                </div>
                <div class="button-group">
                    <button on:click={() => (confirmationDelete = true)} class="btn btn-danger">Delete</button>
                    <button on:click={console.log("edit")} class="btn btn-primary">Edit</button>
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