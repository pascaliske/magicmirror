package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server Server) setupMiddlewares() {
	// basics
	server.router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))
	server.router.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	// security
	server.router.Use(middleware.CORS())
	server.router.Use(middleware.Secure())

	// logging
	server.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, ip=${remote_ip}, latency=${latency_human}\n",
		Skipper: func(c echo.Context) bool {
			return c.Response().Status >= 200 && c.Response().Status <= 299
		},
	}))
}
