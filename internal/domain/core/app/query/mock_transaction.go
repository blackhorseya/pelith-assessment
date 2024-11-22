// Code generated by MockGen. DO NOT EDIT.
// Source: transaction.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_transaction.go -package=query -source=transaction.go
//

// Package query is a generated GoMock package.
package query

import (
	context "context"
	reflect "reflect"

	biz "github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	gomock "go.uber.org/mock/gomock"
)

// MockTransactionGetter is a mock of TransactionGetter interface.
type MockTransactionGetter struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionGetterMockRecorder
}

// MockTransactionGetterMockRecorder is the mock recorder for MockTransactionGetter.
type MockTransactionGetterMockRecorder struct {
	mock *MockTransactionGetter
}

// NewMockTransactionGetter creates a new mock instance.
func NewMockTransactionGetter(ctrl *gomock.Controller) *MockTransactionGetter {
	mock := &MockTransactionGetter{ctrl: ctrl}
	mock.recorder = &MockTransactionGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionGetter) EXPECT() *MockTransactionGetterMockRecorder {
	return m.recorder
}

// ListByAddress mocks base method.
func (m *MockTransactionGetter) ListByAddress(c context.Context, address string, cond ListTransactionCondition) (biz.TransactionList, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByAddress", c, address, cond)
	ret0, _ := ret[0].(biz.TransactionList)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByAddress indicates an expected call of ListByAddress.
func (mr *MockTransactionGetterMockRecorder) ListByAddress(c, address, cond any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByAddress", reflect.TypeOf((*MockTransactionGetter)(nil).ListByAddress), c, address, cond)
}