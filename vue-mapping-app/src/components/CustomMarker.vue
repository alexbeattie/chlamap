<!-- CustomMarker.vue -->
<template>
  <GMapMarker
    :position="position"
    :icon="getMarkerIcon()"
    :clickable="true"
    @click="$emit('click', hospital)"
  />
</template>

<script>
export default {
  name: 'CustomMarker',
  
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

  methods: {
    getMarkerIcon() {
      return {
        path: "M 12,0 L 24,24 L 12,16 L 0,24 Z",
        fillColor: this.isSelected ? "#4CAF50" : "#2196F3",
        fillOpacity: 1,
        strokeWeight: 1,
        strokeColor: "#000000",
        scale: 1,
        anchor: new window.google.maps.Point(12, 12)
      };
    }
  }
};
</script>