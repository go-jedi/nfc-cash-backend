package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type TodoAuth interface {
	CreateUser(userForm appl_row.CreateUser) (string, int, error)
	GetUser(username string, password string) ([]appl_row.User, error)
	CheckEmailExist(userForm appl_row.CheckEmailExist) (bool, int, error)
	CheckUsernameExist(userForm appl_row.CheckUsernameExist) (bool, int, error)
}

type TodoVerify interface {
	CheckEmailVerify(uid string) (bool, int, error)
	EmailVerify(uid string) (bool, int, error)
}

type TodoRecovery interface {
	GetUserUidByEmail(userForm appl_row.RecoveryPasswordSendMessage) (string, int, error)
	CheckRecoveryPassword(uid string) (bool, int, error)
	LaunchRecoveryPassword(uid string) (int, error)
	CompleteRecoveryPassword(uid string) (int, error)
	RecoveryPasswordCompare(uid string, password string) (bool, int, error)
	RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error)
}

type Repository struct {
	TodoAuth
	TodoVerify
	TodoRecovery
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoAuth:     NewAuthPostgres(db),
		TodoVerify:   NewVerifyPostgres(db),
		TodoRecovery: NewRecoveryPostgres(db),
	}
}
