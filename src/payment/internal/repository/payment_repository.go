package repository

import (
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/models"
	"github.com/jmoiron/sqlx"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}

func (r *PaymentPostgres) GetPaymentInfo(paymentUid string) (*models.PaymentInfo, error) {
	var payment models.PaymentInfo

	query := fmt.Sprintf("SELECT status, price FROM payment WHERE payment_uid = $1")

	if err := r.db.Get(&payment, query, paymentUid); err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentPostgres) CancelPayment(paymentUid string) error {
	query := fmt.Sprintf("UPDATE payment SET status = 'CANCELED' WHERE payment_uid = $1")
	_, err := r.db.Exec(query, paymentUid)
	return err
}

func (r *PaymentPostgres) CreatePayment(price int) (*models.PaymentFullInfo, error) {
	query := fmt.Sprintf("INSERT INTO payment (payment_uid, status, price) VALUES (gen_random_uuid(), 'PAID', $1) RETURNING id, payment_uid")
	row := r.db.QueryRow(query, price)
	var id int
	var paymentUid string
	if err := row.Scan(&id, &paymentUid); err != nil {
		return nil, err
	}
	paymentFullInfo := models.PaymentFullInfo{
		PaymentUid: paymentUid,
		Status:     "PAID",
		Price:      price,
	}
	return &paymentFullInfo, nil
}
