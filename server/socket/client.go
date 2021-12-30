package socket

import (
	"encoding/json"

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

type Settings struct {
	Language  string  `json:"language"`
	Units     string  `json:"units"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var clients map[Client]bool = make(map[Client]bool)

func CreateClient(cfg config.Config, socket *websocket.Conn) (client Client) {
	// create client
	client = Client{UUID: uuid.NewString(), socket: socket, send: make(chan SocketMessage)}

	// save client reference
	clients[client] = true

	// build client settings
	settings := client.BuildSettings(cfg)
	payload, _ := json.Marshal(settings)

	// send register message
	message := SocketMessage{UUID: client.UUID, Action: "register", Payload: string(payload)}
	client.socket.WriteJSON(message)
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

		// broadcast message to all clients
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

func (client Client) BuildSettings(cfg config.Config) (settings Settings) {
	settings = Settings{Language: cfg.Language, Units: cfg.Units, Latitude: cfg.Location.Latitude, Longitude: cfg.Location.Longitude}
	return
}
