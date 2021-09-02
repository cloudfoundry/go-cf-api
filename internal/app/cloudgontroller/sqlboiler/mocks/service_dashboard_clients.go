// Code generated by MockGen. DO NOT EDIT.
// Source: psql_service_dashboard_clients.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockServiceDashboardClientFinisher is a mock of ServiceDashboardClientFinisher interface.
type MockServiceDashboardClientFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDashboardClientFinisherMockRecorder
}

// MockServiceDashboardClientFinisherMockRecorder is the mock recorder for MockServiceDashboardClientFinisher.
type MockServiceDashboardClientFinisherMockRecorder struct {
	mock *MockServiceDashboardClientFinisher
}

// NewMockServiceDashboardClientFinisher creates a new mock instance.
func NewMockServiceDashboardClientFinisher(ctrl *gomock.Controller) *MockServiceDashboardClientFinisher {
	mock := &MockServiceDashboardClientFinisher{ctrl: ctrl}
	mock.recorder = &MockServiceDashboardClientFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDashboardClientFinisher) EXPECT() *MockServiceDashboardClientFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockServiceDashboardClientFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.ServiceDashboardClientSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.ServiceDashboardClientSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockServiceDashboardClientFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceDashboardClientFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockServiceDashboardClientFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockServiceDashboardClientFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockServiceDashboardClientFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockServiceDashboardClientFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockServiceDashboardClientFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockServiceDashboardClientFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockServiceDashboardClientFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.ServiceDashboardClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.ServiceDashboardClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockServiceDashboardClientFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceDashboardClientFinisher)(nil).One), ctx, exec)
}

// MockServiceDashboardClientFinder is a mock of ServiceDashboardClientFinder interface.
type MockServiceDashboardClientFinder struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDashboardClientFinderMockRecorder
}

// MockServiceDashboardClientFinderMockRecorder is the mock recorder for MockServiceDashboardClientFinder.
type MockServiceDashboardClientFinderMockRecorder struct {
	mock *MockServiceDashboardClientFinder
}

// NewMockServiceDashboardClientFinder creates a new mock instance.
func NewMockServiceDashboardClientFinder(ctrl *gomock.Controller) *MockServiceDashboardClientFinder {
	mock := &MockServiceDashboardClientFinder{ctrl: ctrl}
	mock.recorder = &MockServiceDashboardClientFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDashboardClientFinder) EXPECT() *MockServiceDashboardClientFinderMockRecorder {
	return m.recorder
}

// FindServiceDashboardClient mocks base method.
func (m *MockServiceDashboardClientFinder) FindServiceDashboardClient(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.ServiceDashboardClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindServiceDashboardClient", varargs...)
	ret0, _ := ret[0].(*models.ServiceDashboardClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceDashboardClient indicates an expected call of FindServiceDashboardClient.
func (mr *MockServiceDashboardClientFinderMockRecorder) FindServiceDashboardClient(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceDashboardClient", reflect.TypeOf((*MockServiceDashboardClientFinder)(nil).FindServiceDashboardClient), varargs...)
}

// MockServiceDashboardClientInserter is a mock of ServiceDashboardClientInserter interface.
type MockServiceDashboardClientInserter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDashboardClientInserterMockRecorder
}

// MockServiceDashboardClientInserterMockRecorder is the mock recorder for MockServiceDashboardClientInserter.
type MockServiceDashboardClientInserterMockRecorder struct {
	mock *MockServiceDashboardClientInserter
}

// NewMockServiceDashboardClientInserter creates a new mock instance.
func NewMockServiceDashboardClientInserter(ctrl *gomock.Controller) *MockServiceDashboardClientInserter {
	mock := &MockServiceDashboardClientInserter{ctrl: ctrl}
	mock.recorder = &MockServiceDashboardClientInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDashboardClientInserter) EXPECT() *MockServiceDashboardClientInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockServiceDashboardClientInserter) Insert(o *models.ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceDashboardClientInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceDashboardClientInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockServiceDashboardClientUpdater is a mock of ServiceDashboardClientUpdater interface.
type MockServiceDashboardClientUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDashboardClientUpdaterMockRecorder
}

// MockServiceDashboardClientUpdaterMockRecorder is the mock recorder for MockServiceDashboardClientUpdater.
type MockServiceDashboardClientUpdaterMockRecorder struct {
	mock *MockServiceDashboardClientUpdater
}

// NewMockServiceDashboardClientUpdater creates a new mock instance.
func NewMockServiceDashboardClientUpdater(ctrl *gomock.Controller) *MockServiceDashboardClientUpdater {
	mock := &MockServiceDashboardClientUpdater{ctrl: ctrl}
	mock.recorder = &MockServiceDashboardClientUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDashboardClientUpdater) EXPECT() *MockServiceDashboardClientUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockServiceDashboardClientUpdater) Update(o *models.ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceDashboardClientUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceDashboardClientUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockServiceDashboardClientUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockServiceDashboardClientUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockServiceDashboardClientUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockServiceDashboardClientUpdater) UpdateAllSlice(o models.ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockServiceDashboardClientUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockServiceDashboardClientUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockServiceDashboardClientDeleter is a mock of ServiceDashboardClientDeleter interface.
type MockServiceDashboardClientDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDashboardClientDeleterMockRecorder
}

// MockServiceDashboardClientDeleterMockRecorder is the mock recorder for MockServiceDashboardClientDeleter.
type MockServiceDashboardClientDeleterMockRecorder struct {
	mock *MockServiceDashboardClientDeleter
}

// NewMockServiceDashboardClientDeleter creates a new mock instance.
func NewMockServiceDashboardClientDeleter(ctrl *gomock.Controller) *MockServiceDashboardClientDeleter {
	mock := &MockServiceDashboardClientDeleter{ctrl: ctrl}
	mock.recorder = &MockServiceDashboardClientDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDashboardClientDeleter) EXPECT() *MockServiceDashboardClientDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServiceDashboardClientDeleter) Delete(o *models.ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceDashboardClientDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceDashboardClientDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockServiceDashboardClientDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockServiceDashboardClientDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockServiceDashboardClientDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockServiceDashboardClientDeleter) DeleteAllSlice(o models.ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockServiceDashboardClientDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockServiceDashboardClientDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockServiceDashboardClientReloader is a mock of ServiceDashboardClientReloader interface.
type MockServiceDashboardClientReloader struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDashboardClientReloaderMockRecorder
}

// MockServiceDashboardClientReloaderMockRecorder is the mock recorder for MockServiceDashboardClientReloader.
type MockServiceDashboardClientReloaderMockRecorder struct {
	mock *MockServiceDashboardClientReloader
}

// NewMockServiceDashboardClientReloader creates a new mock instance.
func NewMockServiceDashboardClientReloader(ctrl *gomock.Controller) *MockServiceDashboardClientReloader {
	mock := &MockServiceDashboardClientReloader{ctrl: ctrl}
	mock.recorder = &MockServiceDashboardClientReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDashboardClientReloader) EXPECT() *MockServiceDashboardClientReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockServiceDashboardClientReloader) Reload(o *models.ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockServiceDashboardClientReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockServiceDashboardClientReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockServiceDashboardClientReloader) ReloadAll(o *models.ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockServiceDashboardClientReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockServiceDashboardClientReloader)(nil).ReloadAll), o, ctx, exec)
}
