import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/cadastro',
      name: 'signup',
      component: () => import('../views/SignupView.vue')
    },
    {
      path: '/login',
      name: 'signin',
      component: () => import('../views/SigninView.vue')
    },
    {
      path: '/primeiros-passos',
      name: 'first-steps',
      component: () => import('../views/FirstStepsView.vue')
    },
  ]
})

export default router
