/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primaryColor: '#12372A',
        bg1: '#436850',
        bg2: '#ADBC9F',
        accent: '#FBFADA',
      },
      fontFamily: {
        sans: ['CustomFont', 'sans-serif'],
      },
    },
  },
  plugins: [],
}
