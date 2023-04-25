package service

import (
	"net/http"

	"github.com/rob-bender/meetsite-backend/appl_row"
	"github.com/rob-bender/meetsite-backend/pkg/email"
	"github.com/rob-bender/meetsite-backend/pkg/hash"
	"github.com/rob-bender/meetsite-backend/pkg/repository"
)

type RecoveryService struct {
	repo repository.TodoRecovery
}

func NewRecoveryService(r repository.TodoRecovery) *RecoveryService {
	return &RecoveryService{
		repo: r,
	}
}

func (s *RecoveryService) RecoveryPasswordSendMessage(userForm appl_row.RecoveryPasswordSendMessage) (bool, int, error) {
	resGetUserUidByEmail, statusCode, err := s.repo.GetUserUidByEmail(userForm)
	if err != nil {
		return false, statusCode, err
	}
	resSendRecoveryPasswordMail, err := email.SendRecoveryPasswordMail(userForm.Email, resGetUserUidByEmail)
	if err != nil {
		return false, resSendRecoveryPasswordMail, err
	}
	return true, http.StatusOK, nil
}

func (s *RecoveryService) RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error) {
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(userForm.Password)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}
	userForm.Password = resGeneratePasswordHash
	return s.repo.RecoveryPassword(userForm)
}
