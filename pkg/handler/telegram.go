package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

// @Summary		BotCreate
// @Tags			telegram
// @Description	bot create
// @ID				bot-create
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.BotCreate	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/telegram/create-bot [post]
func (h *Handler) botCreate(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "идентификатор пользователя имеет недопустимый тип",
		})
		return
	}
	if userId > 0 {
		type Body struct {
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
		resBotCreate, statusCode, err := h.services.BotCreate(userId, appl_row.BotCreate(body))
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
}

// @Summary		BotDelete
// @Tags			telegram
// @Description	bot delete
// @ID				bot-delete
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.BotDelete	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/telegram/delete-bot [post]
func (h *Handler) botDelete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "идентификатор пользователя имеет недопустимый тип",
		})
		return
	}
	if userId > 0 {
		type Body struct {
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
		resBotDelete, statusCode, err := h.services.BotDelete(userId, appl_row.BotDelete(body))
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
}

// @Summary		TurnOnBot
// @Tags			telegram
// @Description	turn on bot
// @ID				turn-on-bot
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.TurnOnBot	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/telegram/turn-on-bot [post]
func (h *Handler) turnOnBot(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "идентификатор пользователя имеет недопустимый тип",
		})
		return
	}
	if userId > 0 {
		type Body struct {
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
		resTurnOnBot, statusCode, err := h.services.TurnOnBot(userId, body.Token)
		if err != nil {
			c.JSON(statusCode, map[string]interface{}{
				"status":  statusCode,
				"message": err.Error(),
				"result":  resTurnOnBot,
			})
			return
		}
		if resTurnOnBot {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное включение бота",
				"result":  resTurnOnBot,
			})
		}
		if !resTurnOnBot {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "ошибка включения бота",
				"result":  resTurnOnBot,
			})
		}
	}
}

// @Summary		SwitchOffBot
// @Tags			telegram
// @Description	switch off bot
// @ID				switch-off-bot
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.SwitchOffBot	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/telegram/switch-off-bot [post]
func (h *Handler) switchOffBot(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "идентификатор пользователя имеет недопустимый тип",
		})
		return
	}
	if userId > 0 {
		type Body struct {
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
		resSwitchOffBot, statusCode, err := h.services.SwitchOffBot(userId, body.Token)
		if err != nil {
			c.JSON(statusCode, map[string]interface{}{
				"status":  statusCode,
				"message": err.Error(),
				"result":  resSwitchOffBot,
			})
			return
		}
		if resSwitchOffBot {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное выключение бота",
				"result":  resSwitchOffBot,
			})
		}
		if !resSwitchOffBot {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "ошибка выключения бота",
				"result":  resSwitchOffBot,
			})
		}
	}
}

// @Summary		GetBots
// @Tags			telegram
// @Description	get bots
// @ID				get-bots
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/telegram/get-bots [get]
func (h *Handler) getBots(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "идентификатор пользователя имеет недопустимый тип",
		})
		return
	}
	if userId > 0 {
		resGetBots, statusCode, err := h.services.GetBots(userId)
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
}
