<script>
  import { fetchIncome } from '../routes/api/fetchIncome';
  import { postIncome, postListofIncomes } from '../routes/api/addIncome';
  import Toast from './Toast.svelte';

  let componentView = 'single';

  let name = '';
  let amount = 0;
  let category = '';
  let frequency = '';
  let startDate = '';
  let endDate = '';

  const categories = [
    'Salary',
    'Investment',
    'Rental',
    'Interest',
    'Dividends',
    'Business',
    'Other'
  ];

  let message = '';
  let theme = 'info'; // Possible values: 'success', 'error', 'info'
  let duration = 3000;

  async function addIncome() {
    const start = new Date(startDate); 
    let end;
    if (endDate) {
        end = new Date(endDate);
    }

    const incomes = [];

    if (frequency === 'One-time') {
        incomes.push(createIncomeObject(start));
    } else if (end) { 
        let currentDate = new Date(start.getTime());

        while (currentDate <= end) {
            incomes.push(createIncomeObject(currentDate));
            incrementDate(currentDate, frequency);
        }
    }

    try {
        let response;
        if (incomes.length === 1) {
            response = await postIncome(incomes[0]);
        } else {
            response = await postListofIncomes(incomes);
        }

        handleResponse(response);
    } catch (error) {
        console.error("Error adding income:", error);
        message = "Error adding income";
        theme = "error";
    }

    resetFormFields();
    fetchIncome();
}

function createIncomeObject(date) {
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
        throw new Error("Failed to add income");
    }
    message = "Income added successfully!";
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
    width: 18rem;
    margin: auto;
    padding: 20px;
    background-color: var(--component-bg-color);
    border-style: solid;
    border-color: var(--component-border-color);
    border-width: 1px;
    border-radius: var(--component-border-radius);
    box-shadow: var(--component-box-shadow);
    font-family: var(--font-family);
  }

  h2 {
    color: var(--text-color);
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

  @media (max-width: 480px) {
      .form-container {
          padding: 10px;
      }
  }
</style>

<div class="form-container">
<h2>Add Income</h2>
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
<button on:click={addIncome}>Add Income</button>
<Toast {message} {theme} {duration} />
</div>