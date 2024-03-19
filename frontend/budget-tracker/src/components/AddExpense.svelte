<script>
    import { FetchExpenses } from '../routes/api/fetchExpenses';
    import { postExpense, postListofExpenses } from '../routes/api/addExpense';
    import Toast from './Toast.svelte';

    let name = '';
    let amount = 0;
    let category = '';
    let frequency = '';
    let startDate = '';
    let endDate = '';
  
    const categories = [
      'Food',
      'Transportation',
      'Housing',
      'Utilities',
      'Insurance',
      'Healthcare',
      'Entertainment',
      'Clothing',
      'Miscellaneous'
    ];

    let message = '';
    let theme = 'info'; // Possible values: 'success', 'error', 'info'
    let duration = 3000;

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

    resetFormFields();
    FetchExpenses();
}

function createExpenseObject(date) {
    return {
        Name: name,
        Amount: parseInt(amount, 10),
        Category: category,
        Frequency: frequency,
        Date: date.toISOString()
    };
}

function incrementDate(date, frequency) {
    switch (frequency) {
        case 'Daily':
            date.setDate(date.getDate() + 1);
            break;
        case 'Weekly':
            date.setDate(date.getDate() + 7);
            break;
        case 'Monthly':
            date.setMonth(date.getMonth() + 1);
            break;
        case 'Annually':
            date.setFullYear(date.getFullYear() + 1);
            break;
    }
}

function handleResponse(response) {
    if (!response.ok) {
        throw new Error("Failed to add expense");
    }
    message = "Expense added successfully!";
    theme = "success";
}

function resetFormFields() {
    name = '';
    amount = 0;
    category = '';
    frequency = '';
    startDate = '';
    endDate = '';
}

  </script>
  
  
  <style>
    .form-container {
      width: 400px;
      margin: auto;
      padding: 20px;
      background-color: #E8EFF1;
      border-style: solid;
      border-color: black;
      border-width: 1px;
      border-radius: 10px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      font-family: 'Roboto', sans-serif;
    }

    h2 {
      color: #333;
      text-align: center;
      margin-bottom: 20px;
    }

    .form-control {
      margin-bottom: 20px;
      width: 100%;
    }

    input[type="text"],
    input[type="number"],
    input[type="date"],
    select {
      width: calc(100% - 20px);
      padding: 10px;
      margin-top: 5px;
      margin-right: 10px;
      border: 1px solid black;
      border-radius: 5px;
      box-sizing: border-box; 
    }

    button {
      display: block;
      width: 100%;
      padding: 10px;
      margin-top: 20px;
      background-color: #3A87F2;
      color: white;
      border: none;
      border-radius: 5px;
      cursor: pointer;
      transition: background-color 0.3s;
    }

    button:hover {
      background-color: #0056b3;
    }

    @media (max-width: 480px) {
        .form-container {
            padding: 10px;
        }
    }
</style>

<div class="form-container">
  <h2>Add Expense</h2>
  <div class="form-control">
    <input id="name" placeholder="Name" type="text" bind:value={name} />
  </div>
  <div class="form-control">
    <input id="amount" placeholder="Amount" type="number" bind:value={amount} />
  </div>
  <div class="form-control">
    <select id="category" bind:value={category}>
        <option value="">Select category</option>
        {#each categories as cat} 
          <option value={cat}>{cat}</option>
        {/each}
      </select>
  </div>
  <div class="form-control">
    <select id="frequency" bind:value={frequency}>
      <option value="">Select frequency</option>
      <option value="One-time">One-time</option>
      <option value="Daily">Daily</option>
      <option value="Weekly">Weekly</option>
      <option value="Monthly">Monthly</option>
      <option value="Annually">Annually</option>
    </select>
  </div>
  <div class="form-control">
    <input id="date" type="date" bind:value={startDate} />
  </div>
  {#if !(frequency === 'One-time' || frequency === '')}
    <div class="form-control">
      <input id="date" type="date" bind:value={endDate} />
    </div>
  {/if} 
  <button on:click={addExpense}>Add Expense</button>
  <Toast {message} {theme} {duration} />
</div>