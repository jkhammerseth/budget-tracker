<script>
    import { createUser } from '../api/createUser';
  
    let firstname = '';
    let lastname = '';
    let email = '';
    let username = '';
    let password = '';
    let retypedPassword = '';
    let errorMessage = '';
  
    async function handleCreateUser() {
      const userData = {
        "FirstName" : firstname,
        "LastName" : lastname,
        "Email" : email,
        "Username": username,
        "Password" : password,
      };
      
      try {
        await createUser(userData);
        window.location.href = '/login'; 
      } catch (error) {
        errorMessage = error.message; 
      }
    }

    // TODO: implement
    function passwordMatch() {
      return password === retypedPassword;
    }

  </script>
  
  <style>
    .signup-container {
      max-width: 400px;
      margin: 50px auto;
      padding: 30px;
      background-color: #E8EFF1;
      border-radius: 12px;
      box-shadow: 0 8px 16px rgba(0,0,0,0.1); 
      font-family: 'Roboto', sans-serif;
    }
  
    h2 {
      text-align: center;
      color: #2C3E50;
      margin-bottom: 25px;
    }
  
    input {
      width: 100%;
      padding: 10px;
      margin-bottom: 20px;
      border: 1px solid #BDC3C7;
      border-radius: 4px;
      background-color: #FFFFFF;
    }
  
    .signup-button {
      width: 100%;
      padding: 12px;
      background-color: #3498DB; 
      color: white;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      font-weight: bold;
    }
  
    .signup-button:hover {
      background-color: #2980B9; 
    }
  
    .signup-link {
      display: block;
      text-align: center;
      margin-top: 20px;
      color: #3498DB;
      text-decoration: none;
    }
  
    .signup-link:hover {
      text-decoration: underline;
    }
  
    .error-message {
      color: #E74C3C;
      text-align: center;
      margin-top: 15px;
    }
  </style>
  
  <div class="signup-container">
    <h2>Create Your Account</h2>
    <form on:submit|preventDefault={handleCreateUser}>
      <input type="text" bind:value={firstname} placeholder="First Name" />
      <input type="text" bind:value={lastname} placeholder="Last Name" />
      <input type="email" bind:value={email} placeholder="Email" />
      <input type="text" bind:value={username} placeholder="Username" />
      <input type="password" bind:value={password} placeholder="Password" />
      
      
      <button type="submit" class="signup-button">Sign Up</button>
    </form>
    {#if errorMessage}
      <p class="error-message">{errorMessage}</p>
    {/if}
    <a href="/login" class="signup-link">Already have an account? Log in</a>
  </div>
  