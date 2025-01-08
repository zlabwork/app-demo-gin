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

	// api
	pub := r.Group("/api")
	{
		pub.GET("/public_key", api.PublicKeyHandler)
		pub.GET("/token", api.GenerateTokenHandler)
		pub.PUT("/token", api.RefreshTokenHandler)
	}

	// v1
	v1 := r.Group("/v1")
	v1.Use(listener.Acl())
	{
	}

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/assets", "./public/assets")
	r.GET("/", web.DefaultHandler)
	r.GET("/sample", web.SampleHandler)
	r.Match([]string{"GET", "POST"}, "/ping", web.PingHandler)

	return r
}
