package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

// @Summary		RecoveryPasswordSendMessage
// @Tags			recovery
// @Description	recovery passwordSendMessage
// @ID				recovery-passwordSendMessage
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.RecoveryPasswordSendMessage	true	"account info"
// @Success		200		{integer}	integer					1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/recovery/recoveryPasswordSendMessage [post]
func (h *Handler) recoveryPasswordSendMessage(c *gin.Context) { // отправка письма на почту пользователя для восстановления пароля
	type Body struct {
		Email string `json:"email"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resRecoveryPasswordSendMessage, statusCode, err := h.services.RecoveryPasswordSendMessage(appl_row.RecoveryPasswordSendMessage(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resRecoveryPasswordSendMessage,
		})
		return
	}
	if resRecoveryPasswordSendMessage {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешная отправка письма для восстановления пароля",
			"result":  resRecoveryPasswordSendMessage,
		})
	}
	if !resRecoveryPasswordSendMessage {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка отправки письма для восстановления пароля",
			"result":  resRecoveryPasswordSendMessage,
		})
	}
}

// @Summary		RecoveryPassword
// @Tags			recovery
// @Description	recovery password
// @ID				recovery-password
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.RecoveryPassword	true	"account info"
// @Success		200		{integer}	integer					1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/recovery/recoveryPassword [post]
func (h *Handler) recoveryPassword(c *gin.Context) { // изменение пароля
	type Body struct {
		Uid      string `json:"uid"`
		Password string `json:"password"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resRecoveryPassword, statusCode, err := h.services.RecoveryPassword(appl_row.RecoveryPassword(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resRecoveryPassword,
		})
		return
	}
	if resRecoveryPassword {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешное изменение пароля",
			"result":  resRecoveryPassword,
		})
	}
	if !resRecoveryPassword {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка изменения пароля",
			"result":  resRecoveryPassword,
		})
	}
}
