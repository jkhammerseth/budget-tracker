<script>
    import { FetchExpenses } from "../../routes/api/fetchExpenses";
  
    export let expense;
    const today = new Date();
  
    async function updateStatusOfExpense(status) {
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
        updateStatusOfExpense(newStatus);
    }
  
    function getStatusLabel(expense) {
        const today = new Date().setHours(0, 0, 0, 0); // Set the time to midnight to ignore time part
      
        if (expense.Paid) return 'Paid';
  
        const paymentDate = new Date(expense.PaymentDate).setHours(0, 0, 0, 0); // Normalize the payment date to midnight
        if (!expense.Paid && paymentDate < today) return 'Overdue';
  
        return 'Not Paid';
      }
  
    function getStatusClass() {
      const today = new Date().setHours(0, 0, 0, 0); // Set the time to midnight to ignore time part
      
      if (expense.Paid) return 'expense-status-button--paid';
  
      const paymentDate = new Date(expense.PaymentDate).setHours(0, 0, 0, 0); // Normalize the payment date to midnight
      if (!expense.Paid && paymentDate < today) return 'expense-status-button--overdue';
  
      return 'expense-status-button--not-paid';
    }
  </script>
  
  <button class="expense-status-button {getStatusClass()}" on:click={handleStatusClick}>
    {#if expense}
      <span class="expense-status-button-text">{getStatusLabel(expense)}</span>
    {:else}
      <span>Invalid Expense</span>  <!-- Or a different message if the expense is not defined -->
    {/if}
  </button>
  
  <style>
    :root {
        --success: #63bf66; 
        --danger: #ba0f03; 
        --warning: #ff9800;
    }
  
    .expense-status-button {
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 4px 8px;
        border-radius: 5px; /* Creates pill shape */
        border: none;
        color: white;
        font-size: 0.8rem;
        font-weight: bold;
        cursor: pointer;
        transition: background-color 0.3s ease;
        box-shadow: var(--component-box-shadow);
    }
  
    .expense-status-button--paid {
        background-color: var(--success);
    }
  
    .expense-status-button--not-paid {
        background-color: var(--danger);
    }
  
    .expense-status-button--overdue {
        background-color: var(--warning);
    }
  
    .expense-status-button:hover {
        filter: brightness(0.9);
    }
  
    /* Mobile-responsive styles */
    @media (max-width: 768px) {
      .expense-status-button {
        padding: 4px 6px; /* Reduced padding on smaller screens */
        font-size: 0.7rem; /* Smaller font size */
        width: auto; /* Ensure the button width adjusts based on content */
      }
  
      .expense-status-button-text {
        font-size: 0.7rem; /* Smaller text size inside the button */
      }
    }
  
    @media (max-width: 480px) {
      .expense-status-button {
        padding: 3px 5px; /* Further reduce padding on very small screens */
        font-size: 0.6rem; /* Further reduce font size */
      }
  
      .expense-status-button-text {
        font-size: 0.6rem; /* Further reduce font size */
      }
    }
  </style>
  