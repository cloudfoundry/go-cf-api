// Code generated by MockGen. DO NOT EDIT.
// Source: psql_domains.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockDomainFinisher is a mock of DomainFinisher interface.
type MockDomainFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockDomainFinisherMockRecorder
}

// MockDomainFinisherMockRecorder is the mock recorder for MockDomainFinisher.
type MockDomainFinisherMockRecorder struct {
	mock *MockDomainFinisher
}

// NewMockDomainFinisher creates a new mock instance.
func NewMockDomainFinisher(ctrl *gomock.Controller) *MockDomainFinisher {
	mock := &MockDomainFinisher{ctrl: ctrl}
	mock.recorder = &MockDomainFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainFinisher) EXPECT() *MockDomainFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockDomainFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.DomainSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.DomainSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockDomainFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockDomainFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockDomainFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDomainFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDomainFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockDomainFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDomainFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDomainFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockDomainFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockDomainFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockDomainFinisher)(nil).One), ctx, exec)
}

// MockDomainFinder is a mock of DomainFinder interface.
type MockDomainFinder struct {
	ctrl     *gomock.Controller
	recorder *MockDomainFinderMockRecorder
}

// MockDomainFinderMockRecorder is the mock recorder for MockDomainFinder.
type MockDomainFinderMockRecorder struct {
	mock *MockDomainFinder
}

// NewMockDomainFinder creates a new mock instance.
func NewMockDomainFinder(ctrl *gomock.Controller) *MockDomainFinder {
	mock := &MockDomainFinder{ctrl: ctrl}
	mock.recorder = &MockDomainFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainFinder) EXPECT() *MockDomainFinderMockRecorder {
	return m.recorder
}

// FindDomain mocks base method.
func (m *MockDomainFinder) FindDomain(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Domain, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindDomain", varargs...)
	ret0, _ := ret[0].(*models.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDomain indicates an expected call of FindDomain.
func (mr *MockDomainFinderMockRecorder) FindDomain(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDomain", reflect.TypeOf((*MockDomainFinder)(nil).FindDomain), varargs...)
}

// MockDomainInserter is a mock of DomainInserter interface.
type MockDomainInserter struct {
	ctrl     *gomock.Controller
	recorder *MockDomainInserterMockRecorder
}

// MockDomainInserterMockRecorder is the mock recorder for MockDomainInserter.
type MockDomainInserterMockRecorder struct {
	mock *MockDomainInserter
}

// NewMockDomainInserter creates a new mock instance.
func NewMockDomainInserter(ctrl *gomock.Controller) *MockDomainInserter {
	mock := &MockDomainInserter{ctrl: ctrl}
	mock.recorder = &MockDomainInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainInserter) EXPECT() *MockDomainInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockDomainInserter) Insert(o *models.Domain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockDomainInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDomainInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockDomainUpdater is a mock of DomainUpdater interface.
type MockDomainUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDomainUpdaterMockRecorder
}

// MockDomainUpdaterMockRecorder is the mock recorder for MockDomainUpdater.
type MockDomainUpdaterMockRecorder struct {
	mock *MockDomainUpdater
}

// NewMockDomainUpdater creates a new mock instance.
func NewMockDomainUpdater(ctrl *gomock.Controller) *MockDomainUpdater {
	mock := &MockDomainUpdater{ctrl: ctrl}
	mock.recorder = &MockDomainUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainUpdater) EXPECT() *MockDomainUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockDomainUpdater) Update(o *models.Domain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDomainUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDomainUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockDomainUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockDomainUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockDomainUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockDomainUpdater) UpdateAllSlice(o models.DomainSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockDomainUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockDomainUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockDomainDeleter is a mock of DomainDeleter interface.
type MockDomainDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDomainDeleterMockRecorder
}

// MockDomainDeleterMockRecorder is the mock recorder for MockDomainDeleter.
type MockDomainDeleterMockRecorder struct {
	mock *MockDomainDeleter
}

// NewMockDomainDeleter creates a new mock instance.
func NewMockDomainDeleter(ctrl *gomock.Controller) *MockDomainDeleter {
	mock := &MockDomainDeleter{ctrl: ctrl}
	mock.recorder = &MockDomainDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainDeleter) EXPECT() *MockDomainDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDomainDeleter) Delete(o *models.Domain, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDomainDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDomainDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockDomainDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockDomainDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockDomainDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockDomainDeleter) DeleteAllSlice(o models.DomainSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockDomainDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockDomainDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockDomainReloader is a mock of DomainReloader interface.
type MockDomainReloader struct {
	ctrl     *gomock.Controller
	recorder *MockDomainReloaderMockRecorder
}

// MockDomainReloaderMockRecorder is the mock recorder for MockDomainReloader.
type MockDomainReloaderMockRecorder struct {
	mock *MockDomainReloader
}

// NewMockDomainReloader creates a new mock instance.
func NewMockDomainReloader(ctrl *gomock.Controller) *MockDomainReloader {
	mock := &MockDomainReloader{ctrl: ctrl}
	mock.recorder = &MockDomainReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainReloader) EXPECT() *MockDomainReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockDomainReloader) Reload(o *models.Domain, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockDomainReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockDomainReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockDomainReloader) ReloadAll(o *models.DomainSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockDomainReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockDomainReloader)(nil).ReloadAll), o, ctx, exec)
}
