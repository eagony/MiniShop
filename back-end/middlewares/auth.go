package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
}

// AuthMiddleware 简单的授权验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		// fmt.Println(token, "<==>", os.Getenv("SECRET_KEY"))
		if token == "" || token != os.Getenv("SECRET_KEY") {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"Auth error": "Unauthorized.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
