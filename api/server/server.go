package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pascaliske/magicmirror/config"
)

type Server struct {
	port   int
	router *echo.Echo
	http   http.Server
}

func SetupAndListen() {
	server := Server{port: config.GetInt("Port")}

	// setup router
	server.router = echo.New()
	server.http = http.Server{Addr: fmt.Sprintf(":%d", server.port), Handler: server.router}

	// setup middlewares & routes
	server.setupLogger()
	server.setupMiddlewares()
	server.setupRoutes()
	server.setupProxy("http://localhost:4200")

	// start server
	go server.listen()

	// graceful shutdown
	server.shutdown()
}
