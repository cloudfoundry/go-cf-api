// Code generated by MockGen. DO NOT EDIT.
// Source: psql_kpack_lifecycle_data.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockKpackLifecycleDatumFinisher is a mock of KpackLifecycleDatumFinisher interface.
type MockKpackLifecycleDatumFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockKpackLifecycleDatumFinisherMockRecorder
}

// MockKpackLifecycleDatumFinisherMockRecorder is the mock recorder for MockKpackLifecycleDatumFinisher.
type MockKpackLifecycleDatumFinisherMockRecorder struct {
	mock *MockKpackLifecycleDatumFinisher
}

// NewMockKpackLifecycleDatumFinisher creates a new mock instance.
func NewMockKpackLifecycleDatumFinisher(ctrl *gomock.Controller) *MockKpackLifecycleDatumFinisher {
	mock := &MockKpackLifecycleDatumFinisher{ctrl: ctrl}
	mock.recorder = &MockKpackLifecycleDatumFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKpackLifecycleDatumFinisher) EXPECT() *MockKpackLifecycleDatumFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockKpackLifecycleDatumFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.KpackLifecycleDatumSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.KpackLifecycleDatumSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockKpackLifecycleDatumFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockKpackLifecycleDatumFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockKpackLifecycleDatumFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockKpackLifecycleDatumFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockKpackLifecycleDatumFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockKpackLifecycleDatumFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockKpackLifecycleDatumFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockKpackLifecycleDatumFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockKpackLifecycleDatumFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.KpackLifecycleDatum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.KpackLifecycleDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockKpackLifecycleDatumFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockKpackLifecycleDatumFinisher)(nil).One), ctx, exec)
}

// MockKpackLifecycleDatumFinder is a mock of KpackLifecycleDatumFinder interface.
type MockKpackLifecycleDatumFinder struct {
	ctrl     *gomock.Controller
	recorder *MockKpackLifecycleDatumFinderMockRecorder
}

// MockKpackLifecycleDatumFinderMockRecorder is the mock recorder for MockKpackLifecycleDatumFinder.
type MockKpackLifecycleDatumFinderMockRecorder struct {
	mock *MockKpackLifecycleDatumFinder
}

// NewMockKpackLifecycleDatumFinder creates a new mock instance.
func NewMockKpackLifecycleDatumFinder(ctrl *gomock.Controller) *MockKpackLifecycleDatumFinder {
	mock := &MockKpackLifecycleDatumFinder{ctrl: ctrl}
	mock.recorder = &MockKpackLifecycleDatumFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKpackLifecycleDatumFinder) EXPECT() *MockKpackLifecycleDatumFinderMockRecorder {
	return m.recorder
}

