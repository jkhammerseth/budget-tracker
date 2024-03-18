import { income } from '../../stores/income.js';

export async function fetchIncome() {
    console.log("Populating budget table");
    try {
      const inc_response = await fetch('http://localhost:8080/api/users/incomes', {
        credentials: 'include',
      });

      if (!inc_response.ok) {
        throw new Error(`HTTP error! status: ${inc_response.status}`);
      }

      const inc_data = await inc_response.json();
      const sortedIncome = inc_data.sort((a, b) => new Date(b.Date) - new Date(a.Date));
      income.set(sortedIncome);
    } catch (error) {
      console.error("Failed to fetch income:", error);
    }
  }
  