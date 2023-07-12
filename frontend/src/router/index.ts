import { createRouter, createWebHistory } from 'vue-router'
// import FeedView from '../views/FeedView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/FeedView.vue')
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
    {
      path: '/inbox',
      name: 'inbox',
      component: () => null
    },
    {
      path: '/notificacoes',
      name: 'notifications',
      component: () => null
    },
    {
      path: '/configuracoes',
      name: 'configurations',
      component: () => null
    },
    {
      path: '/explorar',
      name: 'search',
      component: () => null
    },
  ]
})

export default router
