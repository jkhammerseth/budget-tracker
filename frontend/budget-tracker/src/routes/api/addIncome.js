
export async function postIncome(income) {
    try {
        const response = await fetch('http://localhost:8080/api/users/incomes', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(income),
        credentials: 'include',
        mode: 'no-cors'
        });

        if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to add income');
        }
        console.log(JSON.stringify(income))
        console.log('Income added successfully');
    } catch (error) {
        console.error('Error adding income:', error.message);
    }
    }