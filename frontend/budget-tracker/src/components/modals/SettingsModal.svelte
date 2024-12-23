<script>
    import FaTimes from 'svelte-icons/fa/FaTimes.svelte'
    import Profile from '../Profile.svelte';
    import { writable } from 'svelte/store';

    export let show = false;
    export let showModal = false;

    if (showModal) {
        show = false;
    }

    // Store for settings
    const settings = writable({
        general: {
            theme: 'light',
            language: 'English',
            timezone: 'UTC',
            autoSave: true
        },
        privacy: {
            profileVisibility: 'public',
            showEmail: false,
            showActivity: true,
            twoFactorAuth: false
        },
        notifications: {
            email: true,
            push: true,
            sound: true,
            desktop: false,
            marketing: false
        }
    });

    let modalContent;
    let activeMenu = 'general';

    function setActiveMenu(menu) {
        activeMenu = menu;
    }

    function close(event) {
        // Only close if clicking directly on the backdrop
        if (event.target.classList.contains('modal-backdrop')) {
            showModal = false;
        }
    }

    function handleKeyup(event) {
        if (event.key === 'Escape') {
            showModal = false;
        }
    }
</script>

{#if showModal}
    <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
    <div 
        role="dialog" 
        class="modal-backdrop" 
        on:click={close} 
        on:keyup={handleKeyup} 
        aria-label="Close modal"
    >
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div 
            bind:this={modalContent} 
            class="modal-content"
            on:click|stopPropagation={() => {}} 
            tabindex="0"
        >
            <div class="top-of-box">
                <h2>Settings</h2>
                <button class="icon-button" on:click={() => showModal = false}>
                    <span class="icon"><FaTimes/></span>
                </button>
            </div>
            <hr />
            <div class="settings-content">
                <div class="settings-nav">
                    <button class="nav-button" on:click={() => setActiveMenu('general')}>General</button>
                    <button class="nav-button" on:click={() => setActiveMenu('profile')}>Profile</button>
                    <button class="nav-button" on:click={() => setActiveMenu('privacy')}>Privacy</button>
                    <button class="nav-button" on:click={() => setActiveMenu('notifications')}>Notifications</button>
                    <button class="nav-button" on:click={() => setActiveMenu('help')}>Help</button>
                </div>
                <div class="menu-content">
                    {#if activeMenu === 'general'}
                        <div class="settings-section">
                            <h3>General Settings</h3>
                            <hr />
                            <div class="setting-item">
                                <label for="theme">Theme</label>
                                <select id="theme" bind:value={$settings.general.theme}>
                                    <option value="light">Light</option>
                                    <option value="dark">Dark</option>
                                    <option value="system">System Default</option>
                                </select>
                            </div>
                            <div class="setting-item">
                                <label for="language">Language</label>
                                <select id="language" bind:value={$settings.general.language}>
                                    <option value="English">English</option>
                                    <option value="Spanish">Spanish</option>
                                    <option value="French">French</option>
                                    <option value="German">German</option>
                                </select>
                            </div>
                            <div class="setting-item">
                                <label for="timezone">Timezone</label>
                                <select id="timezone" bind:value={$settings.general.timezone}>
                                    <option value="UTC">UTC</option>
                                    <option value="EST">EST</option>
                                    <option value="PST">PST</option>
                                    <option value="GMT">GMT</option>
                                </select>
                            </div>
                            <div class="setting-item checkbox">
                                <label>
                                    <input 
                                        type="checkbox" 
                                        bind:checked={$settings.general.autoSave}
                                    />
                                    Enable Auto-Save
                                </label>
                            </div>
                        </div>
                    {/if}
                    {#if activeMenu === 'profile'}
                        <Profile/>
                    {/if}
                    {#if activeMenu === 'privacy'}
                        <p>Privacy settings here</p>
                    {/if}
                    {#if activeMenu === 'notifications'}
                        <p>Notification settings here</p>
                    {/if}
                    {#if activeMenu === 'help'}
                        <p>Help settings here</p>
                    {/if}
                </div>
            </div>
        </div>
    </div>
{/if}

<style>

.settings-section {
        padding: 1rem;
    }

    .setting-item {
        margin-bottom: 1.5rem;
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    .setting-item.checkbox {
        flex-direction: row;
        align-items: center;
        gap: 1rem;
    }

    .setting-item label {
        font-weight: 500;
    }

    .setting-item select {
        padding: 0.5rem;
        border-radius: 4px;
        border: 1px solid #ccc;
        background-color: white;
    }

    .icon-button {
        background: none;
        border: none;
        padding: 0;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .top-of-box {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    

    .icon {
        width: 24px; 
        height: 24px; 
        color: #333; 
    }

    .settings-content {
        display: grid;
        gap: 20px;
        grid-template-columns: repeat(4, 1fr);
        margin-bottom: 20px;
    }

    .settings-nav {
        grid-column: 1;
        grid-row: 1;
        margin-top: 0;

        margin-bottom: 20px;
    }

    .nav-button {
        text-align: left;
        padding-left: 30px;
        width: 100%;
        padding: 10px 0;
        border: 0;
        border-radius: 5px;
        box-sizing: border-box;
        background-color: #E8EFF1;
    }

    .nav-button:hover {
        background-color: #e0e0e0;
        border: 1px solid #ccc;
    }

    .nav-button:focus {
        outline: none;
    }

    .menu-content {
        grid-column: 2 / 5;
        grid-row: 1;
        margin-top: 0;
        margin-bottom: 20px;
    }


    .modal-backdrop {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.4);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
    }

    .modal-content {
        background-color: #E8EFF1;
        border: 1px solid black;
        border-radius: 10px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        height: 600px;
        width: 600px;
        max-width: 800px;
        padding: 0 20px;
        width: calc(100% - 40px);
        box-sizing: border-box;
        display: flex;
        flex-direction: column;

    }


    @media (max-width: 600px) {
        .modal-content {
            width: calc(100% - 20px);
            padding: 10px;
        }
    }
</style>