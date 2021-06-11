module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      spacing: {
        '130': '30rem',
        '140': '40rem',
        '150': '50rem',
        '160': '60rem',
      },
      colors: {
        'deep-purple': '#190624',
        'base-purple': '#2A1437',
        'light-purple': '#ad92b7',
        'base-green': '#00ff4a',
        'base-red': '#ff005c',
      },
      fontFamily: {
        'nova-light': 'nova-light',
        'nova': 'nova',
        'nova-semi': 'nova-semi',
        'nova-bold': 'nova-bold',
        'nova-extra': 'nova-extra',
        'nova-black': 'nova-black',
      },
      boxShadow: {
        brand: 'rgba(0, 0, 0, 0.12) 0px 1px 3px, rgba(0, 0, 0, 0.24) 0px 1px 2px;',
        big: 'rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;',
        bigger: 'rgba(0, 0, 0, 0.19) 0px 10px 20px, rgba(0, 0, 0, 0.23) 0px 6px 6px;'
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}