<script>
    import { activeModal } from "../../stores/activeModal";
    import { X } from 'lucide-svelte';

    export let modalId; // Unique ID for the modal
    export let title = ''; // Title for the modal
    export let onClose = () => {};

    function close() {
        activeModal.set(null);
        onClose();
    }
</script>

{#if $activeModal === modalId}
    <div class="modal-overlay" on:click={close}>
        <div class="modal-content" on:click|stopPropagation>
            <div class="modal-header">
                <span class="modal-title">{title}</span>
                <button class="cross" on:click={close}>
                    <X size="40" />
                </button>
            </div>
            <slot />
            <button class="close-button" on:click={close}>Close</button>
        </div>
    </div>
{/if}

<style>
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .modal-content {
        background: white;
        padding: 20px;
        border-radius: 8px;
        max-width: 500px;
        width: 90%;
    }

    .modal-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 1rem;
    }

    .modal-title {
        font-size: 1.5rem;
        font-weight: 600;
        padding-left: 1.5rem;
        color: var(--text-color);
    }

    .cross {
        background: none;
        border: none;
        cursor: pointer;
        color: var(--text-color);
        font-size: 1.5rem;
        transition: color 0.3s ease;
    }

    .close-button {
        width: 100%;
        padding: 14px;
        background-color: var(--primary-button-color);
        color: white;
        border: none;
        border-radius: 8px;
        cursor: pointer;
        font-size: 1.1rem;
        font-weight: 600;
        transition: background-color 0.3s ease, transform 0.2s ease, box-shadow 0.3s ease;
    }

    .close-button:hover {
        background-color: var(--primary-button-hover-color);
        transform: translateY(-2px);
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }
</style>
