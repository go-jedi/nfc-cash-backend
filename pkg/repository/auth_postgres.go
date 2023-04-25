package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(userForm appl_row.CreateUser) (string, int, error) {
	var uid string
	userFormJson, err := json.Marshal(userForm)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userForm, %s", err)
	}
	err = r.db.QueryRow("SELECT uid($1)", 8).Scan(&uid)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
	}
	_, err = r.db.Exec("SELECT user_create($1, $2)", userFormJson, uid)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_create из базы данных, %s", err)
	}
	return uid, http.StatusOK, nil
}

func (r *AuthPostgres) GetUser(username string, password string) ([]appl_row.User, error) {
	var user []appl_row.User
	var userByte []byte
	err := r.db.QueryRow("SELECT user_get_data($1, $2)", username, password).Scan(&userByte)
	if err != nil {
		return []appl_row.User{}, fmt.Errorf("ошибка выполнения функции user_get_data из базы данных, %s", err)
	}
	err = json.Unmarshal(userByte, &user)
	if err != nil {
		return []appl_row.User{}, fmt.Errorf("ошибка конвертации в функции GetUser, %s", err)
	}
	return user, nil
}

func (r *AuthPostgres) CheckEmailExist(userForm appl_row.CheckEmailExist) (bool, int, error) {
	var isEmailExist bool
	userFormJson, err := json.Marshal(userForm)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userForm, %s", err)
	}
	err = r.db.QueryRow("SELECT user_check_exist_email($1)", userFormJson).Scan(&isEmailExist)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_exist_email из базы данных, %s", err)
	}
	return isEmailExist, http.StatusOK, nil
}

func (r *AuthPostgres) CheckUsernameExist(userForm appl_row.CheckUsernameExist) (bool, int, error) {
	var isUsernameExist bool
	userFormJson, err := json.Marshal(userForm)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userForm, %s", err)
	}
	err = r.db.QueryRow("SELECT user_check_exist_username($1)", userFormJson).Scan(&isUsernameExist)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_exist_username из базы данных, %s", err)
	}
	return isUsernameExist, http.StatusOK, nil
}
