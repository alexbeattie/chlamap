<template>
  <div class="flex h-screen">
    <!-- Sidebar -->
    <div
      class="w-80 flex-shrink-0 border-r border-gray-200 bg-white overflow-y-auto"
    >
      <HospitalSidebar
        :hospitals="resources"
        :selected-hospital="selectedResource"
        @hospital-selected="handleHospitalSelection"
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
    v-for="resource in filteredResources"
    :key="resource.id"
    :position="{
      lat: parseFloat(resource.latitude),
      lng: parseFloat(resource.longitude),
    }"
    @marker-click="handleMarkerClick(resource)"
    :hospital="resource"
    :is-selected="selectedResource?.id === resource.id"
  />
  <CustomInfoWindow
          v-if="selectedResource && isGoogleMapsLoaded"
          :position="{
            lat: parseFloat(selectedResource.latitude),
            lng: parseFloat(selectedResource.longitude),
          }"
          :hospital="selectedResource"
          @closeclick="handleInfoWindowClose"
        />
  <!-- ✅ Ensure CustomInfoWindow is inside Map -->
 
</Map>

    </div>

    <!-- Hospital Details Modal -->
    <HospitalDetailsModal
      v-if="showDetailsModal"
      :hospital="selectedResource"
      @close="handleModalClose"
      @get-directions="handleGetDirections"
    />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import apiClient from "@/plugins/axios";
import HospitalSidebar from "./HospitalSidebar.vue";
import HospitalDetailsModal from "./HospitalDetailsModal.vue";
import { Map } from "@fawmi/vue-google-maps";
import CustomMarker from "./CustomMarker.vue";
import CustomInfoWindow from "./CustomInfoWindow.vue";
// const { AdvancedMarkerElement, PinElement } = await google.maps.importLibrary("marker");

export default defineComponent({
  name: "MapView",

  components: {
    Map,
    CustomMarker,
    CustomInfoWindow,
    HospitalSidebar,
    HospitalDetailsModal,
  },

  data() {
    return {
      mapKey: 0, // Used to force re-render
      mapCenter: { lat: 36.7783, lng: -119.4179 },
      resources: [],
      selectedResource: null,
      showDetailsModal: false,
      isGoogleMapsLoaded: false, // Track if the API is loaded
      map: null,
    };
  },
  computed: {
    googleMapsStatus() {
      return !!window.google && !!window.google.maps;
    },
   filteredResources() {
    return this.isGoogleMapsLoaded ? this.resources : [];
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
        this.fetchResources();
      })
      .catch((error) => {
        console.error("Google Maps API failed to load:", error);
      });

    // Continue with the rest of your mounted logic
    this.fetchResources();
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
        this.selectedResource = null;
      }
    },
    //es lint-disable-next-lin
    handleWindowFocus() {
      if (this.$refs.mapRef && window.google) {
        // Add window. prefix
        const map = this.$refs.mapRef.$mapObject;
        if (map) {
          window.google.maps.event.trigger(map, "resize"); // Add window. prefix
          if (this.selectedResource) {
            map.panTo({
              lat: parseFloat(this.selectedResource.latitude),
              lng: parseFloat(this.selectedResource.longitude),
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

      // Uncomment if you need a test marker
      // const testMarker = new window.google.maps.Marker({
      //   position: { lat: 36.7783, lng: -119.4179 },
      //   map: map,
      //   title: "Test Marker",
      // });
    },

    async fetchResources() {
      try {
        console.log("Fetching resources from:", apiClient.defaults.baseURL + "/regional-centers");
        const { data } = await apiClient.get("/regional-centers");
        console.log("Raw resources data:", data);

        this.resources = (Array.isArray(data) ? data : []).map((resource) => ({
          ...resource,
          latitude: parseFloat(resource.latitude),
          longitude: parseFloat(resource.longitude),
        }));
        this.mapKey++; // Force Map to re-render after data load


        if (this.resources.length > 0) {
          this.mapCenter = {
            lat: this.resources[0].latitude,
            lng: this.resources[0].longitude,
          };
        }
      } catch (error) {
        console.error("Failed to fetch resources:", error);
      }
    },

    handleHospitalSelection(hospital) {
      this.selectedResource = hospital;
      if (this.$refs.mapRef?.$mapObject && hospital) {
        const position = {
          lat: parseFloat(hospital.latitude),
          lng: parseFloat(hospital.longitude),
        };
        this.$refs.mapRef.$mapObject.panTo(position);
        this.$refs.mapRef.$mapObject.setZoom(12);
      }
    },

    handleMapClick() {
      if (this.selectedResource) {
        this.selectedResource = null;
      }
    },

handleMarkerClick(hospital) {
  try {
    if (!hospital || !hospital.latitude || !hospital.longitude) {
      console.error("Invalid hospital object:", hospital);
      return;
    }

    console.log("Marker clicked:", hospital);
    this.selectedResource = hospital; // ✅ Update state

    this.$nextTick(() => {
      const sidebarItem = document.querySelector(
        `[data-hospital-id="${hospital.id}"]`
      );
      if (sidebarItem) {
        sidebarItem.scrollIntoView({ behavior: "smooth", block: "center" });
      }
    });

    console.log("Selected Resource Updated:", this.selectedResource);
  } catch (error) {
    console.error("Error handling marker click:", error);
  }
},


  handleInfoWindowClose() {
  setTimeout(() => {
    this.selectedResource = null;
  }, 100); // ✅ Delay to prevent Vue from removing too fast
},
    handleModalClose() {
      this.showDetailsModal = false;
    },

    handleGetDirections() {
      if (this.selectedResource) {
        const { latitude, longitude } = this.selectedResource;
        if (latitude == null || longitude == null) {
          console.error("Invalid coordinates for selected resource.");
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
