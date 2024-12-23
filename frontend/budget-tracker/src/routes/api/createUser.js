import { goto } from '$app/navigation';

export async function createUser(userformdata) {
  try {
    const response = await fetch("http://localhost:8080/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(userformdata)
    });
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Create user failed");
    }
    goto('/login');
  } catch (error) {
    console.error("Create user failed", error.message);
    throw error;
  }
}