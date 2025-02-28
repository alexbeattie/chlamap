import { createApp } from "vue";
import VueGoogleMaps from "@fawmi/vue-google-maps";
import App from "./App.vue";
import './assets/main.css';
import router from './router'  // Make sure this import is correct

const app = createApp(App);
app.use(router)  // This line is critical - it registers the router with the app

app.use(VueGoogleMaps, {
  load: {
    key: process.env.VUE_APP_GOOGLE_MAPS_API_KEY, // ✅ Check API Key
    libraries: ["places", "geometry"], // ✅ Ensure required libraries are included
    language: "en",
    region: "US",
    v: "weekly", // ✅ Always use the latest version
    async: true,
    defer:
      true, // ✅ Ensure the async and defer attributes are set to true
  },
});

app.mount("#app");
