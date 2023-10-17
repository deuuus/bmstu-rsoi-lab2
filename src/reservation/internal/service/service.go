package service

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Reservation interface {
	GetListOfHotels() (*[]models.HotelFullInfo, error)
	GetReservationShortInfo(username string, reservationUid string) (*models.ReservationShortInfo, error)
	GetListOfReservationShortInfo(username string) (*[]models.ReservationShortInfo, error)
	CancelReservation(username string, reservationUid string) (*models.ReservationUpdateInfo, error)
	GetHotelInfo(hotelUid string) (*models.HotelCheckInfo, error)
	CreateNewReservation(username string, requestBody models.ReserveInfo) (*models.ReservedShortInfo, error)
}

type Service struct {
	Reservation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Reservation: NewReservationService(repos.Reservation)}
}
