//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_droplets.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
)

// MockDropletUpserter is a mock of DropletUpserter interface.
type MockDropletUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockDropletUpserterMockRecorder
}

// MockDropletUpserterMockRecorder is the mock recorder for MockDropletUpserter.
type MockDropletUpserterMockRecorder struct {
	mock *MockDropletUpserter
}

// NewMockDropletUpserter creates a new mock instance.
func NewMockDropletUpserter(ctrl *gomock.Controller) *MockDropletUpserter {
	mock := &MockDropletUpserter{ctrl: ctrl}
	mock.recorder = &MockDropletUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletUpserter) EXPECT() *MockDropletUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockDropletUpserter) Upsert(o *models.Droplet, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockDropletUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDropletUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockDropletFinisher is a mock of DropletFinisher interface.
type MockDropletFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockDropletFinisherMockRecorder
}

// MockDropletFinisherMockRecorder is the mock recorder for MockDropletFinisher.
type MockDropletFinisherMockRecorder struct {
	mock *MockDropletFinisher
}

// NewMockDropletFinisher creates a new mock instance.
func NewMockDropletFinisher(ctrl *gomock.Controller) *MockDropletFinisher {
	mock := &MockDropletFinisher{ctrl: ctrl}
	mock.recorder = &MockDropletFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletFinisher) EXPECT() *MockDropletFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockDropletFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.DropletSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.DropletSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockDropletFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockDropletFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockDropletFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDropletFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDropletFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockDropletFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDropletFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDropletFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockDropletFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.Droplet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.Droplet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockDropletFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockDropletFinisher)(nil).One), ctx, exec)
}

// MockDropletFinder is a mock of DropletFinder interface.
type MockDropletFinder struct {
	ctrl     *gomock.Controller
	recorder *MockDropletFinderMockRecorder
}

// MockDropletFinderMockRecorder is the mock recorder for MockDropletFinder.
type MockDropletFinderMockRecorder struct {
	mock *MockDropletFinder
}

// NewMockDropletFinder creates a new mock instance.
func NewMockDropletFinder(ctrl *gomock.Controller) *MockDropletFinder {
	mock := &MockDropletFinder{ctrl: ctrl}
	mock.recorder = &MockDropletFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletFinder) EXPECT() *MockDropletFinderMockRecorder {
	return m.recorder
}

// FindDroplet mocks base method.
func (m *MockDropletFinder) FindDroplet(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.Droplet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindDroplet", varargs...)
	ret0, _ := ret[0].(*models.Droplet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDroplet indicates an expected call of FindDroplet.
func (mr *MockDropletFinderMockRecorder) FindDroplet(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDroplet", reflect.TypeOf((*MockDropletFinder)(nil).FindDroplet), varargs...)
}

// MockDropletInserter is a mock of DropletInserter interface.
type MockDropletInserter struct {
	ctrl     *gomock.Controller
	recorder *MockDropletInserterMockRecorder
}

// MockDropletInserterMockRecorder is the mock recorder for MockDropletInserter.
type MockDropletInserterMockRecorder struct {
	mock *MockDropletInserter
}

// NewMockDropletInserter creates a new mock instance.
func NewMockDropletInserter(ctrl *gomock.Controller) *MockDropletInserter {
	mock := &MockDropletInserter{ctrl: ctrl}
	mock.recorder = &MockDropletInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletInserter) EXPECT() *MockDropletInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockDropletInserter) Insert(o *models.Droplet, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockDropletInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDropletInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockDropletUpdater is a mock of DropletUpdater interface.
type MockDropletUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDropletUpdaterMockRecorder
}

// MockDropletUpdaterMockRecorder is the mock recorder for MockDropletUpdater.
type MockDropletUpdaterMockRecorder struct {
	mock *MockDropletUpdater
}

// NewMockDropletUpdater creates a new mock instance.
func NewMockDropletUpdater(ctrl *gomock.Controller) *MockDropletUpdater {
	mock := &MockDropletUpdater{ctrl: ctrl}
	mock.recorder = &MockDropletUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletUpdater) EXPECT() *MockDropletUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockDropletUpdater) Update(o *models.Droplet, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDropletUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDropletUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockDropletUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockDropletUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockDropletUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockDropletUpdater) UpdateAllSlice(o models.DropletSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockDropletUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockDropletUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockDropletDeleter is a mock of DropletDeleter interface.
type MockDropletDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDropletDeleterMockRecorder
}

// MockDropletDeleterMockRecorder is the mock recorder for MockDropletDeleter.
type MockDropletDeleterMockRecorder struct {
	mock *MockDropletDeleter
}

// NewMockDropletDeleter creates a new mock instance.
func NewMockDropletDeleter(ctrl *gomock.Controller) *MockDropletDeleter {
	mock := &MockDropletDeleter{ctrl: ctrl}
	mock.recorder = &MockDropletDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletDeleter) EXPECT() *MockDropletDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDropletDeleter) Delete(o *models.Droplet, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDropletDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDropletDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockDropletDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockDropletDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockDropletDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockDropletDeleter) DeleteAllSlice(o models.DropletSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockDropletDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockDropletDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockDropletReloader is a mock of DropletReloader interface.
type MockDropletReloader struct {
	ctrl     *gomock.Controller
	recorder *MockDropletReloaderMockRecorder
}

// MockDropletReloaderMockRecorder is the mock recorder for MockDropletReloader.
type MockDropletReloaderMockRecorder struct {
	mock *MockDropletReloader
}

// NewMockDropletReloader creates a new mock instance.
func NewMockDropletReloader(ctrl *gomock.Controller) *MockDropletReloader {
	mock := &MockDropletReloader{ctrl: ctrl}
	mock.recorder = &MockDropletReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDropletReloader) EXPECT() *MockDropletReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockDropletReloader) Reload(o *models.Droplet, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockDropletReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockDropletReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockDropletReloader) ReloadAll(o *models.DropletSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockDropletReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockDropletReloader)(nil).ReloadAll), o, ctx, exec)
}
