export async function updateExpense(expenseData) {
    try {
        console.log("Updating expense with data:", expenseData);

        const response = await fetch(`http://localhost:8080/api/users/expenses/${expenseData.ID}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(expenseData), // Send only fields provided
            credentials: "include",
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || "Failed to update expense");
        }

        console.log("Expense updated successfully");
        return true; // Indicate success
    } catch (error) {
        console.error("Error updating expense:", error.message);
        throw error;
    }
}
