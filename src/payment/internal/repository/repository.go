package repository

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/models"
	"github.com/jmoiron/sqlx"
)

type Payment interface {
	GetPaymentInfo(paymentUid string) (*models.PaymentInfo, error)
	CancelPayment(paymentUid string) error
	CreatePayment(price int) (*models.PaymentFullInfo, error)
}

type Repository struct {
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Payment: NewPaymentPostgres(db),
	}
}
