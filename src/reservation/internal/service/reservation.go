package service

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/repository"
)

type ReservationService struct {
	repo repository.Reservation
}

func NewReservationService(repo repository.Reservation) *ReservationService {
	return &ReservationService{repo: repo}
}

func (s *ReservationService) GetListOfHotels() (*[]models.HotelFullInfo, error) {
	return s.repo.GetListOfHotels()
}

func (s *ReservationService) GetReservationShortInfo(username string, reservationUid string) (*models.ReservationShortInfo, error) {
	return s.repo.GetReservationShortInfo(username, reservationUid)
}

func (s *ReservationService) GetListOfReservationShortInfo(username string) (*[]models.ReservationShortInfo, error) {
	return s.repo.GetListOfReservationShortInfo(username)
}

func (s *ReservationService) CancelReservation(username string, reservationUid string) (*models.ReservationUpdateInfo, error) {
	return s.repo.CancelReservation(username, reservationUid)
}

func (s *ReservationService) GetHotelInfo(hotelUid string) (*models.HotelCheckInfo, error) {
	return s.repo.GetHotelInfo(hotelUid)
}

func (s *ReservationService) CreateNewReservation(username string, requestBody models.ReserveInfo) (*models.ReservedShortInfo, error) {
	return s.repo.CreateNewReservation(username, requestBody)
}
