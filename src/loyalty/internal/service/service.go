package service

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Loyalty interface {
	GetLoyaltyStatus(username string) (*models.LoyaltyStatus, error)
	UpdateReservationCount(username string, acc int) (*models.LoyaltyStatusShort, error)
}

type Service struct {
	Loyalty
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Loyalty: NewLoyaltyService(repos.Loyalty)}
}
