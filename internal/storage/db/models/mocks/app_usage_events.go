//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_app_usage_events.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockAppUsageEventUpserter is a mock of AppUsageEventUpserter interface.
type MockAppUsageEventUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppUsageEventUpserterMockRecorder
}

// MockAppUsageEventUpserterMockRecorder is the mock recorder for MockAppUsageEventUpserter.
type MockAppUsageEventUpserterMockRecorder struct {
	mock *MockAppUsageEventUpserter
}

// NewMockAppUsageEventUpserter creates a new mock instance.
func NewMockAppUsageEventUpserter(ctrl *gomock.Controller) *MockAppUsageEventUpserter {
	mock := &MockAppUsageEventUpserter{ctrl: ctrl}
	mock.recorder = &MockAppUsageEventUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUsageEventUpserter) EXPECT() *MockAppUsageEventUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockAppUsageEventUpserter) Upsert(o *models.AppUsageEvent, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockAppUsageEventUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockAppUsageEventUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockAppUsageEventFinisher is a mock of AppUsageEventFinisher interface.
type MockAppUsageEventFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockAppUsageEventFinisherMockRecorder
}

// MockAppUsageEventFinisherMockRecorder is the mock recorder for MockAppUsageEventFinisher.
type MockAppUsageEventFinisherMockRecorder struct {
	mock *MockAppUsageEventFinisher
}

// NewMockAppUsageEventFinisher creates a new mock instance.
func NewMockAppUsageEventFinisher(ctrl *gomock.Controller) *MockAppUsageEventFinisher {
	mock := &MockAppUsageEventFinisher{ctrl: ctrl}
	mock.recorder = &MockAppUsageEventFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUsageEventFinisher) EXPECT() *MockAppUsageEventFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockAppUsageEventFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.AppUsageEventSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.AppUsageEventSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockAppUsageEventFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockAppUsageEventFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockAppUsageEventFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockAppUsageEventFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockAppUsageEventFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockAppUsageEventFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockAppUsageEventFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockAppUsageEventFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockAppUsageEventFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.AppUsageEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.AppUsageEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockAppUsageEventFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockAppUsageEventFinisher)(nil).One), ctx, exec)
}

// MockAppUsageEventFinder is a mock of AppUsageEventFinder interface.
type MockAppUsageEventFinder struct {
	ctrl     *gomock.Controller
	recorder *MockAppUsageEventFinderMockRecorder
}

// MockAppUsageEventFinderMockRecorder is the mock recorder for MockAppUsageEventFinder.
type MockAppUsageEventFinderMockRecorder struct {
	mock *MockAppUsageEventFinder
}

// NewMockAppUsageEventFinder creates a new mock instance.
func NewMockAppUsageEventFinder(ctrl *gomock.Controller) *MockAppUsageEventFinder {
	mock := &MockAppUsageEventFinder{ctrl: ctrl}
	mock.recorder = &MockAppUsageEventFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUsageEventFinder) EXPECT() *MockAppUsageEventFinderMockRecorder {
	return m.recorder
}

// FindAppUsageEvent mocks base method.
func (m *MockAppUsageEventFinder) FindAppUsageEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.AppUsageEvent, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAppUsageEvent", varargs...)
	ret0, _ := ret[0].(*models.AppUsageEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAppUsageEvent indicates an expected call of FindAppUsageEvent.
func (mr *MockAppUsageEventFinderMockRecorder) FindAppUsageEvent(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAppUsageEvent", reflect.TypeOf((*MockAppUsageEventFinder)(nil).FindAppUsageEvent), varargs...)
}

// MockAppUsageEventInserter is a mock of AppUsageEventInserter interface.
type MockAppUsageEventInserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppUsageEventInserterMockRecorder
}

// MockAppUsageEventInserterMockRecorder is the mock recorder for MockAppUsageEventInserter.
type MockAppUsageEventInserterMockRecorder struct {
	mock *MockAppUsageEventInserter
}

// NewMockAppUsageEventInserter creates a new mock instance.
func NewMockAppUsageEventInserter(ctrl *gomock.Controller) *MockAppUsageEventInserter {
	mock := &MockAppUsageEventInserter{ctrl: ctrl}
	mock.recorder = &MockAppUsageEventInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUsageEventInserter) EXPECT() *MockAppUsageEventInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockAppUsageEventInserter) Insert(o *models.AppUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAppUsageEventInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAppUsageEventInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockAppUsageEventUpdater is a mock of AppUsageEventUpdater interface.
type MockAppUsageEventUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAppUsageEventUpdaterMockRecorder
}

// MockAppUsageEventUpdaterMockRecorder is the mock recorder for MockAppUsageEventUpdater.
type MockAppUsageEventUpdaterMockRecorder struct {
	mock *MockAppUsageEventUpdater
}

// NewMockAppUsageEventUpdater creates a new mock instance.
func NewMockAppUsageEventUpdater(ctrl *gomock.Controller) *MockAppUsageEventUpdater {
	mock := &MockAppUsageEventUpdater{ctrl: ctrl}
	mock.recorder = &MockAppUsageEventUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUsageEventUpdater) EXPECT() *MockAppUsageEventUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockAppUsageEventUpdater) Update(o *models.AppUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAppUsageEventUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppUsageEventUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockAppUsageEventUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockAppUsageEventUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockAppUsageEventUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockAppUsageEventUpdater) UpdateAllSlice(o models.AppUsageEventSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockAppUsageEventUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockAppUsageEventUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockAppUsageEventDeleter is a mock of AppUsageEventDeleter interface.
type MockAppUsageEventDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockAppUsageEventDeleterMockRecorder
}

// MockAppUsageEventDeleterMockRecorder is the mock recorder for MockAppUsageEventDeleter.
type MockAppUsageEventDeleterMockRecorder struct {
	mock *MockAppUsageEventDeleter
}

// NewMockAppUsageEventDeleter creates a new mock instance.
func NewMockAppUsageEventDeleter(ctrl *gomock.Controller) *MockAppUsageEventDeleter {
	mock := &MockAppUsageEventDeleter{ctrl: ctrl}
	mock.recorder = &MockAppUsageEventDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUsageEventDeleter) EXPECT() *MockAppUsageEventDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAppUsageEventDeleter) Delete(o *models.AppUsageEvent, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAppUsageEventDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppUsageEventDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockAppUsageEventDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockAppUsageEventDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockAppUsageEventDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockAppUsageEventDeleter) DeleteAllSlice(o models.AppUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockAppUsageEventDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockAppUsageEventDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockAppUsageEventReloader is a mock of AppUsageEventReloader interface.
type MockAppUsageEventReloader struct {
	ctrl     *gomock.Controller
	recorder *MockAppUsageEventReloaderMockRecorder
}

// MockAppUsageEventReloaderMockRecorder is the mock recorder for MockAppUsageEventReloader.
type MockAppUsageEventReloaderMockRecorder struct {
	mock *MockAppUsageEventReloader
}

// NewMockAppUsageEventReloader creates a new mock instance.
func NewMockAppUsageEventReloader(ctrl *gomock.Controller) *MockAppUsageEventReloader {
	mock := &MockAppUsageEventReloader{ctrl: ctrl}
	mock.recorder = &MockAppUsageEventReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUsageEventReloader) EXPECT() *MockAppUsageEventReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockAppUsageEventReloader) Reload(o *models.AppUsageEvent, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockAppUsageEventReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockAppUsageEventReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockAppUsageEventReloader) ReloadAll(o *models.AppUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockAppUsageEventReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockAppUsageEventReloader)(nil).ReloadAll), o, ctx, exec)
}
