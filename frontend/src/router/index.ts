import { createRouter, createWebHistory } from 'vue-router'
// import navigationGuard from './navigationGuard'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/cadastro',
      name: 'signup',
      component: () => import('../views/SignupView.vue'),
      meta: { title: 'Cadastro' }
    },
    {
      path: '/login',
      name: 'signin',
      component: () => import('../views/SigninView.vue'),
      meta: { title: 'Login' }
    },
    {
      path: '/',
      name: 'home',
      component: () => import('../views/FeedView.vue'),
      meta: { requireAuth: true }
    },
    {
      path: '/primeiros-passos',
      name: 'first-steps',
      component: () => import('../views/FirstStepsView.vue'),
      meta: { requireAuth: true }
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
    {
      path: '/perfil/:arroba',
      name: 'profile',
      component: () => null
    },
  ]
})

// router.beforeEach(navigationGuard)
export default router
