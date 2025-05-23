// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: proto/core/campaign.proto

package core

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockCampaignService_ListCampaignsClient is a mock of CampaignService_ListCampaignsClient interface.
type MockCampaignService_ListCampaignsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignService_ListCampaignsClientMockRecorder
}

// MockCampaignService_ListCampaignsClientMockRecorder is the mock recorder for MockCampaignService_ListCampaignsClient.
type MockCampaignService_ListCampaignsClientMockRecorder struct {
	mock *MockCampaignService_ListCampaignsClient
}

// NewMockCampaignService_ListCampaignsClient creates a new mock instance.
func NewMockCampaignService_ListCampaignsClient(ctrl *gomock.Controller) *MockCampaignService_ListCampaignsClient {
	mock := &MockCampaignService_ListCampaignsClient{ctrl: ctrl}
	mock.recorder = &MockCampaignService_ListCampaignsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignService_ListCampaignsClient) EXPECT() *MockCampaignService_ListCampaignsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockCampaignService_ListCampaignsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockCampaignService_ListCampaignsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockCampaignService_ListCampaignsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockCampaignService_ListCampaignsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockCampaignService_ListCampaignsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCampaignService_ListCampaignsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockCampaignService_ListCampaignsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockCampaignService_ListCampaignsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockCampaignService_ListCampaignsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockCampaignService_ListCampaignsClient) Recv() (*GetCampaignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*GetCampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockCampaignService_ListCampaignsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockCampaignService_ListCampaignsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockCampaignService_ListCampaignsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockCampaignService_ListCampaignsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCampaignService_ListCampaignsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockCampaignService_ListCampaignsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockCampaignService_ListCampaignsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCampaignService_ListCampaignsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockCampaignService_ListCampaignsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockCampaignService_ListCampaignsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockCampaignService_ListCampaignsClient)(nil).Trailer))
}

// MockCampaignService_ListCampaignsServer is a mock of CampaignService_ListCampaignsServer interface.
type MockCampaignService_ListCampaignsServer struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignService_ListCampaignsServerMockRecorder
}

// MockCampaignService_ListCampaignsServerMockRecorder is the mock recorder for MockCampaignService_ListCampaignsServer.
type MockCampaignService_ListCampaignsServerMockRecorder struct {
	mock *MockCampaignService_ListCampaignsServer
}

// NewMockCampaignService_ListCampaignsServer creates a new mock instance.
func NewMockCampaignService_ListCampaignsServer(ctrl *gomock.Controller) *MockCampaignService_ListCampaignsServer {
	mock := &MockCampaignService_ListCampaignsServer{ctrl: ctrl}
	mock.recorder = &MockCampaignService_ListCampaignsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignService_ListCampaignsServer) EXPECT() *MockCampaignService_ListCampaignsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockCampaignService_ListCampaignsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockCampaignService_ListCampaignsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCampaignService_ListCampaignsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockCampaignService_ListCampaignsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockCampaignService_ListCampaignsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCampaignService_ListCampaignsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockCampaignService_ListCampaignsServer) Send(arg0 *GetCampaignResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockCampaignService_ListCampaignsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockCampaignService_ListCampaignsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockCampaignService_ListCampaignsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockCampaignService_ListCampaignsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockCampaignService_ListCampaignsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockCampaignService_ListCampaignsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockCampaignService_ListCampaignsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCampaignService_ListCampaignsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockCampaignService_ListCampaignsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockCampaignService_ListCampaignsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockCampaignService_ListCampaignsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockCampaignService_ListCampaignsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockCampaignService_ListCampaignsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockCampaignService_ListCampaignsServer)(nil).SetTrailer), arg0)
}

// MockCampaignService_RunBacktestByCampaignClient is a mock of CampaignService_RunBacktestByCampaignClient interface.
type MockCampaignService_RunBacktestByCampaignClient struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignService_RunBacktestByCampaignClientMockRecorder
}

// MockCampaignService_RunBacktestByCampaignClientMockRecorder is the mock recorder for MockCampaignService_RunBacktestByCampaignClient.
type MockCampaignService_RunBacktestByCampaignClientMockRecorder struct {
	mock *MockCampaignService_RunBacktestByCampaignClient
}

// NewMockCampaignService_RunBacktestByCampaignClient creates a new mock instance.
func NewMockCampaignService_RunBacktestByCampaignClient(ctrl *gomock.Controller) *MockCampaignService_RunBacktestByCampaignClient {
	mock := &MockCampaignService_RunBacktestByCampaignClient{ctrl: ctrl}
	mock.recorder = &MockCampaignService_RunBacktestByCampaignClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignService_RunBacktestByCampaignClient) EXPECT() *MockCampaignService_RunBacktestByCampaignClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockCampaignService_RunBacktestByCampaignClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockCampaignService_RunBacktestByCampaignClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignClient)(nil).Context))
}

// Header mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockCampaignService_RunBacktestByCampaignClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignClient) Recv() (*BacktestResultResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*BacktestResultResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockCampaignService_RunBacktestByCampaignClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockCampaignService_RunBacktestByCampaignClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockCampaignService_RunBacktestByCampaignClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockCampaignService_RunBacktestByCampaignClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignClient)(nil).Trailer))
}

// MockCampaignService_RunBacktestByCampaignServer is a mock of CampaignService_RunBacktestByCampaignServer interface.
type MockCampaignService_RunBacktestByCampaignServer struct {
	ctrl     *gomock.Controller
	recorder *MockCampaignService_RunBacktestByCampaignServerMockRecorder
}

