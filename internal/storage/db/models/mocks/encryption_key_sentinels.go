//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_encryption_key_sentinels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockEncryptionKeySentinelUpserter is a mock of EncryptionKeySentinelUpserter interface.
type MockEncryptionKeySentinelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionKeySentinelUpserterMockRecorder
}

// MockEncryptionKeySentinelUpserterMockRecorder is the mock recorder for MockEncryptionKeySentinelUpserter.
type MockEncryptionKeySentinelUpserterMockRecorder struct {
	mock *MockEncryptionKeySentinelUpserter
}

// NewMockEncryptionKeySentinelUpserter creates a new mock instance.
func NewMockEncryptionKeySentinelUpserter(ctrl *gomock.Controller) *MockEncryptionKeySentinelUpserter {
	mock := &MockEncryptionKeySentinelUpserter{ctrl: ctrl}
	mock.recorder = &MockEncryptionKeySentinelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionKeySentinelUpserter) EXPECT() *MockEncryptionKeySentinelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockEncryptionKeySentinelUpserter) Upsert(o *models.EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockEncryptionKeySentinelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockEncryptionKeySentinelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockEncryptionKeySentinelFinisher is a mock of EncryptionKeySentinelFinisher interface.
type MockEncryptionKeySentinelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionKeySentinelFinisherMockRecorder
}

// MockEncryptionKeySentinelFinisherMockRecorder is the mock recorder for MockEncryptionKeySentinelFinisher.
type MockEncryptionKeySentinelFinisherMockRecorder struct {
	mock *MockEncryptionKeySentinelFinisher
}

// NewMockEncryptionKeySentinelFinisher creates a new mock instance.
func NewMockEncryptionKeySentinelFinisher(ctrl *gomock.Controller) *MockEncryptionKeySentinelFinisher {
	mock := &MockEncryptionKeySentinelFinisher{ctrl: ctrl}
	mock.recorder = &MockEncryptionKeySentinelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionKeySentinelFinisher) EXPECT() *MockEncryptionKeySentinelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockEncryptionKeySentinelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.EncryptionKeySentinelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.EncryptionKeySentinelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockEncryptionKeySentinelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockEncryptionKeySentinelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockEncryptionKeySentinelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockEncryptionKeySentinelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockEncryptionKeySentinelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockEncryptionKeySentinelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockEncryptionKeySentinelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockEncryptionKeySentinelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockEncryptionKeySentinelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.EncryptionKeySentinel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.EncryptionKeySentinel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockEncryptionKeySentinelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockEncryptionKeySentinelFinisher)(nil).One), ctx, exec)
}

// MockEncryptionKeySentinelFinder is a mock of EncryptionKeySentinelFinder interface.
type MockEncryptionKeySentinelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionKeySentinelFinderMockRecorder
}

// MockEncryptionKeySentinelFinderMockRecorder is the mock recorder for MockEncryptionKeySentinelFinder.
type MockEncryptionKeySentinelFinderMockRecorder struct {
	mock *MockEncryptionKeySentinelFinder
}

// NewMockEncryptionKeySentinelFinder creates a new mock instance.
func NewMockEncryptionKeySentinelFinder(ctrl *gomock.Controller) *MockEncryptionKeySentinelFinder {
	mock := &MockEncryptionKeySentinelFinder{ctrl: ctrl}
	mock.recorder = &MockEncryptionKeySentinelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionKeySentinelFinder) EXPECT() *MockEncryptionKeySentinelFinderMockRecorder {
	return m.recorder
}

// FindEncryptionKeySentinel mocks base method.
func (m *MockEncryptionKeySentinelFinder) FindEncryptionKeySentinel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.EncryptionKeySentinel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindEncryptionKeySentinel", varargs...)
	ret0, _ := ret[0].(*models.EncryptionKeySentinel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEncryptionKeySentinel indicates an expected call of FindEncryptionKeySentinel.
func (mr *MockEncryptionKeySentinelFinderMockRecorder) FindEncryptionKeySentinel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEncryptionKeySentinel", reflect.TypeOf((*MockEncryptionKeySentinelFinder)(nil).FindEncryptionKeySentinel), varargs...)
}

// MockEncryptionKeySentinelInserter is a mock of EncryptionKeySentinelInserter interface.
type MockEncryptionKeySentinelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionKeySentinelInserterMockRecorder
}

