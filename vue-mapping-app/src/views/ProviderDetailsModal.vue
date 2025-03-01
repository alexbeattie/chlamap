<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-hidden">
      <!-- Header -->
      <div class="flex justify-between items-center border-b p-4">
        <h2 class="text-xl font-semibold">{{ provider.name }}</h2>
        <button
          class="text-gray-500 hover:text-gray-700"
          @click="$emit('close')"
        >
          <span class="text-2xl">&times;</span>
        </button>
      </div>

      <!-- Content -->
      <div class="p-4 overflow-y-auto max-h-[calc(90vh-120px)]">
        <div class="flex flex-col space-y-4">
          <!-- Contact Information -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">Contact Information</h3>
            <div class="flex items-start mb-2">
              <div class="flex-shrink-0 w-6 text-gray-500">
                <i class="fas fa-phone-alt"></i>
              </div>
              <div class="ml-2">
                <p class="text-gray-800">{{ provider.phone }}</p>
              </div>
            </div>
            <div v-if="provider.address" class="flex items-start">
              <div class="flex-shrink-0 w-6 text-gray-500">
                <i class="fas fa-map-marker-alt"></i>
              </div>
              <div class="ml-2">
                <p class="text-gray-800">{{ provider.address }}</p>
              </div>
            </div>
          </div>

          <!-- Service Areas -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">Service Areas</h3>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="(area, index) in provider.areas"
                :key="index"
                class="bg-blue-100 text-blue-800 px-2 py-1 rounded text-sm inline-block"
              >
                {{ area }}
              </span>
            </div>
          </div>

          <!-- Center Based Services -->
          <div v-if="provider.center_based_services">
            <h3 class="text-lg font-medium text-gray-900 mb-2">Center-Based Locations</h3>
            <p class="text-gray-800">{{ provider.center_based_services }}</p>
          </div>

          <!-- Description/Notes (if available) -->
          <div v-if="provider.description">
            <h3 class="text-lg font-medium text-gray-900 mb-2">Description</h3>
            <p class="text-gray-800">{{ provider.description }}</p>
          </div>

          <!-- Additional Information (can be added as needed) -->
          <div v-if="provider.website">
            <h3 class="text-lg font-medium text-gray-900 mb-2">Website</h3>
            <a 
              :href="formatWebsiteUrl(provider.website)" 
              target="_blank" 
              class="text-blue-600 hover:underline"
            >
              {{ provider.website }}
            </a>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="border-t p-4 flex justify-end space-x-2">
        <button
          class="px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300"
          @click="$emit('close')"
        >
          Close
        </button>
        <button
          class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
          @click="$emit('get-directions')"
        >
          Get Directions
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProviderDetailsModal",
  
  props: {
    provider: {
      type: Object,
      required: true,
    },
  },
  
  methods: {
    formatWebsiteUrl(url) {
      if (!url) return '#';
      return url.startsWith('http') ? url : `https://${url}`;
    }
  },
};
</script>