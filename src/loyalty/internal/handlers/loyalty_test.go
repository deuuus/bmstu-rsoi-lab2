package handlers

import (
	"errors"
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/service"
	mock_service "github.com/deuuus/bmsru-rsoi-lab2/src/loyalty/internal/service/mocks"
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

func TestHandler_GetLoyaltyStatus(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockLoyalty, username string, output *models.LoyaltyStatus)
	tests := []struct {
		name                 string
		output               models.LoyaltyStatus
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			output: models.LoyaltyStatus{
				Username:         "Test Max",
				ReservationCount: 10,
				Status:           "GOLD",
				Discount:         10,
			},
			username: "Test Max",
			mockBehaviour: func(r *mock_service.MockLoyalty, username string, output *models.LoyaltyStatus) {
				r.EXPECT().GetLoyaltyStatus(username).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"username":"Test Max","reservationCount":10,"status":"GOLD","discount":10}`,
		},
		{
			name: "Not Found",
			output: models.LoyaltyStatus{
				Username:         "Test Max",
				ReservationCount: 10,
				Status:           "GOLD",
				Discount:         10,
			},
			username: "Test Min",
			mockBehaviour: func(r *mock_service.MockLoyalty, username string, output *models.LoyaltyStatus) {
				r.EXPECT().GetLoyaltyStatus(username).Return(output, errors.New("not found"))
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `{"message":"not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockLoyalty(c)
			test.mockBehaviour(repo, test.username, &test.output)

			services := &service.Service{Loyalty: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/loyalty", handler.GetLoyaltyStatus)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/loyalty"), nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_UpdateReservationsCount(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockLoyalty, username string, acc int, output *models.LoyaltyStatusShort)
	tests := []struct {
		name                 string
		output               models.LoyaltyStatusShort
		username             string
		acc                  int
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			output: models.LoyaltyStatusShort{
				Status: "GOLD",
			},
			username: "Test Max",
			acc:      1,
			mockBehaviour: func(r *mock_service.MockLoyalty, username string, acc int, output *models.LoyaltyStatusShort) {
				r.EXPECT().UpdateReservationCount(username, acc).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"GOLD"}`,
		},
		{
			name: "Internal error",
			output: models.LoyaltyStatusShort{
				Status: "GOLD",
			},
			username: "Test Max",
			acc:      1,
			mockBehaviour: func(r *mock_service.MockLoyalty, username string, acc int, output *models.LoyaltyStatusShort) {
				r.EXPECT().UpdateReservationCount(username, acc).Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockLoyalty(c)
			test.mockBehaviour(repo, test.username, test.acc, &test.output)

			services := &service.Service{Loyalty: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/api/v1/loyalty", handler.UpdateReservationsCount)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/loyalty"), nil)
			req.Header.Set("X-User-Name", test.username)
			req.Header.Set("Acc", strconv.Itoa(test.acc))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
