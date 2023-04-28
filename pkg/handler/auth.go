package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
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

// @Summary		CheckEmailExist
// @Tags			auth
// @Description	check email exist
// @ID				check-email-exist
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.CheckEmailExist	true	"credentials"
// @Success		200		{string}	string				"res"
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/auth/check-email-exist [post]
func (h *Handler) checkEmailExist(c *gin.Context) { // Есть ли в базе данных зарегистрированный email
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
	resCheckEmailExist, statusCode, err := h.services.CheckEmailExist(appl_row.CheckEmailExist(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if resCheckEmailExist {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "пользователь с такой электронной почтой уже существует",
			"result":  resCheckEmailExist,
		})
	}
	if !resCheckEmailExist {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "пользователь с такой электронной почтой не существует",
			"result":  resCheckEmailExist,
		})
	}
}

// @Summary		CheckUsernameExist
// @Tags			auth
// @Description	check username exist
// @ID				check-username-exist
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.CheckUsernameExist	true	"credentials"
// @Success		200		{string}	string				"res"
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/auth/check-username-exist [post]
func (h *Handler) checkUsernameExist(c *gin.Context) { // Есть ли в базе данных зарегистрированный username
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
	resCheckUsernameExist, statusCode, err := h.services.CheckUsernameExist(appl_row.CheckUsernameExist(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if resCheckUsernameExist {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "пользователь с таким username уже существует",
			"result":  resCheckUsernameExist,
		})
	}
	if !resCheckUsernameExist {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "пользователь с таким username не существует",
			"result":  resCheckUsernameExist,
		})
	}
}

// @Summary		CheckConfirmAccount
// @Tags			auth
// @Description	check confirm account
// @ID				check-confirm-account
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.CheckConfirmAccount	true	"credentials"
// @Success		200		{string}	string				"res"
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/auth/check-confirm-account [post]
func (h *Handler) checkConfirmAccount(c *gin.Context) {
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
	resCheckConfirmAccount, statusCode, err := h.services.CheckConfirmAccount(appl_row.CheckConfirmAccount(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if resCheckConfirmAccount {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "аккаунт пользователя успешно подтверждён администратором",
			"result":  resCheckConfirmAccount,
		})
	}
	if !resCheckConfirmAccount {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "аккаунт пользователя ещё не подтверждён администратором",
			"result":  resCheckConfirmAccount,
		})
	}
}
