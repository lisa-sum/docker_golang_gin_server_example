package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// CORS 跨域处理配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"} // 允许的请求方法
	corsConfig.AllowAllOrigins = true                                           // 是否允许所有IP的请求
	r.Use(cors.New(corsConfig))

	r.Any("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run("0.0.0.0:8080")
}
