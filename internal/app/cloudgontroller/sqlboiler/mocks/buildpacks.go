// Code generated by MockGen. DO NOT EDIT.
// Source: psql_buildpacks.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockBuildpackFinisher is a mock of BuildpackFinisher interface.
type MockBuildpackFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackFinisherMockRecorder
}

// MockBuildpackFinisherMockRecorder is the mock recorder for MockBuildpackFinisher.
type MockBuildpackFinisherMockRecorder struct {
	mock *MockBuildpackFinisher
}

// NewMockBuildpackFinisher creates a new mock instance.
func NewMockBuildpackFinisher(ctrl *gomock.Controller) *MockBuildpackFinisher {
	mock := &MockBuildpackFinisher{ctrl: ctrl}
	mock.recorder = &MockBuildpackFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackFinisher) EXPECT() *MockBuildpackFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockBuildpackFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.BuildpackSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.BuildpackSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockBuildpackFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockBuildpackFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockBuildpackFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockBuildpackFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockBuildpackFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockBuildpackFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockBuildpackFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockBuildpackFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockBuildpackFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Buildpack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Buildpack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockBuildpackFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockBuildpackFinisher)(nil).One), ctx, exec)
}

// MockBuildpackFinder is a mock of BuildpackFinder interface.
type MockBuildpackFinder struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackFinderMockRecorder
}

// MockBuildpackFinderMockRecorder is the mock recorder for MockBuildpackFinder.
type MockBuildpackFinderMockRecorder struct {
	mock *MockBuildpackFinder
}

// NewMockBuildpackFinder creates a new mock instance.
func NewMockBuildpackFinder(ctrl *gomock.Controller) *MockBuildpackFinder {
	mock := &MockBuildpackFinder{ctrl: ctrl}
	mock.recorder = &MockBuildpackFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackFinder) EXPECT() *MockBuildpackFinderMockRecorder {
	return m.recorder
}

// FindBuildpack mocks base method.
func (m *MockBuildpackFinder) FindBuildpack(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Buildpack, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindBuildpack", varargs...)
	ret0, _ := ret[0].(*models.Buildpack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBuildpack indicates an expected call of FindBuildpack.
func (mr *MockBuildpackFinderMockRecorder) FindBuildpack(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBuildpack", reflect.TypeOf((*MockBuildpackFinder)(nil).FindBuildpack), varargs...)
}

// MockBuildpackInserter is a mock of BuildpackInserter interface.
type MockBuildpackInserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackInserterMockRecorder
}

// MockBuildpackInserterMockRecorder is the mock recorder for MockBuildpackInserter.
type MockBuildpackInserterMockRecorder struct {
	mock *MockBuildpackInserter
}

// NewMockBuildpackInserter creates a new mock instance.
func NewMockBuildpackInserter(ctrl *gomock.Controller) *MockBuildpackInserter {
	mock := &MockBuildpackInserter{ctrl: ctrl}
	mock.recorder = &MockBuildpackInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackInserter) EXPECT() *MockBuildpackInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockBuildpackInserter) Insert(o *models.Buildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockBuildpackInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBuildpackInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockBuildpackUpdater is a mock of BuildpackUpdater interface.
type MockBuildpackUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackUpdaterMockRecorder
}

// MockBuildpackUpdaterMockRecorder is the mock recorder for MockBuildpackUpdater.
type MockBuildpackUpdaterMockRecorder struct {
	mock *MockBuildpackUpdater
}

// NewMockBuildpackUpdater creates a new mock instance.
func NewMockBuildpackUpdater(ctrl *gomock.Controller) *MockBuildpackUpdater {
	mock := &MockBuildpackUpdater{ctrl: ctrl}
	mock.recorder = &MockBuildpackUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackUpdater) EXPECT() *MockBuildpackUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockBuildpackUpdater) Update(o *models.Buildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBuildpackUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBuildpackUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockBuildpackUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockBuildpackUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockBuildpackUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockBuildpackUpdater) UpdateAllSlice(o models.BuildpackSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockBuildpackUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockBuildpackUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockBuildpackUpserter is a mock of BuildpackUpserter interface.
type MockBuildpackUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackUpserterMockRecorder
}

// MockBuildpackUpserterMockRecorder is the mock recorder for MockBuildpackUpserter.
type MockBuildpackUpserterMockRecorder struct {
	mock *MockBuildpackUpserter
}

// NewMockBuildpackUpserter creates a new mock instance.
func NewMockBuildpackUpserter(ctrl *gomock.Controller) *MockBuildpackUpserter {
	mock := &MockBuildpackUpserter{ctrl: ctrl}
	mock.recorder = &MockBuildpackUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackUpserter) EXPECT() *MockBuildpackUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockBuildpackUpserter) Upsert(o *models.Buildpack, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockBuildpackUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockBuildpackUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockBuildpackDeleter is a mock of BuildpackDeleter interface.
type MockBuildpackDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackDeleterMockRecorder
}

// MockBuildpackDeleterMockRecorder is the mock recorder for MockBuildpackDeleter.
type MockBuildpackDeleterMockRecorder struct {
	mock *MockBuildpackDeleter
}

// NewMockBuildpackDeleter creates a new mock instance.
func NewMockBuildpackDeleter(ctrl *gomock.Controller) *MockBuildpackDeleter {
	mock := &MockBuildpackDeleter{ctrl: ctrl}
	mock.recorder = &MockBuildpackDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackDeleter) EXPECT() *MockBuildpackDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBuildpackDeleter) Delete(o *models.Buildpack, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBuildpackDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBuildpackDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockBuildpackDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockBuildpackDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockBuildpackDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockBuildpackDeleter) DeleteAllSlice(o models.BuildpackSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockBuildpackDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockBuildpackDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockBuildpackReloader is a mock of BuildpackReloader interface.
type MockBuildpackReloader struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackReloaderMockRecorder
}

// MockBuildpackReloaderMockRecorder is the mock recorder for MockBuildpackReloader.
type MockBuildpackReloaderMockRecorder struct {
	mock *MockBuildpackReloader
}

// NewMockBuildpackReloader creates a new mock instance.
func NewMockBuildpackReloader(ctrl *gomock.Controller) *MockBuildpackReloader {
	mock := &MockBuildpackReloader{ctrl: ctrl}
	mock.recorder = &MockBuildpackReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackReloader) EXPECT() *MockBuildpackReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockBuildpackReloader) Reload(o *models.Buildpack, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockBuildpackReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockBuildpackReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockBuildpackReloader) ReloadAll(o *models.BuildpackSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockBuildpackReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockBuildpackReloader)(nil).ReloadAll), o, ctx, exec)
}
