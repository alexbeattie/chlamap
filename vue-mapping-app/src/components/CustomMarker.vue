<!-- CustomMarker.vue -->
<template>
  <Marker
    :position="position"
    :icon="getMarkerIcon()"
    :clickable="true"
    @click="$emit('marker-click')"
  />
</template>

<script>
import { Marker } from '@fawmi/vue-google-maps';

export default {
  name: 'CustomMarker',
  
  components: {
    Marker
  },

  props: {
    position: {
      type: Object,
      required: true,
      validator: value => value.lat !== undefined && value.lng !== undefined
    },
    hospital: {
      type: Object,
      required: true
    },
    isSelected: {
      type: Boolean,
      default: false
    }
  },
    emits: ['marker-click'],

  methods: {
    getMarkerIcon() {
      // Hospital cross SVG path
      // This creates a symmetrical cross with slightly rounded edges
      const path = `
        M 15,0 
        A 3,3 0 0 1 18,3 
        L 18,12 
        L 27,12 
        A 3,3 0 0 1 30,15 
        L 30,18 
        A 3,3 0 0 1 27,21 
        L 18,21 
        L 18,30 
        A 3,3 0 0 1 15,33 
        L 12,33 
        A 3,3 0 0 1 9,30 
        L 9,21 
        L 0,21 
        A 3,3 0 0 1 -3,18 
        L -3,15 
        A 3,3 0 0 1 0,12 
        L 9,12 
        L 9,3 
        A 3,3 0 0 1 12,0 
        Z
      `.trim();

      return {
        path: path,
        fillColor: this.isSelected ? "#FF4444" : "#E11D48", // Red shades
        fillOpacity: 1,
        strokeWeight: 2,
        strokeColor: "#FFFFFF", // White border
        scale: 1,
        anchor: new window.google.maps.Point(15, 16.5) // Center point of the cross
      };
    }
  }
};
</script>