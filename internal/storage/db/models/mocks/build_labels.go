//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_build_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
)

// MockBuildLabelUpserter is a mock of BuildLabelUpserter interface.
type MockBuildLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLabelUpserterMockRecorder
}

// MockBuildLabelUpserterMockRecorder is the mock recorder for MockBuildLabelUpserter.
type MockBuildLabelUpserterMockRecorder struct {
	mock *MockBuildLabelUpserter
}

// NewMockBuildLabelUpserter creates a new mock instance.
func NewMockBuildLabelUpserter(ctrl *gomock.Controller) *MockBuildLabelUpserter {
	mock := &MockBuildLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockBuildLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLabelUpserter) EXPECT() *MockBuildLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockBuildLabelUpserter) Upsert(o *models.BuildLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockBuildLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockBuildLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockBuildLabelFinisher is a mock of BuildLabelFinisher interface.
type MockBuildLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLabelFinisherMockRecorder
}

// MockBuildLabelFinisherMockRecorder is the mock recorder for MockBuildLabelFinisher.
type MockBuildLabelFinisherMockRecorder struct {
	mock *MockBuildLabelFinisher
}

// NewMockBuildLabelFinisher creates a new mock instance.
func NewMockBuildLabelFinisher(ctrl *gomock.Controller) *MockBuildLabelFinisher {
	mock := &MockBuildLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockBuildLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLabelFinisher) EXPECT() *MockBuildLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockBuildLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.BuildLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.BuildLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockBuildLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockBuildLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockBuildLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockBuildLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockBuildLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockBuildLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockBuildLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockBuildLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockBuildLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.BuildLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.BuildLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockBuildLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockBuildLabelFinisher)(nil).One), ctx, exec)
}

// MockBuildLabelFinder is a mock of BuildLabelFinder interface.
type MockBuildLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLabelFinderMockRecorder
}

// MockBuildLabelFinderMockRecorder is the mock recorder for MockBuildLabelFinder.
type MockBuildLabelFinderMockRecorder struct {
	mock *MockBuildLabelFinder
}

// NewMockBuildLabelFinder creates a new mock instance.
func NewMockBuildLabelFinder(ctrl *gomock.Controller) *MockBuildLabelFinder {
	mock := &MockBuildLabelFinder{ctrl: ctrl}
	mock.recorder = &MockBuildLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLabelFinder) EXPECT() *MockBuildLabelFinderMockRecorder {
	return m.recorder
}

// FindBuildLabel mocks base method.
func (m *MockBuildLabelFinder) FindBuildLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.BuildLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindBuildLabel", varargs...)
	ret0, _ := ret[0].(*models.BuildLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBuildLabel indicates an expected call of FindBuildLabel.
func (mr *MockBuildLabelFinderMockRecorder) FindBuildLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBuildLabel", reflect.TypeOf((*MockBuildLabelFinder)(nil).FindBuildLabel), varargs...)
}

// MockBuildLabelInserter is a mock of BuildLabelInserter interface.
type MockBuildLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLabelInserterMockRecorder
}

// MockBuildLabelInserterMockRecorder is the mock recorder for MockBuildLabelInserter.
type MockBuildLabelInserterMockRecorder struct {
	mock *MockBuildLabelInserter
}

// NewMockBuildLabelInserter creates a new mock instance.
func NewMockBuildLabelInserter(ctrl *gomock.Controller) *MockBuildLabelInserter {
	mock := &MockBuildLabelInserter{ctrl: ctrl}
	mock.recorder = &MockBuildLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLabelInserter) EXPECT() *MockBuildLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockBuildLabelInserter) Insert(o *models.BuildLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockBuildLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBuildLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockBuildLabelUpdater is a mock of BuildLabelUpdater interface.
type MockBuildLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLabelUpdaterMockRecorder
}

// MockBuildLabelUpdaterMockRecorder is the mock recorder for MockBuildLabelUpdater.
type MockBuildLabelUpdaterMockRecorder struct {
	mock *MockBuildLabelUpdater
}

// NewMockBuildLabelUpdater creates a new mock instance.
func NewMockBuildLabelUpdater(ctrl *gomock.Controller) *MockBuildLabelUpdater {
	mock := &MockBuildLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockBuildLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLabelUpdater) EXPECT() *MockBuildLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockBuildLabelUpdater) Update(o *models.BuildLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBuildLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBuildLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockBuildLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockBuildLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockBuildLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockBuildLabelUpdater) UpdateAllSlice(o models.BuildLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockBuildLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockBuildLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockBuildLabelDeleter is a mock of BuildLabelDeleter interface.
type MockBuildLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLabelDeleterMockRecorder
}

// MockBuildLabelDeleterMockRecorder is the mock recorder for MockBuildLabelDeleter.
type MockBuildLabelDeleterMockRecorder struct {
	mock *MockBuildLabelDeleter
}

// NewMockBuildLabelDeleter creates a new mock instance.
func NewMockBuildLabelDeleter(ctrl *gomock.Controller) *MockBuildLabelDeleter {
	mock := &MockBuildLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockBuildLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLabelDeleter) EXPECT() *MockBuildLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBuildLabelDeleter) Delete(o *models.BuildLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBuildLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBuildLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockBuildLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockBuildLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockBuildLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockBuildLabelDeleter) DeleteAllSlice(o models.BuildLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockBuildLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockBuildLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockBuildLabelReloader is a mock of BuildLabelReloader interface.
type MockBuildLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLabelReloaderMockRecorder
}

// MockBuildLabelReloaderMockRecorder is the mock recorder for MockBuildLabelReloader.
type MockBuildLabelReloaderMockRecorder struct {
	mock *MockBuildLabelReloader
}

// NewMockBuildLabelReloader creates a new mock instance.
func NewMockBuildLabelReloader(ctrl *gomock.Controller) *MockBuildLabelReloader {
	mock := &MockBuildLabelReloader{ctrl: ctrl}
	mock.recorder = &MockBuildLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLabelReloader) EXPECT() *MockBuildLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockBuildLabelReloader) Reload(o *models.BuildLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockBuildLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockBuildLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockBuildLabelReloader) ReloadAll(o *models.BuildLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockBuildLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockBuildLabelReloader)(nil).ReloadAll), o, ctx, exec)
}
