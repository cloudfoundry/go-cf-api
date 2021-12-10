//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_isolation_segment_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockIsolationSegmentAnnotationUpserter is a mock of IsolationSegmentAnnotationUpserter interface.
type MockIsolationSegmentAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentAnnotationUpserterMockRecorder
}

// MockIsolationSegmentAnnotationUpserterMockRecorder is the mock recorder for MockIsolationSegmentAnnotationUpserter.
type MockIsolationSegmentAnnotationUpserterMockRecorder struct {
	mock *MockIsolationSegmentAnnotationUpserter
}

// NewMockIsolationSegmentAnnotationUpserter creates a new mock instance.
func NewMockIsolationSegmentAnnotationUpserter(ctrl *gomock.Controller) *MockIsolationSegmentAnnotationUpserter {
	mock := &MockIsolationSegmentAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentAnnotationUpserter) EXPECT() *MockIsolationSegmentAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockIsolationSegmentAnnotationUpserter) Upsert(o *models.IsolationSegmentAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockIsolationSegmentAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockIsolationSegmentAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockIsolationSegmentAnnotationFinisher is a mock of IsolationSegmentAnnotationFinisher interface.
type MockIsolationSegmentAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentAnnotationFinisherMockRecorder
}

// MockIsolationSegmentAnnotationFinisherMockRecorder is the mock recorder for MockIsolationSegmentAnnotationFinisher.
type MockIsolationSegmentAnnotationFinisherMockRecorder struct {
	mock *MockIsolationSegmentAnnotationFinisher
}

// NewMockIsolationSegmentAnnotationFinisher creates a new mock instance.
func NewMockIsolationSegmentAnnotationFinisher(ctrl *gomock.Controller) *MockIsolationSegmentAnnotationFinisher {
	mock := &MockIsolationSegmentAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentAnnotationFinisher) EXPECT() *MockIsolationSegmentAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockIsolationSegmentAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.IsolationSegmentAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.IsolationSegmentAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockIsolationSegmentAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIsolationSegmentAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockIsolationSegmentAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIsolationSegmentAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIsolationSegmentAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockIsolationSegmentAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockIsolationSegmentAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIsolationSegmentAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockIsolationSegmentAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.IsolationSegmentAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.IsolationSegmentAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockIsolationSegmentAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockIsolationSegmentAnnotationFinisher)(nil).One), ctx, exec)
}

// MockIsolationSegmentAnnotationFinder is a mock of IsolationSegmentAnnotationFinder interface.
type MockIsolationSegmentAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentAnnotationFinderMockRecorder
}

// MockIsolationSegmentAnnotationFinderMockRecorder is the mock recorder for MockIsolationSegmentAnnotationFinder.
type MockIsolationSegmentAnnotationFinderMockRecorder struct {
	mock *MockIsolationSegmentAnnotationFinder
}

// NewMockIsolationSegmentAnnotationFinder creates a new mock instance.
func NewMockIsolationSegmentAnnotationFinder(ctrl *gomock.Controller) *MockIsolationSegmentAnnotationFinder {
	mock := &MockIsolationSegmentAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentAnnotationFinder) EXPECT() *MockIsolationSegmentAnnotationFinderMockRecorder {
	return m.recorder
}

// FindIsolationSegmentAnnotation mocks base method.
func (m *MockIsolationSegmentAnnotationFinder) FindIsolationSegmentAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.IsolationSegmentAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindIsolationSegmentAnnotation", varargs...)
	ret0, _ := ret[0].(*models.IsolationSegmentAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIsolationSegmentAnnotation indicates an expected call of FindIsolationSegmentAnnotation.
func (mr *MockIsolationSegmentAnnotationFinderMockRecorder) FindIsolationSegmentAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIsolationSegmentAnnotation", reflect.TypeOf((*MockIsolationSegmentAnnotationFinder)(nil).FindIsolationSegmentAnnotation), varargs...)
}

// MockIsolationSegmentAnnotationInserter is a mock of IsolationSegmentAnnotationInserter interface.
type MockIsolationSegmentAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentAnnotationInserterMockRecorder
}

