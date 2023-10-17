package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/gateway/internal/models"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func UpdateReservationStatus(reservationServiceAddress string, reservationUid string, username string) (*models.ReservationUpdateInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/reservations/%s", reservationServiceAddress, reservationUid)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	request.Header.Set("X-User-Name", username)

	client := http.Client{Timeout: 1 * time.Minute}

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mismatched status code: wanted 200 OK, got %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	var reservation models.ReservationUpdateInfo
	if err := json.Unmarshal(body, &reservation); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &reservation, nil
}

func GetUserReservationsShortInfo(reservationServiceAddress string, username string) (*[]models.ReservationShortInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/reservations", reservationServiceAddress)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	request.Header.Set("X-User-Name", username)

	client := http.Client{Timeout: 1 * time.Minute}

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mismatched status code: wanted 200 OK, got %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	var reservations []models.ReservationShortInfo
	if err := json.Unmarshal(body, &reservations); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &reservations, nil
}

func GetReservationShortInfo(reservationServiceAddress string, reservationUid string, username string) (*models.ReservationShortInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/reservations/%s", reservationServiceAddress, reservationUid)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	request.Header.Set("X-User-Name", username)

	client := http.Client{Timeout: 1 * time.Minute}

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mismatched status code: wanted 200 OK, got %s", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	var reservation models.ReservationShortInfo
	if err := json.Unmarshal(body, &reservation); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &reservation, nil
}

func CreateNewReservation(reservationServiceAddress string, username string, requestBody *models.ReserveFullInfo) (*models.ReservedShortInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/reservations", reservationServiceAddress)

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(jsonData))
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	request.Header.Set("X-User-Name", username)

	client := http.Client{Timeout: 1 * time.Minute}

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mismatched status code: wanted 200 OK, got %s", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	var reservedInfo models.ReservedShortInfo
	if err := json.Unmarshal(body, &reservedInfo); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &reservedInfo, nil
}

func GetListOfHotels(reservationServiceAddress string) (*[]models.HotelFullInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/hotels", reservationServiceAddress)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	client := http.Client{Timeout: 1 * time.Minute}

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mismatched status code: wanted 200 OK, got %s", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	var hotels []models.HotelFullInfo
	if err := json.Unmarshal(body, &hotels); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &hotels, nil
}

func CheckHotelExists(reservationServiceAddress string, hotelUid string) (*models.HotelCheckInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/hotels/%s", reservationServiceAddress, hotelUid)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	client := http.Client{Timeout: 1 * time.Minute}

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mismatched status code: wanted 200 OK, got %s", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	var hotel models.HotelCheckInfo
	if err := json.Unmarshal(body, &hotel); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &hotel, nil
}
