package repository

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type VerifyPostgres struct {
	db *sqlx.DB
}

func NewVerifyPostgres(db *sqlx.DB) *VerifyPostgres {
	return &VerifyPostgres{
		db: db,
	}
}

func (r *VerifyPostgres) CheckEmailVerify(uid string) (bool, int, error) {
	var isEmailVerify bool
	err := r.db.QueryRow("SELECT user_check_email_verify($1)", uid).Scan(&isEmailVerify)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_email_verify из базы данных, %s", err)
	}
	return isEmailVerify, http.StatusOK, nil
}

func (r *VerifyPostgres) EmailVerify(uid string) (bool, int, error) {
	_, err := r.db.Exec("SELECT user_verify_email($1)", uid)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_verify_email из базы данных, %s", err)
	}
	return true, http.StatusOK, nil
}
