package service

import (
	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
)

type AdminService struct {
	repo repository.TodoAdmin
}

func NewAdminService(r repository.TodoAdmin) *AdminService {
	return &AdminService{
		repo: r,
	}
}

func (s *AdminService) GetUsersConfirm(id int) ([]appl_row.GetUsersConfirm, int, error) {
	return s.repo.GetUsersConfirm(id)
}

func (s *AdminService) GetUsersUnConfirm(id int) ([]appl_row.GetUsersUnConfirm, int, error) {
	return s.repo.GetUsersUnConfirm(id)
}

func (s *AdminService) UserConfirmAccount(id int, userForm appl_row.UserConfirmAccount) (bool, int, error) {
	return s.repo.UserConfirmAccount(id, userForm)
}
