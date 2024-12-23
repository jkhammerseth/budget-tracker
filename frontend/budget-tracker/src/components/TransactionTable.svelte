<script>
    import { onMount } from 'svelte';
    import FaRegTrashAlt from 'svelte-icons/fa/FaRegTrashAlt.svelte'
    import { filteredExpenses } from '../stores/filteredExpenses';
    import { filteredIncome } from '../stores/filteredIncome.js';
    import { fetchIncome } from '../routes/api/fetchIncome';
    import { FetchExpenses } from '../routes/api/fetchExpenses';
    import { derived } from 'svelte/store';
    import ExpenseStatusButton from './ui/ExpenseStatusButton.svelte';
    import IncomeStatusButton from './ui/IncomeStatusButton.svelte';

    let nameFilter = '';
    let sortKey = 'date';
    let sortOrder = 'asc';

    const filteredSortedIncome = derived(filteredIncome, $filteredIncome => {
      return $filteredIncome
        .filter(income => income.Name.toLowerCase().includes(nameFilter.toLowerCase()))
        .sort((a, b) => {
          let comparison = 0;
          if (a[sortKey] < b[sortKey]) {
            comparison = -1;
          } else if (a[sortKey] > b[sortKey]) {
            comparison = 1;
          }
          return sortOrder === 'asc' ? comparison : -comparison;
        });
    });

    const filteredSortedExpenses = derived(filteredExpenses, $filteredExpenses => {
      return $filteredExpenses
        .filter(expense => expense.Name.toLowerCase().includes(nameFilter.toLowerCase()))
        .sort((a, b) => {
          let comparison = 0;
          if (a[sortKey] < b[sortKey]) {
            comparison = -1;
          } else if (a[sortKey] > b[sortKey]) {
            comparison = 1;
          }
          return sortOrder === 'asc' ? comparison : -comparison;
        });
  });

    const mergedData = derived(
        [filteredSortedIncome, filteredSortedExpenses],
        ([$filteredSortedIncome, $filteredSortedExpenses]) => {
            const merged = [...$filteredSortedIncome.map(item => ({...item, Type: 'Income'})),
                            ...$filteredSortedExpenses.map(item => ({...item, Type: 'Expense'}))];
            
            // Sort the merged array if necessary
            return merged.sort((a, b) => {
            // Example sort by date
            return new Date(a.Date) - new Date(b.Date);
            });
        }
        );
    
    onMount(() =>{
        fetchIncome();
        FetchExpenses();
    });


    function formatAmount(amount) {
      return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'NOK' }).format(amount);
    }

    function resetFilter() {
      nameFilter = '';
      sortKey = 'date';
      sortOrder = 'asc';
      FetchExpenses();
    }

    function fromISOString(isoString) {
    const date = new Date(isoString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0'); 
    const year = date.getFullYear().toString().substr(-2); 

    // Formatting the date as DD.MM.YY
    return `${day}.${month}.${year}`;
    }

  </script>

<div class="container">
  <div class="filter-sort-container">
    <input class="search-field" type="text" placeholder="Search by name..." bind:value="{nameFilter}">
    
    <select class="select-filter" bind:value="{sortKey}">
      <option value="amount">Sort by Amount</option>
      <option value="category">Sort by Category</option>
      <option value="date">Sort by Date</option>
    </select>
  
    <select class="filter-order" bind:value="{sortOrder}">
      <option value="asc">Ascending</option>
      <option value="desc">Descending</option>
    </select>

    <button class="apply-filter-button" on:click={FetchExpenses} title="Apply">
      <span>Apply</span>
    </button>
    <button class="icon-button" on:click={resetFilter} title="Remove filter">
      <span class="icon"><FaRegTrashAlt/></span>
  </div>
  {#if $mergedData.length > 0}
<table>
  <thead>
    <tr>
      <th>...</th>
      <th>Type</th> 
      <th>Name</th>
      <th>Amount</th>
      <th>Category</th>
      <th>Status</th>
      <th>Date</th>
      <th>Actions</th>
    </tr>
  </thead>
  <tbody>
    {#each $mergedData as item}
      <tr class="{item.Type === 'Expense' ? 'expense-row' : 'income-row'}">
        <td>
          <input class="checkbox" type="checkbox" />
        </td>
        <td>{item.Type}</td>
        <td>{item.Name}</td>
        <td>{formatAmount(item.Amount)}</td>
        <td>{item.Category}</td>
        <td>
          {#if item.Type === 'Expense'}
              <ExpenseStatusButton expense={item}/>
          {:else if item.Type === 'Income'}
            <IncomeStatusButton income={item}/>
          {/if}
        </td>
        <td>{fromISOString(item.PaymentDate)}</td>
      
      </tr>
    {/each}
  </tbody>
</table>
{:else}
<p>No data found.</p>
{/if}


  {#if $mergedData.length === 0}
    <p>No expenses found.</p>
  {/if}
</div>

  
  <style>
    .container {
      max-width: full-width;
      font-family: 'Roboto', sans-serif;
      border-radius: 20px;
      margin: 40px auto;
      margin-top: 20px;
    }

    .filter-sort-container {
      display: flex;
      align-items: center;
      gap: 10px;
      padding: 10px;
      background-color: #E8EFF1;
      border: black;
      border-style: solid;
      border-width: 1px;
      border-radius: 10px;
      border-bottom: none;
      border-bottom-left-radius: 0;
      border-bottom-right-radius: 0;

    }

    .search-field, .select-filter, .filter-order, .apply-filter-button, .icon-button {
      border: 1px solid #ccc; 
      padding: 8px 12px;
      border-radius: 4px;
      outline: none;
    }

    .search-field:focus, .select-filter:focus, .filter-order:focus {
      border-color: #3A87F2;
    }

    .apply-filter-button, .icon-button {
      cursor: pointer; 
      background-color: #3A87F2;
      color: white;
      border: none; 
      transition: background-color 0.2s;
    }

    .apply-filter-button:hover {
      background-color: #2a6db2;
    }

    input[type="text"],
    select {
      width: calc(100% - 20px);
      padding: 10px;
      margin-top: 5px;
      border: 1px solid #ccc;
      border-radius: 5px;
      box-sizing: border-box;
    }
  
    table {
      width: 100%;
      border-collapse: separate;
      border-spacing: 0;
      background-color: #E8EFF1;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      font-family: 'Roboto', sans-serif;
      overflow: hidden;
      border-style: solid;
      border-color: black;
      border-width: 1px;
    }
  
    th, td {
      padding: 12px 15px; 
      text-align: left;
      border-bottom: solid 1px #ddd;
    }
  
    th {
      background-color: #3A87F2;
      color: white;
      font-weight: normal;
    }
  
    tr:hover {
      background-color: #f5f5f5; 
    }
  
    tr:last-child td {
      border-bottom: none;
    }
    .icon-button {
      background: none;
      border: none;
      padding: 0;
      cursor: pointer;
      outline: inherit;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      height: 16px;
      width: 16px;
  }
  .icon {
      width: 16px; 
      height: 16px;
      color: #333; 
  }

  .icon:hover {
      color: #e74c3c;
  }

  .icon:active {
      transform: scale(0.9);
  }
  </style>
  
