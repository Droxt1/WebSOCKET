package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}

}

type CreateRoomRequest struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//hub lock for safe access to Rooms
	h.hub.mu.Lock()
	defer h.hub.mu.Unlock()

	if _, exists := h.hub.Rooms[req.ID]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room already exists"})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true

	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	// Parse the room ID from the request URL
	roomID := c.Param("roomId")

	// Check if the room exists
	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room does not exist"})

		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upgrade connection"})
		return
	}

	clientID := c.Query("userId")
	username := c.Query("username")

	client := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		Username: username,
		RoomID:   roomID,
	}

	m := &Message{
		Content:  "New user joined the room",
		RoomID:   roomID,
		Username: username,
		IsDirect: true, // Mark as a direct message
	}

	h.hub.Register <- client
	h.hub.Broadcast <- m

	go client.writeMessage()
	client.readMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)
	for _, room := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   room.ID,
			Name: room.Name,
		})
	}
	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	roomId := c.Param("roomId")

	// Check if the room exists
	h.hub.mu.Lock()
	room, ok := h.hub.Rooms[roomId]
	h.hub.mu.Unlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	var clients []ClientRes
	h.hub.mu.Lock()
	for _, c := range room.Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}
	h.hub.mu.Unlock()

	c.JSON(http.StatusOK, clients)
}
