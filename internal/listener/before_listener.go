package listener

import (
	"app/internal/msg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"regexp"
)

func Before() gin.HandlerFunc {
	return func(c *gin.Context) {

		// maintenance response
		if os.Getenv("APP_MAINTENANCE") == "true" {
			pattern := `^/(v[0-9]+|api)/.*`
			re := regexp.MustCompile(pattern)
			if re.MatchString(c.Request.URL.Path) {
				c.AbortWithStatusJSON(http.StatusServiceUnavailable, msg.DataWrap{
					Status:  msg.StatusMaintenance,
					Message: "Under Maintenance.",
				})
			} else {
				c.HTML(http.StatusServiceUnavailable, "errors/maintenance.html", gin.H{})
			}
			c.Abort()
			return
		}

		trace := uuid.New().String()
		// c.Set("traceId", trace)
		c.Header("X-Request-Id", trace)

		c.Header("X-Frame-Options", "SAMEORIGIN")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "origin")
		// c.Header("Content-Security-Policy", "default-src 'self';")

		c.Next()
	}
}
