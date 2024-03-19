<script>
    import { writable } from 'svelte/store';

    export let message = "";
    export let theme = "info"; // "success", "error", or "info"
    export let duration = 3000; // in milliseconds

    let show = writable(false);

    function showToast() {
        show.set(true);
        setTimeout(() => {
            show.set(false);
        }, duration);
    }

    // Reactive statement to react to changes in `message`
    $: if (message) {
        showToast();
    }
</script>

{#if $show}
    <div class={`toast ${theme}`}>
        {message}
    </div>
{/if}

<style>
    .toast {
        position: fixed;
        bottom: 20px;
        left: 50%;
        transform: translateX(-50%);
        padding: 10px 20px;
        border-radius: 5px;
        color: white;
        animation: fadeInOut 2.5s forwards;
    }

    .success {
        background-color: #4CAF50; 
    }

    .error {
        background-color: #F44336;
    }

    .info {
        background-color: #2196F3;
    }

    @keyframes fadeInOut {
        0%, 100% { opacity: 0; }
        10%, 90% { opacity: 1; }
    }
</style>
