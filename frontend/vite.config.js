import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// Ganz simples Svelte-Setup ohne zickige Tailwind-Plugins
export default defineConfig({
  plugins: [svelte()]
})