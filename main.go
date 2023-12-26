package main

import (
	"app/internal/app"
	"app/internal/route"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	app.Boot()

	port := ":3000"
	if strings.TrimSpace(os.Getenv("APP_PORT")) != "" {
		port = ":" + os.Getenv("APP_PORT")
	}
	srv := &http.Server{
		Addr:         port,
		Handler:      route.GetRoute(),
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
