package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rob-bender/nfc-cash-backend/pkg/ws"
)

type CreateRoomReq struct {
	ID string `json:"id"`
}

func (h *Handler) createRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &ws.Room{
		ID:      req.ID,
		Clients: make(map[string]*ws.Client),
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

func (h *Handler) joinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	uid := c.Query("uid")

	cl := &ws.Client{
		Conn:    conn,
		Message: make(chan *ws.Message, 10),
		RoomID:  roomID,
		Uid:     uid,
	}

	m := &ws.Message{
		Content: "A new user has joined the room",
		RoomID:  roomID,
		Uid:     uid,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

type RoomRes struct {
	ID string `json:"id"`
}

func (h *Handler) getRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID: r.ID,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	Uid string `json:"uid"`
}

func (h *Handler) getClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			Uid: c.Uid,
		})
	}

	c.JSON(http.StatusOK, clients)
}
