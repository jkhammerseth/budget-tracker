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
      font-family: var(--font-family);
      border: 1px solid var(--component-border-color);
      font-size: 2rem;
      text-align: center;
      padding: 20px;
      background: var(--primary-color);
      color: #fff;
      border-radius: 8px;
      width: 200px;
      margin: auto;
      height: 2rem;
    }
  </style>
  
  <div class="clock">{currentTime}</div>
  