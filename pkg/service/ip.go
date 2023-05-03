package service

import "github.com/rob-bender/nfc-cash-backend/pkg/repository"

type IpService struct {
	repo repository.TodoIp
}

func NewIpService(r repository.TodoIp) *IpService {
	return &IpService{
		repo: r,
	}
}

func (s *IpService) BlockIp(ipAddress string) (bool, int, error) {
	return s.repo.BlockIp(ipAddress)
}

func (s *IpService) CheckIpBlock(ipAddress string) (bool, int, error) {
	return s.repo.CheckIpBlock(ipAddress)
}
