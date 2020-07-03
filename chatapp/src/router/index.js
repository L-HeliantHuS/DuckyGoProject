import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: "主页"
    }
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: {
      title: "关于"
    }
  },
  {
    path: '/user/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {
      title: "登录"
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

// router.js
const baseTitle = "ChatAPP";
router.beforeEach((to, from, next) => {
  document.title = (to.meta.title ? to.meta.title : " ") + "  --  " + baseTitle;
  next()
});


export default router
