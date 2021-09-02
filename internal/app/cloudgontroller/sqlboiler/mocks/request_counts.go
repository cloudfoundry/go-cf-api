// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/cloudgontroller/sqlboiler/psql_request_counts.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockRequestCountFinisher is a mock of RequestCountFinisher interface.
type MockRequestCountFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockRequestCountFinisherMockRecorder
}

// MockRequestCountFinisherMockRecorder is the mock recorder for MockRequestCountFinisher.
type MockRequestCountFinisherMockRecorder struct {
	mock *MockRequestCountFinisher
}

// NewMockRequestCountFinisher creates a new mock instance.
func NewMockRequestCountFinisher(ctrl *gomock.Controller) *MockRequestCountFinisher {
	mock := &MockRequestCountFinisher{ctrl: ctrl}
	mock.recorder = &MockRequestCountFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestCountFinisher) EXPECT() *MockRequestCountFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockRequestCountFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.RequestCountSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.RequestCountSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockRequestCountFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRequestCountFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockRequestCountFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRequestCountFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRequestCountFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockRequestCountFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRequestCountFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRequestCountFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockRequestCountFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.RequestCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.RequestCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockRequestCountFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockRequestCountFinisher)(nil).One), ctx, exec)
}

// MockRequestCountFinder is a mock of RequestCountFinder interface.
type MockRequestCountFinder struct {
	ctrl     *gomock.Controller
	recorder *MockRequestCountFinderMockRecorder
}

// MockRequestCountFinderMockRecorder is the mock recorder for MockRequestCountFinder.
type MockRequestCountFinderMockRecorder struct {
	mock *MockRequestCountFinder
}

// NewMockRequestCountFinder creates a new mock instance.
func NewMockRequestCountFinder(ctrl *gomock.Controller) *MockRequestCountFinder {
	mock := &MockRequestCountFinder{ctrl: ctrl}
	mock.recorder = &MockRequestCountFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestCountFinder) EXPECT() *MockRequestCountFinderMockRecorder {
	return m.recorder
}

// FindRequestCount mocks base method.
func (m *MockRequestCountFinder) FindRequestCount(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.RequestCount, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindRequestCount", varargs...)
	ret0, _ := ret[0].(*models.RequestCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRequestCount indicates an expected call of FindRequestCount.
func (mr *MockRequestCountFinderMockRecorder) FindRequestCount(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRequestCount", reflect.TypeOf((*MockRequestCountFinder)(nil).FindRequestCount), varargs...)
}

// MockRequestCountInserter is a mock of RequestCountInserter interface.
type MockRequestCountInserter struct {
	ctrl     *gomock.Controller
	recorder *MockRequestCountInserterMockRecorder
}

// MockRequestCountInserterMockRecorder is the mock recorder for MockRequestCountInserter.
type MockRequestCountInserterMockRecorder struct {
	mock *MockRequestCountInserter
}

// NewMockRequestCountInserter creates a new mock instance.
func NewMockRequestCountInserter(ctrl *gomock.Controller) *MockRequestCountInserter {
	mock := &MockRequestCountInserter{ctrl: ctrl}
	mock.recorder = &MockRequestCountInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestCountInserter) EXPECT() *MockRequestCountInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockRequestCountInserter) Insert(o *models.RequestCount, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRequestCountInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRequestCountInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockRequestCountUpdater is a mock of RequestCountUpdater interface.
type MockRequestCountUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockRequestCountUpdaterMockRecorder
}

// MockRequestCountUpdaterMockRecorder is the mock recorder for MockRequestCountUpdater.
type MockRequestCountUpdaterMockRecorder struct {
	mock *MockRequestCountUpdater
}

// NewMockRequestCountUpdater creates a new mock instance.
func NewMockRequestCountUpdater(ctrl *gomock.Controller) *MockRequestCountUpdater {
	mock := &MockRequestCountUpdater{ctrl: ctrl}
	mock.recorder = &MockRequestCountUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestCountUpdater) EXPECT() *MockRequestCountUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockRequestCountUpdater) Update(o *models.RequestCount, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRequestCountUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRequestCountUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockRequestCountUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockRequestCountUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockRequestCountUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockRequestCountUpdater) UpdateAllSlice(o models.RequestCountSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockRequestCountUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockRequestCountUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockRequestCountDeleter is a mock of RequestCountDeleter interface.
type MockRequestCountDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockRequestCountDeleterMockRecorder
}

// MockRequestCountDeleterMockRecorder is the mock recorder for MockRequestCountDeleter.
type MockRequestCountDeleterMockRecorder struct {
	mock *MockRequestCountDeleter
}

// NewMockRequestCountDeleter creates a new mock instance.
func NewMockRequestCountDeleter(ctrl *gomock.Controller) *MockRequestCountDeleter {
	mock := &MockRequestCountDeleter{ctrl: ctrl}
	mock.recorder = &MockRequestCountDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestCountDeleter) EXPECT() *MockRequestCountDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRequestCountDeleter) Delete(o *models.RequestCount, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRequestCountDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRequestCountDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockRequestCountDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockRequestCountDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockRequestCountDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockRequestCountDeleter) DeleteAllSlice(o models.RequestCountSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockRequestCountDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockRequestCountDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockRequestCountReloader is a mock of RequestCountReloader interface.
type MockRequestCountReloader struct {
	ctrl     *gomock.Controller
	recorder *MockRequestCountReloaderMockRecorder
}

// MockRequestCountReloaderMockRecorder is the mock recorder for MockRequestCountReloader.
type MockRequestCountReloaderMockRecorder struct {
	mock *MockRequestCountReloader
}

// NewMockRequestCountReloader creates a new mock instance.
func NewMockRequestCountReloader(ctrl *gomock.Controller) *MockRequestCountReloader {
	mock := &MockRequestCountReloader{ctrl: ctrl}
	mock.recorder = &MockRequestCountReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestCountReloader) EXPECT() *MockRequestCountReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockRequestCountReloader) Reload(o *models.RequestCount, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockRequestCountReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockRequestCountReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockRequestCountReloader) ReloadAll(o *models.RequestCountSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockRequestCountReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockRequestCountReloader)(nil).ReloadAll), o, ctx, exec)
}
