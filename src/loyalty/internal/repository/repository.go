package repository

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/models"
	"github.com/jmoiron/sqlx"
)

type Loyalty interface {
	GetLoyaltyStatus(username string) (*models.LoyaltyStatus, error)
	UpdateReservationCount(username string, acc int) (*models.LoyaltyStatusShort, error)
}

type Repository struct {
	Loyalty
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Loyalty: NewLoyaltyPostgres(db),
	}
}
