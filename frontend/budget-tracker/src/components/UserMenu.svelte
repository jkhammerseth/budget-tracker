<script>
    import { logout } from "../routes/api/auth";
    import FaCog from 'svelte-icons/fa/FaCog.svelte'
    import FaSignOutAlt from 'svelte-icons/fa/FaSignOutAlt.svelte'
    import { fade } from 'svelte/transition';
    import SettingsModal from "./modals/SettingsModal.svelte";
    import { onMount, onDestroy } from 'svelte';

    let showModal = false;

    export let show = false;
    let menuRef;

    function handleGlobalClick(event) {
        if (!menuRef.contains(event.target)) {
            close();
        }
    }

    function handleSettingsClick() {
        showModal = true;
        close();
    }

    function handleLogoutClick() {
        logout();
        close();
    }

    function close() {
        show = false;
    }
  
    function handleKeydown(event) {
        if (event.key === 'Escape') {
            close();
        }
    }
    
  </script>
  {#if show}
  <div transition:fade={{ duration: 100 }}>
  <div class="user-menu" class:hidden={!show}>
    <ul>
      <li>
        <span class="icon">
            <FaCog/>
        </span>
        <button on:click={handleSettingsClick}>
            Settings
        </button>
    </li>
    <hr>
      <li>
          <span class="icon">
              <FaSignOutAlt/>
          </span>
        <button class="logout" on:click={handleLogoutClick}>
            Log out
    </button>
    </li>
    </ul>
  </div>
</div>
{/if}

<SettingsModal bind:showModal/>
  
  <style>
    .user-menu {
      position: fixed;
      bottom: 62px;
      left: 0;
      margin-left: 12px;
      background-color: white;
      border: 1px solid #ccc;
      width: 247px;
      border-radius: 10px;
      padding: 10px;
      display: none;
    }
    
    .user-menu.hidden {
      display: none;
    }
    
    .user-menu:not(.hidden) {
      display: block;
    }

    .logout {
      display: flex;
      align-items: left;
      height: auto;
      background: none;
      border: none;
      font-size: 16px;
      cursor: pointer;
      color: #333;
      margin: 10px;
      margin-left: 0px;

      margin-bottom: 10px;
    }

    .icon {
      width: 20px;
      height: 20px;
      margin: 5px;
      margin-top: 10px;
    }
    
    ul {
      list-style: none;
      padding: 0;
      margin: 0;
    }
    
    li {
      display: flex;
      align-items: left;
      height: auto;

      margin-bottom: 10px;

    }

    button {
      background: none;
      border: none;
      font-size: 16px;
      cursor: pointer;
      color: #333;
      margin: 10px;
      margin-left: 0px;
    }
    
    li:last-child {
      margin-bottom: 0;
    }

    li:hover {
      background-color: #f5f5f5;
      border-radius: 4px;
    }

  </style>
  