import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // murid route 
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/murid',
      name: 'murid.index',
      component: () => import('../views/siswa/index.vue')
    },
    {
      path: '/murid/create',
      name: 'murid.create',
      component: () => import('../views/siswa/create.vue')
    },
    {
      path: '/murid/:id',
      name: 'murid.edit',
      component: () => import('../views/siswa/edit.vue')
    },

    // class route 
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
