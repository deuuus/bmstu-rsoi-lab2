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

func CancelPaymentStatus(reservationServiceAddress string, paymentUid string) error {
	requestURL := fmt.Sprintf("%s/api/v1/reservations/%s", reservationServiceAddress, paymentUid)

	request, err := http.NewRequest(http.MethodPost, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return err
	}

	client := http.Client{Timeout: 1 * time.Minute}

	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalf(err.Error())
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("mismatched status code: wanted 200 OK, got %d", response.StatusCode)
	}

	return nil
}

func CreateNewPayment(paymentServiceAddress string, price int) (*models.PaymentFullInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/reservations", paymentServiceAddress)

	jsonData, err := json.Marshal(price)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(jsonData))
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
		return nil, fmt.Errorf("mismatched status code: wanted 200 OK, got %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	var payment models.PaymentFullInfo
	if err := json.Unmarshal(body, &payment); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &payment, nil
}

func GetPaymentInfo(paymentServiceAddress string, paymentUid string) (*models.PaymentInfo, error) {
	requestURL := fmt.Sprintf("%s/api/v1/reservations", paymentServiceAddress)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	request.Header.Set("paymentUid", paymentUid)

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

	var payments models.PaymentInfo
	if err := json.Unmarshal(body, &payments); err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return &payments, nil
}
