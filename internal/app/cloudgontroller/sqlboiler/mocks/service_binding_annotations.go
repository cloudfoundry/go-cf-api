// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_binding_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockServiceBindingAnnotationFinisher is a mock of ServiceBindingAnnotationFinisher interface.
type MockServiceBindingAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingAnnotationFinisherMockRecorder
}

// MockServiceBindingAnnotationFinisherMockRecorder is the mock recorder for MockServiceBindingAnnotationFinisher.
type MockServiceBindingAnnotationFinisherMockRecorder struct {
	mock *MockServiceBindingAnnotationFinisher
}

// NewMockServiceBindingAnnotationFinisher creates a new mock instance.
func NewMockServiceBindingAnnotationFinisher(ctrl *gomock.Controller) *MockServiceBindingAnnotationFinisher {
	mock := &MockServiceBindingAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceBindingAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingAnnotationFinisher) EXPECT() *MockServiceBindingAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceBindingAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceBindingAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceBindingAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceBindingAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceBindingAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceBindingAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceBindingAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceBindingAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceBindingAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceBindingAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceBindingAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceBindingAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceBindingAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceBindingAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceBindingAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceBindingAnnotationFinisher)(nil).One), ctx, exec)
}

// MockServiceBindingAnnotationFinder is a mock of ServiceBindingAnnotationFinder interface.
type MockServiceBindingAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingAnnotationFinderMockRecorder
}

// MockServiceBindingAnnotationFinderMockRecorder is the mock recorder for MockServiceBindingAnnotationFinder.
type MockServiceBindingAnnotationFinderMockRecorder struct {
	mock *MockServiceBindingAnnotationFinder
}

// NewMockServiceBindingAnnotationFinder creates a new mock instance.
func NewMockServiceBindingAnnotationFinder(ctrl *gomock.Controller) *MockServiceBindingAnnotationFinder {
	mock := &MockServiceBindingAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockServiceBindingAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingAnnotationFinder) EXPECT() *MockServiceBindingAnnotationFinderMockRecorder {
	return m.recorder
}

// FindServiceBindingAnnotation mocks base method.
func (m *MockServiceBindingAnnotationFinder) FindServiceBindingAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceBindingAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceBindingAnnotation", varargs...)
	ret0, _ := ret[0].(*models.ServiceBindingAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceBindingAnnotation indicates an expected call of FindServiceBindingAnnotation.
func (mr *MockServiceBindingAnnotationFinderMockRecorder) FindServiceBindingAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceBindingAnnotation", reflect.TypeOf((*MockServiceBindingAnnotationFinder)(nil).FindServiceBindingAnnotation), varargs...)
}

// MockServiceBindingAnnotationInserter is a mock of ServiceBindingAnnotationInserter interface.
type MockServiceBindingAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingAnnotationInserterMockRecorder
}

// MockServiceBindingAnnotationInserterMockRecorder is the mock recorder for MockServiceBindingAnnotationInserter.
type MockServiceBindingAnnotationInserterMockRecorder struct {
	mock *MockServiceBindingAnnotationInserter
}

// NewMockServiceBindingAnnotationInserter creates a new mock instance.
func NewMockServiceBindingAnnotationInserter(ctrl *gomock.Controller) *MockServiceBindingAnnotationInserter {
	mock := &MockServiceBindingAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockServiceBindingAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingAnnotationInserter) EXPECT() *MockServiceBindingAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceBindingAnnotationInserter) Insert(o *models.ServiceBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceBindingAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceBindingAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceBindingAnnotationUpdater is a mock of ServiceBindingAnnotationUpdater interface.
type MockServiceBindingAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingAnnotationUpdaterMockRecorder
}

// MockServiceBindingAnnotationUpdaterMockRecorder is the mock recorder for MockServiceBindingAnnotationUpdater.
type MockServiceBindingAnnotationUpdaterMockRecorder struct {
	mock *MockServiceBindingAnnotationUpdater
}

// NewMockServiceBindingAnnotationUpdater creates a new mock instance.
func NewMockServiceBindingAnnotationUpdater(ctrl *gomock.Controller) *MockServiceBindingAnnotationUpdater {
	mock := &MockServiceBindingAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceBindingAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingAnnotationUpdater) EXPECT() *MockServiceBindingAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceBindingAnnotationUpdater) Update(o *models.ServiceBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceBindingAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceBindingAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceBindingAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceBindingAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceBindingAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceBindingAnnotationUpdater) UpdateAllSlice(o models.ServiceBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceBindingAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceBindingAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceBindingAnnotationDeleter is a mock of ServiceBindingAnnotationDeleter interface.
type MockServiceBindingAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingAnnotationDeleterMockRecorder
}

// MockServiceBindingAnnotationDeleterMockRecorder is the mock recorder for MockServiceBindingAnnotationDeleter.
type MockServiceBindingAnnotationDeleterMockRecorder struct {
	mock *MockServiceBindingAnnotationDeleter
}

// NewMockServiceBindingAnnotationDeleter creates a new mock instance.
func NewMockServiceBindingAnnotationDeleter(ctrl *gomock.Controller) *MockServiceBindingAnnotationDeleter {
	mock := &MockServiceBindingAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceBindingAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingAnnotationDeleter) EXPECT() *MockServiceBindingAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceBindingAnnotationDeleter) Delete(o *models.ServiceBindingAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceBindingAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceBindingAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceBindingAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceBindingAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceBindingAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceBindingAnnotationDeleter) DeleteAllSlice(o models.ServiceBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceBindingAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceBindingAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceBindingAnnotationReloader is a mock of ServiceBindingAnnotationReloader interface.
type MockServiceBindingAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBindingAnnotationReloaderMockRecorder
}

// MockServiceBindingAnnotationReloaderMockRecorder is the mock recorder for MockServiceBindingAnnotationReloader.
type MockServiceBindingAnnotationReloaderMockRecorder struct {
	mock *MockServiceBindingAnnotationReloader
}

// NewMockServiceBindingAnnotationReloader creates a new mock instance.
func NewMockServiceBindingAnnotationReloader(ctrl *gomock.Controller) *MockServiceBindingAnnotationReloader {
	mock := &MockServiceBindingAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockServiceBindingAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBindingAnnotationReloader) EXPECT() *MockServiceBindingAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceBindingAnnotationReloader) Reload(o *models.ServiceBindingAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceBindingAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceBindingAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceBindingAnnotationReloader) ReloadAll(o *models.ServiceBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceBindingAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceBindingAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
