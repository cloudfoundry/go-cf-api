//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_env_groups.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockEnvGroupUpserter is a mock of EnvGroupUpserter interface.
type MockEnvGroupUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvGroupUpserterMockRecorder
}

// MockEnvGroupUpserterMockRecorder is the mock recorder for MockEnvGroupUpserter.
type MockEnvGroupUpserterMockRecorder struct {
	mock *MockEnvGroupUpserter
}

// NewMockEnvGroupUpserter creates a new mock instance.
func NewMockEnvGroupUpserter(ctrl *gomock.Controller) *MockEnvGroupUpserter {
	mock := &MockEnvGroupUpserter{ctrl: ctrl}
	mock.recorder = &MockEnvGroupUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvGroupUpserter) EXPECT() *MockEnvGroupUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockEnvGroupUpserter) Upsert(o *models.EnvGroup, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockEnvGroupUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockEnvGroupUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockEnvGroupFinisher is a mock of EnvGroupFinisher interface.
type MockEnvGroupFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockEnvGroupFinisherMockRecorder
}

// MockEnvGroupFinisherMockRecorder is the mock recorder for MockEnvGroupFinisher.
type MockEnvGroupFinisherMockRecorder struct {
	mock *MockEnvGroupFinisher
}

// NewMockEnvGroupFinisher creates a new mock instance.
func NewMockEnvGroupFinisher(ctrl *gomock.Controller) *MockEnvGroupFinisher {
	mock := &MockEnvGroupFinisher{ctrl: ctrl}
	mock.recorder = &MockEnvGroupFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvGroupFinisher) EXPECT() *MockEnvGroupFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockEnvGroupFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.EnvGroupSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.EnvGroupSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockEnvGroupFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockEnvGroupFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockEnvGroupFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockEnvGroupFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockEnvGroupFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockEnvGroupFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockEnvGroupFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockEnvGroupFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockEnvGroupFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.EnvGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.EnvGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockEnvGroupFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockEnvGroupFinisher)(nil).One), ctx, exec)
}

// MockEnvGroupFinder is a mock of EnvGroupFinder interface.
type MockEnvGroupFinder struct {
	ctrl     *gomock.Controller
	recorder *MockEnvGroupFinderMockRecorder
}

// MockEnvGroupFinderMockRecorder is the mock recorder for MockEnvGroupFinder.
type MockEnvGroupFinderMockRecorder struct {
	mock *MockEnvGroupFinder
}

// NewMockEnvGroupFinder creates a new mock instance.
func NewMockEnvGroupFinder(ctrl *gomock.Controller) *MockEnvGroupFinder {
	mock := &MockEnvGroupFinder{ctrl: ctrl}
	mock.recorder = &MockEnvGroupFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvGroupFinder) EXPECT() *MockEnvGroupFinderMockRecorder {
	return m.recorder
}

// FindEnvGroup mocks base method.
func (m *MockEnvGroupFinder) FindEnvGroup(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.EnvGroup, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindEnvGroup", varargs...)
	ret0, _ := ret[0].(*models.EnvGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEnvGroup indicates an expected call of FindEnvGroup.
func (mr *MockEnvGroupFinderMockRecorder) FindEnvGroup(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEnvGroup", reflect.TypeOf((*MockEnvGroupFinder)(nil).FindEnvGroup), varargs...)
}

// MockEnvGroupInserter is a mock of EnvGroupInserter interface.
type MockEnvGroupInserter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvGroupInserterMockRecorder
}

// MockEnvGroupInserterMockRecorder is the mock recorder for MockEnvGroupInserter.
type MockEnvGroupInserterMockRecorder struct {
	mock *MockEnvGroupInserter
}

// NewMockEnvGroupInserter creates a new mock instance.
func NewMockEnvGroupInserter(ctrl *gomock.Controller) *MockEnvGroupInserter {
	mock := &MockEnvGroupInserter{ctrl: ctrl}
	mock.recorder = &MockEnvGroupInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvGroupInserter) EXPECT() *MockEnvGroupInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockEnvGroupInserter) Insert(o *models.EnvGroup, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockEnvGroupInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockEnvGroupInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockEnvGroupUpdater is a mock of EnvGroupUpdater interface.
type MockEnvGroupUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockEnvGroupUpdaterMockRecorder
}

// MockEnvGroupUpdaterMockRecorder is the mock recorder for MockEnvGroupUpdater.
type MockEnvGroupUpdaterMockRecorder struct {
	mock *MockEnvGroupUpdater
}

// NewMockEnvGroupUpdater creates a new mock instance.
func NewMockEnvGroupUpdater(ctrl *gomock.Controller) *MockEnvGroupUpdater {
	mock := &MockEnvGroupUpdater{ctrl: ctrl}
	mock.recorder = &MockEnvGroupUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvGroupUpdater) EXPECT() *MockEnvGroupUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockEnvGroupUpdater) Update(o *models.EnvGroup, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEnvGroupUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEnvGroupUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockEnvGroupUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockEnvGroupUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockEnvGroupUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockEnvGroupUpdater) UpdateAllSlice(o models.EnvGroupSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockEnvGroupUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockEnvGroupUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockEnvGroupDeleter is a mock of EnvGroupDeleter interface.
type MockEnvGroupDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvGroupDeleterMockRecorder
}

// MockEnvGroupDeleterMockRecorder is the mock recorder for MockEnvGroupDeleter.
type MockEnvGroupDeleterMockRecorder struct {
	mock *MockEnvGroupDeleter
}

// NewMockEnvGroupDeleter creates a new mock instance.
func NewMockEnvGroupDeleter(ctrl *gomock.Controller) *MockEnvGroupDeleter {
	mock := &MockEnvGroupDeleter{ctrl: ctrl}
	mock.recorder = &MockEnvGroupDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvGroupDeleter) EXPECT() *MockEnvGroupDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockEnvGroupDeleter) Delete(o *models.EnvGroup, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEnvGroupDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEnvGroupDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockEnvGroupDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockEnvGroupDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockEnvGroupDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockEnvGroupDeleter) DeleteAllSlice(o models.EnvGroupSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockEnvGroupDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockEnvGroupDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockEnvGroupReloader is a mock of EnvGroupReloader interface.
type MockEnvGroupReloader struct {
	ctrl     *gomock.Controller
	recorder *MockEnvGroupReloaderMockRecorder
}

// MockEnvGroupReloaderMockRecorder is the mock recorder for MockEnvGroupReloader.
type MockEnvGroupReloaderMockRecorder struct {
	mock *MockEnvGroupReloader
}

// NewMockEnvGroupReloader creates a new mock instance.
func NewMockEnvGroupReloader(ctrl *gomock.Controller) *MockEnvGroupReloader {
	mock := &MockEnvGroupReloader{ctrl: ctrl}
	mock.recorder = &MockEnvGroupReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvGroupReloader) EXPECT() *MockEnvGroupReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockEnvGroupReloader) Reload(o *models.EnvGroup, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockEnvGroupReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockEnvGroupReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockEnvGroupReloader) ReloadAll(o *models.EnvGroupSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockEnvGroupReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockEnvGroupReloader)(nil).ReloadAll), o, ctx, exec)
}
