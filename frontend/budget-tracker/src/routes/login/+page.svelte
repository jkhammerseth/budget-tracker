<script>
  import { login } from '../api/auth';

  let username = '';
  let password = '';
  let errorMessage = '';

  async function handleLogin() {
    try {
      await login(username, password);
    } catch (error) {
      errorMessage = error.message;
    }
  }

</script>

<div class="login-container">
  <h2>Welcome Back!</h2>
  
  <input
    class="login-input"
    type="text"
    bind:value={username}
    placeholder="Username"
  />
  <input
    class="login-input"
    type="password"
    bind:value={password}
    placeholder="Password"
  />
  {#if errorMessage}
    <div class="login-error">{errorMessage}</div>
  {/if}
  <button class="login-button" on:click={handleLogin}>Login</button>
  <span class="no-account">Don't have an account?</span>

  <a href="/signup" class="login-link">Create an Account!</a>

</div>


<style>
  .login-container {
    max-width: 400px;
    margin: 100px auto;
    padding: 50px;
    background: white;
    border-radius: 8px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  }

  .login-container h2 {
    margin-bottom: 20px;
    text-align: center;
    font-size: 24px;
    color: #333;
  }

  .login-input {
    width: 100%;
    padding: 15px;
    margin-bottom: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
    box-sizing: border-box;
  }

  .login-button {
    width: 100%;
    padding: 15px;
    background-color: #0052cc;
    border: none;
    border-radius: 8px;
    color: white;
    cursor: pointer;
    margin-bottom: 10px;
  }

  .login-button:hover {
    background-color: #003399;
  }

  .login-link {
    color: #0052cc;
    text-decoration: none;
  }

  .login-error {
    color: red;
    font-size: 0.9em;
    margin-top: -15px;
    margin-bottom: 20px;
    text-align: center;
  }

  .no-account, .login-link {
    
    display: flex;
    justify-content: center;
    align-items: space-between;
    
    padding: 2px;
  }
</style>