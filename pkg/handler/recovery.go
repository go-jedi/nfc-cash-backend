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
// @Router			/recovery/recovery-password-send-message [post]
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

func (h *Handler) checkRecoveryPassword(c *gin.Context) {
	type Body struct {
		Uid string `json:"uid"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resCheckRecoveryPassword, statusCode, err := h.services.CheckRecoveryPassword(body.Uid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resCheckRecoveryPassword,
		})
		return
	}
	if resCheckRecoveryPassword {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "запуск восстановления пароля запущен",
			"result":  resCheckRecoveryPassword,
		})
	}
	if !resCheckRecoveryPassword {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "запуск восстановления пароля не запущен",
			"result":  resCheckRecoveryPassword,
		})
	}
}

func (h *Handler) completeRecoveryPassword(c *gin.Context) {
	type Body struct {
		Uid string `json:"uid"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.CompleteRecoveryPassword(body.Uid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное завершение восстановления пароля",
	})
}

// @Summary		RecoveryPasswordCompare
// @Tags			recovery
// @Description	recovery password compare
// @ID				recovery-password-compare
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.RecoveryPasswordCompare	true	"account info"
// @Success		200		{integer}	integer					1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/recovery-password-compare [post]
func (h *Handler) recoveryPasswordCompare(c *gin.Context) {
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
	resRecoveryPasswordCompare, statusCode, err := h.services.RecoveryPasswordCompare(body.Uid, body.Password)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resRecoveryPasswordCompare,
		})
		return
	}
	if resRecoveryPasswordCompare {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешное сравнение паролей",
			"result":  resRecoveryPasswordCompare,
		})
	}
	if !resRecoveryPasswordCompare {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка сравнения паролей",
			"result":  resRecoveryPasswordCompare,
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
// @Router			/recovery/recovery-password [post]
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
