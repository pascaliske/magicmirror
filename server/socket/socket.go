package socket

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
)

type SocketMessage struct {
	UUID    string `json:"uuid,omitempty"`
	Action  string `json:"action"`
	Payload string `json:"payload,omitempty"`
}

var upgrader = websocket.Upgrader{}

func Handler(cfg config.Config, server *echo.Echo) echo.HandlerFunc {
	return func(c echo.Context) error {
		// upgrade connection to socket
		socket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

		// could not upgrade connection
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		// create client for connection
		client := CreateClient(cfg, socket)

		// handle client io
		go client.Read(c)
		go client.Write(c)

		return nil
	}
}
