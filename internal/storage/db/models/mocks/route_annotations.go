//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_route_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockRouteAnnotationUpserter is a mock of RouteAnnotationUpserter interface.
type MockRouteAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockRouteAnnotationUpserterMockRecorder
}

// MockRouteAnnotationUpserterMockRecorder is the mock recorder for MockRouteAnnotationUpserter.
type MockRouteAnnotationUpserterMockRecorder struct {
	mock *MockRouteAnnotationUpserter
}

// NewMockRouteAnnotationUpserter creates a new mock instance.
func NewMockRouteAnnotationUpserter(ctrl *gomock.Controller) *MockRouteAnnotationUpserter {
	mock := &MockRouteAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockRouteAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteAnnotationUpserter) EXPECT() *MockRouteAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockRouteAnnotationUpserter) Upsert(o *models.RouteAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRouteAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRouteAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockRouteAnnotationFinisher is a mock of RouteAnnotationFinisher interface.
type MockRouteAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockRouteAnnotationFinisherMockRecorder
}

// MockRouteAnnotationFinisherMockRecorder is the mock recorder for MockRouteAnnotationFinisher.
type MockRouteAnnotationFinisherMockRecorder struct {
	mock *MockRouteAnnotationFinisher
}

// NewMockRouteAnnotationFinisher creates a new mock instance.
func NewMockRouteAnnotationFinisher(ctrl *gomock.Controller) *MockRouteAnnotationFinisher {
	mock := &MockRouteAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockRouteAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteAnnotationFinisher) EXPECT() *MockRouteAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockRouteAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.RouteAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.RouteAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockRouteAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRouteAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockRouteAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRouteAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRouteAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockRouteAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRouteAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRouteAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockRouteAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.RouteAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.RouteAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockRouteAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockRouteAnnotationFinisher)(nil).One), ctx, exec)
}

// MockRouteAnnotationFinder is a mock of RouteAnnotationFinder interface.
type MockRouteAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockRouteAnnotationFinderMockRecorder
}

// MockRouteAnnotationFinderMockRecorder is the mock recorder for MockRouteAnnotationFinder.
type MockRouteAnnotationFinderMockRecorder struct {
	mock *MockRouteAnnotationFinder
}

// NewMockRouteAnnotationFinder creates a new mock instance.
func NewMockRouteAnnotationFinder(ctrl *gomock.Controller) *MockRouteAnnotationFinder {
	mock := &MockRouteAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockRouteAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteAnnotationFinder) EXPECT() *MockRouteAnnotationFinderMockRecorder {
	return m.recorder
}

// FindRouteAnnotation mocks base method.
func (m *MockRouteAnnotationFinder) FindRouteAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.RouteAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindRouteAnnotation", varargs...)
	ret0, _ := ret[0].(*models.RouteAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRouteAnnotation indicates an expected call of FindRouteAnnotation.
func (mr *MockRouteAnnotationFinderMockRecorder) FindRouteAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRouteAnnotation", reflect.TypeOf((*MockRouteAnnotationFinder)(nil).FindRouteAnnotation), varargs...)
}

// MockRouteAnnotationInserter is a mock of RouteAnnotationInserter interface.
type MockRouteAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockRouteAnnotationInserterMockRecorder
}

// MockRouteAnnotationInserterMockRecorder is the mock recorder for MockRouteAnnotationInserter.
type MockRouteAnnotationInserterMockRecorder struct {
	mock *MockRouteAnnotationInserter
}

// NewMockRouteAnnotationInserter creates a new mock instance.
func NewMockRouteAnnotationInserter(ctrl *gomock.Controller) *MockRouteAnnotationInserter {
	mock := &MockRouteAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockRouteAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteAnnotationInserter) EXPECT() *MockRouteAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockRouteAnnotationInserter) Insert(o *models.RouteAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRouteAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRouteAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockRouteAnnotationUpdater is a mock of RouteAnnotationUpdater interface.
type MockRouteAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockRouteAnnotationUpdaterMockRecorder
}

// MockRouteAnnotationUpdaterMockRecorder is the mock recorder for MockRouteAnnotationUpdater.
type MockRouteAnnotationUpdaterMockRecorder struct {
	mock *MockRouteAnnotationUpdater
}

// NewMockRouteAnnotationUpdater creates a new mock instance.
func NewMockRouteAnnotationUpdater(ctrl *gomock.Controller) *MockRouteAnnotationUpdater {
	mock := &MockRouteAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockRouteAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteAnnotationUpdater) EXPECT() *MockRouteAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockRouteAnnotationUpdater) Update(o *models.RouteAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRouteAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRouteAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockRouteAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockRouteAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockRouteAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockRouteAnnotationUpdater) UpdateAllSlice(o models.RouteAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockRouteAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockRouteAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockRouteAnnotationDeleter is a mock of RouteAnnotationDeleter interface.
type MockRouteAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockRouteAnnotationDeleterMockRecorder
}

// MockRouteAnnotationDeleterMockRecorder is the mock recorder for MockRouteAnnotationDeleter.
type MockRouteAnnotationDeleterMockRecorder struct {
	mock *MockRouteAnnotationDeleter
}

// NewMockRouteAnnotationDeleter creates a new mock instance.
func NewMockRouteAnnotationDeleter(ctrl *gomock.Controller) *MockRouteAnnotationDeleter {
	mock := &MockRouteAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockRouteAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteAnnotationDeleter) EXPECT() *MockRouteAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRouteAnnotationDeleter) Delete(o *models.RouteAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRouteAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRouteAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockRouteAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockRouteAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockRouteAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockRouteAnnotationDeleter) DeleteAllSlice(o models.RouteAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockRouteAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockRouteAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockRouteAnnotationReloader is a mock of RouteAnnotationReloader interface.
type MockRouteAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockRouteAnnotationReloaderMockRecorder
}

// MockRouteAnnotationReloaderMockRecorder is the mock recorder for MockRouteAnnotationReloader.
type MockRouteAnnotationReloaderMockRecorder struct {
	mock *MockRouteAnnotationReloader
}

// NewMockRouteAnnotationReloader creates a new mock instance.
func NewMockRouteAnnotationReloader(ctrl *gomock.Controller) *MockRouteAnnotationReloader {
	mock := &MockRouteAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockRouteAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteAnnotationReloader) EXPECT() *MockRouteAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockRouteAnnotationReloader) Reload(o *models.RouteAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockRouteAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockRouteAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockRouteAnnotationReloader) ReloadAll(o *models.RouteAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockRouteAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockRouteAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
