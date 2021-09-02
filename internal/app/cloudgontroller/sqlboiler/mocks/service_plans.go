// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_plans.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockServicePlanFinisher is a mock of ServicePlanFinisher interface.
type MockServicePlanFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanFinisherMockRecorder
}

// MockServicePlanFinisherMockRecorder is the mock recorder for MockServicePlanFinisher.
type MockServicePlanFinisherMockRecorder struct {
	mock *MockServicePlanFinisher
}

// NewMockServicePlanFinisher creates a new mock instance.
func NewMockServicePlanFinisher(ctrl *gomock.Controller) *MockServicePlanFinisher {
	mock := &MockServicePlanFinisher{ctrl: ctrl}
	mock.recorder = &MockServicePlanFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanFinisher) EXPECT() *MockServicePlanFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServicePlanFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServicePlanSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServicePlanSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServicePlanFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServicePlanFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServicePlanFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServicePlanFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServicePlanFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServicePlanFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServicePlanFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServicePlanFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServicePlanFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServicePlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServicePlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServicePlanFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServicePlanFinisher)(nil).One), ctx, exec)
}

// MockServicePlanFinder is a mock of ServicePlanFinder interface.
type MockServicePlanFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanFinderMockRecorder
}

// MockServicePlanFinderMockRecorder is the mock recorder for MockServicePlanFinder.
type MockServicePlanFinderMockRecorder struct {
	mock *MockServicePlanFinder
}

// NewMockServicePlanFinder creates a new mock instance.
func NewMockServicePlanFinder(ctrl *gomock.Controller) *MockServicePlanFinder {
	mock := &MockServicePlanFinder{ctrl: ctrl}
	mock.recorder = &MockServicePlanFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanFinder) EXPECT() *MockServicePlanFinderMockRecorder {
	return m.recorder
}

// FindServicePlan mocks base method.
func (m *MockServicePlanFinder) FindServicePlan(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServicePlan, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServicePlan", varargs...)
	ret0, _ := ret[0].(*models.ServicePlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServicePlan indicates an expected call of FindServicePlan.
func (mr *MockServicePlanFinderMockRecorder) FindServicePlan(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServicePlan", reflect.TypeOf((*MockServicePlanFinder)(nil).FindServicePlan), varargs...)
}

// MockServicePlanInserter is a mock of ServicePlanInserter interface.
type MockServicePlanInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanInserterMockRecorder
}

// MockServicePlanInserterMockRecorder is the mock recorder for MockServicePlanInserter.
type MockServicePlanInserterMockRecorder struct {
	mock *MockServicePlanInserter
}

// NewMockServicePlanInserter creates a new mock instance.
func NewMockServicePlanInserter(ctrl *gomock.Controller) *MockServicePlanInserter {
	mock := &MockServicePlanInserter{ctrl: ctrl}
	mock.recorder = &MockServicePlanInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanInserter) EXPECT() *MockServicePlanInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServicePlanInserter) Insert(o *models.ServicePlan, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServicePlanInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServicePlanInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServicePlanUpdater is a mock of ServicePlanUpdater interface.
type MockServicePlanUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanUpdaterMockRecorder
}

// MockServicePlanUpdaterMockRecorder is the mock recorder for MockServicePlanUpdater.
type MockServicePlanUpdaterMockRecorder struct {
	mock *MockServicePlanUpdater
}

// NewMockServicePlanUpdater creates a new mock instance.
func NewMockServicePlanUpdater(ctrl *gomock.Controller) *MockServicePlanUpdater {
	mock := &MockServicePlanUpdater{ctrl: ctrl}
	mock.recorder = &MockServicePlanUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanUpdater) EXPECT() *MockServicePlanUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServicePlanUpdater) Update(o *models.ServicePlan, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServicePlanUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServicePlanUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServicePlanUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServicePlanUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServicePlanUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServicePlanUpdater) UpdateAllSlice(o models.ServicePlanSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServicePlanUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServicePlanUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServicePlanDeleter is a mock of ServicePlanDeleter interface.
type MockServicePlanDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanDeleterMockRecorder
}

// MockServicePlanDeleterMockRecorder is the mock recorder for MockServicePlanDeleter.
type MockServicePlanDeleterMockRecorder struct {
	mock *MockServicePlanDeleter
}

// NewMockServicePlanDeleter creates a new mock instance.
func NewMockServicePlanDeleter(ctrl *gomock.Controller) *MockServicePlanDeleter {
	mock := &MockServicePlanDeleter{ctrl: ctrl}
	mock.recorder = &MockServicePlanDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanDeleter) EXPECT() *MockServicePlanDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServicePlanDeleter) Delete(o *models.ServicePlan, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServicePlanDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServicePlanDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServicePlanDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServicePlanDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServicePlanDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServicePlanDeleter) DeleteAllSlice(o models.ServicePlanSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServicePlanDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServicePlanDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServicePlanReloader is a mock of ServicePlanReloader interface.
type MockServicePlanReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServicePlanReloaderMockRecorder
}

// MockServicePlanReloaderMockRecorder is the mock recorder for MockServicePlanReloader.
type MockServicePlanReloaderMockRecorder struct {
	mock *MockServicePlanReloader
}

// NewMockServicePlanReloader creates a new mock instance.
func NewMockServicePlanReloader(ctrl *gomock.Controller) *MockServicePlanReloader {
	mock := &MockServicePlanReloader{ctrl: ctrl}
	mock.recorder = &MockServicePlanReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePlanReloader) EXPECT() *MockServicePlanReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServicePlanReloader) Reload(o *models.ServicePlan, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServicePlanReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServicePlanReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServicePlanReloader) ReloadAll(o *models.ServicePlanSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServicePlanReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServicePlanReloader)(nil).ReloadAll), o, ctx, exec)
}
