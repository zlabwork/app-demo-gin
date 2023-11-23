package listener

import (
	"github.com/gin-gonic/gin"
)

// Acl @docs https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/
func Acl() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		l := len(token)
		if l == 0 {
			token, _ = c.GetQuery("token")
		} else if l > 7 && token[:6] == "Bearer" {
			token = token[7:]
		}
		c.Set("userId", 123456)

		c.Next()

	}
}
