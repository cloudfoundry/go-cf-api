// Code generated by MockGen. DO NOT EDIT.
// Source: psql_revision_sidecar_process_types.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockRevisionSidecarProcessTypeFinisher is a mock of RevisionSidecarProcessTypeFinisher interface.
type MockRevisionSidecarProcessTypeFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionSidecarProcessTypeFinisherMockRecorder
}

// MockRevisionSidecarProcessTypeFinisherMockRecorder is the mock recorder for MockRevisionSidecarProcessTypeFinisher.
type MockRevisionSidecarProcessTypeFinisherMockRecorder struct {
	mock *MockRevisionSidecarProcessTypeFinisher
}

// NewMockRevisionSidecarProcessTypeFinisher creates a new mock instance.
func NewMockRevisionSidecarProcessTypeFinisher(ctrl *gomock.Controller) *MockRevisionSidecarProcessTypeFinisher {
	mock := &MockRevisionSidecarProcessTypeFinisher{ctrl: ctrl}
	mock.recorder = &MockRevisionSidecarProcessTypeFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionSidecarProcessTypeFinisher) EXPECT() *MockRevisionSidecarProcessTypeFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockRevisionSidecarProcessTypeFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.RevisionSidecarProcessTypeSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.RevisionSidecarProcessTypeSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockRevisionSidecarProcessTypeFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRevisionSidecarProcessTypeFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockRevisionSidecarProcessTypeFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRevisionSidecarProcessTypeFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRevisionSidecarProcessTypeFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockRevisionSidecarProcessTypeFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRevisionSidecarProcessTypeFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRevisionSidecarProcessTypeFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockRevisionSidecarProcessTypeFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.RevisionSidecarProcessType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.RevisionSidecarProcessType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockRevisionSidecarProcessTypeFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockRevisionSidecarProcessTypeFinisher)(nil).One), ctx, exec)
}

// MockRevisionSidecarProcessTypeFinder is a mock of RevisionSidecarProcessTypeFinder interface.
type MockRevisionSidecarProcessTypeFinder struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionSidecarProcessTypeFinderMockRecorder
}

// MockRevisionSidecarProcessTypeFinderMockRecorder is the mock recorder for MockRevisionSidecarProcessTypeFinder.
type MockRevisionSidecarProcessTypeFinderMockRecorder struct {
	mock *MockRevisionSidecarProcessTypeFinder
}

// NewMockRevisionSidecarProcessTypeFinder creates a new mock instance.
func NewMockRevisionSidecarProcessTypeFinder(ctrl *gomock.Controller) *MockRevisionSidecarProcessTypeFinder {
	mock := &MockRevisionSidecarProcessTypeFinder{ctrl: ctrl}
	mock.recorder = &MockRevisionSidecarProcessTypeFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionSidecarProcessTypeFinder) EXPECT() *MockRevisionSidecarProcessTypeFinderMockRecorder {
	return m.recorder
}

// FindRevisionSidecarProcessType mocks base method.
func (m *MockRevisionSidecarProcessTypeFinder) FindRevisionSidecarProcessType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.RevisionSidecarProcessType, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindRevisionSidecarProcessType", varargs...)
	ret0, _ := ret[0].(*models.RevisionSidecarProcessType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRevisionSidecarProcessType indicates an expected call of FindRevisionSidecarProcessType.
func (mr *MockRevisionSidecarProcessTypeFinderMockRecorder) FindRevisionSidecarProcessType(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRevisionSidecarProcessType", reflect.TypeOf((*MockRevisionSidecarProcessTypeFinder)(nil).FindRevisionSidecarProcessType), varargs...)
}

// MockRevisionSidecarProcessTypeInserter is a mock of RevisionSidecarProcessTypeInserter interface.
type MockRevisionSidecarProcessTypeInserter struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionSidecarProcessTypeInserterMockRecorder
}

// MockRevisionSidecarProcessTypeInserterMockRecorder is the mock recorder for MockRevisionSidecarProcessTypeInserter.
type MockRevisionSidecarProcessTypeInserterMockRecorder struct {
	mock *MockRevisionSidecarProcessTypeInserter
}

// NewMockRevisionSidecarProcessTypeInserter creates a new mock instance.
func NewMockRevisionSidecarProcessTypeInserter(ctrl *gomock.Controller) *MockRevisionSidecarProcessTypeInserter {
	mock := &MockRevisionSidecarProcessTypeInserter{ctrl: ctrl}
	mock.recorder = &MockRevisionSidecarProcessTypeInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionSidecarProcessTypeInserter) EXPECT() *MockRevisionSidecarProcessTypeInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockRevisionSidecarProcessTypeInserter) Insert(o *models.RevisionSidecarProcessType, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRevisionSidecarProcessTypeInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRevisionSidecarProcessTypeInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockRevisionSidecarProcessTypeUpdater is a mock of RevisionSidecarProcessTypeUpdater interface.
type MockRevisionSidecarProcessTypeUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionSidecarProcessTypeUpdaterMockRecorder
}

