//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_organizations_auditors.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
)

// MockOrganizationsAuditorUpserter is a mock of OrganizationsAuditorUpserter interface.
type MockOrganizationsAuditorUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsAuditorUpserterMockRecorder
}

// MockOrganizationsAuditorUpserterMockRecorder is the mock recorder for MockOrganizationsAuditorUpserter.
type MockOrganizationsAuditorUpserterMockRecorder struct {
	mock *MockOrganizationsAuditorUpserter
}

// NewMockOrganizationsAuditorUpserter creates a new mock instance.
func NewMockOrganizationsAuditorUpserter(ctrl *gomock.Controller) *MockOrganizationsAuditorUpserter {
	mock := &MockOrganizationsAuditorUpserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsAuditorUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsAuditorUpserter) EXPECT() *MockOrganizationsAuditorUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockOrganizationsAuditorUpserter) Upsert(o *models.OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOrganizationsAuditorUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOrganizationsAuditorUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockOrganizationsAuditorFinisher is a mock of OrganizationsAuditorFinisher interface.
type MockOrganizationsAuditorFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsAuditorFinisherMockRecorder
}

// MockOrganizationsAuditorFinisherMockRecorder is the mock recorder for MockOrganizationsAuditorFinisher.
type MockOrganizationsAuditorFinisherMockRecorder struct {
	mock *MockOrganizationsAuditorFinisher
}

// NewMockOrganizationsAuditorFinisher creates a new mock instance.
func NewMockOrganizationsAuditorFinisher(ctrl *gomock.Controller) *MockOrganizationsAuditorFinisher {
	mock := &MockOrganizationsAuditorFinisher{ctrl: ctrl}
	mock.recorder = &MockOrganizationsAuditorFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsAuditorFinisher) EXPECT() *MockOrganizationsAuditorFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrganizationsAuditorFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrganizationsAuditorSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrganizationsAuditorSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrganizationsAuditorFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrganizationsAuditorFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrganizationsAuditorFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrganizationsAuditorFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrganizationsAuditorFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrganizationsAuditorFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrganizationsAuditorFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrganizationsAuditorFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrganizationsAuditorFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.OrganizationsAuditor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.OrganizationsAuditor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrganizationsAuditorFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrganizationsAuditorFinisher)(nil).One), ctx, exec)
}

// MockOrganizationsAuditorFinder is a mock of OrganizationsAuditorFinder interface.
type MockOrganizationsAuditorFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsAuditorFinderMockRecorder
}

// MockOrganizationsAuditorFinderMockRecorder is the mock recorder for MockOrganizationsAuditorFinder.
type MockOrganizationsAuditorFinderMockRecorder struct {
	mock *MockOrganizationsAuditorFinder
}

// NewMockOrganizationsAuditorFinder creates a new mock instance.
func NewMockOrganizationsAuditorFinder(ctrl *gomock.Controller) *MockOrganizationsAuditorFinder {
	mock := &MockOrganizationsAuditorFinder{ctrl: ctrl}
	mock.recorder = &MockOrganizationsAuditorFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsAuditorFinder) EXPECT() *MockOrganizationsAuditorFinderMockRecorder {
	return m.recorder
}

