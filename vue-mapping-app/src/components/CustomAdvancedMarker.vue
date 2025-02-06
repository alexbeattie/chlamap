<!-- CustomAdvancedMarker.vue -->

<template>
  <div ref="markerContainer"></div>
</template>

<script>
export default {
  name: 'CustomAdvancedMarker',
  
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
    },
    map: {
      type: Object,
      required: true
    }
  },

  data() {
    return {
      marker: null
    };
  },

  mounted() {
    this.initializeMarker();
  },

  methods: {
    async initializeMarker() {
      // eslint-disable-next-line
      const { AdvancedMarkerElement } = await google.maps.importLibrary("marker");
      
      const markerElement = document.createElement('div');
      markerElement.className = 'custom-marker';
      markerElement.innerHTML = `
        <div class="marker-content" style="
          background-color: ${this.isSelected ? '#4CAF50' : '#2196F3'};
          width: 24px;
          height: 24px;
          border-radius: 50%;
          border: 2px solid white;
          box-shadow: 0 2px 4px rgba(0,0,0,0.3);
          cursor: pointer;
          transition: transform 0.2s;
        "></div>
      `;

      this.marker = new AdvancedMarkerElement({
        position: this.position,
        map: this.map,
        content: markerElement,
        title: this.hospital.name
      });

      this.marker.addListener('click', () => {
        this.$emit('click', this.hospital);
      });
    }
  },

  watch: {
    isSelected(newValue) {
      if (this.marker) {
        const markerContent = this.marker.content.querySelector('.marker-content');
        if (markerContent) {
          markerContent.style.backgroundColor = newValue ? '#4CAF50' : '#2196F3';
        }
      }
    },
    position(newPos) {
      if (this.marker) {
        this.marker.position = newPos;
      }
    }
  },

  beforeUnmount() {
    if (this.marker) {
      this.marker.map = null;
    }
  }
};
</script>

<style>
.custom-marker {
  cursor: pointer;
}
.custom-marker:hover .marker-content {
  transform: scale(1.1);
}
</style>