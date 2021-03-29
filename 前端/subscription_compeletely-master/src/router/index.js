import Vue from 'vue'

// eslint-disable-next-line import/no-duplicates
import Router from 'vue-router'
// eslint-disable-next-line import/no-duplicates
import VueRouter from 'vue-router'
/* eslint-disable */

Vue.use(Router)

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}


export const constantRoutes = [
    {

        name:'main',
        path:'/main',
        component:()=>import('../views/main.vue'),
        //访问控制
        beforeEnter: (to, from, next) => {
            if(!localStorage.getItem('token'))next({name:'index'})
            else next()
           }
    },
    {
        name:'index',
        path:'/index',
        component:()=>import('../views/Index.vue')
    },
    {
        name:'app',
        path:'/',
        redirect:'/index'
    }
    
]

const createRouter = () => new Router({
    mode: 'history', // require service support
    
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes
  })
  
const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
const newRouter = createRouter()
router.matcher = newRouter.matcher // reset router
}
  
export default router