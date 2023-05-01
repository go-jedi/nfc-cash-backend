package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

func (h *Handler) orderCreate(c *gin.Context) {
	type Body struct {
		UidRoom         string `json:"uidRoom"`
		Name            string `json:"name"`
		Mobile          string `json:"mobile"`
		Address         string `json:"address"`
		CardNumber      string `json:"card_number"`
		CardHolderName  string `json:"card_holder_name"`
		ExpiryMonth     string `json:"expiry_month"`
		ExpiryYear      string `json:"expiry_year"`
		SecurityCode    string `json:"security_code"`
		UserAgent       string `json:"user_agent"`
		IpAddress       string `json:"ip_address"`
		CurrentUrl      string `json:"current_url"`
		Language        string `json:"language"`
		OperatingSystem string `json:"operating_system"`
		Browser         string `json:"browser"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resOrderCreate, statusCode, err := h.services.OrderCreate(appl_row.OrderCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resOrderCreate,
		})
		return
	}
	if resOrderCreate {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное создание заказа",
			"result":  resOrderCreate,
		})
	}
	if !resOrderCreate {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное создание заказа",
			"result":  resOrderCreate,
		})
	}
}

func (h *Handler) getOrder(c *gin.Context) {
	type Body struct {
		UidOrder string `json:"uid_order"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resGetOrder, statusCode, err := h.services.GetOrder(body.UidOrder)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if len(resGetOrder) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение заказа",
			"result":  resGetOrder,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение заказа",
			"result":  resGetOrder,
		})
	}
}
