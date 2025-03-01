// src/plugins/axios.js
import axios from 'axios';
import { useAuthStore } from '../store/auth';

// Create axios instance with base URL
const api = axios.create({
  baseURL: 'http://localhost:3000/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor
api.interceptors.request.use(
  config => {
    // Get the auth store at request time, not during module initialization
    const token = localStorage.getItem('token');

    // Add auth token to requests if available
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// Response interceptor
api.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    // Handle authentication errors
    if (error.response && error.response.status === 401) {
      // Get the auth store at error time
      const { logout } = useAuthStore();
      // If token is expired or invalid, logout
      if (logout) logout();
    }

    return Promise.reject(error);
  }
);

export default api;