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

func (s *OrderService) GetOrder(uidOrder string) ([]appl_row.Order, int, error) {
	return s.repo.GetOrder(uidOrder)
}

func (s *OrderService) GetOrders() ([]appl_row.Orders, int, error) {
	return s.repo.GetOrders()
}
