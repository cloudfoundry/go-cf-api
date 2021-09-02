// Code generated by MockGen. DO NOT EDIT.
// Source: psql_deployment_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockDeploymentLabelFinisher is a mock of DeploymentLabelFinisher interface.
type MockDeploymentLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentLabelFinisherMockRecorder
}

// MockDeploymentLabelFinisherMockRecorder is the mock recorder for MockDeploymentLabelFinisher.
type MockDeploymentLabelFinisherMockRecorder struct {
	mock *MockDeploymentLabelFinisher
}

// NewMockDeploymentLabelFinisher creates a new mock instance.
func NewMockDeploymentLabelFinisher(ctrl *gomock.Controller) *MockDeploymentLabelFinisher {
	mock := &MockDeploymentLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockDeploymentLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentLabelFinisher) EXPECT() *MockDeploymentLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockDeploymentLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.DeploymentLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.DeploymentLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockDeploymentLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockDeploymentLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockDeploymentLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDeploymentLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDeploymentLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockDeploymentLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDeploymentLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDeploymentLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockDeploymentLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.DeploymentLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.DeploymentLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockDeploymentLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockDeploymentLabelFinisher)(nil).One), ctx, exec)
}

// MockDeploymentLabelFinder is a mock of DeploymentLabelFinder interface.
type MockDeploymentLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentLabelFinderMockRecorder
}

// MockDeploymentLabelFinderMockRecorder is the mock recorder for MockDeploymentLabelFinder.
type MockDeploymentLabelFinderMockRecorder struct {
	mock *MockDeploymentLabelFinder
}

// NewMockDeploymentLabelFinder creates a new mock instance.
func NewMockDeploymentLabelFinder(ctrl *gomock.Controller) *MockDeploymentLabelFinder {
	mock := &MockDeploymentLabelFinder{ctrl: ctrl}
	mock.recorder = &MockDeploymentLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentLabelFinder) EXPECT() *MockDeploymentLabelFinderMockRecorder {
	return m.recorder
}

// FindDeploymentLabel mocks base method.
func (m *MockDeploymentLabelFinder) FindDeploymentLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.DeploymentLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindDeploymentLabel", varargs...)
	ret0, _ := ret[0].(*models.DeploymentLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDeploymentLabel indicates an expected call of FindDeploymentLabel.
func (mr *MockDeploymentLabelFinderMockRecorder) FindDeploymentLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeploymentLabel", reflect.TypeOf((*MockDeploymentLabelFinder)(nil).FindDeploymentLabel), varargs...)
}

// MockDeploymentLabelInserter is a mock of DeploymentLabelInserter interface.
type MockDeploymentLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentLabelInserterMockRecorder
}

// MockDeploymentLabelInserterMockRecorder is the mock recorder for MockDeploymentLabelInserter.
type MockDeploymentLabelInserterMockRecorder struct {
	mock *MockDeploymentLabelInserter
}

// NewMockDeploymentLabelInserter creates a new mock instance.
func NewMockDeploymentLabelInserter(ctrl *gomock.Controller) *MockDeploymentLabelInserter {
	mock := &MockDeploymentLabelInserter{ctrl: ctrl}
	mock.recorder = &MockDeploymentLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentLabelInserter) EXPECT() *MockDeploymentLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockDeploymentLabelInserter) Insert(o *models.DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockDeploymentLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDeploymentLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockDeploymentLabelUpdater is a mock of DeploymentLabelUpdater interface.
type MockDeploymentLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentLabelUpdaterMockRecorder
}

// MockDeploymentLabelUpdaterMockRecorder is the mock recorder for MockDeploymentLabelUpdater.
type MockDeploymentLabelUpdaterMockRecorder struct {
	mock *MockDeploymentLabelUpdater
}

// NewMockDeploymentLabelUpdater creates a new mock instance.
func NewMockDeploymentLabelUpdater(ctrl *gomock.Controller) *MockDeploymentLabelUpdater {
	mock := &MockDeploymentLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockDeploymentLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentLabelUpdater) EXPECT() *MockDeploymentLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockDeploymentLabelUpdater) Update(o *models.DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDeploymentLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeploymentLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockDeploymentLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockDeploymentLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockDeploymentLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockDeploymentLabelUpdater) UpdateAllSlice(o models.DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockDeploymentLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockDeploymentLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockDeploymentLabelDeleter is a mock of DeploymentLabelDeleter interface.
type MockDeploymentLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentLabelDeleterMockRecorder
}

// MockDeploymentLabelDeleterMockRecorder is the mock recorder for MockDeploymentLabelDeleter.
type MockDeploymentLabelDeleterMockRecorder struct {
	mock *MockDeploymentLabelDeleter
}

// NewMockDeploymentLabelDeleter creates a new mock instance.
func NewMockDeploymentLabelDeleter(ctrl *gomock.Controller) *MockDeploymentLabelDeleter {
	mock := &MockDeploymentLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockDeploymentLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentLabelDeleter) EXPECT() *MockDeploymentLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeploymentLabelDeleter) Delete(o *models.DeploymentLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDeploymentLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeploymentLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockDeploymentLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockDeploymentLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockDeploymentLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockDeploymentLabelDeleter) DeleteAllSlice(o models.DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockDeploymentLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockDeploymentLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockDeploymentLabelReloader is a mock of DeploymentLabelReloader interface.
type MockDeploymentLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentLabelReloaderMockRecorder
}

// MockDeploymentLabelReloaderMockRecorder is the mock recorder for MockDeploymentLabelReloader.
type MockDeploymentLabelReloaderMockRecorder struct {
	mock *MockDeploymentLabelReloader
}

// NewMockDeploymentLabelReloader creates a new mock instance.
func NewMockDeploymentLabelReloader(ctrl *gomock.Controller) *MockDeploymentLabelReloader {
	mock := &MockDeploymentLabelReloader{ctrl: ctrl}
	mock.recorder = &MockDeploymentLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentLabelReloader) EXPECT() *MockDeploymentLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockDeploymentLabelReloader) Reload(o *models.DeploymentLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockDeploymentLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockDeploymentLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockDeploymentLabelReloader) ReloadAll(o *models.DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockDeploymentLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockDeploymentLabelReloader)(nil).ReloadAll), o, ctx, exec)
}