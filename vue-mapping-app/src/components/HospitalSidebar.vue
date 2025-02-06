<!-- HospitalSidebar.vue -->
<template>
  <div class="h-full flex flex-col">
    <div class="p-4">
      <h2 class="text-lg font-semibold mb-4">Hospital Resources</h2>
      
      <!-- Search/Filter Section -->
      <div class="mb-4">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Search hospitals..."
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>
    </div>

    <!-- Hospital List -->
    <div class="flex-1 overflow-y-auto">
      <div class="space-y-2 p-4">
        <div
          v-for="hospital in filteredHospitals"
          :key="hospital.id"
          :data-hospital-id="hospital.id"
          @click="$emit('hospital-selected', hospital)"
          class="p-3 rounded-lg cursor-pointer transition-colors duration-200"
          :class="{
            'bg-blue-50 border-blue-500': selectedHospital?.id === hospital.id,
            'bg-gray-50 hover:bg-gray-100 border-transparent': selectedHospital?.id !== hospital.id,
            'border': true
          }"
        >
          <div class="flex justify-between items-start">
            <div>
              <h3 class="font-medium text-gray-900">{{ hospital.name }}</h3>
              <p class="text-sm text-gray-500 mt-1">{{ hospital.address }}</p>
              <p class="text-sm text-gray-500" v-if="hospital.phone">
                📞 {{ hospital.phone }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HospitalSidebar',
  
  props: {
    hospitals: {
      type: Array,
      required: true
    },
    selectedHospital: {
      type: Object,
      default: null
    }
  },

  data() {
    return {
      searchQuery: ''
    };
  },

  computed: {
    filteredHospitals() {
      if (!this.searchQuery) return this.hospitals;
      
      const query = this.searchQuery.toLowerCase();
      return this.hospitals.filter(hospital => 
        hospital.name.toLowerCase().includes(query) ||
        hospital.address.toLowerCase().includes(query) ||
        (hospital.phone && hospital.phone.toLowerCase().includes(query))
      );
    }
  },

  watch: {
    selectedHospital(newHospital) {
      if (newHospital) {
        this.$nextTick(() => {
          const element = document.querySelector(`[data-hospital-id="${newHospital.id}"]`);
          if (element) {
            element.scrollIntoView({ behavior: 'smooth', block: 'center' });
          }
        });
      }
    }
  }
};
</script>