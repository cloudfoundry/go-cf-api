//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_services.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockServiceUpserter is a mock of ServiceUpserter interface.
type MockServiceUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUpserterMockRecorder
}

// MockServiceUpserterMockRecorder is the mock recorder for MockServiceUpserter.
type MockServiceUpserterMockRecorder struct {
	mock *MockServiceUpserter
}

// NewMockServiceUpserter creates a new mock instance.
func NewMockServiceUpserter(ctrl *gomock.Controller) *MockServiceUpserter {
	mock := &MockServiceUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUpserter) EXPECT() *MockServiceUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceUpserter) Upsert(o *models.Service, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceFinisher is a mock of ServiceFinisher interface.
type MockServiceFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceFinisherMockRecorder
}

// MockServiceFinisherMockRecorder is the mock recorder for MockServiceFinisher.
type MockServiceFinisherMockRecorder struct {
	mock *MockServiceFinisher
}

// NewMockServiceFinisher creates a new mock instance.
func NewMockServiceFinisher(ctrl *gomock.Controller) *MockServiceFinisher {
	mock := &MockServiceFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceFinisher) EXPECT() *MockServiceFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceFinisher)(nil).One), ctx, exec)
}

// MockServiceFinder is a mock of ServiceFinder interface.
type MockServiceFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceFinderMockRecorder
}

// MockServiceFinderMockRecorder is the mock recorder for MockServiceFinder.
type MockServiceFinderMockRecorder struct {
	mock *MockServiceFinder
}

// NewMockServiceFinder creates a new mock instance.
func NewMockServiceFinder(ctrl *gomock.Controller) *MockServiceFinder {
	mock := &MockServiceFinder{ctrl: ctrl}
	mock.recorder = &MockServiceFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceFinder) EXPECT() *MockServiceFinderMockRecorder {
	return m.recorder
}

// FindService mocks base method.
func (m *MockServiceFinder) FindService(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Service, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindService", varargs...)
	ret0, _ := ret[0].(*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindService indicates an expected call of FindService.
func (mr *MockServiceFinderMockRecorder) FindService(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindService", reflect.TypeOf((*MockServiceFinder)(nil).FindService), varargs...)
}

// MockServiceInserter is a mock of ServiceInserter interface.
type MockServiceInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInserterMockRecorder
}

// MockServiceInserterMockRecorder is the mock recorder for MockServiceInserter.
type MockServiceInserterMockRecorder struct {
	mock *MockServiceInserter
}

// NewMockServiceInserter creates a new mock instance.
func NewMockServiceInserter(ctrl *gomock.Controller) *MockServiceInserter {
	mock := &MockServiceInserter{ctrl: ctrl}
	mock.recorder = &MockServiceInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInserter) EXPECT() *MockServiceInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceInserter) Insert(o *models.Service, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceUpdater is a mock of ServiceUpdater interface.
type MockServiceUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUpdaterMockRecorder
}

// MockServiceUpdaterMockRecorder is the mock recorder for MockServiceUpdater.
type MockServiceUpdaterMockRecorder struct {
	mock *MockServiceUpdater
}

// NewMockServiceUpdater creates a new mock instance.
func NewMockServiceUpdater(ctrl *gomock.Controller) *MockServiceUpdater {
	mock := &MockServiceUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUpdater) EXPECT() *MockServiceUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceUpdater) Update(o *models.Service, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceUpdater) UpdateAllSlice(o models.ServiceSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceDeleter is a mock of ServiceDeleter interface.
type MockServiceDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDeleterMockRecorder
}

// MockServiceDeleterMockRecorder is the mock recorder for MockServiceDeleter.
type MockServiceDeleterMockRecorder struct {
	mock *MockServiceDeleter
}

// NewMockServiceDeleter creates a new mock instance.
func NewMockServiceDeleter(ctrl *gomock.Controller) *MockServiceDeleter {
	mock := &MockServiceDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDeleter) EXPECT() *MockServiceDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceDeleter) Delete(o *models.Service, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceDeleter) DeleteAllSlice(o models.ServiceSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceReloader is a mock of ServiceReloader interface.
type MockServiceReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceReloaderMockRecorder
}

// MockServiceReloaderMockRecorder is the mock recorder for MockServiceReloader.
type MockServiceReloaderMockRecorder struct {
	mock *MockServiceReloader
}

// NewMockServiceReloader creates a new mock instance.
func NewMockServiceReloader(ctrl *gomock.Controller) *MockServiceReloader {
	mock := &MockServiceReloader{ctrl: ctrl}
	mock.recorder = &MockServiceReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceReloader) EXPECT() *MockServiceReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceReloader) Reload(o *models.Service, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceReloader) ReloadAll(o *models.ServiceSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceReloader)(nil).ReloadAll), o, ctx, exec)
}
