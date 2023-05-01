package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rob-bender/nfc-cash-backend/pkg/wsRoom"
)

// @Summary		CreateRoom
// @Tags			room
// @Description	create room
// @ID				create-room
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/room/create-room [post]
func (h *Handler) createRoom(c *gin.Context) {
	resCreateRoom, statusCode, err := h.services.CreateRoom()
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resCreateRoom,
		})
		return
	}
	if len(resCreateRoom) > 0 {
		h.hub.Rooms[resCreateRoom] = &wsRoom.Room{
			UidRoom: resCreateRoom,
			Clients: make(map[string]*wsRoom.Client),
		}
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешное создание комнаты",
			"result":  resCreateRoom,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка создания комнаты",
			"result":  resCreateRoom,
		})
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @Summary		JoinRoom
// @Tags			room
// @Description	join room
// @ID				join-room
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/room/join-room/:roomId [get]
func (h *Handler) joinRoom(c *gin.Context) {
	uidRoom := c.Param("roomId")
	uidUser := c.Query("uidUser")
	resJoinRoom, statusCode, err := h.services.JoinRoom(uidRoom, uidUser)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cl := &wsRoom.Client{
		Conn:    conn,
		Message: make(chan *wsRoom.Message, 10),
		RoomID:  uidRoom,
		UidUser: resJoinRoom,
	}

	m := &wsRoom.Message{
		Content: "A new user has joined the room",
		RoomID:  uidRoom,
		UidUser: resJoinRoom,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

// @Summary		LeaveRoom
// @Tags			room
// @Description	leave room
// @ID				leave-room
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.LeaveRoom	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/room/leave-room [post]
func (h *Handler) leaveRoom(c *gin.Context) {
	type Body struct {
		UidRoom string `json:"uidRoom"`
		UidUser string `json:"uidUser"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.LeaveRoom(body.UidRoom, body.UidUser)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "пользователь успешно покинул чат",
	})
}

type RoomRes struct {
	UidRoom string `json:"uidRoom"`
}

// func (h *Handler) getRooms(c *gin.Context) {
// 	rooms := make([]RoomRes, 0)

// 	for _, r := range h.hub.Rooms {
// 		rooms = append(rooms, RoomRes{
// 			UidRoom: r.UidRoom,
// 		})
// 	}

// 	c.JSON(http.StatusOK, rooms)
// }

// type ClientRes struct {
// 	UidUser string `json:"uidUser"`
// }

// func (h *Handler) getClients(c *gin.Context) {
// 	var clients []ClientRes
// 	roomId := c.Param("roomId")

// 	if _, ok := h.hub.Rooms[roomId]; !ok {
// 		clients = make([]ClientRes, 0)
// 		c.JSON(http.StatusOK, clients)
// 	}

// 	for _, c := range h.hub.Rooms[roomId].Clients {
// 		clients = append(clients, ClientRes{
// 			UidUser: c.UidUser,
// 		})
// 	}

// 	c.JSON(http.StatusOK, clients)
// }
