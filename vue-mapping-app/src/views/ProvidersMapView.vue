<template>
  <div class="flex h-screen">
    <!-- Sidebar -->
    <div
      class="w-80 flex-shrink-0 border-r border-gray-200 bg-white overflow-y-auto"
    >
      <ProviderSidebar
        :providers="providers"
        :selected-provider="selectedProvider"
        @provider-selected="handleProviderSelection"
      />
    </div>

    <!-- Main Content (Map) -->
    <div class="flex-1">
      <!-- Only render the Map component if the API is loaded -->
      <Map
        :center="mapCenter"
        :zoom="6"
        class="w-full h-full"
        ref="mapRef"
        @loaded="onMapLoaded"
        @click="handleMapClick"
      >
        <CustomMarker
          v-for="provider in filteredProviders"
          :key="provider.id"
          :position="{
            lat: parseFloat(provider.latitude),
            lng: parseFloat(provider.longitude),
          }"
          @marker-click="handleMarkerClick(provider)"
          :provider="provider"
          :is-selected="selectedProvider?.id === provider.id"
        />
        <CustomInfoWindow
          v-if="selectedProvider && isGoogleMapsLoaded"
          :position="{
            lat: parseFloat(selectedProvider.latitude),
            lng: parseFloat(selectedProvider.longitude),
          }"
          :provider="selectedProvider"
          @closeclick="handleInfoWindowClose"
        />
      </Map>
    </div>

    <!-- Provider Details Modal -->
    <ProviderDetailsModal
      v-if="showDetailsModal"
      :provider="selectedProvider"
      @close="handleModalClose"
      @get-directions="handleGetDirections"
    />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import apiClient from "@/plugins/axios";
import ProviderSidebar from "./ProviderSidebar.vue";
import ProviderDetailsModal from "./ProviderDetailsModal.vue";
import { Map } from "@fawmi/vue-google-maps";
import CustomMarker from "../components/CustomMarker.vue";
import CustomInfoWindow from "./CustomInfoWindow.vue";

