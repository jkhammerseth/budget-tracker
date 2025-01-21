export async function deleteIncome(incomeID) {
    try {
        const response = await fetch(`http://localhost:8080/api/users/incomes/${incomeID}`, {
        method: 'DELETE',
        credentials: 'include',
        });

        if (!response.status == 204) {
            const errorData = await response.json();
            throw new Error(errorData.message || 'Failed to delete income');
          }
        console.log('Income deleted successfully');
        return response
    } catch (error) {
        console.error('Error deleting income:', error.message);
    }
}