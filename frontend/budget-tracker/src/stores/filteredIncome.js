import { writable, derived } from 'svelte/store';
import { income } from './income';
import { selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate } from './selectionMode';

export const filteredIncome = derived(
    [income, selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate],
    ([$income, $selectionMode, $selectedYear, $selectedMonth, $selectedStartDate, $selectedEndDate]) => {
      // Implement filtering logic based on the selectionMode
      switch ($selectionMode) {
        case 'year':
          return $income.filter(inc => new Date(inc.Date).getFullYear() === $selectedYear);
        case 'month':
          return $income.filter(inc => {
            const incomeDate = new Date(inc.Date);
            return incomeDate.getFullYear() === $selectedYear && incomeDate.getMonth() === $selectedMonth;
          });
        case 'range':
          return $income.filter(inc => {
            const incomeDate = new Date(inc.Date);
            return incomeDate >= $selectedStartDate && incomeDate <= $selectedEndDate;
          });
        default:
          return $income;
      }
    }
  );