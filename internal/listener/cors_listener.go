package listener

import (
	"github.com/gin-gonic/gin"
)

// @docs https://github.com/gin-contrib/cors

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Keep-Alive, User-Agent, Cache-Control, Content-Type, X-CSRF-Token, X-User-Token, X-Requested-With, If-Modified-Since, Authorization")

		c.Next()

	}
}
