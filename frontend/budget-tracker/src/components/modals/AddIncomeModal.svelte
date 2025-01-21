<script>
    import { fetchIncome } from '../../routes/api/fetchIncome';
    import { postIncome, postListofIncomes } from '../../routes/api/addIncome';
    import Toast from '../Toast.svelte';
    import Modal from '../modals/Modal.svelte';
  
    let name = '';
    let amount = 0;
    let category = '';
    let frequency = '';
    let startDate = '';
    let endDate = '';

    let isRecurring = false;
  
    const categories = [
      'Salary',
      'Investment',
      'Rental',
      'Interest',
      'Dividends',
      'Business',
      'Other'
    ];
  
    let message = '';
    let theme = 'info'; // Possible values: 'success', 'error', 'info'
    let duration = 3000;
  
    async function addIncome() {
      const start = new Date(startDate); 
      let end;
      if (endDate) {
          end = new Date(endDate);
      }
  
      const incomes = [];
  
      if (frequency === 'One-time') {
          incomes.push(createIncomeObject(start));
      } else if (end) { 
          let currentDate = new Date(start.getTime());
  
          while (currentDate <= end) {
              incomes.push(createIncomeObject(currentDate));
              incrementDate(currentDate, frequency);
          }
      }
  
      try {
          let response;
          if (incomes.length === 1) {
              response = await postIncome(incomes[0]);
          } else {
              response = await postListofIncomes(incomes);
          }
  
          handleResponse(response);
      } catch (error) {
          console.error("Error adding income:", error);
          message = "Error adding income";
          theme = "error";
      }
  
      resetFormFields();
      fetchIncome();
  }
  
  function createIncomeObject(date) {
      return {
          Name: name,
          Amount: parseInt(amount, 10),
          Category: category,
          Frequency: frequency,
          Date: date.toISOString()
      };
  }
  
  function incrementDate(date, frequency) {
      switch (frequency) {
          case 'Daily':
              date.setDate(date.getDate() + 1);
              break;
          case 'Weekly':
              date.setDate(date.getDate() + 7);
              break;
          case 'Monthly':
              date.setMonth(date.getMonth() + 1);
              break;
          case 'Annually':
              date.setFullYear(date.getFullYear() + 1);
              break;
      }
  }
  
  function handleResponse(response) {
      if (!response.ok) {
          throw new Error("Failed to add income");
      }
      message = "Income added successfully!";
      theme = "success";
  }
  
  function resetFormFields() {
      name = '';
      amount = 0;
      category = '';
      frequency = '';
      startDate = '';
      endDate = '';
  }

  function closeModal() {
        activeModal.set(null); // Clear the active modal
    }
  
    $: frequency = isRecurring ? frequency : 'One-time';
  </script>

<Modal modalId="addIncome" onClose={closeModal} title="Add Income">
  <div class="modal-container">
    <div class="form-control">
      <input id="name" placeholder="Name" type="text" bind:value={name} />
      <input id="amount" placeholder="Amount" type="number" bind:value={amount} />
      <select id="category" bind:value={category}>
          <option value="">Select category</option>
          {#each categories as cat} 
            <option value={cat}>{cat}</option>
          {/each}
        </select>

        <div class="recurring-box">
          <label for="isRecurring">Recurring?</label>
          <input class="checkbox" type="checkbox" id="isRecurring" bind:checked={isRecurring} />
          </div>  

        {#if isRecurring}
                <select id="frequency" bind:value={frequency}>
                    <option value="">Select frequency</option>
                    <option value="Daily">Daily</option>
                    <option value="Weekly">Weekly</option>
                    <option value="Monthly">Monthly</option>
                    <option value="Annually">Annually</option>
                </select>
        {/if}

      <input id="date" type="date" bind:value={startDate} />

    {#if !(frequency === 'One-time' || frequency === '')}
        <input id="date" type="date" bind:value={endDate} />
    {/if} 

    <button on:click={addIncome}>Add Income</button>
  </div>
</div>
    <Toast {message} {theme} {duration} />
</Modal>

<style>
  .modal-container {
      display: grid;
      grid-template-columns: 1fr; 
      grid-template-rows: auto 1fr;
      gap: 20px;
      padding: 20px;
      height: 100%;
      margin: auto;
      justify-items: center;
  }
  
  .recurring-box {
      display: flex;
      align-items: center;
      gap: 10px; 
      font-size: 1rem; 
      color: #555; 
      margin-bottom: 15px; 
  }
  
  
  .recurring-box label {
      font-size: 1rem;
      cursor: pointer; 
      color: #333;
  }
  
  
  .recurring-box .checkbox {
      width: 18px; 
      height: 18px;
      border: 2px solid #ccc;
      border-radius: 4px; 
      background-color: #fff;
      cursor: pointer;
      transition: background-color 0.3s ease, border-color 0.3s ease;
  }
  
  .recurring-box .checkbox:checked {
      background-color: var(--primary-button-color); 
  }
  
  .recurring-box .checkbox:checked::after {
      content: '';
      position: absolute;
      top: 3px;
      left: 6px;
      width: 6px;
      height: 6px;
      background-color: #fff; 
      clip-path: polygon(0 50%, 40% 100%, 100% 0, 40% 25%);
      transform: rotate(45deg);
  }
  
  .form-control {
      display: flex;
      flex-direction: column;
      align-items: stretch;
      gap: 15px;
      width: 100%;
      max-width: 600px; 
      overflow-y: auto;
  }
      .form-control {
          display: flex;
          flex-direction: column;
          align-items: stretch; /* Ensures uniform width */
          gap: 15px;
      }
  
          input,
          select {
          max-width: 100%;
      }
  
    
      .form-control label {
        font-size: 1rem;
        color: #555;
      }
    
      input[type="text"],
      input[type="number"],
      input[type="date"],
      select {
        width: 100%;
        padding: 12px 16px;
        border: 1px solid #ccc;
        border-radius: 8px;
        box-sizing: border-box;
        font-size: 1rem;
        background-color: #f9f9f9;
        transition: border-color 0.3s ease, box-shadow 0.3s ease;
      }
    
      input[type="text"]:focus,
      input[type="number"]:focus,
      input[type="date"]:focus,
      select:focus {
        border-color: #5c6bc0;
        outline: none;
        box-shadow: 0 0 5px rgba(92, 107, 192, 0.5);
      }
    
      /* Button styles */
      button {
        width: 100%;
        padding: 12px;
        background-color: var(--primary-button-color);
        color: white;
        border: none;
        border-radius: 8px;
        cursor: pointer;
        font-size: 1rem;
        font-weight: bold;
        transition: background-color 0.3s ease, transform 0.2s ease, box-shadow 0.3s ease;
      }
    
      button:hover {
        background-color: var(--primary-button-hover-color);
        transform: translateY(-2px);
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      }
    
      button:active {
        transform: translateY(0);
        box-shadow: none;
      }
  
    
      /* Responsive adjustments */
      @media (max-width: 500px) {
        .modal-container {
          padding: 15px;
        }
    
        input[type="text"],
        input[type="number"],
        input[type="date"],
        select {
          padding: 10px 12px;
        }
    
        button {
          font-size: 1rem;
        }
      }
    </style>