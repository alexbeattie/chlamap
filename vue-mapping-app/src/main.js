import { createApp } from "vue";
import { createPinia } from 'pinia';
import VueGoogleMaps from "@fawmi/vue-google-maps";
import App from "./App.vue";
import './assets/main.css';
import router from './router'
import { useAuthStore } from './store/auth';

const app = createApp(App);

// Setup Pinia first
const pinia = createPinia();
app.use(pinia);
app.use(router);

// Configure Google Maps
app.use(VueGoogleMaps, {
  load: {
    key: process.env.VUE_APP_GOOGLE_MAPS_API_KEY,
    libraries: ["places", "geometry"],
    language: "en",
    region: "US",
    v: "weekly",
    async: true,
    defer: true,
  },
});

// Get the auth store AFTER Pinia is initialized
const authStore = useAuthStore();

// Only mount once after auth initialization
authStore.initialize().finally(() => {
  app.mount('#app');
});