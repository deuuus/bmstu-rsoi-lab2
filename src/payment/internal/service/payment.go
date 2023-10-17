package service

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/repository"
)

type PaymentService struct {
	repo repository.Payment
}

func NewPaymentService(repo repository.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(price int) (*models.PaymentFullInfo, error) {
	return s.repo.CreatePayment(price)
}

func (s *PaymentService) GetPaymentInfo(paymentUid string) (*models.PaymentInfo, error) {
	return s.repo.GetPaymentInfo(paymentUid)
}

func (s *PaymentService) CancelPayment(paymentUid string) error {
	return s.repo.CancelPayment(paymentUid)
}
