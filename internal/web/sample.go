package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
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
