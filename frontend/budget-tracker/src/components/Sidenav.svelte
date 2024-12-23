<script>
  import { onMount } from 'svelte';
  import FaRegUserCircle from 'svelte-icons/fa/FaRegUserCircle.svelte';
  import { goto } from '$app/navigation';
  import FaUserAltSlash from 'svelte-icons/fa/FaUserAltSlash.svelte';
  import { checkAuthStatus } from '../routes/api/auth.js';
  import { user } from '../stores/user.js';
  import Calendar from './Calendar.svelte';
  import FaCog from 'svelte-icons/fa/FaCog.svelte';
  import SettingsModal from "./modals/SettingsModal.svelte";
  import ExpensesByCategoryBox from './ExpensesByCategoryBox.svelte';
  import FaBars from 'svelte-icons/fa/FaBars.svelte';
  import FaTimes from 'svelte-icons/fa/FaTimes.svelte';

  let showModal = false;
  let isCollapsed = false; // To toggle the sidenav state

  onMount(() => {
    checkAuthStatus();
  });

  function handleSettingsClick() {
    showModal = true;
  }

  function toggleSidenav() {
    isCollapsed = !isCollapsed;
  }
</script>

<div class="sidenav {isCollapsed ? 'collapsed' : ''}">
  <!-- Collapse/Expand Button -->
  <button class="collapse-button" on:click={toggleSidenav}>
    {#if isCollapsed}
      <FaBars />
    {:else}
      <FaTimes />
    {/if}
  </button>

  <!-- User Menu Section -->
  <div class="user-menu">
    {#if $user.loggedIn}
      <button class="settings-button {isCollapsed ? 'collapsed' : ''}" on:click={() => handleSettingsClick()}>
        <FaCog />
      </button>
      <button class="user-image {isCollapsed ? 'collapsed' : ''}">
        <FaRegUserCircle />
      </button>
    {:else}
      <button class="user-button" on:click|preventDefault={() => goto('/login')}>
        <FaUserAltSlash />
        <span class="full-name" class:hide={isCollapsed}>Sign In</span>
      </button>
    {/if}
    <SettingsModal bind:showModal />
  </div>

  <!-- Categories Section -->
  <div class="categories-section {isCollapsed ? 'collapsed' : ''}">
    <ExpensesByCategoryBox />
  </div>

  <!-- Calendar Section -->
  <div class="calendar-container {isCollapsed ? 'collapsed' : ''}">
    <Calendar />
  </div>
</div>

<style>
  :root {
    --primary-color: #007bff;
    --text-color: #fff;
    --icon-size: 24px;
    --border-radius: 8px;
    --font-family: 'Roboto', sans-serif;
  }

  .sidenav {
    height: 100%;
    width: 22rem;
    position: fixed;
    z-index: 1;
    top: 0;
    right: 0;
    background: var(--sidenav-color);
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 20px;
    font-family: var(--font-family);
    transition: width 0.3s ease, background 0.3s ease;
    overflow-x: hidden;
    border-radius: 0 0 8px 8px;
  }

  .sidenav.collapsed {
    width: 1%;
    background-color: var(--background-color);
  }

  .user-image.collapsed,
  .settings-button.collapsed,
  .categories-section.collapsed,
  .calendar-container.collapsed {
    display: none;
  }

  .collapse-button {
    width: var(--icon-size);
    height: var(--icon-size);
    position: absolute;
    top: 10px;
    left: 10px;
    background: none;
    border: none;
    cursor: pointer;
    color: var(--text-color);
    font-size: 1.5rem;
    transition: color 0.3s ease;
  }

  .collapse-button:hover {
    color: var(--primary-color);
  }

  .user-menu {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    gap: 10px;
  }

  .user-button,
  .user-image {
    width: var(--icon-size);
    height: var(--icon-size);
    border-radius: 50%;
    background: #f0f0f0;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    overflow: hidden;
    transition: background 0.3s ease;
  }

  .user-button {
    display: flex;
    align-items: center;
    gap: 10px;
    background: none;
    border: none;
    color: var(--text-color);
    cursor: pointer;
    font-size: 1rem;
    font-weight: bold;
    transition: opacity 0.3s ease;
  }

  .user-button .hide {
    display: none;
  }

  .settings-button {
    width: var(--icon-size);
    height: var(--icon-size);
    background: none;
    border: none;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: transform 0.3s ease;
  }

  .settings-button:hover {
    transform: scale(1.1);
  }

  .categories-section {
    flex-grow: 1;
    overflow-y: auto;
    margin-bottom: 10px;
    transition: opacity 0.3s ease;
  }

  .categories-section.collapsed {
    opacity: 0.5;
  }

  .calendar-container {
    margin-bottom: 50px;
  }

  .calendar-container :global(.calendar) {
    width: 100%;
  }

  /* Scrollbar styling */
  .categories-section::-webkit-scrollbar {
    width: 5px;
  }

  .categories-section::-webkit-scrollbar-thumb {
    background: #ccc;
    border-radius: var(--border-radius);
  }

  .categories-section::-webkit-scrollbar-thumb:hover {
    background: #aaa;
  }
</style>
