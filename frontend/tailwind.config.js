/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,svelte}"],
  theme: {
    extend: {
      colors: {
        background: '#0A0A0A',
        foreground: '#F0FAFA',
        card: '#141414',
        muted: '#1C1C1C',
        'muted-foreground': '#A8C8CA',
        primary: {
          DEFAULT: '#3B82F6',
          dark: '#1D4ED8',
        },
        secondary: '#8B3D8B',
        accent: '#FF8C42',
        border: 'rgba(59, 130, 246, 0.2)',
        'border-dim': 'rgba(255, 255, 255, 0.06)',
        success: '#4ADE80',
        danger: '#FF5555',
      },
      fontFamily: {
        display: ['Baskerville', 'Libre Baskerville', 'serif'],
        body: ['Inter', 'sans-serif'],
        mono: ['JetBrains Mono', 'Fira Code', 'monospace'],
      },
    },
  },
  plugins: [],
};
