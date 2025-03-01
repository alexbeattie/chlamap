<!-- src/App.vue -->
<template>
  <div class="app-container">
     <header>
      <nav>
        <div>
          <router-link to="/">Home</router-link>
          <router-link to="/providers">Providers Map</router-link> <!-- ✅ Add this -->
        </div>
        
        <div v-if="!authStore.isAuthenticated">
          <router-link to="/login">Login</router-link>
          <router-link to="/register">Register</router-link>
        </div>
        
        <div v-else>
          <router-link to="/profile">Profile</router-link>
          <router-link to="/admin" v-if="authStore.hasPermission('read:users')">Admin</router-link>
          <button @click="authStore.logout" class="logout-btn">Logout</button>
        </div>
      </nav>
    </header>
    
    <main>
      <router-view />
    </main>
    
    <footer>
      <p>&copy; {{ new Date().getFullYear() }} Your App Name</p>
    </footer>
  </div>
</template>

<script setup>
import { useAuthStore } from './store/auth';

const authStore = useAuthStore();
</script>

<style>
body {
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
    Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  background-color: #f5f5f5;
}

.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

header {
  background-color: #333;
  color: white;
  padding: 1rem;
}

nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

nav a {
  color: white;
  margin-right: 1rem;
  text-decoration: none;
}

nav a:hover {
  text-decoration: underline;
}

main {
  flex: 1;
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}

footer {
  background-color: #333;
  color: white;
  padding: 1rem;
  text-align: center;
}

.logout-btn {
  background: none;
  border: none;
  color: white;
  cursor: pointer;
  font-size: 1rem;
}

.logout-btn:hover {
  text-decoration: underline;
}
</style>