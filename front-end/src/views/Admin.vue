<template>
  <div class="">
    <nav class="navbar navbar-expand-lg navbar-light rounded-0 shadow-sm bg-white p-3 mb-3">
      <a class="navbar-brand" href="/"><h3>3008小卖部</h3></a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
         <button class="btn btn-light float-right ml-auto" type="button">
           <router-link class="text-decoration-none" to="/" style="color: black;">返回</router-link>
         </button>
      </div>
    </nav>
    <div class="container">
      <!-- 库存清单 -->
      <table class="table">
        <thead>
          <tr>
            <th scope="col">商品名</th>
            <th scope="col">单价</th>
            <th scope="col">销量</th>
            <th scope="col">库存</th>
            <th scope="col">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="product in productList" v-bind:key="product.name">
            <td>{{ product.name }}</td>
            <td>{{ product.price }}</td>
            <td>{{ product.sales }}</td>
            <td>{{ product.stock }}</td>
            <td>
              <button class="btn btn-danger">删除</button>
              <button class="btn btn-primary">更新</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="row mt-5">
        <h3>今日营业额：{{}}元</h3><hr>
        <h3>总营业额：{{}}元</h3>
      </div>
    </div>
 </div>
</template>

<script>
  export default {
    name: 'Admin',
    data() {
      return {
        productList: '',
        }
    },
    methods: {
      getProductList () {
        const path = 'http://localhost:8000/api/v1/goods'
        this.$http.get(path
        ).then((res) => {
          this.productList = res.data.data
          console.log("商品加载完毕。")
          this.show = false
         }).catch((err) => { console.log(err) })
      },
    },
    created() {
      this.getProductList()
    }
  }
</script>
