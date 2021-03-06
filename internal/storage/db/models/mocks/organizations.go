//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_organizations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockOrganizationUpserter is a mock of OrganizationUpserter interface.
type MockOrganizationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationUpserterMockRecorder
}

// MockOrganizationUpserterMockRecorder is the mock recorder for MockOrganizationUpserter.
type MockOrganizationUpserterMockRecorder struct {
	mock *MockOrganizationUpserter
}

// NewMockOrganizationUpserter creates a new mock instance.
func NewMockOrganizationUpserter(ctrl *gomock.Controller) *MockOrganizationUpserter {
	mock := &MockOrganizationUpserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationUpserter) EXPECT() *MockOrganizationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockOrganizationUpserter) Upsert(o *models.Organization, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOrganizationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOrganizationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockOrganizationFinisher is a mock of OrganizationFinisher interface.
type MockOrganizationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationFinisherMockRecorder
}

// MockOrganizationFinisherMockRecorder is the mock recorder for MockOrganizationFinisher.
type MockOrganizationFinisherMockRecorder struct {
	mock *MockOrganizationFinisher
}

// NewMockOrganizationFinisher creates a new mock instance.
func NewMockOrganizationFinisher(ctrl *gomock.Controller) *MockOrganizationFinisher {
	mock := &MockOrganizationFinisher{ctrl: ctrl}
	mock.recorder = &MockOrganizationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationFinisher) EXPECT() *MockOrganizationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrganizationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrganizationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrganizationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrganizationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrganizationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrganizationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrganizationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrganizationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrganizationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrganizationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrganizationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrganizationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrganizationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrganizationFinisher)(nil).One), ctx, exec)
}

// MockOrganizationFinder is a mock of OrganizationFinder interface.
type MockOrganizationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationFinderMockRecorder
}

// MockOrganizationFinderMockRecorder is the mock recorder for MockOrganizationFinder.
type MockOrganizationFinderMockRecorder struct {
	mock *MockOrganizationFinder
}

// NewMockOrganizationFinder creates a new mock instance.
func NewMockOrganizationFinder(ctrl *gomock.Controller) *MockOrganizationFinder {
	mock := &MockOrganizationFinder{ctrl: ctrl}
	mock.recorder = &MockOrganizationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationFinder) EXPECT() *MockOrganizationFinderMockRecorder {
	return m.recorder
}

// FindOrganization mocks base method.
func (m *MockOrganizationFinder) FindOrganization(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Organization, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganization", varargs...)
	ret0, _ := ret[0].(*models.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganization indicates an expected call of FindOrganization.
func (mr *MockOrganizationFinderMockRecorder) FindOrganization(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganization", reflect.TypeOf((*MockOrganizationFinder)(nil).FindOrganization), varargs...)
}

// MockOrganizationInserter is a mock of OrganizationInserter interface.
type MockOrganizationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationInserterMockRecorder
}

// MockOrganizationInserterMockRecorder is the mock recorder for MockOrganizationInserter.
type MockOrganizationInserterMockRecorder struct {
	mock *MockOrganizationInserter
}

// NewMockOrganizationInserter creates a new mock instance.
func NewMockOrganizationInserter(ctrl *gomock.Controller) *MockOrganizationInserter {
	mock := &MockOrganizationInserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationInserter) EXPECT() *MockOrganizationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrganizationInserter) Insert(o *models.Organization, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrganizationUpdater is a mock of OrganizationUpdater interface.
type MockOrganizationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationUpdaterMockRecorder
}

// MockOrganizationUpdaterMockRecorder is the mock recorder for MockOrganizationUpdater.
type MockOrganizationUpdaterMockRecorder struct {
	mock *MockOrganizationUpdater
}

// NewMockOrganizationUpdater creates a new mock instance.
func NewMockOrganizationUpdater(ctrl *gomock.Controller) *MockOrganizationUpdater {
	mock := &MockOrganizationUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationUpdater) EXPECT() *MockOrganizationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrganizationUpdater) Update(o *models.Organization, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrganizationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrganizationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrganizationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrganizationUpdater) UpdateAllSlice(o models.OrganizationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrganizationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrganizationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrganizationDeleter is a mock of OrganizationDeleter interface.
type MockOrganizationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationDeleterMockRecorder
}

// MockOrganizationDeleterMockRecorder is the mock recorder for MockOrganizationDeleter.
type MockOrganizationDeleterMockRecorder struct {
	mock *MockOrganizationDeleter
}

// NewMockOrganizationDeleter creates a new mock instance.
func NewMockOrganizationDeleter(ctrl *gomock.Controller) *MockOrganizationDeleter {
	mock := &MockOrganizationDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationDeleter) EXPECT() *MockOrganizationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationDeleter) Delete(o *models.Organization, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrganizationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrganizationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrganizationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrganizationDeleter) DeleteAllSlice(o models.OrganizationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrganizationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrganizationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrganizationReloader is a mock of OrganizationReloader interface.
type MockOrganizationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationReloaderMockRecorder
}

// MockOrganizationReloaderMockRecorder is the mock recorder for MockOrganizationReloader.
type MockOrganizationReloaderMockRecorder struct {
	mock *MockOrganizationReloader
}

// NewMockOrganizationReloader creates a new mock instance.
func NewMockOrganizationReloader(ctrl *gomock.Controller) *MockOrganizationReloader {
	mock := &MockOrganizationReloader{ctrl: ctrl}
	mock.recorder = &MockOrganizationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationReloader) EXPECT() *MockOrganizationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrganizationReloader) Reload(o *models.Organization, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrganizationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrganizationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrganizationReloader) ReloadAll(o *models.OrganizationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrganizationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrganizationReloader)(nil).ReloadAll), o, ctx, exec)
}
