# MiniShop
一款简易的小商城应用，使用gin+vue开发。实现了管理商品，挑选商品，提交订单。

# 安装
1. 下载
    
    `git clone https://github.com/eagony/MiniShop.git`
2. 后端

    `cd backe-end`

    修改.env里的数据库等配置
    
    `go run main.go` 启动后端
3. 前端(自行安装node.js)
    
    `cd front-end`

    尽量不用cnpm安装，可以用淘宝的镜像源

    `npm install --registry=https://registry.npm.taobao.org` 

    `npm run dev` 启动前端

    浏览器打开localhost:8080
