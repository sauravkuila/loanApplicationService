package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aspire/dao"
)

var ctx context.Context
var srv *http.Server

func main() {
	ctx = context.Background()
	initRouter()
	dao.Init()
	startRouter()
	addShutdownHook()
}

func startRouter() {
	srv = &http.Server{
		Addr:    ":8080",
		Handler: getRouter(ctx),
	}

	// run api router
	log.Println("starting router")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error starting server")
		}
	}()
}

func addShutdownHook() {
	//when receive interruption from system shutdown server
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Println("Shutting down server")
	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server forced to shutdown")
	}

	log.Println("Server exiting")
}
