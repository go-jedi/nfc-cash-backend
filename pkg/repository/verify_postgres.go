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

func (r *VerifyPostgres) EmailVerify(uid string) (bool, int, error) {
	_, err := r.db.Exec("SELECT user_verify_email($1)", uid)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_verify_email из базы данных, %s", err)
	}
	return true, http.StatusOK, nil
}
