<script>
    import { writable } from 'svelte/store';
    import { postExpense } from '../routes/api/addExpense';
    import { FetchExpenses } from '../routes/api/fetchExpenses';

    let expensesList = writable([{ name: '', amount: 0, category: '', frequency: '', startDate: '', endDate: '' }]);
    const categories = [
        'Food', 'Transportation', 'Housing', 'Utilities', 'Insurance', 'Healthcare', 'Entertainment', 'Clothing', 'Miscellaneous'
    ];

    function addExpenseField() {
        expensesList.update(list => [...list, { name: '', amount: 0, category: '', frequency: '', startDate: '', endDate: '' }]);
    }

    async function submitAllExpenses() {
        const expenses = $expensesList;
        for (const expense of expenses) {
            await postExpense({
                Name: expense.name,
                Amount: parseInt(expense.amount, 10),
                Category: expense.category,
                Frequency: expense.frequency,
                Date: new Date(expense.startDate).toISOString(),
                // Additional logic here to handle endDate and frequency
            });
        }
        FetchExpenses(); // Refresh the expenses list
    }
</script>

<div class="form-container">
    <h2>Add Expenses</h2>
    {#each $expensesList as expense, index (index)}
        <div class="expense-item">
            <div class="form-control">
                <input placeholder="Name" type="text" bind:value={expense.name} />
                <input placeholder="Amount" type="number" bind:value={expense.amount} />
                <select bind:value={expense.category}>
                    <option value="">Select category</option>
                    {#each categories as cat}
                        <option value={cat}>{cat}</option>
                    {/each}
                </select>
                <select bind:value={expense.frequency}>
                    <option value="">Select frequency</option>
                    <option value="One-time">One-time</option>
                    <option value="Daily">Daily</option>
                    <option value="Weekly">Weekly</option>
                    <option value="Monthly">Monthly</option>
                    <option value="Annually">Annually</option>
                </select>
                <input type="date" bind:value={expense.startDate} />
                {#if !(expense.frequency === 'One-time' || expense.frequency === '')}
                    <input type="date" bind:value={expense.endDate} />
                {/if}
            </div>
        </div>
    {/each}
    <button on:click={addExpenseField}>Add Another Expense</button>
    <button on:click={submitAllExpenses}>Submit All</button>
</div>
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