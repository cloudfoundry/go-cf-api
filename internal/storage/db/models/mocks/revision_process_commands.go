//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_revision_process_commands.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockRevisionProcessCommandUpserter is a mock of RevisionProcessCommandUpserter interface.
type MockRevisionProcessCommandUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionProcessCommandUpserterMockRecorder
}

// MockRevisionProcessCommandUpserterMockRecorder is the mock recorder for MockRevisionProcessCommandUpserter.
type MockRevisionProcessCommandUpserterMockRecorder struct {
	mock *MockRevisionProcessCommandUpserter
}

// NewMockRevisionProcessCommandUpserter creates a new mock instance.
func NewMockRevisionProcessCommandUpserter(ctrl *gomock.Controller) *MockRevisionProcessCommandUpserter {
	mock := &MockRevisionProcessCommandUpserter{ctrl: ctrl}
	mock.recorder = &MockRevisionProcessCommandUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionProcessCommandUpserter) EXPECT() *MockRevisionProcessCommandUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockRevisionProcessCommandUpserter) Upsert(o *models.RevisionProcessCommand, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRevisionProcessCommandUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRevisionProcessCommandUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockRevisionProcessCommandFinisher is a mock of RevisionProcessCommandFinisher interface.
type MockRevisionProcessCommandFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionProcessCommandFinisherMockRecorder
}

// MockRevisionProcessCommandFinisherMockRecorder is the mock recorder for MockRevisionProcessCommandFinisher.
type MockRevisionProcessCommandFinisherMockRecorder struct {
	mock *MockRevisionProcessCommandFinisher
}

// NewMockRevisionProcessCommandFinisher creates a new mock instance.
func NewMockRevisionProcessCommandFinisher(ctrl *gomock.Controller) *MockRevisionProcessCommandFinisher {
	mock := &MockRevisionProcessCommandFinisher{ctrl: ctrl}
	mock.recorder = &MockRevisionProcessCommandFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionProcessCommandFinisher) EXPECT() *MockRevisionProcessCommandFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockRevisionProcessCommandFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.RevisionProcessCommandSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.RevisionProcessCommandSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockRevisionProcessCommandFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRevisionProcessCommandFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockRevisionProcessCommandFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRevisionProcessCommandFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRevisionProcessCommandFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockRevisionProcessCommandFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRevisionProcessCommandFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRevisionProcessCommandFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockRevisionProcessCommandFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.RevisionProcessCommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.RevisionProcessCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockRevisionProcessCommandFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockRevisionProcessCommandFinisher)(nil).One), ctx, exec)
}

// MockRevisionProcessCommandFinder is a mock of RevisionProcessCommandFinder interface.
type MockRevisionProcessCommandFinder struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionProcessCommandFinderMockRecorder
}

// MockRevisionProcessCommandFinderMockRecorder is the mock recorder for MockRevisionProcessCommandFinder.
type MockRevisionProcessCommandFinderMockRecorder struct {
	mock *MockRevisionProcessCommandFinder
}

// NewMockRevisionProcessCommandFinder creates a new mock instance.
func NewMockRevisionProcessCommandFinder(ctrl *gomock.Controller) *MockRevisionProcessCommandFinder {
	mock := &MockRevisionProcessCommandFinder{ctrl: ctrl}
	mock.recorder = &MockRevisionProcessCommandFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionProcessCommandFinder) EXPECT() *MockRevisionProcessCommandFinderMockRecorder {
	return m.recorder
}

// FindRevisionProcessCommand mocks base method.
func (m *MockRevisionProcessCommandFinder) FindRevisionProcessCommand(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.RevisionProcessCommand, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindRevisionProcessCommand", varargs...)
	ret0, _ := ret[0].(*models.RevisionProcessCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRevisionProcessCommand indicates an expected call of FindRevisionProcessCommand.
func (mr *MockRevisionProcessCommandFinderMockRecorder) FindRevisionProcessCommand(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRevisionProcessCommand", reflect.TypeOf((*MockRevisionProcessCommandFinder)(nil).FindRevisionProcessCommand), varargs...)
}

// MockRevisionProcessCommandInserter is a mock of RevisionProcessCommandInserter interface.
type MockRevisionProcessCommandInserter struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionProcessCommandInserterMockRecorder
}

// MockRevisionProcessCommandInserterMockRecorder is the mock recorder for MockRevisionProcessCommandInserter.
type MockRevisionProcessCommandInserterMockRecorder struct {
	mock *MockRevisionProcessCommandInserter
}

// NewMockRevisionProcessCommandInserter creates a new mock instance.
func NewMockRevisionProcessCommandInserter(ctrl *gomock.Controller) *MockRevisionProcessCommandInserter {
	mock := &MockRevisionProcessCommandInserter{ctrl: ctrl}
	mock.recorder = &MockRevisionProcessCommandInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionProcessCommandInserter) EXPECT() *MockRevisionProcessCommandInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockRevisionProcessCommandInserter) Insert(o *models.RevisionProcessCommand, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRevisionProcessCommandInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRevisionProcessCommandInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockRevisionProcessCommandUpdater is a mock of RevisionProcessCommandUpdater interface.
type MockRevisionProcessCommandUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionProcessCommandUpdaterMockRecorder
}

