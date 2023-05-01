package handler

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nfc-cash-backend/pkg/service"
	"github.com/rob-bender/nfc-cash-backend/pkg/wsRoom"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/rob-bender/nfc-cash-backend/docs"
)

type Handler struct {
	hub      *wsRoom.Hub
	services *service.Service
}

func NewHandler(s *service.Service, h *wsRoom.Hub) *Handler { // создаём новый handler с полем services
	return &Handler{
		services: s,
		hub:      h,
	}
}

func (h *Handler) InitRoutes() *gin.Engine { // обработчик роутов, Создание роутов
	router := gin.New() // инициализация роутов

	// router.Use(cors.Default())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)                            // Регистрация пользователя
		auth.POST("/sign-in", h.signIn)                            // Авторизация пользователя
		auth.POST("/check-email-exist", h.checkEmailExist)         // Есть ли в базе данных зарегистрированный email
		auth.POST("/check-username-exist", h.checkUsernameExist)   // Есть ли в базе данных зарегистрированный username
		auth.POST("/check-confirm-account", h.checkConfirmAccount) // подтверждён ли аккаунт администратором
	}

	email := router.Group("/verify")
	{
		email.POST("/check-email-verify", h.checkEmailVerify) // проверка является ли почта подтвержденной
		email.GET("/emailver/:uid", h.emailVerify)            // переход по ссылки из почты для подтверждения почты
	}

	recovery := router.Group("/recovery")
	{
		recovery.POST("/recovery-password-send-message", h.recoveryPasswordSendMessage) // отправка пользователю на почту сообщения о восстановлении пароля
		recovery.POST("/check-recovery-password", h.checkRecoveryPassword)              // проверка запущен ли процесс восстановления пароля
		recovery.POST("/recovery-password-complete", h.completeRecoveryPassword)        // завершение восстановления пароля
		recovery.POST("/recovery-password-compare", h.recoveryPasswordCompare)          // сравнивание нового пароля и старого
		recovery.POST("/recovery-password", h.recoveryPassword)                         // изменение пароля пользователя на новый
	}

	validate := router.Group("/validate")
	{
		validate.POST("/validate-email", h.validateEmail)       // проверка почты на валидность
		validate.POST("/validate-password", h.validatePassword) // проверка пароль на валидность
		validate.POST("/validate-username", h.validateUsername) // проверка username на валидность
	}

	webSocket := router.Group("/room")
	{
		webSocket.GET("/join-room/:roomId", h.joinRoom) // вступить в нужную комнату
		// webSocket.GET("/get-rooms", h.getRooms)
		// webSocket.GET("/get-clients/:roomId", h.getClients)
		webSocket.POST("/create-room", h.createRoom) // создать комнату для чата
		webSocket.POST("/leave-room", h.leaveRoom)   // покинуть комнату
	}

	order := router.Group("/order")
	{
		order.POST("/create-order", h.orderCreate) // создание ордера (заказа)
	}

	message := router.Group("/message")
	{
		message.POST("/create-message", h.createMessage)    // создать сообщение
		message.POST("/get-all-messages", h.getAllMessages) // получить все сообщения в чате
	}

	api := router.Group("/api-v1", h.userIdentity)
	{
		validateToken := api.Group("/validate-token")
		{
			validateToken.GET("/", h.checkValidateToken) // проверка на валидный токен
		}
		user := api.Group("/user")
		{
			user.GET("/get-user-profile", h.getUserProfile) // получение профиля пользователя
			user.GET("/check-is-admin", h.checkIsAdmin)     // проверка пользователя на администратора
		}
		admin := api.Group("/admin")
		{
			admin.GET("/get-users-confirm", h.getUsersConfirm)        // получить пользователей, которые с подтвержденными аккаунтами (без супер администратора)
			admin.GET("/get-users-un-confirm", h.getUsersUnConfirm)   // получить пользователей не подтвержденными аккаунтами
			admin.POST("/user-confirm-account", h.userConfirmAccount) // подтверждение аккаунта пользователя администратором
		}
	}

	return router
}
