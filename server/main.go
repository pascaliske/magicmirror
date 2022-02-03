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
	"github.com/pascaliske/magicmirror/metrics"
	"github.com/pascaliske/magicmirror/proxy"
	"github.com/pascaliske/magicmirror/socket"
)

var Version string
var GitCommit string

func main() {
	// parse config
	cfg, err := config.Parse()
	if err != nil {
		fmt.Println("Error: Couldn't parse config")
		return
	}

	// build information
	figure.NewFigure("MagicMirror", "graffiti", true).Print()
	fmt.Printf("\nVersion %s @ %s\n", color.CyanString(Version), color.CyanString(GitCommit))

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
	server.GET("/health", health.Handler(cfg, server))
	server.GET("/socket", socket.Handler(cfg, server))
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		Index: "index.html",
		HTML5: true,
	}))

	// metrics
	if cfg.Metrics.Enabled {
		server.Use(metrics.Middleware(cfg))
		server.GET(cfg.Metrics.Path, metrics.Handler(cfg, server))
	}

	// development proxy
	if cfg.Environment != "production" {
		server.Use(proxy.Handler(cfg, server, "http://localhost:4200"))
	}

	// start server
	go listen(cfg, server)

	// graceful shutdown
	shutdown(server)
}

func listen(cfg config.Config, server *echo.Echo) {
	fmt.Printf("Server is listening on %s\n", color.CyanString(fmt.Sprintf(":%d", cfg.Port)))

	if cfg.Environment != "production" {
		fmt.Printf("Using %s proxy for %s\n", cfg.Environment, color.CyanString("http://localhost:4200"))
	}

	// start server
	if err := server.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil && err != http.ErrServerClosed {
		server.Logger.Fatal(err)
	}
}

func shutdown(server *echo.Echo) {
	// wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("\nGracefully shutting down server...")

	// timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// shutdown server
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
