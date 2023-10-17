package service

import (
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/gateway/internal/models"
	"time"
)

func GetUserLoyaltyInfoController(loyaltyServiceAddress string, username string) (*models.LoyaltyInfo, error) {
	loyaltyInfo, err := GetUserLoyaltyInfo(loyaltyServiceAddress, username)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user loyalty info: %s", err)
	}
	return loyaltyInfo, nil
}

func GetListOfHotelsController(reservationServiceAddress string) (*[]models.HotelFullInfo, error) {
	hotelsInfo, err := GetListOfHotels(reservationServiceAddress)
	if err != nil {
		return nil, fmt.Errorf("error while fetching hotels info: %s", err)
	}
	return hotelsInfo, nil
}

func GetListOfUserReservationsController(reservationServiceAddress string, paymentServiceAddress string, username string) (*[]models.ReservationInfo, error) {
	reservationsShortInfo, err := GetUserReservationsShortInfo(reservationServiceAddress, username)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user reservations info: %s", err)
	}

	reservationsInfo := make([]models.ReservationInfo, 0)
	for _, reservationShortInfo := range *reservationsShortInfo {
		paymentUid := reservationShortInfo.PaymentUid
		payment, err := GetPaymentInfo(paymentServiceAddress, paymentUid)
		if err != nil {
			return nil, fmt.Errorf("error while fetching user payment info: %s", err)
		}

		reservationInfo := models.ReservationInfo{
			ReservationUid: reservationShortInfo.ReservationUid,
			Hotel:          reservationShortInfo.Hotel,
			StartDate:      reservationShortInfo.StartDate[:10],
			EndDate:        reservationShortInfo.EndDate[:10],
			Status:         reservationShortInfo.Status,
			Payment:        payment,
		}

		reservationsInfo = append(reservationsInfo, reservationInfo)
	}

	return &reservationsInfo, nil
}

func GetReservationInfoController(reservationServiceAddress string, paymentServiceAddress string, reservationUid string, username string) (*models.ReservationInfo, error) {
	reservationShortInfo, err := GetReservationShortInfo(reservationServiceAddress, reservationUid, username)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user reservation info: %s", err)
	}

	paymentUid := reservationShortInfo.PaymentUid
	payment, err := GetPaymentInfo(paymentServiceAddress, paymentUid)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user payment info: %s", err)
	}

	reservationInfo := models.ReservationInfo{
		ReservationUid: reservationShortInfo.ReservationUid,
		Hotel:          reservationShortInfo.Hotel,
		StartDate:      reservationShortInfo.StartDate[:10],
		EndDate:        reservationShortInfo.EndDate[:10],
		Status:         reservationShortInfo.Status,
		Payment:        payment,
	}

	return &reservationInfo, nil
}

func GetUserInfoController(reservationServiceAddress string, paymentServiceAddress string, loyaltyServiceAddress string, username string) (*models.UserInfo, error) {
	reservationsShortInfo, err := GetUserReservationsShortInfo(reservationServiceAddress, username)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user reservations info: %s", err)
	}

	reservationsInfo := make([]models.ReservationInfo, 0)
	for _, reservationShortInfo := range *reservationsShortInfo {
		paymentUid := reservationShortInfo.PaymentUid
		payment, err := GetPaymentInfo(paymentServiceAddress, paymentUid)
		if err != nil {
			return nil, fmt.Errorf("error while fetching user payment info: %s", err)
		}

		reservationInfo := models.ReservationInfo{
			ReservationUid: reservationShortInfo.ReservationUid,
			Hotel:          reservationShortInfo.Hotel,
			StartDate:      reservationShortInfo.StartDate[:10],
			EndDate:        reservationShortInfo.EndDate[:10],
			Status:         reservationShortInfo.Status,
			Payment:        payment,
		}

		reservationsInfo = append(reservationsInfo, reservationInfo)
	}

	loyaltyInfo, err := GetUserLoyaltyInfo(loyaltyServiceAddress, username)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user loyalty info: %s", err)
	}

	loyaltyShortInfo := &models.LoyaltyShortInfo{
		Status:   loyaltyInfo.Status,
		Discount: loyaltyInfo.Discount,
	}

	userInfo := &models.UserInfo{
		Reservations: &reservationsInfo,
		Loyalty:      loyaltyShortInfo,
	}

	return userInfo, err
}

func CreateReservationController(reservationServiceAddress string, paymentServiceAddress string, loyaltyServiceAddress string, username string, requestBody *models.ReserveInfo) (*models.ReservedInfo, error) {
	hotel, err := CheckHotelExists(reservationServiceAddress, requestBody.HotelUid)
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02"
	startTimestamp, err := time.Parse(layout, requestBody.StartDate)
	if err != nil {
		fmt.Println("err while parsing timestamp:", err)
		return nil, err
	}
	endTimestamp, err := time.Parse(layout, requestBody.EndDate)
	if err != nil {
		fmt.Println("err while parsing timestamp:", err)
		return nil, err
	}

	differenceInDays := endTimestamp.Sub(startTimestamp).Hours() / 24
	cost := int(differenceInDays) * hotel.Price

	loyaltyInfo, err := GetUserLoyaltyInfo(loyaltyServiceAddress, username)
	if err != nil {
		return nil, fmt.Errorf("error while fetching user loyalty info: %s", err)
	}

	status := loyaltyInfo.Status
	if status == "BRONZE" {
		cost = int(0.95 * float64(cost))
	}
	if status == "SILVER" {
		cost = int(0.93 * float64(cost))
	}
	if status == "GOLD" {
		cost = int(0.9 * float64(cost))
	}

	payment, err := CreateNewPayment(paymentServiceAddress, cost)
	if err != nil {
		return nil, fmt.Errorf("error while creating new payment: %s", err)
	}

	loyaltyStatus, err := UpdateLoyaltyReservationCount(loyaltyServiceAddress, username, 1)
	if err != nil {
		return nil, fmt.Errorf("error while updating loyalty status: %s", err)
	}

	var discount int
	if loyaltyStatus.Status == "BRONZE" {
		discount = 5
	}
	if loyaltyStatus.Status == "SILVER" {
		discount = 7
	}
	if loyaltyStatus.Status == "GOLD" {
		discount = 10
	}

	newRequestBody := &models.ReserveFullInfo{
		PaymentUid: payment.PaymentUid,
		HotelUid:   requestBody.HotelUid,
		StartDate:  requestBody.StartDate,
		EndDate:    requestBody.EndDate,
	}

	reservation, err := CreateNewReservation(reservationServiceAddress, username, newRequestBody)

	reserveInfo := &models.ReservedInfo{
		ReservationUid: reservation.ReservationUid,
		HotelUid:       requestBody.HotelUid,
		StartDate:      requestBody.StartDate,
		EndDate:        requestBody.EndDate,
		Discount:       discount,
		Status:         reservation.Status,
		Payment: &models.PaymentInfo{
			Status: payment.Status,
			Price:  payment.Price,
		},
	}

	return reserveInfo, nil
}

func CancelReservationController(reservationServiceAddress string, paymentServiceAddress string, loyaltyServiceAddress string, reservationUid string, username string) error {
	reservation, err := UpdateReservationStatus(reservationServiceAddress, reservationUid, username)
	if err != nil {
		return err
	}

	err = CancelPaymentStatus(paymentServiceAddress, reservation.PaymentUid)
	if err != nil {
		return err
	}

	_, err = UpdateLoyaltyReservationCount(loyaltyServiceAddress, username, -1)
	if err != nil {
		return err
	}

	return nil
}
