package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rob-bender/meetsite-backend/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/rob-bender/meetsite-backend/docs"
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
		auth.POST("/sign-up", h.signUp) // Регистрация пользователя
		auth.POST("/sign-in", h.signIn) // Авторизация пользователя
	}

	email := router.Group("/verify")
	{
		email.GET("/emailver/:uid", h.emailVerify)
	}

	recovery := router.Group("/recovery")
	{
		recovery.POST("/recoveryPasswordSendMessage", h.recoveryPasswordSendMessage)
		recovery.POST("/recoveryPassword", h.recoveryPassword)
	}

	validate := router.Group("/validate")
	{
		validate.POST("/validateEmail", h.validateEmail)
		validate.POST("/validatePassword", h.validatePassword)
		validate.POST("/validateUsername", h.validateUsername)
	}

	api := router.Group("/api-v1", h.userIdentity)
	{

		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
		}
	}

	return router
}
