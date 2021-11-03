// +build unit

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_builds.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockBuildUpserter is a mock of BuildUpserter interface.
type MockBuildUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildUpserterMockRecorder
}

// MockBuildUpserterMockRecorder is the mock recorder for MockBuildUpserter.
type MockBuildUpserterMockRecorder struct {
	mock *MockBuildUpserter
}

// NewMockBuildUpserter creates a new mock instance.
func NewMockBuildUpserter(ctrl *gomock.Controller) *MockBuildUpserter {
	mock := &MockBuildUpserter{ctrl: ctrl}
	mock.recorder = &MockBuildUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildUpserter) EXPECT() *MockBuildUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockBuildUpserter) Upsert(o *models.Build, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockBuildUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockBuildUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockBuildFinisher is a mock of BuildFinisher interface.
type MockBuildFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockBuildFinisherMockRecorder
}

// MockBuildFinisherMockRecorder is the mock recorder for MockBuildFinisher.
type MockBuildFinisherMockRecorder struct {
	mock *MockBuildFinisher
}

// NewMockBuildFinisher creates a new mock instance.
func NewMockBuildFinisher(ctrl *gomock.Controller) *MockBuildFinisher {
	mock := &MockBuildFinisher{ctrl: ctrl}
	mock.recorder = &MockBuildFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildFinisher) EXPECT() *MockBuildFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockBuildFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.BuildSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.BuildSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockBuildFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockBuildFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockBuildFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockBuildFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockBuildFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockBuildFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockBuildFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockBuildFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockBuildFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Build, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockBuildFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockBuildFinisher)(nil).One), ctx, exec)
}

// MockBuildFinder is a mock of BuildFinder interface.
type MockBuildFinder struct {
	ctrl     *gomock.Controller
	recorder *MockBuildFinderMockRecorder
}

// MockBuildFinderMockRecorder is the mock recorder for MockBuildFinder.
type MockBuildFinderMockRecorder struct {
	mock *MockBuildFinder
}

// NewMockBuildFinder creates a new mock instance.
func NewMockBuildFinder(ctrl *gomock.Controller) *MockBuildFinder {
	mock := &MockBuildFinder{ctrl: ctrl}
	mock.recorder = &MockBuildFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildFinder) EXPECT() *MockBuildFinderMockRecorder {
	return m.recorder
}

// FindBuild mocks base method.
func (m *MockBuildFinder) FindBuild(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Build, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindBuild", varargs...)
	ret0, _ := ret[0].(*models.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBuild indicates an expected call of FindBuild.
func (mr *MockBuildFinderMockRecorder) FindBuild(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBuild", reflect.TypeOf((*MockBuildFinder)(nil).FindBuild), varargs...)
}

// MockBuildInserter is a mock of BuildInserter interface.
type MockBuildInserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildInserterMockRecorder
}

// MockBuildInserterMockRecorder is the mock recorder for MockBuildInserter.
type MockBuildInserterMockRecorder struct {
	mock *MockBuildInserter
}

// NewMockBuildInserter creates a new mock instance.
func NewMockBuildInserter(ctrl *gomock.Controller) *MockBuildInserter {
	mock := &MockBuildInserter{ctrl: ctrl}
	mock.recorder = &MockBuildInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildInserter) EXPECT() *MockBuildInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockBuildInserter) Insert(o *models.Build, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockBuildInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBuildInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockBuildUpdater is a mock of BuildUpdater interface.
type MockBuildUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockBuildUpdaterMockRecorder
}

// MockBuildUpdaterMockRecorder is the mock recorder for MockBuildUpdater.
type MockBuildUpdaterMockRecorder struct {
	mock *MockBuildUpdater
}

// NewMockBuildUpdater creates a new mock instance.
func NewMockBuildUpdater(ctrl *gomock.Controller) *MockBuildUpdater {
	mock := &MockBuildUpdater{ctrl: ctrl}
	mock.recorder = &MockBuildUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildUpdater) EXPECT() *MockBuildUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockBuildUpdater) Update(o *models.Build, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBuildUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBuildUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockBuildUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockBuildUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockBuildUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockBuildUpdater) UpdateAllSlice(o models.BuildSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockBuildUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockBuildUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockBuildDeleter is a mock of BuildDeleter interface.
type MockBuildDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildDeleterMockRecorder
}

// MockBuildDeleterMockRecorder is the mock recorder for MockBuildDeleter.
type MockBuildDeleterMockRecorder struct {
	mock *MockBuildDeleter
}

// NewMockBuildDeleter creates a new mock instance.
func NewMockBuildDeleter(ctrl *gomock.Controller) *MockBuildDeleter {
	mock := &MockBuildDeleter{ctrl: ctrl}
	mock.recorder = &MockBuildDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildDeleter) EXPECT() *MockBuildDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBuildDeleter) Delete(o *models.Build, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBuildDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBuildDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockBuildDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockBuildDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockBuildDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockBuildDeleter) DeleteAllSlice(o models.BuildSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockBuildDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockBuildDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockBuildReloader is a mock of BuildReloader interface.
type MockBuildReloader struct {
	ctrl     *gomock.Controller
	recorder *MockBuildReloaderMockRecorder
}

// MockBuildReloaderMockRecorder is the mock recorder for MockBuildReloader.
type MockBuildReloaderMockRecorder struct {
	mock *MockBuildReloader
}

// NewMockBuildReloader creates a new mock instance.
func NewMockBuildReloader(ctrl *gomock.Controller) *MockBuildReloader {
	mock := &MockBuildReloader{ctrl: ctrl}
	mock.recorder = &MockBuildReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildReloader) EXPECT() *MockBuildReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockBuildReloader) Reload(o *models.Build, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockBuildReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockBuildReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockBuildReloader) ReloadAll(o *models.BuildSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockBuildReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockBuildReloader)(nil).ReloadAll), o, ctx, exec)
}
