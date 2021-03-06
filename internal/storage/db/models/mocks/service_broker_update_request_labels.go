//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_broker_update_request_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockServiceBrokerUpdateRequestLabelUpserter is a mock of ServiceBrokerUpdateRequestLabelUpserter interface.
type MockServiceBrokerUpdateRequestLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestLabelUpserterMockRecorder
}

// MockServiceBrokerUpdateRequestLabelUpserterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestLabelUpserter.
type MockServiceBrokerUpdateRequestLabelUpserterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestLabelUpserter
}

// NewMockServiceBrokerUpdateRequestLabelUpserter creates a new mock instance.
func NewMockServiceBrokerUpdateRequestLabelUpserter(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestLabelUpserter {
	mock := &MockServiceBrokerUpdateRequestLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestLabelUpserter) EXPECT() *MockServiceBrokerUpdateRequestLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelUpserter) Upsert(o *models.ServiceBrokerUpdateRequestLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceBrokerUpdateRequestLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceBrokerUpdateRequestLabelFinisher is a mock of ServiceBrokerUpdateRequestLabelFinisher interface.
type MockServiceBrokerUpdateRequestLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestLabelFinisherMockRecorder
}

// MockServiceBrokerUpdateRequestLabelFinisherMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestLabelFinisher.
type MockServiceBrokerUpdateRequestLabelFinisherMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestLabelFinisher
}

// NewMockServiceBrokerUpdateRequestLabelFinisher creates a new mock instance.
func NewMockServiceBrokerUpdateRequestLabelFinisher(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestLabelFinisher {
	mock := &MockServiceBrokerUpdateRequestLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestLabelFinisher) EXPECT() *MockServiceBrokerUpdateRequestLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceBrokerUpdateRequestLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceBrokerUpdateRequestLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceBrokerUpdateRequestLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceBrokerUpdateRequestLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceBrokerUpdateRequestLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceBrokerUpdateRequestLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceBrokerUpdateRequestLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceBrokerUpdateRequestLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelFinisher)(nil).One), ctx, exec)
}

// MockServiceBrokerUpdateRequestLabelFinder is a mock of ServiceBrokerUpdateRequestLabelFinder interface.
type MockServiceBrokerUpdateRequestLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestLabelFinderMockRecorder
}

// MockServiceBrokerUpdateRequestLabelFinderMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestLabelFinder.
type MockServiceBrokerUpdateRequestLabelFinderMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestLabelFinder
}

// NewMockServiceBrokerUpdateRequestLabelFinder creates a new mock instance.
func NewMockServiceBrokerUpdateRequestLabelFinder(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestLabelFinder {
	mock := &MockServiceBrokerUpdateRequestLabelFinder{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestLabelFinder) EXPECT() *MockServiceBrokerUpdateRequestLabelFinderMockRecorder {
	return m.recorder
}

// FindServiceBrokerUpdateRequestLabel mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelFinder) FindServiceBrokerUpdateRequestLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceBrokerUpdateRequestLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceBrokerUpdateRequestLabel", varargs...)
	ret0, _ := ret[0].(*models.ServiceBrokerUpdateRequestLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceBrokerUpdateRequestLabel indicates an expected call of FindServiceBrokerUpdateRequestLabel.
func (mr *MockServiceBrokerUpdateRequestLabelFinderMockRecorder) FindServiceBrokerUpdateRequestLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceBrokerUpdateRequestLabel", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelFinder)(nil).FindServiceBrokerUpdateRequestLabel), varargs...)
}

// MockServiceBrokerUpdateRequestLabelInserter is a mock of ServiceBrokerUpdateRequestLabelInserter interface.
type MockServiceBrokerUpdateRequestLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestLabelInserterMockRecorder
}

// MockServiceBrokerUpdateRequestLabelInserterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestLabelInserter.
type MockServiceBrokerUpdateRequestLabelInserterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestLabelInserter
}

// NewMockServiceBrokerUpdateRequestLabelInserter creates a new mock instance.
func NewMockServiceBrokerUpdateRequestLabelInserter(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestLabelInserter {
	mock := &MockServiceBrokerUpdateRequestLabelInserter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestLabelInserter) EXPECT() *MockServiceBrokerUpdateRequestLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelInserter) Insert(o *models.ServiceBrokerUpdateRequestLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceBrokerUpdateRequestLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceBrokerUpdateRequestLabelUpdater is a mock of ServiceBrokerUpdateRequestLabelUpdater interface.
type MockServiceBrokerUpdateRequestLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder
}

// MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestLabelUpdater.
type MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestLabelUpdater
}

// NewMockServiceBrokerUpdateRequestLabelUpdater creates a new mock instance.
func NewMockServiceBrokerUpdateRequestLabelUpdater(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestLabelUpdater {
	mock := &MockServiceBrokerUpdateRequestLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestLabelUpdater) EXPECT() *MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelUpdater) Update(o *models.ServiceBrokerUpdateRequestLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelUpdater) UpdateAllSlice(o models.ServiceBrokerUpdateRequestLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceBrokerUpdateRequestLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceBrokerUpdateRequestLabelDeleter is a mock of ServiceBrokerUpdateRequestLabelDeleter interface.
type MockServiceBrokerUpdateRequestLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestLabelDeleterMockRecorder
}

// MockServiceBrokerUpdateRequestLabelDeleterMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestLabelDeleter.
type MockServiceBrokerUpdateRequestLabelDeleterMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestLabelDeleter
}

// NewMockServiceBrokerUpdateRequestLabelDeleter creates a new mock instance.
func NewMockServiceBrokerUpdateRequestLabelDeleter(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestLabelDeleter {
	mock := &MockServiceBrokerUpdateRequestLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestLabelDeleter) EXPECT() *MockServiceBrokerUpdateRequestLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelDeleter) Delete(o *models.ServiceBrokerUpdateRequestLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceBrokerUpdateRequestLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceBrokerUpdateRequestLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelDeleter) DeleteAllSlice(o models.ServiceBrokerUpdateRequestLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceBrokerUpdateRequestLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceBrokerUpdateRequestLabelReloader is a mock of ServiceBrokerUpdateRequestLabelReloader interface.
type MockServiceBrokerUpdateRequestLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerUpdateRequestLabelReloaderMockRecorder
}

// MockServiceBrokerUpdateRequestLabelReloaderMockRecorder is the mock recorder for MockServiceBrokerUpdateRequestLabelReloader.
type MockServiceBrokerUpdateRequestLabelReloaderMockRecorder struct {
	mock *MockServiceBrokerUpdateRequestLabelReloader
}

// NewMockServiceBrokerUpdateRequestLabelReloader creates a new mock instance.
func NewMockServiceBrokerUpdateRequestLabelReloader(ctrl *gomock.Controller) *MockServiceBrokerUpdateRequestLabelReloader {
	mock := &MockServiceBrokerUpdateRequestLabelReloader{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerUpdateRequestLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerUpdateRequestLabelReloader) EXPECT() *MockServiceBrokerUpdateRequestLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelReloader) Reload(o *models.ServiceBrokerUpdateRequestLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceBrokerUpdateRequestLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceBrokerUpdateRequestLabelReloader) ReloadAll(o *models.ServiceBrokerUpdateRequestLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceBrokerUpdateRequestLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceBrokerUpdateRequestLabelReloader)(nil).ReloadAll), o, ctx, exec)
}
