<template>
  <div id="app">
    <MapView />
  </div>
</template>

<script>
import MapView from './components/MapView.vue';

export default {
  name: 'App',
  components: {
    MapView, // Register the Map component
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 0;
}
</style>
