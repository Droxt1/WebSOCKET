package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// Check if the room exists
			if r, ok := h.Rooms[client.RoomID]; ok {
				// Check if the client is already in the room
				if _, exists := r.Clients[client.ID]; !exists {
					r.Clients[client.ID] = client
				}
			}

		case client := <-h.Unregister:
			// Check if the room exists
			if r, ok := h.Rooms[client.RoomID]; ok {
				// Check if the client is in the room
				if _, exists := r.Clients[client.ID]; exists {
					// Broadcast to all clients in the room that the client has left
					if len(r.Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  client.Username + " has left the room",
							RoomID:   client.RoomID,
							Username: client.Username,
						}
					}

					delete(r.Clients, client.ID)
					close(client.Message)
				}
			}

		case m := <-h.Broadcast:
			// Check if the room exists
			if r, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range r.Clients {
					cl.Message <- m
				}
			}
		}
	}
}
