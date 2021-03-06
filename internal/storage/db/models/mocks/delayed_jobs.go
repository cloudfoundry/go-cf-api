//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_delayed_jobs.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockDelayedJobUpserter is a mock of DelayedJobUpserter interface.
type MockDelayedJobUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockDelayedJobUpserterMockRecorder
}

// MockDelayedJobUpserterMockRecorder is the mock recorder for MockDelayedJobUpserter.
type MockDelayedJobUpserterMockRecorder struct {
	mock *MockDelayedJobUpserter
}

// NewMockDelayedJobUpserter creates a new mock instance.
func NewMockDelayedJobUpserter(ctrl *gomock.Controller) *MockDelayedJobUpserter {
	mock := &MockDelayedJobUpserter{ctrl: ctrl}
	mock.recorder = &MockDelayedJobUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelayedJobUpserter) EXPECT() *MockDelayedJobUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockDelayedJobUpserter) Upsert(o *models.DelayedJob, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockDelayedJobUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDelayedJobUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockDelayedJobFinisher is a mock of DelayedJobFinisher interface.
type MockDelayedJobFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockDelayedJobFinisherMockRecorder
}

// MockDelayedJobFinisherMockRecorder is the mock recorder for MockDelayedJobFinisher.
type MockDelayedJobFinisherMockRecorder struct {
	mock *MockDelayedJobFinisher
}

// NewMockDelayedJobFinisher creates a new mock instance.
func NewMockDelayedJobFinisher(ctrl *gomock.Controller) *MockDelayedJobFinisher {
	mock := &MockDelayedJobFinisher{ctrl: ctrl}
	mock.recorder = &MockDelayedJobFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelayedJobFinisher) EXPECT() *MockDelayedJobFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockDelayedJobFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.DelayedJobSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.DelayedJobSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockDelayedJobFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockDelayedJobFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockDelayedJobFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDelayedJobFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDelayedJobFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockDelayedJobFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDelayedJobFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDelayedJobFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockDelayedJobFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.DelayedJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.DelayedJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockDelayedJobFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockDelayedJobFinisher)(nil).One), ctx, exec)
}

// MockDelayedJobFinder is a mock of DelayedJobFinder interface.
type MockDelayedJobFinder struct {
	ctrl     *gomock.Controller
	recorder *MockDelayedJobFinderMockRecorder
}

// MockDelayedJobFinderMockRecorder is the mock recorder for MockDelayedJobFinder.
type MockDelayedJobFinderMockRecorder struct {
	mock *MockDelayedJobFinder
}

// NewMockDelayedJobFinder creates a new mock instance.
func NewMockDelayedJobFinder(ctrl *gomock.Controller) *MockDelayedJobFinder {
	mock := &MockDelayedJobFinder{ctrl: ctrl}
	mock.recorder = &MockDelayedJobFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelayedJobFinder) EXPECT() *MockDelayedJobFinderMockRecorder {
	return m.recorder
}

// FindDelayedJob mocks base method.
func (m *MockDelayedJobFinder) FindDelayedJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.DelayedJob, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindDelayedJob", varargs...)
	ret0, _ := ret[0].(*models.DelayedJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDelayedJob indicates an expected call of FindDelayedJob.
func (mr *MockDelayedJobFinderMockRecorder) FindDelayedJob(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDelayedJob", reflect.TypeOf((*MockDelayedJobFinder)(nil).FindDelayedJob), varargs...)
}

// MockDelayedJobInserter is a mock of DelayedJobInserter interface.
type MockDelayedJobInserter struct {
	ctrl     *gomock.Controller
	recorder *MockDelayedJobInserterMockRecorder
}

// MockDelayedJobInserterMockRecorder is the mock recorder for MockDelayedJobInserter.
type MockDelayedJobInserterMockRecorder struct {
	mock *MockDelayedJobInserter
}

// NewMockDelayedJobInserter creates a new mock instance.
func NewMockDelayedJobInserter(ctrl *gomock.Controller) *MockDelayedJobInserter {
	mock := &MockDelayedJobInserter{ctrl: ctrl}
	mock.recorder = &MockDelayedJobInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelayedJobInserter) EXPECT() *MockDelayedJobInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockDelayedJobInserter) Insert(o *models.DelayedJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockDelayedJobInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDelayedJobInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockDelayedJobUpdater is a mock of DelayedJobUpdater interface.
type MockDelayedJobUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDelayedJobUpdaterMockRecorder
}

// MockDelayedJobUpdaterMockRecorder is the mock recorder for MockDelayedJobUpdater.
type MockDelayedJobUpdaterMockRecorder struct {
	mock *MockDelayedJobUpdater
}

// NewMockDelayedJobUpdater creates a new mock instance.
func NewMockDelayedJobUpdater(ctrl *gomock.Controller) *MockDelayedJobUpdater {
	mock := &MockDelayedJobUpdater{ctrl: ctrl}
	mock.recorder = &MockDelayedJobUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelayedJobUpdater) EXPECT() *MockDelayedJobUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockDelayedJobUpdater) Update(o *models.DelayedJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDelayedJobUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDelayedJobUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockDelayedJobUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockDelayedJobUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockDelayedJobUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockDelayedJobUpdater) UpdateAllSlice(o models.DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockDelayedJobUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockDelayedJobUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockDelayedJobDeleter is a mock of DelayedJobDeleter interface.
type MockDelayedJobDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDelayedJobDeleterMockRecorder
}

// MockDelayedJobDeleterMockRecorder is the mock recorder for MockDelayedJobDeleter.
type MockDelayedJobDeleterMockRecorder struct {
	mock *MockDelayedJobDeleter
}

// NewMockDelayedJobDeleter creates a new mock instance.
func NewMockDelayedJobDeleter(ctrl *gomock.Controller) *MockDelayedJobDeleter {
	mock := &MockDelayedJobDeleter{ctrl: ctrl}
	mock.recorder = &MockDelayedJobDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelayedJobDeleter) EXPECT() *MockDelayedJobDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDelayedJobDeleter) Delete(o *models.DelayedJob, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDelayedJobDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDelayedJobDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockDelayedJobDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockDelayedJobDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockDelayedJobDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockDelayedJobDeleter) DeleteAllSlice(o models.DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockDelayedJobDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockDelayedJobDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockDelayedJobReloader is a mock of DelayedJobReloader interface.
type MockDelayedJobReloader struct {
	ctrl     *gomock.Controller
	recorder *MockDelayedJobReloaderMockRecorder
}

// MockDelayedJobReloaderMockRecorder is the mock recorder for MockDelayedJobReloader.
type MockDelayedJobReloaderMockRecorder struct {
	mock *MockDelayedJobReloader
}

// NewMockDelayedJobReloader creates a new mock instance.
func NewMockDelayedJobReloader(ctrl *gomock.Controller) *MockDelayedJobReloader {
	mock := &MockDelayedJobReloader{ctrl: ctrl}
	mock.recorder = &MockDelayedJobReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelayedJobReloader) EXPECT() *MockDelayedJobReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockDelayedJobReloader) Reload(o *models.DelayedJob, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockDelayedJobReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockDelayedJobReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockDelayedJobReloader) ReloadAll(o *models.DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockDelayedJobReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockDelayedJobReloader)(nil).ReloadAll), o, ctx, exec)
}
