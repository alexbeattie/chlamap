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
      const { AdvancedMarkerElement } = await google.maps.importLibrary("marker");
      
      const markerElement = document.createElement('div');
      markerElement.className = 'custom-marker';
      markerElement.innerHTML = `
        <div class="marker-content">
          <svg 
            width="36" 
            height="36" 
            viewBox="0 0 36 36" 
            fill="none" 
            xmlns="http://www.w3.org/2000/svg"
            style="
              filter: drop-shadow(0 2px 4px rgba(0,0,0,0.3));
              transition: transform 0.2s;
            "
          >
            <!-- Main cross background for white border effect -->
            <path d="
              M 15,3 
              A 3,3 0 0 1 18,0 
              L 21,0 
              A 3,3 0 0 1 24,3 
              L 24,12 
              L 33,12 
              A 3,3 0 0 1 36,15 
              L 36,18 
              A 3,3 0 0 1 33,21 
              L 24,21 
              L 24,30 
              A 3,3 0 0 1 21,33 
              L 18,33 
              A 3,3 0 0 1 15,30 
              L 15,21 
              L 6,21 
              A 3,3 0 0 1 3,18 
              L 3,15 
              A 3,3 0 0 1 6,12 
              L 15,12 
              Z
            " 
            fill="white"
            />
            <!-- Main cross foreground -->
            <path d="
              M 15,3 
              A 3,3 0 0 1 18,0 
              L 21,0 
              A 3,3 0 0 1 24,3 
              L 24,12 
              L 33,12 
              A 3,3 0 0 1 36,15 
              L 36,18 
              A 3,3 0 0 1 33,21 
              L 24,21 
              L 24,30 
              A 3,3 0 0 1 21,33 
              L 18,33 
              A 3,3 0 0 1 15,30 
              L 15,21 
              L 6,21 
              A 3,3 0 0 1 3,18 
              L 3,15 
              A 3,3 0 0 1 6,12 
              L 15,12 
              Z
            " 
            fill="${this.isSelected ? '#E11D48' : '#2196F3'}"
            transform="scale(0.9) translate(2, 2)"
            />
          </svg>
        </div>
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
        const path = this.marker.content.querySelector('path:last-child');
        if (path) {
          path.setAttribute('fill', newValue ? '#E11D48' : '#2196F3');
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
.custom-marker:hover svg {
  transform: scale(1.1);
}
</style>