package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/rob-bender/nfc-cash-backend/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler { // создаём новый handler с полем services
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine { // обработчик роутов, Создание роутов
	router := gin.New() // инициализация роутов

	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)                          // Регистрация пользователя
		auth.POST("/sign-in", h.signIn)                          // Авторизация пользователя
		auth.POST("/check-email-exist", h.checkEmailExist)       // Есть ли в базе данных зарегистрированный email
		auth.POST("/check-username-exist", h.checkUsernameExist) // Есть ли в базе данных зарегистрированный username
	}

	email := router.Group("/verify")
	{
		email.POST("/check-email-verify", h.checkEmailVerify)
		email.GET("/emailver/:uid", h.emailVerify)
	}

	recovery := router.Group("/recovery")
	{
		recovery.POST("/recovery-password-send-message", h.recoveryPasswordSendMessage)
		recovery.POST("/check-recovery-password", h.checkRecoveryPassword)
		recovery.POST("/recovery-password-complete", h.completeRecoveryPassword)
		recovery.POST("/recovery-password-compare", h.recoveryPasswordCompare)
		recovery.POST("/recovery-password", h.recoveryPassword)
	}

	validate := router.Group("/validate")
	{
		validate.POST("/validate-email", h.validateEmail)
		validate.POST("/validate-password", h.validatePassword)
		validate.POST("/validate-username", h.validateUsername)
	}

	api := router.Group("/api-v1", h.userIdentity)
	{

		validateToken := api.Group("/validate-token")
		{
			validateToken.GET("/", h.checkValidateToken)
		}
		users := api.Group("/user")
		{
			users.GET("/get-user-profile", h.getUserProfile)
		}
	}

	return router
}
