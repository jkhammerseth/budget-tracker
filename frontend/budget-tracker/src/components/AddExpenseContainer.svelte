<script>
    import { writable } from 'svelte/store';
    import { postExpense, postListofExpenses } from '../routes/api/addExpense';
    import { FetchExpenses } from '../routes/api/fetchExpenses';
    import Toast from './Toast.svelte';

    let componentView = 'single';

    let name = '', amount = 0, category = '', frequency = '', startDate = '', endDate = '';
    let message = '', theme = 'info', duration = 3000;
    let numberOfPeople = 1, userShare = 0;
    
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

    // Submits the expenses based on the view
    async function submitExpenses() {
        if (componentView === 'single') {
            await addExpense();
        } else if (componentView === 'multiple') {
            await addMultipleExpenses();
        } else if (componentView === 'shared') {
            await addSharedExpense();
        }
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
        frequency = '';
        startDate = '';
        endDate = '';
        FetchExpenses();
    }

    async function addMultipleExpenses() {
        const expenses = $expensesList;
        for (const expense of expenses) {
            try {
                let response;
                response = await postExpense({
                    Name: expense.name,
                    Amount: parseInt(expense.amount, 10),
                    Category: expense.category,
                    Frequency: expense.frequency,
                    Date: new Date(expense.startDate).toISOString(),
                });
                message = "Expense added successfully!";
                theme = "success";

                handleResponse(response);
            } catch (error) {
                console.error("Error adding expense:", error);
                message = "Error adding expense";
                theme = "error";
            }
        }
        FetchExpenses();
        

    }

    async function addSharedExpense() {
        calculateShare();
        const expense = {
            Name: name + " (shared)",
            Amount: userShare,
            Category: category,
            Frequency: frequency,
            Date: new Date(startDate).toISOString()
        };
        try {
            let response;
            response = await postExpense(expense);
            message = "Expense added successfully!";
            theme = "success";

            handleResponse(response);
        } catch (error) {
            console.error("Error adding expense:", error);
            message = "Error adding expense";
            theme = "error";
        }

        FetchExpenses(); 
        //reset form fields
        name = '';
        amount = 0;
        category = '';
        frequency = 'One-time';
        startDate = '';
        endDate = '';
        numberOfPeople = 1;
    }

    function calculateShare() {
        userShare = parseFloat((amount / numberOfPeople).toFixed(2));
    }

    function setView(view) {
        return function() {
            componentView = view;
        };
    }
</script>

<div class="container">
    <div class="button-row">
        
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <span class="view-tab {componentView === 'single' ? 'active' : ''}" on:click={setView('single')}>Add Expense</span>
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <span class="view-tab {componentView === 'multiple' ? 'active' : ''}" on:click={setView('multiple')}>Add Expenses</span>
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <span class="view-tab {componentView === 'shared' ? 'active' : ''}" on:click={setView('shared')}>Shared Expense</span>
    </div>


