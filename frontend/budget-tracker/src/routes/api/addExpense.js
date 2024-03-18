
 export async function postExpense(expense) {
    try {
        const response = await fetch('http://localhost:8080/api/users/expenses', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(expense),
        credentials: 'include',
        });

        if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to add expense');
        }
        console.log(JSON.stringify(expense))
        console.log('Expense added successfully');
    } catch (error) {
        console.error('Error adding expense:', error.message);
    }
    }