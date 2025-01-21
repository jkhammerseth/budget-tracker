<script>
    import FaTimes from 'svelte-icons/fa/FaTimes.svelte'
    import Profile from '../Profile.svelte';
    import { writable } from 'svelte/store';
    import Modal from "./Modal.svelte";
    import { activeModal } from "../../stores/activeModal.js";
    import { Menu, X, Settings, User, UserRoundX } from 'lucide-svelte';

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

    function closeModal() {
        activeModal.set(null); // Clear the active modal
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

<Modal modalId="settings" onClose={closeModal} title="Settings">
    <div class="modal-container">
        <div class="settings-content">
            <nav class="settings-nav">
                <button
                    class="nav-button {activeMenu === 'general' ? 'active' : ''}"
                    on:click={() => setActiveMenu('general')}
                >General</button>
                <button
                    class="nav-button {activeMenu === 'profile' ? 'active' : ''}"
                    on:click={() => setActiveMenu('profile')}
                >Profile</button>
                <button
                    class="nav-button {activeMenu === 'privacy' ? 'active' : ''}"
                    on:click={() => setActiveMenu('privacy')}
                >Privacy</button>
                <button
                    class="nav-button {activeMenu === 'notifications' ? 'active' : ''}"
                    on:click={() => setActiveMenu('notifications')}
                >Notifications</button>
                <button
                    class="nav-button {activeMenu === 'help' ? 'active' : ''}"
                    on:click={() => setActiveMenu('help')}
                >Help</button>
            </nav>
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
                    <Profile />
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
</Modal>

<style>
.modal-container {
    background: #fff;
    border-radius: 10px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    padding: 1rem 1.5rem;
    max-width: 600px;
    margin: auto;
}

.top-of-box {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.icon-button {
    background: none;
    border: none;
    cursor: pointer;
}

.icon {
    width: 24px;
    height: 24px;
}

.settings-content {
    display: flex;
    gap: 1.5rem;
    margin-top: 1rem;
}

.settings-nav {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.nav-button {
    text-align: left;
    padding: 0.75rem 1rem;
    border: none;
    border-radius: 5px;
    background: #f9f9f9;
    cursor: pointer;
    font-size: 0.9rem;
    font-weight: 500;
    transition: all 0.2s ease-in-out;
}

.nav-button.active {
    background: #e0e0e0;
    font-weight: bold;
}

.nav-button:hover {
    background: #e8e8e8;
}

.menu-content {
    flex: 3;
    padding: 1rem;
    border: 1px solid #ddd;
    border-radius: 5px;
    background: #fafafa;
}

.setting-item {
    margin-bottom: 1rem;
}

.setting-item label {
    font-weight: bold;
}

.setting-item select,
.setting-item input[type="checkbox"] {
    margin-top: 0.5rem;
}
</style>
