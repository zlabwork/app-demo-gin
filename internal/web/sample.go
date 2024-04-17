package web

import (
	"app/internal/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  msg.StatusSuccess,
		"message": msg.StatusSuccess,
		"data":    "pong",
	})
}

func DefaultHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home/main.html", gin.H{
		"title":      "Gin Web Framework " + gin.Version,
		"go_version": runtime.Version(),
		"framework":  "Gin " + gin.Version,
	})
}

func SampleHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home/sample.html", gin.H{
		"title": "Main website",
	})
}
