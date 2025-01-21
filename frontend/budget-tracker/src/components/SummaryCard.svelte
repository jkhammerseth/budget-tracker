<script>
    export let title;
    export let value;
  </script>
  
  <div class="summary-card">
      <div>
        <p class="title">{title}</p>
        <h3 class="value">{value}</h3>
      </div>
  </div>
  
  <style>
    .summary-card {
      padding: 20px;
      border-radius: 8px;
      color: #333;
      background: #fff;
    }

    .title {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .value {
      font-size: 24px;
      font-weight: bold;
      margin-top: 10px;
    }

  </style>
  