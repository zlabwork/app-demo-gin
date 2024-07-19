package main

import (
	"app/internal/help"
	"app/internal/route"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {

	// server
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
			log.Fatalf("Listen: %s\n", err)
		}
	}()
	fmt.Println("---- Server Is Started ----")

	// logs
	logFile, err := os.OpenFile(help.Dir.Logs+"app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("can not open the log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("Server Is Started")

	// shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server Exiting")
}
