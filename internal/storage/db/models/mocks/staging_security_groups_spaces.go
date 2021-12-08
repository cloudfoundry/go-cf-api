//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_staging_security_groups_spaces.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
)

// MockStagingSecurityGroupsSpaceUpserter is a mock of StagingSecurityGroupsSpaceUpserter interface.
type MockStagingSecurityGroupsSpaceUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockStagingSecurityGroupsSpaceUpserterMockRecorder
}

// MockStagingSecurityGroupsSpaceUpserterMockRecorder is the mock recorder for MockStagingSecurityGroupsSpaceUpserter.
type MockStagingSecurityGroupsSpaceUpserterMockRecorder struct {
	mock *MockStagingSecurityGroupsSpaceUpserter
}

// NewMockStagingSecurityGroupsSpaceUpserter creates a new mock instance.
func NewMockStagingSecurityGroupsSpaceUpserter(ctrl *gomock.Controller) *MockStagingSecurityGroupsSpaceUpserter {
	mock := &MockStagingSecurityGroupsSpaceUpserter{ctrl: ctrl}
	mock.recorder = &MockStagingSecurityGroupsSpaceUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagingSecurityGroupsSpaceUpserter) EXPECT() *MockStagingSecurityGroupsSpaceUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockStagingSecurityGroupsSpaceUpserter) Upsert(o *models.StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockStagingSecurityGroupsSpaceUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockStagingSecurityGroupsSpaceUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockStagingSecurityGroupsSpaceFinisher is a mock of StagingSecurityGroupsSpaceFinisher interface.
type MockStagingSecurityGroupsSpaceFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockStagingSecurityGroupsSpaceFinisherMockRecorder
}

// MockStagingSecurityGroupsSpaceFinisherMockRecorder is the mock recorder for MockStagingSecurityGroupsSpaceFinisher.
type MockStagingSecurityGroupsSpaceFinisherMockRecorder struct {
	mock *MockStagingSecurityGroupsSpaceFinisher
}

// NewMockStagingSecurityGroupsSpaceFinisher creates a new mock instance.
func NewMockStagingSecurityGroupsSpaceFinisher(ctrl *gomock.Controller) *MockStagingSecurityGroupsSpaceFinisher {
	mock := &MockStagingSecurityGroupsSpaceFinisher{ctrl: ctrl}
	mock.recorder = &MockStagingSecurityGroupsSpaceFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagingSecurityGroupsSpaceFinisher) EXPECT() *MockStagingSecurityGroupsSpaceFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockStagingSecurityGroupsSpaceFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.StagingSecurityGroupsSpaceSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.StagingSecurityGroupsSpaceSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockStagingSecurityGroupsSpaceFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockStagingSecurityGroupsSpaceFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockStagingSecurityGroupsSpaceFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockStagingSecurityGroupsSpaceFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockStagingSecurityGroupsSpaceFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockStagingSecurityGroupsSpaceFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockStagingSecurityGroupsSpaceFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockStagingSecurityGroupsSpaceFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockStagingSecurityGroupsSpaceFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.StagingSecurityGroupsSpace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.StagingSecurityGroupsSpace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockStagingSecurityGroupsSpaceFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockStagingSecurityGroupsSpaceFinisher)(nil).One), ctx, exec)
}

// MockStagingSecurityGroupsSpaceFinder is a mock of StagingSecurityGroupsSpaceFinder interface.
type MockStagingSecurityGroupsSpaceFinder struct {
	ctrl     *gomock.Controller
	recorder *MockStagingSecurityGroupsSpaceFinderMockRecorder
}

// MockStagingSecurityGroupsSpaceFinderMockRecorder is the mock recorder for MockStagingSecurityGroupsSpaceFinder.
type MockStagingSecurityGroupsSpaceFinderMockRecorder struct {
	mock *MockStagingSecurityGroupsSpaceFinder
}

// NewMockStagingSecurityGroupsSpaceFinder creates a new mock instance.
func NewMockStagingSecurityGroupsSpaceFinder(ctrl *gomock.Controller) *MockStagingSecurityGroupsSpaceFinder {
	mock := &MockStagingSecurityGroupsSpaceFinder{ctrl: ctrl}
	mock.recorder = &MockStagingSecurityGroupsSpaceFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagingSecurityGroupsSpaceFinder) EXPECT() *MockStagingSecurityGroupsSpaceFinderMockRecorder {
	return m.recorder
}

