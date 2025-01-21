<script>
    import Modal from "./Modal.svelte";
    import { activeModal } from "../../stores/activeModal.js";

    let selectedFile = null;
    let isLoading = false;

    function handleFileChange(event) {
        selectedFile = event.target.files[0];
    }

    async function uploadFile() {
        if (!selectedFile) {
            alert("Please select a file to upload!");
            return;
        }

        // Validate file type
        if (!selectedFile.name.toLowerCase().endsWith('.csv')) {
            alert('Please select a CSV file');
            return;
        }

        isLoading = true;
        const uploadButton = document.querySelector('.upload-button');
        uploadButton.classList.add('loading');

        try {
            const formData = new FormData();
            formData.append("file", selectedFile);

            const response = await fetch('http://localhost:8080/api/users/expenses/import', {
                method: 'POST',
                credentials: 'include',
                body: formData
            });

            const result = await response.json();

            if (!response.ok) {
                throw new Error(result.error || 'Failed to upload file');
            }

            alert(`Upload successful! Imported ${result.expense_count} expenses and ${result.income_count} income entries.`);
            closeModal();
        } catch (error) {
            console.error("Upload failed:", error);
            alert(`Failed to upload the file: ${error.message}`);
        } finally {
            isLoading = false;
            uploadButton.classList.remove('loading');
        }
    }

    function closeModal() {
        activeModal.set(null);
    }
</script>

<Modal modalId="uploadFile" onClose={closeModal} title="Upload a File">
    <div>
        <label for="fileInput" class="file-input-label">
            Choose a CSV file:
        </label>
        <input
            id="fileInput"
            type="file"
            accept=".csv"
            on:change={handleFileChange}
            class="file-input"
            disabled={isLoading}
        />
        {#if selectedFile}
            <p>Selected File: {selectedFile.name}</p>
        {/if}

        <div class="modal-buttons">
            <button 
                class="upload-button" 
                on:click={uploadFile} 
                disabled={isLoading || !selectedFile}
            >
                {isLoading ? 'Uploading...' : 'Upload'}
            </button>
            <button 
                class="cancel-button" 
                on:click={closeModal}
                disabled={isLoading}
            >
                Cancel
            </button>
        </div>
    </div>
</Modal>

    <style>
    .file-input-label {
        display: block;
        margin-bottom: 1rem;
        font-weight: 500;
        color: #374151;
        font-size: 0.95rem;
    }

    .file-input {
        width: 95%;
        padding: 0.75rem;
        margin-bottom: 1.5rem;
        background-color: #f9fafb;
        border: 2px dashed #e5e7eb;
        border-radius: 0.5rem;
        transition: all 0.2s ease-in-out;
        cursor: pointer;
    }

    .file-input:hover {
        background-color: #f3f4f6;
        border-color: #d1d5db;
    }

    .file-input:focus {
        outline: none;
        border-color: #60a5fa;
        box-shadow: 0 0 0 3px rgba(96, 165, 250, 0.1);
    }

    p {
        background-color: #f3f4f6;
        padding: 0.75rem 1rem;
        border-radius: 0.5rem;
        margin-bottom: 1.5rem;
        font-size: 0.9rem;
        color: #4b5563;
        word-break: break-all;
    }

    .modal-buttons {
        display: flex;
        gap: 1rem;
        justify-content: flex-end;
        margin-top: 2rem;
        padding-bottom: 1rem;
    }

    .upload-button,
    .cancel-button {
        padding: 0.625rem 1.25rem;
        border-radius: 0.5rem;
        font-weight: 500;
        font-size: 0.95rem;
        transition: all 0.2s ease;
        border: none;
        cursor: pointer;
    }

    .upload-button {
        background-color: #3b82f6;
        color: white;
    }

    .upload-button:hover {
        background-color: #2563eb;
        transform: translateY(-1px);
    }

    .upload-button:active {
        transform: translateY(0);
    }

    .cancel-button {
        background-color: #f3f4f6;
        color: #4b5563;
    }

    .cancel-button:hover {
        background-color: #e5e7eb;
        color: #374151;
    }

    :global(.modal-content) {
        animation: modalSlide 0.3s ease-out;
    }

    @keyframes modalSlide {
        from {
            opacity: 0;
            transform: translateY(-10px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .file-input::-webkit-file-upload-button {
        padding: 0.5rem 1rem;
        background-color: #60a5fa;
        color: white;
        border: none;
        border-radius: 0.375rem;
        margin-right: 1rem;
        cursor: pointer;
        transition: background-color 0.2s ease;
    }

    .file-input::-webkit-file-upload-button:hover {
        background-color: #3b82f6;
    }

    .upload-button.loading {
        position: relative;
        cursor: not-allowed;
        opacity: 0.7;
    }

    .upload-button.loading::after {
        content: "";
        position: absolute;
        width: 1rem;
        height: 1rem;
        border: 2px solid transparent;
        border-radius: 50%;
        border-top-color: white;
        animation: spin 0.8s linear infinite;
        margin-left: 0.5rem;
    }

    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }
</style>