// FindOrganizationsAuditor mocks base method.
func (m *MockOrganizationsAuditorFinder) FindOrganizationsAuditor(ctx context.Context, exec boil.ContextExecutor, organizationsAuditorsPK int, selectCols ...string) (*models.OrganizationsAuditor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, organizationsAuditorsPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganizationsAuditor", varargs...)
	ret0, _ := ret[0].(*models.OrganizationsAuditor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationsAuditor indicates an expected call of FindOrganizationsAuditor.
func (mr *MockOrganizationsAuditorFinderMockRecorder) FindOrganizationsAuditor(ctx, exec, organizationsAuditorsPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, organizationsAuditorsPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationsAuditor", reflect.TypeOf((*MockOrganizationsAuditorFinder)(nil).FindOrganizationsAuditor), varargs...)
}

// MockOrganizationsAuditorInserter is a mock of OrganizationsAuditorInserter interface.
type MockOrganizationsAuditorInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsAuditorInserterMockRecorder
}

// MockOrganizationsAuditorInserterMockRecorder is the mock recorder for MockOrganizationsAuditorInserter.
type MockOrganizationsAuditorInserterMockRecorder struct {
	mock *MockOrganizationsAuditorInserter
}

// NewMockOrganizationsAuditorInserter creates a new mock instance.
func NewMockOrganizationsAuditorInserter(ctrl *gomock.Controller) *MockOrganizationsAuditorInserter {
	mock := &MockOrganizationsAuditorInserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsAuditorInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsAuditorInserter) EXPECT() *MockOrganizationsAuditorInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrganizationsAuditorInserter) Insert(o *models.OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationsAuditorInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationsAuditorInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrganizationsAuditorUpdater is a mock of OrganizationsAuditorUpdater interface.
type MockOrganizationsAuditorUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsAuditorUpdaterMockRecorder
}

// MockOrganizationsAuditorUpdaterMockRecorder is the mock recorder for MockOrganizationsAuditorUpdater.
type MockOrganizationsAuditorUpdaterMockRecorder struct {
	mock *MockOrganizationsAuditorUpdater
}

// NewMockOrganizationsAuditorUpdater creates a new mock instance.
func NewMockOrganizationsAuditorUpdater(ctrl *gomock.Controller) *MockOrganizationsAuditorUpdater {
	mock := &MockOrganizationsAuditorUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationsAuditorUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsAuditorUpdater) EXPECT() *MockOrganizationsAuditorUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrganizationsAuditorUpdater) Update(o *models.OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationsAuditorUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationsAuditorUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrganizationsAuditorUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrganizationsAuditorUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrganizationsAuditorUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrganizationsAuditorUpdater) UpdateAllSlice(o models.OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrganizationsAuditorUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrganizationsAuditorUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrganizationsAuditorDeleter is a mock of OrganizationsAuditorDeleter interface.
type MockOrganizationsAuditorDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsAuditorDeleterMockRecorder
}

// MockOrganizationsAuditorDeleterMockRecorder is the mock recorder for MockOrganizationsAuditorDeleter.
type MockOrganizationsAuditorDeleterMockRecorder struct {
	mock *MockOrganizationsAuditorDeleter
}

// NewMockOrganizationsAuditorDeleter creates a new mock instance.
func NewMockOrganizationsAuditorDeleter(ctrl *gomock.Controller) *MockOrganizationsAuditorDeleter {
	mock := &MockOrganizationsAuditorDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsAuditorDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsAuditorDeleter) EXPECT() *MockOrganizationsAuditorDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationsAuditorDeleter) Delete(o *models.OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationsAuditorDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationsAuditorDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrganizationsAuditorDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrganizationsAuditorDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrganizationsAuditorDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrganizationsAuditorDeleter) DeleteAllSlice(o models.OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrganizationsAuditorDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrganizationsAuditorDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrganizationsAuditorReloader is a mock of OrganizationsAuditorReloader interface.
type MockOrganizationsAuditorReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsAuditorReloaderMockRecorder
}

// MockOrganizationsAuditorReloaderMockRecorder is the mock recorder for MockOrganizationsAuditorReloader.
type MockOrganizationsAuditorReloaderMockRecorder struct {
	mock *MockOrganizationsAuditorReloader
}

// NewMockOrganizationsAuditorReloader creates a new mock instance.
func NewMockOrganizationsAuditorReloader(ctrl *gomock.Controller) *MockOrganizationsAuditorReloader {
	mock := &MockOrganizationsAuditorReloader{ctrl: ctrl}
	mock.recorder = &MockOrganizationsAuditorReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsAuditorReloader) EXPECT() *MockOrganizationsAuditorReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrganizationsAuditorReloader) Reload(o *models.OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrganizationsAuditorReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrganizationsAuditorReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrganizationsAuditorReloader) ReloadAll(o *models.OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrganizationsAuditorReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrganizationsAuditorReloader)(nil).ReloadAll), o, ctx, exec)
}
