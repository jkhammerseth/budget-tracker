export async function postLoan(loan) {
    try {
        const response = await fetch('http://localhost:8080/api/users/loan', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(loan),
            credentials: 'include',
        });

        if (!response.ok) {
            throw new Error('Failed to add loan');
        }

        return response;
    } catch (error) {
        console.error('Error adding loan:', error.message);
        throw error;
    }
}