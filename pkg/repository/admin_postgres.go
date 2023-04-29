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

func (r *AdminPostgres) GetUsersConfirm(id int) ([]appl_row.GetUsersConfirm, int, error) {
	var usersConfirm []appl_row.GetUsersConfirm
	var usersConfirmByte []byte
	err := r.db.QueryRow("SELECT admin_get_users_confirm($1)", id).Scan(&usersConfirmByte)
	if err != nil {
		return []appl_row.GetUsersConfirm{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_get_users_confirm из базы данных, %s", err)
	}
	err = json.Unmarshal(usersConfirmByte, &usersConfirm)
	if err != nil {
		return []appl_row.GetUsersConfirm{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUsersConfirm, %s", err)
	}
	return usersConfirm, http.StatusOK, nil
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
		return []appl_row.GetUsersUnConfirm{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUsersUnConfirm, %s", err)
	}
	return usersUnConfirm, http.StatusOK, nil
}

func (r *AdminPostgres) GetUserProfile(id int) ([]appl_row.UserProfile, int, error) {
	var userProfile []appl_row.UserProfile
	var userProfileByte []byte
	err := r.db.QueryRow("SELECT user_get_profile($1)", id).Scan(&userProfileByte)
	if err != nil {
		return []appl_row.UserProfile{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_profile из базы данных, %s", err)
	}
	err = json.Unmarshal(userProfileByte, &userProfile)
	if err != nil {
		return []appl_row.UserProfile{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserProfile, %s", err)
	}
	return userProfile, http.StatusOK, nil
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
