<template>
  <div v-if="isVisible" class="custom-info-window">
    <div class="info-window-content">
      <h3>{{ hospital.name }}</h3>
      <p>{{ hospital.address }}</p>
      <p v-if="hospital.phone">📞 {{ hospital.phone }}</p>
      <button @click="handleClose">Close</button>
    </div>
  </div>
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
  },

  data() {
    return {
      isVisible: true
    };
  },

  methods: {
    handleClose() {
      this.isVisible = false;
      this.$emit('closeclick');
    }
  }
};
</script>

<style scoped>
.custom-info-window {
  position: absolute;
  background: white;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.info-window-content {
  text-align: left;
}

button {
  margin-top: 8px;
  padding: 4px 8px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background: #0056b3;
}
</style>