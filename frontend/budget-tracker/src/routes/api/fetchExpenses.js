import { expenses } from '../../stores/expenses.js';

let currentPage = 0; // Tracks the current page
const limit = 50;    // Number of expenses per page

export async function FetchExpensesPaginated({ append = false, mode, year, month }) {
  try {
    const url = `/api/expenses?mode=${mode}&year=${year}&month=${month}&page=${currentPage}&limit=${limit}`;
    console.log(url);
    const response = await fetch(url, { credentials: 'include' });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();

    if (data.expenses.length) {
      if (append) {
        expenses.update(existing => [...existing, ...data.expenses]);
      } else {
        expenses.set(data.expenses);
      }
      currentPage += 1; // Increment the page only if data is fetched successfully
    }

    return data.hasMore; // Backend should indicate if more data exists
  } catch (error) {
    console.error("Error fetching paginated expenses:", error);
    return false;
  }
}




export async function FetchExpenses() {
  try {
    const exp_response = await fetch('http://localhost:8080/api/users/allexpenses', {
      credentials: 'include'
    });

    if (!exp_response.ok) {
      throw new Error(`HTTP error! status: ${exp_response.status}`);
    }

    const exp_data = await exp_response.json();
    const sortedExpenses = exp_data.sort((a, b) => new Date(b.Date) - new Date(a.Date));
    console.log("fetched expenses");
    expenses.set(sortedExpenses); // Replace existing expenses with all expenses
  } catch (error) {
    console.error("Failed to fetch all expenses:", error);
  }
}
