<script>
    import { onMount } from 'svelte';
    import { filteredIncome } from './stores/filteredIncome.js'; // Ensure correct import path
    import { derived } from 'svelte/store';
    import { fetchIncome } from '../routes/api/fetchIncome';
    import { FetchExpenses } from '../routes/api/fetchExpenses';
  
    onMount(() =>{
        fetchIncome();
    });
  
    const incomeByCategory = derived(filteredIncome, $filteredIncome => {
        return $filteredIncome.reduce((acc, income) => {
            if (acc[income.Category]) {
                acc[income.Category] += income.Amount;
            } else {
                acc[income.Category] = income.Amount;
            }
            return acc;
        }, {});
    });
  
    // Function to format the amount in a more readable way
    function formatAmount(amount) {
        return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'NOK' }).format(amount); // Adjust currency code as needed
    }
  </script>
  <div class="asset-box">
      <h1>Income Sources</h1>
      <hr />
  <div class="container">
        <div class="income-by-category-container">
            <ul>
                {#each Object.entries($incomeByCategory) as [category, totalAmount]}
                <div class="income-source-item">
                    <li><strong>{category}:</strong> {formatAmount(totalAmount)}</li>
                </div>
                {/each}
            </ul>
        </div>
     
  </div>
  </div>
  
  <style>
    .asset-box {
        border-radius: 20px;
        padding: 20px;
        margin: 20px auto;
        max-width: 600px;
        background-color: white;
        border: black 1px solid;
        box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        font-family: 'Roboto', sans-serif;
    }
  
    .container {
        display: flex;
        flex-direction: column;
        gap: 20px;
    }
  
    .income-source-item {
        list-style-type: none;
        margin: 10px 0;
        padding: 10px;
        background-color: #FFFFFF;
        border-radius: 10px;
        box-shadow: 0 2px 4px rgba(0,0,0,0.05);
        transition: box-shadow 0.3s ease;
    }
  
    .income-source-item:hover {
        box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    }
  
    h1 {
        margin-bottom: 20px;
        color: #333;
        font-size: 24px;
        text-align: center;
    }
  
    hr {
        border-color: #DDD;
        margin-bottom: 20px;
    }
  
    li {
        color: #333;
        font-size: 18px;
        line-height: 1.5;
    }

  </style>
  