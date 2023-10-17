// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockReservation is a mock of Reservation interface.
type MockReservation struct {
	ctrl     *gomock.Controller
	recorder *MockReservationMockRecorder
}

// MockReservationMockRecorder is the mock recorder for MockReservation.
type MockReservationMockRecorder struct {
	mock *MockReservation
}

// NewMockReservation creates a new mock instance.
func NewMockReservation(ctrl *gomock.Controller) *MockReservation {
	mock := &MockReservation{ctrl: ctrl}
	mock.recorder = &MockReservationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReservation) EXPECT() *MockReservationMockRecorder {
	return m.recorder
}

// CancelReservation mocks base method.
func (m *MockReservation) CancelReservation(username, reservationUid string) (*models.ReservationUpdateInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelReservation", username, reservationUid)
	ret0, _ := ret[0].(*models.ReservationUpdateInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelReservation indicates an expected call of CancelReservation.
func (mr *MockReservationMockRecorder) CancelReservation(username, reservationUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelReservation", reflect.TypeOf((*MockReservation)(nil).CancelReservation), username, reservationUid)
}

// CreateNewReservation mocks base method.
func (m *MockReservation) CreateNewReservation(username string, requestBody models.ReserveInfo) (*models.ReservedShortInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewReservation", username, requestBody)
	ret0, _ := ret[0].(*models.ReservedShortInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewReservation indicates an expected call of CreateNewReservation.
func (mr *MockReservationMockRecorder) CreateNewReservation(username, requestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewReservation", reflect.TypeOf((*MockReservation)(nil).CreateNewReservation), username, requestBody)
}

// GetHotelInfo mocks base method.
func (m *MockReservation) GetHotelInfo(hotelUid string) (*models.HotelCheckInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHotelInfo", hotelUid)
	ret0, _ := ret[0].(*models.HotelCheckInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHotelInfo indicates an expected call of GetHotelInfo.
func (mr *MockReservationMockRecorder) GetHotelInfo(hotelUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHotelInfo", reflect.TypeOf((*MockReservation)(nil).GetHotelInfo), hotelUid)
}

// GetListOfHotels mocks base method.
func (m *MockReservation) GetListOfHotels() (*[]models.HotelFullInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListOfHotels")
	ret0, _ := ret[0].(*[]models.HotelFullInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListOfHotels indicates an expected call of GetListOfHotels.
func (mr *MockReservationMockRecorder) GetListOfHotels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListOfHotels", reflect.TypeOf((*MockReservation)(nil).GetListOfHotels))
}

// GetListOfReservationShortInfo mocks base method.
func (m *MockReservation) GetListOfReservationShortInfo(username string) (*[]models.ReservationShortInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListOfReservationShortInfo", username)
	ret0, _ := ret[0].(*[]models.ReservationShortInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListOfReservationShortInfo indicates an expected call of GetListOfReservationShortInfo.
func (mr *MockReservationMockRecorder) GetListOfReservationShortInfo(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListOfReservationShortInfo", reflect.TypeOf((*MockReservation)(nil).GetListOfReservationShortInfo), username)
}

// GetReservationShortInfo mocks base method.
func (m *MockReservation) GetReservationShortInfo(username, reservationUid string) (*models.ReservationShortInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservationShortInfo", username, reservationUid)
	ret0, _ := ret[0].(*models.ReservationShortInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationShortInfo indicates an expected call of GetReservationShortInfo.
func (mr *MockReservationMockRecorder) GetReservationShortInfo(username, reservationUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationShortInfo", reflect.TypeOf((*MockReservation)(nil).GetReservationShortInfo), username, reservationUid)
}
