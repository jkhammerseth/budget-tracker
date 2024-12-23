<script>
  import { onMount } from 'svelte';
  import { loans } from '../stores/loans'; 
  import { fetchLoans } from '../routes/api/fetchLoans'; 
  import { expenses } from '../stores/expenses'; 
  import { FetchExpenses } from '../routes/api/fetchExpenses';

  // Function to get the next upcoming expense for a loan
  const getUpcomingExpense = (loanID) => {
      const sortedExpenses = $expenses
          .filter(exp => exp.LoanID === loanID && !exp.Paid)
          .sort((a, b) => a.StartDate - b.StartDate); // Sort by date ascending
      return sortedExpenses.length > 0 ? sortedExpenses[0] : null;
  };

  function paidExpenses(loanID) {
    const paidExpensesList = $expenses.filter(exp => exp.LoanID === loanID && exp.Paid);
    return paidExpensesList.reduce((total, expense) => total + expense.Amount, 0);
}

  onMount(() => {
      fetchLoans();
      FetchExpenses();
  });

  function fromISOString(isoString) {
    const date = new Date(isoString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0'); 
    const year = date.getFullYear().toString().substr(-2); 
    return `${day}.${month}.${year}`;
  }

  function calculateRemaining(loan) {
    const paidSum = paidExpenses(loan.ID);
    return loan.Amount - paidSum;
  }
  
</script>

<div class="loans-container">
  {#each $loans as loan (loan.ID)}
      <div class="loan-expense-container">
          <div class="loan-card">
              <div class="loan-title">{loan.Type} Loan</div>
              <hr>
              <div class="loan-detail">
                  <label for="loan-amount">Amount:</label>
                  <span id="loan-amount">{loan.Amount} kr</span>
              </div>
              <div class="loan-detail">
                  <label for="loan-interest">Interest:</label>
                  <span id="loan-interest">{loan.Interest}%</span>
              </div>
              <div class="loan-detail">
                  <label for="loan-monthly-payment">Monthly Payment:</label>
                  <span id="loan-monthly-payment">{loan.MonthlyPayment} kr</span>
              </div>
              <div class="loan-detail">
                  <label for="loan-remaining">Remaining:</label>
                  <span id="loan-remaining">{calculateRemaining(loan)} kr</span>
              </div>
              <div class="loan-detail">
                  <label for="loan-status">Status:</label>
                  <span id="loan-status">{loan.Status}</span>
              </div>
          </div>

          <div class="linked-expenses">
              <div class="upcoming-expense-container">
                  <h4>Upcoming Expense</h4>
                  {#if getUpcomingExpense(loan.ID)}
                      <div class="expense-item">
                          <span>{getUpcomingExpense(loan.ID).Name}</span> - 
                          <span>{getUpcomingExpense(loan.ID).Amount} kr</span>
                          <span>{fromISOString(getUpcomingExpense(loan.ID).PaymentDate)}</span>
                      </div>
                  {:else}
                      <p>No upcoming expense.</p>
                  {/if}
              </div>

              <div class="button-row">
                  <button>View Details</button>
                  <button>Upload payment plan</button>
                  <button>Edit loan details</button>
              </div>
          </div>
      </div>
  {/each}
</div>

<style>
  .button-row {
      display: flex;
      gap: 10px;
  }

  .loans-container {
      display: flex;
      flex-direction: column;
      gap: 20px;
      padding: 20px;
  }

  .loan-expense-container {
      display: flex;
      background-color: #E8EFF1;
      border: 1px solid black;
      border-radius: 10px;
  }

  .loan-card, .linked-expenses {
      flex: 1;
      padding: 20px;
  }

  .loan-title {
      font-size: 1.2em;
      font-weight: bold;
  }

  .loan-detail, .upcoming-expense-container {
      display: flex;
      justify-content: space-between;
      margin-bottom: 10px;
  }

  .loan-detail label, .upcoming-expense-container h4 {
      font-weight: bold;
  }

  .expense-item {
      margin-top: 10px;
  }

  button {
      display: block;
      width: 100%;
      padding: 10px;
      margin-top: 20px;
      background-color: #3A87F2;
      color: white;
      border: none;
      border-radius: 5px;
      cursor: pointer;
      transition: background-color 0.3s;
  }

  button:hover {
      background-color: #2C6FBF;
  }
</style>
