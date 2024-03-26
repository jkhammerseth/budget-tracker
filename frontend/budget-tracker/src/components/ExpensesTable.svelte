<script>
  import { filteredExpenses } from '../stores/filteredExpenses';
  import { onMount } from 'svelte';
  import FaRegTrashAlt from 'svelte-icons/fa/FaRegTrashAlt.svelte';
  import { FetchExpenses } from '../routes/api/fetchExpenses';
  import { derived } from 'svelte/store';
  import ExpenseStatusButton from './ui/ExpenseStatusButton.svelte';
  import ExpenseActionMenu from './ExpenseActionMenu.svelte';

  function toggleMenu(expense) {
    menuStates[expense] = !menuStates[expense];
    menuStates = { ...menuStates };
  }

  let menuStates = {};

  let nameFilter = '';
  let sortKey = 'date';
  let sortOrder = 'asc';

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

  onMount(FetchExpenses);

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
  {#if $filteredExpenses.length > 0}
    <table>
      <thead>
        <tr>
          <th>...</th>
          <th>Name</th>
          <th>Amount</th>
          <th>Category</th>
          <th>Status</th>
          <th>Date</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {#each $filteredSortedExpenses as expense}
          <tr>
            <td>
              <input class="checkbox" type="checkbox" />
            </td>
            <td>
              {expense.Name}
            </td>
            <td>{formatAmount(expense.Amount)}</td>
            <td>{expense.Category}</td>
            <td>
              <ExpenseStatusButton {expense}/>
            </td>
            <td>{fromISOString(expense.Date)}</td>
            <td>
              <button class="dots" on:click={() => toggleMenu(expense)}>
                <span class="icon">...</span>
              </button>
              {#if menuStates[expense]}
                <ExpenseActionMenu expense={expense} />
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}

  {#if $filteredExpenses.length === 0}
    <p>No expenses found.</p>
  {/if}
</div>

  
  <style>
    .container {
      max-width: 100%;
      font-family: var(--font-family);
      border-radius: var(--component-border-radius);

    }

    .filter-sort-container {
      display: flex;
      align-items: center;
      width: 41.5rem;
      gap: 10px;
      padding: 10px;
      background-color: var(--component-bg-color);
      border: var(--component-border-color);
      border-style: solid;
      border-width: 1px;
      border-radius: var(--component-border-radius);
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
      background-color: var(--primary-button-color);
      color: var(--primary-button-text-color);
      border: none; 
      transition: background-color 0.2s;
    }

    .apply-filter-button:hover {
      background-color: var(--primary-button-hover-color);
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

    .checkbox {
      width: 16px;
      height: 16px;
    }
  
    table {
      width: 42.86rem;
      border-collapse: separate;
      border-spacing: 0;
      background-color: var(--component-bg-color);
      box-shadow: var(--component-box-shadow);
      font-family: var(--font-family);
      overflow: hidden;
      border-style: solid;
      border-color: var(--component-border-color);
      border-width: 1px;
    }
  
    th, td {
      padding: 10px 12px; 
      text-align: left;
      border-bottom: solid 1px #ddd;
    }
  
    th {
      background-color: var(--primary-color);
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
      height: 46px;
      width: 46px;
  }

  .dots {
      background: none;
      border: none;
      padding: 2px;
      cursor: pointer;
      outline: inherit;
      display: inline-flex;
      align-items: center;
      justify-content: center;

  }

  .icon {
      font-size: 16px;
      line-height: 1.5;
      color: black; 
  }

  .icon:hover {
      color: #3A87F2;
  }

  .icon:active {
      transform: scale(0.9);
  }
  </style>
  
