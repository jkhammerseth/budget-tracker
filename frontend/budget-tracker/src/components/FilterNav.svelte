<style>
    .sidenav {
      height: 100%;
      width: 300px; 
      position: fixed;
      z-index: 1;
      top: 0;
      right: 0;
      background: #9FB6D0;
      overflow-x: hidden;
      padding-top: 20px;
    }
  
    .sidenav h1 {
      color: black;
      text-align: center;
      margin-bottom: 30px;
      font-family: 'Roboto', sans-serif;
      font-size: 34px; 
      font-weight: 300; 
    }
  
    .sidenav a {
      padding: 10px 20px; 
      text-decoration: none;
      font-size: 20px;
      color: black;
      display: flex;
      align-items: center;
      gap: 10px; 
      transition: 0.3s;
      border-left: 3px solid transparent;
    }

    .user-button-container {
      position: absolute;
      bottom: 20px;
      width: 100%;
      padding-bottom: 10px;
      font-family: 'Roboto', sans-serif;
    }

    .icon {
      width: 40px;
      height: 40px;
      margin-top: 3px;
    }

    .full-name {
      display: flex;
      font-size: 1.1rem;
      margin: 15px;
      margin-left: 15px;
    }

    .user-button {
      display: flex;
      flex-direction: row;
      width: 90%;
      margin: 10px;
      margin-bottom: 0;
 
      background-color: transparent; 
      color: black;
      border: none;
      border-radius: 10px;
      cursor: pointer;
    }

    .user-button.active {
      background-color: #2980b9; 
    }
  
    .sidenav a:hover {
      background-color: #6d9dd2; 
    }
  
    .sidenav .active {
      background-color: #3A87F2; 
    }
  </style>

<script>
    import { page } from '$app/stores';
    import { derived } from 'svelte/store';
    import UserMenu from './UserMenu.svelte';
    import { onMount } from 'svelte';
    import FaRegUserCircle from 'svelte-icons/fa/FaRegUserCircle.svelte'
    import { goto } from '$app/navigation';
    import FaUserAltSlash from 'svelte-icons/fa/FaUserAltSlash.svelte'
    import { checkAuthStatus } from '../routes/api/auth.js';
    import { user } from '../stores/user.js';
  import Calendar from './Calendar.svelte';



  onMount(() => {
    checkAuthStatus();
  });

  function goToLogin() {
    goto('/login');
  }
  </script>
  
  <div class="sidenav">
    <Calendar />
  </div>
  