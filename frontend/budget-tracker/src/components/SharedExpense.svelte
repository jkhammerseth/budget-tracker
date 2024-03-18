<script>
    import { postExpense } from '../routes/api/addExpense';
    import { FetchExpenses } from '../routes/api/fetchExpenses';

    let name = '';
    let amount = 0;
    let category = '';
    let frequency = 'One-time';
    let startDate = '';
    let endDate = '';
    let numberOfPeople = 1;
    let userShare = 0;

    const categories = [
        'Food', 'Transportation', 'Housing', 'Utilities', 'Insurance', 'Healthcare', 'Entertainment', 'Clothing', 'Miscellaneous'
    ];

    function calculateShare() {
        userShare = parseFloat((amount / numberOfPeople).toFixed(2));
    }

    async function addExpense() {
        calculateShare(); // Ensure userShare is updated before posting
        const expense = {
            Name: name + " (shared)",
            Amount: userShare,
            Category: category,
            Frequency: frequency,
            Date: new Date(startDate).toISOString()
        };
        await postExpense(expense);
        FetchExpenses(); // Refresh expenses after adding
        resetForm();
    }

    function resetForm() {
        name = '';
        amount = 0;
        category = '';
        frequency = 'One-time';
        startDate = '';
        endDate = '';
        numberOfPeople = 1;
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
    <h2>Shared Expense</h2>
    <div class="form-control">
        <input id="name" placeholder="Name" type="text" bind:value={name} />
    </div>
    <div class="form-control">
        <input id="amount" placeholder="Total amount" type="number" bind:value={amount} min="0" />
    </div>
    <div class="form-control">
        <input id="people" placeholder="Number of People" type="number" bind:value={numberOfPeople} min="1" on:input={calculateShare} />
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
            <option value="One-time">One-time</option>
            <!-- Additional options if needed -->
        </select>
    </div>
    <div class="form-control">
        <input id="date" type="date" bind:value={startDate} />
    </div>
    <!-- Only show end date for recurring expenses -->
    {#if frequency !== 'One-time'}
    <div class="form-control">
        <label for="endDate">End Date</label>
        <input id="endDate" type="date" bind:value={endDate} />
    </div>
    {/if}
    <div>Your share: {userShare} kr</div>
    <button on:click={addExpense}>Add Your Share</button>
</div>
