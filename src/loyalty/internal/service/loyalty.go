package service

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/repository"
)

type LoyaltyService struct {
	repo repository.Loyalty
}

func NewLoyaltyService(repo repository.Loyalty) *LoyaltyService {
	return &LoyaltyService{repo: repo}
}

func (s *LoyaltyService) GetLoyaltyStatus(username string) (*models.LoyaltyStatus, error) {
	return s.repo.GetLoyaltyStatus(username)
}

func (s *LoyaltyService) UpdateReservationCount(username string, acc int) (*models.LoyaltyStatusShort, error) {
	return s.repo.UpdateReservationCount(username, acc)
}
