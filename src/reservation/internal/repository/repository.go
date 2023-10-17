package repository

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/models"
	"github.com/jmoiron/sqlx"
)

type Reservation interface {
	GetListOfHotels() (*[]models.HotelFullInfo, error)
	GetReservationShortInfo(username string, reservationUid string) (*models.ReservationShortInfo, error)
	GetListOfReservationShortInfo(username string) (*[]models.ReservationShortInfo, error)
	CancelReservation(username string, reservationUid string) (*models.ReservationUpdateInfo, error)
	GetHotelInfo(hotelUid string) (*models.HotelCheckInfo, error)
	CreateNewReservation(username string, requestBody models.ReserveInfo) (*models.ReservedShortInfo, error)
}

type Repository struct {
	Reservation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reservation: NewReservationPostgres(db),
	}
}
