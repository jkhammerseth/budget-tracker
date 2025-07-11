<script>
  import { onMount } from 'svelte';
  import { fetchIncome } from '../routes/api/fetchIncome';
  import { derived } from 'svelte/store';
  import { fromISOString, formatAmount } from '../utility/functions'
  import { activeModal } from '../stores/activeModal';
  import { FontAwesomeIcon } from '@fortawesome/svelte-fontawesome';
  import { filteredIncome } from '../stores/filteredIncome';
  import { getIconComponent } from '../utility/icons'
  import { Calendar, Info } from 'lucide-svelte';
  import IncomeModal from './modals/IncomeModal.svelte';

  let selectedIncome = null;
  let selectedIncomes = new Set();
  let selectAll = false;

  onMount(fetchIncome);

  function openIncomeModal(income) {
    selectedIncome = income;
    console.log(income)
    activeModal.set('income');
  }

  function closeIncomeModal() {
    selectedIncome = null;
    activeModal.set(null); // Close modal by resetting activeModal state
  }

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

  function daysUntil(dateString) {
    const today = new Date();
    const targetDate = new Date(dateString);

    const diffTime = targetDate - today;
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24)); // Convert milliseconds to days
    return diffDays >= 0 ? `${diffDays}d left` : `${Math.abs(diffDays)}d ago`;
  }
  
  function toggleIncomeSelection(incomeId) {
    if (selectedIncomes.has(incomeId)) selectedIncomes.delete(incomeId);
    else selectedIncomes.add(incomeId);
  }

  function toggleSelectAll() {
    selectAll = !selectAll;
    selectedIncomes = new Set(selectAll ? $filteredIncome.map((e) => e.id) : []);
  }

</script>

<div class="container">
  {#if $filteredIncome === undefined}
    <p>Loading income...</p>
  {:else if $filteredIncome.length === 0}
    <p>No income found.</p>
  {:else}
    <div class="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>
              <input
              type="checkbox"
              bind:checked={selectAll}
              on:change={toggleSelectAll}
            /></th>
            <th>Category</th>
            <th>Name</th>
            <th class="text-right">Date</th>
            <th class="text-right">Amount</th>
          </tr>
        </thead>
        <tbody>
          {#each $filteredSortedIncome as income}
            <tr on:click={() => openIncomeModal(income)}>
              <td><input
                type="checkbox"
                checked={selectedIncomes.has(income.ID)}
                on:click|stopPropagation={() => toggleIncomeSelection(income.ID)}
              /></td>
              <td class="category-cell">
                <div class="category-content">
                  {#if getIconComponent(income.Category)}
                    <span class="category-icon">
                      <svelte:component this={getIconComponent(income.Category)} size="20" />
                    </span>
                  {:else}
                    📁
                  {/if}
                  <span class="category-name">{income.Category}</span>
                </div>
              </td>
              <td class="name">{income.Name}</td>
              <td class="text-right">
                <div class="date-in-table">
                  <span class="date-logo"><Calendar size="20" /></span>
                  <div class="date-info">
                    <span class="date">{fromISOString(income.Date)}</span>
                    <span class="days-until">{daysUntil(income.Date)}</span>
                  </div>
                </div>
              </td>
              <td class="text-right-number">{formatAmount(income.Amount)}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}

  {#if $activeModal === 'income' && selectedIncome}
  <IncomeModal income={selectedIncome} />
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

  .category-cell {
  display: flex;
  align-items: center;
  gap: 8px; /* Space between icon and name */
}

  .category-content {
    display: flex;
    align-items: center;
    justify-items: left;
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
