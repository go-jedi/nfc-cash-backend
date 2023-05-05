package service

import (
	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
)

type TelegramService struct {
	repo repository.TodoTelegram
}

func NewTelegramService(r repository.TodoTelegram) *TelegramService {
	return &TelegramService{
		repo: r,
	}
}

func (s *TelegramService) BotCreate(id int, botForm appl_row.BotCreate) (bool, int, error) {
	return s.repo.BotCreate(id, botForm)
}

func (s *TelegramService) BotDelete(id int, botForm appl_row.BotDelete) (bool, int, error) {
	return s.repo.BotDelete(id, botForm)
}

func (s *TelegramService) GetBots(id int) ([]appl_row.Bot, int, error) {
	return s.repo.GetBots(id)
}
