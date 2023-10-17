package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/service"
	mock_service "github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler_CreatePayment(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPayment, price string, output *models.PaymentFullInfo)
	tests := []struct {
		name                 string
		price                string
		output               models.PaymentFullInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:  "Ok",
			price: "27000",
			output: models.PaymentFullInfo{
				PaymentUid: "test_uid",
				Status:     "PAID",
				Price:      27000,
			},
			mockBehaviour: func(r *mock_service.MockPayment, price string, output *models.PaymentFullInfo) {
				priceInt, _ := strconv.Atoi(price)
				r.EXPECT().CreatePayment(priceInt).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"paymentUid":"test_uid","status":"PAID","price":27000}`,
		},
		{
			name:  "Bad request",
			price: "ABC",
			output: models.PaymentFullInfo{
				PaymentUid: "test_uid",
				Status:     "PAID",
				Price:      27000,
			},
			mockBehaviour:        func(r *mock_service.MockPayment, price string, output *models.PaymentFullInfo) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid character 'A' looking for beginning of value"}`,
		},
		{
			name:  "Internal error",
			price: "27000",
			output: models.PaymentFullInfo{
				PaymentUid: "test_uid",
				Status:     "PAID",
				Price:      27000,
			},
			mockBehaviour: func(r *mock_service.MockPayment, price string, output *models.PaymentFullInfo) {
				priceInt, _ := strconv.Atoi(price)
				r.EXPECT().CreatePayment(priceInt).Return(output, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPayment(c)
			test.mockBehaviour(repo, test.price, &test.output)

			services := &service.Service{Payment: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/api/v1/reservations", handler.CreatePayment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/reservations"), bytes.NewBufferString(test.price))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_GetPaymentInfo(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPayment, paymentUid string, output *models.PaymentInfo)
	tests := []struct {
		name                 string
		paymentUid           string
		output               models.PaymentInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "Ok",
			paymentUid: "test_uid",
			output: models.PaymentInfo{
				Status: "PAID",
				Price:  27000,
			},
			mockBehaviour: func(r *mock_service.MockPayment, paymentUid string, output *models.PaymentInfo) {
				r.EXPECT().GetPaymentInfo(paymentUid).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"PAID","price":27000}`,
		},
		{
			name:       "Bad request",
			paymentUid: "",
			output: models.PaymentInfo{
				Status: "PAID",
				Price:  27000,
			},
			mockBehaviour:        func(r *mock_service.MockPayment, paymentUid string, output *models.PaymentInfo) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Header paymentUid is not presented."}`,
		},
		{
			name:       "Internal error",
			paymentUid: "test_uid",
			output: models.PaymentInfo{
				Status: "PAID",
				Price:  27000,
			},
			mockBehaviour: func(r *mock_service.MockPayment, paymentUid string, output *models.PaymentInfo) {
				r.EXPECT().GetPaymentInfo(paymentUid).Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPayment(c)
			test.mockBehaviour(repo, test.paymentUid, &test.output)

			services := &service.Service{Payment: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/reservations", handler.GetPaymentInfo)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/reservations"), nil)
			req.Header.Set("paymentUid", test.paymentUid)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_CancelPayment(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockPayment, paymentUid string)
	tests := []struct {
		name                 string
		paymentUid           string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "Ok",
			paymentUid: "test_uid",
			mockBehaviour: func(r *mock_service.MockPayment, paymentUid string) {
				r.EXPECT().CancelPayment(paymentUid).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Bad request",
			paymentUid:         "",
			mockBehaviour:      func(r *mock_service.MockPayment, paymentUid string) {},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:       "Internal server error",
			paymentUid: "test_uid",
			mockBehaviour: func(r *mock_service.MockPayment, paymentUid string) {
				r.EXPECT().CancelPayment(paymentUid).Return(errors.New("some internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPayment(c)
			test.mockBehaviour(repo, test.paymentUid)

			services := &service.Service{Payment: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/reservations/:paymentUid", handler.CancelPayment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/reservations/%s", test.paymentUid), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
