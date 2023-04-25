package handler

import (
	"github.com/gin-gonic/gin"
)

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
