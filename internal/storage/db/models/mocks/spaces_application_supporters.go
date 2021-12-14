//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_spaces_application_supporters.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockSpacesApplicationSupporterUpserter is a mock of SpacesApplicationSupporterUpserter interface.
type MockSpacesApplicationSupporterUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesApplicationSupporterUpserterMockRecorder
}

// MockSpacesApplicationSupporterUpserterMockRecorder is the mock recorder for MockSpacesApplicationSupporterUpserter.
type MockSpacesApplicationSupporterUpserterMockRecorder struct {
	mock *MockSpacesApplicationSupporterUpserter
}

// NewMockSpacesApplicationSupporterUpserter creates a new mock instance.
func NewMockSpacesApplicationSupporterUpserter(ctrl *gomock.Controller) *MockSpacesApplicationSupporterUpserter {
	mock := &MockSpacesApplicationSupporterUpserter{ctrl: ctrl}
	mock.recorder = &MockSpacesApplicationSupporterUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesApplicationSupporterUpserter) EXPECT() *MockSpacesApplicationSupporterUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockSpacesApplicationSupporterUpserter) Upsert(o *models.SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSpacesApplicationSupporterUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSpacesApplicationSupporterUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockSpacesApplicationSupporterFinisher is a mock of SpacesApplicationSupporterFinisher interface.
type MockSpacesApplicationSupporterFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesApplicationSupporterFinisherMockRecorder
}

// MockSpacesApplicationSupporterFinisherMockRecorder is the mock recorder for MockSpacesApplicationSupporterFinisher.
type MockSpacesApplicationSupporterFinisherMockRecorder struct {
	mock *MockSpacesApplicationSupporterFinisher
}

// NewMockSpacesApplicationSupporterFinisher creates a new mock instance.
func NewMockSpacesApplicationSupporterFinisher(ctrl *gomock.Controller) *MockSpacesApplicationSupporterFinisher {
	mock := &MockSpacesApplicationSupporterFinisher{ctrl: ctrl}
	mock.recorder = &MockSpacesApplicationSupporterFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesApplicationSupporterFinisher) EXPECT() *MockSpacesApplicationSupporterFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockSpacesApplicationSupporterFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.SpacesApplicationSupporterSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.SpacesApplicationSupporterSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockSpacesApplicationSupporterFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockSpacesApplicationSupporterFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockSpacesApplicationSupporterFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSpacesApplicationSupporterFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSpacesApplicationSupporterFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockSpacesApplicationSupporterFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSpacesApplicationSupporterFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSpacesApplicationSupporterFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockSpacesApplicationSupporterFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.SpacesApplicationSupporter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.SpacesApplicationSupporter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockSpacesApplicationSupporterFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockSpacesApplicationSupporterFinisher)(nil).One), ctx, exec)
}

// MockSpacesApplicationSupporterFinder is a mock of SpacesApplicationSupporterFinder interface.
type MockSpacesApplicationSupporterFinder struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesApplicationSupporterFinderMockRecorder
}

// MockSpacesApplicationSupporterFinderMockRecorder is the mock recorder for MockSpacesApplicationSupporterFinder.
type MockSpacesApplicationSupporterFinderMockRecorder struct {
	mock *MockSpacesApplicationSupporterFinder
}

// NewMockSpacesApplicationSupporterFinder creates a new mock instance.
func NewMockSpacesApplicationSupporterFinder(ctrl *gomock.Controller) *MockSpacesApplicationSupporterFinder {
	mock := &MockSpacesApplicationSupporterFinder{ctrl: ctrl}
	mock.recorder = &MockSpacesApplicationSupporterFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesApplicationSupporterFinder) EXPECT() *MockSpacesApplicationSupporterFinderMockRecorder {
	return m.recorder
}

// FindSpacesApplicationSupporter mocks base method.
func (m *MockSpacesApplicationSupporterFinder) FindSpacesApplicationSupporter(ctx context.Context, exec boil.ContextExecutor, spacesApplicationSupportersPK int, selectCols ...string) (*models.SpacesApplicationSupporter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, spacesApplicationSupportersPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSpacesApplicationSupporter", varargs...)
	ret0, _ := ret[0].(*models.SpacesApplicationSupporter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSpacesApplicationSupporter indicates an expected call of FindSpacesApplicationSupporter.
func (mr *MockSpacesApplicationSupporterFinderMockRecorder) FindSpacesApplicationSupporter(ctx, exec, spacesApplicationSupportersPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, spacesApplicationSupportersPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSpacesApplicationSupporter", reflect.TypeOf((*MockSpacesApplicationSupporterFinder)(nil).FindSpacesApplicationSupporter), varargs...)
}

// MockSpacesApplicationSupporterInserter is a mock of SpacesApplicationSupporterInserter interface.
type MockSpacesApplicationSupporterInserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesApplicationSupporterInserterMockRecorder
}

// MockSpacesApplicationSupporterInserterMockRecorder is the mock recorder for MockSpacesApplicationSupporterInserter.
type MockSpacesApplicationSupporterInserterMockRecorder struct {
	mock *MockSpacesApplicationSupporterInserter
}

// NewMockSpacesApplicationSupporterInserter creates a new mock instance.
func NewMockSpacesApplicationSupporterInserter(ctrl *gomock.Controller) *MockSpacesApplicationSupporterInserter {
	mock := &MockSpacesApplicationSupporterInserter{ctrl: ctrl}
	mock.recorder = &MockSpacesApplicationSupporterInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesApplicationSupporterInserter) EXPECT() *MockSpacesApplicationSupporterInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockSpacesApplicationSupporterInserter) Insert(o *models.SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSpacesApplicationSupporterInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSpacesApplicationSupporterInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockSpacesApplicationSupporterUpdater is a mock of SpacesApplicationSupporterUpdater interface.
type MockSpacesApplicationSupporterUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesApplicationSupporterUpdaterMockRecorder
}

