package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

// @Summary		GetUsersUnConfirm
// @Tags			admin
// @Description	get users un confirm
// @ID				get-users-un-confirm
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/admin/get-users-un-confirm [get]
func (h *Handler) getUsersUnConfirm(c *gin.Context) {
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
		resGetUsersUnConfirm, statusCode, err := h.services.GetUsersUnConfirm(userId)
		if err != nil {
			c.JSON(statusCode, map[string]interface{}{
				"status":  statusCode,
				"message": err.Error(),
			})
			return
		}
		if len(resGetUsersUnConfirm) > 0 {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное получение не подтвержденных пользователей",
				"result":  resGetUsersUnConfirm,
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное получение не подтвержденных пользователей",
				"result":  resGetUsersUnConfirm,
			})
		}
	}
}

// @Summary		UserConfirmAccount
// @Tags			admin
// @Description	user confirm account
// @ID				user-confirm-account
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.UserConfirmAccount	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/admin/user-confirm-account [post]
func (h *Handler) userConfirmAccount(c *gin.Context) {
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
			Id int `json:"id"`
		}
		var body Body
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "некорректно переданы данные в body",
			})
			return
		}
		resUserConfirmAccount, statusCode, err := h.services.UserConfirmAccount(userId, appl_row.UserConfirmAccount(body))
		if err != nil {
			c.JSON(statusCode, map[string]interface{}{
				"status":  statusCode,
				"message": err.Error(),
				"result":  resUserConfirmAccount,
			})
			return
		}
		if resUserConfirmAccount {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное подтверждение аккаунта пользователя",
				"result":  resUserConfirmAccount,
			})
		}
		if !resUserConfirmAccount {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "ошибка подтверждения аккаунта пользователя",
				"result":  resUserConfirmAccount,
			})
		}
	}
}
