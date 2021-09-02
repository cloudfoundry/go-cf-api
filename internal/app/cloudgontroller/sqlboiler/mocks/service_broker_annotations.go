// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_broker_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockServiceBrokerAnnotationFinisher is a mock of ServiceBrokerAnnotationFinisher interface.
type MockServiceBrokerAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerAnnotationFinisherMockRecorder
}

// MockServiceBrokerAnnotationFinisherMockRecorder is the mock recorder for MockServiceBrokerAnnotationFinisher.
type MockServiceBrokerAnnotationFinisherMockRecorder struct {
	mock *MockServiceBrokerAnnotationFinisher
}

// NewMockServiceBrokerAnnotationFinisher creates a new mock instance.
func NewMockServiceBrokerAnnotationFinisher(ctrl *gomock.Controller) *MockServiceBrokerAnnotationFinisher {
	mock := &MockServiceBrokerAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerAnnotationFinisher) EXPECT() *MockServiceBrokerAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceBrokerAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceBrokerAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceBrokerAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceBrokerAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceBrokerAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceBrokerAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceBrokerAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceBrokerAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceBrokerAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceBrokerAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceBrokerAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceBrokerAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceBrokerAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceBrokerAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceBrokerAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceBrokerAnnotationFinisher)(nil).One), ctx, exec)
}

// MockServiceBrokerAnnotationFinder is a mock of ServiceBrokerAnnotationFinder interface.
type MockServiceBrokerAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerAnnotationFinderMockRecorder
}

// MockServiceBrokerAnnotationFinderMockRecorder is the mock recorder for MockServiceBrokerAnnotationFinder.
type MockServiceBrokerAnnotationFinderMockRecorder struct {
	mock *MockServiceBrokerAnnotationFinder
}

// NewMockServiceBrokerAnnotationFinder creates a new mock instance.
func NewMockServiceBrokerAnnotationFinder(ctrl *gomock.Controller) *MockServiceBrokerAnnotationFinder {
	mock := &MockServiceBrokerAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerAnnotationFinder) EXPECT() *MockServiceBrokerAnnotationFinderMockRecorder {
	return m.recorder
}

// FindServiceBrokerAnnotation mocks base method.
func (m *MockServiceBrokerAnnotationFinder) FindServiceBrokerAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceBrokerAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceBrokerAnnotation", varargs...)
	ret0, _ := ret[0].(*models.ServiceBrokerAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceBrokerAnnotation indicates an expected call of FindServiceBrokerAnnotation.
func (mr *MockServiceBrokerAnnotationFinderMockRecorder) FindServiceBrokerAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceBrokerAnnotation", reflect.TypeOf((*MockServiceBrokerAnnotationFinder)(nil).FindServiceBrokerAnnotation), varargs...)
}

// MockServiceBrokerAnnotationInserter is a mock of ServiceBrokerAnnotationInserter interface.
type MockServiceBrokerAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerAnnotationInserterMockRecorder
}

// MockServiceBrokerAnnotationInserterMockRecorder is the mock recorder for MockServiceBrokerAnnotationInserter.
type MockServiceBrokerAnnotationInserterMockRecorder struct {
	mock *MockServiceBrokerAnnotationInserter
}

// NewMockServiceBrokerAnnotationInserter creates a new mock instance.
func NewMockServiceBrokerAnnotationInserter(ctrl *gomock.Controller) *MockServiceBrokerAnnotationInserter {
	mock := &MockServiceBrokerAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerAnnotationInserter) EXPECT() *MockServiceBrokerAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceBrokerAnnotationInserter) Insert(o *models.ServiceBrokerAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceBrokerAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceBrokerAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceBrokerAnnotationUpdater is a mock of ServiceBrokerAnnotationUpdater interface.
type MockServiceBrokerAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerAnnotationUpdaterMockRecorder
}

// MockServiceBrokerAnnotationUpdaterMockRecorder is the mock recorder for MockServiceBrokerAnnotationUpdater.
type MockServiceBrokerAnnotationUpdaterMockRecorder struct {
	mock *MockServiceBrokerAnnotationUpdater
}

// NewMockServiceBrokerAnnotationUpdater creates a new mock instance.
func NewMockServiceBrokerAnnotationUpdater(ctrl *gomock.Controller) *MockServiceBrokerAnnotationUpdater {
	mock := &MockServiceBrokerAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerAnnotationUpdater) EXPECT() *MockServiceBrokerAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceBrokerAnnotationUpdater) Update(o *models.ServiceBrokerAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceBrokerAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceBrokerAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceBrokerAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceBrokerAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceBrokerAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceBrokerAnnotationUpdater) UpdateAllSlice(o models.ServiceBrokerAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceBrokerAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceBrokerAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceBrokerAnnotationDeleter is a mock of ServiceBrokerAnnotationDeleter interface.
type MockServiceBrokerAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerAnnotationDeleterMockRecorder
}

// MockServiceBrokerAnnotationDeleterMockRecorder is the mock recorder for MockServiceBrokerAnnotationDeleter.
type MockServiceBrokerAnnotationDeleterMockRecorder struct {
	mock *MockServiceBrokerAnnotationDeleter
}

// NewMockServiceBrokerAnnotationDeleter creates a new mock instance.
func NewMockServiceBrokerAnnotationDeleter(ctrl *gomock.Controller) *MockServiceBrokerAnnotationDeleter {
	mock := &MockServiceBrokerAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerAnnotationDeleter) EXPECT() *MockServiceBrokerAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceBrokerAnnotationDeleter) Delete(o *models.ServiceBrokerAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceBrokerAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceBrokerAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceBrokerAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceBrokerAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceBrokerAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceBrokerAnnotationDeleter) DeleteAllSlice(o models.ServiceBrokerAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceBrokerAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceBrokerAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceBrokerAnnotationReloader is a mock of ServiceBrokerAnnotationReloader interface.
type MockServiceBrokerAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBrokerAnnotationReloaderMockRecorder
}

// MockServiceBrokerAnnotationReloaderMockRecorder is the mock recorder for MockServiceBrokerAnnotationReloader.
type MockServiceBrokerAnnotationReloaderMockRecorder struct {
	mock *MockServiceBrokerAnnotationReloader
}

// NewMockServiceBrokerAnnotationReloader creates a new mock instance.
func NewMockServiceBrokerAnnotationReloader(ctrl *gomock.Controller) *MockServiceBrokerAnnotationReloader {
	mock := &MockServiceBrokerAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockServiceBrokerAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBrokerAnnotationReloader) EXPECT() *MockServiceBrokerAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceBrokerAnnotationReloader) Reload(o *models.ServiceBrokerAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceBrokerAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceBrokerAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceBrokerAnnotationReloader) ReloadAll(o *models.ServiceBrokerAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceBrokerAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceBrokerAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
