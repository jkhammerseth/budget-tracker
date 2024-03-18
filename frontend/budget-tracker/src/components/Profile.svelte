<script>
  import { user } from '../stores/user';
  import FaPencilAlt from 'svelte-icons/fa/FaPencilAlt.svelte';
  import { writable } from 'svelte/store';

  let isEditingName = false;
  let isEditingEmail = false;

  const handleUpdateName = async () => {
      isEditingName = false;
  };

  const handleUpdateEmail = async () => {
      // Similar to handleUpdateName
      isEditingEmail = false;
  };

  const handleChangeName = () => {
      isEditingName = true;
  };

  const handleChangeEmail = () => {
      isEditingEmail = true;
  };
</script>

<div class="profile-container">
  <h2>Profile</h2>
  <hr />
  <div class="content-box">
      <div class="name-box">
          <label class="name label" for="name">Name</label>
          <div class="name">
              {#if isEditingName}
                  <input id="firstName" type="text" bind:value={$user.firstName} />
                  <input id="lastName" type="text" bind:value={$user.lastName} />
                  <button on:click={handleUpdateName}>Save</button>
              {:else}
                  <span>{$user.firstName + ' ' + $user.lastName}</span>
                  <span class="edit-icon" on:click={handleChangeName}><FaPencilAlt/></span>
              {/if}
          </div>
      </div>
      <div class="email-box">
        <label for="email-label">Email</label>
        <div class="email">
        {#if isEditingEmail}
            
              <input id="email" type="text" bind:value={$user.email} />
              <button on:click={handleUpdateEmail}>Save</button>
          {:else}
              <span>{$user.email}</span>
              <span class="edit-icon" on:click={handleChangeEmail}><FaPencilAlt/></span>
          {/if}
        </div>
      </div>
  </div>
</div>

  
  <style>
    .content-box {
      display: flex;
      flex-direction: column;
      align-items: left;
      margin-bottom: 20px;
    }

    .name{
      display: flex;
      justify-content: space-between;
    }

    .email{
      display: flex;
      justify-content: space-between;
    }

    .name-box {
      display: flex;
      flex-direction: column;
      align-items: left;
      margin-bottom: 20px;
    }

    .edit-icon {
      width: 16px;
      height: 16px;
      color: #333;
      cursor: pointer;
      background: none;
    }
    
    .email-box {
      display: flex;
      flex-direction: column;
      align-items: left;
      margin-bottom: 20px;
    }

    .profile-container {
        margin: 0 auto;
        padding-left: 20px;
        border-radius: 4px;
    }
  
    label {
      font-weight: bold;
    }
  </style>
  