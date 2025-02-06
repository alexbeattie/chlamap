import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:3000', // Replace with your Go backend's URL
  headers: {
    'Content-Type': 'application/json',
  },
});

export default apiClient;
