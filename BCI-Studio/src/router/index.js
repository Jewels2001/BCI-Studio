import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RecordView from '../views/RecordView.vue'
import PlayView from '../views/PlayView.vue'
import StudioView from '../views/StudioView.vue'
import AboutView from '../views/AboutView.vue'
import ContactView from '../views/ContactView.vue'
import NotFound from '../views/NotFound.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/recording',
    name: 'record',
    component: RecordView
  },
  {
    path: '/play',
    name: 'play',
    component: PlayView
  },
  {
    path: '/studio',
    name: 'studio',
    component: StudioView
  },
  {
    path: '/about',
    name: 'about',
    component: AboutView
  },
  {
    path: '/contact',
    name: 'contact',
    component: ContactView
  },
  {
    path: '/:pathMatch(.*)',
    component: NotFound
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