// MockSpacesApplicationSupporterUpdaterMockRecorder is the mock recorder for MockSpacesApplicationSupporterUpdater.
type MockSpacesApplicationSupporterUpdaterMockRecorder struct {
	mock *MockSpacesApplicationSupporterUpdater
}

// NewMockSpacesApplicationSupporterUpdater creates a new mock instance.
func NewMockSpacesApplicationSupporterUpdater(ctrl *gomock.Controller) *MockSpacesApplicationSupporterUpdater {
	mock := &MockSpacesApplicationSupporterUpdater{ctrl: ctrl}
	mock.recorder = &MockSpacesApplicationSupporterUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesApplicationSupporterUpdater) EXPECT() *MockSpacesApplicationSupporterUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockSpacesApplicationSupporterUpdater) Update(o *models.SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSpacesApplicationSupporterUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSpacesApplicationSupporterUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockSpacesApplicationSupporterUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockSpacesApplicationSupporterUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockSpacesApplicationSupporterUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockSpacesApplicationSupporterUpdater) UpdateAllSlice(o models.SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockSpacesApplicationSupporterUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockSpacesApplicationSupporterUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockSpacesApplicationSupporterDeleter is a mock of SpacesApplicationSupporterDeleter interface.
type MockSpacesApplicationSupporterDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesApplicationSupporterDeleterMockRecorder
}

// MockSpacesApplicationSupporterDeleterMockRecorder is the mock recorder for MockSpacesApplicationSupporterDeleter.
type MockSpacesApplicationSupporterDeleterMockRecorder struct {
	mock *MockSpacesApplicationSupporterDeleter
}

// NewMockSpacesApplicationSupporterDeleter creates a new mock instance.
func NewMockSpacesApplicationSupporterDeleter(ctrl *gomock.Controller) *MockSpacesApplicationSupporterDeleter {
	mock := &MockSpacesApplicationSupporterDeleter{ctrl: ctrl}
	mock.recorder = &MockSpacesApplicationSupporterDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesApplicationSupporterDeleter) EXPECT() *MockSpacesApplicationSupporterDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSpacesApplicationSupporterDeleter) Delete(o *models.SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSpacesApplicationSupporterDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSpacesApplicationSupporterDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockSpacesApplicationSupporterDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockSpacesApplicationSupporterDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockSpacesApplicationSupporterDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockSpacesApplicationSupporterDeleter) DeleteAllSlice(o models.SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockSpacesApplicationSupporterDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockSpacesApplicationSupporterDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockSpacesApplicationSupporterReloader is a mock of SpacesApplicationSupporterReloader interface.
type MockSpacesApplicationSupporterReloader struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesApplicationSupporterReloaderMockRecorder
}

// MockSpacesApplicationSupporterReloaderMockRecorder is the mock recorder for MockSpacesApplicationSupporterReloader.
type MockSpacesApplicationSupporterReloaderMockRecorder struct {
	mock *MockSpacesApplicationSupporterReloader
}

// NewMockSpacesApplicationSupporterReloader creates a new mock instance.
func NewMockSpacesApplicationSupporterReloader(ctrl *gomock.Controller) *MockSpacesApplicationSupporterReloader {
	mock := &MockSpacesApplicationSupporterReloader{ctrl: ctrl}
	mock.recorder = &MockSpacesApplicationSupporterReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesApplicationSupporterReloader) EXPECT() *MockSpacesApplicationSupporterReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockSpacesApplicationSupporterReloader) Reload(o *models.SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockSpacesApplicationSupporterReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSpacesApplicationSupporterReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockSpacesApplicationSupporterReloader) ReloadAll(o *models.SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockSpacesApplicationSupporterReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockSpacesApplicationSupporterReloader)(nil).ReloadAll), o, ctx, exec)
}
