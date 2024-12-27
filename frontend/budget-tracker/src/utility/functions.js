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
  return '-' + amount + ' kr';
}

export function daysFromNow(isoString) {
  const targetDate = new Date(isoString); // Parse the ISO string to a date object
  const today = new Date();

  // Reset time parts of both dates to only consider the date difference
  today.setHours(0, 0, 0, 0);
  targetDate.setHours(0, 0, 0, 0);

  const diffInTime = targetDate - today; // Time difference in milliseconds
  const diffInDays = Math.round(diffInTime / (1000 * 60 * 60 * 24)); // Convert to days

  if (diffInDays === 0) {
    return 'Today';
  } else if (diffInDays === -1) {
    return 'Yesterday';
  } else if (diffInDays === 1) {
    return 'Tomorrow';
  } else if (diffInDays > 0) {
    return `In ${diffInDays} days`;
  } else {
    return `${Math.abs(diffInDays)} days ago`;
  }
}
