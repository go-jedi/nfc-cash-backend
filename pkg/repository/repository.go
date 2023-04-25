package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/meetsite-backend/appl_row"
)

type TodoAuth interface {
	CreateUser(userForm appl_row.CreateUser) (string, int, error)
	GetUser(username string, password string) ([]appl_row.User, error)
}

type TodoVerify interface {
	EmailVerify(uid string) (bool, int, error)
}

type TodoRecovery interface {
	RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error)
	GetUserUidByEmail(userForm appl_row.RecoveryPasswordSendMessage) (string, int, error)
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
