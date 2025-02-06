<!-- HospitalDetailsModal.vue -->
<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg p-6 max-w-2xl w-full mx-4">
      <!-- Modal Header -->
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-xl font-semibold text-gray-900">{{ hospital.name }}</h3>
        <button
          @click="$emit('close')"
          class="text-gray-500 hover:text-gray-700"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>

      <!-- Hospital Details -->
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <!-- Contact Information -->
          <div class="space-y-2">
            <h4 class="font-medium text-gray-900">Contact Information</h4>
            <p class="text-sm text-gray-600">
              <span class="font-medium">Address:</span><br />
              {{ hospital.address }}
            </p>
            <p class="text-sm text-gray-600">
              <span class="font-medium">Phone:</span><br />
              {{ hospital.phone }}
            </p>
            <p class="text-sm text-gray-600">
              <span class="font-medium">Website:</span><br />
              <a
                :href="'https://' + hospital.website"
                target="_blank"
                class="text-blue-600 hover:underline"
              >
                {{ hospital.website }}
              </a>
            </p>
          </div>

          <!-- Specialties -->
          <div class="space-y-2" v-if="hospital.specialties">
            <h4 class="font-medium text-gray-900">Specialties</h4>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="specialty in hospital.specialties"
                :key="specialty"
                class="px-2 py-1 bg-blue-100 text-blue-800 text-sm rounded-full"
              >
                {{ specialty }}
              </span>
            </div>
          </div>
        </div>

        <!-- Additional Information -->
        <div class="mt-4" v-if="hospital.description">
          <h4 class="font-medium text-gray-900 mb-2">About</h4>
          <p class="text-sm text-gray-600">{{ hospital.description }}</p>
        </div>

        <!-- Operating Hours -->
        <div class="mt-4" v-if="hospital.hours">
          <h4 class="font-medium text-gray-900 mb-2">Operating Hours</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
            <div
              v-for="(hours, day) in hospital.hours"
              :key="day"
              class="text-sm text-gray-600"
            >
              <span class="font-medium">{{ day }}:</span> {{ hours }}
            </div>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="mt-6 flex justify-end space-x-3">
        <button
          @click="$emit('get-directions')"
          class="px-4 py-2 bg-gray-100 text-gray-700 rounded hover:bg-gray-200"
        >
          Get Directions
        </button>
        <button
          @click="$emit('close')"
          class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HospitalDetailsModal',
  
  props: {
    hospital: {
      type: Object,
      required: true
    }
  },

  emits: ['close', 'get-directions']
};
</script>