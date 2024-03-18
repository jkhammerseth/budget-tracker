<script>
    import { onMount, onDestroy } from 'svelte';
  
    let currentTime = new Date().toLocaleTimeString();
  
    // Function to update the current time
    function updateTime() {
      currentTime = new Date().toLocaleTimeString();
    }
  
    onMount(() => {
      // Update time every second
      const interval = setInterval(updateTime, 1000);
  
      return () => {
        // Cleanup interval on component destruction
        clearInterval(interval);
      };
    });
  </script>
  
  <style>
    .clock {
      font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
      font-size: 2rem;
      text-align: center;
      padding: 20px;
      background: #3A87F2;
      color: #fff;
      border-radius: 8px;
      width: 200px;
      margin: auto;
    }
  </style>
  
  <div class="clock">{currentTime}</div>
  