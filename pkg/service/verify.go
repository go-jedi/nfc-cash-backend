package service

import "github.com/rob-bender/nfc-cash-backend/pkg/repository"

type VerifyService struct {
	repo repository.TodoVerify
}

func NewVerifyService(r repository.TodoVerify) *VerifyService {
	return &VerifyService{
		repo: r,
	}
}

func (s *VerifyService) CheckEmailVerify(uid string) (bool, int, error) {
	return s.repo.CheckEmailVerify(uid)
}

func (s *VerifyService) EmailVerify(uid string) (bool, int, error) {
	return s.repo.EmailVerify(uid)
}
