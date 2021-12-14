//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_app_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockAppAnnotationUpserter is a mock of AppAnnotationUpserter interface.
type MockAppAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppAnnotationUpserterMockRecorder
}

// MockAppAnnotationUpserterMockRecorder is the mock recorder for MockAppAnnotationUpserter.
type MockAppAnnotationUpserterMockRecorder struct {
	mock *MockAppAnnotationUpserter
}

// NewMockAppAnnotationUpserter creates a new mock instance.
func NewMockAppAnnotationUpserter(ctrl *gomock.Controller) *MockAppAnnotationUpserter {
	mock := &MockAppAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockAppAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppAnnotationUpserter) EXPECT() *MockAppAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockAppAnnotationUpserter) Upsert(o *models.AppAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockAppAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockAppAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockAppAnnotationFinisher is a mock of AppAnnotationFinisher interface.
type MockAppAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockAppAnnotationFinisherMockRecorder
}

// MockAppAnnotationFinisherMockRecorder is the mock recorder for MockAppAnnotationFinisher.
type MockAppAnnotationFinisherMockRecorder struct {
	mock *MockAppAnnotationFinisher
}

// NewMockAppAnnotationFinisher creates a new mock instance.
func NewMockAppAnnotationFinisher(ctrl *gomock.Controller) *MockAppAnnotationFinisher {
	mock := &MockAppAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockAppAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppAnnotationFinisher) EXPECT() *MockAppAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockAppAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.AppAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.AppAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockAppAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockAppAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockAppAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockAppAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockAppAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockAppAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockAppAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockAppAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockAppAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.AppAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.AppAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockAppAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockAppAnnotationFinisher)(nil).One), ctx, exec)
}

// MockAppAnnotationFinder is a mock of AppAnnotationFinder interface.
type MockAppAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockAppAnnotationFinderMockRecorder
}

// MockAppAnnotationFinderMockRecorder is the mock recorder for MockAppAnnotationFinder.
type MockAppAnnotationFinderMockRecorder struct {
	mock *MockAppAnnotationFinder
}

// NewMockAppAnnotationFinder creates a new mock instance.
func NewMockAppAnnotationFinder(ctrl *gomock.Controller) *MockAppAnnotationFinder {
	mock := &MockAppAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockAppAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppAnnotationFinder) EXPECT() *MockAppAnnotationFinderMockRecorder {
	return m.recorder
}

// FindAppAnnotation mocks base method.
func (m *MockAppAnnotationFinder) FindAppAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.AppAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAppAnnotation", varargs...)
	ret0, _ := ret[0].(*models.AppAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAppAnnotation indicates an expected call of FindAppAnnotation.
func (mr *MockAppAnnotationFinderMockRecorder) FindAppAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAppAnnotation", reflect.TypeOf((*MockAppAnnotationFinder)(nil).FindAppAnnotation), varargs...)
}

// MockAppAnnotationInserter is a mock of AppAnnotationInserter interface.
type MockAppAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppAnnotationInserterMockRecorder
}

// MockAppAnnotationInserterMockRecorder is the mock recorder for MockAppAnnotationInserter.
type MockAppAnnotationInserterMockRecorder struct {
	mock *MockAppAnnotationInserter
}

// NewMockAppAnnotationInserter creates a new mock instance.
func NewMockAppAnnotationInserter(ctrl *gomock.Controller) *MockAppAnnotationInserter {
	mock := &MockAppAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockAppAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppAnnotationInserter) EXPECT() *MockAppAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockAppAnnotationInserter) Insert(o *models.AppAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAppAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAppAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockAppAnnotationUpdater is a mock of AppAnnotationUpdater interface.
type MockAppAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAppAnnotationUpdaterMockRecorder
}

// MockAppAnnotationUpdaterMockRecorder is the mock recorder for MockAppAnnotationUpdater.
type MockAppAnnotationUpdaterMockRecorder struct {
	mock *MockAppAnnotationUpdater
}

// NewMockAppAnnotationUpdater creates a new mock instance.
func NewMockAppAnnotationUpdater(ctrl *gomock.Controller) *MockAppAnnotationUpdater {
	mock := &MockAppAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockAppAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppAnnotationUpdater) EXPECT() *MockAppAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockAppAnnotationUpdater) Update(o *models.AppAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAppAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockAppAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockAppAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockAppAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockAppAnnotationUpdater) UpdateAllSlice(o models.AppAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockAppAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockAppAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockAppAnnotationDeleter is a mock of AppAnnotationDeleter interface.
type MockAppAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockAppAnnotationDeleterMockRecorder
}

// MockAppAnnotationDeleterMockRecorder is the mock recorder for MockAppAnnotationDeleter.
type MockAppAnnotationDeleterMockRecorder struct {
	mock *MockAppAnnotationDeleter
}

// NewMockAppAnnotationDeleter creates a new mock instance.
func NewMockAppAnnotationDeleter(ctrl *gomock.Controller) *MockAppAnnotationDeleter {
	mock := &MockAppAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockAppAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppAnnotationDeleter) EXPECT() *MockAppAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAppAnnotationDeleter) Delete(o *models.AppAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAppAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockAppAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockAppAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockAppAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockAppAnnotationDeleter) DeleteAllSlice(o models.AppAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockAppAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockAppAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockAppAnnotationReloader is a mock of AppAnnotationReloader interface.
type MockAppAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockAppAnnotationReloaderMockRecorder
}

// MockAppAnnotationReloaderMockRecorder is the mock recorder for MockAppAnnotationReloader.
type MockAppAnnotationReloaderMockRecorder struct {
	mock *MockAppAnnotationReloader
}

// NewMockAppAnnotationReloader creates a new mock instance.
func NewMockAppAnnotationReloader(ctrl *gomock.Controller) *MockAppAnnotationReloader {
	mock := &MockAppAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockAppAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppAnnotationReloader) EXPECT() *MockAppAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockAppAnnotationReloader) Reload(o *models.AppAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockAppAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockAppAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockAppAnnotationReloader) ReloadAll(o *models.AppAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockAppAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockAppAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
