<script>
    import { activeModal } from "../../stores/activeModal.js";
    import Modal from "./Modal.svelte";
    import Toast from "../Toast.svelte";
    import { writable } from 'svelte/store';
    import { postExpense, postListofExpenses } from '../../routes/api/addExpense.js';
    import { FetchExpenses } from '../../routes/api/fetchExpenses';
    import Calendar from '../Calendar.svelte'
    import { selectedDay } from '../../stores/selectionMode.js'
 

    let name = '', amount = 0, category = '', frequency = 'One-time', startDate = '', endDate = '';
    let message = '', theme = 'info', duration = 3000;

    let isRecurring = false; // track the checkbox state
    
    let expensesList = writable([{ name: '', amount: 0, category: '', frequency: '', startDate: '', endDate: '' }]);
    
    const categories = ['Food', 'Transportation', 'Housing', 'Utilities', 'Insurance', 'Healthcare', 'Entertainment', 'Clothing', 'Miscellaneous'];

    function addExpenseField() {
        expensesList.update(list => [...list, { name: '', amount: 0, category: '', frequency: '', startDate: '', endDate: '' }]);
    }

    function createExpenseObject(date) {
    return {
        Name: name,
        Amount: parseInt(amount, 10),
        Category: category,
        Frequency: frequency,
        PaymentDate: date.toISOString()
    };
}

    function closeModal() {
        activeModal.set(null); // Clear the active modal
    }

    // Submits the expenses based on the view
    async function submitExpenses() {
        await addExpense(); 
    }

    function handleResponse(response) {
    if (!response.ok) {
        throw new Error("Failed to add expense");
    }
        message = "Expense added successfully!";
        theme = "success";
    }

    async function addExpense() {
        const start = new Date(startDate); 
        let end;
        if (endDate) {
            end = new Date(endDate);
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

            handleResponse(response);
        } catch (error) {
            console.error("Error adding expense:", error);
            message = "Error adding expense";
            theme = "error";
        }

        name = '';
        amount = 0;
        category = '';
        frequency = 'One-time'; // Reset to default
        startDate = '';
        endDate = '';
        FetchExpenses();
    }


    // Update frequency based on checkbox
    $: frequency = isRecurring ? frequency : 'One-time';
</script>

<!-- Your existing HTML and Modal code goes here -->



<!-- Reusable Modal -->
<Modal modalId="addExpense">
    <h2>Add Expense</h2>
    <div class="modal-container">
         
          
          <div class="form-control">

            <label for="name" class="sr-only">Name</label>
            <input id="name" placeholder="Name" type="text" bind:value={name} />

            <label for="amount" class="sr-only">Amount</label>
            <input id="amount" placeholder="Amount" type="number" bind:value={amount} />

            <label for="category">Category</label>
            <select id="category" bind:value={category}>
                <option value="">Select category</option>
                {#each categories as cat} 
                  <option value={cat}>{cat}</option>
                {/each}
            </select>
                  
            <div class="recurring-box">
            <label for="isRecurring">Recurring?</label>
            <input class="checkbox" type="checkbox" id="isRecurring" bind:checked={isRecurring} />
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
            {/if}
            <label for="date" class="sr-only">Date</label>
            <input id="date" type="date" bind:value={startDate} />
        
          {#if !(frequency === 'One-time' || frequency === '')}
            <input id="date" type="date" bind:value={endDate} />
          {/if} 
          <button on:click={submitExpenses}>Add Expense</button>
        </div>
    </div>
    <Toast {message} {theme} {duration} />
</Modal>

<style>
.modal-container {
    display: grid;
    grid-template-columns: 1fr; /* 1 column layout for smaller screens */
    grid-template-rows: auto 1fr; /* Ensure the form fills remaining space */
    gap: 20px;
    padding: 20px;
    max-width: 800px;
    height: 100%; /* Ensure the modal container takes up all available space */
    margin: auto;
    justify-items: center; /* Ensures content is centered */
}

/* Container for the recurring box */
.recurring-box {
    display: flex;
    align-items: center; /* Center the checkbox and label vertically */
    gap: 10px; /* Space between label and checkbox */
    font-size: 1rem; /* Set a nice font size */
    color: #555; /* Lighter text color */
    margin-bottom: 15px; /* Space below the recurring box */
}

/* Style the label */
.recurring-box label {
    font-size: 1rem;
    cursor: pointer; /* Indicate that the label is clickable */
    color: #333;
}

/* Style the checkbox */
.recurring-box .checkbox {
    width: 18px; /* Adjust the size of the checkbox */
    height: 18px;
    border: 2px solid #ccc;
    border-radius: 4px; /* Rounded corners */
    background-color: #fff;
    cursor: pointer; /* Change the cursor to indicate it's clickable */
    transition: background-color 0.3s ease, border-color 0.3s ease; /* Smooth transition */
}

/* Change the appearance when checkbox is checked */
.recurring-box .checkbox:checked {
    background-color: #5c6bc0; /* Color for checked box */
    border-color: #5c6bc0; /* Match the border color to the background */
}

/* Add a custom checkmark appearance */
.recurring-box .checkbox:checked::after {
    content: ''; /* Create a pseudo-element for the checkmark */
    position: absolute;
    top: 3px;
    left: 6px;
    width: 6px;
    height: 6px;
    background-color: #fff; /* White checkmark */
    clip-path: polygon(0 50%, 40% 100%, 100% 0, 40% 25%); /* Checkmark shape */
    transform: rotate(45deg); /* Rotate the checkmark to make it appear correctly */
}

/* Focus state for accessibility */
.recurring-box .checkbox:focus {
    outline: none;
    border-color: #5c6bc0;
}


.calendar {
    width: 100%;
    max-width: 400px; /* Adjust to fit within modal, can also increase or remove */
    margin-bottom: 20px; /* Add spacing between calendar and form */
}

.form-control {
    display: flex;
    flex-direction: column;
    align-items: stretch;
    gap: 15px;
    width: 100%;
    max-width: 600px; /* Limit form width */
    overflow-y: auto; /* Allow scrolling if content overflows */
}

    h2 {
      font-size: 1.75rem;
      color: #333;
      text-align: center;
      margin: 0;
      font-weight: bold;
    }
  
    .form-control {
        display: flex;
        flex-direction: column;
        align-items: stretch; /* Ensures uniform width */
        gap: 15px;
    }

        input,
        select {
        max-width: 100%;
    }

  
    .form-control label {
      font-size: 1rem;
      color: #555;
    }
  
    /* Input and select field styles */
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
      background-color: #5c6bc0;
      color: white;
      border: none;
      border-radius: 8px;
      cursor: pointer;
      font-size: 1rem;
      font-weight: bold;
      transition: background-color 0.3s ease, transform 0.2s ease, box-shadow 0.3s ease;
    }
  
    button:hover {
      background-color: #3f51b5;
      transform: translateY(-2px);
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }
  
    button:active {
      transform: translateY(0);
      box-shadow: none;
    }
  
  
    .calendar {
        width: 100%;
        max-width: 400px; /* Adjust to fit within modal */
    }

  
    /* Responsive adjustments */
    @media (max-width: 500px) {
      .modal-container {
        padding: 15px;
      }
  
      h2 {
        font-size: 1.5rem;
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
  