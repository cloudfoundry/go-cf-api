// Code generated by MockGen. DO NOT EDIT.
// Source: psql_organizations_private_domains.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockOrganizationsPrivateDomainFinisher is a mock of OrganizationsPrivateDomainFinisher interface.
type MockOrganizationsPrivateDomainFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsPrivateDomainFinisherMockRecorder
}

// MockOrganizationsPrivateDomainFinisherMockRecorder is the mock recorder for MockOrganizationsPrivateDomainFinisher.
type MockOrganizationsPrivateDomainFinisherMockRecorder struct {
	mock *MockOrganizationsPrivateDomainFinisher
}

// NewMockOrganizationsPrivateDomainFinisher creates a new mock instance.
func NewMockOrganizationsPrivateDomainFinisher(ctrl *gomock.Controller) *MockOrganizationsPrivateDomainFinisher {
	mock := &MockOrganizationsPrivateDomainFinisher{ctrl: ctrl}
	mock.recorder = &MockOrganizationsPrivateDomainFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsPrivateDomainFinisher) EXPECT() *MockOrganizationsPrivateDomainFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrganizationsPrivateDomainFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrganizationsPrivateDomainSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrganizationsPrivateDomainSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrganizationsPrivateDomainFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrganizationsPrivateDomainFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrganizationsPrivateDomainFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrganizationsPrivateDomainFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrganizationsPrivateDomainFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrganizationsPrivateDomainFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrganizationsPrivateDomainFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrganizationsPrivateDomainFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrganizationsPrivateDomainFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.OrganizationsPrivateDomain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.OrganizationsPrivateDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrganizationsPrivateDomainFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrganizationsPrivateDomainFinisher)(nil).One), ctx, exec)
}

// MockOrganizationsPrivateDomainFinder is a mock of OrganizationsPrivateDomainFinder interface.
type MockOrganizationsPrivateDomainFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsPrivateDomainFinderMockRecorder
}

// MockOrganizationsPrivateDomainFinderMockRecorder is the mock recorder for MockOrganizationsPrivateDomainFinder.
type MockOrganizationsPrivateDomainFinderMockRecorder struct {
	mock *MockOrganizationsPrivateDomainFinder
}

// NewMockOrganizationsPrivateDomainFinder creates a new mock instance.
func NewMockOrganizationsPrivateDomainFinder(ctrl *gomock.Controller) *MockOrganizationsPrivateDomainFinder {
	mock := &MockOrganizationsPrivateDomainFinder{ctrl: ctrl}
	mock.recorder = &MockOrganizationsPrivateDomainFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsPrivateDomainFinder) EXPECT() *MockOrganizationsPrivateDomainFinderMockRecorder {
	return m.recorder
}

// FindOrganizationsPrivateDomain mocks base method.
func (m *MockOrganizationsPrivateDomainFinder) FindOrganizationsPrivateDomain(ctx context.Context, exec boil.ContextExecutor, organizationsPrivateDomainsPK int, selectCols ...string) (*models.OrganizationsPrivateDomain, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, organizationsPrivateDomainsPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganizationsPrivateDomain", varargs...)
	ret0, _ := ret[0].(*models.OrganizationsPrivateDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationsPrivateDomain indicates an expected call of FindOrganizationsPrivateDomain.
func (mr *MockOrganizationsPrivateDomainFinderMockRecorder) FindOrganizationsPrivateDomain(ctx, exec, organizationsPrivateDomainsPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, organizationsPrivateDomainsPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationsPrivateDomain", reflect.TypeOf((*MockOrganizationsPrivateDomainFinder)(nil).FindOrganizationsPrivateDomain), varargs...)
}

// MockOrganizationsPrivateDomainInserter is a mock of OrganizationsPrivateDomainInserter interface.
type MockOrganizationsPrivateDomainInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsPrivateDomainInserterMockRecorder
}

// MockOrganizationsPrivateDomainInserterMockRecorder is the mock recorder for MockOrganizationsPrivateDomainInserter.
type MockOrganizationsPrivateDomainInserterMockRecorder struct {
	mock *MockOrganizationsPrivateDomainInserter
}

// NewMockOrganizationsPrivateDomainInserter creates a new mock instance.
func NewMockOrganizationsPrivateDomainInserter(ctrl *gomock.Controller) *MockOrganizationsPrivateDomainInserter {
	mock := &MockOrganizationsPrivateDomainInserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsPrivateDomainInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsPrivateDomainInserter) EXPECT() *MockOrganizationsPrivateDomainInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrganizationsPrivateDomainInserter) Insert(o *models.OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationsPrivateDomainInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationsPrivateDomainInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrganizationsPrivateDomainUpdater is a mock of OrganizationsPrivateDomainUpdater interface.
type MockOrganizationsPrivateDomainUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsPrivateDomainUpdaterMockRecorder
}