{#if componentView === 'single'}
<div class="form-container">
  <h2>Add Expense</h2>
  <div class="form-control">
    <label for="name" class="sr-only">Name</label>
    <input id="name" placeholder="Name" type="text" bind:value={name} />
    <label for="amount" class="sr-only">Amount</label>
    <input id="amount" placeholder="Amount" type="number" bind:value={amount} />
    <select id="category" bind:value={category}>
        <option value="">Select category</option>
        {#each categories as cat} 
          <option value={cat}>{cat}</option>
        {/each}
      </select>
    <select id="frequency" bind:value={frequency}>
      <option value="">Select frequency</option>
      <option value="One-time">One-time</option>
      <option value="Daily">Daily</option>
      <option value="Weekly">Weekly</option>
      <option value="Monthly">Monthly</option>
      <option value="Annually">Annually</option>
    </select>
    <label for="date" class="sr-only">Date</label>
    <input id="date" type="date" bind:value={startDate} />

  {#if !(frequency === 'One-time' || frequency === '')}
      <input id="date" type="date" bind:value={endDate} />

  {/if} 
  <button on:click={submitExpenses}>Add Expense</button>
  <Toast {message} {theme} {duration} />
</div>
</div>
{:else if componentView === 'multiple'}
<div class="form-container">
    <h2>Add Expenses</h2>
    {#each $expensesList as expense, index (index)}
        <div class="expense-item">
            <div class="form-control">
                <input id="name" placeholder="Name" type="text" bind:value={expense.name} />
                <input id="amount" placeholder="Amount" type="number" bind:value={expense.amount} />
                <select id="category" bind:value={expense.category}>
                    <option value="">Select category</option>
                    {#each categories as cat}
                        <option value={cat}>{cat}</option>
                    {/each}
                </select>
                <select id="frequency" bind:value={expense.frequency}>
                    <option value="">Select frequency</option>
                    <option value="One-time">One-time</option>
                    <option value="Daily">Daily</option>
                    <option value="Weekly">Weekly</option>
                    <option value="Monthly">Monthly</option>
                    <option value="Annually">Annually</option>
                </select>
                <input id="date" type="date" bind:value={expense.startDate} />
                {#if !(expense.frequency === 'One-time' || expense.frequency === '')}
                    <input id="endDate" type="date" bind:value={expense.endDate} />
                {/if}
            </div>
        </div>
    {/each}
    <button on:click={addExpenseField}>Add Another Expense</button>
    <button on:click={submitExpenses}>Submit All</button>
    <Toast {message} {theme} {duration} />
</div>
{:else if componentView === 'shared'}
<div class="form-container">
    <h2>Shared Expense</h2>
    <div class="form-control">
        <input id="name" placeholder="Name" type="text" bind:value={name} />
        <input id="amount" placeholder="Total amount" type="number" bind:value={amount} min="0" />
        <input id="people" placeholder="Number of People" type="number" bind:value={numberOfPeople} min="1" on:input={calculateShare} />
        <select id="category" bind:value={category}>
            <option value="">Select category</option>
            {#each categories as cat} 
            <option value={cat}>{cat}</option>
            {/each}
        </select>
        <select id="frequency" bind:value={frequency}>
            <option value="One-time">One-time</option>
            <option value="Daily">Daily</option>
            <option value="Weekly">Weekly</option>
            <option value="Monthly">Monthly</option>
            <option value="Annually">Annually</option>
        </select>
        <input id="date" type="date" bind:value={startDate} />
    {#if frequency !== 'One-time'}
        <input id="endDate" type="date" bind:value={endDate} />
    {/if}
    </div>
    <div>Your share: {userShare} kr</div>
    <button on:click={submitExpenses}>Add Your Share</button>
    <Toast {message} {theme} {duration} />
</div>
{/if}
</div>
<style>
    .button-row {
        display: flex;
        align-items: center;
        justify-content: center;
        
    }

    .view-tab {
        display: flex;
        justify-content: center;
        text-align: center;
        padding: 10px;
        background-color: var(--primary-button-color);
        color: var(--primary-button-text-color);
        border: none;
        cursor: pointer;
        transition: background-color 0.3s;
        width: 33.33%;
        border-color: var(--component-border-color);
        border-style: solid;
        border-width: 1px;
    }

    .sr-only {
      position: absolute;
      width: 1px;
      height: 1px;
      padding: 0;
      margin: -1px;
      overflow: hidden;
      clip: rect(0, 0, 0, 0);
      white-space: nowrap;
      border: 0;
    }

    .view-tab:first-of-type {
        border-top-left-radius: var(--component-border-radius);
        border-right: 0;
    }

    .view-tab:last-of-type {
        border-top-right-radius: var(--component-border-radius);
        border-left: 0;
    }


    .view-tab.active {
        background-color: var(--component-bg-color);
        border-bottom-color: var(--component-bg-color);
        color: var(--text-color);

        cursor: pointer;
    }

    .view-tab.active:hover {
        background-color: var(--component-bg-color);
    }

    .view-tab:hover {
        background-color: var(--primary-button-hover-color);
    }

    .form-container {
      max-width: 18rem;
      margin: auto;
      padding: 20px;
      background-color: var(--component-bg-color);
      border-style: solid;
      border-color: var(--component-border-color);
      border-width: 1px;
      border-radius: var(--component-border-radius);
      border-top-left-radius: 0;
      border-top-right-radius: 0;
      box-shadow: var(--component-box-shadow);
      font-family: var(--font-family);
      border-top: none;
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
      background-color: var(--primary-button-color);
      color: var(--primary-button-text-color);
      border: none;
      border-radius: 5px;
      cursor: pointer;
      transition: background-color 0.3s;
    }

    button:hover {
      background-color: var(--primary-button-hover-color);
    }

    .container {
      justify-content: center;
      align-items: center;
      height: 100%;
      width: 20rem;
    }

    @media (max-width: 480px) {
        .form-container {
            padding: 10px;
        }
    }

    
</style>