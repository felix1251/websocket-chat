package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message

		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()

		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		hub.Broadcast <- msg
	}
}
