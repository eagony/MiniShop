// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

// 引入配置后的axios
import axios from './http.js'
Vue.prototype.$http = axios

// 引入Vuex
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    shoppingCart: []
  },
  getters: {
    // 分类筛选
  },
  mutations: {
    // 添加商品到购物车
    addGoodToCart(state, obj){
      const isAdded = state.shoppingCart.find(item => item.name === obj.name)
      // 不存在新增，存在数量+1
      if (isAdded) {
        isAdded.count++
      } else {
        state.shoppingCart.push({
          name: obj.name,
          price: obj.price,
          count: 1,
        })
      }
    },

    // 从购物车移除商品
    removeGoodFromCart(state, obj){
      const index = state.shoppingCart.findIndex(item => item.name === obj.name);
      console.log(name, index)
      if (index != -1) {
        if (state.shoppingCart[index].count > 1) {
          state.shoppingCart[index].count--
        } else {
          state.shoppingCart.splice(index, 1)
        }
      } else {
        alert("购物车里没有哦")
      }
    },
        // 清空购物车
    emptyCart(state){
      state.shoppingCart = []
    }
  }
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
