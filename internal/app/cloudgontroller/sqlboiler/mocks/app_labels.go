// +build unit

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_app_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockAppLabelUpserter is a mock of AppLabelUpserter interface.
type MockAppLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppLabelUpserterMockRecorder
}

// MockAppLabelUpserterMockRecorder is the mock recorder for MockAppLabelUpserter.
type MockAppLabelUpserterMockRecorder struct {
	mock *MockAppLabelUpserter
}

// NewMockAppLabelUpserter creates a new mock instance.
func NewMockAppLabelUpserter(ctrl *gomock.Controller) *MockAppLabelUpserter {
	mock := &MockAppLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockAppLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLabelUpserter) EXPECT() *MockAppLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockAppLabelUpserter) Upsert(o *models.AppLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockAppLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockAppLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockAppLabelFinisher is a mock of AppLabelFinisher interface.
type MockAppLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockAppLabelFinisherMockRecorder
}

// MockAppLabelFinisherMockRecorder is the mock recorder for MockAppLabelFinisher.
type MockAppLabelFinisherMockRecorder struct {
	mock *MockAppLabelFinisher
}

// NewMockAppLabelFinisher creates a new mock instance.
func NewMockAppLabelFinisher(ctrl *gomock.Controller) *MockAppLabelFinisher {
	mock := &MockAppLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockAppLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLabelFinisher) EXPECT() *MockAppLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockAppLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.AppLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.AppLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockAppLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockAppLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockAppLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockAppLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockAppLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockAppLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockAppLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockAppLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockAppLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.AppLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.AppLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockAppLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockAppLabelFinisher)(nil).One), ctx, exec)
}

// MockAppLabelFinder is a mock of AppLabelFinder interface.
type MockAppLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockAppLabelFinderMockRecorder
}

// MockAppLabelFinderMockRecorder is the mock recorder for MockAppLabelFinder.
type MockAppLabelFinderMockRecorder struct {
	mock *MockAppLabelFinder
}

// NewMockAppLabelFinder creates a new mock instance.
func NewMockAppLabelFinder(ctrl *gomock.Controller) *MockAppLabelFinder {
	mock := &MockAppLabelFinder{ctrl: ctrl}
	mock.recorder = &MockAppLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLabelFinder) EXPECT() *MockAppLabelFinderMockRecorder {
	return m.recorder
}

// FindAppLabel mocks base method.
func (m *MockAppLabelFinder) FindAppLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.AppLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAppLabel", varargs...)
	ret0, _ := ret[0].(*models.AppLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAppLabel indicates an expected call of FindAppLabel.
func (mr *MockAppLabelFinderMockRecorder) FindAppLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAppLabel", reflect.TypeOf((*MockAppLabelFinder)(nil).FindAppLabel), varargs...)
}

// MockAppLabelInserter is a mock of AppLabelInserter interface.
type MockAppLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockAppLabelInserterMockRecorder
}

// MockAppLabelInserterMockRecorder is the mock recorder for MockAppLabelInserter.
type MockAppLabelInserterMockRecorder struct {
	mock *MockAppLabelInserter
}

// NewMockAppLabelInserter creates a new mock instance.
func NewMockAppLabelInserter(ctrl *gomock.Controller) *MockAppLabelInserter {
	mock := &MockAppLabelInserter{ctrl: ctrl}
	mock.recorder = &MockAppLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLabelInserter) EXPECT() *MockAppLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockAppLabelInserter) Insert(o *models.AppLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAppLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAppLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockAppLabelUpdater is a mock of AppLabelUpdater interface.
type MockAppLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAppLabelUpdaterMockRecorder
}

// MockAppLabelUpdaterMockRecorder is the mock recorder for MockAppLabelUpdater.
type MockAppLabelUpdaterMockRecorder struct {
	mock *MockAppLabelUpdater
}

// NewMockAppLabelUpdater creates a new mock instance.
func NewMockAppLabelUpdater(ctrl *gomock.Controller) *MockAppLabelUpdater {
	mock := &MockAppLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockAppLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLabelUpdater) EXPECT() *MockAppLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockAppLabelUpdater) Update(o *models.AppLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAppLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockAppLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockAppLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockAppLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockAppLabelUpdater) UpdateAllSlice(o models.AppLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockAppLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockAppLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockAppLabelDeleter is a mock of AppLabelDeleter interface.
type MockAppLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockAppLabelDeleterMockRecorder
}

// MockAppLabelDeleterMockRecorder is the mock recorder for MockAppLabelDeleter.
type MockAppLabelDeleterMockRecorder struct {
	mock *MockAppLabelDeleter
}

// NewMockAppLabelDeleter creates a new mock instance.
func NewMockAppLabelDeleter(ctrl *gomock.Controller) *MockAppLabelDeleter {
	mock := &MockAppLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockAppLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLabelDeleter) EXPECT() *MockAppLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAppLabelDeleter) Delete(o *models.AppLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAppLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockAppLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockAppLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockAppLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockAppLabelDeleter) DeleteAllSlice(o models.AppLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockAppLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockAppLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockAppLabelReloader is a mock of AppLabelReloader interface.
type MockAppLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockAppLabelReloaderMockRecorder
}

// MockAppLabelReloaderMockRecorder is the mock recorder for MockAppLabelReloader.
type MockAppLabelReloaderMockRecorder struct {
	mock *MockAppLabelReloader
}

// NewMockAppLabelReloader creates a new mock instance.
func NewMockAppLabelReloader(ctrl *gomock.Controller) *MockAppLabelReloader {
	mock := &MockAppLabelReloader{ctrl: ctrl}
	mock.recorder = &MockAppLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppLabelReloader) EXPECT() *MockAppLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockAppLabelReloader) Reload(o *models.AppLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockAppLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockAppLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockAppLabelReloader) ReloadAll(o *models.AppLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockAppLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockAppLabelReloader)(nil).ReloadAll), o, ctx, exec)
}
