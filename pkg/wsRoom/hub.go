package wsRoom

import "fmt"

type Room struct {
	UidRoom string             `json:"uidRoom"`
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
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.UidUser]; !ok {
					r.Clients[cl.UidUser] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.UidUser]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						fmt.Println("cl.RoomID -->", cl.RoomID)
						fmt.Println("cl.UidUser -->", cl.UidUser)
						h.Broadcast <- &Message{
							Content: "user left the chat",
							RoomID:  cl.RoomID,
							UidUser: cl.UidUser,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.UidUser)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast: // здесь происходит отправка сообщения клиентам
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Clients {
					if cl.UidUser != m.UidUser {
						cl.Message <- m
					}
				}
			}
		}
	}
}
