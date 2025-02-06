<!-- MapView.vue -->
<template>
  <div class="flex h-screen">
    <!-- Sidebar -->
    <div class="w-80 flex-shrink-0 border-r border-gray-200 bg-white overflow-y-auto">
      <HospitalSidebar 
        :hospitals="resources" 
        @hospital-selected="handleHospitalSelection" 
      />
    </div>

    <!-- Main Content (Map) -->
    <div class="flex-1">
      <Map
        :center="mapCenter"
        :zoom="6"
        class="w-full h-full"
        @loaded="onMapLoaded"
      >
        <Marker
          v-for="resource in resources"
          :key="resource.id"
          :position="{
            lat: typeof resource.latitude === 'string' ? parseFloat(resource.latitude) : resource.latitude,
            lng: typeof resource.longitude === 'string' ? parseFloat(resource.longitude) : resource.longitude
          }"
          @click="() => handleMarkerClick(resource)"
        />

        <InfoWindow
          v-if="selectedResource"
          :position="{
            lat: typeof selectedResource.latitude === 'string' ? parseFloat(selectedResource.latitude) : selectedResource.latitude,
            lng: typeof selectedResource.longitude === 'string' ? parseFloat(selectedResource.longitude) : selectedResource.longitude
          }"
          @closeclick="handleInfoWindowClose"
        >
          <div class="p-4">
            <h3 class="font-semibold mb-2">{{ selectedResource.name }}</h3>
            <p>{{ selectedResource.address }}</p>
            <button 
              @click="showDetailsModal = true"
              class="mt-2 bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
            >
              View Details
            </button>
          </div>
        </InfoWindow>
      </Map>
    </div>

    <!-- Hospital Details Modal -->
    <HospitalDetailsModal
      v-if="showDetailsModal"
      :hospital="selectedResource"
      @close="handleModalClose"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import apiClient from "@/plugins/axios";
import HospitalSidebar from "./HospitalSidebar.vue";
import HospitalDetailsModal from "./HospitalDetailsModal.vue";
import { Map, Marker, InfoWindow } from '@fawmi/vue-google-maps';

export default defineComponent({
  name: 'MapView',
  
  components: {
    Map,
    Marker,
    InfoWindow,
    HospitalSidebar,
    HospitalDetailsModal,
  },

  data() {
    return {
      mapCenter: { lat: 36.7783, lng: -119.4179 },
      resources: [],
      selectedResource: null,
      showDetailsModal: false
    };
  },

  async mounted() {
    // Fetch resources when component mounts
    await this.fetchResources();
  },

  methods: {
    onMapLoaded(map) {
      console.log('Map loaded:', map);
      // Log current resources when map loads to verify data
      console.log('Current resources:', this.resources);
    },

    async fetchResources() {
      try {
        const { data } = await apiClient.get("/api/regional-centers");
        console.log("Raw resources data:", data);
        
        this.resources = (Array.isArray(data) ? data : []).map(resource => ({
          ...resource,
          latitude: typeof resource.latitude === 'string' ? parseFloat(resource.latitude) : resource.latitude,
          longitude: typeof resource.longitude === 'string' ? parseFloat(resource.longitude) : resource.longitude
        }));

        if (this.resources.length > 0) {
          this.mapCenter = {
            lat: this.resources[0].latitude,
            lng: this.resources[0].longitude
          };
        }
        
        console.log("Processed resources:", this.resources);
      } catch (error) {
        console.error("Failed to fetch resources:", error);
      }
    },

    handleHospitalSelection(hospital) {
      this.selectedResource = hospital;
      this.mapCenter = {
        lat: typeof hospital.latitude === 'string' ? parseFloat(hospital.latitude) : hospital.latitude,
        lng: typeof hospital.longitude === 'string' ? parseFloat(hospital.longitude) : hospital.longitude
      };
    },

    handleMarkerClick(hospital) {
      console.log('Marker clicked:', hospital);
      this.selectedResource = hospital;
    },

    handleInfoWindowClose() {
      this.selectedResource = null;
    },

    handleModalClose() {
      this.showDetailsModal = false;
    }
  }
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