package main

import (
	"app/internal/api"
	"app/internal/app"
	"app/internal/web"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app.Boot()

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/assets", "./assets")
	r.GET("/", web.DefaultHandler)
	r.GET("/sample", web.SampleHandler)
	r.GET("/ping", web.PingHandler)
	v1 := r.Group("/v1")
	{
		v1.GET("/public_key", api.PublicKeyHandler)
		v1.GET("/token", api.GenerateTokenHandler)
		v1.PUT("/token", api.RefreshTokenHandler)
	}

	port := ":3000"
	if strings.TrimSpace(os.Getenv("APP_PORT")) != "" {
		port = ":" + os.Getenv("APP_PORT")
	}
	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
