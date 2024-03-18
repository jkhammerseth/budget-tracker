<script>
    export let text = 'Click Me'; 
    export let type = 'button'; 
    export let onClick = () => {};
  
    // Optional props for more customization
    export let buttonClass = '';
  </script>
  
  <button type={type} class={`btn ${buttonClass}`} on:click={onClick}>
    {text}
  </button>
  
  <style>
    .btn {
      padding: 10px 20px;
      border: none;
      border-radius: 5px;
      background-color: #007bff;
      color: white;
      font-size: 16px;
      cursor: pointer;
      outline: none;
      transition: background-color 0.3s, transform 0.2s;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
  
    .btn:hover {
      background-color: #0056b3;
      transform: translateY(-2px);
    }
  
    .btn:active {
      transform: translateY(1px);
    }
  
    .btn:focus {
      box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.5);
    }
  </style>
  