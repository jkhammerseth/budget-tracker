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