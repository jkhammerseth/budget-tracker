export async function deleteExpense(expenseID) {
    try {
        const response = await fetch(`http://localhost:8080/api/users/expenses/${expenseID}`, {
        method: 'DELETE',
        credentials: 'include',
        });

        if (!response.status == 204) {
            const errorData = await response.json();
            throw new Error(errorData.message || 'Failed to delete expense');
          }
        console.log('Expense deleted successfully');
        return response
    } catch (error) {
        console.error('Error deleting expense:', error.message);
    }
}