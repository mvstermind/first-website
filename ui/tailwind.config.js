/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./html/*.html"],
  theme: {
    extend: {
      colors:{
         'shamrock': {
        '50': '#eefbf5',
        '100': '#d5f6e6',
        '200': '#afebd1',
        '300': '#7bdab8',
        '400': '#4ac49c',
        '500': '#22a780',
        '600': '#148767',
        '700': '#106c55',
        '800': '#0f5645',
        '900': '#0e4639',
        '950': '#062821',
    },
      }
    },
  },
  plugins: [],
}

