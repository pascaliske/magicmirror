package server

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/health"
	"github.com/pascaliske/magicmirror/metrics"
	"github.com/pascaliske/magicmirror/socket"
)

func (server Server) setupRoutes() {
	// health
	server.router.GET("/health", health.Handler())

	// socket
	server.router.GET("/socket", socket.Handler())

	// metrics
	if config.GetBool("Metrics.Enabled") {
		server.router.Use(metrics.Middleware())
		server.router.GET(config.GetString("Metrics.Path"), metrics.Handler())
	}

	// static files
	server.router.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		Index: "index.html",
		HTML5: true,
	}))
}
