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
      <!-- 购物清单 -->
      <table class="table">
        <thead>
          <tr>
            <th scope="col">商品名</th>
            <th scope="col">单价</th>
            <th scope="col">数量</th>
            <th scope="col">总价</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="good in shoppingList" v-bind:key="good.name">
            <td>{{ good.name }}</td>
            <td>{{ good.price }}</td>
            <td>{{ good.count }}</td>
            <td>{{ good.price * good.count }}</td>
          </tr>
        </tbody>
      </table>
      <div class="row mt-5 float-right">
        <form>
          <div class="form-group">
            <label for="inputAddress">地址：</label>
            <input type="text" class="form-control" id="inputAddress" v-model="orderForm.address"/>
          </div>
          <div class="form-group">
            <label for="inputRemark">备注：</label>
            <input type="text" class="form-control" id="inputRemark" v-model="orderForm.remark">
          </div>
          <div class="form-group text-center">
            <h3 >总共应付：{{ costAll }}元</h3>
            <button type="submit" class="btn btn-primary mt-5 float-right" @click="commitOrder">下单</button>
          </div>
        </form>
      </div>
    </div>
 </div>
</template>

<script>
  export default {
    name: 'ShoppingList',
    data () {
      return {
        orderForm: {
          detail: '',
          status: 'Created',
          amount: 0,
          remark:'',
          address:'',
        }
      }
    },
    computed: {
      shoppingList() {
        return this.$store.state.shoppingCart
      },
      costAll(){
        let cost = 0;
        this.shoppingList.forEach(item => {
          cost += item.price * item.count;
        });
        return cost;
      }
    },
    methods: {
      commitOrder () {
        this.orderForm.amount = this.costAll
        this.shoppingList.forEach(item => {
          this.orderForm.detail += (item.name + " " + item.price + " " + item.count + ";")
        });
        const path = '/api/v1/orders'
        const payload = {
          detail: this.orderForm.detail,
          status: 'Commit',
          amount: this.orderForm.amount,
          remark: this.orderForm.remark,
          address: this.orderForm.address,
        }
        const json = JSON.stringify(payload)
        this.$http.post(path, json, {headers: {'Content-Type': 'application/json'}})
        .then((res) => {
          console.log(res.data.data)
        })
        .catch((err) => {
          alert(err)
        })
      }
    },
    created() {
    }
  }
</script>
