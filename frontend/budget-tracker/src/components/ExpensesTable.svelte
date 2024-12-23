<script>
  import { filteredExpenses } from '../stores/filteredExpenses';
  import { onMount } from 'svelte';
  import { FetchExpenses } from '../routes/api/fetchExpenses';
  import { derived } from 'svelte/store';
  import ExpenseStatusButton from './ui/ExpenseStatusButton.svelte';
  import ExpenseModal from './modals/ExpenseModal.svelte';
  import { fromISOString, formatExpenseAmount } from '../utility/functions'
  import { activeModal } from '../stores/activeModal';

  import { faUtensils, faCar, faHome, faBolt, faFileShield, faHeartbeat, faFilm, faTshirt, faBoxOpen, faCalendar } from '@fortawesome/free-solid-svg-icons';
  import { FontAwesomeIcon } from '@fortawesome/svelte-fontawesome';

  let selectedExpense = null;

  const categoryIcons = {
    Food: faUtensils,
    Transportation: faCar,
    Housing: faHome,
    Utilities:  faBolt, 
    Insurance: faFileShield, 
    Healthcare: faHeartbeat, 
    Entertainment: faFilm, 
    Clothing: faTshirt, 
    Miscellaneous: faBoxOpen,
  };

  function getCategoryIcon(category) {
    return categoryIcons[category] || '📁'; // Default icon if category not found
  }

  function openExpenseModal(expense) {
    selectedExpense = expense;
    activeModal.set('expense'); // Make sure 'expense' is set to trigger the modal
  }

  function closeExpenseModal() {
    selectedExpense = null;
    activeModal.set(null); // Close modal by resetting activeModal state
  }

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

  function daysUntil(dateString) {
    const today = new Date();
    const targetDate = new Date(dateString);

    const diffTime = targetDate - today;
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24)); // Convert milliseconds to days
    return diffDays >= 0 ? `${diffDays}d left` : `${Math.abs(diffDays)}d ago`;
  }
</script>

<div class="container">
  {#if $filteredExpenses === undefined}
    <p>Loading expenses...</p>
  {:else if $filteredExpenses.length === 0}
    <p>No expenses found.</p>
  {:else}
    <div class="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>Category</th>
            <th>Name</th>
            <th>Status</th>
            <th class="text-right">Date</th>
            <th class="text-right">Amount</th>
          </tr>
        </thead>
        <tbody>
          {#each $filteredSortedExpenses as expense}
            <tr on:click={() => openExpenseModal(expense)}>
              <td >
                <span class="category-icon"><FontAwesomeIcon icon={getCategoryIcon(expense.Category)}/></span>
                <span>{expense.Category}</span>
              </td>
              <td class="name">{expense.Name}</td>
              <td><ExpenseStatusButton {expense} /></td>
              <td class="text-right">
                <div class="date-in-table">
                  <span class="date-logo"><FontAwesomeIcon icon={faCalendar}/></span>
                  <div class="date-info">
                    <span class="date">{fromISOString(expense.PaymentDate)}</span>
                    <span class="days-until">{daysUntil(expense.PaymentDate)}</span>
                  </div>
                </div>
              </td>
              <td class="text-right-number">{formatExpenseAmount(expense.Amount)}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}

  {#if $activeModal === 'expense' && selectedExpense}
    <ExpenseModal 
      expense={selectedExpense} 
    />
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

  .category-icon {
    font-size: 1.5rem;
    margin-right: 12px;
  }

  .days-until {
    font-size: 0.85rem;
    color: #777;
  }

  .date-logo {
    margin-right: 10px;
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
