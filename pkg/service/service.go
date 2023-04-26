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
}

type TodoVerify interface {
	CheckEmailVerify(uid string) (bool, int, error)
	EmailVerify(uid string) (bool, int, error)
}

type TodoRecovery interface {
	RecoveryPasswordSendMessage(userForm appl_row.RecoveryPasswordSendMessage) (bool, int, error)
	RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error)
}

type Service struct {
	TodoAuth
	TodoVerify
	TodoRecovery
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoAuth:     NewAuthService(r.TodoAuth),
		TodoVerify:   NewVerifyService(r.TodoVerify),
		TodoRecovery: NewRecoveryService(r.TodoRecovery),
	}
}
