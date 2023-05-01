package service

import (
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
)

type RoomService struct {
	repo repository.TodoRoom
}

func NewRoomService(r repository.TodoRoom) *RoomService {
	return &RoomService{
		repo: r,
	}
}

func (s *RoomService) CreateRoom() (string, int, error) {
	return s.repo.CreateRoom()
}

func (s *RoomService) JoinRoom(uidRoom string, uidUser string) (string, int, error) {
	return s.repo.JoinRoom(uidRoom, uidUser)
}
