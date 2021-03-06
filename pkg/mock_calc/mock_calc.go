// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hazbo/calc-svc/pkg/calc (interfaces: CalculatorClient)

// Package mock_calc is a generated GoMock package.
package mock_calc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	calc "github.com/hazbo/calc-svc/pkg/calc"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockCalculatorClient is a mock of CalculatorClient interface
type MockCalculatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorClientMockRecorder
}

// MockCalculatorClientMockRecorder is the mock recorder for MockCalculatorClient
type MockCalculatorClientMockRecorder struct {
	mock *MockCalculatorClient
}

// NewMockCalculatorClient creates a new mock instance
func NewMockCalculatorClient(ctrl *gomock.Controller) *MockCalculatorClient {
	mock := &MockCalculatorClient{ctrl: ctrl}
	mock.recorder = &MockCalculatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCalculatorClient) EXPECT() *MockCalculatorClientMockRecorder {
	return m.recorder
}

// Div mocks base method
func (m *MockCalculatorClient) Div(arg0 context.Context, arg1 *calc.DivRequest, arg2 ...grpc.CallOption) (*calc.DivResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Div", varargs...)
	ret0, _ := ret[0].(*calc.DivResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Div indicates an expected call of Div
func (mr *MockCalculatorClientMockRecorder) Div(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Div", reflect.TypeOf((*MockCalculatorClient)(nil).Div), varargs...)
}

// Mul mocks base method
func (m *MockCalculatorClient) Mul(arg0 context.Context, arg1 *calc.MulRequest, arg2 ...grpc.CallOption) (*calc.MulResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Mul", varargs...)
	ret0, _ := ret[0].(*calc.MulResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mul indicates an expected call of Mul
func (mr *MockCalculatorClientMockRecorder) Mul(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mul", reflect.TypeOf((*MockCalculatorClient)(nil).Mul), varargs...)
}

// Sum mocks base method
func (m *MockCalculatorClient) Sum(arg0 context.Context, arg1 *calc.SumRequest, arg2 ...grpc.CallOption) (*calc.SumResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sum", varargs...)
	ret0, _ := ret[0].(*calc.SumResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sum indicates an expected call of Sum
func (mr *MockCalculatorClientMockRecorder) Sum(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockCalculatorClient)(nil).Sum), varargs...)
}
