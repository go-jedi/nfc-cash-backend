package service

import (
	"github.com/rob-bender/nfc-cash-backend/appl_row"
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

func (s *RoomService) LeaveRoom(uidRoom string, uidUser string) (int, error) {
	return s.repo.LeaveRoom(uidRoom, uidUser)
}

func (s *RoomService) GetRoom(uidRoom string) ([]appl_row.Room, int, error) {
	return s.repo.GetRoom(uidRoom)
}
