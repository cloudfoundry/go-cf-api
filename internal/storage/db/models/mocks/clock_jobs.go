//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_clock_jobs.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockClockJobUpserter is a mock of ClockJobUpserter interface.
type MockClockJobUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockClockJobUpserterMockRecorder
}

// MockClockJobUpserterMockRecorder is the mock recorder for MockClockJobUpserter.
type MockClockJobUpserterMockRecorder struct {
	mock *MockClockJobUpserter
}

// NewMockClockJobUpserter creates a new mock instance.
func NewMockClockJobUpserter(ctrl *gomock.Controller) *MockClockJobUpserter {
	mock := &MockClockJobUpserter{ctrl: ctrl}
	mock.recorder = &MockClockJobUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockJobUpserter) EXPECT() *MockClockJobUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockClockJobUpserter) Upsert(o *models.ClockJob, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockClockJobUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockClockJobUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockClockJobFinisher is a mock of ClockJobFinisher interface.
type MockClockJobFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockClockJobFinisherMockRecorder
}

// MockClockJobFinisherMockRecorder is the mock recorder for MockClockJobFinisher.
type MockClockJobFinisherMockRecorder struct {
	mock *MockClockJobFinisher
}

// NewMockClockJobFinisher creates a new mock instance.
func NewMockClockJobFinisher(ctrl *gomock.Controller) *MockClockJobFinisher {
	mock := &MockClockJobFinisher{ctrl: ctrl}
	mock.recorder = &MockClockJobFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockJobFinisher) EXPECT() *MockClockJobFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockClockJobFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ClockJobSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ClockJobSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockClockJobFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockClockJobFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockClockJobFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockClockJobFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockClockJobFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockClockJobFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockClockJobFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockClockJobFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockClockJobFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ClockJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ClockJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockClockJobFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockClockJobFinisher)(nil).One), ctx, exec)
}

// MockClockJobFinder is a mock of ClockJobFinder interface.
type MockClockJobFinder struct {
	ctrl     *gomock.Controller
	recorder *MockClockJobFinderMockRecorder
}

// MockClockJobFinderMockRecorder is the mock recorder for MockClockJobFinder.
type MockClockJobFinderMockRecorder struct {
	mock *MockClockJobFinder
}

// NewMockClockJobFinder creates a new mock instance.
func NewMockClockJobFinder(ctrl *gomock.Controller) *MockClockJobFinder {
	mock := &MockClockJobFinder{ctrl: ctrl}
	mock.recorder = &MockClockJobFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockJobFinder) EXPECT() *MockClockJobFinderMockRecorder {
	return m.recorder
}

// FindClockJob mocks base method.
func (m *MockClockJobFinder) FindClockJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ClockJob, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindClockJob", varargs...)
	ret0, _ := ret[0].(*models.ClockJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindClockJob indicates an expected call of FindClockJob.
func (mr *MockClockJobFinderMockRecorder) FindClockJob(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindClockJob", reflect.TypeOf((*MockClockJobFinder)(nil).FindClockJob), varargs...)
}

// MockClockJobInserter is a mock of ClockJobInserter interface.
type MockClockJobInserter struct {
	ctrl     *gomock.Controller
	recorder *MockClockJobInserterMockRecorder
}

// MockClockJobInserterMockRecorder is the mock recorder for MockClockJobInserter.
type MockClockJobInserterMockRecorder struct {
	mock *MockClockJobInserter
}

// NewMockClockJobInserter creates a new mock instance.
func NewMockClockJobInserter(ctrl *gomock.Controller) *MockClockJobInserter {
	mock := &MockClockJobInserter{ctrl: ctrl}
	mock.recorder = &MockClockJobInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockJobInserter) EXPECT() *MockClockJobInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockClockJobInserter) Insert(o *models.ClockJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockClockJobInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClockJobInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockClockJobUpdater is a mock of ClockJobUpdater interface.
type MockClockJobUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockClockJobUpdaterMockRecorder
}

// MockClockJobUpdaterMockRecorder is the mock recorder for MockClockJobUpdater.
type MockClockJobUpdaterMockRecorder struct {
	mock *MockClockJobUpdater
}

// NewMockClockJobUpdater creates a new mock instance.
func NewMockClockJobUpdater(ctrl *gomock.Controller) *MockClockJobUpdater {
	mock := &MockClockJobUpdater{ctrl: ctrl}
	mock.recorder = &MockClockJobUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockJobUpdater) EXPECT() *MockClockJobUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockClockJobUpdater) Update(o *models.ClockJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockClockJobUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClockJobUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockClockJobUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockClockJobUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockClockJobUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockClockJobUpdater) UpdateAllSlice(o models.ClockJobSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockClockJobUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockClockJobUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockClockJobDeleter is a mock of ClockJobDeleter interface.
type MockClockJobDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockClockJobDeleterMockRecorder
}

// MockClockJobDeleterMockRecorder is the mock recorder for MockClockJobDeleter.
type MockClockJobDeleterMockRecorder struct {
	mock *MockClockJobDeleter
}

// NewMockClockJobDeleter creates a new mock instance.
func NewMockClockJobDeleter(ctrl *gomock.Controller) *MockClockJobDeleter {
	mock := &MockClockJobDeleter{ctrl: ctrl}
	mock.recorder = &MockClockJobDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockJobDeleter) EXPECT() *MockClockJobDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockClockJobDeleter) Delete(o *models.ClockJob, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockClockJobDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClockJobDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockClockJobDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockClockJobDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockClockJobDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockClockJobDeleter) DeleteAllSlice(o models.ClockJobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockClockJobDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockClockJobDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockClockJobReloader is a mock of ClockJobReloader interface.
type MockClockJobReloader struct {
	ctrl     *gomock.Controller
	recorder *MockClockJobReloaderMockRecorder
}

// MockClockJobReloaderMockRecorder is the mock recorder for MockClockJobReloader.
type MockClockJobReloaderMockRecorder struct {
	mock *MockClockJobReloader
}

// NewMockClockJobReloader creates a new mock instance.
func NewMockClockJobReloader(ctrl *gomock.Controller) *MockClockJobReloader {
	mock := &MockClockJobReloader{ctrl: ctrl}
	mock.recorder = &MockClockJobReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClockJobReloader) EXPECT() *MockClockJobReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockClockJobReloader) Reload(o *models.ClockJob, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockClockJobReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockClockJobReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockClockJobReloader) ReloadAll(o *models.ClockJobSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockClockJobReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockClockJobReloader)(nil).ReloadAll), o, ctx, exec)
}
