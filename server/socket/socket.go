package socket

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/metrics"
)

type SocketMessage struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload,omitempty"`
}

var upgrader = websocket.Upgrader{}

func Handler(server *echo.Echo) echo.HandlerFunc {
	logger.Debug("Socket endpoint ready at %s", color.CyanString("/socket"))

	// update metric
	if config.GetBool("Metrics.Enabled") {
		metrics.SocketClients.WithLabelValues().Set(float64(len(clients)))
	}

	return func(c echo.Context) error {
		// upgrade connection to socket
		socket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

		// could not upgrade connection
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		// create client for connection
		client := CreateClient(socket)

		// handle client io
		go client.Read(c)
		go client.Write(c)

		return nil
	}
}
