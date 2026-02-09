package server

import (
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
)

type Server struct {
	router *echo.Echo
}

func SetupAndListen() {
	server := Server{}

	// setup router
	server.router = echo.New()
	server.router.HidePort = true
	server.router.HideBanner = true

	// setup middlewares & routes
	server.setupLogger()
	server.setupMiddlewares()
	server.setupRoutes()
	server.setupProxy("http://localhost:4200")

	// start server
	go server.listen(config.GetInt("Port"))

	// graceful shutdown
	server.shutdown()
}
