package web

import (
	"app/internal/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

// @docs https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/
// 可以通过 go-playground/validator 扩展自定义验证规则。
// 参考关键字 required omitempty min max len eq ne email url uuid alphanum numeric datetime=2006-01-02

type Person struct {
	Name     string `json:"name" binding:"required"`                // 必填字段
	Age      int    `json:"age" binding:"omitempty,min=18,max=130"` // 选填且最小值为 18, 最大 130
	Email    string `json:"email" binding:"required,email"`         // 必填且格式为 email
	Passcode string `json:"passcode" binding:"omitempty,min=6"`     // 选填且至少 6 个字符
}

func PingHandler(c *gin.Context) {

	// POST 请求做参数验证
	if c.Request.Method == http.MethodPost {
		var user Person
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": consts.StatusError, "message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  consts.StatusSuccess,
		"message": consts.StatusSuccess,
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
