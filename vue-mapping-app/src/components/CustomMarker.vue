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
     

      return {
        url: "/marker.svg", // Make sure the SVG is in the 'public' folder
    scaledSize: new window.google.maps.Size(40, 40), // Adjust size
        anchor: new window.google.maps.Point(20, 40)
    
      };
    }
  }
};
</script>