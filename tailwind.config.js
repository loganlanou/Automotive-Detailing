/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/templates/**/*.templ",
    "./web/static/**/*.js",
  ],
  theme: {
    extend: {
      colors: {
        'brand-primary': '#003DA5',
        'brand-secondary': '#111827',
        'brand-accent': '#60A5FA',
        'brand-accent-bright': '#3B82F6', // Brighter blue for CTAs
        'brand-green': '#10B981', // Green accent like exquisitedetailing
        'brand-bg': '#0B0F13',
        'brand-fg': '#F3F4F6',
        'muted': '#9CA3AF',
        'border': '#1F2937',
      },
      fontFamily: {
        heading: ['Poppins', 'sans-serif'],
        body: ['Inter', 'sans-serif'],
        script: ['Dancing Script', 'cursive'],
      },
      borderRadius: {
        DEFAULT: '0.375rem', // Squared corners for rigid aesthetic
        'sm': '0.25rem',
        'md': '0.375rem',
        'lg': '0.5rem',
      },
      boxShadow: {
        'sm': '0 1px 2px rgba(0,0,0,.05)',
        'md': '0 4px 12px rgba(0,0,0,.08)',
      },
    },
  },
  plugins: [],
}
