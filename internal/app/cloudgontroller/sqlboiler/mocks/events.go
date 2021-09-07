// +build unit

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_events.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockEventFinisher is a mock of EventFinisher interface.
type MockEventFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockEventFinisherMockRecorder
}

// MockEventFinisherMockRecorder is the mock recorder for MockEventFinisher.
type MockEventFinisherMockRecorder struct {
	mock *MockEventFinisher
}

// NewMockEventFinisher creates a new mock instance.
func NewMockEventFinisher(ctrl *gomock.Controller) *MockEventFinisher {
	mock := &MockEventFinisher{ctrl: ctrl}
	mock.recorder = &MockEventFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventFinisher) EXPECT() *MockEventFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockEventFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.EventSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.EventSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockEventFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockEventFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockEventFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockEventFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockEventFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockEventFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockEventFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockEventFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockEventFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockEventFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockEventFinisher)(nil).One), ctx, exec)
}

// MockEventFinder is a mock of EventFinder interface.
type MockEventFinder struct {
	ctrl     *gomock.Controller
	recorder *MockEventFinderMockRecorder
}

// MockEventFinderMockRecorder is the mock recorder for MockEventFinder.
type MockEventFinderMockRecorder struct {
	mock *MockEventFinder
}

// NewMockEventFinder creates a new mock instance.
func NewMockEventFinder(ctrl *gomock.Controller) *MockEventFinder {
	mock := &MockEventFinder{ctrl: ctrl}
	mock.recorder = &MockEventFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventFinder) EXPECT() *MockEventFinderMockRecorder {
	return m.recorder
}

// FindEvent mocks base method.
func (m *MockEventFinder) FindEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Event, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindEvent", varargs...)
	ret0, _ := ret[0].(*models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEvent indicates an expected call of FindEvent.
func (mr *MockEventFinderMockRecorder) FindEvent(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEvent", reflect.TypeOf((*MockEventFinder)(nil).FindEvent), varargs...)
}

// MockEventInserter is a mock of EventInserter interface.
type MockEventInserter struct {
	ctrl     *gomock.Controller
	recorder *MockEventInserterMockRecorder
}

// MockEventInserterMockRecorder is the mock recorder for MockEventInserter.
type MockEventInserterMockRecorder struct {
	mock *MockEventInserter
}

// NewMockEventInserter creates a new mock instance.
func NewMockEventInserter(ctrl *gomock.Controller) *MockEventInserter {
	mock := &MockEventInserter{ctrl: ctrl}
	mock.recorder = &MockEventInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventInserter) EXPECT() *MockEventInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockEventInserter) Insert(o *models.Event, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockEventInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockEventInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockEventUpdater is a mock of EventUpdater interface.
type MockEventUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockEventUpdaterMockRecorder
}

// MockEventUpdaterMockRecorder is the mock recorder for MockEventUpdater.
type MockEventUpdaterMockRecorder struct {
	mock *MockEventUpdater
}

// NewMockEventUpdater creates a new mock instance.
func NewMockEventUpdater(ctrl *gomock.Controller) *MockEventUpdater {
	mock := &MockEventUpdater{ctrl: ctrl}
	mock.recorder = &MockEventUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventUpdater) EXPECT() *MockEventUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockEventUpdater) Update(o *models.Event, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEventUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockEventUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockEventUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockEventUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockEventUpdater) UpdateAllSlice(o models.EventSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockEventUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockEventUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockEventUpserter is a mock of EventUpserter interface.
type MockEventUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockEventUpserterMockRecorder
}

// MockEventUpserterMockRecorder is the mock recorder for MockEventUpserter.
type MockEventUpserterMockRecorder struct {
	mock *MockEventUpserter
}

// NewMockEventUpserter creates a new mock instance.
func NewMockEventUpserter(ctrl *gomock.Controller) *MockEventUpserter {
	mock := &MockEventUpserter{ctrl: ctrl}
	mock.recorder = &MockEventUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventUpserter) EXPECT() *MockEventUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockEventUpserter) Upsert(o *models.Event, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockEventUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockEventUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockEventDeleter is a mock of EventDeleter interface.
type MockEventDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockEventDeleterMockRecorder
}

// MockEventDeleterMockRecorder is the mock recorder for MockEventDeleter.
type MockEventDeleterMockRecorder struct {
	mock *MockEventDeleter
}

// NewMockEventDeleter creates a new mock instance.
func NewMockEventDeleter(ctrl *gomock.Controller) *MockEventDeleter {
	mock := &MockEventDeleter{ctrl: ctrl}
	mock.recorder = &MockEventDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventDeleter) EXPECT() *MockEventDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockEventDeleter) Delete(o *models.Event, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEventDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockEventDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockEventDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockEventDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockEventDeleter) DeleteAllSlice(o models.EventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockEventDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockEventDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockEventReloader is a mock of EventReloader interface.
type MockEventReloader struct {
	ctrl     *gomock.Controller
	recorder *MockEventReloaderMockRecorder
}

// MockEventReloaderMockRecorder is the mock recorder for MockEventReloader.
type MockEventReloaderMockRecorder struct {
	mock *MockEventReloader
}

// NewMockEventReloader creates a new mock instance.
func NewMockEventReloader(ctrl *gomock.Controller) *MockEventReloader {
	mock := &MockEventReloader{ctrl: ctrl}
	mock.recorder = &MockEventReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventReloader) EXPECT() *MockEventReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockEventReloader) Reload(o *models.Event, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockEventReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockEventReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockEventReloader) ReloadAll(o *models.EventSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockEventReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockEventReloader)(nil).ReloadAll), o, ctx, exec)
}
