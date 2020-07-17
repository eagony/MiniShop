<template>
  <div class="container">
    <!-- 标题 -->
    <div class="row d-block mt-5">
      <div class="card mb-3">
        <div class="card-body">
          <h3 class="card-title text-center align-center">3008小卖部</h3>
        </div>
      </div>
    </div>

    <!-- 商品列表 -->
    <div class="row row-cols-1 row-cols-md-4 mt-5">
      <div v-for="good in goods" :key="good.name" class="col mb-4">
        <div class="card h-80">
          <p class="text-center mt-3">{{ good.name }}</p>
          <img :src="good.pictureURL" class="card-img-top" alt="good.name">
          <div class="card-body text-center">
            <h5 class="card-title text-center">
            单价：{{ good.price }} 库存：{{ good.stock }}
            </h5>
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
        goods: [
          {
            name: '柠檬茶',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t1/25746/22/15383/185658/5caea739E0f0d863a/d042ffcbc4e42384.jpg',
            price: 3,
            stock: 10
          },
          {
            name: '1',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t25249/58/2315766440/164091/9488a877/5bc84c56N50ffba2a.jpg',
            price: 3,
            stock: 10
          },
          {
            name: '2',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t1/33888/16/2940/319842/5cb30cdeEe9df33e4/681f71f7b55ca696.jpg',
            price: 3,
            stock: 10
          },
          {
            name: '3',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t1/112078/5/3705/133960/5ea951ddEd704ce5b/eb47e1f6656ec4d7.jpg',
            price: 3,
            stock: 10
          },
          {
            name: '3',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t1/66985/23/8361/710344/5d66657aEd541ee41/f02bd0b70d53bb9b.jpg',
            price: 3,
            stock: 10
          },
          {
            name: '4',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t5584/298/7207844933/163376/21ec1cc8/596ed0dfN0a868c6c.jpg',
            price: 3,
            stock: 10
          },
          {
            name: '5',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t1/101949/34/17038/79321/5e82d311Ea1719801/d5c36227db453cb9.jpg',
            price: 3,
            stock: 10
          },
          {
            name: '6',
            pictureURL: 'https://img14.360buyimg.com/n0/jfs/t1/66489/24/8658/617013/5d68ed99E73fe0dab/c174af452d1034a3.jpg',
            price: 3,
            stock: 10
          },
        ]
      }
    }
  }
</script>

<style scope>
body {
  background: #f0f0f4;
}
</style>