<template>
  <nav class="navbar">
    <div class="navbar-brand">
      <router-link to="/" class="navbar-item">Resources</router-link>
    </div>
    <div class="navbar-menu">
      <div class="navbar-start">
        <router-link to="/">Home</router-link>
        <router-link to="/aba-centers">ABA Centers</router-link>
        <router-link to="/providers">Providers Map</router-link>
        <!-- ✅ Add this link -->
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavigationBar", // Corrected component name
};
</script>

<style scoped>
.navbar {
  background-color: #f0f0f0;
  padding: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navbar-brand .navbar-item {
  font-weight: bold;
  font-size: 1.2rem;
}

.navbar-menu {
  display: flex;
}

.navbar-start .navbar-item {
  margin-left: 1rem;
  text-decoration: none;
  color: #333;
}

.navbar-start .navbar-item.router-link-active {
  font-weight: bold;
  color: blue;
}
</style>
