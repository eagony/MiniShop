<template>
  <div class="">

    <nav class="navbar navbar-expand-lg navbar-light rounded-0 shadow-sm bg-white p-3 mb-3">
      <a class="navbar-brand" href="/"><h3>3008小卖部</h3></a>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <button class="btn btn-light float-right ml-auto" type="button">
          <router-link class="text-decoration-none" to="/list" style="color: black;"><span class="badge badge-danger mr-2">{{ countAll }}</span>购物车</router-link>
        </button>
      </div>
    </nav>
    <div class="container">

      <!-- 商品列表 -->
      <div class="row row-cols-1 row-cols-md-4 mt-5">
        <div v-for="good in goods" :key="good.id" class="col mb-4">
          <div class="card h-80">
            <p class="text-center mt-3">{{ good.name }}</p>
            <img :src="good.pictureurl" class="card-img-top" alt="商品图">
            <div class="card-body text-center">
              <h5 class="card-title text-center">
                单价：{{ good.price }} 销量：{{ good.sales }}
              </h5>
              <h5 class="card-title text-center">
                库存：{{ good.stock }}
              </h5>
              <button class="btn btn-secondary rounded mr-3" @click="countMinus(good)">减一</button>
              <button class="btn btn-info rounded" @click="countPlus(good)">加一</button>
            </div>
          </div>
        </div>
      </div>
    </div>
 </div>
</template>


<script>
  export default {
    name: 'Home',
    data () {
      return {
        goods: '',
        show: true,
      }
    },
    methods: {
      getData () {
        const path = '/api/v1/goods'
        this.$http.get(path
        ).then((res) => {
          this.goods = res.data.data
          console.log("商品加载完毕。")
          this.show = false
         }).catch((err) => { console.log(err) })
      },
      countPlus (good) {
        const isAdded = this.$store.state.shoppingCart.find(item => item.name === good.name)

        if (good.stock > 0 ) {
          if (isAdded === undefined || isAdded.count < good.stock) {
            this.$store.commit('addGoodToCart', {
              name: good.name,
              price: good.price
            })
          } else {
            alert("库存不足")
            return
          }
        } else {
          alert("库存不足")
          return
        }
      },
      countMinus (good) {
        this.$store.commit('removeGoodFromCart', {
          name: good.name,
        })
      }
    },
    computed: {
      shoppingList() {
        return this.$store.state.shoppingCart
        },

      //购物车商品总数
      countAll(){
        let count = 0;
        this.shoppingList.forEach(item => {
          count += item.count;
        });
        return count;
      },
    },
    created() {
      this.getData()
    }
  }
</script>

<style>
</style>
