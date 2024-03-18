<script>
  import { fetchIncome } from '../routes/api/fetchIncome';
  import { postIncome } from '../routes/api/addIncome';

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

  //#TODO: Support for daily, weekly, and annual income
  async function addIncome() {
    const start = new Date(startDate);
    const end = new Date(endDate);

    if (frequency === 'One-time') {
      const income = {
        Name: name,
        Amount: parseInt(amount, 10),
        Category: category,
        Frequency: frequency,
        Date: start.toISOString()
      };
      postIncome(income);
      fetchIncome();
    } else {
      const monthDiff = (end.getFullYear() - start.getFullYear()) * 12 + end.getMonth() - start.getMonth() + 1;

      for (let i = 0; i < monthDiff; i++) {
          const incomeDate = new Date(start.getFullYear(), start.getMonth() + i, start.getDate());

          const income = {
              Name: name,
              Amount: parseInt(amount, 10),
              Category: category,
              Frequency: frequency,
              Date: incomeDate.toISOString()
          };
          postIncome(income);
        }
    
        name = '';
        amount = 0;
        category = '';
        frequency = '';
        startDate = '';
        endDate = '';

        fetchIncome();
      }
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
</div>