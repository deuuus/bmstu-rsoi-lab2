package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/service"
	mock_service "github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetListOfHotels(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockReservation, output *[]models.HotelFullInfo)
	tests := []struct {
		name                 string
		output               []models.HotelFullInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			output: []models.HotelFullInfo{
				{
					HotelUid: "hotel_uid 1",
					Name:     "name 1",
					Country:  "country 1",
					City:     "city 1",
					Address:  "address 1",
					Stars:    5,
					Price:    10000,
				},
				{
					HotelUid: "hotel_uid 2",
					Name:     "name 2",
					Country:  "country 2",
					City:     "city 2",
					Address:  "address 2",
					Stars:    3,
					Price:    5000,
				},
			},
			mockBehaviour: func(r *mock_service.MockReservation, output *[]models.HotelFullInfo) {
				r.EXPECT().GetListOfHotels().Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `[{"hotelUid":"hotel_uid 1","name":"name 1","country":"country 1","city":"city 1","address":"address 1","stars":5,"price":10000},` +
				`{"hotelUid":"hotel_uid 2","name":"name 2","country":"country 2","city":"city 2","address":"address 2","stars":3,"price":5000}]`,
		},
		{
			name: "Internal error",
			output: []models.HotelFullInfo{
				{
					HotelUid: "hotel_uid 1",
					Name:     "name 1",
					Country:  "country 1",
					City:     "city 1",
					Address:  "address 1",
					Stars:    5,
					Price:    10000,
				},
				{
					HotelUid: "hotel_uid 2",
					Name:     "name 2",
					Country:  "country 2",
					City:     "city 2",
					Address:  "address 2",
					Stars:    3,
					Price:    5000,
				},
			},
			mockBehaviour: func(r *mock_service.MockReservation, output *[]models.HotelFullInfo) {
				r.EXPECT().GetListOfHotels().Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockReservation(c)
			test.mockBehaviour(repo, &test.output)

			services := &service.Service{Reservation: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/hotels", handler.GetListOfHotels)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/hotels"), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_GetHotelInfo(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockReservation, hotelUid string, output *models.HotelCheckInfo)
	tests := []struct {
		name                 string
		hotelUid             string
		output               models.HotelCheckInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			hotelUid: "hotel_uid",
			output: models.HotelCheckInfo{
				Price: 10000,
			},
			mockBehaviour: func(r *mock_service.MockReservation, hotelUid string, output *models.HotelCheckInfo) {
				r.EXPECT().GetHotelInfo(hotelUid).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"price":10000}`,
		},
		{
			name:     "Internal error",
			hotelUid: "hotel_uid",
			output: models.HotelCheckInfo{
				Price: 10000,
			},
			mockBehaviour: func(r *mock_service.MockReservation, hotelUid string, output *models.HotelCheckInfo) {
				r.EXPECT().GetHotelInfo(hotelUid).Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockReservation(c)
			test.mockBehaviour(repo, test.hotelUid, &test.output)

			services := &service.Service{Reservation: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/hotels/:hotelUid", handler.GetHotelInfo)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/hotels/%s", test.hotelUid), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_CreateReservation(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockReservation, username string, input models.ReserveInfo, output *models.ReservedShortInfo)
	tests := []struct {
		name                 string
		username             string
		input                models.ReserveInfo
		output               models.ReservedShortInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			username: "Test Max",
			input: models.ReserveInfo{
				PaymentUid: "payment_uid",
				HotelUid:   "hotel_uid",
				StartDate:  "start",
				EndDate:    "end",
			},
			output: models.ReservedShortInfo{
				ReservationUid: "reservation_uid",
				Status:         "PAID",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, input models.ReserveInfo, output *models.ReservedShortInfo) {
				r.EXPECT().CreateNewReservation(username, input).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"reservationUid":"reservation_uid","status":"PAID"}`,
		},
		{
			name:     "Bad request",
			username: "",
			input: models.ReserveInfo{
				PaymentUid: "payment_uid",
				HotelUid:   "hotel_uid",
				StartDate:  "start",
				EndDate:    "end",
			},
			output: models.ReservedShortInfo{
				ReservationUid: "reservation_uid",
				Status:         "PAID",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, input models.ReserveInfo, output *models.ReservedShortInfo) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Header X-User-Name is not presented."}`,
		},
		{
			name:     "Internal error",
			username: "Test Max",
			input: models.ReserveInfo{
				PaymentUid: "payment_uid",
				HotelUid:   "hotel_uid",
				StartDate:  "start",
				EndDate:    "end",
			},
			output: models.ReservedShortInfo{
				ReservationUid: "reservation_uid",
				Status:         "PAID",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, input models.ReserveInfo, output *models.ReservedShortInfo) {
				r.EXPECT().CreateNewReservation(username, input).Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockReservation(c)
			test.mockBehaviour(repo, test.username, test.input, &test.output)

			services := &service.Service{Reservation: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/api/v1/reservations", handler.CreateReservation)

			body, _ := json.Marshal(test.input)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/reservations"), bytes.NewBuffer(body))
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_GetReservationShortInfo(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationShortInfo)
	tests := []struct {
		name                 string
		username             string
		reservationUid       string
		output               models.ReservationShortInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "Ok",
			username:       "Test Max",
			reservationUid: "reservation_uid",
			output: models.ReservationShortInfo{
				ReservationUid: "reservation_uid",
				PaymentUid:     "payment_uid",
				Hotel: &models.HotelInfo{
					HotelUid:    "hotel_uid",
					Name:        "name",
					FullAddress: "full_address",
					Stars:       5,
				},
				StartDate: "start",
				EndDate:   "end",
				Status:    "PAID",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationShortInfo) {
				r.EXPECT().GetReservationShortInfo(username, reservationUid).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"reservationUid":"reservation_uid","paymentUid":"payment_uid","hotel":{"hotelUid":"hotel_uid","name":"name","fullAddress":"full_address","stars":5},"startDate":"start","endDate":"end","status":"PAID"}`,
		},
		{
			name:           "Bad request (header)",
			username:       "",
			reservationUid: "reservation_uid",
			output: models.ReservationShortInfo{
				ReservationUid: "reservation_uid",
				PaymentUid:     "payment_uid",
				Hotel: &models.HotelInfo{
					HotelUid:    "hotel_uid",
					Name:        "name",
					FullAddress: "full_address",
					Stars:       5,
				},
				StartDate: "start",
				EndDate:   "end",
				Status:    "PAID",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationShortInfo) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Header X-User-Name is not presented."}`,
		},
		{
			name:           "Not found",
			username:       "Test Max",
			reservationUid: "reservation_uid",
			output: models.ReservationShortInfo{
				ReservationUid: "reservation_uid",
				PaymentUid:     "payment_uid",
				Hotel: &models.HotelInfo{
					HotelUid:    "hotel_uid",
					Name:        "name",
					FullAddress: "full_address",
					Stars:       5,
				},
				StartDate: "start",
				EndDate:   "end",
				Status:    "PAID",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationShortInfo) {
				r.EXPECT().GetReservationShortInfo(username, reservationUid).Return(nil, errors.New("not found"))
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `{"message":"not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockReservation(c)
			test.mockBehaviour(repo, test.username, test.reservationUid, &test.output)

			services := &service.Service{Reservation: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/reservations/:reservationUid", handler.GetReservationShortInfo)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/reservations/%s", test.reservationUid), nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_GetReservationsShortInfo(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockReservation, username string, output *[]models.ReservationShortInfo)
	tests := []struct {
		name                 string
		username             string
		output               []models.ReservationShortInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			username: "Test Max",
			output: []models.ReservationShortInfo{
				{
					ReservationUid: "reservation_uid",
					PaymentUid:     "payment_uid",
					Hotel: &models.HotelInfo{
						HotelUid:    "hotel_uid",
						Name:        "name",
						FullAddress: "full_address",
						Stars:       5,
					},
					StartDate: "start",
					EndDate:   "end",
					Status:    "PAID",
				},
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, output *[]models.ReservationShortInfo) {
				r.EXPECT().GetListOfReservationShortInfo(username).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"reservationUid":"reservation_uid","paymentUid":"payment_uid","hotel":{"hotelUid":"hotel_uid","name":"name","fullAddress":"full_address","stars":5},"startDate":"start","endDate":"end","status":"PAID"}]`,
		},
		{
			name:     "Bad request (header)",
			username: "",
			output: []models.ReservationShortInfo{
				{
					ReservationUid: "reservation_uid",
					PaymentUid:     "payment_uid",
					Hotel: &models.HotelInfo{
						HotelUid:    "hotel_uid",
						Name:        "name",
						FullAddress: "full_address",
						Stars:       5,
					},
					StartDate: "start",
					EndDate:   "end",
					Status:    "PAID",
				},
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, output *[]models.ReservationShortInfo) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Header X-User-Name is not presented."}`,
		},
		{
			name:     "Not found",
			username: "Test Max",
			output: []models.ReservationShortInfo{
				{
					ReservationUid: "reservation_uid",
					PaymentUid:     "payment_uid",
					Hotel: &models.HotelInfo{
						HotelUid:    "hotel_uid",
						Name:        "name",
						FullAddress: "full_address",
						Stars:       5,
					},
					StartDate: "start",
					EndDate:   "end",
					Status:    "PAID",
				},
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, output *[]models.ReservationShortInfo) {
				r.EXPECT().GetListOfReservationShortInfo(username).Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockReservation(c)
			test.mockBehaviour(repo, test.username, &test.output)

			services := &service.Service{Reservation: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/api/v1/reservations", handler.GetReservationsShortInfo)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/reservations"), nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_CancelReservation(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetOutput(ioutil.Discard)

	type mockBehaviour func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationUpdateInfo)
	tests := []struct {
		name                 string
		username             string
		reservationUid       string
		output               models.ReservationUpdateInfo
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "Ok",
			username:       "Test Max",
			reservationUid: "reservation_uid",
			output: models.ReservationUpdateInfo{
				PaymentUid: "payment_uid",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationUpdateInfo) {
				r.EXPECT().CancelReservation(username, reservationUid).Return(output, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"paymentUid":"payment_uid"}`,
		},
		{
			name:           "Bad request",
			username:       "",
			reservationUid: "reservation_uid",
			output: models.ReservationUpdateInfo{
				PaymentUid: "payment_uid",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationUpdateInfo) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Header X-User-Name is not presented."}`,
		},
		{
			name:           "Internal error",
			username:       "Test Max",
			reservationUid: "reservation_uid",
			output: models.ReservationUpdateInfo{
				PaymentUid: "payment_uid",
			},
			mockBehaviour: func(r *mock_service.MockReservation, username string, reservationUid string, output *models.ReservationUpdateInfo) {
				r.EXPECT().CancelReservation(username, reservationUid).Return(nil, errors.New("some internal error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some internal error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockReservation(c)
			test.mockBehaviour(repo, test.username, test.reservationUid, &test.output)

			services := &service.Service{Reservation: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/api/v1/reservations/:reservationUid", handler.CancelReservation)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/reservations/%s", test.reservationUid), nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
