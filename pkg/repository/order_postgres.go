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

func (r *OrderPostgres) GetOrder(uidOrder string) ([]appl_row.GetOrder, int, error) {
	var order []appl_row.GetOrder
	var orderByte []byte
	err := r.db.QueryRow("SELECT order_get($1)", uidOrder).Scan(&orderByte)
	if err != nil {
		return []appl_row.GetOrder{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции order_get из базы данных, %s", err)
	}
	err = json.Unmarshal(orderByte, &order)
	if err != nil {
		return []appl_row.GetOrder{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUsersConfirm, %s", err)
	}
	return order, http.StatusOK, nil
}