// MockRevisionProcessCommandUpdaterMockRecorder is the mock recorder for MockRevisionProcessCommandUpdater.
type MockRevisionProcessCommandUpdaterMockRecorder struct {
	mock *MockRevisionProcessCommandUpdater
}

// NewMockRevisionProcessCommandUpdater creates a new mock instance.
func NewMockRevisionProcessCommandUpdater(ctrl *gomock.Controller) *MockRevisionProcessCommandUpdater {
	mock := &MockRevisionProcessCommandUpdater{ctrl: ctrl}
	mock.recorder = &MockRevisionProcessCommandUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionProcessCommandUpdater) EXPECT() *MockRevisionProcessCommandUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockRevisionProcessCommandUpdater) Update(o *models.RevisionProcessCommand, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRevisionProcessCommandUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRevisionProcessCommandUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockRevisionProcessCommandUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockRevisionProcessCommandUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockRevisionProcessCommandUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockRevisionProcessCommandUpdater) UpdateAllSlice(o models.RevisionProcessCommandSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockRevisionProcessCommandUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockRevisionProcessCommandUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockRevisionProcessCommandDeleter is a mock of RevisionProcessCommandDeleter interface.
type MockRevisionProcessCommandDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionProcessCommandDeleterMockRecorder
}

// MockRevisionProcessCommandDeleterMockRecorder is the mock recorder for MockRevisionProcessCommandDeleter.
type MockRevisionProcessCommandDeleterMockRecorder struct {
	mock *MockRevisionProcessCommandDeleter
}

// NewMockRevisionProcessCommandDeleter creates a new mock instance.
func NewMockRevisionProcessCommandDeleter(ctrl *gomock.Controller) *MockRevisionProcessCommandDeleter {
	mock := &MockRevisionProcessCommandDeleter{ctrl: ctrl}
	mock.recorder = &MockRevisionProcessCommandDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionProcessCommandDeleter) EXPECT() *MockRevisionProcessCommandDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRevisionProcessCommandDeleter) Delete(o *models.RevisionProcessCommand, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRevisionProcessCommandDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRevisionProcessCommandDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockRevisionProcessCommandDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockRevisionProcessCommandDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockRevisionProcessCommandDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockRevisionProcessCommandDeleter) DeleteAllSlice(o models.RevisionProcessCommandSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockRevisionProcessCommandDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockRevisionProcessCommandDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockRevisionProcessCommandReloader is a mock of RevisionProcessCommandReloader interface.
type MockRevisionProcessCommandReloader struct {
	ctrl     *gomock.Controller
	recorder *MockRevisionProcessCommandReloaderMockRecorder
}

// MockRevisionProcessCommandReloaderMockRecorder is the mock recorder for MockRevisionProcessCommandReloader.
type MockRevisionProcessCommandReloaderMockRecorder struct {
	mock *MockRevisionProcessCommandReloader
}

// NewMockRevisionProcessCommandReloader creates a new mock instance.
func NewMockRevisionProcessCommandReloader(ctrl *gomock.Controller) *MockRevisionProcessCommandReloader {
	mock := &MockRevisionProcessCommandReloader{ctrl: ctrl}
	mock.recorder = &MockRevisionProcessCommandReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevisionProcessCommandReloader) EXPECT() *MockRevisionProcessCommandReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockRevisionProcessCommandReloader) Reload(o *models.RevisionProcessCommand, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockRevisionProcessCommandReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockRevisionProcessCommandReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockRevisionProcessCommandReloader) ReloadAll(o *models.RevisionProcessCommandSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockRevisionProcessCommandReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockRevisionProcessCommandReloader)(nil).ReloadAll), o, ctx, exec)
}
