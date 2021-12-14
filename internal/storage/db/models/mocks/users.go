//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_users.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockUserUpserter is a mock of UserUpserter interface.
type MockUserUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockUserUpserterMockRecorder
}

// MockUserUpserterMockRecorder is the mock recorder for MockUserUpserter.
type MockUserUpserterMockRecorder struct {
	mock *MockUserUpserter
}

// NewMockUserUpserter creates a new mock instance.
func NewMockUserUpserter(ctrl *gomock.Controller) *MockUserUpserter {
	mock := &MockUserUpserter{ctrl: ctrl}
	mock.recorder = &MockUserUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUpserter) EXPECT() *MockUserUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockUserUpserter) Upsert(o *models.User, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockUserUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockUserUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockUserFinisher is a mock of UserFinisher interface.
type MockUserFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockUserFinisherMockRecorder
}

// MockUserFinisherMockRecorder is the mock recorder for MockUserFinisher.
type MockUserFinisherMockRecorder struct {
	mock *MockUserFinisher
}

// NewMockUserFinisher creates a new mock instance.
func NewMockUserFinisher(ctrl *gomock.Controller) *MockUserFinisher {
	mock := &MockUserFinisher{ctrl: ctrl}
	mock.recorder = &MockUserFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserFinisher) EXPECT() *MockUserFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockUserFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockUserFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockUserFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockUserFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockUserFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUserFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockUserFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockUserFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockUserFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockUserFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockUserFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockUserFinisher)(nil).One), ctx, exec)
}

// MockUserFinder is a mock of UserFinder interface.
type MockUserFinder struct {
	ctrl     *gomock.Controller
	recorder *MockUserFinderMockRecorder
}

// MockUserFinderMockRecorder is the mock recorder for MockUserFinder.
type MockUserFinderMockRecorder struct {
	mock *MockUserFinder
}

// NewMockUserFinder creates a new mock instance.
func NewMockUserFinder(ctrl *gomock.Controller) *MockUserFinder {
	mock := &MockUserFinder{ctrl: ctrl}
	mock.recorder = &MockUserFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserFinder) EXPECT() *MockUserFinderMockRecorder {
	return m.recorder
}

// FindUser mocks base method.
func (m *MockUserFinder) FindUser(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindUser", varargs...)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUserFinderMockRecorder) FindUser(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUserFinder)(nil).FindUser), varargs...)
}

// MockUserInserter is a mock of UserInserter interface.
type MockUserInserter struct {
	ctrl     *gomock.Controller
	recorder *MockUserInserterMockRecorder
}

// MockUserInserterMockRecorder is the mock recorder for MockUserInserter.
type MockUserInserterMockRecorder struct {
	mock *MockUserInserter
}

// NewMockUserInserter creates a new mock instance.
func NewMockUserInserter(ctrl *gomock.Controller) *MockUserInserter {
	mock := &MockUserInserter{ctrl: ctrl}
	mock.recorder = &MockUserInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInserter) EXPECT() *MockUserInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockUserInserter) Insert(o *models.User, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockUserInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockUserUpdater is a mock of UserUpdater interface.
type MockUserUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockUserUpdaterMockRecorder
}

// MockUserUpdaterMockRecorder is the mock recorder for MockUserUpdater.
type MockUserUpdaterMockRecorder struct {
	mock *MockUserUpdater
}

// NewMockUserUpdater creates a new mock instance.
func NewMockUserUpdater(ctrl *gomock.Controller) *MockUserUpdater {
	mock := &MockUserUpdater{ctrl: ctrl}
	mock.recorder = &MockUserUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUpdater) EXPECT() *MockUserUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockUserUpdater) Update(o *models.User, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockUserUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockUserUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockUserUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockUserUpdater) UpdateAllSlice(o models.UserSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockUserUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockUserUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockUserDeleter is a mock of UserDeleter interface.
type MockUserDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockUserDeleterMockRecorder
}

// MockUserDeleterMockRecorder is the mock recorder for MockUserDeleter.
type MockUserDeleterMockRecorder struct {
	mock *MockUserDeleter
}

// NewMockUserDeleter creates a new mock instance.
func NewMockUserDeleter(ctrl *gomock.Controller) *MockUserDeleter {
	mock := &MockUserDeleter{ctrl: ctrl}
	mock.recorder = &MockUserDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDeleter) EXPECT() *MockUserDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserDeleter) Delete(o *models.User, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockUserDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockUserDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockUserDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockUserDeleter) DeleteAllSlice(o models.UserSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockUserDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockUserDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockUserReloader is a mock of UserReloader interface.
type MockUserReloader struct {
	ctrl     *gomock.Controller
	recorder *MockUserReloaderMockRecorder
}

// MockUserReloaderMockRecorder is the mock recorder for MockUserReloader.
type MockUserReloaderMockRecorder struct {
	mock *MockUserReloader
}

// NewMockUserReloader creates a new mock instance.
func NewMockUserReloader(ctrl *gomock.Controller) *MockUserReloader {
	mock := &MockUserReloader{ctrl: ctrl}
	mock.recorder = &MockUserReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserReloader) EXPECT() *MockUserReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockUserReloader) Reload(o *models.User, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockUserReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockUserReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockUserReloader) ReloadAll(o *models.UserSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockUserReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockUserReloader)(nil).ReloadAll), o, ctx, exec)
}
