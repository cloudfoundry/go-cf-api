// Code generated by MockGen. DO NOT EDIT.
// Source: psql_spaces_developers.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockSpacesDeveloperFinisher is a mock of SpacesDeveloperFinisher interface.
type MockSpacesDeveloperFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesDeveloperFinisherMockRecorder
}

// MockSpacesDeveloperFinisherMockRecorder is the mock recorder for MockSpacesDeveloperFinisher.
type MockSpacesDeveloperFinisherMockRecorder struct {
	mock *MockSpacesDeveloperFinisher
}

// NewMockSpacesDeveloperFinisher creates a new mock instance.
func NewMockSpacesDeveloperFinisher(ctrl *gomock.Controller) *MockSpacesDeveloperFinisher {
	mock := &MockSpacesDeveloperFinisher{ctrl: ctrl}
	mock.recorder = &MockSpacesDeveloperFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesDeveloperFinisher) EXPECT() *MockSpacesDeveloperFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockSpacesDeveloperFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.SpacesDeveloperSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.SpacesDeveloperSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockSpacesDeveloperFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockSpacesDeveloperFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockSpacesDeveloperFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSpacesDeveloperFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSpacesDeveloperFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockSpacesDeveloperFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSpacesDeveloperFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSpacesDeveloperFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockSpacesDeveloperFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.SpacesDeveloper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.SpacesDeveloper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockSpacesDeveloperFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockSpacesDeveloperFinisher)(nil).One), ctx, exec)
}

// MockSpacesDeveloperFinder is a mock of SpacesDeveloperFinder interface.
type MockSpacesDeveloperFinder struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesDeveloperFinderMockRecorder
}

// MockSpacesDeveloperFinderMockRecorder is the mock recorder for MockSpacesDeveloperFinder.
type MockSpacesDeveloperFinderMockRecorder struct {
	mock *MockSpacesDeveloperFinder
}

// NewMockSpacesDeveloperFinder creates a new mock instance.
func NewMockSpacesDeveloperFinder(ctrl *gomock.Controller) *MockSpacesDeveloperFinder {
	mock := &MockSpacesDeveloperFinder{ctrl: ctrl}
	mock.recorder = &MockSpacesDeveloperFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesDeveloperFinder) EXPECT() *MockSpacesDeveloperFinderMockRecorder {
	return m.recorder
}

// FindSpacesDeveloper mocks base method.
func (m *MockSpacesDeveloperFinder) FindSpacesDeveloper(ctx context.Context, exec boil.ContextExecutor, spacesDevelopersPK int, selectCols ...string) (*models.SpacesDeveloper, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, spacesDevelopersPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSpacesDeveloper", varargs...)
	ret0, _ := ret[0].(*models.SpacesDeveloper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSpacesDeveloper indicates an expected call of FindSpacesDeveloper.
func (mr *MockSpacesDeveloperFinderMockRecorder) FindSpacesDeveloper(ctx, exec, spacesDevelopersPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, spacesDevelopersPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSpacesDeveloper", reflect.TypeOf((*MockSpacesDeveloperFinder)(nil).FindSpacesDeveloper), varargs...)
}

// MockSpacesDeveloperInserter is a mock of SpacesDeveloperInserter interface.
type MockSpacesDeveloperInserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesDeveloperInserterMockRecorder
}

// MockSpacesDeveloperInserterMockRecorder is the mock recorder for MockSpacesDeveloperInserter.
type MockSpacesDeveloperInserterMockRecorder struct {
	mock *MockSpacesDeveloperInserter
}

// NewMockSpacesDeveloperInserter creates a new mock instance.
func NewMockSpacesDeveloperInserter(ctrl *gomock.Controller) *MockSpacesDeveloperInserter {
	mock := &MockSpacesDeveloperInserter{ctrl: ctrl}
	mock.recorder = &MockSpacesDeveloperInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesDeveloperInserter) EXPECT() *MockSpacesDeveloperInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockSpacesDeveloperInserter) Insert(o *models.SpacesDeveloper, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSpacesDeveloperInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSpacesDeveloperInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockSpacesDeveloperUpdater is a mock of SpacesDeveloperUpdater interface.
type MockSpacesDeveloperUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesDeveloperUpdaterMockRecorder
}

// MockSpacesDeveloperUpdaterMockRecorder is the mock recorder for MockSpacesDeveloperUpdater.
type MockSpacesDeveloperUpdaterMockRecorder struct {
	mock *MockSpacesDeveloperUpdater
}

// NewMockSpacesDeveloperUpdater creates a new mock instance.
func NewMockSpacesDeveloperUpdater(ctrl *gomock.Controller) *MockSpacesDeveloperUpdater {
	mock := &MockSpacesDeveloperUpdater{ctrl: ctrl}
	mock.recorder = &MockSpacesDeveloperUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesDeveloperUpdater) EXPECT() *MockSpacesDeveloperUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockSpacesDeveloperUpdater) Update(o *models.SpacesDeveloper, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSpacesDeveloperUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSpacesDeveloperUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockSpacesDeveloperUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockSpacesDeveloperUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockSpacesDeveloperUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockSpacesDeveloperUpdater) UpdateAllSlice(o models.SpacesDeveloperSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockSpacesDeveloperUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockSpacesDeveloperUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockSpacesDeveloperDeleter is a mock of SpacesDeveloperDeleter interface.
type MockSpacesDeveloperDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesDeveloperDeleterMockRecorder
}

// MockSpacesDeveloperDeleterMockRecorder is the mock recorder for MockSpacesDeveloperDeleter.
type MockSpacesDeveloperDeleterMockRecorder struct {
	mock *MockSpacesDeveloperDeleter
}

// NewMockSpacesDeveloperDeleter creates a new mock instance.
func NewMockSpacesDeveloperDeleter(ctrl *gomock.Controller) *MockSpacesDeveloperDeleter {
	mock := &MockSpacesDeveloperDeleter{ctrl: ctrl}
	mock.recorder = &MockSpacesDeveloperDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesDeveloperDeleter) EXPECT() *MockSpacesDeveloperDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSpacesDeveloperDeleter) Delete(o *models.SpacesDeveloper, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSpacesDeveloperDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSpacesDeveloperDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockSpacesDeveloperDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockSpacesDeveloperDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockSpacesDeveloperDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockSpacesDeveloperDeleter) DeleteAllSlice(o models.SpacesDeveloperSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockSpacesDeveloperDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockSpacesDeveloperDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockSpacesDeveloperReloader is a mock of SpacesDeveloperReloader interface.
type MockSpacesDeveloperReloader struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesDeveloperReloaderMockRecorder
}

// MockSpacesDeveloperReloaderMockRecorder is the mock recorder for MockSpacesDeveloperReloader.
type MockSpacesDeveloperReloaderMockRecorder struct {
	mock *MockSpacesDeveloperReloader
}

// NewMockSpacesDeveloperReloader creates a new mock instance.
func NewMockSpacesDeveloperReloader(ctrl *gomock.Controller) *MockSpacesDeveloperReloader {
	mock := &MockSpacesDeveloperReloader{ctrl: ctrl}
	mock.recorder = &MockSpacesDeveloperReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesDeveloperReloader) EXPECT() *MockSpacesDeveloperReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockSpacesDeveloperReloader) Reload(o *models.SpacesDeveloper, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockSpacesDeveloperReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSpacesDeveloperReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockSpacesDeveloperReloader) ReloadAll(o *models.SpacesDeveloperSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockSpacesDeveloperReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockSpacesDeveloperReloader)(nil).ReloadAll), o, ctx, exec)
}
