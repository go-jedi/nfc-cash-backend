package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) orderCreate(c *gin.Context) {
	type Body struct {
		NameCard        string `json:"name_card"`
		MobileCard      string `json:"mobile_card"`
		AddressCard     string `json:"address_card"`
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
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
}
