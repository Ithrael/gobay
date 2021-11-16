// Code generated by MockGen. DO NOT EDIT.
// Source: .//gen/grpc/git.17bdc.com/shanbay/protos-go/health.pb.go

// Package mock_protos_go is a generated GoMock package.
package mock_protos_go

import (
	context "context"
	protos_go "github.com/shanbay/gobay/testdata/health_pb"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockHealthClient is a mock of HealthClient interface.
type MockHealthClient struct {
	ctrl     *gomock.Controller
	recorder *MockHealthClientMockRecorder
}

// MockHealthClientMockRecorder is the mock recorder for MockHealthClient.
type MockHealthClientMockRecorder struct {
	mock *MockHealthClient
}

// NewMockHealthClient creates a new mock instance.
func NewMockHealthClient(ctrl *gomock.Controller) *MockHealthClient {
	mock := &MockHealthClient{ctrl: ctrl}
	mock.recorder = &MockHealthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthClient) EXPECT() *MockHealthClientMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockHealthClient) Check(ctx context.Context, in *protos_go.HealthCheckRequest, opts ...grpc.CallOption) (*protos_go.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Check", varargs...)
	ret0, _ := ret[0].(*protos_go.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockHealthClientMockRecorder) Check(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockHealthClient)(nil).Check), varargs...)
}

// MockHealthServer is a mock of HealthServer interface.
type MockHealthServer struct {
	ctrl     *gomock.Controller
	recorder *MockHealthServerMockRecorder
}

// MockHealthServerMockRecorder is the mock recorder for MockHealthServer.
type MockHealthServerMockRecorder struct {
	mock *MockHealthServer
}

// NewMockHealthServer creates a new mock instance.
func NewMockHealthServer(ctrl *gomock.Controller) *MockHealthServer {
	mock := &MockHealthServer{ctrl: ctrl}
	mock.recorder = &MockHealthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthServer) EXPECT() *MockHealthServerMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockHealthServer) Check(arg0 context.Context, arg1 *protos_go.HealthCheckRequest) (*protos_go.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(*protos_go.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockHealthServerMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockHealthServer)(nil).Check), arg0, arg1)
}
