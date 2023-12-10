package websockets

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// client is a websocket client that represents a connected user
type Client struct {
	conn     *websocket.Conn
	userName string
	Room     *Chat.Room
}

// message represent message sent by clients
type Message struct {
	sender  string `json:"sender"`
	content string `json:"content"`
}
