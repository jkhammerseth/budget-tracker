export async function updateIncome(incomeData) {
    try {
        console.log("Updating income with data:", incomeData);

        const response = await fetch(`http://localhost:8080/api/users/incomes/${incomeData.ID}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(incomeData), 
            credentials: "include",
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || "Failed to update income");
        }

        console.log("income updated successfully");
        return true;
    } catch (error) {
        console.error("Error updating income:", error.message);
        throw error;
    }
}
