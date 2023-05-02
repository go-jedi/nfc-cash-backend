package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type RoomPostgres struct {
	db *sqlx.DB
}

func NewRoomPostgres(db *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{
		db: db,
	}
}

func (r *RoomPostgres) CreateRoom() (string, int, error) {
	var uidRoom string
	err := r.db.QueryRow("SELECT room_uid($1)", 18).Scan(&uidRoom)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции room_uid из базы данных, %s", err)
	}
	_, err = r.db.Exec("SELECT room_create($1)", uidRoom)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции room_create из базы данных, %s", err)
	}
	return uidRoom, http.StatusOK, nil
}

func (r *RoomPostgres) JoinRoom(uidRoom string, uidUser string) (string, int, error) {
	var userUidCheck string = uidUser
	if userUidCheck == "none" {
		err := r.db.QueryRow("SELECT room_user_uid($1)", 20).Scan(&userUidCheck)
		if err != nil {
			return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции room_user_uid из базы данных, %s", err)
		}
	}
	_, err := r.db.Exec("SELECT room_join($1, $2)", uidRoom, userUidCheck)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции room_join из базы данных, %s", err)
	}
	return userUidCheck, http.StatusOK, nil
}

func (r *RoomPostgres) LeaveRoom(uidRoom string, uidUser string) (int, error) {
	_, err := r.db.Exec("SELECT room_leave($1, $2)", uidRoom, uidUser)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции room_leave из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *RoomPostgres) GetRoom(uidRoom string) ([]appl_row.Room, int, error) {
	var room []appl_row.Room
	var roomByte []byte
	err := r.db.QueryRow("SELECT room_get($1)", uidRoom).Scan(&roomByte)
	if err != nil {
		return []appl_row.Room{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции room_get из базы данных, %s", err)
	}
	err = json.Unmarshal(roomByte, &room)
	if err != nil {
		return []appl_row.Room{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetRoom, %s", err)
	}
	return room, http.StatusOK, nil
}
