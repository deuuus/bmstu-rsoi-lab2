// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPayment is a mock of Payment interface.
type MockPayment struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentMockRecorder
}

// MockPaymentMockRecorder is the mock recorder for MockPayment.
type MockPaymentMockRecorder struct {
	mock *MockPayment
}

// NewMockPayment creates a new mock instance.
func NewMockPayment(ctrl *gomock.Controller) *MockPayment {
	mock := &MockPayment{ctrl: ctrl}
	mock.recorder = &MockPaymentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPayment) EXPECT() *MockPaymentMockRecorder {
	return m.recorder
}

// CancelPayment mocks base method.
func (m *MockPayment) CancelPayment(paymentUid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelPayment", paymentUid)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelPayment indicates an expected call of CancelPayment.
func (mr *MockPaymentMockRecorder) CancelPayment(paymentUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelPayment", reflect.TypeOf((*MockPayment)(nil).CancelPayment), paymentUid)
}

// CreatePayment mocks base method.
func (m *MockPayment) CreatePayment(price int) (*models.PaymentFullInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", price)
	ret0, _ := ret[0].(*models.PaymentFullInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockPaymentMockRecorder) CreatePayment(price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockPayment)(nil).CreatePayment), price)
}

// GetPaymentInfo mocks base method.
func (m *MockPayment) GetPaymentInfo(paymentUid string) (*models.PaymentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentInfo", paymentUid)
	ret0, _ := ret[0].(*models.PaymentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentInfo indicates an expected call of GetPaymentInfo.
func (mr *MockPaymentMockRecorder) GetPaymentInfo(paymentUid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentInfo", reflect.TypeOf((*MockPayment)(nil).GetPaymentInfo), paymentUid)
}
