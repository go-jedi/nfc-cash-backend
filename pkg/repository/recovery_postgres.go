package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type RecoveryPostgres struct {
	db *sqlx.DB
}

func NewRecoveryPostgres(db *sqlx.DB) *RecoveryPostgres {
	return &RecoveryPostgres{
		db: db,
	}
}

func (r *RecoveryPostgres) GetUserUidByEmail(userForm appl_row.RecoveryPasswordSendMessage) (string, int, error) {
	var uid string
	userFormJson, err := json.Marshal(userForm)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userForm, %s", err)
	}
	err = r.db.QueryRow("SELECT user_get_uid_by_email($1)", userFormJson).Scan(&uid)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_uid_by_email из базы данных, %s", err)
	}
	return uid, http.StatusOK, nil
}

func (r *RecoveryPostgres) CheckRecoveryPassword(uid string) (bool, int, error) {
	var isRecoveryPassword bool
	err := r.db.QueryRow("SELECT user_check_recovery_password($1)", uid).Scan(&isRecoveryPassword)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_recovery_password из базы данных, %s", err)
	}
	return isRecoveryPassword, http.StatusOK, nil
}

func (r *RecoveryPostgres) LaunchRecoveryPassword(uid string) (int, error) {
	_, err := r.db.Exec("SELECT user_launch_recovery_password($1)", uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_launch_recovery_password из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *RecoveryPostgres) CompleteRecoveryPassword(uid string) (int, error) {
	_, err := r.db.Exec("SELECT user_complete_recovery_password($1)", uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_complete_recovery_password из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *RecoveryPostgres) RecoveryPasswordCompare(uid string, password string) (bool, int, error) {
	var isComparePasswords bool
	err := r.db.QueryRow("SELECT user_compare_password($1, $2)", uid, password).Scan(&isComparePasswords)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_compare_password из базы данных, %s", err)
	}
	return isComparePasswords, http.StatusOK, nil
}

func (r *RecoveryPostgres) RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error) {
	userFormJson, err := json.Marshal(userForm)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации userForm, %s", err)
	}
	_, err = r.db.Exec("SELECT user_recovery_password($1)", userFormJson)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_recovery_password из базы данных, %s", err)
	}
	return true, http.StatusOK, nil
}
