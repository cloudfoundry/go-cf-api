//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_isolation_segments.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
)

// MockIsolationSegmentUpserter is a mock of IsolationSegmentUpserter interface.
type MockIsolationSegmentUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentUpserterMockRecorder
}

// MockIsolationSegmentUpserterMockRecorder is the mock recorder for MockIsolationSegmentUpserter.
type MockIsolationSegmentUpserterMockRecorder struct {
	mock *MockIsolationSegmentUpserter
}

// NewMockIsolationSegmentUpserter creates a new mock instance.
func NewMockIsolationSegmentUpserter(ctrl *gomock.Controller) *MockIsolationSegmentUpserter {
	mock := &MockIsolationSegmentUpserter{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentUpserter) EXPECT() *MockIsolationSegmentUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockIsolationSegmentUpserter) Upsert(o *models.IsolationSegment, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockIsolationSegmentUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockIsolationSegmentUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockIsolationSegmentFinisher is a mock of IsolationSegmentFinisher interface.
type MockIsolationSegmentFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentFinisherMockRecorder
}

// MockIsolationSegmentFinisherMockRecorder is the mock recorder for MockIsolationSegmentFinisher.
type MockIsolationSegmentFinisherMockRecorder struct {
	mock *MockIsolationSegmentFinisher
}

// NewMockIsolationSegmentFinisher creates a new mock instance.
func NewMockIsolationSegmentFinisher(ctrl *gomock.Controller) *MockIsolationSegmentFinisher {
	mock := &MockIsolationSegmentFinisher{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentFinisher) EXPECT() *MockIsolationSegmentFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockIsolationSegmentFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.IsolationSegmentSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.IsolationSegmentSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockIsolationSegmentFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIsolationSegmentFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockIsolationSegmentFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIsolationSegmentFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIsolationSegmentFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockIsolationSegmentFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockIsolationSegmentFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIsolationSegmentFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockIsolationSegmentFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.IsolationSegment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.IsolationSegment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockIsolationSegmentFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockIsolationSegmentFinisher)(nil).One), ctx, exec)
}

// MockIsolationSegmentFinder is a mock of IsolationSegmentFinder interface.
type MockIsolationSegmentFinder struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentFinderMockRecorder
}

// MockIsolationSegmentFinderMockRecorder is the mock recorder for MockIsolationSegmentFinder.
type MockIsolationSegmentFinderMockRecorder struct {
	mock *MockIsolationSegmentFinder
}

// NewMockIsolationSegmentFinder creates a new mock instance.
func NewMockIsolationSegmentFinder(ctrl *gomock.Controller) *MockIsolationSegmentFinder {
	mock := &MockIsolationSegmentFinder{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentFinder) EXPECT() *MockIsolationSegmentFinderMockRecorder {
	return m.recorder
}

// FindIsolationSegment mocks base method.
func (m *MockIsolationSegmentFinder) FindIsolationSegment(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.IsolationSegment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindIsolationSegment", varargs...)
	ret0, _ := ret[0].(*models.IsolationSegment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIsolationSegment indicates an expected call of FindIsolationSegment.
func (mr *MockIsolationSegmentFinderMockRecorder) FindIsolationSegment(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIsolationSegment", reflect.TypeOf((*MockIsolationSegmentFinder)(nil).FindIsolationSegment), varargs...)
}

// MockIsolationSegmentInserter is a mock of IsolationSegmentInserter interface.
type MockIsolationSegmentInserter struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentInserterMockRecorder
}

// MockIsolationSegmentInserterMockRecorder is the mock recorder for MockIsolationSegmentInserter.
type MockIsolationSegmentInserterMockRecorder struct {
	mock *MockIsolationSegmentInserter
}

// NewMockIsolationSegmentInserter creates a new mock instance.
func NewMockIsolationSegmentInserter(ctrl *gomock.Controller) *MockIsolationSegmentInserter {
	mock := &MockIsolationSegmentInserter{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentInserter) EXPECT() *MockIsolationSegmentInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockIsolationSegmentInserter) Insert(o *models.IsolationSegment, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIsolationSegmentInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIsolationSegmentInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockIsolationSegmentUpdater is a mock of IsolationSegmentUpdater interface.
type MockIsolationSegmentUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentUpdaterMockRecorder
}

// MockIsolationSegmentUpdaterMockRecorder is the mock recorder for MockIsolationSegmentUpdater.
type MockIsolationSegmentUpdaterMockRecorder struct {
	mock *MockIsolationSegmentUpdater
}

// NewMockIsolationSegmentUpdater creates a new mock instance.
func NewMockIsolationSegmentUpdater(ctrl *gomock.Controller) *MockIsolationSegmentUpdater {
	mock := &MockIsolationSegmentUpdater{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentUpdater) EXPECT() *MockIsolationSegmentUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockIsolationSegmentUpdater) Update(o *models.IsolationSegment, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIsolationSegmentUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIsolationSegmentUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockIsolationSegmentUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockIsolationSegmentUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockIsolationSegmentUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockIsolationSegmentUpdater) UpdateAllSlice(o models.IsolationSegmentSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockIsolationSegmentUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockIsolationSegmentUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockIsolationSegmentDeleter is a mock of IsolationSegmentDeleter interface.
type MockIsolationSegmentDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentDeleterMockRecorder
}

// MockIsolationSegmentDeleterMockRecorder is the mock recorder for MockIsolationSegmentDeleter.
type MockIsolationSegmentDeleterMockRecorder struct {
	mock *MockIsolationSegmentDeleter
}

// NewMockIsolationSegmentDeleter creates a new mock instance.
func NewMockIsolationSegmentDeleter(ctrl *gomock.Controller) *MockIsolationSegmentDeleter {
	mock := &MockIsolationSegmentDeleter{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentDeleter) EXPECT() *MockIsolationSegmentDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIsolationSegmentDeleter) Delete(o *models.IsolationSegment, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIsolationSegmentDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIsolationSegmentDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockIsolationSegmentDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockIsolationSegmentDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockIsolationSegmentDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockIsolationSegmentDeleter) DeleteAllSlice(o models.IsolationSegmentSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockIsolationSegmentDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockIsolationSegmentDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockIsolationSegmentReloader is a mock of IsolationSegmentReloader interface.
type MockIsolationSegmentReloader struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentReloaderMockRecorder
}

// MockIsolationSegmentReloaderMockRecorder is the mock recorder for MockIsolationSegmentReloader.
type MockIsolationSegmentReloaderMockRecorder struct {
	mock *MockIsolationSegmentReloader
}

// NewMockIsolationSegmentReloader creates a new mock instance.
func NewMockIsolationSegmentReloader(ctrl *gomock.Controller) *MockIsolationSegmentReloader {
	mock := &MockIsolationSegmentReloader{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentReloader) EXPECT() *MockIsolationSegmentReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockIsolationSegmentReloader) Reload(o *models.IsolationSegment, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockIsolationSegmentReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockIsolationSegmentReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockIsolationSegmentReloader) ReloadAll(o *models.IsolationSegmentSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockIsolationSegmentReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockIsolationSegmentReloader)(nil).ReloadAll), o, ctx, exec)
}
