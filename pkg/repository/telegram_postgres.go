package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type TelegramPostgres struct {
	db *sqlx.DB
}

func NewTelegramPostgres(db *sqlx.DB) *TelegramPostgres {
	return &TelegramPostgres{
		db: db,
	}
}

func (r *TelegramPostgres) BotCreate(id int, botForm appl_row.BotCreate) (bool, int, error) {
	var isBotCreate bool
	botFormJson, err := json.Marshal(botForm)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации botForm, %s", err)
	}
	err = r.db.QueryRow("SELECT bot_create($1, $2)", id, botFormJson).Scan(&isBotCreate)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции bot_create из базы данных, %s", err)
	}
	return isBotCreate, http.StatusOK, nil
}

func (r *TelegramPostgres) BotDelete(id int, botForm appl_row.BotDelete) (bool, int, error) {
	var isBotDelete bool
	botFormJson, err := json.Marshal(botForm)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации botForm, %s", err)
	}
	err = r.db.QueryRow("SELECT bot_delete($1, $2)", id, botFormJson).Scan(&isBotDelete)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции bot_delete из базы данных, %s", err)
	}
	return isBotDelete, http.StatusOK, nil
}

func (r *TelegramPostgres) GetBots(id int) ([]appl_row.Bot, int, error) {
	var bots []appl_row.Bot
	var botsByte []byte
	err := r.db.QueryRow("SELECT bots_get($1)", id).Scan(&botsByte)
	if err != nil {
		return []appl_row.Bot{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции bots_get из базы данных, %s", err)
	}
	err = json.Unmarshal(botsByte, &bots)
	if err != nil {
		return []appl_row.Bot{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetBots, %s", err)
	}
	return bots, http.StatusOK, nil
}