// FindStagingSecurityGroupsSpace mocks base method.
func (m *MockStagingSecurityGroupsSpaceFinder) FindStagingSecurityGroupsSpace(ctx context.Context, exec boil.ContextExecutor, stagingSecurityGroupsSpacesPK int, selectCols ...string) (*models.StagingSecurityGroupsSpace, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, stagingSecurityGroupsSpacesPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindStagingSecurityGroupsSpace", varargs...)
	ret0, _ := ret[0].(*models.StagingSecurityGroupsSpace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStagingSecurityGroupsSpace indicates an expected call of FindStagingSecurityGroupsSpace.
func (mr *MockStagingSecurityGroupsSpaceFinderMockRecorder) FindStagingSecurityGroupsSpace(ctx, exec, stagingSecurityGroupsSpacesPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, stagingSecurityGroupsSpacesPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStagingSecurityGroupsSpace", reflect.TypeOf((*MockStagingSecurityGroupsSpaceFinder)(nil).FindStagingSecurityGroupsSpace), varargs...)
}

// MockStagingSecurityGroupsSpaceInserter is a mock of StagingSecurityGroupsSpaceInserter interface.
type MockStagingSecurityGroupsSpaceInserter struct {
	ctrl     *gomock.Controller
	recorder *MockStagingSecurityGroupsSpaceInserterMockRecorder
}

// MockStagingSecurityGroupsSpaceInserterMockRecorder is the mock recorder for MockStagingSecurityGroupsSpaceInserter.
type MockStagingSecurityGroupsSpaceInserterMockRecorder struct {
	mock *MockStagingSecurityGroupsSpaceInserter
}

// NewMockStagingSecurityGroupsSpaceInserter creates a new mock instance.
func NewMockStagingSecurityGroupsSpaceInserter(ctrl *gomock.Controller) *MockStagingSecurityGroupsSpaceInserter {
	mock := &MockStagingSecurityGroupsSpaceInserter{ctrl: ctrl}
	mock.recorder = &MockStagingSecurityGroupsSpaceInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagingSecurityGroupsSpaceInserter) EXPECT() *MockStagingSecurityGroupsSpaceInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockStagingSecurityGroupsSpaceInserter) Insert(o *models.StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockStagingSecurityGroupsSpaceInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockStagingSecurityGroupsSpaceInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockStagingSecurityGroupsSpaceUpdater is a mock of StagingSecurityGroupsSpaceUpdater interface.
type MockStagingSecurityGroupsSpaceUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockStagingSecurityGroupsSpaceUpdaterMockRecorder
}

// MockStagingSecurityGroupsSpaceUpdaterMockRecorder is the mock recorder for MockStagingSecurityGroupsSpaceUpdater.
type MockStagingSecurityGroupsSpaceUpdaterMockRecorder struct {
	mock *MockStagingSecurityGroupsSpaceUpdater
}

// NewMockStagingSecurityGroupsSpaceUpdater creates a new mock instance.
func NewMockStagingSecurityGroupsSpaceUpdater(ctrl *gomock.Controller) *MockStagingSecurityGroupsSpaceUpdater {
	mock := &MockStagingSecurityGroupsSpaceUpdater{ctrl: ctrl}
	mock.recorder = &MockStagingSecurityGroupsSpaceUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagingSecurityGroupsSpaceUpdater) EXPECT() *MockStagingSecurityGroupsSpaceUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockStagingSecurityGroupsSpaceUpdater) Update(o *models.StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStagingSecurityGroupsSpaceUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStagingSecurityGroupsSpaceUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockStagingSecurityGroupsSpaceUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockStagingSecurityGroupsSpaceUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockStagingSecurityGroupsSpaceUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockStagingSecurityGroupsSpaceUpdater) UpdateAllSlice(o models.StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockStagingSecurityGroupsSpaceUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockStagingSecurityGroupsSpaceUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockStagingSecurityGroupsSpaceDeleter is a mock of StagingSecurityGroupsSpaceDeleter interface.
type MockStagingSecurityGroupsSpaceDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockStagingSecurityGroupsSpaceDeleterMockRecorder
}

// MockStagingSecurityGroupsSpaceDeleterMockRecorder is the mock recorder for MockStagingSecurityGroupsSpaceDeleter.
type MockStagingSecurityGroupsSpaceDeleterMockRecorder struct {
	mock *MockStagingSecurityGroupsSpaceDeleter
}

// NewMockStagingSecurityGroupsSpaceDeleter creates a new mock instance.
func NewMockStagingSecurityGroupsSpaceDeleter(ctrl *gomock.Controller) *MockStagingSecurityGroupsSpaceDeleter {
	mock := &MockStagingSecurityGroupsSpaceDeleter{ctrl: ctrl}
	mock.recorder = &MockStagingSecurityGroupsSpaceDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagingSecurityGroupsSpaceDeleter) EXPECT() *MockStagingSecurityGroupsSpaceDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockStagingSecurityGroupsSpaceDeleter) Delete(o *models.StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockStagingSecurityGroupsSpaceDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStagingSecurityGroupsSpaceDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockStagingSecurityGroupsSpaceDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockStagingSecurityGroupsSpaceDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockStagingSecurityGroupsSpaceDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockStagingSecurityGroupsSpaceDeleter) DeleteAllSlice(o models.StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockStagingSecurityGroupsSpaceDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockStagingSecurityGroupsSpaceDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockStagingSecurityGroupsSpaceReloader is a mock of StagingSecurityGroupsSpaceReloader interface.
type MockStagingSecurityGroupsSpaceReloader struct {
	ctrl     *gomock.Controller
	recorder *MockStagingSecurityGroupsSpaceReloaderMockRecorder
}

// MockStagingSecurityGroupsSpaceReloaderMockRecorder is the mock recorder for MockStagingSecurityGroupsSpaceReloader.
type MockStagingSecurityGroupsSpaceReloaderMockRecorder struct {
	mock *MockStagingSecurityGroupsSpaceReloader
}

// NewMockStagingSecurityGroupsSpaceReloader creates a new mock instance.
func NewMockStagingSecurityGroupsSpaceReloader(ctrl *gomock.Controller) *MockStagingSecurityGroupsSpaceReloader {
	mock := &MockStagingSecurityGroupsSpaceReloader{ctrl: ctrl}
	mock.recorder = &MockStagingSecurityGroupsSpaceReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStagingSecurityGroupsSpaceReloader) EXPECT() *MockStagingSecurityGroupsSpaceReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockStagingSecurityGroupsSpaceReloader) Reload(o *models.StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockStagingSecurityGroupsSpaceReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockStagingSecurityGroupsSpaceReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockStagingSecurityGroupsSpaceReloader) ReloadAll(o *models.StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockStagingSecurityGroupsSpaceReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockStagingSecurityGroupsSpaceReloader)(nil).ReloadAll), o, ctx, exec)
}
