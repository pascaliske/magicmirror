package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/health"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/metrics"
	"github.com/pascaliske/magicmirror/proxy"
	"github.com/pascaliske/magicmirror/socket"
)

var Version string
var GitCommit string

func main() {
	// build information
	figure.NewFigure("MagicMirror", "graffiti", true).Print()
	logger.Raw("\nVersion %s @ %s\n", color.CyanString(Version), color.CyanString(GitCommit))

	// parse config
	if err := config.Parse(); err != nil {
		logger.Error("Couldn't parse config")
		return
	}

	// configure logging
	logger.SetLevel(config.GetString("Log.Level"))
	config.OnChange(time.Now().String(), func() {
		logger.SetLevel(config.GetString("Log.Level"))
	})

	// setup server
	server := echo.New()
	server.HidePort = true
	server.HideBanner = true
	server.Use(middleware.CORS())
	server.Use(middleware.Secure())
	server.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, ip=${remote_ip}, latency=${latency_human}\n",
		Skipper: func(c echo.Context) bool {
			return c.Response().Status >= 200 && c.Response().Status <= 299
		},
	}))

	// endpoints
	server.GET("/health", health.Handler(server))
	server.GET("/socket", socket.Handler(server))
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		Index: "index.html",
		HTML5: true,
	}))

	// metrics
	if config.GetBool("Metrics.Enabled") {
		server.Use(metrics.Middleware())
		server.GET(config.GetString("Metrics.Path"), metrics.Handler(server))
	}

	// development proxy
	if config.GetString("Environment") != "production" {
		server.Use(proxy.Handler(server, "http://localhost:4200"))
	}

	// start server
	go listen(server)

	// graceful shutdown
	shutdown(server)
}

func listen(server *echo.Echo) {
	logger.Info("Server is listening on %s", color.CyanString(fmt.Sprintf(":%d", config.GetInt("Port"))))

	// start server
	if err := server.Start(fmt.Sprintf(":%d", config.GetInt("Port"))); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
	}
}

func shutdown(server *echo.Echo) {
	// wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Raw("\nGracefully shutting down server...")

	// timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// shutdown server
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal(err.Error())
	}
}
