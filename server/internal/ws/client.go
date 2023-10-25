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
}

// write message to client
func (c *Client) writeMessage() {
	//close the connection when the function returns because we don't want to leave the connection open
	//defer is used to ensure that the connection is closed when the function returns
	defer func() {
		c.Conn.Close()
	}()
	//loop over the message channel and write the message to the client
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

// read message from client
func (c *Client) readMessage(hub *Hub) {
	defer func() {
		//unregister the client from the hub when the function returns
		//because we don't want to leave the client registered in the hub
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
		}
		//send the message to the hub
		hub.Broadcast <- message

	}

}
