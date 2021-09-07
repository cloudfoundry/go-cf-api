// +build unit

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_tasks.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockTaskFinisher is a mock of TaskFinisher interface.
type MockTaskFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockTaskFinisherMockRecorder
}

// MockTaskFinisherMockRecorder is the mock recorder for MockTaskFinisher.
type MockTaskFinisherMockRecorder struct {
	mock *MockTaskFinisher
}

// NewMockTaskFinisher creates a new mock instance.
func NewMockTaskFinisher(ctrl *gomock.Controller) *MockTaskFinisher {
	mock := &MockTaskFinisher{ctrl: ctrl}
	mock.recorder = &MockTaskFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskFinisher) EXPECT() *MockTaskFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockTaskFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.TaskSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.TaskSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockTaskFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockTaskFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockTaskFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockTaskFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockTaskFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockTaskFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockTaskFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockTaskFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockTaskFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockTaskFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockTaskFinisher)(nil).One), ctx, exec)
}

// MockTaskFinder is a mock of TaskFinder interface.
type MockTaskFinder struct {
	ctrl     *gomock.Controller
	recorder *MockTaskFinderMockRecorder
}

// MockTaskFinderMockRecorder is the mock recorder for MockTaskFinder.
type MockTaskFinderMockRecorder struct {
	mock *MockTaskFinder
}

// NewMockTaskFinder creates a new mock instance.
func NewMockTaskFinder(ctrl *gomock.Controller) *MockTaskFinder {
	mock := &MockTaskFinder{ctrl: ctrl}
	mock.recorder = &MockTaskFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskFinder) EXPECT() *MockTaskFinderMockRecorder {
	return m.recorder
}

// FindTask mocks base method.
func (m *MockTaskFinder) FindTask(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Task, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindTask", varargs...)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTask indicates an expected call of FindTask.
func (mr *MockTaskFinderMockRecorder) FindTask(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTask", reflect.TypeOf((*MockTaskFinder)(nil).FindTask), varargs...)
}

// MockTaskInserter is a mock of TaskInserter interface.
type MockTaskInserter struct {
	ctrl     *gomock.Controller
	recorder *MockTaskInserterMockRecorder
}

// MockTaskInserterMockRecorder is the mock recorder for MockTaskInserter.
type MockTaskInserterMockRecorder struct {
	mock *MockTaskInserter
}

// NewMockTaskInserter creates a new mock instance.
func NewMockTaskInserter(ctrl *gomock.Controller) *MockTaskInserter {
	mock := &MockTaskInserter{ctrl: ctrl}
	mock.recorder = &MockTaskInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskInserter) EXPECT() *MockTaskInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockTaskInserter) Insert(o *models.Task, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockTaskInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTaskInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockTaskUpdater is a mock of TaskUpdater interface.
type MockTaskUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUpdaterMockRecorder
}

// MockTaskUpdaterMockRecorder is the mock recorder for MockTaskUpdater.
type MockTaskUpdaterMockRecorder struct {
	mock *MockTaskUpdater
}

// NewMockTaskUpdater creates a new mock instance.
func NewMockTaskUpdater(ctrl *gomock.Controller) *MockTaskUpdater {
	mock := &MockTaskUpdater{ctrl: ctrl}
	mock.recorder = &MockTaskUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUpdater) EXPECT() *MockTaskUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockTaskUpdater) Update(o *models.Task, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTaskUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockTaskUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockTaskUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockTaskUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockTaskUpdater) UpdateAllSlice(o models.TaskSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockTaskUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockTaskUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockTaskUpserter is a mock of TaskUpserter interface.
type MockTaskUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockTaskUpserterMockRecorder
}

// MockTaskUpserterMockRecorder is the mock recorder for MockTaskUpserter.
type MockTaskUpserterMockRecorder struct {
	mock *MockTaskUpserter
}

// NewMockTaskUpserter creates a new mock instance.
func NewMockTaskUpserter(ctrl *gomock.Controller) *MockTaskUpserter {
	mock := &MockTaskUpserter{ctrl: ctrl}
	mock.recorder = &MockTaskUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskUpserter) EXPECT() *MockTaskUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockTaskUpserter) Upsert(o *models.Task, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockTaskUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTaskUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockTaskDeleter is a mock of TaskDeleter interface.
type MockTaskDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockTaskDeleterMockRecorder
}

// MockTaskDeleterMockRecorder is the mock recorder for MockTaskDeleter.
type MockTaskDeleterMockRecorder struct {
	mock *MockTaskDeleter
}

// NewMockTaskDeleter creates a new mock instance.
func NewMockTaskDeleter(ctrl *gomock.Controller) *MockTaskDeleter {
	mock := &MockTaskDeleter{ctrl: ctrl}
	mock.recorder = &MockTaskDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskDeleter) EXPECT() *MockTaskDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTaskDeleter) Delete(o *models.Task, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTaskDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTaskDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockTaskDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockTaskDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockTaskDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockTaskDeleter) DeleteAllSlice(o models.TaskSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockTaskDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockTaskDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockTaskReloader is a mock of TaskReloader interface.
type MockTaskReloader struct {
	ctrl     *gomock.Controller
	recorder *MockTaskReloaderMockRecorder
}

// MockTaskReloaderMockRecorder is the mock recorder for MockTaskReloader.
type MockTaskReloaderMockRecorder struct {
	mock *MockTaskReloader
}

// NewMockTaskReloader creates a new mock instance.
func NewMockTaskReloader(ctrl *gomock.Controller) *MockTaskReloader {
	mock := &MockTaskReloader{ctrl: ctrl}
	mock.recorder = &MockTaskReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskReloader) EXPECT() *MockTaskReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockTaskReloader) Reload(o *models.Task, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockTaskReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockTaskReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockTaskReloader) ReloadAll(o *models.TaskSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockTaskReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockTaskReloader)(nil).ReloadAll), o, ctx, exec)
}
