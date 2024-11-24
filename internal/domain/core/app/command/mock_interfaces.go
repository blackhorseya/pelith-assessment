// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_interfaces.go -package=command -source=interfaces.go
//

// Package command is a generated GoMock package.
package command

import (
	context "context"
	reflect "reflect"

	biz "github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	model "github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskCreator is a mock of TaskCreator interface.
type MockTaskCreator struct {
	ctrl     *gomock.Controller
	recorder *MockTaskCreatorMockRecorder
}

// MockTaskCreatorMockRecorder is the mock recorder for MockTaskCreator.
type MockTaskCreatorMockRecorder struct {
	mock *MockTaskCreator
}

// NewMockTaskCreator creates a new mock instance.
func NewMockTaskCreator(ctrl *gomock.Controller) *MockTaskCreator {
	mock := &MockTaskCreator{ctrl: ctrl}
	mock.recorder = &MockTaskCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskCreator) EXPECT() *MockTaskCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTaskCreator) Create(c context.Context, task *biz.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTaskCreatorMockRecorder) Create(c, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskCreator)(nil).Create), c, task)
}

// MockTaskUpdater is a mock of TaskUpdater interface.
type MockTaskUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUpdaterMockRecorder
}

// MockTaskUpdaterMockRecorder is the mock recorder for MockTaskUpdater.
type MockTaskUpdaterMockRecorder struct {
	mock *MockTaskUpdater
}

// NewMockTaskUpdater creates a new mock instance.
func NewMockTaskUpdater(ctrl *gomock.Controller) *MockTaskUpdater {
	mock := &MockTaskUpdater{ctrl: ctrl}
	mock.recorder = &MockTaskUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUpdater) EXPECT() *MockTaskUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockTaskUpdater) Update(c context.Context, task *biz.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTaskUpdaterMockRecorder) Update(c, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskUpdater)(nil).Update), c, task)
}

// MockCampaignCreator is a mock of CampaignCreator interface.
type MockCampaignCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignCreatorMockRecorder
}

// MockCampaignCreatorMockRecorder is the mock recorder for MockCampaignCreator.
type MockCampaignCreatorMockRecorder struct {
	mock *MockCampaignCreator
}

// NewMockCampaignCreator creates a new mock instance.
func NewMockCampaignCreator(ctrl *gomock.Controller) *MockCampaignCreator {
	mock := &MockCampaignCreator{ctrl: ctrl}
	mock.recorder = &MockCampaignCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignCreator) EXPECT() *MockCampaignCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCampaignCreator) Create(c context.Context, campaign *biz.Campaign) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, campaign)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCampaignCreatorMockRecorder) Create(c, campaign any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCampaignCreator)(nil).Create), c, campaign)
}

// MockCampaignUpdater is a mock of CampaignUpdater interface.
type MockCampaignUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignUpdaterMockRecorder
}

// MockCampaignUpdaterMockRecorder is the mock recorder for MockCampaignUpdater.
type MockCampaignUpdaterMockRecorder struct {
	mock *MockCampaignUpdater
}

// NewMockCampaignUpdater creates a new mock instance.
func NewMockCampaignUpdater(ctrl *gomock.Controller) *MockCampaignUpdater {
	mock := &MockCampaignUpdater{ctrl: ctrl}
	mock.recorder = &MockCampaignUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignUpdater) EXPECT() *MockCampaignUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockCampaignUpdater) Update(c context.Context, campaign *biz.Campaign) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, campaign)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCampaignUpdaterMockRecorder) Update(c, campaign any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCampaignUpdater)(nil).Update), c, campaign)
}

// MockUserCreator is a mock of UserCreator interface.
type MockUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockUserCreatorMockRecorder
}

// MockUserCreatorMockRecorder is the mock recorder for MockUserCreator.
type MockUserCreatorMockRecorder struct {
	mock *MockUserCreator
}

// NewMockUserCreator creates a new mock instance.
func NewMockUserCreator(ctrl *gomock.Controller) *MockUserCreator {
	mock := &MockUserCreator{ctrl: ctrl}
	mock.recorder = &MockUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCreator) EXPECT() *MockUserCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserCreator) Create(c context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserCreatorMockRecorder) Create(c, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserCreator)(nil).Create), c, user)
}

// MockUserUpdater is a mock of UserUpdater interface.
type MockUserUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockUserUpdaterMockRecorder
}

// MockUserUpdaterMockRecorder is the mock recorder for MockUserUpdater.
type MockUserUpdaterMockRecorder struct {
	mock *MockUserUpdater
}

// NewMockUserUpdater creates a new mock instance.
func NewMockUserUpdater(ctrl *gomock.Controller) *MockUserUpdater {
	mock := &MockUserUpdater{ctrl: ctrl}
	mock.recorder = &MockUserUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUpdater) EXPECT() *MockUserUpdaterMockRecorder {
	return m.recorder
}

// IncrementPoints mocks base method.
func (m *MockUserUpdater) IncrementPoints(c context.Context, userID string, points int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementPoints", c, userID, points)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementPoints indicates an expected call of IncrementPoints.
func (mr *MockUserUpdaterMockRecorder) IncrementPoints(c, userID, points any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementPoints", reflect.TypeOf((*MockUserUpdater)(nil).IncrementPoints), c, userID, points)
}

// Update mocks base method.
func (m *MockUserUpdater) Update(c context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserUpdaterMockRecorder) Update(c, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserUpdater)(nil).Update), c, user)
}