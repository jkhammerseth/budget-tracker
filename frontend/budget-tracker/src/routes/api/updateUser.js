export async function updateUser(userformdata) {
  try {
    const response = await fetch("http://localhost:8080/api/users/", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(userformdata),
      credentials: 'include'
    });
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Failed to update user");
    }
  } catch (error) {
    console.error("Update user failed", error.message);
    throw error;
  }
}