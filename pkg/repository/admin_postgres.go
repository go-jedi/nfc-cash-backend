package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{
		db: db,
	}
}

func (r *AdminPostgres) GetUsersUnConfirm(id int) ([]appl_row.GetUsersUnConfirm, int, error) {
	var usersUnConfirm []appl_row.GetUsersUnConfirm
	var usersUnConfirmByte []byte
	err := r.db.QueryRow("SELECT admin_get_users_unconfirm($1)", id).Scan(&usersUnConfirmByte)
	if err != nil {
		return []appl_row.GetUsersUnConfirm{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_get_users_unconfirm из базы данных, %s", err)
	}
	err = json.Unmarshal(usersUnConfirmByte, &usersUnConfirm)
	if err != nil {
		return []appl_row.GetUsersUnConfirm{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserProfile, %s", err)
	}
	return usersUnConfirm, http.StatusOK, nil
}

func (r *AdminPostgres) UserConfirmAccount(id int, userForm appl_row.UserConfirmAccount) (bool, int, error) {
	var isUserConfirmAccount bool
	userFormJson, err := json.Marshal(userForm)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userForm, %s", err)
	}
	err = r.db.QueryRow("SELECT admin_user_confirm_account($1, $2)", id, userFormJson).Scan(&isUserConfirmAccount)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_user_confirm_account из базы данных, %s", err)
	}
	return isUserConfirmAccount, http.StatusOK, nil
}
