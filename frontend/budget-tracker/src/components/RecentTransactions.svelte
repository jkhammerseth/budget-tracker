<script>
  import { formatAmount, formatExpenseAmount, fromISOString, daysFromNow } from "../utility/functions";

  import { recentTransactions } from '../stores/filteredTransactions'
  import { Info } from 'lucide-svelte';
</script>

<div class="recent-transactions">
    <div class="upcoming-header">
      <h2>Recent Transactions</h2>
      <span class="calendar-icon"><Info size="24" /></span>
    </div>
    <div class="transactions-list">
      {#if $recentTransactions.length > 0}
        {#each $recentTransactions as transaction}
          <div class="transaction">
            <div>
              <p class="transaction-name">
                {transaction.Name}
              </p>
              <p class="transaction-date">{fromISOString(transaction.PaymentDate || transaction.Date)}</p>
            </div>
            <p class="transaction-amount {transaction.type === 'income' ? 'income' : 'expense'}">
              {transaction.type === 'income' 
                ? formatAmount(transaction.Amount) 
                : formatExpenseAmount(transaction.Amount)}
            </p>
    </div>
        {/each}
        <div class="view">
          <a class="view-link" href="/" aria-label="View all transactions">View All Transactions</a>
        </div>
      {:else}
        <p>No recent transactions to display.</p>
      {/if}
    </div>
  </div>

  <style>
    .view {
    display: flex;
    justify-content: center;
}

  .view-link {
    text-decoration: none;
    font-size: 1.1rem;
    padding: 16px;
}

    .view-link:hover {
    color: #4a86ff;
    text-decoration: none;
}

.upcoming-header {
    display: flex;
    justify-content: space-between;
}

  .upcoming-header h2 {
  margin: 0;
  }

  .upcoming-header .calendar-icon {
    display: flex;
    align-items: center;
  }


  .calendar-icon {
    padding: 16px;
  }


  .recent-transactions {
  background: #fff;
  padding: 16px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.recent-transactions h2 {
  margin-bottom: 16px;
}

.transactions-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

    .transaction {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px;
    border-radius: 8px;
    background-color: #f9f9f9;
    transition: background-color 0.2s ease;
    }

    .transaction:hover {
    background-color: #e8f0fe;
    }

    .transaction-name {
    font-weight: 600;
    margin: 0;
    text-transform: capitalize;
    }

    .transaction-date {
    font-size: 0.9rem;
    color: #666;
    margin: 4px 0 0;
    }

    .transaction-amount {
    font-weight: 700;
    font-size: 1.2rem;
    }
    .transaction-amount.income {
    color: #228927;
    }

    .transaction-amount.expense {
    color: #f46b64;
    }
  </style>
