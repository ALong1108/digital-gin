package main

import (
	"context"
	"digital-gin/configs"
	"digital-gin/middleware/db"
	"digital-gin/middleware/logger"
	"digital-gin/router"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	var err error
	if err = configs.Init(); err != nil {
		panic(err)
	}

	var r = gin.New()
	r.Use(logger.Logger(), gin.Recovery())

	// register routers
	router.Groups(r)

	// http server
	var httpConfig = configs.GetHTTPServer()
	srv := &http.Server{
		Addr:         httpConfig.Address,
		Handler:      r,
		ReadTimeout:  httpConfig.ReadTimeout * time.Millisecond,
		WriteTimeout: httpConfig.WriteTimeout * time.Millisecond,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	<-ctx.Done()
	stop()
	logger.Info("graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		logger.Fatal("Register forced to shutdown: ", err)
	}
	db.Close()

	logger.Info("Register exiting")
}