// MockOrganizationsPrivateDomainUpdaterMockRecorder is the mock recorder for MockOrganizationsPrivateDomainUpdater.
type MockOrganizationsPrivateDomainUpdaterMockRecorder struct {
	mock *MockOrganizationsPrivateDomainUpdater
}

// NewMockOrganizationsPrivateDomainUpdater creates a new mock instance.
func NewMockOrganizationsPrivateDomainUpdater(ctrl *gomock.Controller) *MockOrganizationsPrivateDomainUpdater {
	mock := &MockOrganizationsPrivateDomainUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationsPrivateDomainUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsPrivateDomainUpdater) EXPECT() *MockOrganizationsPrivateDomainUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrganizationsPrivateDomainUpdater) Update(o *models.OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationsPrivateDomainUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationsPrivateDomainUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrganizationsPrivateDomainUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrganizationsPrivateDomainUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrganizationsPrivateDomainUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrganizationsPrivateDomainUpdater) UpdateAllSlice(o models.OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrganizationsPrivateDomainUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrganizationsPrivateDomainUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrganizationsPrivateDomainDeleter is a mock of OrganizationsPrivateDomainDeleter interface.
type MockOrganizationsPrivateDomainDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsPrivateDomainDeleterMockRecorder
}

// MockOrganizationsPrivateDomainDeleterMockRecorder is the mock recorder for MockOrganizationsPrivateDomainDeleter.
type MockOrganizationsPrivateDomainDeleterMockRecorder struct {
	mock *MockOrganizationsPrivateDomainDeleter
}

// NewMockOrganizationsPrivateDomainDeleter creates a new mock instance.
func NewMockOrganizationsPrivateDomainDeleter(ctrl *gomock.Controller) *MockOrganizationsPrivateDomainDeleter {
	mock := &MockOrganizationsPrivateDomainDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationsPrivateDomainDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsPrivateDomainDeleter) EXPECT() *MockOrganizationsPrivateDomainDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationsPrivateDomainDeleter) Delete(o *models.OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationsPrivateDomainDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationsPrivateDomainDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrganizationsPrivateDomainDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrganizationsPrivateDomainDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrganizationsPrivateDomainDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrganizationsPrivateDomainDeleter) DeleteAllSlice(o models.OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrganizationsPrivateDomainDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrganizationsPrivateDomainDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrganizationsPrivateDomainReloader is a mock of OrganizationsPrivateDomainReloader interface.
type MockOrganizationsPrivateDomainReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsPrivateDomainReloaderMockRecorder
}

// MockOrganizationsPrivateDomainReloaderMockRecorder is the mock recorder for MockOrganizationsPrivateDomainReloader.
type MockOrganizationsPrivateDomainReloaderMockRecorder struct {
	mock *MockOrganizationsPrivateDomainReloader
}

// NewMockOrganizationsPrivateDomainReloader creates a new mock instance.
func NewMockOrganizationsPrivateDomainReloader(ctrl *gomock.Controller) *MockOrganizationsPrivateDomainReloader {
	mock := &MockOrganizationsPrivateDomainReloader{ctrl: ctrl}
	mock.recorder = &MockOrganizationsPrivateDomainReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsPrivateDomainReloader) EXPECT() *MockOrganizationsPrivateDomainReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrganizationsPrivateDomainReloader) Reload(o *models.OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrganizationsPrivateDomainReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrganizationsPrivateDomainReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrganizationsPrivateDomainReloader) ReloadAll(o *models.OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrganizationsPrivateDomainReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrganizationsPrivateDomainReloader)(nil).ReloadAll), o, ctx, exec)
}