// FindKpackLifecycleDatum mocks base method.
func (m *MockKpackLifecycleDatumFinder) FindKpackLifecycleDatum(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.KpackLifecycleDatum, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindKpackLifecycleDatum", varargs...)
	ret0, _ := ret[0].(*models.KpackLifecycleDatum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKpackLifecycleDatum indicates an expected call of FindKpackLifecycleDatum.
func (mr *MockKpackLifecycleDatumFinderMockRecorder) FindKpackLifecycleDatum(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKpackLifecycleDatum", reflect.TypeOf((*MockKpackLifecycleDatumFinder)(nil).FindKpackLifecycleDatum), varargs...)
}

// MockKpackLifecycleDatumInserter is a mock of KpackLifecycleDatumInserter interface.
type MockKpackLifecycleDatumInserter struct {
	ctrl     *gomock.Controller
	recorder *MockKpackLifecycleDatumInserterMockRecorder
}

// MockKpackLifecycleDatumInserterMockRecorder is the mock recorder for MockKpackLifecycleDatumInserter.
type MockKpackLifecycleDatumInserterMockRecorder struct {
	mock *MockKpackLifecycleDatumInserter
}

// NewMockKpackLifecycleDatumInserter creates a new mock instance.
func NewMockKpackLifecycleDatumInserter(ctrl *gomock.Controller) *MockKpackLifecycleDatumInserter {
	mock := &MockKpackLifecycleDatumInserter{ctrl: ctrl}
	mock.recorder = &MockKpackLifecycleDatumInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKpackLifecycleDatumInserter) EXPECT() *MockKpackLifecycleDatumInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockKpackLifecycleDatumInserter) Insert(o *models.KpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockKpackLifecycleDatumInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockKpackLifecycleDatumInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockKpackLifecycleDatumUpdater is a mock of KpackLifecycleDatumUpdater interface.
type MockKpackLifecycleDatumUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockKpackLifecycleDatumUpdaterMockRecorder
}

// MockKpackLifecycleDatumUpdaterMockRecorder is the mock recorder for MockKpackLifecycleDatumUpdater.
type MockKpackLifecycleDatumUpdaterMockRecorder struct {
	mock *MockKpackLifecycleDatumUpdater
}

// NewMockKpackLifecycleDatumUpdater creates a new mock instance.
func NewMockKpackLifecycleDatumUpdater(ctrl *gomock.Controller) *MockKpackLifecycleDatumUpdater {
	mock := &MockKpackLifecycleDatumUpdater{ctrl: ctrl}
	mock.recorder = &MockKpackLifecycleDatumUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKpackLifecycleDatumUpdater) EXPECT() *MockKpackLifecycleDatumUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockKpackLifecycleDatumUpdater) Update(o *models.KpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockKpackLifecycleDatumUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockKpackLifecycleDatumUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockKpackLifecycleDatumUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockKpackLifecycleDatumUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockKpackLifecycleDatumUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockKpackLifecycleDatumUpdater) UpdateAllSlice(o models.KpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockKpackLifecycleDatumUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockKpackLifecycleDatumUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockKpackLifecycleDatumDeleter is a mock of KpackLifecycleDatumDeleter interface.
type MockKpackLifecycleDatumDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockKpackLifecycleDatumDeleterMockRecorder
}

// MockKpackLifecycleDatumDeleterMockRecorder is the mock recorder for MockKpackLifecycleDatumDeleter.
type MockKpackLifecycleDatumDeleterMockRecorder struct {
	mock *MockKpackLifecycleDatumDeleter
}

// NewMockKpackLifecycleDatumDeleter creates a new mock instance.
func NewMockKpackLifecycleDatumDeleter(ctrl *gomock.Controller) *MockKpackLifecycleDatumDeleter {
	mock := &MockKpackLifecycleDatumDeleter{ctrl: ctrl}
	mock.recorder = &MockKpackLifecycleDatumDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKpackLifecycleDatumDeleter) EXPECT() *MockKpackLifecycleDatumDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockKpackLifecycleDatumDeleter) Delete(o *models.KpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockKpackLifecycleDatumDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKpackLifecycleDatumDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockKpackLifecycleDatumDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockKpackLifecycleDatumDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockKpackLifecycleDatumDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockKpackLifecycleDatumDeleter) DeleteAllSlice(o models.KpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockKpackLifecycleDatumDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockKpackLifecycleDatumDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockKpackLifecycleDatumReloader is a mock of KpackLifecycleDatumReloader interface.
type MockKpackLifecycleDatumReloader struct {
	ctrl     *gomock.Controller
	recorder *MockKpackLifecycleDatumReloaderMockRecorder
}

// MockKpackLifecycleDatumReloaderMockRecorder is the mock recorder for MockKpackLifecycleDatumReloader.
type MockKpackLifecycleDatumReloaderMockRecorder struct {
	mock *MockKpackLifecycleDatumReloader
}

// NewMockKpackLifecycleDatumReloader creates a new mock instance.
func NewMockKpackLifecycleDatumReloader(ctrl *gomock.Controller) *MockKpackLifecycleDatumReloader {
	mock := &MockKpackLifecycleDatumReloader{ctrl: ctrl}
	mock.recorder = &MockKpackLifecycleDatumReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKpackLifecycleDatumReloader) EXPECT() *MockKpackLifecycleDatumReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockKpackLifecycleDatumReloader) Reload(o *models.KpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockKpackLifecycleDatumReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockKpackLifecycleDatumReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockKpackLifecycleDatumReloader) ReloadAll(o *models.KpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockKpackLifecycleDatumReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockKpackLifecycleDatumReloader)(nil).ReloadAll), o, ctx, exec)
}
