/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./components/**/*.templ", "internal/services/markdown.go"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}

