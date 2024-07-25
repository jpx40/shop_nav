/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,js,ts,tsx,jsx,svelte,templ}",
    "./pkg/**/*.{html,js,ts,tsx,jsx,svelte,templ}",
    "./template/**/*.{html,js,ts,tsx,jsx,svelte,templ}",
    "./stactic/**/*.{html,js,ts,tsx,jsx,svelte,templ}",
    "./lib/**/*.{html,js,ts,tsx,jsx,svelte,templ}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
