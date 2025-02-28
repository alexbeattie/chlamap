<template>
  <div id="app">
    <NavigationBar />
    <router-view></router-view>  <!-- Use closing tag to be safe -->
  </div>
</template>

<script>
import NavigationBar from './components/NavigationBar.vue';

export default {
   components: {
    NavigationBar,
  },
  // name: 'App'
}
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