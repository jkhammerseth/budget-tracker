export async function postIncome(income) {
    try {
        const response = await fetch('http://localhost:8080/api/users/income', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(income),
            credentials: 'include',
        });

        if (!response.ok) {
            throw new Error('Failed to add income');
        }

        console.log('Income added successfully');
        return response;
    } catch (error) {
        console.error('Error adding income:', error.message);
        throw error;
    }
}

export async function postListofIncomes(incomes) {
    try {
        const response = await fetch('http://localhost:8080/api/users/incomes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(incomes),
            credentials: 'include',
        });

        if (!response.ok) {
            throw new Error('Failed to add incomes');
        }

        console.log('Incomes added successfully');
        return response;
    } catch (error) {
        console.error('Error adding incomes:', error.message);
        throw error;
    }
}