// MockIsolationSegmentAnnotationInserterMockRecorder is the mock recorder for MockIsolationSegmentAnnotationInserter.
type MockIsolationSegmentAnnotationInserterMockRecorder struct {
	mock *MockIsolationSegmentAnnotationInserter
}

// NewMockIsolationSegmentAnnotationInserter creates a new mock instance.
func NewMockIsolationSegmentAnnotationInserter(ctrl *gomock.Controller) *MockIsolationSegmentAnnotationInserter {
	mock := &MockIsolationSegmentAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentAnnotationInserter) EXPECT() *MockIsolationSegmentAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockIsolationSegmentAnnotationInserter) Insert(o *models.IsolationSegmentAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIsolationSegmentAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIsolationSegmentAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockIsolationSegmentAnnotationUpdater is a mock of IsolationSegmentAnnotationUpdater interface.
type MockIsolationSegmentAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentAnnotationUpdaterMockRecorder
}

// MockIsolationSegmentAnnotationUpdaterMockRecorder is the mock recorder for MockIsolationSegmentAnnotationUpdater.
type MockIsolationSegmentAnnotationUpdaterMockRecorder struct {
	mock *MockIsolationSegmentAnnotationUpdater
}

// NewMockIsolationSegmentAnnotationUpdater creates a new mock instance.
func NewMockIsolationSegmentAnnotationUpdater(ctrl *gomock.Controller) *MockIsolationSegmentAnnotationUpdater {
	mock := &MockIsolationSegmentAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentAnnotationUpdater) EXPECT() *MockIsolationSegmentAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockIsolationSegmentAnnotationUpdater) Update(o *models.IsolationSegmentAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIsolationSegmentAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIsolationSegmentAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockIsolationSegmentAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockIsolationSegmentAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockIsolationSegmentAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockIsolationSegmentAnnotationUpdater) UpdateAllSlice(o models.IsolationSegmentAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockIsolationSegmentAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockIsolationSegmentAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockIsolationSegmentAnnotationDeleter is a mock of IsolationSegmentAnnotationDeleter interface.
type MockIsolationSegmentAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentAnnotationDeleterMockRecorder
}

// MockIsolationSegmentAnnotationDeleterMockRecorder is the mock recorder for MockIsolationSegmentAnnotationDeleter.
type MockIsolationSegmentAnnotationDeleterMockRecorder struct {
	mock *MockIsolationSegmentAnnotationDeleter
}

// NewMockIsolationSegmentAnnotationDeleter creates a new mock instance.
func NewMockIsolationSegmentAnnotationDeleter(ctrl *gomock.Controller) *MockIsolationSegmentAnnotationDeleter {
	mock := &MockIsolationSegmentAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentAnnotationDeleter) EXPECT() *MockIsolationSegmentAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIsolationSegmentAnnotationDeleter) Delete(o *models.IsolationSegmentAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIsolationSegmentAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIsolationSegmentAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockIsolationSegmentAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockIsolationSegmentAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockIsolationSegmentAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockIsolationSegmentAnnotationDeleter) DeleteAllSlice(o models.IsolationSegmentAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockIsolationSegmentAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockIsolationSegmentAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockIsolationSegmentAnnotationReloader is a mock of IsolationSegmentAnnotationReloader interface.
type MockIsolationSegmentAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockIsolationSegmentAnnotationReloaderMockRecorder
}

// MockIsolationSegmentAnnotationReloaderMockRecorder is the mock recorder for MockIsolationSegmentAnnotationReloader.
type MockIsolationSegmentAnnotationReloaderMockRecorder struct {
	mock *MockIsolationSegmentAnnotationReloader
}

// NewMockIsolationSegmentAnnotationReloader creates a new mock instance.
func NewMockIsolationSegmentAnnotationReloader(ctrl *gomock.Controller) *MockIsolationSegmentAnnotationReloader {
	mock := &MockIsolationSegmentAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockIsolationSegmentAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIsolationSegmentAnnotationReloader) EXPECT() *MockIsolationSegmentAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockIsolationSegmentAnnotationReloader) Reload(o *models.IsolationSegmentAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockIsolationSegmentAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockIsolationSegmentAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockIsolationSegmentAnnotationReloader) ReloadAll(o *models.IsolationSegmentAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockIsolationSegmentAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockIsolationSegmentAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
