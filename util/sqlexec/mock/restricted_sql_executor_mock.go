// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/util/sqlexec (interfaces: RestrictedSQLExecutor)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	ast "github.com/pingcap/tidb/parser/ast"
	chunk "github.com/pingcap/tidb/util/chunk"
	sqlexec "github.com/pingcap/tidb/util/sqlexec"
	gomock "go.uber.org/mock/gomock"
)

// MockRestrictedSQLExecutor is a mock of RestrictedSQLExecutor interface.
type MockRestrictedSQLExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockRestrictedSQLExecutorMockRecorder
}

// MockRestrictedSQLExecutorMockRecorder is the mock recorder for MockRestrictedSQLExecutor.
type MockRestrictedSQLExecutorMockRecorder struct {
	mock *MockRestrictedSQLExecutor
}

// NewMockRestrictedSQLExecutor creates a new mock instance.
func NewMockRestrictedSQLExecutor(ctrl *gomock.Controller) *MockRestrictedSQLExecutor {
	mock := &MockRestrictedSQLExecutor{ctrl: ctrl}
	mock.recorder = &MockRestrictedSQLExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestrictedSQLExecutor) EXPECT() *MockRestrictedSQLExecutorMockRecorder {
	return m.recorder
}

// ExecRestrictedSQL mocks base method.
func (m *MockRestrictedSQLExecutor) ExecRestrictedSQL(arg0 context.Context, arg1 []func(*sqlexec.ExecOption), arg2 string, arg3 ...interface{}) ([]chunk.Row, []*ast.ResultField, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecRestrictedSQL", varargs...)
	ret0, _ := ret[0].([]chunk.Row)
	ret1, _ := ret[1].([]*ast.ResultField)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExecRestrictedSQL indicates an expected call of ExecRestrictedSQL.
func (mr *MockRestrictedSQLExecutorMockRecorder) ExecRestrictedSQL(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecRestrictedSQL", reflect.TypeOf((*MockRestrictedSQLExecutor)(nil).ExecRestrictedSQL), varargs...)
}

// ExecRestrictedStmt mocks base method.
func (m *MockRestrictedSQLExecutor) ExecRestrictedStmt(arg0 context.Context, arg1 ast.StmtNode, arg2 ...func(*sqlexec.ExecOption)) ([]chunk.Row, []*ast.ResultField, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecRestrictedStmt", varargs...)
	ret0, _ := ret[0].([]chunk.Row)
	ret1, _ := ret[1].([]*ast.ResultField)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExecRestrictedStmt indicates an expected call of ExecRestrictedStmt.
func (mr *MockRestrictedSQLExecutorMockRecorder) ExecRestrictedStmt(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecRestrictedStmt", reflect.TypeOf((*MockRestrictedSQLExecutor)(nil).ExecRestrictedStmt), varargs...)
}

// ParseWithParams mocks base method.
func (m *MockRestrictedSQLExecutor) ParseWithParams(arg0 context.Context, arg1 string, arg2 ...interface{}) (ast.StmtNode, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ParseWithParams", varargs...)
	ret0, _ := ret[0].(ast.StmtNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseWithParams indicates an expected call of ParseWithParams.
func (mr *MockRestrictedSQLExecutorMockRecorder) ParseWithParams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseWithParams", reflect.TypeOf((*MockRestrictedSQLExecutor)(nil).ParseWithParams), varargs...)
}
