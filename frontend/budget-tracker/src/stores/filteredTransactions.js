import { derived } from 'svelte/store';
import { filteredIncome } from './filteredIncome'; // Ensure correct path
import { filteredExpenses } from './filteredExpenses'; // Ensure correct path

export const filteredTransactions = derived(
  [filteredIncome, filteredExpenses],
  ([$filteredIncome, $filteredExpenses]) => {
    // Combine the filtered income and expenses
    return [
      ...$filteredIncome.map(income => ({ ...income, type: 'income' })), // Tag incomes with type
      ...$filteredExpenses.map(expense => ({ ...expense, type: 'expense' })) // Tag expenses with type
    ];
  }
);

export const totalTransactions = derived(
  filteredTransactions,
  $filteredTransactions => {
    // Calculate the net total (income - expenses)
    return $filteredTransactions.reduce((total, transaction) => {
      return transaction.type === 'income'
        ? total + transaction.Amount
        : total - transaction.Amount;
    }, 0);
  }
);


export const recentTransactions = derived(
    filteredTransactions,
    $filteredTransactions => {
      // Filter only expenses where Paid is true
      const paidTransactions = $filteredTransactions.filter(transaction => transaction.Paid);
  
      // Sort the paid transactions by date (newest first)
      const sortedTransactions = paidTransactions.sort(
        (a, b) => new Date(b.Date || b.PaymentDate) - new Date(a.Date || a.PaymentDate)
      );
  
      const recentLimit = 3; 
      return sortedTransactions.slice(0, recentLimit);
    }
  );
  