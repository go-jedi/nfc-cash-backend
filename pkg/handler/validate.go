package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/pkg/validate"
)

// @Summary		ValidateEmail
// @Tags			validate
// @Description	validate email
// @ID				validate-email
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.ValidateEmail	true	"account info"
// @Success		200		{integer}	integer					1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/validate/validate-email [post]
func (h *Handler) validateEmail(c *gin.Context) { // Валидация электронной почты
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
	resValidateEmail, statusCode, err := validate.ValidateEmail(body.Email)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resValidateEmail,
		})
		return
	}
	if resValidateEmail {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешная валидация электронной почты",
			"result":  resValidateEmail,
		})
	}
	if !resValidateEmail {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка валидации электронной почты",
			"result":  resValidateEmail,
		})
	}
}

// @Summary		ValidatePassword
// @Tags			validate
// @Description	validate password
// @ID				validate-password
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.ValidatePassword	true	"account info"
// @Success		200		{integer}	integer					1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/validate/validate-password [post]
func (h *Handler) validatePassword(c *gin.Context) { // Валидация пароля
	type Body struct {
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
	resValidatePassword, statusCode, err := validate.ValidatePassword(body.Password)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resValidatePassword,
		})
		return
	}
	if resValidatePassword {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешная валидация пароля",
			"result":  resValidatePassword,
		})
	}
	if !resValidatePassword {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка валидации пароля",
			"result":  resValidatePassword,
		})
	}
}

// @Summary		ValidateUsername
// @Tags			validate
// @Description	validate username
// @ID				validate-username
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.ValidateUsername	true	"account info"
// @Success		200		{integer}	integer					1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/validate/validate-username [post]
func (h *Handler) validateUsername(c *gin.Context) { // Валидация username
	type Body struct {
		Username string `json:"username"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resValidateUsername, statusCode, err := validate.ValidateUsername(body.Username)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resValidateUsername,
		})
		return
	}
	if resValidateUsername {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "успешная валидация username",
			"result":  resValidateUsername,
		})
	}
	if !resValidateUsername {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "ошибка валидации username",
			"result":  resValidateUsername,
		})
	}
}
