import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Home from '../components/Home.vue'
import Welcome from '../components/Welcome.vue'
import User from '../components/user/User.vue'

Vue.use(VueRouter)

const routes = [{
  path: '/',
  redirect: '/login'
},
{
  path: '/login',
  component: Login
},
{
  path: '/home',
  component: Home,
  redirect: '/welcome',
  children: [{
    path: '/welcome',
    component: Welcome
  },
  {
    path: '/users',
    component: User
  }]
}
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  // next() 放行 next('/ddd') 强制跳转
  if (to.path === '/login') return next()
  const tokerStr = window.sessionStorage.getItem('token')
  if (!tokerStr) return next('/login')
  next()
})
export default router
