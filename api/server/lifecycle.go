package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pascaliske/magicmirror/logger"
)

func (server *Server) listen() {
	logger.Info("Server is listening on %s", fmt.Sprintf(":%d", server.port))

	// start server
	if err := server.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}

func (server *Server) shutdown() {
	// wait for interrupt or terminate signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// defer clean-up
	defer signal.Stop(quit)
	defer cancel()

	logger.Raw("")
	logger.Info("Gracefully shutting down server...")

	// shutdown server
	if err := server.http.Shutdown(ctx); err != nil {
		logger.Debug(err.Error())
		logger.Fatal("Could not shutdown server gracefully!")
		os.Exit(1)
	}

	os.Exit(0)
}
