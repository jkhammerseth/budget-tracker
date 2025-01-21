<script>
  import { formatAmount, formatExpenseAmount, fromISOString, daysFromNow } from "../utility/functions";

  import { recentTransactions } from '../stores/filteredTransactions'
  import { Info } from 'lucide-svelte';
  import { getIconComponent } from '../utility/icons'
</script>

<div class="recent-transactions">
  <div class="section-header">
    <h2>Recent Activity</h2>
    <span class="icon"><Info size="24" /></span>
  </div>
  <div class="transactions-list">
    {#if $recentTransactions.length > 0}
      {#each $recentTransactions as transaction}
        <div class="transaction">
          <div class="transaction-item">
            <div class="category-icon">
              {#if getIconComponent(transaction.category.name)}
                <svelte:component this={getIconComponent(transaction.category.name)} size="24" />
              {:else}
                📁 <!-- Fallback icon -->
              {/if}
            </div>
            <div class="transaction-details">
              <p class="transaction-name">{transaction.name}</p>
              <p class="transaction-date">{fromISOString(transaction.payment_date || transaction.date)}</p>
            </div>
          </div>
          <p class="transaction-amount {transaction.type === 'income' ? 'income' : 'expense'}">
            {transaction.type === 'income' 
              ? formatAmount(transaction.amount) 
              : formatExpenseAmount(transaction.amount)}
          </p>
        </div>
      {/each}
      <div class="view">
        <a class="view-link" href="/expenses" aria-label="View all transactions">View All Transactions</a>
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

.section-header {
    display: flex;
    justify-content: space-between;
}

  .section-header h2 {
  margin: 0;
  }

  .section-header .icon {
    display: flex;
    align-items: center;
  }


  .icon {
    padding: 16px;
  }

  .category-icon {
    padding: 10px;
    border-radius: 50%;
    background-color: #f1f1f1;
    
  }


  .recent-transactions {
  background: #fff;
  padding: 16px;
  border-radius: 8px;
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

    .transaction-item {
      display: flex;
      gap: 10px;
      align-items: center;
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
  </style>
