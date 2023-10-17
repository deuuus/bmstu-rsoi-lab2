package repository

import (
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type ReservationPostgres struct {
	db *sqlx.DB
}

func NewReservationPostgres(db *sqlx.DB) *ReservationPostgres {
	return &ReservationPostgres{db: db}
}

func (r *ReservationPostgres) CreateNewReservation(username string, requestBody models.ReserveInfo) (*models.ReservedShortInfo, error) {
	var hotelId string
	query := fmt.Sprintf("SELECT id FROM hotels WHERE hotel_uid = $1")
	if err := r.db.Get(&hotelId, query, requestBody.HotelUid); err != nil {
		return nil, err
	}

	query = fmt.Sprintf("INSERT INTO reservation (reservation_uid, username, payment_uid, hotel_id, status, start_date, end_data) VALUES ('ac409926-7dd3-47d2-b337-03ba0f865d74', $1, $2, $3, 'PAID', $4, $5) RETURNING reservation_uid, status")

	row := r.db.QueryRow(query, username, requestBody.PaymentUid, hotelId, requestBody.StartDate, requestBody.EndDate)

	var reservationUid string
	var status string
	if err := row.Scan(&reservationUid, &status); err != nil {
		return nil, err
	}
	return &models.ReservedShortInfo{ReservationUid: reservationUid, Status: status}, nil
}

func (r *ReservationPostgres) GetListOfHotels() (*[]models.HotelFullInfo, error) {
	var hotels []models.HotelFullInfo

	query := fmt.Sprintf("SELECT hotel_uid, name, country, city, address, stars, price FROM hotels")
	err := r.db.Select(&hotels, query)
	return &hotels, err
}

func (r *ReservationPostgres) GetHotelInfo(hotelUid string) (*models.HotelCheckInfo, error) {
	var hotelCheckInfo models.HotelCheckInfo

	query := "SELECT price FROM hotels WHERE hotel_uid = $1"
	if err := r.db.Get(&hotelCheckInfo, query, hotelUid); err != nil {
		return nil, err
	}

	return &hotelCheckInfo, nil
}

func (r *ReservationPostgres) CancelReservation(username string, reservationUid string) (*models.ReservationUpdateInfo, error) {
	var reservationUpdateInfo models.ReservationUpdateInfo

	query := fmt.Sprintf("UPDATE reservation SET status = `CANCELED` WHERE reservation_uid = $1 AND username = $2")
	_, err := r.db.Exec(query, username, reservationUid)
	if err != nil {
		return nil, err
	}

	query = fmt.Sprintf("SELECT payment_uid FROM reservation WHERE reservation_uid = $1 AND username = $2")
	if err = r.db.Get(&reservationUpdateInfo, query, username, reservationUid); err != nil {
		return nil, err
	}

	return &reservationUpdateInfo, nil
}

func (r *ReservationPostgres) GetReservationShortInfo(username string, reservationUid string) (*models.ReservationShortInfo, error) {
	var reservationSingleInfo models.ReservationSingleInfo

	query := fmt.Sprintf("SELECT payment_uid, hotel_id, status, start_date, end_data FROM reservation WHERE reservation_uid = $1 and username = $2")

	if err := r.db.Get(&reservationSingleInfo, query, reservationUid, username); err != nil {
		return nil, err
	}

	var hotelShortInfo models.HotelShortInfo
	query = fmt.Sprintf("SELECT hotel_uid, name, country, city, address, stars FROM hotels WHERE id = $1")
	if err := r.db.Get(&hotelShortInfo, query, reservationSingleInfo.HotelId); err != nil {
		return nil, err
	}

	strSlice := []string{hotelShortInfo.Country, hotelShortInfo.City, hotelShortInfo.Address}
	fullAddress := strings.Join(strSlice, ", ")

	reservationShortInfo := models.ReservationShortInfo{
		ReservationUid: reservationUid,
		PaymentUid:     reservationSingleInfo.PaymentUid,
		Hotel: &models.HotelInfo{
			HotelUid:    hotelShortInfo.HotelUid,
			Name:        hotelShortInfo.Name,
			FullAddress: fullAddress,
			Stars:       hotelShortInfo.Stars,
		},
		StartDate: reservationSingleInfo.StartDate,
		EndDate:   reservationSingleInfo.EndDate,
		Status:    reservationSingleInfo.Status,
	}

	return &reservationShortInfo, nil
}

func (r *ReservationPostgres) GetListOfReservationShortInfo(username string) (*[]models.ReservationShortInfo, error) {
	var reservations []models.ReservationFullInfo

	query := fmt.Sprintf("SELECT reservation_uid, payment_uid, hotel_id, status, start_date, end_data FROM reservation WHERE username = $1")
	if err := r.db.Select(&reservations, query, username); err != nil {
		logrus.Info(err.Error())
		return nil, err
	}

	reservationsShortInfo := make([]models.ReservationShortInfo, 0)
	query = fmt.Sprintf("SELECT hotel_uid, name, country, city, address, stars FROM hotels WHERE id = $1")

	for _, reservation := range reservations {
		hotelId := reservation.HotelId
		var hotel models.HotelShortInfo
		if err := r.db.Get(&hotel, query, hotelId); err != nil {
			logrus.Info(err.Error())
			return nil, err
		}

		strSlice := []string{hotel.Country, hotel.City, hotel.Address}
		fullAddress := strings.Join(strSlice, ", ")
		reservationShortInfo := models.ReservationShortInfo{
			ReservationUid: reservation.ReservationUid,
			PaymentUid:     reservation.PaymentUid,
			Hotel: &models.HotelInfo{
				HotelUid:    hotel.HotelUid,
				Name:        hotel.Name,
				FullAddress: fullAddress,
				Stars:       hotel.Stars,
			},
			StartDate: reservation.StartDate,
			EndDate:   reservation.EndDate,
			Status:    reservation.Status,
		}
		reservationsShortInfo = append(reservationsShortInfo, reservationShortInfo)
	}
	return &reservationsShortInfo, nil
}
