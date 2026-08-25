import type { RouteRecordRaw } from 'vue-router'

export const routes: RouteRecordRaw[] = [
  { path: '/', component: () => import('./pages/Home.vue') },
  { path: '/community', component: () => import('./pages/Community.vue') },
  { path: '/faq', component: () => import('./pages/Faq.vue') },
]
