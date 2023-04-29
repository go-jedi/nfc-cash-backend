package service

import (
	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
)

type TodoAuth interface {
	CreateUser(userForm appl_row.CreateUser) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
	CheckEmailExist(userForm appl_row.CheckEmailExist) (bool, int, error)
	CheckUsernameExist(userForm appl_row.CheckUsernameExist) (bool, int, error)
	CheckConfirmAccount(userForm appl_row.CheckConfirmAccount) (bool, int, error)
}

type TodoVerify interface {
	CheckEmailVerify(uid string) (bool, int, error)
	EmailVerify(uid string) (bool, int, error)
}

type TodoRecovery interface {
	RecoveryPasswordSendMessage(userForm appl_row.RecoveryPasswordSendMessage) (bool, int, error)
	CheckRecoveryPassword(uid string) (bool, int, error)
	CompleteRecoveryPassword(uid string) (int, error)
	RecoveryPasswordCompare(uid string, password string) (bool, int, error)
	RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error)
}

type TodoUser interface {
	GetUserProfile(id int) ([]appl_row.UserProfile, int, error)
	CheckIsAdmin(id int) (bool, int, error)
}

type TodoAdmin interface {
	GetUsersConfirm(id int) ([]appl_row.GetUsersConfirm, int, error)
	GetUsersUnConfirm(id int) ([]appl_row.GetUsersUnConfirm, int, error)
	UserConfirmAccount(id int, userForm appl_row.UserConfirmAccount) (bool, int, error)
}

type Service struct {
	TodoAuth
	TodoVerify
	TodoRecovery
	TodoUser
	TodoAdmin
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoAuth:     NewAuthService(r.TodoAuth),
		TodoVerify:   NewVerifyService(r.TodoVerify),
		TodoRecovery: NewRecoveryService(r.TodoRecovery),
		TodoUser:     NewUserService(r.TodoUser),
		TodoAdmin:    NewAdminService(r.TodoAdmin),
	}
}
