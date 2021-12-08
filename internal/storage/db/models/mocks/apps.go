//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_apps.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
)

// MockAppUpserter is a mock of AppUpserter interface.
type MockAppUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppUpserterMockRecorder
}

// MockAppUpserterMockRecorder is the mock recorder for MockAppUpserter.
type MockAppUpserterMockRecorder struct {
	mock *MockAppUpserter
}

// NewMockAppUpserter creates a new mock instance.
func NewMockAppUpserter(ctrl *gomock.Controller) *MockAppUpserter {
	mock := &MockAppUpserter{ctrl: ctrl}
	mock.recorder = &MockAppUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUpserter) EXPECT() *MockAppUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockAppUpserter) Upsert(o *models.App, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockAppUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockAppUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockAppFinisher is a mock of AppFinisher interface.
type MockAppFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockAppFinisherMockRecorder
}

// MockAppFinisherMockRecorder is the mock recorder for MockAppFinisher.
type MockAppFinisherMockRecorder struct {
	mock *MockAppFinisher
}

// NewMockAppFinisher creates a new mock instance.
func NewMockAppFinisher(ctrl *gomock.Controller) *MockAppFinisher {
	mock := &MockAppFinisher{ctrl: ctrl}
	mock.recorder = &MockAppFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppFinisher) EXPECT() *MockAppFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockAppFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.AppSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.AppSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockAppFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockAppFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockAppFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockAppFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockAppFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockAppFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockAppFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockAppFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockAppFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockAppFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockAppFinisher)(nil).One), ctx, exec)
}

// MockAppFinder is a mock of AppFinder interface.
type MockAppFinder struct {
	ctrl     *gomock.Controller
	recorder *MockAppFinderMockRecorder
}

// MockAppFinderMockRecorder is the mock recorder for MockAppFinder.
type MockAppFinderMockRecorder struct {
	mock *MockAppFinder
}

// NewMockAppFinder creates a new mock instance.
func NewMockAppFinder(ctrl *gomock.Controller) *MockAppFinder {
	mock := &MockAppFinder{ctrl: ctrl}
	mock.recorder = &MockAppFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppFinder) EXPECT() *MockAppFinderMockRecorder {
	return m.recorder
}

// FindApp mocks base method.
func (m *MockAppFinder) FindApp(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.App, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindApp", varargs...)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindApp indicates an expected call of FindApp.
func (mr *MockAppFinderMockRecorder) FindApp(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindApp", reflect.TypeOf((*MockAppFinder)(nil).FindApp), varargs...)
}

// MockAppInserter is a mock of AppInserter interface.
type MockAppInserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppInserterMockRecorder
}

// MockAppInserterMockRecorder is the mock recorder for MockAppInserter.
type MockAppInserterMockRecorder struct {
	mock *MockAppInserter
}

// NewMockAppInserter creates a new mock instance.
func NewMockAppInserter(ctrl *gomock.Controller) *MockAppInserter {
	mock := &MockAppInserter{ctrl: ctrl}
	mock.recorder = &MockAppInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppInserter) EXPECT() *MockAppInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockAppInserter) Insert(o *models.App, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAppInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAppInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockAppUpdater is a mock of AppUpdater interface.
type MockAppUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAppUpdaterMockRecorder
}

// MockAppUpdaterMockRecorder is the mock recorder for MockAppUpdater.
type MockAppUpdaterMockRecorder struct {
	mock *MockAppUpdater
}

// NewMockAppUpdater creates a new mock instance.
func NewMockAppUpdater(ctrl *gomock.Controller) *MockAppUpdater {
	mock := &MockAppUpdater{ctrl: ctrl}
	mock.recorder = &MockAppUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppUpdater) EXPECT() *MockAppUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockAppUpdater) Update(o *models.App, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAppUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockAppUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockAppUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockAppUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockAppUpdater) UpdateAllSlice(o models.AppSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockAppUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockAppUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockAppDeleter is a mock of AppDeleter interface.
type MockAppDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockAppDeleterMockRecorder
}

// MockAppDeleterMockRecorder is the mock recorder for MockAppDeleter.
type MockAppDeleterMockRecorder struct {
	mock *MockAppDeleter
}

// NewMockAppDeleter creates a new mock instance.
func NewMockAppDeleter(ctrl *gomock.Controller) *MockAppDeleter {
	mock := &MockAppDeleter{ctrl: ctrl}
	mock.recorder = &MockAppDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppDeleter) EXPECT() *MockAppDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAppDeleter) Delete(o *models.App, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAppDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockAppDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockAppDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockAppDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockAppDeleter) DeleteAllSlice(o models.AppSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockAppDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockAppDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockAppReloader is a mock of AppReloader interface.
type MockAppReloader struct {
	ctrl     *gomock.Controller
	recorder *MockAppReloaderMockRecorder
}

// MockAppReloaderMockRecorder is the mock recorder for MockAppReloader.
type MockAppReloaderMockRecorder struct {
	mock *MockAppReloader
}

// NewMockAppReloader creates a new mock instance.
func NewMockAppReloader(ctrl *gomock.Controller) *MockAppReloader {
	mock := &MockAppReloader{ctrl: ctrl}
	mock.recorder = &MockAppReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppReloader) EXPECT() *MockAppReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockAppReloader) Reload(o *models.App, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockAppReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockAppReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockAppReloader) ReloadAll(o *models.AppSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockAppReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockAppReloader)(nil).ReloadAll), o, ctx, exec)
}
