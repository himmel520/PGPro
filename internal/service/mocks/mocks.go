// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/http/handler.go
//
// Generated by this command:
//
//	mockgen -source=internal/handler/http/handler.go -destination=internal/service/mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/himmel520/pgPro/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateCommand mocks base method.
func (m *MockService) CreateCommand(ctx context.Context, c *model.Command) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommand", ctx, c)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommand indicates an expected call of CreateCommand.
func (mr *MockServiceMockRecorder) CreateCommand(ctx, c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommand", reflect.TypeOf((*MockService)(nil).CreateCommand), ctx, c)
}

// DeleteCommand mocks base method.
func (m *MockService) DeleteCommand(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommand", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCommand indicates an expected call of DeleteCommand.
func (mr *MockServiceMockRecorder) DeleteCommand(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommand", reflect.TypeOf((*MockService)(nil).DeleteCommand), ctx, id)
}

// GetCommandByID mocks base method.
func (m *MockService) GetCommandByID(ctx context.Context, id string) (*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommandByID", ctx, id)
	ret0, _ := ret[0].(*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommandByID indicates an expected call of GetCommandByID.
func (mr *MockServiceMockRecorder) GetCommandByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommandByID", reflect.TypeOf((*MockService)(nil).GetCommandByID), ctx, id)
}

// GetCommandInfoByID mocks base method.
func (m *MockService) GetCommandInfoByID(ctx context.Context, id string) (*model.CommandInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommandInfoByID", ctx, id)
	ret0, _ := ret[0].(*model.CommandInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommandInfoByID indicates an expected call of GetCommandInfoByID.
func (mr *MockServiceMockRecorder) GetCommandInfoByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommandInfoByID", reflect.TypeOf((*MockService)(nil).GetCommandInfoByID), ctx, id)
}

// GetCommands mocks base method.
func (m *MockService) GetCommands(ctx context.Context) ([]*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommands", ctx)
	ret0, _ := ret[0].([]*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommands indicates an expected call of GetCommands.
func (mr *MockServiceMockRecorder) GetCommands(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommands", reflect.TypeOf((*MockService)(nil).GetCommands), ctx)
}

// RunCommand mocks base method.
func (m *MockService) RunCommand(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCommand", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCommand indicates an expected call of RunCommand.
func (mr *MockServiceMockRecorder) RunCommand(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockService)(nil).RunCommand), ctx, id)
}

// StopCommand mocks base method.
func (m *MockService) StopCommand(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopCommand", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopCommand indicates an expected call of StopCommand.
func (mr *MockServiceMockRecorder) StopCommand(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopCommand", reflect.TypeOf((*MockService)(nil).StopCommand), ctx, id)
}

// UpdateCommand mocks base method.
func (m *MockService) UpdateCommand(ctx context.Context, c *model.Command, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommand", ctx, c, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCommand indicates an expected call of UpdateCommand.
func (mr *MockServiceMockRecorder) UpdateCommand(ctx, c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommand", reflect.TypeOf((*MockService)(nil).UpdateCommand), ctx, c, id)
}
