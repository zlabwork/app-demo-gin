package listener

import (
	"app/internal/help"
	"app/internal/msg"
	"github.com/gin-gonic/gin"
	"net/http"
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

		// check
		if help.Env.IsLocal {
			c.Set("userId", int64(123456))
		} else {
			tk, err := help.Libs.Token.ParseTokenString(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, msg.DataWrap{
					Status:  msg.StatusUnauthorized,
					Message: "invalid authorization token",
				})
				return
			}
			c.Set("userId", tk.UserId)
		}

		c.Next()

	}
}
