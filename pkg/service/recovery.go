package service

import (
	"net/http"

	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/email"
	"github.com/rob-bender/nfc-cash-backend/pkg/hash"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
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
	statusCodeLaunch, err := s.repo.LaunchRecoveryPassword(resGetUserUidByEmail)
	if err != nil {
		return false, statusCodeLaunch, err
	}
	resSendRecoveryPasswordMail, err := email.SendRecoveryPasswordMail(userForm.Email, resGetUserUidByEmail)
	if err != nil {
		return false, resSendRecoveryPasswordMail, err
	}
	return true, http.StatusOK, nil
}

func (s *RecoveryService) CheckRecoveryPassword(uid string) (bool, int, error) {
	return s.repo.CheckRecoveryPassword(uid)
}

func (s *RecoveryService) RecoveryPasswordCompare(uid string, password string) (bool, int, error) {
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(password)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}
	password = resGeneratePasswordHash
	return s.repo.RecoveryPasswordCompare(uid, password)
}

func (s *RecoveryService) CompleteRecoveryPassword(uid string) (int, error) {
	return s.repo.CompleteRecoveryPassword(uid)
}

func (s *RecoveryService) RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error) {
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(userForm.Password)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}
	userForm.Password = resGeneratePasswordHash
	return s.repo.RecoveryPassword(userForm)
}
