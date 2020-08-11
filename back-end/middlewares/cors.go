package middlewares

import (
	"github.com/gin-gonic/gin"
)

// CorsHandler 跨域请求处理
func CorsHandler() gin.HandlerFunc {
	// return func(c *gin.Context) {
	// 	// 允许访问所有域
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	// 是否需要带cookie信息 默认为true
	// 	c.Header("Access-Control-Allow-Credentials", "true")
	// 	// Header的类型
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 	// 支持的跨域请求方法
	// 	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	// 	// 缓存预检请求的时间，单位秒
	// 	c.Header("Access-Control-Max-Age", "21600")
	// 	// 设置返回格式是Json
	// 	c.Set("content-type", "application/json")

	// 	// 放行所有OPTIONS方法，本项目直接返回204
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}

	// 	c.Next()
	// }
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
