// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/service.go -destination=internal/repository/mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/himmel520/pgPro/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateCommand mocks base method.
func (m *MockRepository) CreateCommand(ctx context.Context, c *model.Command) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommand", ctx, c)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommand indicates an expected call of CreateCommand.
func (mr *MockRepositoryMockRecorder) CreateCommand(ctx, c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommand", reflect.TypeOf((*MockRepository)(nil).CreateCommand), ctx, c)
}

// DeleteCommand mocks base method.
func (m *MockRepository) DeleteCommand(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommand", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCommand indicates an expected call of DeleteCommand.
func (mr *MockRepositoryMockRecorder) DeleteCommand(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommand", reflect.TypeOf((*MockRepository)(nil).DeleteCommand), ctx, id)
}

// GetCommandByID mocks base method.
func (m *MockRepository) GetCommandByID(ctx context.Context, id string) (*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommandByID", ctx, id)
	ret0, _ := ret[0].(*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommandByID indicates an expected call of GetCommandByID.
func (mr *MockRepositoryMockRecorder) GetCommandByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommandByID", reflect.TypeOf((*MockRepository)(nil).GetCommandByID), ctx, id)
}

// GetCommandInfo mocks base method.
func (m *MockRepository) GetCommandInfo(ctx context.Context, id string) (*model.CommandInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommandInfo", ctx, id)
	ret0, _ := ret[0].(*model.CommandInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommandInfo indicates an expected call of GetCommandInfo.
func (mr *MockRepositoryMockRecorder) GetCommandInfo(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommandInfo", reflect.TypeOf((*MockRepository)(nil).GetCommandInfo), ctx, id)
}

// GetCommands mocks base method.
func (m *MockRepository) GetCommands(ctx context.Context) ([]*model.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommands", ctx)
	ret0, _ := ret[0].([]*model.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommands indicates an expected call of GetCommands.
func (mr *MockRepositoryMockRecorder) GetCommands(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommands", reflect.TypeOf((*MockRepository)(nil).GetCommands), ctx)
}

// UpdateCommand mocks base method.
func (m *MockRepository) UpdateCommand(ctx context.Context, c *model.Command, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommand", ctx, c, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCommand indicates an expected call of UpdateCommand.
func (mr *MockRepositoryMockRecorder) UpdateCommand(ctx, c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommand", reflect.TypeOf((*MockRepository)(nil).UpdateCommand), ctx, c, id)
}

// UpdateCommandInfo mocks base method.
func (m *MockRepository) UpdateCommandInfo(ctx context.Context, c *model.CommandRun) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommandInfo", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCommandInfo indicates an expected call of UpdateCommandInfo.
func (mr *MockRepositoryMockRecorder) UpdateCommandInfo(ctx, c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommandInfo", reflect.TypeOf((*MockRepository)(nil).UpdateCommandInfo), ctx, c)
}