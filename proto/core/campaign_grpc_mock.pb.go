// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: proto/core/campaign.proto

package core

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCampaignServiceClient is a mock of CampaignServiceClient interface.
type MockCampaignServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignServiceClientMockRecorder
}

// MockCampaignServiceClientMockRecorder is the mock recorder for MockCampaignServiceClient.
type MockCampaignServiceClientMockRecorder struct {
	mock *MockCampaignServiceClient
}

// NewMockCampaignServiceClient creates a new mock instance.
func NewMockCampaignServiceClient(ctrl *gomock.Controller) *MockCampaignServiceClient {
	mock := &MockCampaignServiceClient{ctrl: ctrl}
	mock.recorder = &MockCampaignServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignServiceClient) EXPECT() *MockCampaignServiceClientMockRecorder {
	return m.recorder
}

// StartCampaign mocks base method.
func (m *MockCampaignServiceClient) StartCampaign(ctx context.Context, in *StartCampaignRequest, opts ...grpc.CallOption) (*StartCampaignResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartCampaign", varargs...)
	ret0, _ := ret[0].(*StartCampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCampaign indicates an expected call of StartCampaign.
func (mr *MockCampaignServiceClientMockRecorder) StartCampaign(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCampaign", reflect.TypeOf((*MockCampaignServiceClient)(nil).StartCampaign), varargs...)
}

// MockCampaignServiceServer is a mock of CampaignServiceServer interface.
type MockCampaignServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignServiceServerMockRecorder
}

// MockCampaignServiceServerMockRecorder is the mock recorder for MockCampaignServiceServer.
type MockCampaignServiceServerMockRecorder struct {
	mock *MockCampaignServiceServer
}

// NewMockCampaignServiceServer creates a new mock instance.
func NewMockCampaignServiceServer(ctrl *gomock.Controller) *MockCampaignServiceServer {
	mock := &MockCampaignServiceServer{ctrl: ctrl}
	mock.recorder = &MockCampaignServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignServiceServer) EXPECT() *MockCampaignServiceServerMockRecorder {
	return m.recorder
}

// StartCampaign mocks base method.
func (m *MockCampaignServiceServer) StartCampaign(ctx context.Context, in *StartCampaignRequest) (*StartCampaignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartCampaign", ctx, in)
	ret0, _ := ret[0].(*StartCampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartCampaign indicates an expected call of StartCampaign.
func (mr *MockCampaignServiceServerMockRecorder) StartCampaign(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartCampaign", reflect.TypeOf((*MockCampaignServiceServer)(nil).StartCampaign), ctx, in)
}