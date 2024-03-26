import { loans } from '../../stores/loans.js';

export async function fetchLoans() {
    try {
        const response = await fetch('http://localhost:8080/api/users/loans', {
            method: 'GET',
            credentials: 'include',
        });

        if (!response.ok) {
            throw new Error('Failed to fetch loans');
        }

        const data = await response.json();
        loans.set(data);

        } catch (error) {
            console.error('Error fetching loans:', error.message);
            throw error;
    }
}