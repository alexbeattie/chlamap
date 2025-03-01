import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../store/auth';
import Login from '../views/UserLogin.vue';
import Register from '../views/UserRegister.vue';
import Profile from '../views/UserProfile.vue';
import AdminDashboard from '../views/AdminDashboard.vue';
import NotFound from '../views/NotFound.vue';
import ProvidersMapView from '../views/ProvidersMapView.vue'; // Import the new component

import HomeView from '../views/HomeView.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
  },
  {
    path: '/aba-centers',
    name: 'ABACenters',
    component: () => import('../views/ABACentersView.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { guest: true }
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: { guest: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin',
    name: 'AdminDashboard',
    component: AdminDashboard,
    meta: {
      requiresAuth: true,
      requiresPermission: 'read:users'
    }
  },
  {
    path: '/providers',
    name: 'providers',
    component: ProvidersMapView
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound
  }
  // other routes...
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});
// Navigation guards
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();

  // Check if the route requires authentication
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // If not authenticated, redirect to login
    if (!authStore.isAuthenticated) {
      next({ name: 'Login', query: { redirect: to.fullPath } });
      return;
    }

    // If token is expired, refresh or redirect to login
    if (authStore.isTokenExpired()) {
      authStore.logout();
      next({ name: 'Login', query: { redirect: to.fullPath } });
      return;
    }

    // If route requires specific permission
    if (to.meta.requiresPermission && !authStore.hasPermission(to.meta.requiresPermission)) {
      next({ name: 'Home' });
      return;
    }
  }

  // If route is for guests only and user is authenticated
  if (to.matched.some(record => record.meta.guest) && authStore.isAuthenticated) {
    next({ name: 'Home' });
    return;
  }

  next();
});

export default router;