export default defineComponent({
  name: "ProvidersMapView",

  components: {
    Map,
    CustomMarker,
    CustomInfoWindow,
    ProviderSidebar,
    ProviderDetailsModal,
  },

  data() {
    return {
      mapKey: 0, // Used to force re-render
      mapCenter: { lat: 34.0522, lng: -118.2437 }, // Los Angeles as default
      providers: [],
      selectedProvider: null,
      showDetailsModal: false,
      isGoogleMapsLoaded: false, // Track if the API is loaded
      map: null,
    };
  },
  
  computed: {
    googleMapsStatus() {
      return !!window.google && !!window.google.maps;
    },
    filteredProviders() {
      return this.isGoogleMapsLoaded ? this.providers : [];
    },
  },
  
  created() {
    console.log(
      "Google Maps API available:",
      !!window.google && !!window.google.maps
    );
  },

  mounted() {
    // Check if Maps is already loaded first
    this.$gmapApiPromiseLazy()
      .then(() => {
        console.log("Google Maps API has loaded!");
        this.isGoogleMapsLoaded = true;
        this.fetchProviders();
      })
      .catch((error) => {
        console.error("Google Maps API failed to load:", error);
      });

    // Continue with the rest of your mounted logic
    this.fetchProviders();
    window.addEventListener("focus", this.handleWindowFocus);
    window.addEventListener("click", this.handleGlobalClick);
  },

  beforeUnmount() {
    window.removeEventListener("focus", this.handleWindowFocus);
    window.removeEventListener("click", this.handleGlobalClick);
    this.map = null;
  },

  methods: {
    handleGlobalClick(e) {
      const isMarkerClick = e.target.closest(".custom-marker");
      const isInfoWindowClick = e.target.closest(".gm-style-iw");
      if (!isMarkerClick && !isInfoWindowClick) {
        this.selectedProvider = null;
      }
    },

    handleWindowFocus() {
      if (this.$refs.mapRef && window.google) {
        const map = this.$refs.mapRef.$mapObject;
        if (map) {
          window.google.maps.event.trigger(map, "resize");
          if (this.selectedProvider) {
            map.panTo({
              lat: parseFloat(this.selectedProvider.latitude),
              lng: parseFloat(this.selectedProvider.longitude),
            });
          }
        }
      }
    },

    onMapLoaded(map) {
      if (!window.google || !window.google.maps) {
        console.error("Google Maps API not loaded.");
        return;
      }
      console.log("Map loaded:", map);
      this.map = map;
    },

async fetchProviders() {
  try {
    console.log("Fetching providers from:", apiClient.defaults.baseURL + "/providers");
    
    const { data } = await apiClient.get("/providers"); // ✅ Fix: Fetch from `/providers`
    console.log("Raw providers data:", data);

    // ✅ Ensure providers have valid latitude & longitude
    this.providers = (Array.isArray(data) ? data : []).map((provider) => ({
      ...provider,
      latitude: provider.latitude ? parseFloat(provider.latitude) : 0,
      longitude: provider.longitude ? parseFloat(provider.longitude) : 0,
      areas: Array.isArray(provider.areas) ? provider.areas : [],
    }));
    
    this.mapKey++; // Force Map to re-render after data load

    // ✅ Adjust map center to first provider's location
    if (this.providers.length > 0) {
      this.mapCenter = {
        lat: this.providers[0].latitude,
        lng: this.providers[0].longitude,
      };
    }
  } catch (error) {
    console.error("Failed to fetch providers:", error);
  }
},

    handleProviderSelection(provider) {
      this.selectedProvider = provider;
      if (this.$refs.mapRef?.$mapObject && provider) {
        const position = {
          lat: parseFloat(provider.latitude),
          lng: parseFloat(provider.longitude),
        };
        this.$refs.mapRef.$mapObject.panTo(position);
        this.$refs.mapRef.$mapObject.setZoom(12);
      }
    },

    handleMapClick() {
      if (this.selectedProvider) {
        this.selectedProvider = null;
      }
    },

    handleMarkerClick(provider) {
      try {
        if (!provider || !provider.latitude || !provider.longitude) {
          console.error("Invalid provider object:", provider);
          return;
        }

        console.log("Marker clicked:", provider);
        this.selectedProvider = provider;

        this.$nextTick(() => {
          const sidebarItem = document.querySelector(
            `[data-provider-id="${provider.id}"]`
          );
          if (sidebarItem) {
            sidebarItem.scrollIntoView({ behavior: "smooth", block: "center" });
          }
        });

        console.log("Selected Provider Updated:", this.selectedProvider);
      } catch (error) {
        console.error("Error handling marker click:", error);
      }
    },

    handleInfoWindowClose() {
      setTimeout(() => {
        this.selectedProvider = null;
      }, 100);
    },
    
    handleModalClose() {
      this.showDetailsModal = false;
    },

    handleGetDirections() {
      if (this.selectedProvider) {
        const { latitude, longitude } = this.selectedProvider;
        if (latitude == null || longitude == null) {
          console.error("Invalid coordinates for selected provider.");
          return;
        }

        const mapWindow = window.open(
          `https://www.google.com/maps/dir/?api=1&destination=${latitude},${longitude}`,
          "_blank"
        );

        if (mapWindow) {
          const checkWindow = setInterval(() => {
            if (mapWindow.closed) {
              clearInterval(checkWindow);
              this.handleWindowFocus();
            }
          }, 1000);
        }
      }
    },
  },
});
</script>

<style scoped>
.vue-map-container {
  height: 100%;
  width: 100%;
}

.gm-style-iw {
  max-width: 90vw !important;
}

.gm-style-iw-d {
  max-width: 100% !important;
  overflow: auto !important;
}
</style>