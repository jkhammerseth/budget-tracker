<script>
    import { onMount } from 'svelte';
    import { loans } from '../stores/loans'; 
    import { fetchLoans } from '../routes/api/fetchLoans'; 
    import { expenses } from '../stores/expenses'; 
    import { FetchExpenses } from '../routes/api/fetchExpenses';
  import { style } from 'd3';
    onMount(() => {
      fetchLoans();
      FetchExpenses();
    });
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
          <span id="loan-remaining">{loan.Remaining} kr</span>
      </div>
      <div class="loan-detail">
          <label for="loan-status">Status:</label>
          <span id="loan-status">{loan.Status}</span>
      </div>
        </div>
        <div class="linked-expenses">
          <div class="expenses-list-container">
            <h4>Linked Expenses</h4>
            {#if $expenses.filter(exp => exp.LoanID === loan.ID).length === 0}
              <p>No expenses this date.</p>
            {:else}
              <ul class="expenses-list">
                {#each $expenses.filter(exp => exp.LoanID === loan.ID) as expense}
                  <li class="expense-item">{expense.Name} - {expense.Amount} kr</li>
                {/each}
              </ul>
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
  
  .loan-detail, .expenses-list-container {
    display: flex;
    justify-content: space-between;
  }
  
  .loan-detail label, .expenses-list-container h4 {
    font-weight: bold;
  }
  
  .expenses-list {
    margin-top: 10px;
    list-style: none;
    padding-left: 0;
  }
  
  .expense-item {
    padding: 5px 0;
    border-bottom: 1px solid #ccc;
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
  
  </style>