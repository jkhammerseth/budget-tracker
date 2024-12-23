<script>
    import { goto } from '$app/navigation';
    
    let searchQuery = '';
    
    const handleSearch = () => {
      console.log('Search query:', searchQuery);
      // Handle the search logic here (e.g., update a store, send a request, etc.)
    };
    
    const navigateTo = (path) => {
      goto(path);
    };
  </script>
  
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <header class="header">
    <!-- Logo / Application Name -->
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div class="logo" on:click={() => navigateTo('/')}>
      <h1 class="app-name">Budget</h1>
    </div>
  
    <!-- Search Bar -->
    <div class="search-container">
      <input
        type="text"
        class="search-input"
        bind:value={searchQuery}
        placeholder="Search..."
      />
    </div>
  
    <!-- Navigation Links -->
    <nav class="nav">
      <button class="nav-link" on:click={() => navigateTo('/dashboard')}>Dashboard</button>
      <button class="nav-link" on:click={() => navigateTo('/expenses')}>Expenses</button>
      <button class="nav-link" on:click={() => navigateTo('/income')}>Income</button>
      <button class="nav-link" on:click={() => navigateTo('/calendar')}>Calendar</button>
    </nav>
  </header>
  
  <style>
    /* Ensure the header respects the sidenav width */
    .header {
      display: flex;
      position: fixed;
      align-items: center;
      padding: 15px 30px;
      height: var(--header-height);
      width: 100%;
      margin-right: 25rem; 
      background-color: var(--background-color);
      color: black;
      
    }
  
    .logo {
      display: flex;
      align-items: center;
      cursor: pointer;
    }
  
    .logo-image {
      width: 40px;
      height: 40px;
      margin-right: 10px;
    }
  
    .app-name {
      font-size: 1.5rem;
      font-weight: bold;
      margin: 0;
    }
  
    .search-container {
      display: flex;
      align-items: center;
      gap: 5px;
      padding-left: 100px;
    }
  
    .search-input {
      padding: 8px 16px;
      border-radius: 5px;
      border: none;
      outline: none;
      font-size: 1rem;
      width: 200px;
    }
  
    .search-input:focus {
      border-color: #68a1f1;
    }
  
    .nav {
      display: flex;
      gap: 15px;
      
    }
  
    .nav-link {
      background-color: transparent;
      color: black;
      font-size: 1rem;
      border: none;
      cursor: pointer;
      padding: 5px 10px;
      border-radius: 5px;
      transition: background-color 0.2s;
    }
  
    .nav-link:hover {
      background-color: #68a1f1;
      color: white;
    }
  </style>
  