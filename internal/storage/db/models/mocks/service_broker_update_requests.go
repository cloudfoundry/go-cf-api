//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_broker_update_requests.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockServiceBrokerUpdateRequestUpserter is a mock of ServiceBrokerUpdateRequestUpserter interface.
type MockServiceBrokerUpdateRequestUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestUpserterMockRecorder
}

// MockServiceBrokerUpdateRequestUpserterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestUpserter.
type MockServiceBrokerUpdateRequestUpserterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestUpserter
}

// NewMockServiceBrokerUpdateRequestUpserter creates a new mock instance.
func NewMockServiceBrokerUpdateRequestUpserter(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestUpserter {
	mock := &MockServiceBrokerUpdateRequestUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestUpserter) EXPECT() *MockServiceBrokerUpdateRequestUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceBrokerUpdateRequestUpserter) Upsert(o *models.ServiceBrokerUpdateRequest, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceBrokerUpdateRequestUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceBrokerUpdateRequestUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceBrokerUpdateRequestFinisher is a mock of ServiceBrokerUpdateRequestFinisher interface.
type MockServiceBrokerUpdateRequestFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestFinisherMockRecorder
}

// MockServiceBrokerUpdateRequestFinisherMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestFinisher.
type MockServiceBrokerUpdateRequestFinisherMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestFinisher
}

// NewMockServiceBrokerUpdateRequestFinisher creates a new mock instance.
func NewMockServiceBrokerUpdateRequestFinisher(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestFinisher {
	mock := &MockServiceBrokerUpdateRequestFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestFinisher) EXPECT() *MockServiceBrokerUpdateRequestFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceBrokerUpdateRequestFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceBrokerUpdateRequestSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceBrokerUpdateRequestSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceBrokerUpdateRequestFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceBrokerUpdateRequestFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceBrokerUpdateRequestFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceBrokerUpdateRequestFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceBrokerUpdateRequestFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceBrokerUpdateRequestFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceBrokerUpdateRequestFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceBrokerUpdateRequestFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceBrokerUpdateRequestFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceBrokerUpdateRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceBrokerUpdateRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceBrokerUpdateRequestFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceBrokerUpdateRequestFinisher)(nil).One), ctx, exec)
}

// MockServiceBrokerUpdateRequestFinder is a mock of ServiceBrokerUpdateRequestFinder interface.
type MockServiceBrokerUpdateRequestFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestFinderMockRecorder
}

// MockServiceBrokerUpdateRequestFinderMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestFinder.
type MockServiceBrokerUpdateRequestFinderMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestFinder
}

// NewMockServiceBrokerUpdateRequestFinder creates a new mock instance.
func NewMockServiceBrokerUpdateRequestFinder(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestFinder {
	mock := &MockServiceBrokerUpdateRequestFinder{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestFinder) EXPECT() *MockServiceBrokerUpdateRequestFinderMockRecorder {
	return m.recorder
}

// FindServiceBrokerUpdateRequest mocks base method.
func (m *MockServiceBrokerUpdateRequestFinder) FindServiceBrokerUpdateRequest(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceBrokerUpdateRequest, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceBrokerUpdateRequest", varargs...)
	ret0, _ := ret[0].(*models.ServiceBrokerUpdateRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceBrokerUpdateRequest indicates an expected call of FindServiceBrokerUpdateRequest.
func (mr *MockServiceBrokerUpdateRequestFinderMockRecorder) FindServiceBrokerUpdateRequest(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceBrokerUpdateRequest", reflect.TypeOf((*MockServiceBrokerUpdateRequestFinder)(nil).FindServiceBrokerUpdateRequest), varargs...)
}

// MockServiceBrokerUpdateRequestInserter is a mock of ServiceBrokerUpdateRequestInserter interface.
type MockServiceBrokerUpdateRequestInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestInserterMockRecorder
}

// MockServiceBrokerUpdateRequestInserterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestInserter.
type MockServiceBrokerUpdateRequestInserterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestInserter
}

// NewMockServiceBrokerUpdateRequestInserter creates a new mock instance.
func NewMockServiceBrokerUpdateRequestInserter(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestInserter {
	mock := &MockServiceBrokerUpdateRequestInserter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestInserter) EXPECT() *MockServiceBrokerUpdateRequestInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceBrokerUpdateRequestInserter) Insert(o *models.ServiceBrokerUpdateRequest, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceBrokerUpdateRequestInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceBrokerUpdateRequestInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceBrokerUpdateRequestUpdater is a mock of ServiceBrokerUpdateRequestUpdater interface.
type MockServiceBrokerUpdateRequestUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestUpdaterMockRecorder
}

// MockServiceBrokerUpdateRequestUpdaterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestUpdater.
type MockServiceBrokerUpdateRequestUpdaterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestUpdater
}

// NewMockServiceBrokerUpdateRequestUpdater creates a new mock instance.
func NewMockServiceBrokerUpdateRequestUpdater(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestUpdater {
	mock := &MockServiceBrokerUpdateRequestUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestUpdater) EXPECT() *MockServiceBrokerUpdateRequestUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceBrokerUpdateRequestUpdater) Update(o *models.ServiceBrokerUpdateRequest, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceBrokerUpdateRequestUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceBrokerUpdateRequestUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceBrokerUpdateRequestUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceBrokerUpdateRequestUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceBrokerUpdateRequestUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceBrokerUpdateRequestUpdater) UpdateAllSlice(o models.ServiceBrokerUpdateRequestSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceBrokerUpdateRequestUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceBrokerUpdateRequestUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceBrokerUpdateRequestDeleter is a mock of ServiceBrokerUpdateRequestDeleter interface.
type MockServiceBrokerUpdateRequestDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestDeleterMockRecorder
}

// MockServiceBrokerUpdateRequestDeleterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestDeleter.
type MockServiceBrokerUpdateRequestDeleterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestDeleter
}

// NewMockServiceBrokerUpdateRequestDeleter creates a new mock instance.
func NewMockServiceBrokerUpdateRequestDeleter(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestDeleter {
	mock := &MockServiceBrokerUpdateRequestDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestDeleter) EXPECT() *MockServiceBrokerUpdateRequestDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceBrokerUpdateRequestDeleter) Delete(o *models.ServiceBrokerUpdateRequest, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceBrokerUpdateRequestDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceBrokerUpdateRequestDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceBrokerUpdateRequestDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceBrokerUpdateRequestDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceBrokerUpdateRequestDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceBrokerUpdateRequestDeleter) DeleteAllSlice(o models.ServiceBrokerUpdateRequestSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceBrokerUpdateRequestDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceBrokerUpdateRequestDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceBrokerUpdateRequestReloader is a mock of ServiceBrokerUpdateRequestReloader interface.
type MockServiceBrokerUpdateRequestReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestReloaderMockRecorder
}

// MockServiceBrokerUpdateRequestReloaderMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestReloader.
type MockServiceBrokerUpdateRequestReloaderMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestReloader
}

// NewMockServiceBrokerUpdateRequestReloader creates a new mock instance.
func NewMockServiceBrokerUpdateRequestReloader(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestReloader {
	mock := &MockServiceBrokerUpdateRequestReloader{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestReloader) EXPECT() *MockServiceBrokerUpdateRequestReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceBrokerUpdateRequestReloader) Reload(o *models.ServiceBrokerUpdateRequest, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceBrokerUpdateRequestReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceBrokerUpdateRequestReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceBrokerUpdateRequestReloader) ReloadAll(o *models.ServiceBrokerUpdateRequestSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceBrokerUpdateRequestReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceBrokerUpdateRequestReloader)(nil).ReloadAll), o, ctx, exec)
}
