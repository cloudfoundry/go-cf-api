// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_instance_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockServiceInstanceLabelUpserter is a mock of ServiceInstanceLabelUpserter interface.
type MockServiceInstanceLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstanceLabelUpserterMockRecorder
}

// MockServiceInstanceLabelUpserterMockRecorder is the mock recorder for MockServiceInstanceLabelUpserter.
type MockServiceInstanceLabelUpserterMockRecorder struct {
	mock *MockServiceInstanceLabelUpserter
}

// NewMockServiceInstanceLabelUpserter creates a new mock instance.
func NewMockServiceInstanceLabelUpserter(ctrl *gomock.Controller) *MockServiceInstanceLabelUpserter {
	mock := &MockServiceInstanceLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceInstanceLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInstanceLabelUpserter) EXPECT() *MockServiceInstanceLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceInstanceLabelUpserter) Upsert(o *models.ServiceInstanceLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceInstanceLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceInstanceLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceInstanceLabelFinisher is a mock of ServiceInstanceLabelFinisher interface.
type MockServiceInstanceLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstanceLabelFinisherMockRecorder
}

// MockServiceInstanceLabelFinisherMockRecorder is the mock recorder for MockServiceInstanceLabelFinisher.
type MockServiceInstanceLabelFinisherMockRecorder struct {
	mock *MockServiceInstanceLabelFinisher
}

// NewMockServiceInstanceLabelFinisher creates a new mock instance.
func NewMockServiceInstanceLabelFinisher(ctrl *gomock.Controller) *MockServiceInstanceLabelFinisher {
	mock := &MockServiceInstanceLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceInstanceLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInstanceLabelFinisher) EXPECT() *MockServiceInstanceLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceInstanceLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceInstanceLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceInstanceLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceInstanceLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceInstanceLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceInstanceLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceInstanceLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceInstanceLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceInstanceLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceInstanceLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceInstanceLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceInstanceLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceInstanceLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceInstanceLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceInstanceLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceInstanceLabelFinisher)(nil).One), ctx, exec)
}

// MockServiceInstanceLabelFinder is a mock of ServiceInstanceLabelFinder interface.
type MockServiceInstanceLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstanceLabelFinderMockRecorder
}

// MockServiceInstanceLabelFinderMockRecorder is the mock recorder for MockServiceInstanceLabelFinder.
type MockServiceInstanceLabelFinderMockRecorder struct {
	mock *MockServiceInstanceLabelFinder
}

// NewMockServiceInstanceLabelFinder creates a new mock instance.
func NewMockServiceInstanceLabelFinder(ctrl *gomock.Controller) *MockServiceInstanceLabelFinder {
	mock := &MockServiceInstanceLabelFinder{ctrl: ctrl}
	mock.recorder = &MockServiceInstanceLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInstanceLabelFinder) EXPECT() *MockServiceInstanceLabelFinderMockRecorder {
	return m.recorder
}

// FindServiceInstanceLabel mocks base method.
func (m *MockServiceInstanceLabelFinder) FindServiceInstanceLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceInstanceLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceInstanceLabel", varargs...)
	ret0, _ := ret[0].(*models.ServiceInstanceLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceInstanceLabel indicates an expected call of FindServiceInstanceLabel.
func (mr *MockServiceInstanceLabelFinderMockRecorder) FindServiceInstanceLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceInstanceLabel", reflect.TypeOf((*MockServiceInstanceLabelFinder)(nil).FindServiceInstanceLabel), varargs...)
}

// MockServiceInstanceLabelInserter is a mock of ServiceInstanceLabelInserter interface.
type MockServiceInstanceLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstanceLabelInserterMockRecorder
}

// MockServiceInstanceLabelInserterMockRecorder is the mock recorder for MockServiceInstanceLabelInserter.
type MockServiceInstanceLabelInserterMockRecorder struct {
	mock *MockServiceInstanceLabelInserter
}

// NewMockServiceInstanceLabelInserter creates a new mock instance.
func NewMockServiceInstanceLabelInserter(ctrl *gomock.Controller) *MockServiceInstanceLabelInserter {
	mock := &MockServiceInstanceLabelInserter{ctrl: ctrl}
	mock.recorder = &MockServiceInstanceLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInstanceLabelInserter) EXPECT() *MockServiceInstanceLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceInstanceLabelInserter) Insert(o *models.ServiceInstanceLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceInstanceLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceInstanceLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceInstanceLabelUpdater is a mock of ServiceInstanceLabelUpdater interface.
type MockServiceInstanceLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstanceLabelUpdaterMockRecorder
}

// MockServiceInstanceLabelUpdaterMockRecorder is the mock recorder for MockServiceInstanceLabelUpdater.
type MockServiceInstanceLabelUpdaterMockRecorder struct {
	mock *MockServiceInstanceLabelUpdater
}

// NewMockServiceInstanceLabelUpdater creates a new mock instance.
func NewMockServiceInstanceLabelUpdater(ctrl *gomock.Controller) *MockServiceInstanceLabelUpdater {
	mock := &MockServiceInstanceLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceInstanceLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInstanceLabelUpdater) EXPECT() *MockServiceInstanceLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceInstanceLabelUpdater) Update(o *models.ServiceInstanceLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceInstanceLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceInstanceLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceInstanceLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceInstanceLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceInstanceLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceInstanceLabelUpdater) UpdateAllSlice(o models.ServiceInstanceLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceInstanceLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceInstanceLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceInstanceLabelDeleter is a mock of ServiceInstanceLabelDeleter interface.
type MockServiceInstanceLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstanceLabelDeleterMockRecorder
}

// MockServiceInstanceLabelDeleterMockRecorder is the mock recorder for MockServiceInstanceLabelDeleter.
type MockServiceInstanceLabelDeleterMockRecorder struct {
	mock *MockServiceInstanceLabelDeleter
}

// NewMockServiceInstanceLabelDeleter creates a new mock instance.
func NewMockServiceInstanceLabelDeleter(ctrl *gomock.Controller) *MockServiceInstanceLabelDeleter {
	mock := &MockServiceInstanceLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceInstanceLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInstanceLabelDeleter) EXPECT() *MockServiceInstanceLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceInstanceLabelDeleter) Delete(o *models.ServiceInstanceLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceInstanceLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceInstanceLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceInstanceLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceInstanceLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceInstanceLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceInstanceLabelDeleter) DeleteAllSlice(o models.ServiceInstanceLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceInstanceLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceInstanceLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceInstanceLabelReloader is a mock of ServiceInstanceLabelReloader interface.
type MockServiceInstanceLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInstanceLabelReloaderMockRecorder
}

// MockServiceInstanceLabelReloaderMockRecorder is the mock recorder for MockServiceInstanceLabelReloader.
type MockServiceInstanceLabelReloaderMockRecorder struct {
	mock *MockServiceInstanceLabelReloader
}

// NewMockServiceInstanceLabelReloader creates a new mock instance.
func NewMockServiceInstanceLabelReloader(ctrl *gomock.Controller) *MockServiceInstanceLabelReloader {
	mock := &MockServiceInstanceLabelReloader{ctrl: ctrl}
	mock.recorder = &MockServiceInstanceLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInstanceLabelReloader) EXPECT() *MockServiceInstanceLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceInstanceLabelReloader) Reload(o *models.ServiceInstanceLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceInstanceLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceInstanceLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceInstanceLabelReloader) ReloadAll(o *models.ServiceInstanceLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceInstanceLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceInstanceLabelReloader)(nil).ReloadAll), o, ctx, exec)
}