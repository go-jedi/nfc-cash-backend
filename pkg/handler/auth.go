package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/meetsite-backend/appl_row"
)

// @Summary		SignUp
// @Tags			auth
// @Description	create account
// @ID				create-account
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.CreateUser	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) { // Регистрация пользователя
	type Body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
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
	statusCode, err := h.services.CreateUser(appl_row.CreateUser(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная регистрация пользователя",
	})
}

// @Summary		SignIn
// @Tags			auth
// @Description	login
// @ID				login
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.AuthUser	true	"credentials"
// @Success		200		{string}	string				"token"
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) { // Авторизация пользователя
	type Body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	token, err := h.services.GenerateToken(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "успешная авторизация пользователя",
		"token":   token,
	})
}
