//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_route_binding_operations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockRouteBindingOperationUpserter is a mock of RouteBindingOperationUpserter interface.
type MockRouteBindingOperationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockRouteBindingOperationUpserterMockRecorder
}

// MockRouteBindingOperationUpserterMockRecorder is the mock recorder for MockRouteBindingOperationUpserter.
type MockRouteBindingOperationUpserterMockRecorder struct {
	mock *MockRouteBindingOperationUpserter
}

// NewMockRouteBindingOperationUpserter creates a new mock instance.
func NewMockRouteBindingOperationUpserter(ctrl *gomock.Controller) *MockRouteBindingOperationUpserter {
	mock := &MockRouteBindingOperationUpserter{ctrl: ctrl}
	mock.recorder = &MockRouteBindingOperationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteBindingOperationUpserter) EXPECT() *MockRouteBindingOperationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockRouteBindingOperationUpserter) Upsert(o *models.RouteBindingOperation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRouteBindingOperationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRouteBindingOperationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockRouteBindingOperationFinisher is a mock of RouteBindingOperationFinisher interface.
type MockRouteBindingOperationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockRouteBindingOperationFinisherMockRecorder
}

// MockRouteBindingOperationFinisherMockRecorder is the mock recorder for MockRouteBindingOperationFinisher.
type MockRouteBindingOperationFinisherMockRecorder struct {
	mock *MockRouteBindingOperationFinisher
}

// NewMockRouteBindingOperationFinisher creates a new mock instance.
func NewMockRouteBindingOperationFinisher(ctrl *gomock.Controller) *MockRouteBindingOperationFinisher {
	mock := &MockRouteBindingOperationFinisher{ctrl: ctrl}
	mock.recorder = &MockRouteBindingOperationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteBindingOperationFinisher) EXPECT() *MockRouteBindingOperationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockRouteBindingOperationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.RouteBindingOperationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.RouteBindingOperationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockRouteBindingOperationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRouteBindingOperationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockRouteBindingOperationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRouteBindingOperationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRouteBindingOperationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockRouteBindingOperationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRouteBindingOperationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRouteBindingOperationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockRouteBindingOperationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.RouteBindingOperation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.RouteBindingOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockRouteBindingOperationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockRouteBindingOperationFinisher)(nil).One), ctx, exec)
}

// MockRouteBindingOperationFinder is a mock of RouteBindingOperationFinder interface.
type MockRouteBindingOperationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockRouteBindingOperationFinderMockRecorder
}

// MockRouteBindingOperationFinderMockRecorder is the mock recorder for MockRouteBindingOperationFinder.
type MockRouteBindingOperationFinderMockRecorder struct {
	mock *MockRouteBindingOperationFinder
}

// NewMockRouteBindingOperationFinder creates a new mock instance.
func NewMockRouteBindingOperationFinder(ctrl *gomock.Controller) *MockRouteBindingOperationFinder {
	mock := &MockRouteBindingOperationFinder{ctrl: ctrl}
	mock.recorder = &MockRouteBindingOperationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteBindingOperationFinder) EXPECT() *MockRouteBindingOperationFinderMockRecorder {
	return m.recorder
}

// FindRouteBindingOperation mocks base method.
func (m *MockRouteBindingOperationFinder) FindRouteBindingOperation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.RouteBindingOperation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindRouteBindingOperation", varargs...)
	ret0, _ := ret[0].(*models.RouteBindingOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRouteBindingOperation indicates an expected call of FindRouteBindingOperation.
func (mr *MockRouteBindingOperationFinderMockRecorder) FindRouteBindingOperation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRouteBindingOperation", reflect.TypeOf((*MockRouteBindingOperationFinder)(nil).FindRouteBindingOperation), varargs...)
}

// MockRouteBindingOperationInserter is a mock of RouteBindingOperationInserter interface.
type MockRouteBindingOperationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockRouteBindingOperationInserterMockRecorder
}

// MockRouteBindingOperationInserterMockRecorder is the mock recorder for MockRouteBindingOperationInserter.
type MockRouteBindingOperationInserterMockRecorder struct {
	mock *MockRouteBindingOperationInserter
}

