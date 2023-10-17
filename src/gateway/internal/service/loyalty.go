package service

import (
	"encoding/json"
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/gateway/internal/models"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func UpdateLoyaltyReservationCount(loyaltyServiceAddress string, username string, acc int) (*models.LoyaltyStatus, error) {
	requestURL := fmt.Sprintf("%s/api/v1/loyalty", loyaltyServiceAddress)

	request, err := http.NewRequest(http.MethodPost, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}
	request.Header.Set("X-User-Name", username)
	request.Header.Set("Acc", strconv.Itoa(acc))

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

	var loyaltyInfo models.LoyaltyStatus
	if err := json.Unmarshal(body, &loyaltyInfo); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &loyaltyInfo, nil
}

func GetUserLoyaltyInfo(loyaltyServiceAddress string, username string) (*models.LoyaltyInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/loyalty", loyaltyServiceAddress)

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

	var loyaltyInfo models.LoyaltyInfo
	if err := json.Unmarshal(body, &loyaltyInfo); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &loyaltyInfo, nil
}
