<!-- src/views/Profile.vue -->
<template>
  <div class="profile-container">
    <h1>Your Profile</h1>
    
    <div v-if="authStore.loading">Loading profile...</div>
    
    <div v-else-if="authStore.user" class="profile-card">
      <div class="profile-header">
        <h2>{{ authStore.user.first_name }} {{ authStore.user.last_name }}</h2>
        <span class="email">{{ authStore.user.email }}</span>
      </div>
      
      <div class="profile-details">
        <div class="detail-item">
          <strong>User ID:</strong> {{ authStore.user.id }}
        </div>
        <div class="detail-item">
          <strong>Member since:</strong> {{ formatDate(authStore.user.created_at) }}
        </div>
        <div class="detail-item">
          <strong>Roles:</strong> 
          <span v-for="(role, index) in authStore.user.roles" :key="role" class="role-badge">
            {{ role }}{{ index < authStore.user.roles.length - 1 ? ', ' : '' }}
          </span>
        </div>
      </div>
    </div>
    
    <div v-else class="error-message">
      {{ authStore.error || 'Unable to load profile data' }}
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useAuthStore } from '../store/auth';

const authStore = useAuthStore();

// Refresh user data when component mounts
onMounted(() => {
  authStore.fetchUser();
});

// Format date to readable string
function formatDate(dateString) {
  if (!dateString) return 'N/A';
  
  const date = new Date(dateString);
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric', 
    month: 'long', 
    day: 'numeric'
  }).format(date);
}
</script>

<style scoped>
.profile-container {
  max-width: 600px;
  margin: 0 auto;
}

.profile-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
  margin-top: 20px;
}

.profile-header {
  border-bottom: 1px solid #eee;
  padding-bottom: 15px;
  margin-bottom: 15px;
}

.profile-header h2 {
  margin: 0 0 5px 0;
}

.email {
  color: #666;
}

.detail-item {
  margin-bottom: 10px;
}

.role-badge {
  display: inline-block;
  padding: 2px 8px;
  margin-right: 5px;
  background-color: #e0f7fa;
  border-radius: 12px;
  font-size: 0.9em;
}

.error-message {
  color: #f44336;
  margin-top: 20px;
  padding: 10px;
  background-color: rgba(244, 67, 54, 0.1);
  border-radius: 4px;
}
</style>