// NewMockRouteBindingOperationInserter creates a new mock instance.
func NewMockRouteBindingOperationInserter(ctrl *gomock.Controller) *MockRouteBindingOperationInserter {
	mock := &MockRouteBindingOperationInserter{ctrl: ctrl}
	mock.recorder = &MockRouteBindingOperationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteBindingOperationInserter) EXPECT() *MockRouteBindingOperationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockRouteBindingOperationInserter) Insert(o *models.RouteBindingOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRouteBindingOperationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRouteBindingOperationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockRouteBindingOperationUpdater is a mock of RouteBindingOperationUpdater interface.
type MockRouteBindingOperationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockRouteBindingOperationUpdaterMockRecorder
}

// MockRouteBindingOperationUpdaterMockRecorder is the mock recorder for MockRouteBindingOperationUpdater.
type MockRouteBindingOperationUpdaterMockRecorder struct {
	mock *MockRouteBindingOperationUpdater
}

// NewMockRouteBindingOperationUpdater creates a new mock instance.
func NewMockRouteBindingOperationUpdater(ctrl *gomock.Controller) *MockRouteBindingOperationUpdater {
	mock := &MockRouteBindingOperationUpdater{ctrl: ctrl}
	mock.recorder = &MockRouteBindingOperationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteBindingOperationUpdater) EXPECT() *MockRouteBindingOperationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockRouteBindingOperationUpdater) Update(o *models.RouteBindingOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRouteBindingOperationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRouteBindingOperationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockRouteBindingOperationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockRouteBindingOperationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockRouteBindingOperationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockRouteBindingOperationUpdater) UpdateAllSlice(o models.RouteBindingOperationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockRouteBindingOperationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockRouteBindingOperationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockRouteBindingOperationDeleter is a mock of RouteBindingOperationDeleter interface.
type MockRouteBindingOperationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockRouteBindingOperationDeleterMockRecorder
}

// MockRouteBindingOperationDeleterMockRecorder is the mock recorder for MockRouteBindingOperationDeleter.
type MockRouteBindingOperationDeleterMockRecorder struct {
	mock *MockRouteBindingOperationDeleter
}

// NewMockRouteBindingOperationDeleter creates a new mock instance.
func NewMockRouteBindingOperationDeleter(ctrl *gomock.Controller) *MockRouteBindingOperationDeleter {
	mock := &MockRouteBindingOperationDeleter{ctrl: ctrl}
	mock.recorder = &MockRouteBindingOperationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteBindingOperationDeleter) EXPECT() *MockRouteBindingOperationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRouteBindingOperationDeleter) Delete(o *models.RouteBindingOperation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRouteBindingOperationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRouteBindingOperationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockRouteBindingOperationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockRouteBindingOperationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockRouteBindingOperationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockRouteBindingOperationDeleter) DeleteAllSlice(o models.RouteBindingOperationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockRouteBindingOperationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockRouteBindingOperationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockRouteBindingOperationReloader is a mock of RouteBindingOperationReloader interface.
type MockRouteBindingOperationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockRouteBindingOperationReloaderMockRecorder
}

// MockRouteBindingOperationReloaderMockRecorder is the mock recorder for MockRouteBindingOperationReloader.
type MockRouteBindingOperationReloaderMockRecorder struct {
	mock *MockRouteBindingOperationReloader
}

// NewMockRouteBindingOperationReloader creates a new mock instance.
func NewMockRouteBindingOperationReloader(ctrl *gomock.Controller) *MockRouteBindingOperationReloader {
	mock := &MockRouteBindingOperationReloader{ctrl: ctrl}
	mock.recorder = &MockRouteBindingOperationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteBindingOperationReloader) EXPECT() *MockRouteBindingOperationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockRouteBindingOperationReloader) Reload(o *models.RouteBindingOperation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockRouteBindingOperationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockRouteBindingOperationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockRouteBindingOperationReloader) ReloadAll(o *models.RouteBindingOperationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockRouteBindingOperationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockRouteBindingOperationReloader)(nil).ReloadAll), o, ctx, exec)
}
