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
        case '':
          return $income;
        case 'range':
          return $income.filter(inc => {
            const incomeDate = new Date(inc.Date);
            return incomeDate >= $selectedStartDate && incomeDate <= $selectedEndDate;
          });
        default:
          return $income.filter(inc => {
            const incomeDate = new Date(inc.Date);
            return incomeDate.getFullYear() === $selectedYear && incomeDate.getMonth() === $selectedMonth;
          });
      }
    }
  );

  export const filteredPreviousIncome = derived(
    [income, selectionMode, selectedYear, selectedMonth, selectedStartDate, selectedEndDate],
    ([$income, $selectionMode, $selectedYear, $selectedMonth, $selectedStartDate, $selectedEndDate]) => {
      switch ($selectionMode) {
        case 'year':
          return $income.filter(income => {
            const incomeDate = new Date(income.Date);
            return incomeDate.getFullYear() === $selectedYear - 1;
          });
        case 'month':
          return $income.filter(income => {
            const incomeDate = new Date(income.Date);
            const previousMonth = $selectedMonth === 0 ? 11 : $selectedMonth - 1;
            const previousYear = $selectedMonth === 0 ? $selectedYear - 1 : $selectedYear;
            return incomeDate.getFullYear() === previousYear && incomeDate.getMonth() === previousMonth;
          });
        case 'range':
          return $income.filter(income => {
            const incomeDate = new Date(income.Date);
            const previousStartDate = new Date($selectedStartDate);
            const previousEndDate = new Date($selectedEndDate);
            previousStartDate.setDate(previousStartDate.getDate() - ($selectedEndDate - $selectedStartDate + 1));
            previousEndDate.setDate(previousEndDate.getDate() - ($selectedEndDate - $selectedStartDate + 1));
            return incomeDate >= previousStartDate && incomeDate <= previousEndDate;
          });
        default:
          return [];
      }
    }
  );

  const totalIncome = derived(filteredIncome, $filteredIncome => {
    return $filteredIncome.reduce((total, income) => total + income.Amount, 0);
  });
  
  export const totalPreviousIncome = derived(filteredPreviousIncome, $filteredPreviousIncome => {
    return $filteredPreviousIncome.reduce((total, income) => total + income.Amount, 0);
  });
  

  export const incomeChange = derived(
    [totalIncome, totalPreviousIncome],
    ([$totalIncome, $totalPreviousIncome]) => {
      if ($totalPreviousIncome === 0) return 0; // Avoid division by zero
      return (($totalIncome - $totalPreviousIncome) / $totalPreviousIncome) * 100;
    }
  );