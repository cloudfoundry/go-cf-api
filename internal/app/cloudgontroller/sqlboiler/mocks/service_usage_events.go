// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_usage_events.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockServiceUsageEventFinisher is a mock of ServiceUsageEventFinisher interface.
type MockServiceUsageEventFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUsageEventFinisherMockRecorder
}

// MockServiceUsageEventFinisherMockRecorder is the mock recorder for MockServiceUsageEventFinisher.
type MockServiceUsageEventFinisherMockRecorder struct {
	mock *MockServiceUsageEventFinisher
}

// NewMockServiceUsageEventFinisher creates a new mock instance.
func NewMockServiceUsageEventFinisher(ctrl *gomock.Controller) *MockServiceUsageEventFinisher {
	mock := &MockServiceUsageEventFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceUsageEventFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUsageEventFinisher) EXPECT() *MockServiceUsageEventFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceUsageEventFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceUsageEventSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceUsageEventSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceUsageEventFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceUsageEventFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceUsageEventFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceUsageEventFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceUsageEventFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceUsageEventFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceUsageEventFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceUsageEventFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceUsageEventFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceUsageEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceUsageEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceUsageEventFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceUsageEventFinisher)(nil).One), ctx, exec)
}

// MockServiceUsageEventFinder is a mock of ServiceUsageEventFinder interface.
type MockServiceUsageEventFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUsageEventFinderMockRecorder
}

// MockServiceUsageEventFinderMockRecorder is the mock recorder for MockServiceUsageEventFinder.
type MockServiceUsageEventFinderMockRecorder struct {
	mock *MockServiceUsageEventFinder
}

// NewMockServiceUsageEventFinder creates a new mock instance.
func NewMockServiceUsageEventFinder(ctrl *gomock.Controller) *MockServiceUsageEventFinder {
	mock := &MockServiceUsageEventFinder{ctrl: ctrl}
	mock.recorder = &MockServiceUsageEventFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUsageEventFinder) EXPECT() *MockServiceUsageEventFinderMockRecorder {
	return m.recorder
}

// FindServiceUsageEvent mocks base method.
func (m *MockServiceUsageEventFinder) FindServiceUsageEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceUsageEvent, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceUsageEvent", varargs...)
	ret0, _ := ret[0].(*models.ServiceUsageEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceUsageEvent indicates an expected call of FindServiceUsageEvent.
func (mr *MockServiceUsageEventFinderMockRecorder) FindServiceUsageEvent(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceUsageEvent", reflect.TypeOf((*MockServiceUsageEventFinder)(nil).FindServiceUsageEvent), varargs...)
}

// MockServiceUsageEventInserter is a mock of ServiceUsageEventInserter interface.
type MockServiceUsageEventInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUsageEventInserterMockRecorder
}

// MockServiceUsageEventInserterMockRecorder is the mock recorder for MockServiceUsageEventInserter.
type MockServiceUsageEventInserterMockRecorder struct {
	mock *MockServiceUsageEventInserter
}

// NewMockServiceUsageEventInserter creates a new mock instance.
func NewMockServiceUsageEventInserter(ctrl *gomock.Controller) *MockServiceUsageEventInserter {
	mock := &MockServiceUsageEventInserter{ctrl: ctrl}
	mock.recorder = &MockServiceUsageEventInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUsageEventInserter) EXPECT() *MockServiceUsageEventInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceUsageEventInserter) Insert(o *models.ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceUsageEventInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceUsageEventInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceUsageEventUpdater is a mock of ServiceUsageEventUpdater interface.
type MockServiceUsageEventUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUsageEventUpdaterMockRecorder
}

// MockServiceUsageEventUpdaterMockRecorder is the mock recorder for MockServiceUsageEventUpdater.
type MockServiceUsageEventUpdaterMockRecorder struct {
	mock *MockServiceUsageEventUpdater
}

// NewMockServiceUsageEventUpdater creates a new mock instance.
func NewMockServiceUsageEventUpdater(ctrl *gomock.Controller) *MockServiceUsageEventUpdater {
	mock := &MockServiceUsageEventUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceUsageEventUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUsageEventUpdater) EXPECT() *MockServiceUsageEventUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceUsageEventUpdater) Update(o *models.ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceUsageEventUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceUsageEventUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceUsageEventUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceUsageEventUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceUsageEventUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceUsageEventUpdater) UpdateAllSlice(o models.ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceUsageEventUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceUsageEventUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceUsageEventDeleter is a mock of ServiceUsageEventDeleter interface.
type MockServiceUsageEventDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUsageEventDeleterMockRecorder
}

// MockServiceUsageEventDeleterMockRecorder is the mock recorder for MockServiceUsageEventDeleter.
type MockServiceUsageEventDeleterMockRecorder struct {
	mock *MockServiceUsageEventDeleter
}

// NewMockServiceUsageEventDeleter creates a new mock instance.
func NewMockServiceUsageEventDeleter(ctrl *gomock.Controller) *MockServiceUsageEventDeleter {
	mock := &MockServiceUsageEventDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceUsageEventDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUsageEventDeleter) EXPECT() *MockServiceUsageEventDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceUsageEventDeleter) Delete(o *models.ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceUsageEventDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceUsageEventDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceUsageEventDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceUsageEventDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceUsageEventDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceUsageEventDeleter) DeleteAllSlice(o models.ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceUsageEventDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceUsageEventDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceUsageEventReloader is a mock of ServiceUsageEventReloader interface.
type MockServiceUsageEventReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUsageEventReloaderMockRecorder
}

// MockServiceUsageEventReloaderMockRecorder is the mock recorder for MockServiceUsageEventReloader.
type MockServiceUsageEventReloaderMockRecorder struct {
	mock *MockServiceUsageEventReloader
}

// NewMockServiceUsageEventReloader creates a new mock instance.
func NewMockServiceUsageEventReloader(ctrl *gomock.Controller) *MockServiceUsageEventReloader {
	mock := &MockServiceUsageEventReloader{ctrl: ctrl}
	mock.recorder = &MockServiceUsageEventReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUsageEventReloader) EXPECT() *MockServiceUsageEventReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceUsageEventReloader) Reload(o *models.ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceUsageEventReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceUsageEventReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceUsageEventReloader) ReloadAll(o *models.ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceUsageEventReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceUsageEventReloader)(nil).ReloadAll), o, ctx, exec)
}