<script>
    import { activeModal } from "../../stores/activeModal.js";
    import Modal from "./Modal.svelte";
    import Toast from "../Toast.svelte";
    import { onMount } from "svelte";
    import { postExpense, postListofExpenses } from '../../routes/api/addExpense.js';
    import { FetchExpenses } from '../../routes/api/fetchExpenses';
    import { getCategories } from '../../routes/api/fetchCategories.js';
    import { categories } from "../../stores/categories.js";
    

    onMount(async () => {
        await getCategories();
    });

    let name = '', amount = 0, category_id = 0, subcategory_id = 0, frequency = 'One-time', payment_date = '', start_date = '', end_date = '';
    let message = '', theme = 'info', duration = 3000;
    let isRecurring = false; // Track the checkbox state

    function createExpenseObject(date) {
        return {
            name,
            amount: parseInt(amount, 10),
            category_id,
            subcategory_id: subcategory_id || null, // Subcategory can be optional
            frequency,
            payment_date: date.toISOString(),
            paid: false, // Default to unpaid
            comment: "" // Add comment field if needed
        };
    }

    function closeModal() {
        activeModal.set(null); // Clear the active modal
    }

    async function submitExpenses() {
        await addExpense();
    }

    async function addExpense() {
        const start = new Date(start_date);
        let end;
        if (end_date) {
            end = new Date(end_date);
        }

        const expenses = [];

        if (frequency === 'One-time') {
            expenses.push(createExpenseObject(start));
        } else if (end) {
            let currentDate = new Date(start.getTime());

            while (currentDate <= end) {
                expenses.push(createExpenseObject(currentDate));
                incrementDate(currentDate, frequency);
            }
        }

        try {
            let response;
            if (expenses.length === 1) {
                response = await postExpense(expenses[0]);
                message = "Expense added successfully!";
                theme = "success";
            } else {
                response = await postListofExpenses(expenses);
                message = "All expenses added successfully!";
                theme = "success";
            }
        } catch (error) {
            console.error("Error adding expense:", error);
            message = "Error adding expense";
            theme = "error";
        }

        // Reset fields after submission
        name = '';
        amount = 0;
        category_id = 0;
        subcategory_id = 0;
        frequency = 'One-time';
        payment_date = '';
        start_date = '';
        end_date = '';
        FetchExpenses();
    }

    // Update frequency based on checkbox
    $: frequency = isRecurring ? frequency : 'One-time';
</script>




<!-- Reusable Modal -->
<Modal modalId="addExpense" onClose={closeModal} title="Add Expense">
    <div class="modal-container">
        <div class="add-expense-form">
        <div class="form-group">
            <label for="name">Name</label>
            <input id="name" type="text" bind:value={name} required />
        </div>
        <div class="form-group">
            <label for="amount">Amount</label>
            <div class="input-with-prefix">
                <span class="currency-prefix">kr</span>
                <input id="amount" type="number" bind:value={amount} required />
            </div>
        </div>
        <div class="form-group">
            <label for="category">Category</label>
            <select id="category" bind:value={category_id}>
                <option value="">Select category</option>
                {#each $categories as cat}
                    <option value={cat.id}>{cat.name}</option>
                {/each}
            </select>
            
            <!-- Show subcategories only if a category is selected -->
            {#if category_id}
                <select id="subcategory" bind:value={subcategory_id}>
                    <option value="">Select subcategory (optional)</option>
                    {#each $categories.find(cat => cat.id === category_id)?.subcategories || [] as sub}
                        <option value={sub.id}>{sub.name}</option>
                    {/each}
                </select>
            {/if}
        </div>

        <div>
            <div class="recurring-box">
                <input class="checkbox" type="checkbox" id="isRecurring" bind:checked={isRecurring} />
                <label for="isRecurring">Recurring?</label>
            </div>  

            <!-- Frequency Dropdown (only shown if isRecurring is true) -->
            {#if isRecurring}
                <select id="frequency" bind:value={frequency}>
                    <option value="">Select frequency</option>
                    <option value="Daily">Daily</option>
                    <option value="Weekly">Weekly</option>
                    <option value="Monthly">Monthly</option>
                    <option value="Annually">Annually</option>
                </select>
                <input id="date" type="date" bind:value={start_date} />
            {/if}
            <label for="date">Payment Date</label>
            <div class="input-with-icon">
                <input id="date" type="date" bind:value={payment_date} required />
            </div>
        
          {#if !(frequency === 'One-time' || frequency === '')}
            <input id="date" type="date" bind:value={end_date} />
          {/if} 
          <button on:click={submitExpenses}>Add Expense</button>
        </div>
    </div>
</div>
    <Toast {message} {theme} {duration} />
</Modal>

<style>
    .modal-container {
        padding: 0;
        overflow: hidden;
        max-width: 500px;
        width: 100%;
    }

    .add-expense-form {
        padding: 2rem;
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



    .input-with-icon {
        padding-top: 1rem;
    }

    .recurring-box {
        display: flex;
        align-items: center;
        gap: 10px; 
        font-size: 1rem; 
        color: #555; 
        margin-bottom: 15px; 
    }


    .recurring-box label {
        font-size: 1rem;
        cursor: pointer; 
        color: #333;
    }


    .recurring-box .checkbox {
        width: 18px; 
        height: 18px;
        border: 2px solid #ccc;
        border-radius: 4px; 
        background-color: #fff;
        cursor: pointer;
        transition: background-color 0.3s ease, border-color 0.3s ease;
    }

    .recurring-box .checkbox:checked {
        background-color: var(--primary-button-color); 
    }

    .recurring-box .checkbox:checked::after {
        content: '';
        position: absolute;
        top: 3px;
        left: 6px;
        width: 6px;
        height: 6px;
        background-color: #fff; 
        clip-path: polygon(0 50%, 40% 100%, 100% 0, 40% 25%);
        transform: rotate(45deg);
    }

        input,
        select {
        max-width: 100%;
    }

    input[type="text"],
    input[type="number"],
    input[type="date"],
    select {
      width: 100%;
      padding: 12px 16px;
      border: 1px solid #ccc;
      border-radius: 8px;
      box-sizing: border-box;
      font-size: 1rem;
      background-color: #f9f9f9;
      transition: border-color 0.3s ease, box-shadow 0.3s ease;
    }
  
    input[type="text"]:focus,
    input[type="number"]:focus,
    input[type="date"]:focus,
    select:focus {
      border-color: #5c6bc0;
      outline: none;
      box-shadow: 0 0 5px rgba(92, 107, 192, 0.5);
    }
  
    /* Button styles */
    button {
      width: 100%;
      padding: 12px;
      background-color: var(--primary-button-color);
      color: white;
      border: none;
      border-radius: 8px;
      cursor: pointer;
      font-size: 1rem;
      font-weight: bold;
      transition: background-color 0.3s ease, transform 0.2s ease, box-shadow 0.3s ease;
    }
  
    button:hover {
      background-color: var(--primary-button-hover-color);
      transform: translateY(-2px);
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }
  
    button:active {
      transform: translateY(0);
      box-shadow: none;
    }

  
    /* Responsive adjustments */
    @media (max-width: 500px) {
      .modal-container {
        padding: 15px;
      }
  
      input[type="text"],
      input[type="number"],
      input[type="date"],
      select {
        padding: 10px 12px;
      }
  
      button {
        font-size: 1rem;
      }
    }
  </style>
  