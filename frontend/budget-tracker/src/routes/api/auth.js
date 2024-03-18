import { user } from '../../stores/user.js';

import { goto } from '$app/navigation';

export async function login(username, password) {
  try {
    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
      credentials: 'include',
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Login failed");
    }
    
    user.set({ loggedIn: true, userId: null});

    goto('/');
  } catch (error) {
    console.error("Login failed", error.message);
    throw error;
  }
}

export function logout() {
  try {
    fetch("http://localhost:8080/api/auth/logout", {
      method: "POST",
      credentials: 'include', 
    }).then((response) => {
      if (response.ok) {
        user.set({ loggedIn: false, userId: null });
        goto('/login');
      } else {
        console.error("Logout failed");
        goto('/');
      }
    }).catch((error) => {
      console.error("Logout error", error);
    });
  }
  catch (error) {
    console.error("Logout failed", error.message);
    throw error;
  }
}


export async function checkAuthStatus() {
  try {
    const response = await fetch('http://localhost:8080/api/auth/status', {
      credentials: 'include', // Necessary for cookies to be sent with the request
    });
    if (response.ok) {
      const data = await response.json();
      user.set({ 
        loggedIn: data.isLoggedIn, 
        userId: data.user.id, 
        username: data.user.username, // Adjusted assuming username is within data.user
        firstName: data.user.firstName, 
        lastName: data.user.lastName,
        email: data.user.email
      });
    } else {
      // Handle non-OK responses, e.g., by clearing the user store
      user.set({
        loggedIn: false,
        userId: null,
        username: null,
        firstName: null,
        lastName: null,
        email: null,
      });
    }
  } catch (error) {
    console.error('Failed to check authentication status:', error);
    // Consider how to handle errors, possibly resetting the user store or showing a message
  }
}
