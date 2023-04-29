package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) GetUserProfile(id int) ([]appl_row.UserProfile, int, error) {
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

func (r *UserPostgres) CheckIsAdmin(id int) (bool, int, error) {
	var isAdmin bool
	err := r.db.QueryRow("SELECT user_check_is_admin($1)", id).Scan(&isAdmin)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_is_admin из базы данных, %s", err)
	}
	return isAdmin, http.StatusOK, nil
}
