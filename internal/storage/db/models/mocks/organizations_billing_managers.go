//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_organizations_billing_managers.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockOrganizationsBillingManagerUpserter is a mock of OrganizationsBillingManagerUpserter interface.
type MockOrganizationsBillingManagerUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsBillingManagerUpserterMockRecorder
}

// MockOrganizationsBillingManagerUpserterMockRecorder is the mock recorder for MockOrganizationsBillingManagerUpserter.
type MockOrganizationsBillingManagerUpserterMockRecorder struct {
	mock *MockOrganizationsBillingManagerUpserter
}

// NewMockOrganizationsBillingManagerUpserter creates a new mock instance.
func NewMockOrganizationsBillingManagerUpserter(ctrl *gomock.Controller) *MockOrganizationsBillingManagerUpserter {
	mock := &MockOrganizationsBillingManagerUpserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsBillingManagerUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsBillingManagerUpserter) EXPECT() *MockOrganizationsBillingManagerUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockOrganizationsBillingManagerUpserter) Upsert(o *models.OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOrganizationsBillingManagerUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOrganizationsBillingManagerUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockOrganizationsBillingManagerFinisher is a mock of OrganizationsBillingManagerFinisher interface.
type MockOrganizationsBillingManagerFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsBillingManagerFinisherMockRecorder
}

// MockOrganizationsBillingManagerFinisherMockRecorder is the mock recorder for MockOrganizationsBillingManagerFinisher.
type MockOrganizationsBillingManagerFinisherMockRecorder struct {
	mock *MockOrganizationsBillingManagerFinisher
}

// NewMockOrganizationsBillingManagerFinisher creates a new mock instance.
func NewMockOrganizationsBillingManagerFinisher(ctrl *gomock.Controller) *MockOrganizationsBillingManagerFinisher {
	mock := &MockOrganizationsBillingManagerFinisher{ctrl: ctrl}
	mock.recorder = &MockOrganizationsBillingManagerFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsBillingManagerFinisher) EXPECT() *MockOrganizationsBillingManagerFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrganizationsBillingManagerFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrganizationsBillingManagerSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrganizationsBillingManagerSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrganizationsBillingManagerFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrganizationsBillingManagerFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrganizationsBillingManagerFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrganizationsBillingManagerFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrganizationsBillingManagerFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrganizationsBillingManagerFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrganizationsBillingManagerFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrganizationsBillingManagerFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrganizationsBillingManagerFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.OrganizationsBillingManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.OrganizationsBillingManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrganizationsBillingManagerFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrganizationsBillingManagerFinisher)(nil).One), ctx, exec)
}

// MockOrganizationsBillingManagerFinder is a mock of OrganizationsBillingManagerFinder interface.
type MockOrganizationsBillingManagerFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsBillingManagerFinderMockRecorder
}

// MockOrganizationsBillingManagerFinderMockRecorder is the mock recorder for MockOrganizationsBillingManagerFinder.
type MockOrganizationsBillingManagerFinderMockRecorder struct {
	mock *MockOrganizationsBillingManagerFinder
}

// NewMockOrganizationsBillingManagerFinder creates a new mock instance.
func NewMockOrganizationsBillingManagerFinder(ctrl *gomock.Controller) *MockOrganizationsBillingManagerFinder {
	mock := &MockOrganizationsBillingManagerFinder{ctrl: ctrl}
	mock.recorder = &MockOrganizationsBillingManagerFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsBillingManagerFinder) EXPECT() *MockOrganizationsBillingManagerFinderMockRecorder {
	return m.recorder
}

