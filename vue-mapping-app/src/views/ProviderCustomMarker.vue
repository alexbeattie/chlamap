<template>
  <div>
    <Marker
      :options="markerOptions"
      :position="position"
      @click="$emit('marker-click')"
    >
      <template v-if="false">
        <!-- We don't use the slot content, but we need it for proper event handling -->
      </template>
    </Marker>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Marker } from "@fawmi/vue-google-maps";

export default defineComponent({
  name: "CustomMarker",
  components: {
    Marker,
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
    isSelected: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    markerOptions() {
      return {
        animation: this.isSelected ? window.google?.maps?.Animation?.BOUNCE : null,
        icon: this.getMarkerIcon(),
        zIndex: this.isSelected ? 1000 : 1,
        clickable: true,
        draggable: false,
        title: this.provider.name,
      };
    },
  },
  methods: {
    getMarkerIcon() {
      // Determine marker color - could be based on provider type, etc.
      const color = this.isSelected ? "#FF0000" : "#1E40AF"; // Blue as default, red when selected
      
      // You can customize this SVG to represent different types of providers
      return {
        url: "data:image/svg+xml;charset=UTF-8," + encodeURIComponent(`
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="${color}" stroke="#FFFFFF" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
            <circle cx="12" cy="10" r="3"></circle>
          </svg>
        `),
        scaledSize: new window.google?.maps?.Size(32, 32),
        anchor: new window.google?.maps?.Point(16, 32),
      };
    },
  },
});
</script>