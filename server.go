package main

import (
	"context"
	"go-challenge/config"
	"go-challenge/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func serverInit() {
	cnf := config.GetConfig()
	router.InitLocationService()
	r := router.Router()
	srv := &http.Server{
		Addr:    ":" + cnf.GinPort,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
