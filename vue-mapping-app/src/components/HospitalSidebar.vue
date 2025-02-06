<template>
  <div class="flex h-screen">
    <!-- Sidebar Toggle Button (Mobile) -->
    <button
      @click="toggleSidebar"
      class="fixed top-4 left-4 z-50 bg-white p-2 rounded-md shadow-lg md:hidden"
    >
      <span v-if="isSidebarOpen">✕</span>
      <span v-else>☰</span>
    </button>

    <!-- Sidebar -->
    <div
      :class="{
        'fixed inset-y-0 left-0 transform bg-white border-r border-gray-200 overflow-y-auto transition-transform duration-300 ease-in-out z-40 w-80': true,
        'translate-x-0': isSidebarOpen,
        '-translate-x-full': !isSidebarOpen,
        'md:relative md:translate-x-0': true,
      }"
    >
      <div class="p-4">
        <h2 class="text-lg font-semibold mb-4">Hospital Resources</h2>

        <!-- Filter Controls -->
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-2">Filter by Type</label>
          <select v-model="selectedType" class="w-full border border-gray-300 rounded-md shadow-sm p-2">
            <option value="">All</option>
            <option v-for="type in uniqueHospitalTypes" :key="type" :value="type">
              {{ type }}
            </option>
          </select>
        </div>

        <!-- Hospital List -->
        <div v-if="filteredHospitals.length" class="space-y-2">
          <div
            v-for="hospital in filteredHospitals"
            :key="hospital.id"
            @click="selectHospital(hospital)"
            class="p-3 rounded-lg cursor-pointer border bg-gray-50 hover:bg-gray-100"
          >
            <div class="flex justify-between items-center">
              <span class="font-medium">{{ hospital.name }}</span>
              <span class="px-2 py-1 rounded-full text-xs bg-blue-100 text-blue-700">
                {{ hospital.type }}
              </span>
            </div>
            <div class="text-sm text-gray-500 mt-1">
              <p>{{ hospital.address }}</p>
              <p>📞 {{ hospital.phone }}</p>
            </div>
          </div>
        </div>

        <!-- No Results Message -->
        <div v-else class="text-gray-500 text-center mt-4">No hospitals found.</div>
      </div>
    </div>

    <!-- Main Content (Map) -->
    <div class="flex-1">
      <Map :center="mapCenter" :zoom="6" class="w-full h-screen" @load="handleMapLoad">
        <Marker
          v-for="resource in resources"
          :key="resource.id"
          :position="{ lat: resource.latitude, lng: resource.longitude }"
          @click="selectedResource = resource"
        />

        <InfoWindow
          v-if="selectedResource"
          :position="{ lat: selectedResource.latitude, lng: selectedResource.longitude }"
          @closeclick="selectedResource = null"
        >
          <div>
            <h3 class="font-semibold">{{ selectedResource.name }}</h3>
            <p>{{ selectedResource.address }}</p>
            <p>📞 {{ selectedResource.phone }}</p>
            <a
              :href="'https://' + selectedResource.website"
              target="_blank"
              class="text-blue-600 underline"
            >
              Visit Website
            </a>
          </div>
        </InfoWindow>
      </Map>
    </div>
  </div>
</template>

<script>
import { Map, Marker, InfoWindow } from "@fawmi/vue-google-maps";
import apiClient from "@/plugins/axios";

export default {
  components: {
    Map,
    Marker,
    InfoWindow,
  },
  data() {
    return {
      isSidebarOpen: true,
      mapCenter: { lat: 36.7783, lng: -119.4179 },
      resources: [],
      selectedResource: null,
      selectedType: "",
    };
  },
  computed: {
    uniqueHospitalTypes() {
      return [...new Set(this.resources.map((h) => h.type))];
    },
    filteredHospitals() {
      return this.selectedType
        ? this.resources.filter((h) => h.type === this.selectedType)
        : this.resources;
    },
  },
  async mounted() {
    await this.fetchResources();
  },
  methods: {
   async fetchResources() {
  try {
    const { data } = await apiClient.get("/api/regional-centers");
    console.log("Fetched resources:", data); // Add this line
    this.resources = Array.isArray(data) ? data : [];
    
    if (this.resources.length > 0) {
      console.log("First resource:", this.resources[0]); // Add this line
      this.mapCenter = {
        lat: this.resources[0].latitude,
        lng: this.resources[0].longitude,
      };
    }
  } catch (error) {
    console.error("Failed to fetch resources:", error);
  }
},

  handleMapLoad(map) {
    console.log("Map loaded:", map);
  },


    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
    selectHospital(hospital) {
      this.selectedResource = hospital;
      this.mapCenter = {
        lat: hospital.latitude,
        lng: hospital.longitude,
      };
    },
  },
};
</script>