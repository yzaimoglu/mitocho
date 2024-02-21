/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templ/**/*.{templ,html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
}

