/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ["./templ/**/*.{templ,html,js}", "./node_modules/flowbite/**/*.js"],
  theme: {
    extend: {},
  },
  plugins: [require("flowbite")],
}
