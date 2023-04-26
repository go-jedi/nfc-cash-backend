package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		CheckEmailVerify
// @Tags			checkEmailVerify
// @Description	check email verify
// @ID				check-email-verify
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/verify/checkEmailVerify [post]
func (h *Handler) checkEmailVerify(c *gin.Context) { // проверка на верификацию электронной почты
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
	resCheckEmailVerify, statusCode, err := h.services.CheckEmailVerify(body.Uid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resCheckEmailVerify,
		})
		return
	}
	if resCheckEmailVerify {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "электронная почта подтверждена",
			"result":  resCheckEmailVerify,
		})
	}
	if !resCheckEmailVerify {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "электронная почта не подтверждена",
			"result":  resCheckEmailVerify,
		})
	}
}

// @Summary		VerifyEmail
// @Tags			verify
// @Description	verify email
// @ID				verify-email
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/verify/emailver/:uid [get]
func (h *Handler) emailVerify(c *gin.Context) { // верификация электронной почты
	uid := c.Param("uid")
	resEmailVerify, statusCode, err := h.services.EmailVerify(uid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resEmailVerify,
		})
		return
	}
	if resEmailVerify {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешное подтверждение электронной почты",
			"result":  resEmailVerify,
		})
	}
	if !resEmailVerify {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка подтверждения электронной почты",
			"result":  resEmailVerify,
		})
	}
}
