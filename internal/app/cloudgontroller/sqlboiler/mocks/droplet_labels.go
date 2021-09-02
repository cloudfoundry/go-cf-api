// Code generated by MockGen. DO NOT EDIT.
// Source: psql_droplet_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockDropletLabelFinisher is a mock of DropletLabelFinisher interface.
type MockDropletLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockDropletLabelFinisherMockRecorder
}

// MockDropletLabelFinisherMockRecorder is the mock recorder for MockDropletLabelFinisher.
type MockDropletLabelFinisherMockRecorder struct {
	mock *MockDropletLabelFinisher
}

// NewMockDropletLabelFinisher creates a new mock instance.
func NewMockDropletLabelFinisher(ctrl *gomock.Controller) *MockDropletLabelFinisher {
	mock := &MockDropletLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockDropletLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletLabelFinisher) EXPECT() *MockDropletLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockDropletLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.DropletLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.DropletLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockDropletLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockDropletLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockDropletLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDropletLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDropletLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockDropletLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDropletLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDropletLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockDropletLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.DropletLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.DropletLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockDropletLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockDropletLabelFinisher)(nil).One), ctx, exec)
}

// MockDropletLabelFinder is a mock of DropletLabelFinder interface.
type MockDropletLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockDropletLabelFinderMockRecorder
}

// MockDropletLabelFinderMockRecorder is the mock recorder for MockDropletLabelFinder.
type MockDropletLabelFinderMockRecorder struct {
	mock *MockDropletLabelFinder
}

// NewMockDropletLabelFinder creates a new mock instance.
func NewMockDropletLabelFinder(ctrl *gomock.Controller) *MockDropletLabelFinder {
	mock := &MockDropletLabelFinder{ctrl: ctrl}
	mock.recorder = &MockDropletLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletLabelFinder) EXPECT() *MockDropletLabelFinderMockRecorder {
	return m.recorder
}

// FindDropletLabel mocks base method.
func (m *MockDropletLabelFinder) FindDropletLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.DropletLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindDropletLabel", varargs...)
	ret0, _ := ret[0].(*models.DropletLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDropletLabel indicates an expected call of FindDropletLabel.
func (mr *MockDropletLabelFinderMockRecorder) FindDropletLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDropletLabel", reflect.TypeOf((*MockDropletLabelFinder)(nil).FindDropletLabel), varargs...)
}

// MockDropletLabelInserter is a mock of DropletLabelInserter interface.
type MockDropletLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockDropletLabelInserterMockRecorder
}

// MockDropletLabelInserterMockRecorder is the mock recorder for MockDropletLabelInserter.
type MockDropletLabelInserterMockRecorder struct {
	mock *MockDropletLabelInserter
}

// NewMockDropletLabelInserter creates a new mock instance.
func NewMockDropletLabelInserter(ctrl *gomock.Controller) *MockDropletLabelInserter {
	mock := &MockDropletLabelInserter{ctrl: ctrl}
	mock.recorder = &MockDropletLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletLabelInserter) EXPECT() *MockDropletLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockDropletLabelInserter) Insert(o *models.DropletLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockDropletLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDropletLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockDropletLabelUpdater is a mock of DropletLabelUpdater interface.
type MockDropletLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDropletLabelUpdaterMockRecorder
}

// MockDropletLabelUpdaterMockRecorder is the mock recorder for MockDropletLabelUpdater.
type MockDropletLabelUpdaterMockRecorder struct {
	mock *MockDropletLabelUpdater
}

// NewMockDropletLabelUpdater creates a new mock instance.
func NewMockDropletLabelUpdater(ctrl *gomock.Controller) *MockDropletLabelUpdater {
	mock := &MockDropletLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockDropletLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletLabelUpdater) EXPECT() *MockDropletLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockDropletLabelUpdater) Update(o *models.DropletLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDropletLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDropletLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockDropletLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockDropletLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockDropletLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockDropletLabelUpdater) UpdateAllSlice(o models.DropletLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockDropletLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockDropletLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockDropletLabelDeleter is a mock of DropletLabelDeleter interface.
type MockDropletLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDropletLabelDeleterMockRecorder
}

// MockDropletLabelDeleterMockRecorder is the mock recorder for MockDropletLabelDeleter.
type MockDropletLabelDeleterMockRecorder struct {
	mock *MockDropletLabelDeleter
}

// NewMockDropletLabelDeleter creates a new mock instance.
func NewMockDropletLabelDeleter(ctrl *gomock.Controller) *MockDropletLabelDeleter {
	mock := &MockDropletLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockDropletLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletLabelDeleter) EXPECT() *MockDropletLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDropletLabelDeleter) Delete(o *models.DropletLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDropletLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDropletLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockDropletLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockDropletLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockDropletLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockDropletLabelDeleter) DeleteAllSlice(o models.DropletLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockDropletLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockDropletLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockDropletLabelReloader is a mock of DropletLabelReloader interface.
type MockDropletLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockDropletLabelReloaderMockRecorder
}

// MockDropletLabelReloaderMockRecorder is the mock recorder for MockDropletLabelReloader.
type MockDropletLabelReloaderMockRecorder struct {
	mock *MockDropletLabelReloader
}

// NewMockDropletLabelReloader creates a new mock instance.
func NewMockDropletLabelReloader(ctrl *gomock.Controller) *MockDropletLabelReloader {
	mock := &MockDropletLabelReloader{ctrl: ctrl}
	mock.recorder = &MockDropletLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletLabelReloader) EXPECT() *MockDropletLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockDropletLabelReloader) Reload(o *models.DropletLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockDropletLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockDropletLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockDropletLabelReloader) ReloadAll(o *models.DropletLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockDropletLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockDropletLabelReloader)(nil).ReloadAll), o, ctx, exec)
}