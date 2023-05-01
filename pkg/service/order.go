package service

import (
	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
)

type OrderService struct {
	repo repository.TodoOrder
}

func NewOrderService(r repository.TodoOrder) *OrderService {
	return &OrderService{
		repo: r,
	}
}

func (s *OrderService) OrderCreate(orderForm appl_row.OrderCreate) (bool, int, error) {
	return s.repo.OrderCreate(orderForm)
}
