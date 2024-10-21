import { createRouter, createWebHistory } from 'vue-router';
import UserRegister from '../views/Register.vue';
import UserLogin from '../views/Login.vue';

const routes = [
  {
    path: '/register',
    name: 'Register',
    component: UserRegister
  },
  {
    path: '/login',
    name: 'Login',
    component: UserLogin
  }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;