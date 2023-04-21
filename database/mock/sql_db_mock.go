// Code generated by MockGen. DO NOT EDIT.
// Source: database/sql_db.go

// Package sql_db is a generated GoMock package.
package sql_db

import (
	context "context"
	sql "database/sql"
	driver "database/sql/driver"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockDBInterface is a mock of DBInterface interface.
type MockDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBInterfaceMockRecorder
}

// MockDBInterfaceMockRecorder is the mock recorder for MockDBInterface.
type MockDBInterfaceMockRecorder struct {
	mock *MockDBInterface
}

// NewMockDBInterface creates a new mock instance.
func NewMockDBInterface(ctrl *gomock.Controller) *MockDBInterface {
	mock := &MockDBInterface{ctrl: ctrl}
	mock.recorder = &MockDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBInterface) EXPECT() *MockDBInterfaceMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockDBInterface) Begin() (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockDBInterfaceMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockDBInterface)(nil).Begin))
}

// BeginTx mocks base method.
func (m *MockDBInterface) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, opts)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockDBInterfaceMockRecorder) BeginTx(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockDBInterface)(nil).BeginTx), ctx, opts)
}

// Close mocks base method.
func (m *MockDBInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDBInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDBInterface)(nil).Close))
}

// Conn mocks base method.
func (m *MockDBInterface) Conn(ctx context.Context) (*sql.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn", ctx)
	ret0, _ := ret[0].(*sql.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Conn indicates an expected call of Conn.
func (mr *MockDBInterfaceMockRecorder) Conn(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockDBInterface)(nil).Conn), ctx)
}

// Driver mocks base method.
func (m *MockDBInterface) Driver() driver.Driver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Driver")
	ret0, _ := ret[0].(driver.Driver)
	return ret0
}

// Driver indicates an expected call of Driver.
func (mr *MockDBInterfaceMockRecorder) Driver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Driver", reflect.TypeOf((*MockDBInterface)(nil).Driver))
}

// Exec mocks base method.
func (m *MockDBInterface) Exec(query string, args ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockDBInterfaceMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDBInterface)(nil).Exec), varargs...)
}

// ExecContext mocks base method.
func (m *MockDBInterface) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockDBInterfaceMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockDBInterface)(nil).ExecContext), varargs...)
}

// Ping mocks base method.
func (m *MockDBInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockDBInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDBInterface)(nil).Ping))
}

// PingContext mocks base method.
func (m *MockDBInterface) PingContext(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingContext", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingContext indicates an expected call of PingContext.
func (mr *MockDBInterfaceMockRecorder) PingContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingContext", reflect.TypeOf((*MockDBInterface)(nil).PingContext), ctx)
}

// Prepare mocks base method.
func (m *MockDBInterface) Prepare(query string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockDBInterfaceMockRecorder) Prepare(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockDBInterface)(nil).Prepare), query)
}

// PrepareContext mocks base method.
func (m *MockDBInterface) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareContext", ctx, query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareContext indicates an expected call of PrepareContext.
func (mr *MockDBInterfaceMockRecorder) PrepareContext(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareContext", reflect.TypeOf((*MockDBInterface)(nil).PrepareContext), ctx, query)
}

// Query mocks base method.
func (m *MockDBInterface) Query(query string, args ...any) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDBInterfaceMockRecorder) Query(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDBInterface)(nil).Query), varargs...)
}

// QueryContext mocks base method.
func (m *MockDBInterface) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockDBInterfaceMockRecorder) QueryContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockDBInterface)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockDBInterface) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockDBInterfaceMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockDBInterface)(nil).QueryRowContext), varargs...)
}

// SetConnMaxIdleTime mocks base method.
func (m *MockDBInterface) SetConnMaxIdleTime(d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConnMaxIdleTime", d)
}

// SetConnMaxIdleTime indicates an expected call of SetConnMaxIdleTime.
func (mr *MockDBInterfaceMockRecorder) SetConnMaxIdleTime(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConnMaxIdleTime", reflect.TypeOf((*MockDBInterface)(nil).SetConnMaxIdleTime), d)
}

// SetConnMaxLifetime mocks base method.
func (m *MockDBInterface) SetConnMaxLifetime(d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConnMaxLifetime", d)
}

// SetConnMaxLifetime indicates an expected call of SetConnMaxLifetime.
func (mr *MockDBInterfaceMockRecorder) SetConnMaxLifetime(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConnMaxLifetime", reflect.TypeOf((*MockDBInterface)(nil).SetConnMaxLifetime), d)
}

// SetMaxIdleConns mocks base method.
func (m *MockDBInterface) SetMaxIdleConns(n int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxIdleConns", n)
}

// SetMaxIdleConns indicates an expected call of SetMaxIdleConns.
func (mr *MockDBInterfaceMockRecorder) SetMaxIdleConns(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxIdleConns", reflect.TypeOf((*MockDBInterface)(nil).SetMaxIdleConns), n)
}

// SetMaxOpenConns mocks base method.
func (m *MockDBInterface) SetMaxOpenConns(n int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxOpenConns", n)
}

// SetMaxOpenConns indicates an expected call of SetMaxOpenConns.
func (mr *MockDBInterfaceMockRecorder) SetMaxOpenConns(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxOpenConns", reflect.TypeOf((*MockDBInterface)(nil).SetMaxOpenConns), n)
}

// Stats mocks base method.
func (m *MockDBInterface) Stats() sql.DBStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(sql.DBStats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockDBInterfaceMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockDBInterface)(nil).Stats))
}
