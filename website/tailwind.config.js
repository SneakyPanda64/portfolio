/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				's-blue': '#7289DA',
				's-light-gray': '#6f747a',
				's-gray': '#282B30',
				's-dark-gray': '#1E2124',
				's-purple': '#8C60C4'
			}
		}
	},
	plugins: []
};
