import { createApp } from "vue";
import VueGoogleMaps from "@fawmi/vue-google-maps";
import App from "./App.vue";
import './assets/main.css'

const app = createApp(App);

app.use(VueGoogleMaps, {
  load: {
    key: process.env.VUE_APP_GOOGLE_MAPS_API_KEY,
    libraries: ["places", "geometry"],
    language: "en",
    region: "US",
    
  },
});

app.mount("#app");