import { mount } from 'svelte'
import './assets/tailwind.min.css' // Hier laden wir das lokale Tailwind-CDN
import './style.css'
import App from './App.svelte'

const app = mount(App, {
  target: document.getElementById('app')
})

export default app