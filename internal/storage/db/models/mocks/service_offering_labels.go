//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_offering_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockServiceOfferingLabelUpserter is a mock of ServiceOfferingLabelUpserter interface.
type MockServiceOfferingLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceOfferingLabelUpserterMockRecorder
}

// MockServiceOfferingLabelUpserterMockRecorder is the mock recorder for MockServiceOfferingLabelUpserter.
type MockServiceOfferingLabelUpserterMockRecorder struct {
	mock *MockServiceOfferingLabelUpserter
}

// NewMockServiceOfferingLabelUpserter creates a new mock instance.
func NewMockServiceOfferingLabelUpserter(ctrl *gomock.Controller) *MockServiceOfferingLabelUpserter {
	mock := &MockServiceOfferingLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockServiceOfferingLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceOfferingLabelUpserter) EXPECT() *MockServiceOfferingLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockServiceOfferingLabelUpserter) Upsert(o *models.ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServiceOfferingLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServiceOfferingLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockServiceOfferingLabelFinisher is a mock of ServiceOfferingLabelFinisher interface.
type MockServiceOfferingLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceOfferingLabelFinisherMockRecorder
}

// MockServiceOfferingLabelFinisherMockRecorder is the mock recorder for MockServiceOfferingLabelFinisher.
type MockServiceOfferingLabelFinisherMockRecorder struct {
	mock *MockServiceOfferingLabelFinisher
}

// NewMockServiceOfferingLabelFinisher creates a new mock instance.
func NewMockServiceOfferingLabelFinisher(ctrl *gomock.Controller) *MockServiceOfferingLabelFinisher {
	mock := &MockServiceOfferingLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceOfferingLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceOfferingLabelFinisher) EXPECT() *MockServiceOfferingLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceOfferingLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceOfferingLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceOfferingLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceOfferingLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceOfferingLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceOfferingLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceOfferingLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceOfferingLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceOfferingLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceOfferingLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceOfferingLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceOfferingLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceOfferingLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceOfferingLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceOfferingLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceOfferingLabelFinisher)(nil).One), ctx, exec)
}

// MockServiceOfferingLabelFinder is a mock of ServiceOfferingLabelFinder interface.
type MockServiceOfferingLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceOfferingLabelFinderMockRecorder
}

// MockServiceOfferingLabelFinderMockRecorder is the mock recorder for MockServiceOfferingLabelFinder.
type MockServiceOfferingLabelFinderMockRecorder struct {
	mock *MockServiceOfferingLabelFinder
}

// NewMockServiceOfferingLabelFinder creates a new mock instance.
func NewMockServiceOfferingLabelFinder(ctrl *gomock.Controller) *MockServiceOfferingLabelFinder {
	mock := &MockServiceOfferingLabelFinder{ctrl: ctrl}
	mock.recorder = &MockServiceOfferingLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceOfferingLabelFinder) EXPECT() *MockServiceOfferingLabelFinderMockRecorder {
	return m.recorder
}

// FindServiceOfferingLabel mocks base method.
func (m *MockServiceOfferingLabelFinder) FindServiceOfferingLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceOfferingLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceOfferingLabel", varargs...)
	ret0, _ := ret[0].(*models.ServiceOfferingLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceOfferingLabel indicates an expected call of FindServiceOfferingLabel.
func (mr *MockServiceOfferingLabelFinderMockRecorder) FindServiceOfferingLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceOfferingLabel", reflect.TypeOf((*MockServiceOfferingLabelFinder)(nil).FindServiceOfferingLabel), varargs...)
}

// MockServiceOfferingLabelInserter is a mock of ServiceOfferingLabelInserter interface.
type MockServiceOfferingLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceOfferingLabelInserterMockRecorder
}

// MockServiceOfferingLabelInserterMockRecorder is the mock recorder for MockServiceOfferingLabelInserter.
type MockServiceOfferingLabelInserterMockRecorder struct {
	mock *MockServiceOfferingLabelInserter
}

// NewMockServiceOfferingLabelInserter creates a new mock instance.
func NewMockServiceOfferingLabelInserter(ctrl *gomock.Controller) *MockServiceOfferingLabelInserter {
	mock := &MockServiceOfferingLabelInserter{ctrl: ctrl}
	mock.recorder = &MockServiceOfferingLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceOfferingLabelInserter) EXPECT() *MockServiceOfferingLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceOfferingLabelInserter) Insert(o *models.ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceOfferingLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceOfferingLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceOfferingLabelUpdater is a mock of ServiceOfferingLabelUpdater interface.
type MockServiceOfferingLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceOfferingLabelUpdaterMockRecorder
}

// MockServiceOfferingLabelUpdaterMockRecorder is the mock recorder for MockServiceOfferingLabelUpdater.
type MockServiceOfferingLabelUpdaterMockRecorder struct {
	mock *MockServiceOfferingLabelUpdater
}

// NewMockServiceOfferingLabelUpdater creates a new mock instance.
func NewMockServiceOfferingLabelUpdater(ctrl *gomock.Controller) *MockServiceOfferingLabelUpdater {
	mock := &MockServiceOfferingLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceOfferingLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceOfferingLabelUpdater) EXPECT() *MockServiceOfferingLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceOfferingLabelUpdater) Update(o *models.ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceOfferingLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceOfferingLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceOfferingLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceOfferingLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceOfferingLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceOfferingLabelUpdater) UpdateAllSlice(o models.ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceOfferingLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceOfferingLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceOfferingLabelDeleter is a mock of ServiceOfferingLabelDeleter interface.
type MockServiceOfferingLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceOfferingLabelDeleterMockRecorder
}

// MockServiceOfferingLabelDeleterMockRecorder is the mock recorder for MockServiceOfferingLabelDeleter.
type MockServiceOfferingLabelDeleterMockRecorder struct {
	mock *MockServiceOfferingLabelDeleter
}

// NewMockServiceOfferingLabelDeleter creates a new mock instance.
func NewMockServiceOfferingLabelDeleter(ctrl *gomock.Controller) *MockServiceOfferingLabelDeleter {
	mock := &MockServiceOfferingLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceOfferingLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceOfferingLabelDeleter) EXPECT() *MockServiceOfferingLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceOfferingLabelDeleter) Delete(o *models.ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceOfferingLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceOfferingLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceOfferingLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceOfferingLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceOfferingLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceOfferingLabelDeleter) DeleteAllSlice(o models.ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceOfferingLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceOfferingLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceOfferingLabelReloader is a mock of ServiceOfferingLabelReloader interface.
type MockServiceOfferingLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceOfferingLabelReloaderMockRecorder
}

// MockServiceOfferingLabelReloaderMockRecorder is the mock recorder for MockServiceOfferingLabelReloader.
type MockServiceOfferingLabelReloaderMockRecorder struct {
	mock *MockServiceOfferingLabelReloader
}

// NewMockServiceOfferingLabelReloader creates a new mock instance.
func NewMockServiceOfferingLabelReloader(ctrl *gomock.Controller) *MockServiceOfferingLabelReloader {
	mock := &MockServiceOfferingLabelReloader{ctrl: ctrl}
	mock.recorder = &MockServiceOfferingLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceOfferingLabelReloader) EXPECT() *MockServiceOfferingLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceOfferingLabelReloader) Reload(o *models.ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceOfferingLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceOfferingLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceOfferingLabelReloader) ReloadAll(o *models.ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceOfferingLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceOfferingLabelReloader)(nil).ReloadAll), o, ctx, exec)
}