// MockEncryptionKeySentinelInserterMockRecorder is the mock recorder for MockEncryptionKeySentinelInserter.
type MockEncryptionKeySentinelInserterMockRecorder struct {
	mock *MockEncryptionKeySentinelInserter
}

// NewMockEncryptionKeySentinelInserter creates a new mock instance.
func NewMockEncryptionKeySentinelInserter(ctrl *gomock.Controller) *MockEncryptionKeySentinelInserter {
	mock := &MockEncryptionKeySentinelInserter{ctrl: ctrl}
	mock.recorder = &MockEncryptionKeySentinelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionKeySentinelInserter) EXPECT() *MockEncryptionKeySentinelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockEncryptionKeySentinelInserter) Insert(o *models.EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockEncryptionKeySentinelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockEncryptionKeySentinelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockEncryptionKeySentinelUpdater is a mock of EncryptionKeySentinelUpdater interface.
type MockEncryptionKeySentinelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionKeySentinelUpdaterMockRecorder
}

// MockEncryptionKeySentinelUpdaterMockRecorder is the mock recorder for MockEncryptionKeySentinelUpdater.
type MockEncryptionKeySentinelUpdaterMockRecorder struct {
	mock *MockEncryptionKeySentinelUpdater
}

// NewMockEncryptionKeySentinelUpdater creates a new mock instance.
func NewMockEncryptionKeySentinelUpdater(ctrl *gomock.Controller) *MockEncryptionKeySentinelUpdater {
	mock := &MockEncryptionKeySentinelUpdater{ctrl: ctrl}
	mock.recorder = &MockEncryptionKeySentinelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionKeySentinelUpdater) EXPECT() *MockEncryptionKeySentinelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockEncryptionKeySentinelUpdater) Update(o *models.EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEncryptionKeySentinelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEncryptionKeySentinelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockEncryptionKeySentinelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockEncryptionKeySentinelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockEncryptionKeySentinelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockEncryptionKeySentinelUpdater) UpdateAllSlice(o models.EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockEncryptionKeySentinelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockEncryptionKeySentinelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockEncryptionKeySentinelDeleter is a mock of EncryptionKeySentinelDeleter interface.
type MockEncryptionKeySentinelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionKeySentinelDeleterMockRecorder
}

// MockEncryptionKeySentinelDeleterMockRecorder is the mock recorder for MockEncryptionKeySentinelDeleter.
type MockEncryptionKeySentinelDeleterMockRecorder struct {
	mock *MockEncryptionKeySentinelDeleter
}

// NewMockEncryptionKeySentinelDeleter creates a new mock instance.
func NewMockEncryptionKeySentinelDeleter(ctrl *gomock.Controller) *MockEncryptionKeySentinelDeleter {
	mock := &MockEncryptionKeySentinelDeleter{ctrl: ctrl}
	mock.recorder = &MockEncryptionKeySentinelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionKeySentinelDeleter) EXPECT() *MockEncryptionKeySentinelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockEncryptionKeySentinelDeleter) Delete(o *models.EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEncryptionKeySentinelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEncryptionKeySentinelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockEncryptionKeySentinelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockEncryptionKeySentinelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockEncryptionKeySentinelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockEncryptionKeySentinelDeleter) DeleteAllSlice(o models.EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockEncryptionKeySentinelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockEncryptionKeySentinelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockEncryptionKeySentinelReloader is a mock of EncryptionKeySentinelReloader interface.
type MockEncryptionKeySentinelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionKeySentinelReloaderMockRecorder
}

// MockEncryptionKeySentinelReloaderMockRecorder is the mock recorder for MockEncryptionKeySentinelReloader.
type MockEncryptionKeySentinelReloaderMockRecorder struct {
	mock *MockEncryptionKeySentinelReloader
}

// NewMockEncryptionKeySentinelReloader creates a new mock instance.
func NewMockEncryptionKeySentinelReloader(ctrl *gomock.Controller) *MockEncryptionKeySentinelReloader {
	mock := &MockEncryptionKeySentinelReloader{ctrl: ctrl}
	mock.recorder = &MockEncryptionKeySentinelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionKeySentinelReloader) EXPECT() *MockEncryptionKeySentinelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockEncryptionKeySentinelReloader) Reload(o *models.EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockEncryptionKeySentinelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockEncryptionKeySentinelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockEncryptionKeySentinelReloader) ReloadAll(o *models.EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockEncryptionKeySentinelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockEncryptionKeySentinelReloader)(nil).ReloadAll), o, ctx, exec)
}
