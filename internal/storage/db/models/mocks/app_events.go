//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_app_events.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockAppEventUpserter is a mock of AppEventUpserter interface.
type MockAppEventUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppEventUpserterMockRecorder
}

// MockAppEventUpserterMockRecorder is the mock recorder for MockAppEventUpserter.
type MockAppEventUpserterMockRecorder struct {
	mock *MockAppEventUpserter
}

// NewMockAppEventUpserter creates a new mock instance.
func NewMockAppEventUpserter(ctrl *gomock.Controller) *MockAppEventUpserter {
	mock := &MockAppEventUpserter{ctrl: ctrl}
	mock.recorder = &MockAppEventUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppEventUpserter) EXPECT() *MockAppEventUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockAppEventUpserter) Upsert(o *models.AppEvent, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockAppEventUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockAppEventUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockAppEventFinisher is a mock of AppEventFinisher interface.
type MockAppEventFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockAppEventFinisherMockRecorder
}

// MockAppEventFinisherMockRecorder is the mock recorder for MockAppEventFinisher.
type MockAppEventFinisherMockRecorder struct {
	mock *MockAppEventFinisher
}

// NewMockAppEventFinisher creates a new mock instance.
func NewMockAppEventFinisher(ctrl *gomock.Controller) *MockAppEventFinisher {
	mock := &MockAppEventFinisher{ctrl: ctrl}
	mock.recorder = &MockAppEventFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppEventFinisher) EXPECT() *MockAppEventFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockAppEventFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.AppEventSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.AppEventSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockAppEventFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockAppEventFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockAppEventFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockAppEventFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockAppEventFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockAppEventFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockAppEventFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockAppEventFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockAppEventFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.AppEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.AppEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockAppEventFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockAppEventFinisher)(nil).One), ctx, exec)
}

// MockAppEventFinder is a mock of AppEventFinder interface.
type MockAppEventFinder struct {
	ctrl     *gomock.Controller
	recorder *MockAppEventFinderMockRecorder
}

// MockAppEventFinderMockRecorder is the mock recorder for MockAppEventFinder.
type MockAppEventFinderMockRecorder struct {
	mock *MockAppEventFinder
}

// NewMockAppEventFinder creates a new mock instance.
func NewMockAppEventFinder(ctrl *gomock.Controller) *MockAppEventFinder {
	mock := &MockAppEventFinder{ctrl: ctrl}
	mock.recorder = &MockAppEventFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppEventFinder) EXPECT() *MockAppEventFinderMockRecorder {
	return m.recorder
}

// FindAppEvent mocks base method.
func (m *MockAppEventFinder) FindAppEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.AppEvent, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAppEvent", varargs...)
	ret0, _ := ret[0].(*models.AppEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAppEvent indicates an expected call of FindAppEvent.
func (mr *MockAppEventFinderMockRecorder) FindAppEvent(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAppEvent", reflect.TypeOf((*MockAppEventFinder)(nil).FindAppEvent), varargs...)
}

// MockAppEventInserter is a mock of AppEventInserter interface.
type MockAppEventInserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppEventInserterMockRecorder
}

// MockAppEventInserterMockRecorder is the mock recorder for MockAppEventInserter.
type MockAppEventInserterMockRecorder struct {
	mock *MockAppEventInserter
}

// NewMockAppEventInserter creates a new mock instance.
func NewMockAppEventInserter(ctrl *gomock.Controller) *MockAppEventInserter {
	mock := &MockAppEventInserter{ctrl: ctrl}
	mock.recorder = &MockAppEventInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppEventInserter) EXPECT() *MockAppEventInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockAppEventInserter) Insert(o *models.AppEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAppEventInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAppEventInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockAppEventUpdater is a mock of AppEventUpdater interface.
type MockAppEventUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAppEventUpdaterMockRecorder
}

// MockAppEventUpdaterMockRecorder is the mock recorder for MockAppEventUpdater.
type MockAppEventUpdaterMockRecorder struct {
	mock *MockAppEventUpdater
}

// NewMockAppEventUpdater creates a new mock instance.
func NewMockAppEventUpdater(ctrl *gomock.Controller) *MockAppEventUpdater {
	mock := &MockAppEventUpdater{ctrl: ctrl}
	mock.recorder = &MockAppEventUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppEventUpdater) EXPECT() *MockAppEventUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockAppEventUpdater) Update(o *models.AppEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAppEventUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppEventUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockAppEventUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockAppEventUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockAppEventUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockAppEventUpdater) UpdateAllSlice(o models.AppEventSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockAppEventUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockAppEventUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockAppEventDeleter is a mock of AppEventDeleter interface.
type MockAppEventDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockAppEventDeleterMockRecorder
}

// MockAppEventDeleterMockRecorder is the mock recorder for MockAppEventDeleter.
type MockAppEventDeleterMockRecorder struct {
	mock *MockAppEventDeleter
}

// NewMockAppEventDeleter creates a new mock instance.
func NewMockAppEventDeleter(ctrl *gomock.Controller) *MockAppEventDeleter {
	mock := &MockAppEventDeleter{ctrl: ctrl}
	mock.recorder = &MockAppEventDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppEventDeleter) EXPECT() *MockAppEventDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAppEventDeleter) Delete(o *models.AppEvent, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAppEventDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppEventDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockAppEventDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockAppEventDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockAppEventDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockAppEventDeleter) DeleteAllSlice(o models.AppEventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockAppEventDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockAppEventDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockAppEventReloader is a mock of AppEventReloader interface.
type MockAppEventReloader struct {
	ctrl     *gomock.Controller
	recorder *MockAppEventReloaderMockRecorder
}

// MockAppEventReloaderMockRecorder is the mock recorder for MockAppEventReloader.
type MockAppEventReloaderMockRecorder struct {
	mock *MockAppEventReloader
}

// NewMockAppEventReloader creates a new mock instance.
func NewMockAppEventReloader(ctrl *gomock.Controller) *MockAppEventReloader {
	mock := &MockAppEventReloader{ctrl: ctrl}
	mock.recorder = &MockAppEventReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppEventReloader) EXPECT() *MockAppEventReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockAppEventReloader) Reload(o *models.AppEvent, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockAppEventReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockAppEventReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockAppEventReloader) ReloadAll(o *models.AppEventSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockAppEventReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockAppEventReloader)(nil).ReloadAll), o, ctx, exec)
}