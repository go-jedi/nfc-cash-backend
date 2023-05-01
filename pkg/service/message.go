package service

import (
	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
)

type MessageService struct {
	repo repository.TodoMessage
}

func NewMessageService(r repository.TodoMessage) *MessageService {
	return &MessageService{
		repo: r,
	}
}

func (s *MessageService) CreateMessage(messageForm appl_row.CreateMessage) (bool, int, error) {
	return s.repo.CreateMessage(messageForm)
}

func (s *MessageService) GetRoomMessages(uidRoom string) ([]appl_row.Message, int, error) {
	return s.repo.GetRoomMessages(uidRoom)
}
