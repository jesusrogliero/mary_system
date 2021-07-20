import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import about from '../views/About.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/about', name: 'About', component: about }
]

const router = new VueRouter({
  // mode: 'abstract',  Debe ser activado al compilar a .exe
  mode: 'history', // usar en modo desarrollo
  base: process.env.BASE_URL,
  routes
})

export default router
