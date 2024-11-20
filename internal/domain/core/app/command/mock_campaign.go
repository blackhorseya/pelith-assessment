// Code generated by MockGen. DO NOT EDIT.
// Source: campaign.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_campaign.go -package=command -source=campaign.go
//

// Package command is a generated GoMock package.
package command

import (
	context "context"
	reflect "reflect"

	model "github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
	gomock "go.uber.org/mock/gomock"
)

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
func (m *MockCampaignCreator) Create(c context.Context, campaign *model.Campaign) error {
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
func (m *MockCampaignUpdater) Update(c context.Context, campaign *model.Campaign) error {
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