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
	CheckConfirmAccount(userForm appl_row.CheckConfirmAccount) (bool, int, error)
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

type TodoRoom interface {
	CreateRoom() (string, int, error)
	JoinRoom(uidRoom string, uidUser string) (string, int, error)
}

type TodoUser interface {
	GetUserProfile(id int) ([]appl_row.UserProfile, int, error)
	CheckIsAdmin(id int) (bool, int, error)
}

type TodoAdmin interface {
	GetUsersConfirm(id int) ([]appl_row.GetUsersConfirm, int, error)
	GetUsersUnConfirm(id int) ([]appl_row.GetUsersUnConfirm, int, error)
	GetUserProfile(id int) ([]appl_row.UserProfile, int, error)
	UserConfirmAccount(id int, userForm appl_row.UserConfirmAccount) (bool, int, error)
}

type Repository struct {
	TodoAuth
	TodoVerify
	TodoRecovery
	TodoRoom
	TodoUser
	TodoAdmin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoAuth:     NewAuthPostgres(db),
		TodoVerify:   NewVerifyPostgres(db),
		TodoRecovery: NewRecoveryPostgres(db),
		TodoRoom:     NewRoomPostgres(db),
		TodoUser:     NewUserPostgres(db),
		TodoAdmin:    NewAdminPostgres(db),
	}
}
