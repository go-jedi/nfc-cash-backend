package repository

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type IpPostgres struct {
	db *sqlx.DB
}

func NewIpPostgres(db *sqlx.DB) *IpPostgres {
	return &IpPostgres{
		db: db,
	}
}

func (r *IpPostgres) BlockIp(ipAddress string) (bool, int, error) {
	var isBlockIp bool
	err := r.db.QueryRow("SELECT ip_block($1)", ipAddress).Scan(&isBlockIp)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции ip_block из базы данных, %s", err)
	}
	return isBlockIp, http.StatusOK, nil
}

func (r *IpPostgres) CheckIpBlock(ipAddress string) (bool, int, error) {
	var isBlockIp bool
	err := r.db.QueryRow("SELECT ip_check_block($1)", ipAddress).Scan(&isBlockIp)
	if err != nil {
		return true, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции ip_check_block из базы данных, %s", err)
	}
	return isBlockIp, http.StatusOK, nil
}
