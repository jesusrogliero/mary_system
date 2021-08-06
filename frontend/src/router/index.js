import Vue from 'vue'
import VueRouter from 'vue-router'
import Person from '../views/Personas.vue'


Vue.use(VueRouter)

const routes = [
  { path: '/', name: 'Personas', component: Person }

]

const router = new VueRouter({
  mode: 'abstract',  //Debe ser activado al compilar a .exe
  //mode: 'history', // usar en modo desarrollo
  base: process.env.BASE_URL,
  routes
})

export default router
