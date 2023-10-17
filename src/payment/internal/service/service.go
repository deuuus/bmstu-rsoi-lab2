package service

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Payment interface {
	GetPaymentInfo(paymentUid string) (*models.PaymentInfo, error)
	CancelPayment(paymentUid string) error
	CreatePayment(price int) (*models.PaymentFullInfo, error)
}

type Service struct {
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Payment: NewPaymentService(repos.Payment)}
}
