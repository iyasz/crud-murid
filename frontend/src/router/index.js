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
      path: '/murid',
      name: 'muridIndex',
      component: () => import('../views/siswa/indexView.vue')
    },
    {
      path: '/class',
      name: 'classIndex',
      component: () => import('../views/class/index.vue')
    },
    {
      path: '/class/create',
      name: 'classCreate',
      component: () => import('../views/class/create.vue')
    },
    {
      path: '/class/:id',
      name: 'classEdit',
      component: () => import('../views/class/edit.vue')
    },
    
  ]
})

export default router
