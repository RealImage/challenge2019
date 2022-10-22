package main

import (
	"challenge2019/controller"
	Helper2 "challenge2019/helper"
	"challenge2019/service"
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Helper     = Helper2.NewHelper()
	Service    = service.NewService(Helper)
	Controller = controller.NewController(Service)
)

func main() {
	// crating blank engine - which is instance of framework where we can perform everything
	router := gin.New()

	// passing engine instance to perform related operation
	Controller.InstallRoute(router)

	// server configuration
	server := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info("Starting server")
		// Opens and starts listening to server on the mentioned address
		if err := server.ListenAndServe(); err != nil {
			log.Error(err.Error())
		}
	}()

	// To shutdown server gracefully
	waitForShutDown(server)
}

func waitForShutDown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Info("Shutting down")
	os.Exit(0)
}
