package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"data":    "pong",
	})
}
