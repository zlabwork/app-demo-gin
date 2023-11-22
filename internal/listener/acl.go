package listener

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Acl @docs https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/
func Acl() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}