// FindOrganizationsBillingManager mocks base method.
func (m *MockOrganizationsBillingManagerFinder) FindOrganizationsBillingManager(ctx context.Context, exec boil.ContextExecutor, organizationsBillingManagersPK int, selectCols ...string) (*models.OrganizationsBillingManager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, organizationsBillingManagersPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganizationsBillingManager", varargs...)
	ret0, _ := ret[0].(*models.OrganizationsBillingManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationsBillingManager indicates an expected call of FindOrganizationsBillingManager.
func (mr *MockOrganizationsBillingManagerFinderMockRecorder) FindOrganizationsBillingManager(ctx, exec, organizationsBillingManagersPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, organizationsBillingManagersPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationsBillingManager", reflect.TypeOf((*MockOrganizationsBillingManagerFinder)(nil).FindOrganizationsBillingManager), varargs...)
}

// MockOrganizationsBillingManagerInserter is a mock of OrganizationsBillingManagerInserter interface.
type MockOrganizationsBillingManagerInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsBillingManagerInserterMockRecorder
}

// MockOrganizationsBillingManagerInserterMockRecorder is the mock recorder for MockOrganizationsBillingManagerInserter.
type MockOrganizationsBillingManagerInserterMockRecorder struct {
	mock *MockOrganizationsBillingManagerInserter
}

// NewMockOrganizationsBillingManagerInserter creates a new mock instance.
func NewMockOrganizationsBillingManagerInserter(ctrl *gomock.Controller) *MockOrganizationsBillingManagerInserter {
	mock := &MockOrganizationsBillingManagerInserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsBillingManagerInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsBillingManagerInserter) EXPECT() *MockOrganizationsBillingManagerInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrganizationsBillingManagerInserter) Insert(o *models.OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationsBillingManagerInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationsBillingManagerInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrganizationsBillingManagerUpdater is a mock of OrganizationsBillingManagerUpdater interface.
type MockOrganizationsBillingManagerUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsBillingManagerUpdaterMockRecorder
}

// MockOrganizationsBillingManagerUpdaterMockRecorder is the mock recorder for MockOrganizationsBillingManagerUpdater.
type MockOrganizationsBillingManagerUpdaterMockRecorder struct {
	mock *MockOrganizationsBillingManagerUpdater
}

// NewMockOrganizationsBillingManagerUpdater creates a new mock instance.
func NewMockOrganizationsBillingManagerUpdater(ctrl *gomock.Controller) *MockOrganizationsBillingManagerUpdater {
	mock := &MockOrganizationsBillingManagerUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationsBillingManagerUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsBillingManagerUpdater) EXPECT() *MockOrganizationsBillingManagerUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrganizationsBillingManagerUpdater) Update(o *models.OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationsBillingManagerUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationsBillingManagerUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrganizationsBillingManagerUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrganizationsBillingManagerUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrganizationsBillingManagerUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrganizationsBillingManagerUpdater) UpdateAllSlice(o models.OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrganizationsBillingManagerUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrganizationsBillingManagerUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrganizationsBillingManagerDeleter is a mock of OrganizationsBillingManagerDeleter interface.
type MockOrganizationsBillingManagerDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsBillingManagerDeleterMockRecorder
}

// MockOrganizationsBillingManagerDeleterMockRecorder is the mock recorder for MockOrganizationsBillingManagerDeleter.
type MockOrganizationsBillingManagerDeleterMockRecorder struct {
	mock *MockOrganizationsBillingManagerDeleter
}

// NewMockOrganizationsBillingManagerDeleter creates a new mock instance.
func NewMockOrganizationsBillingManagerDeleter(ctrl *gomock.Controller) *MockOrganizationsBillingManagerDeleter {
	mock := &MockOrganizationsBillingManagerDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsBillingManagerDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsBillingManagerDeleter) EXPECT() *MockOrganizationsBillingManagerDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationsBillingManagerDeleter) Delete(o *models.OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationsBillingManagerDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationsBillingManagerDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrganizationsBillingManagerDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrganizationsBillingManagerDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrganizationsBillingManagerDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrganizationsBillingManagerDeleter) DeleteAllSlice(o models.OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrganizationsBillingManagerDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrganizationsBillingManagerDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrganizationsBillingManagerReloader is a mock of OrganizationsBillingManagerReloader interface.
type MockOrganizationsBillingManagerReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsBillingManagerReloaderMockRecorder
}

// MockOrganizationsBillingManagerReloaderMockRecorder is the mock recorder for MockOrganizationsBillingManagerReloader.
type MockOrganizationsBillingManagerReloaderMockRecorder struct {
	mock *MockOrganizationsBillingManagerReloader
}

// NewMockOrganizationsBillingManagerReloader creates a new mock instance.
func NewMockOrganizationsBillingManagerReloader(ctrl *gomock.Controller) *MockOrganizationsBillingManagerReloader {
	mock := &MockOrganizationsBillingManagerReloader{ctrl: ctrl}
	mock.recorder = &MockOrganizationsBillingManagerReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsBillingManagerReloader) EXPECT() *MockOrganizationsBillingManagerReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrganizationsBillingManagerReloader) Reload(o *models.OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrganizationsBillingManagerReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrganizationsBillingManagerReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrganizationsBillingManagerReloader) ReloadAll(o *models.OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrganizationsBillingManagerReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrganizationsBillingManagerReloader)(nil).ReloadAll), o, ctx, exec)
}
