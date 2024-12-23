<script>
  import { user } from '../stores/user';
  import FaPencilAlt from 'svelte-icons/fa/FaPencilAlt.svelte';
  import { updateUser } from '../routes/api/updateUser';
  import { writable } from 'svelte/store';

  let editState = writable({
    isEditingName: false,
    isEditingEmail: false,
    firstName: $user.firstName,
    lastName: $user.lastName,
    email: $user.email
  });

  let errorMessage = '';
  let successMessage = '';

  $: formData = {
    firstName: $editState.firstName,
    lastName: $editState.lastName,
    email: $editState.email
  };

  const showMessage = (message, isError = false) => {
    if (isError) {
      errorMessage = message;
      successMessage = '';
    } else {
      successMessage = message;
      errorMessage = '';
    }
    setTimeout(() => {
      errorMessage = '';
      successMessage = '';
    }, 3000);
  };

  const handleUpdateName = async () => {
    try {
      await updateUser({
        ...formData,
        email: $user.email
      });
      
      $user.firstName = $editState.firstName;
      $user.lastName = $editState.lastName;
      
      // Reset edit state
      $editState.isEditingName = false;
      showMessage('Name updated successfully');
    } catch (error) {
      showMessage(error.message, true);
    }
  };

  const handleUpdateEmail = async () => {
    try {
      await updateUser({
        ...formData,
        firstName: $user.firstName,
        lastName: $user.lastName
      });
      
      $user.email = $editState.email;
      $editState.isEditingEmail = false;
      showMessage('Email updated successfully');
    } catch (error) {
      showMessage(error.message, true);
    }
  };

  const startEditingName = () => {
    $editState.firstName = $user.firstName;
    $editState.lastName = $user.lastName;
    $editState.isEditingName = true;
  };

  const startEditingEmail = () => {
    $editState.email = $user.email;
    $editState.isEditingEmail = true;
  };

  const cancelEdit = (field) => {
    if (field === 'name') {
      $editState.firstName = $user.firstName;
      $editState.lastName = $user.lastName;
      $editState.isEditingName = false;
    } else if (field === 'email') {
      $editState.email = $user.email;
      $editState.isEditingEmail = false;
    }
  };
</script>

<div class="profile-container">
  <h3>Profile</h3>
  <hr />
  
  {#if errorMessage}
    <div class="error-message">{errorMessage}</div>
  {/if}
  
  {#if successMessage}
    <div class="success-message">{successMessage}</div>
  {/if}

  <div class="content-box">
    <div class="name-box">
      <label class="name label" for="firstName">Name</label>
      <div class="name">
        {#if $editState.isEditingName}
          <div class="input-group">
            <input 
              id="firstName" 
              type="text" 
              bind:value={$editState.firstName} 
              placeholder="First Name"
            />
            <input 
              id="lastName" 
              type="text" 
              bind:value={$editState.lastName} 
              placeholder="Last Name"
            />
            <div class="button-group">
              <button on:click|preventDefault={handleUpdateName}>Save</button>
              <button 
                type="button" 
                class="cancel" 
                on:click|preventDefault={() => cancelEdit('name')}
              >
                Cancel
              </button>
            </div>
          </div>
        {:else}
          <span>{$user.firstName} {$user.lastName}</span>
          <button class="edit-button" on:click|preventDefault={startEditingName}>
            <span class="edit-icon">
              <FaPencilAlt/>
            </span>
          </button>
        {/if}
      </div>
    </div>

    <div class="email-box">
      <label for="email">Email</label>
      <div class="email">
        {#if $editState.isEditingEmail}
          <div class="input-group">
            <input 
              id="email" 
              type="email" 
              bind:value={$editState.email} 
              placeholder="Email"
            />
            <div class="button-group">
              <button on:click|preventDefault={handleUpdateEmail}>Save</button>
              <button 
                type="button" 
                class="cancel" 
                on:click|preventDefault={() => cancelEdit('email')}
              >
                Cancel
              </button>
            </div>
          </div>
        {:else}
          <span>{$user.email}</span>
          <button class="edit-button" on:click|preventDefault={startEditingEmail}>
            <span class="edit-icon">
              <FaPencilAlt/>
            </span>
          </button>
        {/if}
      </div>
    </div>
  </div>
</div>

<style>
  .content-box {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    padding: 1rem;
  }

  .name-box, .email-box {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .name, .email {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 1rem;
  }

  .input-group {
    display: flex;
    gap: 0.5rem;
    width: 100%;
    align-items: center;
  }

  .button-group {
    display: flex;
    gap: 0.5rem;
  }

  .edit-button {
    background: none;
    border: none;
    padding: 0;
    cursor: pointer;
  }

  .edit-icon {
    width: 16px;
    height: 16px;
    color: #333;
    display: flex;
    align-items: center;
  }

  input {
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
    flex: 1;
  }

  button {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    background-color: #4CAF50;
    color: white;
    cursor: pointer;
  }

  button.cancel {
    background-color: #f44336;
  }

  .error-message {
    background-color: #ffebee;
    color: #c62828;
    padding: 1rem;
    margin-bottom: 1rem;
    border-radius: 4px;
  }

  .success-message {
    background-color: #e8f5e9;
    color: #2e7d32;
    padding: 1rem;
    margin-bottom: 1rem;
    border-radius: 4px;
  }

  label {
    font-weight: bold;
  }
</style>