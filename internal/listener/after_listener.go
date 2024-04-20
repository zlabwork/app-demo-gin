package listener

import (
	"app/internal/msg"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func After() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用缓冲区来捕获响应体
		w := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		// 执行处理器
		c.Next()

		c.Writer = w.ResponseWriter

		// 修改响应体
		var resp msg.DataWrap
		if json.Unmarshal(w.body.Bytes(), &resp) != nil {
			log.Println("response listener error")
			return
		}
		resp.Message = resp.Message + " - fixed"
		bs, err := json.Marshal(resp)
		if err != nil {
			log.Println("response listener error")
			return
		}

		c.Writer.Write(bs)
	}
}

// responseWriter 是一个用于捕获响应体的自定义 ResponseWriter
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 方法用于捕获响应体的写入操作
func (w *responseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b) // 将响应体写入缓冲区
}
