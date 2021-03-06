//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_buildpack_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockBuildpackLabelUpserter is a mock of BuildpackLabelUpserter interface.
type MockBuildpackLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackLabelUpserterMockRecorder
}

// MockBuildpackLabelUpserterMockRecorder is the mock recorder for MockBuildpackLabelUpserter.
type MockBuildpackLabelUpserterMockRecorder struct {
	mock *MockBuildpackLabelUpserter
}

// NewMockBuildpackLabelUpserter creates a new mock instance.
func NewMockBuildpackLabelUpserter(ctrl *gomock.Controller) *MockBuildpackLabelUpserter {
	mock := &MockBuildpackLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockBuildpackLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackLabelUpserter) EXPECT() *MockBuildpackLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockBuildpackLabelUpserter) Upsert(o *models.BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockBuildpackLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockBuildpackLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockBuildpackLabelFinisher is a mock of BuildpackLabelFinisher interface.
type MockBuildpackLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackLabelFinisherMockRecorder
}

// MockBuildpackLabelFinisherMockRecorder is the mock recorder for MockBuildpackLabelFinisher.
type MockBuildpackLabelFinisherMockRecorder struct {
	mock *MockBuildpackLabelFinisher
}

// NewMockBuildpackLabelFinisher creates a new mock instance.
func NewMockBuildpackLabelFinisher(ctrl *gomock.Controller) *MockBuildpackLabelFinisher {
	mock := &MockBuildpackLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockBuildpackLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackLabelFinisher) EXPECT() *MockBuildpackLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockBuildpackLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.BuildpackLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.BuildpackLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockBuildpackLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockBuildpackLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockBuildpackLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockBuildpackLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockBuildpackLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockBuildpackLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockBuildpackLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockBuildpackLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockBuildpackLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.BuildpackLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.BuildpackLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockBuildpackLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockBuildpackLabelFinisher)(nil).One), ctx, exec)
}

// MockBuildpackLabelFinder is a mock of BuildpackLabelFinder interface.
type MockBuildpackLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackLabelFinderMockRecorder
}

// MockBuildpackLabelFinderMockRecorder is the mock recorder for MockBuildpackLabelFinder.
type MockBuildpackLabelFinderMockRecorder struct {
	mock *MockBuildpackLabelFinder
}

// NewMockBuildpackLabelFinder creates a new mock instance.
func NewMockBuildpackLabelFinder(ctrl *gomock.Controller) *MockBuildpackLabelFinder {
	mock := &MockBuildpackLabelFinder{ctrl: ctrl}
	mock.recorder = &MockBuildpackLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackLabelFinder) EXPECT() *MockBuildpackLabelFinderMockRecorder {
	return m.recorder
}

// FindBuildpackLabel mocks base method.
func (m *MockBuildpackLabelFinder) FindBuildpackLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.BuildpackLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindBuildpackLabel", varargs...)
	ret0, _ := ret[0].(*models.BuildpackLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBuildpackLabel indicates an expected call of FindBuildpackLabel.
func (mr *MockBuildpackLabelFinderMockRecorder) FindBuildpackLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBuildpackLabel", reflect.TypeOf((*MockBuildpackLabelFinder)(nil).FindBuildpackLabel), varargs...)
}

// MockBuildpackLabelInserter is a mock of BuildpackLabelInserter interface.
type MockBuildpackLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackLabelInserterMockRecorder
}

// MockBuildpackLabelInserterMockRecorder is the mock recorder for MockBuildpackLabelInserter.
type MockBuildpackLabelInserterMockRecorder struct {
	mock *MockBuildpackLabelInserter
}

// NewMockBuildpackLabelInserter creates a new mock instance.
func NewMockBuildpackLabelInserter(ctrl *gomock.Controller) *MockBuildpackLabelInserter {
	mock := &MockBuildpackLabelInserter{ctrl: ctrl}
	mock.recorder = &MockBuildpackLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackLabelInserter) EXPECT() *MockBuildpackLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockBuildpackLabelInserter) Insert(o *models.BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockBuildpackLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBuildpackLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockBuildpackLabelUpdater is a mock of BuildpackLabelUpdater interface.
type MockBuildpackLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackLabelUpdaterMockRecorder
}

// MockBuildpackLabelUpdaterMockRecorder is the mock recorder for MockBuildpackLabelUpdater.
type MockBuildpackLabelUpdaterMockRecorder struct {
	mock *MockBuildpackLabelUpdater
}

// NewMockBuildpackLabelUpdater creates a new mock instance.
func NewMockBuildpackLabelUpdater(ctrl *gomock.Controller) *MockBuildpackLabelUpdater {
	mock := &MockBuildpackLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockBuildpackLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackLabelUpdater) EXPECT() *MockBuildpackLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockBuildpackLabelUpdater) Update(o *models.BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBuildpackLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBuildpackLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockBuildpackLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockBuildpackLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockBuildpackLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockBuildpackLabelUpdater) UpdateAllSlice(o models.BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockBuildpackLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockBuildpackLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockBuildpackLabelDeleter is a mock of BuildpackLabelDeleter interface.
type MockBuildpackLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackLabelDeleterMockRecorder
}

// MockBuildpackLabelDeleterMockRecorder is the mock recorder for MockBuildpackLabelDeleter.
type MockBuildpackLabelDeleterMockRecorder struct {
	mock *MockBuildpackLabelDeleter
}

// NewMockBuildpackLabelDeleter creates a new mock instance.
func NewMockBuildpackLabelDeleter(ctrl *gomock.Controller) *MockBuildpackLabelDeleter {
	mock := &MockBuildpackLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockBuildpackLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackLabelDeleter) EXPECT() *MockBuildpackLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBuildpackLabelDeleter) Delete(o *models.BuildpackLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBuildpackLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBuildpackLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockBuildpackLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockBuildpackLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockBuildpackLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockBuildpackLabelDeleter) DeleteAllSlice(o models.BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockBuildpackLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockBuildpackLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockBuildpackLabelReloader is a mock of BuildpackLabelReloader interface.
type MockBuildpackLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackLabelReloaderMockRecorder
}

// MockBuildpackLabelReloaderMockRecorder is the mock recorder for MockBuildpackLabelReloader.
type MockBuildpackLabelReloaderMockRecorder struct {
	mock *MockBuildpackLabelReloader
}

// NewMockBuildpackLabelReloader creates a new mock instance.
func NewMockBuildpackLabelReloader(ctrl *gomock.Controller) *MockBuildpackLabelReloader {
	mock := &MockBuildpackLabelReloader{ctrl: ctrl}
	mock.recorder = &MockBuildpackLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpackLabelReloader) EXPECT() *MockBuildpackLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockBuildpackLabelReloader) Reload(o *models.BuildpackLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockBuildpackLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockBuildpackLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockBuildpackLabelReloader) ReloadAll(o *models.BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockBuildpackLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockBuildpackLabelReloader)(nil).ReloadAll), o, ctx, exec)
}
