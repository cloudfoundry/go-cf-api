//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_bindings.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockServiceBindingUpserter is a mock of ServiceBindingUpserter interface.
type MockServiceBindingUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingUpserterMockRecorder
}

// MockServiceBindingUpserterMockRecorder is the mock recorder for MockServiceBindingUpserter.
type MockServiceBindingUpserterMockRecorder struct {
	mock *MockServiceBindingUpserter
}

// NewMockServiceBindingUpserter creates a new mock instance.
func NewMockServiceBindingUpserter(ctrl *gomock.Controller) *MockServiceBindingUpserter {
	mock := &MockServiceBindingUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceBindingUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingUpserter) EXPECT() *MockServiceBindingUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceBindingUpserter) Upsert(o *models.ServiceBinding, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceBindingUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceBindingUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceBindingFinisher is a mock of ServiceBindingFinisher interface.
type MockServiceBindingFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingFinisherMockRecorder
}

// MockServiceBindingFinisherMockRecorder is the mock recorder for MockServiceBindingFinisher.
type MockServiceBindingFinisherMockRecorder struct {
	mock *MockServiceBindingFinisher
}

// NewMockServiceBindingFinisher creates a new mock instance.
func NewMockServiceBindingFinisher(ctrl *gomock.Controller) *MockServiceBindingFinisher {
	mock := &MockServiceBindingFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceBindingFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingFinisher) EXPECT() *MockServiceBindingFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceBindingFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceBindingSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceBindingSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceBindingFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceBindingFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceBindingFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceBindingFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceBindingFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceBindingFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceBindingFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceBindingFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceBindingFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceBindingFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceBindingFinisher)(nil).One), ctx, exec)
}

// MockServiceBindingFinder is a mock of ServiceBindingFinder interface.
type MockServiceBindingFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingFinderMockRecorder
}

// MockServiceBindingFinderMockRecorder is the mock recorder for MockServiceBindingFinder.
type MockServiceBindingFinderMockRecorder struct {
	mock *MockServiceBindingFinder
}

// NewMockServiceBindingFinder creates a new mock instance.
func NewMockServiceBindingFinder(ctrl *gomock.Controller) *MockServiceBindingFinder {
	mock := &MockServiceBindingFinder{ctrl: ctrl}
	mock.recorder = &MockServiceBindingFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingFinder) EXPECT() *MockServiceBindingFinderMockRecorder {
	return m.recorder
}

// FindServiceBinding mocks base method.
func (m *MockServiceBindingFinder) FindServiceBinding(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceBinding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceBinding", varargs...)
	ret0, _ := ret[0].(*models.ServiceBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceBinding indicates an expected call of FindServiceBinding.
func (mr *MockServiceBindingFinderMockRecorder) FindServiceBinding(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceBinding", reflect.TypeOf((*MockServiceBindingFinder)(nil).FindServiceBinding), varargs...)
}

// MockServiceBindingInserter is a mock of ServiceBindingInserter interface.
type MockServiceBindingInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingInserterMockRecorder
}

// MockServiceBindingInserterMockRecorder is the mock recorder for MockServiceBindingInserter.
type MockServiceBindingInserterMockRecorder struct {
	mock *MockServiceBindingInserter
}

// NewMockServiceBindingInserter creates a new mock instance.
func NewMockServiceBindingInserter(ctrl *gomock.Controller) *MockServiceBindingInserter {
	mock := &MockServiceBindingInserter{ctrl: ctrl}
	mock.recorder = &MockServiceBindingInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingInserter) EXPECT() *MockServiceBindingInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceBindingInserter) Insert(o *models.ServiceBinding, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceBindingInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceBindingInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceBindingUpdater is a mock of ServiceBindingUpdater interface.
type MockServiceBindingUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingUpdaterMockRecorder
}

// MockServiceBindingUpdaterMockRecorder is the mock recorder for MockServiceBindingUpdater.
type MockServiceBindingUpdaterMockRecorder struct {
	mock *MockServiceBindingUpdater
}

// NewMockServiceBindingUpdater creates a new mock instance.
func NewMockServiceBindingUpdater(ctrl *gomock.Controller) *MockServiceBindingUpdater {
	mock := &MockServiceBindingUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceBindingUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingUpdater) EXPECT() *MockServiceBindingUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceBindingUpdater) Update(o *models.ServiceBinding, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceBindingUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceBindingUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceBindingUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceBindingUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceBindingUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceBindingUpdater) UpdateAllSlice(o models.ServiceBindingSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceBindingUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceBindingUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceBindingDeleter is a mock of ServiceBindingDeleter interface.
type MockServiceBindingDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingDeleterMockRecorder
}

// MockServiceBindingDeleterMockRecorder is the mock recorder for MockServiceBindingDeleter.
type MockServiceBindingDeleterMockRecorder struct {
	mock *MockServiceBindingDeleter
}

// NewMockServiceBindingDeleter creates a new mock instance.
func NewMockServiceBindingDeleter(ctrl *gomock.Controller) *MockServiceBindingDeleter {
	mock := &MockServiceBindingDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceBindingDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingDeleter) EXPECT() *MockServiceBindingDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceBindingDeleter) Delete(o *models.ServiceBinding, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceBindingDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceBindingDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceBindingDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceBindingDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceBindingDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceBindingDeleter) DeleteAllSlice(o models.ServiceBindingSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceBindingDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceBindingDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceBindingReloader is a mock of ServiceBindingReloader interface.
type MockServiceBindingReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingReloaderMockRecorder
}

// MockServiceBindingReloaderMockRecorder is the mock recorder for MockServiceBindingReloader.
type MockServiceBindingReloaderMockRecorder struct {
	mock *MockServiceBindingReloader
}

// NewMockServiceBindingReloader creates a new mock instance.
func NewMockServiceBindingReloader(ctrl *gomock.Controller) *MockServiceBindingReloader {
	mock := &MockServiceBindingReloader{ctrl: ctrl}
	mock.recorder = &MockServiceBindingReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingReloader) EXPECT() *MockServiceBindingReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceBindingReloader) Reload(o *models.ServiceBinding, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceBindingReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceBindingReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceBindingReloader) ReloadAll(o *models.ServiceBindingSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceBindingReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceBindingReloader)(nil).ReloadAll), o, ctx, exec)
}
