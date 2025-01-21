import { writable } from 'svelte/store';

export const selectionMode = writable("month");

export const selectedYear = writable(new Date().getFullYear());

export const selectedMonth = writable(new Date().getMonth());

export const selectedStartDate = writable(null);

export const selectedEndDate = writable(null);

export const selectedDay = writable(null);