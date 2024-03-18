import { writable } from 'svelte/store';

export const selectionMode = writable(null);

export const selectedYear = writable(null);

export const selectedMonth = writable(null);

export const selectedStartDate = writable(null);

export const selectedEndDate = writable(null);

export const selectedDay = writable(null);