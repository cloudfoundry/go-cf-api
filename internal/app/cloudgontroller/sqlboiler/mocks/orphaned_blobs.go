// Code generated by MockGen. DO NOT EDIT.
// Source: psql_orphaned_blobs.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockOrphanedBlobFinisher is a mock of OrphanedBlobFinisher interface.
type MockOrphanedBlobFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanedBlobFinisherMockRecorder
}

// MockOrphanedBlobFinisherMockRecorder is the mock recorder for MockOrphanedBlobFinisher.
type MockOrphanedBlobFinisherMockRecorder struct {
	mock *MockOrphanedBlobFinisher
}

// NewMockOrphanedBlobFinisher creates a new mock instance.
func NewMockOrphanedBlobFinisher(ctrl *gomock.Controller) *MockOrphanedBlobFinisher {
	mock := &MockOrphanedBlobFinisher{ctrl: ctrl}
	mock.recorder = &MockOrphanedBlobFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanedBlobFinisher) EXPECT() *MockOrphanedBlobFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrphanedBlobFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrphanedBlobSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrphanedBlobSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrphanedBlobFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrphanedBlobFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrphanedBlobFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrphanedBlobFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrphanedBlobFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrphanedBlobFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrphanedBlobFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrphanedBlobFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrphanedBlobFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.OrphanedBlob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.OrphanedBlob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrphanedBlobFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrphanedBlobFinisher)(nil).One), ctx, exec)
}

// MockOrphanedBlobFinder is a mock of OrphanedBlobFinder interface.
type MockOrphanedBlobFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanedBlobFinderMockRecorder
}

// MockOrphanedBlobFinderMockRecorder is the mock recorder for MockOrphanedBlobFinder.
type MockOrphanedBlobFinderMockRecorder struct {
	mock *MockOrphanedBlobFinder
}

// NewMockOrphanedBlobFinder creates a new mock instance.
func NewMockOrphanedBlobFinder(ctrl *gomock.Controller) *MockOrphanedBlobFinder {
	mock := &MockOrphanedBlobFinder{ctrl: ctrl}
	mock.recorder = &MockOrphanedBlobFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanedBlobFinder) EXPECT() *MockOrphanedBlobFinderMockRecorder {
	return m.recorder
}

// FindOrphanedBlob mocks base method.
func (m *MockOrphanedBlobFinder) FindOrphanedBlob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.OrphanedBlob, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrphanedBlob", varargs...)
	ret0, _ := ret[0].(*models.OrphanedBlob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrphanedBlob indicates an expected call of FindOrphanedBlob.
func (mr *MockOrphanedBlobFinderMockRecorder) FindOrphanedBlob(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrphanedBlob", reflect.TypeOf((*MockOrphanedBlobFinder)(nil).FindOrphanedBlob), varargs...)
}

// MockOrphanedBlobInserter is a mock of OrphanedBlobInserter interface.
type MockOrphanedBlobInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanedBlobInserterMockRecorder
}

// MockOrphanedBlobInserterMockRecorder is the mock recorder for MockOrphanedBlobInserter.
type MockOrphanedBlobInserterMockRecorder struct {
	mock *MockOrphanedBlobInserter
}

// NewMockOrphanedBlobInserter creates a new mock instance.
func NewMockOrphanedBlobInserter(ctrl *gomock.Controller) *MockOrphanedBlobInserter {
	mock := &MockOrphanedBlobInserter{ctrl: ctrl}
	mock.recorder = &MockOrphanedBlobInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanedBlobInserter) EXPECT() *MockOrphanedBlobInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrphanedBlobInserter) Insert(o *models.OrphanedBlob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrphanedBlobInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrphanedBlobInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrphanedBlobUpdater is a mock of OrphanedBlobUpdater interface.
type MockOrphanedBlobUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanedBlobUpdaterMockRecorder
}

// MockOrphanedBlobUpdaterMockRecorder is the mock recorder for MockOrphanedBlobUpdater.
type MockOrphanedBlobUpdaterMockRecorder struct {
	mock *MockOrphanedBlobUpdater
}

// NewMockOrphanedBlobUpdater creates a new mock instance.
func NewMockOrphanedBlobUpdater(ctrl *gomock.Controller) *MockOrphanedBlobUpdater {
	mock := &MockOrphanedBlobUpdater{ctrl: ctrl}
	mock.recorder = &MockOrphanedBlobUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanedBlobUpdater) EXPECT() *MockOrphanedBlobUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrphanedBlobUpdater) Update(o *models.OrphanedBlob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrphanedBlobUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrphanedBlobUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrphanedBlobUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrphanedBlobUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrphanedBlobUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrphanedBlobUpdater) UpdateAllSlice(o models.OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrphanedBlobUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrphanedBlobUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrphanedBlobDeleter is a mock of OrphanedBlobDeleter interface.
type MockOrphanedBlobDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanedBlobDeleterMockRecorder
}

// MockOrphanedBlobDeleterMockRecorder is the mock recorder for MockOrphanedBlobDeleter.
type MockOrphanedBlobDeleterMockRecorder struct {
	mock *MockOrphanedBlobDeleter
}

// NewMockOrphanedBlobDeleter creates a new mock instance.
func NewMockOrphanedBlobDeleter(ctrl *gomock.Controller) *MockOrphanedBlobDeleter {
	mock := &MockOrphanedBlobDeleter{ctrl: ctrl}
	mock.recorder = &MockOrphanedBlobDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanedBlobDeleter) EXPECT() *MockOrphanedBlobDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrphanedBlobDeleter) Delete(o *models.OrphanedBlob, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrphanedBlobDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrphanedBlobDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrphanedBlobDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrphanedBlobDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrphanedBlobDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrphanedBlobDeleter) DeleteAllSlice(o models.OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrphanedBlobDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrphanedBlobDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrphanedBlobReloader is a mock of OrphanedBlobReloader interface.
type MockOrphanedBlobReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanedBlobReloaderMockRecorder
}

// MockOrphanedBlobReloaderMockRecorder is the mock recorder for MockOrphanedBlobReloader.
type MockOrphanedBlobReloaderMockRecorder struct {
	mock *MockOrphanedBlobReloader
}

// NewMockOrphanedBlobReloader creates a new mock instance.
func NewMockOrphanedBlobReloader(ctrl *gomock.Controller) *MockOrphanedBlobReloader {
	mock := &MockOrphanedBlobReloader{ctrl: ctrl}
	mock.recorder = &MockOrphanedBlobReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanedBlobReloader) EXPECT() *MockOrphanedBlobReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrphanedBlobReloader) Reload(o *models.OrphanedBlob, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrphanedBlobReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrphanedBlobReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrphanedBlobReloader) ReloadAll(o *models.OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrphanedBlobReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrphanedBlobReloader)(nil).ReloadAll), o, ctx, exec)
}
