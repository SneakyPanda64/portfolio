import { writable } from 'svelte/store';
export const clicks = writable({ disabled: false, clicks: 0 });
