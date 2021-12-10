//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_organizations_managers.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockOrganizationsManagerUpserter is a mock of OrganizationsManagerUpserter interface.
type MockOrganizationsManagerUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsManagerUpserterMockRecorder
}

// MockOrganizationsManagerUpserterMockRecorder is the mock recorder for MockOrganizationsManagerUpserter.
type MockOrganizationsManagerUpserterMockRecorder struct {
	mock *MockOrganizationsManagerUpserter
}

// NewMockOrganizationsManagerUpserter creates a new mock instance.
func NewMockOrganizationsManagerUpserter(ctrl *gomock.Controller) *MockOrganizationsManagerUpserter {
	mock := &MockOrganizationsManagerUpserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsManagerUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsManagerUpserter) EXPECT() *MockOrganizationsManagerUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockOrganizationsManagerUpserter) Upsert(o *models.OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOrganizationsManagerUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOrganizationsManagerUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockOrganizationsManagerFinisher is a mock of OrganizationsManagerFinisher interface.
type MockOrganizationsManagerFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsManagerFinisherMockRecorder
}

// MockOrganizationsManagerFinisherMockRecorder is the mock recorder for MockOrganizationsManagerFinisher.
type MockOrganizationsManagerFinisherMockRecorder struct {
	mock *MockOrganizationsManagerFinisher
}

// NewMockOrganizationsManagerFinisher creates a new mock instance.
func NewMockOrganizationsManagerFinisher(ctrl *gomock.Controller) *MockOrganizationsManagerFinisher {
	mock := &MockOrganizationsManagerFinisher{ctrl: ctrl}
	mock.recorder = &MockOrganizationsManagerFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsManagerFinisher) EXPECT() *MockOrganizationsManagerFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrganizationsManagerFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrganizationsManagerSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrganizationsManagerSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrganizationsManagerFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrganizationsManagerFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrganizationsManagerFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrganizationsManagerFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrganizationsManagerFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrganizationsManagerFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrganizationsManagerFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrganizationsManagerFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrganizationsManagerFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.OrganizationsManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.OrganizationsManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrganizationsManagerFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrganizationsManagerFinisher)(nil).One), ctx, exec)
}

// MockOrganizationsManagerFinder is a mock of OrganizationsManagerFinder interface.
type MockOrganizationsManagerFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsManagerFinderMockRecorder
}

// MockOrganizationsManagerFinderMockRecorder is the mock recorder for MockOrganizationsManagerFinder.
type MockOrganizationsManagerFinderMockRecorder struct {
	mock *MockOrganizationsManagerFinder
}

// NewMockOrganizationsManagerFinder creates a new mock instance.
func NewMockOrganizationsManagerFinder(ctrl *gomock.Controller) *MockOrganizationsManagerFinder {
	mock := &MockOrganizationsManagerFinder{ctrl: ctrl}
	mock.recorder = &MockOrganizationsManagerFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsManagerFinder) EXPECT() *MockOrganizationsManagerFinderMockRecorder {
	return m.recorder
}

// FindOrganizationsManager mocks base method.
func (m *MockOrganizationsManagerFinder) FindOrganizationsManager(ctx context.Context, exec boil.ContextExecutor, organizationsManagersPK int, selectCols ...string) (*models.OrganizationsManager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, organizationsManagersPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganizationsManager", varargs...)
	ret0, _ := ret[0].(*models.OrganizationsManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationsManager indicates an expected call of FindOrganizationsManager.
func (mr *MockOrganizationsManagerFinderMockRecorder) FindOrganizationsManager(ctx, exec, organizationsManagersPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, organizationsManagersPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationsManager", reflect.TypeOf((*MockOrganizationsManagerFinder)(nil).FindOrganizationsManager), varargs...)
}

// MockOrganizationsManagerInserter is a mock of OrganizationsManagerInserter interface.
type MockOrganizationsManagerInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsManagerInserterMockRecorder
}

// MockOrganizationsManagerInserterMockRecorder is the mock recorder for MockOrganizationsManagerInserter.
type MockOrganizationsManagerInserterMockRecorder struct {
	mock *MockOrganizationsManagerInserter
}

// NewMockOrganizationsManagerInserter creates a new mock instance.
func NewMockOrganizationsManagerInserter(ctrl *gomock.Controller) *MockOrganizationsManagerInserter {
	mock := &MockOrganizationsManagerInserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsManagerInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsManagerInserter) EXPECT() *MockOrganizationsManagerInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrganizationsManagerInserter) Insert(o *models.OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationsManagerInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationsManagerInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrganizationsManagerUpdater is a mock of OrganizationsManagerUpdater interface.
type MockOrganizationsManagerUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsManagerUpdaterMockRecorder
}

// MockOrganizationsManagerUpdaterMockRecorder is the mock recorder for MockOrganizationsManagerUpdater.
type MockOrganizationsManagerUpdaterMockRecorder struct {
	mock *MockOrganizationsManagerUpdater
}

// NewMockOrganizationsManagerUpdater creates a new mock instance.
func NewMockOrganizationsManagerUpdater(ctrl *gomock.Controller) *MockOrganizationsManagerUpdater {
	mock := &MockOrganizationsManagerUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationsManagerUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsManagerUpdater) EXPECT() *MockOrganizationsManagerUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrganizationsManagerUpdater) Update(o *models.OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationsManagerUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationsManagerUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrganizationsManagerUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrganizationsManagerUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrganizationsManagerUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrganizationsManagerUpdater) UpdateAllSlice(o models.OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrganizationsManagerUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrganizationsManagerUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrganizationsManagerDeleter is a mock of OrganizationsManagerDeleter interface.
type MockOrganizationsManagerDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsManagerDeleterMockRecorder
}

// MockOrganizationsManagerDeleterMockRecorder is the mock recorder for MockOrganizationsManagerDeleter.
type MockOrganizationsManagerDeleterMockRecorder struct {
	mock *MockOrganizationsManagerDeleter
}

// NewMockOrganizationsManagerDeleter creates a new mock instance.
func NewMockOrganizationsManagerDeleter(ctrl *gomock.Controller) *MockOrganizationsManagerDeleter {
	mock := &MockOrganizationsManagerDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsManagerDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsManagerDeleter) EXPECT() *MockOrganizationsManagerDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationsManagerDeleter) Delete(o *models.OrganizationsManager, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationsManagerDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationsManagerDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrganizationsManagerDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrganizationsManagerDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrganizationsManagerDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrganizationsManagerDeleter) DeleteAllSlice(o models.OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrganizationsManagerDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrganizationsManagerDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrganizationsManagerReloader is a mock of OrganizationsManagerReloader interface.
type MockOrganizationsManagerReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsManagerReloaderMockRecorder
}

// MockOrganizationsManagerReloaderMockRecorder is the mock recorder for MockOrganizationsManagerReloader.
type MockOrganizationsManagerReloaderMockRecorder struct {
	mock *MockOrganizationsManagerReloader
}

// NewMockOrganizationsManagerReloader creates a new mock instance.
func NewMockOrganizationsManagerReloader(ctrl *gomock.Controller) *MockOrganizationsManagerReloader {
	mock := &MockOrganizationsManagerReloader{ctrl: ctrl}
	mock.recorder = &MockOrganizationsManagerReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsManagerReloader) EXPECT() *MockOrganizationsManagerReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrganizationsManagerReloader) Reload(o *models.OrganizationsManager, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrganizationsManagerReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrganizationsManagerReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrganizationsManagerReloader) ReloadAll(o *models.OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrganizationsManagerReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrganizationsManagerReloader)(nil).ReloadAll), o, ctx, exec)
}
