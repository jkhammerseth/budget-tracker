import { writable, derived } from 'svelte/store';
import { expenses } from './expenses';
import { selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate } from './selectionMode';

export const filteredExpenses = derived(
    [expenses, selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate],
    ([$expenses, $selectionMode, $selectedYear, $selectedMonth, $selectedStartDate, $selectedEndDate]) => {
      // Implement filtering logic based on the selectionMode
      switch ($selectionMode) {
        case 'year':
          return $expenses.filter(expense => new Date(expense.PaymentDate).getFullYear() === $selectedYear);
        case 'month':
          return $expenses.filter(expense => {
            const expenseDate = new Date(expense.PaymentDate);
            return expenseDate.getFullYear() === $selectedYear && expenseDate.getMonth() === $selectedMonth;
          });
        case 'range':
          return $expenses.filter(expense => {
            const expenseDate = new Date(expense.PaymentDate);
            return expenseDate >= $selectedStartDate && expenseDate <= $selectedEndDate;
          });
        default:
          return $expenses;
      }
    }
  );

  export const filteredPreviousExpenses = derived(
    [expenses, selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate],
    ([$expenses, $selectionMode, $selectedYear, $selectedMonth, $selectedStartDate, $selectedEndDate]) => {
      let filtered = [];
      switch ($selectionMode) {
        case 'year':
          filtered = $expenses.filter(expense => {
            const expenseDate = new Date(expense.PaymentDate);
            return expenseDate.getFullYear() === $selectedYear - 1;
          });
          break;
        case 'month':
          filtered = $expenses.filter(expense => {
            const expenseDate = new Date(expense.PaymentDate);
            const previousMonth = $selectedMonth === 0 ? 11 : $selectedMonth - 1;
            const previousYear = $selectedMonth === 0 ? $selectedYear - 1 : $selectedYear;
            return expenseDate.getFullYear() === previousYear && expenseDate.getMonth() === previousMonth;
          });
          break;
        case 'range':
          filtered = $expenses.filter(expense => {
            const expenseDate = new Date(expense.PaymentDate);
            const previousStartDate = new Date($selectedStartDate);
            const previousEndDate = new Date($selectedEndDate);
            previousStartDate.setDate(previousStartDate.getDate() - ($selectedEndDate - $selectedStartDate + 1));
            previousEndDate.setDate(previousEndDate.getDate() - ($selectedEndDate - $selectedStartDate + 1));
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
    return $filteredExpenses.reduce((total, expense) => total + expense.Amount, 0);
  });
  
  export const totalPreviousExpenses = derived(filteredPreviousExpenses, $filteredPreviousExpenses => {
    const total = $filteredPreviousExpenses.reduce((total, expense) => total + expense.Amount, 0);
    return total;
  });
  

  
  export const expensesChange = derived(
    [totalExpenses, totalPreviousExpenses],
    ([$totalExpenses, $totalPreviousExpenses]) => {
      if ($totalPreviousExpenses === 0) return 0; // Avoid division by zero
      return (($totalExpenses - $totalPreviousExpenses) / $totalPreviousExpenses) * 100;
    }
  );

