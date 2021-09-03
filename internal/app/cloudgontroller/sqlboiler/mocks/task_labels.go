// Code generated by MockGen. DO NOT EDIT.
// Source: psql_task_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockTaskLabelFinisher is a mock of TaskLabelFinisher interface.
type MockTaskLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockTaskLabelFinisherMockRecorder
}

// MockTaskLabelFinisherMockRecorder is the mock recorder for MockTaskLabelFinisher.
type MockTaskLabelFinisherMockRecorder struct {
	mock *MockTaskLabelFinisher
}

// NewMockTaskLabelFinisher creates a new mock instance.
func NewMockTaskLabelFinisher(ctrl *gomock.Controller) *MockTaskLabelFinisher {
	mock := &MockTaskLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockTaskLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskLabelFinisher) EXPECT() *MockTaskLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockTaskLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.TaskLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.TaskLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockTaskLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockTaskLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockTaskLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockTaskLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockTaskLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockTaskLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockTaskLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockTaskLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockTaskLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.TaskLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.TaskLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockTaskLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockTaskLabelFinisher)(nil).One), ctx, exec)
}

// MockTaskLabelFinder is a mock of TaskLabelFinder interface.
type MockTaskLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockTaskLabelFinderMockRecorder
}

// MockTaskLabelFinderMockRecorder is the mock recorder for MockTaskLabelFinder.
type MockTaskLabelFinderMockRecorder struct {
	mock *MockTaskLabelFinder
}

// NewMockTaskLabelFinder creates a new mock instance.
func NewMockTaskLabelFinder(ctrl *gomock.Controller) *MockTaskLabelFinder {
	mock := &MockTaskLabelFinder{ctrl: ctrl}
	mock.recorder = &MockTaskLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskLabelFinder) EXPECT() *MockTaskLabelFinderMockRecorder {
	return m.recorder
}

// FindTaskLabel mocks base method.
func (m *MockTaskLabelFinder) FindTaskLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.TaskLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindTaskLabel", varargs...)
	ret0, _ := ret[0].(*models.TaskLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTaskLabel indicates an expected call of FindTaskLabel.
func (mr *MockTaskLabelFinderMockRecorder) FindTaskLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTaskLabel", reflect.TypeOf((*MockTaskLabelFinder)(nil).FindTaskLabel), varargs...)
}

// MockTaskLabelInserter is a mock of TaskLabelInserter interface.
type MockTaskLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockTaskLabelInserterMockRecorder
}

// MockTaskLabelInserterMockRecorder is the mock recorder for MockTaskLabelInserter.
type MockTaskLabelInserterMockRecorder struct {
	mock *MockTaskLabelInserter
}

// NewMockTaskLabelInserter creates a new mock instance.
func NewMockTaskLabelInserter(ctrl *gomock.Controller) *MockTaskLabelInserter {
	mock := &MockTaskLabelInserter{ctrl: ctrl}
	mock.recorder = &MockTaskLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskLabelInserter) EXPECT() *MockTaskLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockTaskLabelInserter) Insert(o *models.TaskLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockTaskLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTaskLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockTaskLabelUpdater is a mock of TaskLabelUpdater interface.
type MockTaskLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockTaskLabelUpdaterMockRecorder
}

// MockTaskLabelUpdaterMockRecorder is the mock recorder for MockTaskLabelUpdater.
type MockTaskLabelUpdaterMockRecorder struct {
	mock *MockTaskLabelUpdater
}

// NewMockTaskLabelUpdater creates a new mock instance.
func NewMockTaskLabelUpdater(ctrl *gomock.Controller) *MockTaskLabelUpdater {
	mock := &MockTaskLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockTaskLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskLabelUpdater) EXPECT() *MockTaskLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockTaskLabelUpdater) Update(o *models.TaskLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTaskLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockTaskLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockTaskLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockTaskLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockTaskLabelUpdater) UpdateAllSlice(o models.TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockTaskLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockTaskLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockTaskLabelUpserter is a mock of TaskLabelUpserter interface.
type MockTaskLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockTaskLabelUpserterMockRecorder
}

// MockTaskLabelUpserterMockRecorder is the mock recorder for MockTaskLabelUpserter.
type MockTaskLabelUpserterMockRecorder struct {
	mock *MockTaskLabelUpserter
}

// NewMockTaskLabelUpserter creates a new mock instance.
func NewMockTaskLabelUpserter(ctrl *gomock.Controller) *MockTaskLabelUpserter {
	mock := &MockTaskLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockTaskLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskLabelUpserter) EXPECT() *MockTaskLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockTaskLabelUpserter) Upsert(o *models.TaskLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockTaskLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTaskLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockTaskLabelDeleter is a mock of TaskLabelDeleter interface.
type MockTaskLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockTaskLabelDeleterMockRecorder
}

// MockTaskLabelDeleterMockRecorder is the mock recorder for MockTaskLabelDeleter.
type MockTaskLabelDeleterMockRecorder struct {
	mock *MockTaskLabelDeleter
}

// NewMockTaskLabelDeleter creates a new mock instance.
func NewMockTaskLabelDeleter(ctrl *gomock.Controller) *MockTaskLabelDeleter {
	mock := &MockTaskLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockTaskLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskLabelDeleter) EXPECT() *MockTaskLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTaskLabelDeleter) Delete(o *models.TaskLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTaskLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTaskLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockTaskLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockTaskLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockTaskLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockTaskLabelDeleter) DeleteAllSlice(o models.TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockTaskLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockTaskLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockTaskLabelReloader is a mock of TaskLabelReloader interface.
type MockTaskLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockTaskLabelReloaderMockRecorder
}

// MockTaskLabelReloaderMockRecorder is the mock recorder for MockTaskLabelReloader.
type MockTaskLabelReloaderMockRecorder struct {
	mock *MockTaskLabelReloader
}

// NewMockTaskLabelReloader creates a new mock instance.
func NewMockTaskLabelReloader(ctrl *gomock.Controller) *MockTaskLabelReloader {
	mock := &MockTaskLabelReloader{ctrl: ctrl}
	mock.recorder = &MockTaskLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskLabelReloader) EXPECT() *MockTaskLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockTaskLabelReloader) Reload(o *models.TaskLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockTaskLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockTaskLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockTaskLabelReloader) ReloadAll(o *models.TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockTaskLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockTaskLabelReloader)(nil).ReloadAll), o, ctx, exec)
}