// MockCampaignService_RunBacktestByCampaignServerMockRecorder is the mock recorder for MockCampaignService_RunBacktestByCampaignServer.
type MockCampaignService_RunBacktestByCampaignServerMockRecorder struct {
	mock *MockCampaignService_RunBacktestByCampaignServer
}

// NewMockCampaignService_RunBacktestByCampaignServer creates a new mock instance.
func NewMockCampaignService_RunBacktestByCampaignServer(ctrl *gomock.Controller) *MockCampaignService_RunBacktestByCampaignServer {
	mock := &MockCampaignService_RunBacktestByCampaignServer{ctrl: ctrl}
	mock.recorder = &MockCampaignService_RunBacktestByCampaignServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCampaignService_RunBacktestByCampaignServer) EXPECT() *MockCampaignService_RunBacktestByCampaignServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockCampaignService_RunBacktestByCampaignServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockCampaignService_RunBacktestByCampaignServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignServer) Send(arg0 *BacktestResultResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockCampaignService_RunBacktestByCampaignServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockCampaignService_RunBacktestByCampaignServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockCampaignService_RunBacktestByCampaignServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockCampaignService_RunBacktestByCampaignServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockCampaignService_RunBacktestByCampaignServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockCampaignService_RunBacktestByCampaignServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockCampaignService_RunBacktestByCampaignServer)(nil).SetTrailer), arg0)
}

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

// CreateCampaign mocks base method.
func (m *MockCampaignServiceClient) CreateCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...grpc.CallOption) (*CreateCampaignResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCampaign", varargs...)
	ret0, _ := ret[0].(*CreateCampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCampaign indicates an expected call of CreateCampaign.
func (mr *MockCampaignServiceClientMockRecorder) CreateCampaign(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCampaign", reflect.TypeOf((*MockCampaignServiceClient)(nil).CreateCampaign), varargs...)
}

// GetCampaign mocks base method.
func (m *MockCampaignServiceClient) GetCampaign(ctx context.Context, in *GetCampaignRequest, opts ...grpc.CallOption) (*GetCampaignResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCampaign", varargs...)
	ret0, _ := ret[0].(*GetCampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCampaign indicates an expected call of GetCampaign.
func (mr *MockCampaignServiceClientMockRecorder) GetCampaign(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCampaign", reflect.TypeOf((*MockCampaignServiceClient)(nil).GetCampaign), varargs...)
}

// ListCampaigns mocks base method.
func (m *MockCampaignServiceClient) ListCampaigns(ctx context.Context, in *ListCampaignsRequest, opts ...grpc.CallOption) (CampaignService_ListCampaignsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCampaigns", varargs...)
	ret0, _ := ret[0].(CampaignService_ListCampaignsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCampaigns indicates an expected call of ListCampaigns.
func (mr *MockCampaignServiceClientMockRecorder) ListCampaigns(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCampaigns", reflect.TypeOf((*MockCampaignServiceClient)(nil).ListCampaigns), varargs...)
}

// RunBacktestByCampaign mocks base method.
func (m *MockCampaignServiceClient) RunBacktestByCampaign(ctx context.Context, in *GetCampaignRequest, opts ...grpc.CallOption) (CampaignService_RunBacktestByCampaignClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunBacktestByCampaign", varargs...)
	ret0, _ := ret[0].(CampaignService_RunBacktestByCampaignClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunBacktestByCampaign indicates an expected call of RunBacktestByCampaign.
func (mr *MockCampaignServiceClientMockRecorder) RunBacktestByCampaign(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunBacktestByCampaign", reflect.TypeOf((*MockCampaignServiceClient)(nil).RunBacktestByCampaign), varargs...)
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

// CreateCampaign mocks base method.
func (m *MockCampaignServiceServer) CreateCampaign(ctx context.Context, in *CreateCampaignRequest) (*CreateCampaignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCampaign", ctx, in)
	ret0, _ := ret[0].(*CreateCampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCampaign indicates an expected call of CreateCampaign.
func (mr *MockCampaignServiceServerMockRecorder) CreateCampaign(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCampaign", reflect.TypeOf((*MockCampaignServiceServer)(nil).CreateCampaign), ctx, in)
}

// GetCampaign mocks base method.
func (m *MockCampaignServiceServer) GetCampaign(ctx context.Context, in *GetCampaignRequest) (*GetCampaignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCampaign", ctx, in)
	ret0, _ := ret[0].(*GetCampaignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCampaign indicates an expected call of GetCampaign.
func (mr *MockCampaignServiceServerMockRecorder) GetCampaign(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCampaign", reflect.TypeOf((*MockCampaignServiceServer)(nil).GetCampaign), ctx, in)
}

// ListCampaigns mocks base method.
func (m *MockCampaignServiceServer) ListCampaigns(blob *ListCampaignsRequest, server CampaignService_ListCampaignsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCampaigns", blob, server)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListCampaigns indicates an expected call of ListCampaigns.
func (mr *MockCampaignServiceServerMockRecorder) ListCampaigns(blob, server interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCampaigns", reflect.TypeOf((*MockCampaignServiceServer)(nil).ListCampaigns), blob, server)
}

// RunBacktestByCampaign mocks base method.
func (m *MockCampaignServiceServer) RunBacktestByCampaign(blob *GetCampaignRequest, server CampaignService_RunBacktestByCampaignServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunBacktestByCampaign", blob, server)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunBacktestByCampaign indicates an expected call of RunBacktestByCampaign.
func (mr *MockCampaignServiceServerMockRecorder) RunBacktestByCampaign(blob, server interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunBacktestByCampaign", reflect.TypeOf((*MockCampaignServiceServer)(nil).RunBacktestByCampaign), blob, server)
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
