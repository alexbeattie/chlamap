<!-- MapView.vue -->
<template>
  <div class="flex h-screen">
    <!-- Sidebar -->
    <div class="w-80 flex-shrink-0 border-r border-gray-200 bg-white overflow-y-auto">
      <HospitalSidebar 
        :hospitals="resources"
        :selected-hospital="selectedResource"
        @hospital-selected="handleHospitalSelection" 
      />
    </div>

    <!-- Main Content (Map) -->
    <div class="flex-1">
      <Map
        :center="mapCenter"
        :zoom="6"
        class="w-full h-full"
        ref="mapRef"
        @loaded="onMapLoaded"
        @click="handleMapClick"
      >
        <CustomMarker
          v-for="resource in resources"
          :key="resource.id"
          :position="{
            lat: typeof resource.latitude === 'string' ? parseFloat(resource.latitude) : resource.latitude,
            lng: typeof resource.longitude === 'string' ? parseFloat(resource.longitude) : resource.longitude
          }"
          @marker-click="handleMarkerClick(resource)"
          :hospital="resource"
          :is-selected="selectedResource?.id === resource.id"
        />

        <CustomInfoWindow
          v-if="selectedResource"
          :position="{
            lat: typeof selectedResource.latitude === 'string' ? parseFloat(selectedResource.latitude) : selectedResource.latitude,
            lng: typeof selectedResource.longitude === 'string' ? parseFloat(selectedResource.longitude) : selectedResource.longitude
          }"
          @closeclick="handleInfoWindowClose"
          :hospital="selectedResource"
          @show-details="showDetailsModal = true"
        >
          <div class="p-4">
            <h3 class="font-semibold mb-2">{{ selectedResource.name }}</h3>
            <p>{{ selectedResource.address }}</p>
            <div class="flex gap-2 mt-2">
              <button 
                @click="handleGetDirections"
                class="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
              >
                Get Directions
              </button>
              <button 
                @click="showDetailsModal = true"
                class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
              >
                View Details
              </button>
            </div>
          </div>
        </CustomInfoWindow>
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
      mapCenter: { lat: 36.7783, lng: -119.4179 },
      resources: [],
      selectedResource: null,
      showDetailsModal: false,
      map: null,
    };
  },

  async mounted() {
    await this.fetchResources();
    window.addEventListener("focus", this.handleWindowFocus);
  },
  beforeUnmount() {
    // Remove event listener when component is destroyed
    window.removeEventListener("focus", this.handleWindowFocus);
  },

  methods: {
    handleGlobalClick(e) {
      // Only close if clicking outside of markers/info windows
      const isMarkerClick = e.target.closest('.custom-marker');
      const isInfoWindowClick = e.target.closest('.gm-style-iw');
      if (!isMarkerClick && !isInfoWindowClick) {
        this.selectedResource = null;
      }
    },

    handleWindowFocus() {
      // Refresh the map instance when window regains focus
      if (this.$refs.mapRef) {
        const map = this.$refs.mapRef.$mapObject;
        if (map) {
          // eslint-disable-next-line
          google.maps.event.trigger(map, "resize");
          if (this.selectedResource) {
            map.panTo({
              lat: this.selectedResource.latitude,
              lng: this.selectedResource.longitude,
            });
          }
        }
      }
    },
    onMapLoaded(map) {
      console.log("Map loaded:", map);
      this.map = map;
    },
    async fetchResources() {
      try {
        const { data } = await apiClient.get("/api/regional-centers");
        console.log("Raw resources data:", data);

        this.resources = (Array.isArray(data) ? data : []).map((resource) => ({
          ...resource,
          latitude:
            typeof resource.latitude === "string"
              ? parseFloat(resource.latitude)
              : resource.latitude,
          longitude:
            typeof resource.longitude === "string"
              ? parseFloat(resource.longitude)
              : resource.longitude,
        }));

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

      // Update map center and zoom to focus on selected hospital
      if (this.map && hospital) {
        const position = {
          lat:
            typeof hospital.latitude === "string"
              ? parseFloat(hospital.latitude)
              : hospital.latitude,
          lng:
            typeof hospital.longitude === "string"
              ? parseFloat(hospital.longitude)
              : hospital.longitude,
        };

        this.map.panTo(position);
        this.map.setZoom(12); // Zoom in when selecting a hospital
      }
    },

  

   handleMapClick() {
      // Only close if we have a selected resource
      if (this.selectedResource) {
        this.selectedResource = null;
      }
    },



    handleMarkerClick(hospital) {
      // Set the selected resource without trying to handle the event
      this.selectedResource = hospital;
      
      this.$nextTick(() => {
        const sidebarItem = document.querySelector(`[data-hospital-id="${hospital.id}"]`);
        if (sidebarItem) {
          sidebarItem.scrollIntoView({ behavior: 'smooth', block: 'center' });
        }
      });
    },
    handleInfoWindowClose() {
      this.selectedResource = null;
    },
   

    handleModalClose() {
      this.showDetailsModal = false;
    },

    handleGetDirections() {
      if (this.selectedResource) {
        const { latitude, longitude } = this.selectedResource;
        // Open directions in a new window and store the window reference
        const mapWindow = window.open(
          `https://www.google.com/maps/dir/?api=1&destination=${latitude},${longitude}`,
          "_blank"
        );

        // If we got a window reference, we can add a listener for when it closes
        if (mapWindow) {
          const checkWindow = setInterval(() => {
            if (mapWindow.closed) {
              clearInterval(checkWindow);
              // Refresh the map when the directions window is closed
              this.handleWindowFocus();
            }
          }, 1000);
        }
      }
    },
  },
});
</script>
<style>
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
