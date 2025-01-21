import { writable, derived } from 'svelte/store';
import { expenses } from './expenses';
import { selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate } from './selectionMode';
import { searchQuery } from './searchQuery';

export const filteredExpenses = derived(
  [expenses, selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate, searchQuery],
  ([
    $expenses,
    $selectionMode,
    $selectedYear,
    $selectedMonth,
    $selectedStartDate,
    $selectedEndDate,
    $searchQuery
  ]) => {
    // First, filter based on selectionMode
    let filtered = [];
    switch ($selectionMode) {
      case 'year':
        filtered = $expenses.filter(
          expense => new Date(expense.payment_date).getFullYear() === $selectedYear
        );
        break;
      case '':
        filtered = $expenses;
        break;
      case 'range':
        filtered = $expenses.filter(expense => {
          const expenseDate = new Date(expense.payment_date);
          return expenseDate >= $selectedStartDate && expenseDate <= $selectedEndDate;
        });
        break;
      default:
        filtered = $expenses.filter(expense => {
          const expenseDate = new Date(expense.payment_date);
          return (
            expenseDate.getFullYear() === $selectedYear &&
            expenseDate.getMonth() === $selectedMonth
          );
        });
        break;
    }

    // Then, apply the search query
    const lowerCaseQuery = $searchQuery.toLowerCase().trim();
    if (lowerCaseQuery) {
      filtered = filtered.filter(expense =>
        expense.name.toLowerCase().includes(lowerCaseQuery)
      );
    }

    return filtered;
  }
);

export const filteredPreviousExpenses = derived(
  [expenses, selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate],
  ([$expenses, $selectionMode, $selectedYear, $selectedMonth, $selectedStartDate, $selectedEndDate]) => {
    let filtered = [];
    switch ($selectionMode) {
      case 'year':
        filtered = $expenses.filter(expense => {
          const expenseDate = new Date(expense.payment_date);
          return expenseDate.getFullYear() === $selectedYear - 1;
        });
        break;
      case 'month':
        filtered = $expenses.filter(expense => {
          const expenseDate = new Date(expense.payment_date);
          const previousMonth = $selectedMonth === 0 ? 11 : $selectedMonth - 1;
          const previousYear = $selectedMonth === 0 ? $selectedYear - 1 : $selectedYear;
          return (
            expenseDate.getFullYear() === previousYear &&
            expenseDate.getMonth() === previousMonth
          );
        });
        break;
      case 'range':
        filtered = $expenses.filter(expense => {
          const expenseDate = new Date(expense.payment_date);
          const previousStartDate = new Date($selectedStartDate);
          const previousEndDate = new Date($selectedEndDate);
          previousStartDate.setDate(
            previousStartDate.getDate() - ($selectedEndDate - $selectedStartDate + 1)
          );
          previousEndDate.setDate(
            previousEndDate.getDate() - ($selectedEndDate - $selectedStartDate + 1)
          );
          return expenseDate >= previousStartDate && expenseDate <= previousEndDate;
        });
        break;
      default:
        filtered = [];
    }
    return filtered;
  }
);

const totalExpenses = derived(filteredExpenses, $filteredExpenses => {
  return $filteredExpenses.reduce((total, expense) => total + expense.amount, 0);
});

export const totalPreviousExpenses = derived(filteredPreviousExpenses, $filteredPreviousExpenses => {
  const total = $filteredPreviousExpenses.reduce((total, expense) => total + expense.amount, 0);
  return total;
});

export const expensesChange = derived(
  [totalExpenses, totalPreviousExpenses],
  ([$totalExpenses, $totalPreviousExpenses]) => {
    if ($totalPreviousExpenses === 0) return 0; // Avoid division by zero
    return (($totalExpenses - $totalPreviousExpenses) / $totalPreviousExpenses) * 100;
  }
);
