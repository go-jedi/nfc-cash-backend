package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{
		db: db,
	}
}

func (r *MessagePostgres) CreateMessage(messageForm appl_row.CreateMessage) (bool, int, error) {
	var isCreateMessage bool
	messageFormJson, err := json.Marshal(messageForm)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации messageForm, %s", err)
	}
	err = r.db.QueryRow("SELECT message_create($1)", messageFormJson).Scan(&isCreateMessage)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции message_create из базы данных, %s", err)
	}
	return isCreateMessage, http.StatusOK, nil
}

func (r *MessagePostgres) GetRoomMessages(uidRoom string) ([]appl_row.Message, int, error) {
	var roomMessages []appl_row.Message
	var roomMessagesByte []byte
	err := r.db.QueryRow("SELECT messages_get_room($1)", uidRoom).Scan(&roomMessagesByte)
	if err != nil {
		return []appl_row.Message{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции messages_get_room из базы данных, %s", err)
	}
	err = json.Unmarshal(roomMessagesByte, &roomMessages)
	if err != nil {
		return []appl_row.Message{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetRoomMessages, %s", err)
	}
	return roomMessages, http.StatusOK, nil
}
