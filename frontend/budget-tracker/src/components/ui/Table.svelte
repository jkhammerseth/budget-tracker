<script>
    export let data = [];
    export let columns = [];
    export let onRowClick = () => {};
  </script>
  
  <div class="table-wrapper">
    <table>
      <thead>
        <tr>
          {#each columns as column}
            <th class="{column.class}">{column.label}</th>
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each data as row}
          <tr on:click={() => onRowClick(row)}>
            {#each columns as column}
              <td class="{column.class}">{column.cell(row)}</td>
            {/each}
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  
  <style>
    .container {
      width: 100%;
    }
  
    .table-wrapper {
      overflow-x: auto;
    }
    .name,
    .date {
      font-weight: 450;
    }
  
    table {
      width: 100%;
      border-collapse: collapse;
      background-color: #ffffff;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
      border-radius: 8px;
    }
  
    th, td {
      padding: 12px 16px;
      text-align: left;
      border-bottom: 1px solid #eee;
    }
  
    th {
      background-color: #F9FAFB; 
      height: 2.5rem;
      border-bottom: 1px solid #E5E7EB;
    }
  
    td.text-right-number,
    th.text-right, td.text-right {
      text-align: right;
    }
  
    tr:nth-child(odd) {
      background-color: #f9f9f9;
    }
  
    tr:hover {
      background-color: #f1f1f1;
    }
    .date-info {
      display: flex;
      flex-direction: column;
      gap: 4px;
      text-align: left;
      align-items: right;
    }
  
    .text-right {
      text-align: right;
      align-items: right;
    }
  
    .date-in-table {
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: flex-end;
    }
  
    .category-icon {
      font-size: 1.5rem;
      margin-right: 12px;
    }
  
    .days-until {
      font-size: 0.85rem;
      color: #777;
    }
  
    .date-logo {
      margin-right: 5px;
    }
  
    /* Media Query for smaller screens */
    @media (max-width: 768px) {
      table {
        display: block;
        overflow-x: auto;
      }
  
      th, td {
        padding: 8px 12px;
        font-size: 12px; /* Reduce font size on smaller screens */
      }
  
      .category-icon {
        font-size: 1.5rem; /* Reduce icon size on smaller screens */
      }
  
      .days-until {
        font-size: 0.75rem; /* Smaller text for 'days left' on mobile */
      }
    }
  </style>
  