/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ["./templ/**/*.{templ,html,js}", "./templ/html/*"],
  theme: {
    extend: {},
  },
  plugins: [require("flowbite")],
}
