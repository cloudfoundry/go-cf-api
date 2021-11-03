// +build unit

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_user_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockUserAnnotationUpserter is a mock of UserAnnotationUpserter interface.
type MockUserAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnnotationUpserterMockRecorder
}

// MockUserAnnotationUpserterMockRecorder is the mock recorder for MockUserAnnotationUpserter.
type MockUserAnnotationUpserterMockRecorder struct {
	mock *MockUserAnnotationUpserter
}

// NewMockUserAnnotationUpserter creates a new mock instance.
func NewMockUserAnnotationUpserter(ctrl *gomock.Controller) *MockUserAnnotationUpserter {
	mock := &MockUserAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockUserAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnnotationUpserter) EXPECT() *MockUserAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockUserAnnotationUpserter) Upsert(o *models.UserAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockUserAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockUserAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockUserAnnotationFinisher is a mock of UserAnnotationFinisher interface.
type MockUserAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnnotationFinisherMockRecorder
}

// MockUserAnnotationFinisherMockRecorder is the mock recorder for MockUserAnnotationFinisher.
type MockUserAnnotationFinisherMockRecorder struct {
	mock *MockUserAnnotationFinisher
}

// NewMockUserAnnotationFinisher creates a new mock instance.
func NewMockUserAnnotationFinisher(ctrl *gomock.Controller) *MockUserAnnotationFinisher {
	mock := &MockUserAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockUserAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnnotationFinisher) EXPECT() *MockUserAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockUserAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.UserAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.UserAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockUserAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockUserAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockUserAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockUserAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUserAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockUserAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockUserAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockUserAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockUserAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.UserAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.UserAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockUserAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockUserAnnotationFinisher)(nil).One), ctx, exec)
}

// MockUserAnnotationFinder is a mock of UserAnnotationFinder interface.
type MockUserAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnnotationFinderMockRecorder
}

// MockUserAnnotationFinderMockRecorder is the mock recorder for MockUserAnnotationFinder.
type MockUserAnnotationFinderMockRecorder struct {
	mock *MockUserAnnotationFinder
}

// NewMockUserAnnotationFinder creates a new mock instance.
func NewMockUserAnnotationFinder(ctrl *gomock.Controller) *MockUserAnnotationFinder {
	mock := &MockUserAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockUserAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnnotationFinder) EXPECT() *MockUserAnnotationFinderMockRecorder {
	return m.recorder
}

// FindUserAnnotation mocks base method.
func (m *MockUserAnnotationFinder) FindUserAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.UserAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindUserAnnotation", varargs...)
	ret0, _ := ret[0].(*models.UserAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserAnnotation indicates an expected call of FindUserAnnotation.
func (mr *MockUserAnnotationFinderMockRecorder) FindUserAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserAnnotation", reflect.TypeOf((*MockUserAnnotationFinder)(nil).FindUserAnnotation), varargs...)
}

// MockUserAnnotationInserter is a mock of UserAnnotationInserter interface.
type MockUserAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnnotationInserterMockRecorder
}

// MockUserAnnotationInserterMockRecorder is the mock recorder for MockUserAnnotationInserter.
type MockUserAnnotationInserterMockRecorder struct {
	mock *MockUserAnnotationInserter
}

// NewMockUserAnnotationInserter creates a new mock instance.
func NewMockUserAnnotationInserter(ctrl *gomock.Controller) *MockUserAnnotationInserter {
	mock := &MockUserAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockUserAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnnotationInserter) EXPECT() *MockUserAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockUserAnnotationInserter) Insert(o *models.UserAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockUserAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockUserAnnotationUpdater is a mock of UserAnnotationUpdater interface.
type MockUserAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnnotationUpdaterMockRecorder
}

// MockUserAnnotationUpdaterMockRecorder is the mock recorder for MockUserAnnotationUpdater.
type MockUserAnnotationUpdaterMockRecorder struct {
	mock *MockUserAnnotationUpdater
}

// NewMockUserAnnotationUpdater creates a new mock instance.
func NewMockUserAnnotationUpdater(ctrl *gomock.Controller) *MockUserAnnotationUpdater {
	mock := &MockUserAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockUserAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnnotationUpdater) EXPECT() *MockUserAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockUserAnnotationUpdater) Update(o *models.UserAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockUserAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockUserAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockUserAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockUserAnnotationUpdater) UpdateAllSlice(o models.UserAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockUserAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockUserAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockUserAnnotationDeleter is a mock of UserAnnotationDeleter interface.
type MockUserAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnnotationDeleterMockRecorder
}

// MockUserAnnotationDeleterMockRecorder is the mock recorder for MockUserAnnotationDeleter.
type MockUserAnnotationDeleterMockRecorder struct {
	mock *MockUserAnnotationDeleter
}

// NewMockUserAnnotationDeleter creates a new mock instance.
func NewMockUserAnnotationDeleter(ctrl *gomock.Controller) *MockUserAnnotationDeleter {
	mock := &MockUserAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockUserAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnnotationDeleter) EXPECT() *MockUserAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserAnnotationDeleter) Delete(o *models.UserAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockUserAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockUserAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockUserAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockUserAnnotationDeleter) DeleteAllSlice(o models.UserAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockUserAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockUserAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockUserAnnotationReloader is a mock of UserAnnotationReloader interface.
type MockUserAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnnotationReloaderMockRecorder
}

// MockUserAnnotationReloaderMockRecorder is the mock recorder for MockUserAnnotationReloader.
type MockUserAnnotationReloaderMockRecorder struct {
	mock *MockUserAnnotationReloader
}

// NewMockUserAnnotationReloader creates a new mock instance.
func NewMockUserAnnotationReloader(ctrl *gomock.Controller) *MockUserAnnotationReloader {
	mock := &MockUserAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockUserAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnnotationReloader) EXPECT() *MockUserAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockUserAnnotationReloader) Reload(o *models.UserAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockUserAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockUserAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockUserAnnotationReloader) ReloadAll(o *models.UserAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockUserAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockUserAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
