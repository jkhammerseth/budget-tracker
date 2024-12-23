<script>
    import Modal from './Modal.svelte';
    import { FetchExpenses } from '../../routes/api/fetchExpenses';
    import { updateExpense } from '../../routes/api/updateExpense';
    import { deleteExpense } from '../../routes/api/deleteExpense';
    import { activeModal } from "../../stores/activeModal.js";
  
    export let expense = null;
    export const isOpen = false;
  
    let editMode = false;
    let updatedExpense = {};
    let confirmationDelete = false;
  
    // Ensure updatedExpense is populated when expense prop changes
    $: updatedExpense = { ...expense }; 
  
    function toggleEditMode() {
      editMode = !editMode;
      // Make sure updatedExpense is cloned for editing
      updatedExpense = { ...expense }; 
    }
  
    async function handleUpdate(event) {
      event.preventDefault();
      try {
        await updateExpense({ ID: expense.ID, ...updatedExpense });
        console.log('Expense updated successfully');
        await FetchExpenses();
        editMode = false; // Exit edit mode
      } catch (error) {
        console.error('Error updating expense:', error.message);
      }
    }
  
    async function handleDelete() {
      try {
        await deleteExpense(expense.ID);
        console.log('Expense deleted successfully');
        await FetchExpenses();
        confirmationDelete = false;
        closeModal();
      } catch (error) {
        console.error('Error deleting expense:', error.message);
      }
    }
  
    // Function to close the modal by resetting the activeModal store
    function closeModal() {
      activeModal.set(null);  // Close the modal
    }
  
    // When the modal is closed, reset selectedExpense in parent
    $: if (!$activeModal) {
      expense = null;  // Clear expense when modal closes
    }
  </script>
  
  <!-- Modal that is conditionally rendered based on activeModal store -->
  {#if $activeModal === 'editExpense' && expense}
      <Modal isOpen={expense !== null} onClose={closeModal} modalId="expense">
          <h2>{editMode ? "Edit Expense" : "Expense Details"}</h2>
  
          {#if confirmationDelete}
              <p>Are you sure you want to delete this expense?</p>
              <div class="button-group">
                  <button on:click={handleDelete} class="delete-button">Delete</button>
                  <button on:click={() => (confirmationDelete = false)} class="cancel-button">Cancel</button>
              </div>
          {:else if editMode}
              <form on:submit={handleUpdate}>
                  <div class="form-control">
                      <label for="name">Name:</label>
                      <input id="name" type="text" bind:value={updatedExpense.Name} required />
                  </div>
                  <div class="form-control">
                      <label for="amount">Amount:</label>
                      <input id="amount" type="number" bind:value={updatedExpense.Amount} required />
                  </div>
                  <div class="form-control">
                      <label for="category">Category:</label>
                      <select id="category" bind:value={updatedExpense.Category} required>
                          <option value="Food">Food</option>
                          <option value="Transportation">Transportation</option>
                          <option value="Housing">Housing</option>
                          <option value="Utilities">Utilities</option>
                          <option value="Insurance">Insurance</option>
                          <option value="Healthcare">Healthcare</option>
                          <option value="Entertainment">Entertainment</option>
                          <option value="Clothing">Clothing</option>
                          <option value="Miscellaneous">Miscellaneous</option>
                      </select>
                  </div>
                  <div class="form-control">
                      <label for="date">Payment Date:</label>
                      <input id="date" type="date" bind:value={updatedExpense.PaymentDate} required />
                  </div>
                  <div class="button-group">
                      <button type="submit" class="submit-button">Save</button>
                      <button type="button" on:click={toggleEditMode} class="cancel-button">Cancel</button>
                  </div>
              </form>
          {:else}
              <p><strong>Name:</strong> {expense.Name}</p>
              <p><strong>Amount:</strong> {expense.Amount + ' kr'}</p>
              <p><strong>Category:</strong> {expense.Category}</p>
              <p><strong>Status:</strong> {expense.Status}</p>
              <p><strong>Payment Date:</strong> {expense.PaymentDate}</p>
              <p><strong>Comment:</strong> {expense.Comment}</p>
  
              <div class="actions">
                  <button on:click={toggleEditMode} class="submit-button">Edit</button>
                  <button on:click={() => (confirmationDelete = true)} class="delete-button">Delete</button>
                  <button on:click={closeModal} class="cancel-button">Close</button>
              </div>
          {/if}
      </Modal>
  {/if}
  
  <style>    
      .actions, .button-group {
          display: flex;
          gap: 10px;
          margin-top: 20px;
      }
      .form-control {
          margin-bottom: 10px;
      }
      .submit-button {
          background-color: #3A87F2;
          color: white;
          padding: 5px;
          border-radius: 10px;
      }
      .submit-button:hover {
          background-color: #0056b3;
      }
      .delete-button {
          background-color: #f12626;
          color: white;
      }
      .delete-button:hover {
          background-color: #b30000;
      }
      .cancel-button {
          background-color: #ccc;
      }
      .cancel-button:hover {
          background-color: #999;
      }
  </style>
  