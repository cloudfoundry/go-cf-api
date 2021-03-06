//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_build_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockBuildAnnotationUpserter is a mock of BuildAnnotationUpserter interface.
type MockBuildAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildAnnotationUpserterMockRecorder
}

// MockBuildAnnotationUpserterMockRecorder is the mock recorder for MockBuildAnnotationUpserter.
type MockBuildAnnotationUpserterMockRecorder struct {
	mock *MockBuildAnnotationUpserter
}

// NewMockBuildAnnotationUpserter creates a new mock instance.
func NewMockBuildAnnotationUpserter(ctrl *gomock.Controller) *MockBuildAnnotationUpserter {
	mock := &MockBuildAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockBuildAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildAnnotationUpserter) EXPECT() *MockBuildAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockBuildAnnotationUpserter) Upsert(o *models.BuildAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockBuildAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockBuildAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockBuildAnnotationFinisher is a mock of BuildAnnotationFinisher interface.
type MockBuildAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockBuildAnnotationFinisherMockRecorder
}

// MockBuildAnnotationFinisherMockRecorder is the mock recorder for MockBuildAnnotationFinisher.
type MockBuildAnnotationFinisherMockRecorder struct {
	mock *MockBuildAnnotationFinisher
}

// NewMockBuildAnnotationFinisher creates a new mock instance.
func NewMockBuildAnnotationFinisher(ctrl *gomock.Controller) *MockBuildAnnotationFinisher {
	mock := &MockBuildAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockBuildAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildAnnotationFinisher) EXPECT() *MockBuildAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockBuildAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.BuildAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.BuildAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockBuildAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockBuildAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockBuildAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockBuildAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockBuildAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockBuildAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockBuildAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockBuildAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockBuildAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.BuildAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.BuildAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockBuildAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockBuildAnnotationFinisher)(nil).One), ctx, exec)
}

// MockBuildAnnotationFinder is a mock of BuildAnnotationFinder interface.
type MockBuildAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockBuildAnnotationFinderMockRecorder
}

// MockBuildAnnotationFinderMockRecorder is the mock recorder for MockBuildAnnotationFinder.
type MockBuildAnnotationFinderMockRecorder struct {
	mock *MockBuildAnnotationFinder
}

// NewMockBuildAnnotationFinder creates a new mock instance.
func NewMockBuildAnnotationFinder(ctrl *gomock.Controller) *MockBuildAnnotationFinder {
	mock := &MockBuildAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockBuildAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildAnnotationFinder) EXPECT() *MockBuildAnnotationFinderMockRecorder {
	return m.recorder
}

// FindBuildAnnotation mocks base method.
func (m *MockBuildAnnotationFinder) FindBuildAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.BuildAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindBuildAnnotation", varargs...)
	ret0, _ := ret[0].(*models.BuildAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBuildAnnotation indicates an expected call of FindBuildAnnotation.
func (mr *MockBuildAnnotationFinderMockRecorder) FindBuildAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBuildAnnotation", reflect.TypeOf((*MockBuildAnnotationFinder)(nil).FindBuildAnnotation), varargs...)
}

// MockBuildAnnotationInserter is a mock of BuildAnnotationInserter interface.
type MockBuildAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildAnnotationInserterMockRecorder
}

// MockBuildAnnotationInserterMockRecorder is the mock recorder for MockBuildAnnotationInserter.
type MockBuildAnnotationInserterMockRecorder struct {
	mock *MockBuildAnnotationInserter
}

// NewMockBuildAnnotationInserter creates a new mock instance.
func NewMockBuildAnnotationInserter(ctrl *gomock.Controller) *MockBuildAnnotationInserter {
	mock := &MockBuildAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockBuildAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildAnnotationInserter) EXPECT() *MockBuildAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockBuildAnnotationInserter) Insert(o *models.BuildAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockBuildAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBuildAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockBuildAnnotationUpdater is a mock of BuildAnnotationUpdater interface.
type MockBuildAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockBuildAnnotationUpdaterMockRecorder
}

// MockBuildAnnotationUpdaterMockRecorder is the mock recorder for MockBuildAnnotationUpdater.
type MockBuildAnnotationUpdaterMockRecorder struct {
	mock *MockBuildAnnotationUpdater
}

// NewMockBuildAnnotationUpdater creates a new mock instance.
func NewMockBuildAnnotationUpdater(ctrl *gomock.Controller) *MockBuildAnnotationUpdater {
	mock := &MockBuildAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockBuildAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildAnnotationUpdater) EXPECT() *MockBuildAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockBuildAnnotationUpdater) Update(o *models.BuildAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBuildAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBuildAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockBuildAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockBuildAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockBuildAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockBuildAnnotationUpdater) UpdateAllSlice(o models.BuildAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockBuildAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockBuildAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockBuildAnnotationDeleter is a mock of BuildAnnotationDeleter interface.
type MockBuildAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildAnnotationDeleterMockRecorder
}

// MockBuildAnnotationDeleterMockRecorder is the mock recorder for MockBuildAnnotationDeleter.
type MockBuildAnnotationDeleterMockRecorder struct {
	mock *MockBuildAnnotationDeleter
}

// NewMockBuildAnnotationDeleter creates a new mock instance.
func NewMockBuildAnnotationDeleter(ctrl *gomock.Controller) *MockBuildAnnotationDeleter {
	mock := &MockBuildAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockBuildAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildAnnotationDeleter) EXPECT() *MockBuildAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBuildAnnotationDeleter) Delete(o *models.BuildAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBuildAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBuildAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockBuildAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockBuildAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockBuildAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockBuildAnnotationDeleter) DeleteAllSlice(o models.BuildAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockBuildAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockBuildAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockBuildAnnotationReloader is a mock of BuildAnnotationReloader interface.
type MockBuildAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockBuildAnnotationReloaderMockRecorder
}

// MockBuildAnnotationReloaderMockRecorder is the mock recorder for MockBuildAnnotationReloader.
type MockBuildAnnotationReloaderMockRecorder struct {
	mock *MockBuildAnnotationReloader
}

// NewMockBuildAnnotationReloader creates a new mock instance.
func NewMockBuildAnnotationReloader(ctrl *gomock.Controller) *MockBuildAnnotationReloader {
	mock := &MockBuildAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockBuildAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildAnnotationReloader) EXPECT() *MockBuildAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockBuildAnnotationReloader) Reload(o *models.BuildAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockBuildAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockBuildAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockBuildAnnotationReloader) ReloadAll(o *models.BuildAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockBuildAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockBuildAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
