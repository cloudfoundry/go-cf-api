//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_plan_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockServicePlanAnnotationUpserter is a mock of ServicePlanAnnotationUpserter interface.
type MockServicePlanAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanAnnotationUpserterMockRecorder
}

// MockServicePlanAnnotationUpserterMockRecorder is the mock recorder for MockServicePlanAnnotationUpserter.
type MockServicePlanAnnotationUpserterMockRecorder struct {
	mock *MockServicePlanAnnotationUpserter
}

// NewMockServicePlanAnnotationUpserter creates a new mock instance.
func NewMockServicePlanAnnotationUpserter(ctrl *gomock.Controller) *MockServicePlanAnnotationUpserter {
	mock := &MockServicePlanAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockServicePlanAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanAnnotationUpserter) EXPECT() *MockServicePlanAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServicePlanAnnotationUpserter) Upsert(o *models.ServicePlanAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServicePlanAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServicePlanAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServicePlanAnnotationFinisher is a mock of ServicePlanAnnotationFinisher interface.
type MockServicePlanAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanAnnotationFinisherMockRecorder
}

// MockServicePlanAnnotationFinisherMockRecorder is the mock recorder for MockServicePlanAnnotationFinisher.
type MockServicePlanAnnotationFinisherMockRecorder struct {
	mock *MockServicePlanAnnotationFinisher
}

// NewMockServicePlanAnnotationFinisher creates a new mock instance.
func NewMockServicePlanAnnotationFinisher(ctrl *gomock.Controller) *MockServicePlanAnnotationFinisher {
	mock := &MockServicePlanAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockServicePlanAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanAnnotationFinisher) EXPECT() *MockServicePlanAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServicePlanAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServicePlanAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServicePlanAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServicePlanAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServicePlanAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServicePlanAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServicePlanAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServicePlanAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServicePlanAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServicePlanAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServicePlanAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServicePlanAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServicePlanAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServicePlanAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServicePlanAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServicePlanAnnotationFinisher)(nil).One), ctx, exec)
}

// MockServicePlanAnnotationFinder is a mock of ServicePlanAnnotationFinder interface.
type MockServicePlanAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanAnnotationFinderMockRecorder
}

// MockServicePlanAnnotationFinderMockRecorder is the mock recorder for MockServicePlanAnnotationFinder.
type MockServicePlanAnnotationFinderMockRecorder struct {
	mock *MockServicePlanAnnotationFinder
}

// NewMockServicePlanAnnotationFinder creates a new mock instance.
func NewMockServicePlanAnnotationFinder(ctrl *gomock.Controller) *MockServicePlanAnnotationFinder {
	mock := &MockServicePlanAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockServicePlanAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanAnnotationFinder) EXPECT() *MockServicePlanAnnotationFinderMockRecorder {
	return m.recorder
}

// FindServicePlanAnnotation mocks base method.
func (m *MockServicePlanAnnotationFinder) FindServicePlanAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServicePlanAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServicePlanAnnotation", varargs...)
	ret0, _ := ret[0].(*models.ServicePlanAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServicePlanAnnotation indicates an expected call of FindServicePlanAnnotation.
func (mr *MockServicePlanAnnotationFinderMockRecorder) FindServicePlanAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServicePlanAnnotation", reflect.TypeOf((*MockServicePlanAnnotationFinder)(nil).FindServicePlanAnnotation), varargs...)
}

// MockServicePlanAnnotationInserter is a mock of ServicePlanAnnotationInserter interface.
type MockServicePlanAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanAnnotationInserterMockRecorder
}

// MockServicePlanAnnotationInserterMockRecorder is the mock recorder for MockServicePlanAnnotationInserter.
type MockServicePlanAnnotationInserterMockRecorder struct {
	mock *MockServicePlanAnnotationInserter
}

// NewMockServicePlanAnnotationInserter creates a new mock instance.
func NewMockServicePlanAnnotationInserter(ctrl *gomock.Controller) *MockServicePlanAnnotationInserter {
	mock := &MockServicePlanAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockServicePlanAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanAnnotationInserter) EXPECT() *MockServicePlanAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServicePlanAnnotationInserter) Insert(o *models.ServicePlanAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServicePlanAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServicePlanAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServicePlanAnnotationUpdater is a mock of ServicePlanAnnotationUpdater interface.
type MockServicePlanAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanAnnotationUpdaterMockRecorder
}

// MockServicePlanAnnotationUpdaterMockRecorder is the mock recorder for MockServicePlanAnnotationUpdater.
type MockServicePlanAnnotationUpdaterMockRecorder struct {
	mock *MockServicePlanAnnotationUpdater
}

// NewMockServicePlanAnnotationUpdater creates a new mock instance.
func NewMockServicePlanAnnotationUpdater(ctrl *gomock.Controller) *MockServicePlanAnnotationUpdater {
	mock := &MockServicePlanAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockServicePlanAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanAnnotationUpdater) EXPECT() *MockServicePlanAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServicePlanAnnotationUpdater) Update(o *models.ServicePlanAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServicePlanAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServicePlanAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServicePlanAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServicePlanAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServicePlanAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServicePlanAnnotationUpdater) UpdateAllSlice(o models.ServicePlanAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServicePlanAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServicePlanAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServicePlanAnnotationDeleter is a mock of ServicePlanAnnotationDeleter interface.
type MockServicePlanAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanAnnotationDeleterMockRecorder
}

// MockServicePlanAnnotationDeleterMockRecorder is the mock recorder for MockServicePlanAnnotationDeleter.
type MockServicePlanAnnotationDeleterMockRecorder struct {
	mock *MockServicePlanAnnotationDeleter
}

// NewMockServicePlanAnnotationDeleter creates a new mock instance.
func NewMockServicePlanAnnotationDeleter(ctrl *gomock.Controller) *MockServicePlanAnnotationDeleter {
	mock := &MockServicePlanAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockServicePlanAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanAnnotationDeleter) EXPECT() *MockServicePlanAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServicePlanAnnotationDeleter) Delete(o *models.ServicePlanAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServicePlanAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServicePlanAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServicePlanAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServicePlanAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServicePlanAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServicePlanAnnotationDeleter) DeleteAllSlice(o models.ServicePlanAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServicePlanAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServicePlanAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServicePlanAnnotationReloader is a mock of ServicePlanAnnotationReloader interface.
type MockServicePlanAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanAnnotationReloaderMockRecorder
}

// MockServicePlanAnnotationReloaderMockRecorder is the mock recorder for MockServicePlanAnnotationReloader.
type MockServicePlanAnnotationReloaderMockRecorder struct {
	mock *MockServicePlanAnnotationReloader
}

// NewMockServicePlanAnnotationReloader creates a new mock instance.
func NewMockServicePlanAnnotationReloader(ctrl *gomock.Controller) *MockServicePlanAnnotationReloader {
	mock := &MockServicePlanAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockServicePlanAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanAnnotationReloader) EXPECT() *MockServicePlanAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServicePlanAnnotationReloader) Reload(o *models.ServicePlanAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServicePlanAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServicePlanAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServicePlanAnnotationReloader) ReloadAll(o *models.ServicePlanAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServicePlanAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServicePlanAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
