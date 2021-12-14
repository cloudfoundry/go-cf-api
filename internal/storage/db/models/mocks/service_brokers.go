//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_brokers.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockServiceBrokerUpserter is a mock of ServiceBrokerUpserter interface.
type MockServiceBrokerUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpserterMockRecorder
}

// MockServiceBrokerUpserterMockRecorder is the mock recorder for MockServiceBrokerUpserter.
type MockServiceBrokerUpserterMockRecorder struct {
	mock *MockServiceBrokerUpserter
}

// NewMockServiceBrokerUpserter creates a new mock instance.
func NewMockServiceBrokerUpserter(ctrl *gomock.Controller) *MockServiceBrokerUpserter {
	mock := &MockServiceBrokerUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpserter) EXPECT() *MockServiceBrokerUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceBrokerUpserter) Upsert(o *models.ServiceBroker, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceBrokerUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceBrokerUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceBrokerFinisher is a mock of ServiceBrokerFinisher interface.
type MockServiceBrokerFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerFinisherMockRecorder
}

// MockServiceBrokerFinisherMockRecorder is the mock recorder for MockServiceBrokerFinisher.
type MockServiceBrokerFinisherMockRecorder struct {
	mock *MockServiceBrokerFinisher
}

// NewMockServiceBrokerFinisher creates a new mock instance.
func NewMockServiceBrokerFinisher(ctrl *gomock.Controller) *MockServiceBrokerFinisher {
	mock := &MockServiceBrokerFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerFinisher) EXPECT() *MockServiceBrokerFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceBrokerFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceBrokerSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceBrokerSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceBrokerFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceBrokerFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceBrokerFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceBrokerFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceBrokerFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceBrokerFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceBrokerFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceBrokerFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceBrokerFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceBroker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceBroker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceBrokerFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceBrokerFinisher)(nil).One), ctx, exec)
}

// MockServiceBrokerFinder is a mock of ServiceBrokerFinder interface.
type MockServiceBrokerFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerFinderMockRecorder
}

// MockServiceBrokerFinderMockRecorder is the mock recorder for MockServiceBrokerFinder.
type MockServiceBrokerFinderMockRecorder struct {
	mock *MockServiceBrokerFinder
}

// NewMockServiceBrokerFinder creates a new mock instance.
func NewMockServiceBrokerFinder(ctrl *gomock.Controller) *MockServiceBrokerFinder {
	mock := &MockServiceBrokerFinder{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerFinder) EXPECT() *MockServiceBrokerFinderMockRecorder {
	return m.recorder
}

// FindServiceBroker mocks base method.
func (m *MockServiceBrokerFinder) FindServiceBroker(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceBroker, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceBroker", varargs...)
	ret0, _ := ret[0].(*models.ServiceBroker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceBroker indicates an expected call of FindServiceBroker.
func (mr *MockServiceBrokerFinderMockRecorder) FindServiceBroker(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceBroker", reflect.TypeOf((*MockServiceBrokerFinder)(nil).FindServiceBroker), varargs...)
}

// MockServiceBrokerInserter is a mock of ServiceBrokerInserter interface.
type MockServiceBrokerInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerInserterMockRecorder
}

// MockServiceBrokerInserterMockRecorder is the mock recorder for MockServiceBrokerInserter.
type MockServiceBrokerInserterMockRecorder struct {
	mock *MockServiceBrokerInserter
}

// NewMockServiceBrokerInserter creates a new mock instance.
func NewMockServiceBrokerInserter(ctrl *gomock.Controller) *MockServiceBrokerInserter {
	mock := &MockServiceBrokerInserter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerInserter) EXPECT() *MockServiceBrokerInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceBrokerInserter) Insert(o *models.ServiceBroker, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceBrokerInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceBrokerInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceBrokerUpdater is a mock of ServiceBrokerUpdater interface.
type MockServiceBrokerUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdaterMockRecorder
}

// MockServiceBrokerUpdaterMockRecorder is the mock recorder for MockServiceBrokerUpdater.
type MockServiceBrokerUpdaterMockRecorder struct {
	mock *MockServiceBrokerUpdater
}

// NewMockServiceBrokerUpdater creates a new mock instance.
func NewMockServiceBrokerUpdater(ctrl *gomock.Controller) *MockServiceBrokerUpdater {
	mock := &MockServiceBrokerUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdater) EXPECT() *MockServiceBrokerUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceBrokerUpdater) Update(o *models.ServiceBroker, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceBrokerUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceBrokerUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceBrokerUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceBrokerUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceBrokerUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceBrokerUpdater) UpdateAllSlice(o models.ServiceBrokerSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceBrokerUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceBrokerUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceBrokerDeleter is a mock of ServiceBrokerDeleter interface.
type MockServiceBrokerDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerDeleterMockRecorder
}

// MockServiceBrokerDeleterMockRecorder is the mock recorder for MockServiceBrokerDeleter.
type MockServiceBrokerDeleterMockRecorder struct {
	mock *MockServiceBrokerDeleter
}

// NewMockServiceBrokerDeleter creates a new mock instance.
func NewMockServiceBrokerDeleter(ctrl *gomock.Controller) *MockServiceBrokerDeleter {
	mock := &MockServiceBrokerDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerDeleter) EXPECT() *MockServiceBrokerDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceBrokerDeleter) Delete(o *models.ServiceBroker, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceBrokerDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceBrokerDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceBrokerDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceBrokerDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceBrokerDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceBrokerDeleter) DeleteAllSlice(o models.ServiceBrokerSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceBrokerDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceBrokerDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceBrokerReloader is a mock of ServiceBrokerReloader interface.
type MockServiceBrokerReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerReloaderMockRecorder
}

// MockServiceBrokerReloaderMockRecorder is the mock recorder for MockServiceBrokerReloader.
type MockServiceBrokerReloaderMockRecorder struct {
	mock *MockServiceBrokerReloader
}

// NewMockServiceBrokerReloader creates a new mock instance.
func NewMockServiceBrokerReloader(ctrl *gomock.Controller) *MockServiceBrokerReloader {
	mock := &MockServiceBrokerReloader{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerReloader) EXPECT() *MockServiceBrokerReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceBrokerReloader) Reload(o *models.ServiceBroker, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceBrokerReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceBrokerReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceBrokerReloader) ReloadAll(o *models.ServiceBrokerSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceBrokerReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceBrokerReloader)(nil).ReloadAll), o, ctx, exec)
}
