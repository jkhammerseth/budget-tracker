export function fromISOString(isoString) {
    const date = new Date(isoString);

    // Array of month names for better readability
    const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

    const day = date.getDate().toString().padStart(2, '0');
    const month = monthNames[date.getMonth()]; // Get month name
    const year = date.getFullYear(); // Full year (e.g., 2024)

    return `${month} ${day}`;
  }

  export function formatAmount(amount) {
    return amount.toLocaleString('en-US') + ' kr';
  }

  export function formatExpenseAmount(amount) {
    return '- ' + amount + ' kr';
  }