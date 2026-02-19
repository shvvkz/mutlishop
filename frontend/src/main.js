import { createApp } from 'vue'
import { createWebHistory, createRouter } from 'vue-router'
import App from './App.vue'
import PublicView from './views/PublicView.vue'
import AdminView from './views/AdminView.vue'
import LoginView from './views/LoginView.vue'
import { getToken } from './services/api'

const routes = [
  { path: '/', component: PublicView },
  { path: '/login', component: LoginView },
  { path: '/admin', component: AdminView, meta: { requiresAuth: true } }
]

const router = createRouter({ history: createWebHistory(), routes })

router.beforeEach((to) => {
  const isAuthenticated = Boolean(getToken())

  if (to.meta.requiresAuth && !isAuthenticated) {
    return { path: '/login', query: { redirect: to.fullPath } }
  }

  if (to.path === '/login' && isAuthenticated) {
    return { path: '/admin' }
  }

  return true
})

createApp(App).use(router).mount('#app')
