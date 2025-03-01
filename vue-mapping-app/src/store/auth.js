// src/store/auth.js - Using Pinia for state management

import { defineStore } from 'pinia';
import api from '../plugins/axios';
import router from '../router';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token'),
    expiresAt: localStorage.getItem('expiresAt'),
    isAuthenticated: !!localStorage.getItem('token'),
    loading: false,
    error: null
  }),

  getters: {
    isAdmin: (state) => state.user?.roles?.includes('admin'),
    userPermissions: (state) => state.user?.permissions || [],
    hasPermission: (state) => (permission) => {
      return state.user?.permissions?.includes(permission);
    }
  },

  actions: {
    async register(credentials) {
      this.loading = true;
      this.error = null;
      try {
        await api.post('/register', credentials);
        return true;
      } catch (error) {
        this.error = error.response?.data?.error || 'Registration failed';
        return false;
      } finally {
        this.loading = false;
      }
    },

    async login(credentials) {
      this.loading = true;
      this.error = null;

      try {
        const response = await api.post('/login', credentials);
        const { token, expires_at } = response.data;

        // Save to store
        this.token = token;
        this.expiresAt = expires_at;
        this.isAuthenticated = true;

        // Save to localStorage
        localStorage.setItem('token', token);
        localStorage.setItem('expiresAt', expires_at);

        // Get user profile
        await this.fetchUser();

        return true;
      } catch (error) {
        this.error = error.response?.data?.error || 'Login failed';
        return false;
      } finally {
        this.loading = false;
      }
    },

    async fetchUser() {
      this.loading = true;
      this.error = null;

      try {
        const response = await api.get('/me');
        this.user = response.data.user;
        this.user.roles = response.data.roles;

        return this.user;
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to fetch user profile';
        this.logout();
        return null;
      } finally {
        this.loading = false;
      }
    },

    logout() {
      // Clear state
      this.user = null;
      this.token = null;
      this.expiresAt = null;
      this.isAuthenticated = false;

      // Clear localStorage
      localStorage.removeItem('token');
      localStorage.removeItem('expiresAt');

      // Redirect to login
      router.push('/login');
    },

    // Check if the token is expired
    isTokenExpired() {
      if (!this.expiresAt) return true;

      const expiryTime = new Date(this.expiresAt * 1000);
      const currentTime = new Date();

      return currentTime >= expiryTime;
    },

    // Initialize auth state (call this in app initialization)
    async initialize() {
      if (this.token && !this.isTokenExpired()) {
        try {
          await this.fetchUser();
        } catch (error) {
          this.logout();
        }
      } else if (this.token) {
        this.logout();
      }
    }
  }
});