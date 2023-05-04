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

func (s *TelegramService) BotCreate(botForm appl_row.BotCreate) (bool, int, error) {
	return s.repo.BotCreate(botForm)
}

func (s *TelegramService) BotDelete(botForm appl_row.BotDelete) (bool, int, error) {
	return s.repo.BotDelete(botForm)
}

func (s *TelegramService) GetBots(uid string) ([]appl_row.Bot, int, error) {
	return s.repo.GetBots(uid)
}
