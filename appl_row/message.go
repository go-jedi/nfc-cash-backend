package appl_row

type Message struct {
	Id      int    `json:"id"`
	UidRoom string `json:"uid_room"`
	UidUser string `json:"uid_user"`
	Message string `json:"message"`
	Created string `json:"created"`
}

type CreateMessage struct {
	UidRoom string `json:"uidRoom"`
	UidUser string `json:"uidUser"`
	Message string `json:"message"`
}

type GetRoomMessages struct {
	UidRoom string `json:"uidRoom"`
}
