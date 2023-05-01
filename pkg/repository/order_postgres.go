package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/binCard"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

func (r *OrderPostgres) OrderCreate(orderForm appl_row.OrderCreate) (bool, int, error) {
	var isOrderCreate bool
	resCheckBin, err := binCard.CheckBin(orderForm.CardNumber)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции CheckBin, %s", err)
	}
	orderFormJson, err := json.Marshal(orderForm)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации orderForm, %s", err)
	}
	resCheckBinJson, err := json.Marshal(resCheckBin)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации resCheckBin, %s", err)
	}
	err = r.db.QueryRow("SELECT order_create($1, $2)", orderFormJson, resCheckBinJson).Scan(&isOrderCreate)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
	}
	return isOrderCreate, http.StatusOK, nil
}
