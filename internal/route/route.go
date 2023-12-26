package route

import (
	"app/internal/api"
	"app/internal/listener"
	"app/internal/web"
	"github.com/gin-gonic/gin"
)

// @docs https://gin-gonic.com/docs/examples/define-format-for-the-log-of-routes/
// @docs https://gin-gonic.com/docs/examples/grouping-routes/
// @docs https://gin-gonic.com/docs/examples/custom-middleware/
// @docs https://gin-gonic.com/docs/examples/using-middleware/

func GetRoute() *gin.Engine {

	r := gin.Default()
	r.Use(listener.Before())

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/assets", "./assets")
	r.GET("/", web.DefaultHandler)
	r.GET("/sample", web.SampleHandler)
	r.GET("/ping", web.PingHandler)
	// api
	v1 := r.Group("/v1")
	v1.Use(listener.Acl())
	{
		v1.GET("/public_key", api.PublicKeyHandler)
		v1.GET("/token", api.GenerateTokenHandler)
		v1.PUT("/token", api.RefreshTokenHandler)
	}

	return r
}
