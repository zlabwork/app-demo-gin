package listener

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Before() gin.HandlerFunc {
	return func(c *gin.Context) {
		trace := uuid.New().String()
		// c.Set("traceId", trace)
		c.Header("X-Request-Id", trace)
		c.Next()
	}
}
