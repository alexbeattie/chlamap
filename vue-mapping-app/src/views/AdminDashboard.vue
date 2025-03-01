<!-- src/views/AdminDashboard.vue -->
<template>
  <div class="admin-container">
    <h1>Admin Dashboard</h1>
    
    <div v-if="loading" class="loading">Loading user data...</div>
    
    <div v-else-if="error" class="error-message">{{ error }}</div>
    
    <div v-else>
      <div class="users-header">
        <h2>User Management</h2>
        <span class="user-count">{{ users.length }} users</span>
      </div>
      
      <div class="user-table-container">
        <table class="user-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Email</th>
              <th>Created</th>
              <th v-if="authStore.hasPermission('update:users') || authStore.hasPermission('delete:users')">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.id">
              <td>{{ user.id }}</td>
              <td>{{ user.first_name }} {{ user.last_name }}</td>
              <td>{{ user.email }}</td>
              <td>{{ formatDate(user.created_at) }}</td>
              <td v-if="authStore.hasPermission('update:users') || authStore.hasPermission('delete:users')" class="actions">
                <button 
                  v-if="authStore.hasPermission('update:users')" 
                  class="edit-btn"
                  @click="editUser(user)"
                >
                  Edit
                </button>
                <button 
                  v-if="authStore.hasPermission('delete:users')" 
                  class="delete-btn"
                  @click="confirmDelete(user)"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useAuthStore } from '../store/auth';
import api from '../plugins/axios';

const authStore = useAuthStore();
const users = ref([]);
const loading = ref(true);
const error = ref(null);

onMounted(async () => {
  try {
    const response = await api.get('/users');
    users.value = response.data.users;
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to load users';
  } finally {
    loading.value = false;
  }
});

function formatDate(dateString) {
  if (!dateString) return 'N/A';
  
  const date = new Date(dateString);
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric', 
    month: 'short', 
    day: 'numeric'
  }).format(date);
}

function editUser(user) {
  // This would typically open a modal or navigate to an edit page
  alert(`Edit user: ${user.first_name} ${user.last_name}`);
}

function confirmDelete(user) {
  if (confirm(`Are you sure you want to delete ${user.first_name} ${user.last_name}?`)) {
    deleteUser(user.id);
  }
}

async function deleteUser(userId) {
  try {
    await api.delete(`/users/${userId}`);
    // Remove from the list
    users.value = users.value.filter(user => user.id !== userId);
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to delete user';
  }
}
</script>

<style scoped>
.admin-container {
  max-width: 1000px;
  margin: 0 auto;
}

.users-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.user-count {
  background-color: #e0f7fa;
  padding: 5px 10px;
  border-radius: 12px;
  font-size: 0.9em;
}

.user-table-container {
  overflow-x: auto;
}

.user-table {
  width: 100%;
  border-collapse: collapse;
  background-color: white;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}

.user-table th,
.user-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.user-table th {
  background-color: #f5f5f5;
  font-weight: 600;
}

.user-table tr:last-child td {
  border-bottom: none;
}

.actions {
  display: flex;
  gap: 8px;
}

.edit-btn,
.delete-btn {
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9em;
}

.edit-btn {
  background-color: #e3f2fd;
  color: #1976d2;
}

.delete-btn {
  background-color: #ffebee;
  color: #d32f2f;
}

.loading {
  text-align: center;
  padding: 20px;
  color: #666;
}

.error-message {
  color: #f44336;
  padding: 15px;
  background-color: rgba(244, 67, 54, 0.1);
  border-radius: 4px;
  margin-top: 15px;
}
</style>