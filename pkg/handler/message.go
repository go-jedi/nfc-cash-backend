package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

// @Summary		CreateMessage
// @Tags			message
// @Description	create message
// @ID				create-message
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.CreateMessage	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/message/create-message [post]
func (h *Handler) createMessage(c *gin.Context) {
	type Body struct {
		UidRoom string `json:"uidRoom"`
		UidUser string `json:"uidUser"`
		Message string `json:"message"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resCreateMessage, statusCode, err := h.services.CreateMessage(appl_row.CreateMessage(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resCreateMessage,
		})
		return
	}
	if resCreateMessage {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное создание сообщения",
			"result":  resCreateMessage,
		})
	}
	if !resCreateMessage {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "ошибка создания сообщения",
			"result":  resCreateMessage,
		})
	}
}

// @Summary		GetRoomMessages
// @Tags			message
// @Description	get room messages
// @ID				get-room-messages
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.GetRoomMessages	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/message/get-room-messages [post]
func (h *Handler) getRoomMessages(c *gin.Context) {
	type Body struct {
		UidRoom string `json:"uidRoom"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resGetRoomMessages, statusCode, err := h.services.GetRoomMessages(body.UidRoom)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if len(resGetRoomMessages) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение сообщений чата",
			"result":  resGetRoomMessages,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение сообщений чата",
			"result":  resGetRoomMessages,
		})
	}
}
