<template>
  <InfoWindow
    :options="options"
    :position="position"
    @closeclick="$emit('closeclick')"
  >
    <div class="info-window-content">
      <h3 class="text-lg font-semibold mb-1">{{ provider.name }}</h3>
      <p class="text-sm mb-1" v-if="provider.phone">
        <span class="font-medium">Phone:</span> {{ provider.phone }}
      </p>
      <p class="text-sm mb-2" v-if="provider.address">
        {{ provider.address }}
      </p>
      <div v-if="hasAreas" class="text-sm mb-2">
        <span class="font-medium">Areas:</span>
        <div class="flex flex-wrap gap-1 mt-1">
          <span
            v-for="(area, i) in displayedAreas"
            :key="i"
            class="inline-block bg-blue-100 text-blue-800 text-xs px-2 py-0.5 rounded"
          >
            {{ area }}
          </span>
          <span v-if="showMoreAreas" class="text-blue-600 text-xs cursor-pointer">
            +{{ provider.areas.length - 3 }} more
          </span>
        </div>
      </div>
      <div v-if="provider.center_based_services" class="mt-1 mb-2">
        <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
          Center-Based: {{ provider.center_based_services }}
        </span>
      </div>
      <div class="flex mt-2">
        <button
          class="bg-blue-500 hover:bg-blue-600 text-white py-1 px-3 rounded text-sm w-full"
          @click="openDirections"
        >
          Get Directions
        </button>
      </div>
    </div>
  </InfoWindow>
</template>

<script>
import { defineComponent } from "vue";
import { InfoWindow } from "@fawmi/vue-google-maps";

export default defineComponent({
  name: "CustomInfoWindow",
  components: {
    InfoWindow,
  },
  props: {
    position: {
      type: Object,
      required: true,
    },
    provider: {
      type: Object,
      required: true,
    },
  },
  computed: {
    options() {
      return {
        maxWidth: 320,
        pixelOffset: {
          width: 0,
          height: -40,
        },
      };
    },
    hasAreas() {
      return this.provider.areas && this.provider.areas.length > 0;
    },
    displayedAreas() {
      return this.provider.areas ? this.provider.areas.slice(0, 3) : [];
    },
    showMoreAreas() {
      return this.provider.areas && this.provider.areas.length > 3;
    },
  },
  methods: {
    openDirections() {
      if (this.provider.latitude && this.provider.longitude) {
        window.open(
          `https://www.google.com/maps/dir/?api=1&destination=${this.provider.latitude},${this.provider.longitude}`,
          "_blank"
        );
      }
    },
  },
});
</script>

<style scoped>
.info-window-content {
  min-width: 200px;
  max-width: 300px;
}
</style>