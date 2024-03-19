export async function postExpense(expense) {
    try {
        const response = await fetch('http://localhost:8080/api/users/expense', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(expense),
        credentials: 'include',
        });

        if (!response.ok) {
            throw new Error('Failed to add expense');
        }
        return response
    } catch (error) {
        console.error('Error adding expense:', error.message);
    }
}

export async function postListofExpenses(expenses) {
    try {
        const response = await fetch('http://localhost:8080/api/users/expenses', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(expenses),
        credentials: 'include',
        });

        if (!response.ok) {
            throw new Error('Failed to add expenses');
        }
        return response;
    } catch (error) {
        console.error('Error adding expenses:', error.message);
    }
}