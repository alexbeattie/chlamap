<!-- src/views/Login.vue -->
<template>
  <div class="login-container">
    <h1>Login</h1>
    
    <form @submit.prevent="handleLogin" class="login-form">
      <div class="form-group">
        <label for="email">Email</label>
        <input 
          id="email"
          v-model="email"
          type="email" 
          required
          placeholder="Enter your email"
        />
      </div>
      
      <div class="form-group">
        <label for="password">Password</label>
        <input 
          id="password"
          v-model="password"
          type="password" 
          required
          placeholder="Enter your password"
        />
      </div>
      
      <div v-if="authStore.error" class="error-message">
        {{ authStore.error }}
      </div>
      
      <button type="submit" :disabled="authStore.loading">
        {{ authStore.loading ? 'Logging in...' : 'Login' }}
      </button>
      
      <div class="register-link">
        Don't have an account? <router-link to="/register">Register</router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useAuthStore } from '../store/auth';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();

const email = ref('');
const password = ref('');

async function handleLogin() {
  const success = await authStore.login({
    email: email.value,
    password: password.value
  });
  
  if (success) {
    // Redirect to the page the user was trying to access, or home
    const redirectPath = route.query.redirect || '/';
    router.push(redirectPath);
  }
}
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 40px 20px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

input {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button {
  padding: 12px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:disabled {
  background-color: #cccccc;
}

.error-message {
  color: #f44336;
  margin-bottom: 10px;
}

.register-link {
  text-align: center;
  margin-top: 10px;
}
</style>