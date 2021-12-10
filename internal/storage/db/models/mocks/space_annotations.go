//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_space_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockSpaceAnnotationUpserter is a mock of SpaceAnnotationUpserter interface.
type MockSpaceAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceAnnotationUpserterMockRecorder
}

// MockSpaceAnnotationUpserterMockRecorder is the mock recorder for MockSpaceAnnotationUpserter.
type MockSpaceAnnotationUpserterMockRecorder struct {
	mock *MockSpaceAnnotationUpserter
}

// NewMockSpaceAnnotationUpserter creates a new mock instance.
func NewMockSpaceAnnotationUpserter(ctrl *gomock.Controller) *MockSpaceAnnotationUpserter {
	mock := &MockSpaceAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockSpaceAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceAnnotationUpserter) EXPECT() *MockSpaceAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockSpaceAnnotationUpserter) Upsert(o *models.SpaceAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSpaceAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSpaceAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockSpaceAnnotationFinisher is a mock of SpaceAnnotationFinisher interface.
type MockSpaceAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceAnnotationFinisherMockRecorder
}

// MockSpaceAnnotationFinisherMockRecorder is the mock recorder for MockSpaceAnnotationFinisher.
type MockSpaceAnnotationFinisherMockRecorder struct {
	mock *MockSpaceAnnotationFinisher
}

// NewMockSpaceAnnotationFinisher creates a new mock instance.
func NewMockSpaceAnnotationFinisher(ctrl *gomock.Controller) *MockSpaceAnnotationFinisher {
	mock := &MockSpaceAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockSpaceAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceAnnotationFinisher) EXPECT() *MockSpaceAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockSpaceAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.SpaceAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.SpaceAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockSpaceAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockSpaceAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockSpaceAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSpaceAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSpaceAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockSpaceAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSpaceAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSpaceAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockSpaceAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.SpaceAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.SpaceAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockSpaceAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockSpaceAnnotationFinisher)(nil).One), ctx, exec)
}

// MockSpaceAnnotationFinder is a mock of SpaceAnnotationFinder interface.
type MockSpaceAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceAnnotationFinderMockRecorder
}

// MockSpaceAnnotationFinderMockRecorder is the mock recorder for MockSpaceAnnotationFinder.
type MockSpaceAnnotationFinderMockRecorder struct {
	mock *MockSpaceAnnotationFinder
}

// NewMockSpaceAnnotationFinder creates a new mock instance.
func NewMockSpaceAnnotationFinder(ctrl *gomock.Controller) *MockSpaceAnnotationFinder {
	mock := &MockSpaceAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockSpaceAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceAnnotationFinder) EXPECT() *MockSpaceAnnotationFinderMockRecorder {
	return m.recorder
}

// FindSpaceAnnotation mocks base method.
func (m *MockSpaceAnnotationFinder) FindSpaceAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.SpaceAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSpaceAnnotation", varargs...)
	ret0, _ := ret[0].(*models.SpaceAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSpaceAnnotation indicates an expected call of FindSpaceAnnotation.
func (mr *MockSpaceAnnotationFinderMockRecorder) FindSpaceAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSpaceAnnotation", reflect.TypeOf((*MockSpaceAnnotationFinder)(nil).FindSpaceAnnotation), varargs...)
}

// MockSpaceAnnotationInserter is a mock of SpaceAnnotationInserter interface.
type MockSpaceAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceAnnotationInserterMockRecorder
}

// MockSpaceAnnotationInserterMockRecorder is the mock recorder for MockSpaceAnnotationInserter.
type MockSpaceAnnotationInserterMockRecorder struct {
	mock *MockSpaceAnnotationInserter
}

// NewMockSpaceAnnotationInserter creates a new mock instance.
func NewMockSpaceAnnotationInserter(ctrl *gomock.Controller) *MockSpaceAnnotationInserter {
	mock := &MockSpaceAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockSpaceAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceAnnotationInserter) EXPECT() *MockSpaceAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockSpaceAnnotationInserter) Insert(o *models.SpaceAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSpaceAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSpaceAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockSpaceAnnotationUpdater is a mock of SpaceAnnotationUpdater interface.
type MockSpaceAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceAnnotationUpdaterMockRecorder
}

// MockSpaceAnnotationUpdaterMockRecorder is the mock recorder for MockSpaceAnnotationUpdater.
type MockSpaceAnnotationUpdaterMockRecorder struct {
	mock *MockSpaceAnnotationUpdater
}

// NewMockSpaceAnnotationUpdater creates a new mock instance.
func NewMockSpaceAnnotationUpdater(ctrl *gomock.Controller) *MockSpaceAnnotationUpdater {
	mock := &MockSpaceAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockSpaceAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceAnnotationUpdater) EXPECT() *MockSpaceAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockSpaceAnnotationUpdater) Update(o *models.SpaceAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSpaceAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSpaceAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockSpaceAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockSpaceAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockSpaceAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockSpaceAnnotationUpdater) UpdateAllSlice(o models.SpaceAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockSpaceAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockSpaceAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockSpaceAnnotationDeleter is a mock of SpaceAnnotationDeleter interface.
type MockSpaceAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceAnnotationDeleterMockRecorder
}

// MockSpaceAnnotationDeleterMockRecorder is the mock recorder for MockSpaceAnnotationDeleter.
type MockSpaceAnnotationDeleterMockRecorder struct {
	mock *MockSpaceAnnotationDeleter
}

// NewMockSpaceAnnotationDeleter creates a new mock instance.
func NewMockSpaceAnnotationDeleter(ctrl *gomock.Controller) *MockSpaceAnnotationDeleter {
	mock := &MockSpaceAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockSpaceAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceAnnotationDeleter) EXPECT() *MockSpaceAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSpaceAnnotationDeleter) Delete(o *models.SpaceAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSpaceAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSpaceAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockSpaceAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockSpaceAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockSpaceAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockSpaceAnnotationDeleter) DeleteAllSlice(o models.SpaceAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockSpaceAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockSpaceAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockSpaceAnnotationReloader is a mock of SpaceAnnotationReloader interface.
type MockSpaceAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceAnnotationReloaderMockRecorder
}

// MockSpaceAnnotationReloaderMockRecorder is the mock recorder for MockSpaceAnnotationReloader.
type MockSpaceAnnotationReloaderMockRecorder struct {
	mock *MockSpaceAnnotationReloader
}

// NewMockSpaceAnnotationReloader creates a new mock instance.
func NewMockSpaceAnnotationReloader(ctrl *gomock.Controller) *MockSpaceAnnotationReloader {
	mock := &MockSpaceAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockSpaceAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceAnnotationReloader) EXPECT() *MockSpaceAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockSpaceAnnotationReloader) Reload(o *models.SpaceAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockSpaceAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSpaceAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockSpaceAnnotationReloader) ReloadAll(o *models.SpaceAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockSpaceAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockSpaceAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
