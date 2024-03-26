<script>
    import FaTimes from 'svelte-icons/fa/FaTimes.svelte'
    import Profile from '../Profile.svelte';

    export let show = false;

    export let showModal = false;

    if (showModal) {
        show = false;
    }

    let modalContent;

    let activeMenu = 'general';

    function setActiveMenu(menu) {
        activeMenu = menu;
    }

    function close() {
        showModal = false;
    }

    function handleKeyup(event) {
        if (event.key === 'Escape') {
            close();
        }
    }

</script>

{#if showModal}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
    <div role="dialog" class="modal-backdrop" on:click={close} on:keyup={close} aria-label="Close modal">
            <!-- svelte-ignore a11y-no-static-element-interactions -->
    <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
    <div bind:this={modalContent} class="modal-content" on:click|stopPropagation tabindex="0">
        <div class="top-of-box">
            <h2>Settings</h2>
            <span class="icon" on:click={close}><FaTimes/></span>
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
                    <p>General settings here</p>
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