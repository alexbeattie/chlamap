<!-- CustomInfoWindow.vue -->
<template>
  <GMapInfoWindow
    :position="position"
    @closeclick="$emit('closeclick')"
  >
    <div class="bg-white rounded-lg shadow-lg max-w-md">
      <div class="px-4 py-2 border-b border-gray-200">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-semibold text-gray-900">{{ hospital.name }}</h3>
        </div>
      </div>

      <div class="px-4 py-2 space-y-2">
        <div class="flex text-sm">
          <div class="font-medium text-gray-700 min-w-24">Address:</div>
          <div class="text-gray-900">{{ hospital.address }}</div>
        </div>

        <div class="flex text-sm" v-if="hospital.phone">
          <div class="font-medium text-gray-700 min-w-24">Phone:</div>
          <div class="text-gray-900">{{ hospital.phone }}</div>
        </div>

        <div class="flex text-sm" v-if="hospital.website">
          <div class="font-medium text-gray-700 min-w-24">Website:</div>
          <div class="text-gray-900">
            <a :href="'https://' + hospital.website" target="_blank" class="text-blue-600 hover:underline">
              {{ hospital.website }}
            </a>
          </div>
        </div>
      </div>

      <div class="px-4 py-2 mt-2 mb-2 flex space-x-2">
        <button
          @click="$emit('show-details')"
          class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors flex-1"
        >
          View Details
        </button>
      </div>
    </div>
  </GMapInfoWindow>
</template>

<script>
export default {
  name: 'CustomInfoWindow',
  
  props: {
    position: {
      type: Object,
      required: true,
      validator: value => value.lat !== undefined && value.lng !== undefined
    },
    hospital: {
      type: Object,
      required: true
    }
  }
};
</script>