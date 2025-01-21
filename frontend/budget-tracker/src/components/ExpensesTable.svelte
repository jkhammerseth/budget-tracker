<script>
  import { filteredExpenses } from '../stores/filteredExpenses';
  import { writable, derived } from 'svelte/store';
  import { activeModal } from '../stores/activeModal';
  import ExpenseModal from './modals/ExpenseModal.svelte';
  import ExpenseStatusButton from './ui/ExpenseStatusButton.svelte';
  import { fromISOString, formatExpenseAmount } from '../utility/functions';
  import { getIconComponent } from '../utility/icons';
  import { Calendar } from 'lucide-svelte';

  let selectedExpense = null;

  export const sortKey = writable(null); // Initial sort key
  export const sortOrder = writable(null); // Initial sort order

const filteredSortedExpenses = derived(
  [filteredExpenses, sortKey, sortOrder],
  ([$filteredExpenses, $sortKey, $sortOrder]) => {
    if (!$sortKey || !$sortOrder) return $filteredExpenses;

    return [...$filteredExpenses].sort((a, b) => {
      let comparison = 0;

      if ($sortKey === 'category') {
        comparison = a.category.name.localeCompare(b.category.name);
      } else if ($sortKey === 'name') {
        comparison = a.name.localeCompare(b.name);
      } else if ($sortKey === 'date') {
        comparison = new Date(a.payment_date) - new Date(b.payment_date);
      } else if ($sortKey === 'amount') {
        comparison = a.amount - b.amount;
      } else if ($sortKey === 'status') {
        comparison = a.status.localeCompare(b.status);
      }

      return $sortOrder === 'asc' ? comparison : -comparison;
    });
  }
);


function toggleSort(key) {
  sortKey.update((currentKey) => {
    if (currentKey === key) {
      sortOrder.update((currentOrder) => {
        if (currentOrder === 'asc') return 'desc';
        if (currentOrder === 'desc') {
          sortKey.set(null); // Reset sort
          return null;
        }
        return 'asc';
      });
    } else {
      sortKey.set(key);
      sortOrder.set('asc'); // Default to ascending
    }
    return key;
  });
}


  function openExpenseModal(expense) {
    selectedExpense = expense;
    activeModal.set('expense');
  }

  function closeExpenseModal() {
    selectedExpense = null;
    activeModal.set(null);
  }

  function daysUntil(dateString) {
    const today = new Date();
    const targetDate = new Date(dateString);

    const diffTime = targetDate - today;
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    return diffDays >= 0 ? `${diffDays}d left` : `${Math.abs(diffDays)}d ago`;
  }

</script>

<div class="container">
  {#if $filteredExpenses === undefined}
    <p>Loading expenses...</p>
  {:else if ($filteredExpenses.length === 0)}
    <p>No expenses found.</p>
  {:else}
    <div class="table-wrapper">
      <table>
        <thead>
          <tr>
            <th on:click={() => toggleSort('category')}>
              Category 
              {#if $sortKey === 'category'}
              <span class="arrows">
                {#if $sortOrder === 'asc'}
                  <svelte:component this={getIconComponent('ArrowUp')} size="20" />
                {:else}
                  <svelte:component this={getIconComponent('ArrowDown')} size="20" />
                {/if}
              </span>
              {/if}
            </th>
            <th on:click={() => toggleSort('name')}>
              Name {$sortKey === 'name' ? ($sortOrder === 'asc' ? '▲' : '▼') : ''}
            </th>
            <th on:click={() => toggleSort('status')}>
              Status {$sortKey === 'status' ? ($sortOrder === 'asc' ? '▲' : '▼') : ''}
            </th>
            <th on:click={() => toggleSort('date')} class="text-right">
              Date {$sortKey === 'date' ? ($sortOrder === 'asc' ? '▲' : '▼') : ''}
            </th>
            <th on:click={() => toggleSort('amount')} class="text-right">
              Amount {$sortKey === 'amount' ? ($sortOrder === 'asc' ? '▲' : '▼') : ''}
            </th>
          </tr>
        </thead>
        
        <tbody>
          {#each $filteredSortedExpenses as expense}
            <tr on:click={() => openExpenseModal(expense)}>
              <td class="category-cell">
                <div class="category-content">
                  {#if getIconComponent(expense.category.name)}
                    <span class="category-icon">
                      <svelte:component this={getIconComponent(expense.category.name)} size="20" />
                    </span>
                  {:else}
                    📁
                  {/if}
                  <span class="category-name">{expense.category.name}</span>
                </div>
              </td>
              <td>{expense.name}</td>
              <td><ExpenseStatusButton {expense} /></td>
              <td class="text-right">
                <div class="date-in-table">
                  <span class="date-logo"><Calendar size="20" /></span>
                  <div class="date-info">
                    <span class="date">{fromISOString(expense.payment_date)}</span>
                    <span class="days-until">{daysUntil(expense.payment_date)}</span>
                  </div>
                </div>
              </td>
              <td class="text-right-number">{formatExpenseAmount(expense.amount)}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}

  {#if $activeModal === 'expense' && selectedExpense}
    <ExpenseModal expense={selectedExpense} on:close={closeExpenseModal} />
  {/if}
</div>

<style>
  .container {
    width: 100%;
  }

  .table-wrapper {
    overflow-x: auto;
  }
  .name,
  .date {
    font-weight: 450;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    background-color: #ffffff;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
  }

  th, td {
    padding: 12px 16px;
    text-align: left;
    border-bottom: 1px solid #eee;
  }

  .arrows {
    margin-top: 1rem;
  }

  th {
    background-color: #F9FAFB; 
    height: 2.5rem;
    border-bottom: 1px solid #E5E7EB;
  }

  td.text-right-number,
  th.text-right, td.text-right {
    text-align: right;
  }

  tr:nth-child(odd) {
    background-color: #f9f9f9;
  }

  tr:hover {
    background-color: #f1f1f1;
  }
  .date-info {
    display: flex;
    flex-direction: column;
    gap: 4px;
    text-align: left;
    align-items: right;
  }

  .text-right {
    text-align: right;
    align-items: right;
  }

  .date-in-table {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-end;
  }

  .category-cell {
  display: flex;
  align-items: center;
  gap: 8px; /* Space between icon and name */
}

  .category-content {
    display: flex;
    align-items: center;
    justify-items: space-between;
  }

  .category-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.2rem;
  }

  .category-name {
    font-size: 0.95rem; /* Adjust text size as needed */
    color: #333;
    font-weight: 450;
  }


  .days-until {
    font-size: 0.85rem;
    color: #777;
  }

  .date-logo {
    margin-right: 5px;
  }

  /* Media Query for smaller screens */
  @media (max-width: 768px) {
    table {
      display: block;
      overflow-x: auto;
    }

    th, td {
      padding: 8px 12px;
      font-size: 12px; /* Reduce font size on smaller screens */
    }

    .category-icon {
      font-size: 1.5rem; /* Reduce icon size on smaller screens */
    }

    .days-until {
      font-size: 0.75rem; /* Smaller text for 'days left' on mobile */
    }
  }
</style>
