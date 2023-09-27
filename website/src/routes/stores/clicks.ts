import { writable } from 'svelte/store';
export const clicks = writable({ disabled: true, clicks: 0 });
