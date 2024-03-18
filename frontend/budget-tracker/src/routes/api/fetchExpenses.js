import { expenses } from '../../stores/expenses.js';

export async function FetchExpenses() {
    console.log("Populating budget table");
    try {
      const exp_response = await fetch('http://localhost:8080/api/users/expenses', {
        credentials: 'include'
      });

      if (!exp_response.ok) {
        throw new Error(`HTTP error! status: ${exp_response.status}`);
      }

      const exp_data = await exp_response.json();
      const sortedExpenses = exp_data.sort((a, b) => new Date(b.Date) - new Date(a.Date));
      expenses.set(sortedExpenses);
      //expenses.set(exp_data);
    } catch (error) {
      console.error("Failed to fetch expenses:", error);
    }
  }
  