package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	Username string `json:"username"`
	RoomID   string `json:"RoomId"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
	IsDirect bool   `json:"isDirect"`
}

// write message to client
func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		err := c.Conn.WriteJSON(message)
		if err != nil {
			return
		}
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error %v", err)
			}
			break
		}
		message := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
			IsDirect: true, // Mark as direct message
		}

		hub.Broadcast <- message

	}
}
