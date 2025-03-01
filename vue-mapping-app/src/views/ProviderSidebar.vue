<template>
  <div class="h-full flex flex-col bg-white">
    <div class="px-4 py-3 border-b border-gray-200">
      <h1 class="text-xl font-semibold text-gray-800">ABA Providers</h1>
      <div class="mt-2">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Search providers..."
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>
      <div class="mt-2">
        <select 
          v-model="selectedArea" 
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="">All Areas</option>
          <option v-for="area in uniqueAreas" :key="area" :value="area">
            {{ area }}
          </option>
        </select>
      </div>
    </div>

    <div class="flex-1 overflow-y-auto">
      <div v-if="filteredProviders.length === 0" class="p-4 text-center text-gray-500">
        No providers found matching your search.
      </div>
      <div v-else>
        <div
          v-for="provider in filteredProviders"
          :key="provider.id"
          :data-provider-id="provider.id"
          class="p-4 border-b border-gray-200 hover:bg-gray-50 cursor-pointer transition-colors"
          :class="{
            'bg-blue-50 border-l-4 border-l-blue-500': selectedProvider?.id === provider.id,
          }"
          @click="$emit('provider-selected', provider)"
        >
          <h2 class="text-lg font-medium text-gray-900">{{ provider.name }}</h2>
          <p class="mt-1 text-sm text-gray-600">
            <span class="font-semibold">Phone:</span> {{ provider.phone }}
          </p>
          <p class="mt-1 text-sm text-gray-600 line-clamp-2">
            <span class="font-semibold">Areas:</span>
            {{ formatAreas(provider.areas) }}
          </p>
          <div v-if="provider.center_based_services" class="mt-1 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
            Center-Based
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProviderSidebar",
  
  props: {
    providers: {
      type: Array,
      required: true,
    },
    selectedProvider: {
      type: Object,
      default: null,
    },
  },
  
  data() {
    return {
      searchQuery: "",
      selectedArea: "",
    };
  },
  
  computed: {
    uniqueAreas() {
      // Extract all areas from all providers and create a unique list
      const allAreas = new Set();
      
      this.providers.forEach(provider => {
        if (provider.areas && Array.isArray(provider.areas)) {
          provider.areas.forEach(area => {
            if (area) allAreas.add(area);
          });
        }
      });
      
      return [...allAreas].sort();
    },
    
    filteredProviders() {
      return this.providers.filter(provider => {
        // Filter by search query (name or phone)
        const matchesSearch = 
          !this.searchQuery ||
          provider.name.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
          (provider.phone && provider.phone.includes(this.searchQuery));
        
        // Filter by selected area
        const matchesArea = 
          !this.selectedArea ||
          (provider.areas && 
           Array.isArray(provider.areas) && 
           provider.areas.some(area => area && area.includes(this.selectedArea)));
        
        return matchesSearch && matchesArea;
      });
    },
  },
  
  methods: {
    formatAreas(areas) {
      if (!areas || !Array.isArray(areas)) return "No areas listed";
      return areas.slice(0, 3).join(", ") + (areas.length > 3 ? "..." : "");
    },
  },
};
</script>