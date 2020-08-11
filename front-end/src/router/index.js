import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import ShoppingList from '@/views/ShoppingList'
import Admin from '@/views/Admin'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/list',
      name: 'ShoppingList',
      component: ShoppingList
    },
    {
      path: '/admin',
      name: 'Admin',
      component: Admin
    }
  ]
})

//跳转前设置title
// router.beforeEach((to, from, next) => {
//     window.document.title = to.meta.title;
//     next();
// });

//跳转后设置scroll为原点
router.afterEach((to, from, next) => {
    window.scrollTo(0, 0);
});

export default router
