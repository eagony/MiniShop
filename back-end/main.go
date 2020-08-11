package main

import (
	"MiniShop/api"
	"MiniShop/extensions"
	"MiniShop/middlewares"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// 连接一次数据库，完成自动迁移
	db := extensions.GetDB()
	db.Close()
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
}

func main() {
	router := gin.Default()

	// 使用跨域请求中间件
	router.Use(middlewares.CorsHandler())

	// 分组
	apiV1 := router.Group("/api/v1")

	// 注册API
	new(api.GoodAPI).Register(apiV1)
	new(api.OrderAPI).Register(apiV1)

	// 测试连接
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	fmt.Println("SSL: ", os.Getenv("SSL"))
	port := os.Getenv("PORT")

	if os.Getenv("SSL") == "TRUE" {
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{}
		// Generated using sh generate-certificate.sh
		SSLKeys.CERT = "./cert/MyCA.cer"
		SSLKeys.KEY = "./cert/MyCA.key"
		router.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		router.Run(":" + port)
	}
}
