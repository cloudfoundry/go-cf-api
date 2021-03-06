//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_key_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockServiceKeyAnnotationUpserter is a mock of ServiceKeyAnnotationUpserter interface.
type MockServiceKeyAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceKeyAnnotationUpserterMockRecorder
}

// MockServiceKeyAnnotationUpserterMockRecorder is the mock recorder for MockServiceKeyAnnotationUpserter.
type MockServiceKeyAnnotationUpserterMockRecorder struct {
	mock *MockServiceKeyAnnotationUpserter
}

// NewMockServiceKeyAnnotationUpserter creates a new mock instance.
func NewMockServiceKeyAnnotationUpserter(ctrl *gomock.Controller) *MockServiceKeyAnnotationUpserter {
	mock := &MockServiceKeyAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceKeyAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceKeyAnnotationUpserter) EXPECT() *MockServiceKeyAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceKeyAnnotationUpserter) Upsert(o *models.ServiceKeyAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceKeyAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceKeyAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceKeyAnnotationFinisher is a mock of ServiceKeyAnnotationFinisher interface.
type MockServiceKeyAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceKeyAnnotationFinisherMockRecorder
}

// MockServiceKeyAnnotationFinisherMockRecorder is the mock recorder for MockServiceKeyAnnotationFinisher.
type MockServiceKeyAnnotationFinisherMockRecorder struct {
	mock *MockServiceKeyAnnotationFinisher
}

// NewMockServiceKeyAnnotationFinisher creates a new mock instance.
func NewMockServiceKeyAnnotationFinisher(ctrl *gomock.Controller) *MockServiceKeyAnnotationFinisher {
	mock := &MockServiceKeyAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceKeyAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceKeyAnnotationFinisher) EXPECT() *MockServiceKeyAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceKeyAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceKeyAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceKeyAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceKeyAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceKeyAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceKeyAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceKeyAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceKeyAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceKeyAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceKeyAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceKeyAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceKeyAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceKeyAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceKeyAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceKeyAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceKeyAnnotationFinisher)(nil).One), ctx, exec)
}

// MockServiceKeyAnnotationFinder is a mock of ServiceKeyAnnotationFinder interface.
type MockServiceKeyAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceKeyAnnotationFinderMockRecorder
}

// MockServiceKeyAnnotationFinderMockRecorder is the mock recorder for MockServiceKeyAnnotationFinder.
type MockServiceKeyAnnotationFinderMockRecorder struct {
	mock *MockServiceKeyAnnotationFinder
}

// NewMockServiceKeyAnnotationFinder creates a new mock instance.
func NewMockServiceKeyAnnotationFinder(ctrl *gomock.Controller) *MockServiceKeyAnnotationFinder {
	mock := &MockServiceKeyAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockServiceKeyAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceKeyAnnotationFinder) EXPECT() *MockServiceKeyAnnotationFinderMockRecorder {
	return m.recorder
}

// FindServiceKeyAnnotation mocks base method.
func (m *MockServiceKeyAnnotationFinder) FindServiceKeyAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceKeyAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceKeyAnnotation", varargs...)
	ret0, _ := ret[0].(*models.ServiceKeyAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceKeyAnnotation indicates an expected call of FindServiceKeyAnnotation.
func (mr *MockServiceKeyAnnotationFinderMockRecorder) FindServiceKeyAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceKeyAnnotation", reflect.TypeOf((*MockServiceKeyAnnotationFinder)(nil).FindServiceKeyAnnotation), varargs...)
}

// MockServiceKeyAnnotationInserter is a mock of ServiceKeyAnnotationInserter interface.
type MockServiceKeyAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceKeyAnnotationInserterMockRecorder
}

// MockServiceKeyAnnotationInserterMockRecorder is the mock recorder for MockServiceKeyAnnotationInserter.
type MockServiceKeyAnnotationInserterMockRecorder struct {
	mock *MockServiceKeyAnnotationInserter
}

// NewMockServiceKeyAnnotationInserter creates a new mock instance.
func NewMockServiceKeyAnnotationInserter(ctrl *gomock.Controller) *MockServiceKeyAnnotationInserter {
	mock := &MockServiceKeyAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockServiceKeyAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceKeyAnnotationInserter) EXPECT() *MockServiceKeyAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceKeyAnnotationInserter) Insert(o *models.ServiceKeyAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceKeyAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceKeyAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceKeyAnnotationUpdater is a mock of ServiceKeyAnnotationUpdater interface.
type MockServiceKeyAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceKeyAnnotationUpdaterMockRecorder
}

// MockServiceKeyAnnotationUpdaterMockRecorder is the mock recorder for MockServiceKeyAnnotationUpdater.
type MockServiceKeyAnnotationUpdaterMockRecorder struct {
	mock *MockServiceKeyAnnotationUpdater
}

// NewMockServiceKeyAnnotationUpdater creates a new mock instance.
func NewMockServiceKeyAnnotationUpdater(ctrl *gomock.Controller) *MockServiceKeyAnnotationUpdater {
	mock := &MockServiceKeyAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceKeyAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceKeyAnnotationUpdater) EXPECT() *MockServiceKeyAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceKeyAnnotationUpdater) Update(o *models.ServiceKeyAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceKeyAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceKeyAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceKeyAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceKeyAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceKeyAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceKeyAnnotationUpdater) UpdateAllSlice(o models.ServiceKeyAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceKeyAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceKeyAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceKeyAnnotationDeleter is a mock of ServiceKeyAnnotationDeleter interface.
type MockServiceKeyAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceKeyAnnotationDeleterMockRecorder
}

// MockServiceKeyAnnotationDeleterMockRecorder is the mock recorder for MockServiceKeyAnnotationDeleter.
type MockServiceKeyAnnotationDeleterMockRecorder struct {
	mock *MockServiceKeyAnnotationDeleter
}

// NewMockServiceKeyAnnotationDeleter creates a new mock instance.
func NewMockServiceKeyAnnotationDeleter(ctrl *gomock.Controller) *MockServiceKeyAnnotationDeleter {
	mock := &MockServiceKeyAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceKeyAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceKeyAnnotationDeleter) EXPECT() *MockServiceKeyAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceKeyAnnotationDeleter) Delete(o *models.ServiceKeyAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceKeyAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceKeyAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceKeyAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceKeyAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceKeyAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceKeyAnnotationDeleter) DeleteAllSlice(o models.ServiceKeyAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceKeyAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceKeyAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceKeyAnnotationReloader is a mock of ServiceKeyAnnotationReloader interface.
type MockServiceKeyAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceKeyAnnotationReloaderMockRecorder
}

// MockServiceKeyAnnotationReloaderMockRecorder is the mock recorder for MockServiceKeyAnnotationReloader.
type MockServiceKeyAnnotationReloaderMockRecorder struct {
	mock *MockServiceKeyAnnotationReloader
}

// NewMockServiceKeyAnnotationReloader creates a new mock instance.
func NewMockServiceKeyAnnotationReloader(ctrl *gomock.Controller) *MockServiceKeyAnnotationReloader {
	mock := &MockServiceKeyAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockServiceKeyAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceKeyAnnotationReloader) EXPECT() *MockServiceKeyAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceKeyAnnotationReloader) Reload(o *models.ServiceKeyAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceKeyAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceKeyAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceKeyAnnotationReloader) ReloadAll(o *models.ServiceKeyAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceKeyAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceKeyAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
