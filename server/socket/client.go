package socket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
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

	// build client settings
	settings := BuildSettings()

	// send register message
	client.SendAction("register", settings)

	// send reload message on config changes
	config.OnChange(func() {
		client.SendAction("reload", nil)
	})

	return
}

func (client Client) Read(c echo.Context) {
	// unregister client
	defer delete(clients, client)

	for {
		message := SocketMessage{}

		// read message from client
		if err := client.socket.ReadJSON(&message); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger().Error(err)
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
		// send message to client
		message := <-client.send
		client.socket.WriteJSON(message)
	}
}

func (client Client) SendAction(action string, payload interface{}) {
	message := SocketMessage{Action: action, Payload: payload}
	client.socket.WriteJSON(message)
}
