package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/models"
	"github.com/jmoiron/sqlx"
)

type LoyaltyPostgres struct {
	db *sqlx.DB
}

func NewLoyaltyPostgres(db *sqlx.DB) *LoyaltyPostgres {
	return &LoyaltyPostgres{db: db}
}

func (r *LoyaltyPostgres) UpdateReservationCount(username string, acc int) (*models.LoyaltyStatusShort, error) {
	var query string
	if acc == 1 {
		query = fmt.Sprintf("UPDATE loyalty SET reservation_count = reservation_count + 1, status = CASE WHEN reservation_count + 1 >= 20 THEN 'GOLD' WHEN reservation_count + 1 >= 10 THEN 'SILVER' ELSE 'BRONZE' END WHERE username = $1 RETURNING status")
	} else {
		query = fmt.Sprintf("UPDATE loyalty SET reservation_count = GREATEST(reservation_count - 1, 0), status = CASE WHEN GREATEST(reservation_count - 1, 0) >= 20 THEN 'GOLD' WHEN GREATEST(reservation_count - 1, 0) >= 10 THEN 'SILVER' ELSE 'BRONZE' END WHERE username = $1 RETURNING status")
	}

	var status string
	row := r.db.QueryRow(query, username)
	if err := row.Scan(&status); err != nil {
		return nil, err
	}
	loyaltyStatusShort := models.LoyaltyStatusShort{
		Status: status,
	}
	return &loyaltyStatusShort, nil
}

func (r *LoyaltyPostgres) GetLoyaltyStatus(username string) (*models.LoyaltyStatus, error) {
	var loyaltyStatus models.LoyaltyStatus

	query := fmt.Sprintf("SELECT status, discount, reservation_count FROM loyalty WHERE username = $1")

	if err := r.db.Get(&loyaltyStatus, query, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			loyaltyStatus = models.LoyaltyStatus{
				Username:         username,
				Status:           "BRONZE",
				Discount:         0,
				ReservationCount: 0,
			}
			var id int
			query = fmt.Sprintf("INSERT INTO loyalty (username, reservation_count, status, discount) VALUES ($1, $2, $3, $4) RETURNING id")
			row := r.db.QueryRow(query, loyaltyStatus.Username, loyaltyStatus.ReservationCount, loyaltyStatus.Status, loyaltyStatus.Discount)
			if err := row.Scan(&id); err != nil {
				return &loyaltyStatus, err
			}
			return &loyaltyStatus, nil
		}
	}

	loyaltyStatus.Username = username

	return &loyaltyStatus, nil
}
