<script>
    import Modal from "./Modal.svelte";
    import { activeModal } from "../../stores/activeModal.js";
    import { getIconComponent } from '../../utility/icons'

    export let category = null;

    function closeModal() {
        activeModal.set(null);
    }
</script>

<Modal modalId="category" onClose={closeModal} title={category.name}>
    <h3>Subcategories</h3>
    <ul class="subcategories">
        {#if category.subcategories && category.subcategories.length > 0}
            {#each category.subcategories as subcategory}
                <li class="subcategory-item">
                    {subcategory.name}
                </li>
            {/each}
        {:else}
            <li class="subcategory-item empty">No subcategories available</li>
        {/if}
    </ul>
</Modal>

<style>
    h3 {
        font-size: 1.1rem;
        font-weight: 500;
        color: #64748b;
        margin: 0 0 1rem 0;
        padding-left: 0.5rem;
    }

    .subcategories {
        list-style: none;
        padding: 0;
        margin: 0;
        border-radius: 0.75rem;
        overflow: hidden;
        border: 1px solid #e2e8f0;
    }

    .subcategory-item {
        padding: 1rem 1.25rem;
        border-bottom: 1px solid #e2e8f0;
        color: #334155;
        background-color: white;
        transition: all 0.2s ease;
        display: flex;
        align-items: center;
        gap: 0.75rem;
    }

    .subcategory-item:last-child {
        border-bottom: none;
    }

    .subcategory-item:hover {
        background-color: #f8fafc;
        padding-left: 1.5rem;
    }

    .subcategory-item::before {
        content: "•";
        color: #64748b;
        font-size: 1.2rem;
    }

    /* Empty state styling */
    .subcategory-item.empty {
        color: #94a3b8;
        font-style: italic;
        padding: 1.5rem;
        text-align: center;
        justify-content: center;
    }

    .subcategory-item.empty::before {
        display: none;
    }

    /* Animation for list items */
    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(10px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .subcategory-item {
        animation: slideIn 0.2s ease-out forwards;
    }

    .subcategory-item:nth-child(1) { animation-delay: 0.1s; }
    .subcategory-item:nth-child(2) { animation-delay: 0.15s; }
    .subcategory-item:nth-child(3) { animation-delay: 0.2s; }
    .subcategory-item:nth-child(4) { animation-delay: 0.25s; }
    .subcategory-item:nth-child(5) { animation-delay: 0.3s; }
</style>