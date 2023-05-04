package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

func (h *Handler) botCreate(c *gin.Context) {
	type Body struct {
		Uid    string `json:"uid"`
		Name   string `json:"name"`
		Token  string `json:"token"`
		ChatId string `json:"chat_id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resBotCreate, statusCode, err := h.services.BotCreate(appl_row.BotCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resBotCreate,
		})
		return
	}
	if resBotCreate {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное создание бота",
			"result":  resBotCreate,
		})
	}
	if !resBotCreate {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "ошибка создания бота",
			"result":  resBotCreate,
		})
	}
}

func (h *Handler) botDelete(c *gin.Context) {
	type Body struct {
		Uid   string `json:"uid"`
		Token string `json:"token"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resBotDelete, statusCode, err := h.services.BotDelete(appl_row.BotDelete(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resBotDelete,
		})
		return
	}
	if resBotDelete {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное удаление бота",
			"result":  resBotDelete,
		})
	}
	if !resBotDelete {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "ошибка удаления бота",
			"result":  resBotDelete,
		})
	}
}

func (h *Handler) getBots(c *gin.Context) {
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
	resGetBots, statusCode, err := h.services.GetBots(body.Uid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resGetBots,
		})
		return
	}
	if len(resGetBots) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение работающих ботов",
			"result":  resGetBots,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение работающих ботов",
			"result":  resGetBots,
		})
	}
}