// MockRevisionSidecarProcessTypeUpdaterMockRecorder is the mock recorder for MockRevisionSidecarProcessTypeUpdater.
type MockRevisionSidecarProcessTypeUpdaterMockRecorder struct {
	mock *MockRevisionSidecarProcessTypeUpdater
}

// NewMockRevisionSidecarProcessTypeUpdater creates a new mock instance.
func NewMockRevisionSidecarProcessTypeUpdater(ctrl *gomock.Controller) *MockRevisionSidecarProcessTypeUpdater {
	mock := &MockRevisionSidecarProcessTypeUpdater{ctrl: ctrl}
	mock.recorder = &MockRevisionSidecarProcessTypeUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionSidecarProcessTypeUpdater) EXPECT() *MockRevisionSidecarProcessTypeUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockRevisionSidecarProcessTypeUpdater) Update(o *models.RevisionSidecarProcessType, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRevisionSidecarProcessTypeUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRevisionSidecarProcessTypeUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockRevisionSidecarProcessTypeUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockRevisionSidecarProcessTypeUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockRevisionSidecarProcessTypeUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockRevisionSidecarProcessTypeUpdater) UpdateAllSlice(o models.RevisionSidecarProcessTypeSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockRevisionSidecarProcessTypeUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockRevisionSidecarProcessTypeUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockRevisionSidecarProcessTypeDeleter is a mock of RevisionSidecarProcessTypeDeleter interface.
type MockRevisionSidecarProcessTypeDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionSidecarProcessTypeDeleterMockRecorder
}

// MockRevisionSidecarProcessTypeDeleterMockRecorder is the mock recorder for MockRevisionSidecarProcessTypeDeleter.
type MockRevisionSidecarProcessTypeDeleterMockRecorder struct {
	mock *MockRevisionSidecarProcessTypeDeleter
}

// NewMockRevisionSidecarProcessTypeDeleter creates a new mock instance.
func NewMockRevisionSidecarProcessTypeDeleter(ctrl *gomock.Controller) *MockRevisionSidecarProcessTypeDeleter {
	mock := &MockRevisionSidecarProcessTypeDeleter{ctrl: ctrl}
	mock.recorder = &MockRevisionSidecarProcessTypeDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionSidecarProcessTypeDeleter) EXPECT() *MockRevisionSidecarProcessTypeDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRevisionSidecarProcessTypeDeleter) Delete(o *models.RevisionSidecarProcessType, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRevisionSidecarProcessTypeDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRevisionSidecarProcessTypeDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockRevisionSidecarProcessTypeDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockRevisionSidecarProcessTypeDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockRevisionSidecarProcessTypeDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockRevisionSidecarProcessTypeDeleter) DeleteAllSlice(o models.RevisionSidecarProcessTypeSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockRevisionSidecarProcessTypeDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockRevisionSidecarProcessTypeDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockRevisionSidecarProcessTypeReloader is a mock of RevisionSidecarProcessTypeReloader interface.
type MockRevisionSidecarProcessTypeReloader struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionSidecarProcessTypeReloaderMockRecorder
}

// MockRevisionSidecarProcessTypeReloaderMockRecorder is the mock recorder for MockRevisionSidecarProcessTypeReloader.
type MockRevisionSidecarProcessTypeReloaderMockRecorder struct {
	mock *MockRevisionSidecarProcessTypeReloader
}

// NewMockRevisionSidecarProcessTypeReloader creates a new mock instance.
func NewMockRevisionSidecarProcessTypeReloader(ctrl *gomock.Controller) *MockRevisionSidecarProcessTypeReloader {
	mock := &MockRevisionSidecarProcessTypeReloader{ctrl: ctrl}
	mock.recorder = &MockRevisionSidecarProcessTypeReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionSidecarProcessTypeReloader) EXPECT() *MockRevisionSidecarProcessTypeReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockRevisionSidecarProcessTypeReloader) Reload(o *models.RevisionSidecarProcessType, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockRevisionSidecarProcessTypeReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockRevisionSidecarProcessTypeReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockRevisionSidecarProcessTypeReloader) ReloadAll(o *models.RevisionSidecarProcessTypeSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockRevisionSidecarProcessTypeReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockRevisionSidecarProcessTypeReloader)(nil).ReloadAll), o, ctx, exec)
}
