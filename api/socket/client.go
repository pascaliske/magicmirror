package socket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/metrics"
)

type Client struct {
	UUID   string
	socket *websocket.Conn
	send   chan SocketMessage
}

var clients map[Client]bool = make(map[Client]bool)

func CreateClient(socket *websocket.Conn) (client Client) {
	// create client
	client = Client{UUID: uuid.NewString(), socket: socket, send: make(chan SocketMessage)}

	// save client reference
	clients[client] = true

	// update metric
	metrics.SocketConnections.WithLabelValues().Set(float64(len(clients)))

	// build client settings
	settings := client.CreateSettings()

	// send register message
	client.SendAction("register", settings)

	return
}

func (client Client) Read(c echo.Context) {
	// send reload message on config changes
	cancel := config.OnChangeSuccess(client.UUID, func() {
		client.SendAction("reload", nil)
	})

	// unregister client
	defer cancel()
	defer delete(clients, client)

	// update metric
	defer metrics.SocketConnections.WithLabelValues().Set(float64(len(clients)))

	for {
		message := SocketMessage{}

		// read message from client
		if err := client.socket.ReadJSON(&message); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error(err.Error())
			}
			break
		}

		// broadcast message to other clients
		for target := range clients {
			if target.UUID != client.UUID {
				target.send <- message
			}
		}
	}
}

func (client Client) Write(c echo.Context) {
	// close connection
	defer client.socket.Close()

	for {
		// get message for client
		message := <-client.send

		// send message to client
		if err := client.socket.WriteJSON(message); err != nil {
			logger.Error(err.Error())
		}
	}
}

func (client Client) SendAction(action string, payload interface{}) {
	// build message for client
	message := SocketMessage{Action: action, Payload: payload}

	// send message to client
	if err := client.socket.WriteJSON(message); err != nil {
		logger.Error(err.Error())
	}
}
