<!-- src/views/Register.vue -->
<template>
  <div class="register-container">
    <h1>Create Account</h1>
    
    <form @submit.prevent="handleRegister" class="register-form">
      <div class="form-group">
        <label for="firstName">First Name</label>
        <input 
          id="firstName"
          v-model="firstName"
          type="text" 
          required
          placeholder="Enter your first name"
        />
      </div>
      
      <div class="form-group">
        <label for="lastName">Last Name</label>
        <input 
          id="lastName"
          v-model="lastName"
          type="text" 
          required
          placeholder="Enter your last name"
        />
      </div>
      
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
      
      <div class="form-group">
        <label for="confirmPassword">Confirm Password</label>
        <input 
          id="confirmPassword"
          v-model="confirmPassword"
          type="password" 
          required
          placeholder="Confirm your password"
        />
        <div v-if="passwordsDoNotMatch" class="error-message">
          Passwords do not match
        </div>
      </div>
      
      <div v-if="authStore.error" class="error-message">
        {{ authStore.error }}
      </div>
      
      <button type="submit" :disabled="authStore.loading || !canSubmit">
        {{ authStore.loading ? 'Creating Account...' : 'Register' }}
      </button>
      
      <div class="login-link">
        Already have an account? <router-link to="/login">Login</router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';

const router = useRouter();
const authStore = useAuthStore();

const firstName = ref('');
const lastName = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');

const passwordsDoNotMatch = computed(() => {
  return password.value && 
         confirmPassword.value && 
         password.value !== confirmPassword.value;
});

const canSubmit = computed(() => {
  return firstName.value && 
         lastName.value && 
         email.value && 
         password.value && 
         confirmPassword.value && 
         !passwordsDoNotMatch.value;
});

async function handleRegister() {
  if (!canSubmit.value) return;
  
  const success = await authStore.register({
    first_name: firstName.value,
    last_name: lastName.value,
    email: email.value,
    password: password.value
  });
  
  if (success) {
    // Either redirect to login or automatically log them in
    router.push('/login');
  }
}
</script>

<style scoped>
.register-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 40px 20px;
}

.register-form {
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

.login-link {
  text-align: center;
  margin-top: 10px;
}
</style>