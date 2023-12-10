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

// Hanldeconnection upgrades the http connection to a websocket connection
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	params := r.URL.Query()
	userName := params.Get("userName")
	roomName := params.Get("room")

	client := &Client{
		conn:     ws,
		userName: userName,
		Room:     Chat.GetRoom(roomName),
	}

	client.Room.Register <- client

	go client.Read()
	go client.Write()
}

// Read reads messages from the websocket connection
func (c *Client) Read() {
	defer func() {
		c.Room.Unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		c.Room.Broadcast <- Message{sender: c.userName, content: string(msg)}
	}
}

// Write writes messages to the websocket connection
func (c *Client) Write() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Room.Messages:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}
}
