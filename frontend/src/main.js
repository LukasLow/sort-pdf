import { mount } from "svelte";
import "./style.css"; // Das lädt jetzt das dynamische Tailwind über PostCSS
import App from "./App.svelte";

const app = mount(App, {
  target: document.getElementById("app"),
});

export default app;
