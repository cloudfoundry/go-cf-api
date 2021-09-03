// Code generated by MockGen. DO NOT EDIT.
// Source: psql_space_quota_definitions.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockSpaceQuotaDefinitionFinisher is a mock of SpaceQuotaDefinitionFinisher interface.
type MockSpaceQuotaDefinitionFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceQuotaDefinitionFinisherMockRecorder
}

// MockSpaceQuotaDefinitionFinisherMockRecorder is the mock recorder for MockSpaceQuotaDefinitionFinisher.
type MockSpaceQuotaDefinitionFinisherMockRecorder struct {
	mock *MockSpaceQuotaDefinitionFinisher
}

// NewMockSpaceQuotaDefinitionFinisher creates a new mock instance.
func NewMockSpaceQuotaDefinitionFinisher(ctrl *gomock.Controller) *MockSpaceQuotaDefinitionFinisher {
	mock := &MockSpaceQuotaDefinitionFinisher{ctrl: ctrl}
	mock.recorder = &MockSpaceQuotaDefinitionFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceQuotaDefinitionFinisher) EXPECT() *MockSpaceQuotaDefinitionFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockSpaceQuotaDefinitionFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.SpaceQuotaDefinitionSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.SpaceQuotaDefinitionSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockSpaceQuotaDefinitionFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockSpaceQuotaDefinitionFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockSpaceQuotaDefinitionFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSpaceQuotaDefinitionFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSpaceQuotaDefinitionFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockSpaceQuotaDefinitionFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSpaceQuotaDefinitionFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSpaceQuotaDefinitionFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockSpaceQuotaDefinitionFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.SpaceQuotaDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.SpaceQuotaDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockSpaceQuotaDefinitionFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockSpaceQuotaDefinitionFinisher)(nil).One), ctx, exec)
}

// MockSpaceQuotaDefinitionFinder is a mock of SpaceQuotaDefinitionFinder interface.
type MockSpaceQuotaDefinitionFinder struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceQuotaDefinitionFinderMockRecorder
}

// MockSpaceQuotaDefinitionFinderMockRecorder is the mock recorder for MockSpaceQuotaDefinitionFinder.
type MockSpaceQuotaDefinitionFinderMockRecorder struct {
	mock *MockSpaceQuotaDefinitionFinder
}

// NewMockSpaceQuotaDefinitionFinder creates a new mock instance.
func NewMockSpaceQuotaDefinitionFinder(ctrl *gomock.Controller) *MockSpaceQuotaDefinitionFinder {
	mock := &MockSpaceQuotaDefinitionFinder{ctrl: ctrl}
	mock.recorder = &MockSpaceQuotaDefinitionFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceQuotaDefinitionFinder) EXPECT() *MockSpaceQuotaDefinitionFinderMockRecorder {
	return m.recorder
}

// FindSpaceQuotaDefinition mocks base method.
func (m *MockSpaceQuotaDefinitionFinder) FindSpaceQuotaDefinition(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.SpaceQuotaDefinition, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSpaceQuotaDefinition", varargs...)
	ret0, _ := ret[0].(*models.SpaceQuotaDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSpaceQuotaDefinition indicates an expected call of FindSpaceQuotaDefinition.
func (mr *MockSpaceQuotaDefinitionFinderMockRecorder) FindSpaceQuotaDefinition(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSpaceQuotaDefinition", reflect.TypeOf((*MockSpaceQuotaDefinitionFinder)(nil).FindSpaceQuotaDefinition), varargs...)
}

// MockSpaceQuotaDefinitionInserter is a mock of SpaceQuotaDefinitionInserter interface.
type MockSpaceQuotaDefinitionInserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceQuotaDefinitionInserterMockRecorder
}

// MockSpaceQuotaDefinitionInserterMockRecorder is the mock recorder for MockSpaceQuotaDefinitionInserter.
type MockSpaceQuotaDefinitionInserterMockRecorder struct {
	mock *MockSpaceQuotaDefinitionInserter
}

// NewMockSpaceQuotaDefinitionInserter creates a new mock instance.
func NewMockSpaceQuotaDefinitionInserter(ctrl *gomock.Controller) *MockSpaceQuotaDefinitionInserter {
	mock := &MockSpaceQuotaDefinitionInserter{ctrl: ctrl}
	mock.recorder = &MockSpaceQuotaDefinitionInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceQuotaDefinitionInserter) EXPECT() *MockSpaceQuotaDefinitionInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockSpaceQuotaDefinitionInserter) Insert(o *models.SpaceQuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSpaceQuotaDefinitionInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSpaceQuotaDefinitionInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockSpaceQuotaDefinitionUpdater is a mock of SpaceQuotaDefinitionUpdater interface.
type MockSpaceQuotaDefinitionUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceQuotaDefinitionUpdaterMockRecorder
}

// MockSpaceQuotaDefinitionUpdaterMockRecorder is the mock recorder for MockSpaceQuotaDefinitionUpdater.
type MockSpaceQuotaDefinitionUpdaterMockRecorder struct {
	mock *MockSpaceQuotaDefinitionUpdater
}

// NewMockSpaceQuotaDefinitionUpdater creates a new mock instance.
func NewMockSpaceQuotaDefinitionUpdater(ctrl *gomock.Controller) *MockSpaceQuotaDefinitionUpdater {
	mock := &MockSpaceQuotaDefinitionUpdater{ctrl: ctrl}
	mock.recorder = &MockSpaceQuotaDefinitionUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceQuotaDefinitionUpdater) EXPECT() *MockSpaceQuotaDefinitionUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockSpaceQuotaDefinitionUpdater) Update(o *models.SpaceQuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSpaceQuotaDefinitionUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSpaceQuotaDefinitionUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockSpaceQuotaDefinitionUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockSpaceQuotaDefinitionUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockSpaceQuotaDefinitionUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockSpaceQuotaDefinitionUpdater) UpdateAllSlice(o models.SpaceQuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockSpaceQuotaDefinitionUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockSpaceQuotaDefinitionUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockSpaceQuotaDefinitionUpserter is a mock of SpaceQuotaDefinitionUpserter interface.
type MockSpaceQuotaDefinitionUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceQuotaDefinitionUpserterMockRecorder
}

// MockSpaceQuotaDefinitionUpserterMockRecorder is the mock recorder for MockSpaceQuotaDefinitionUpserter.
type MockSpaceQuotaDefinitionUpserterMockRecorder struct {
	mock *MockSpaceQuotaDefinitionUpserter
}

// NewMockSpaceQuotaDefinitionUpserter creates a new mock instance.
func NewMockSpaceQuotaDefinitionUpserter(ctrl *gomock.Controller) *MockSpaceQuotaDefinitionUpserter {
	mock := &MockSpaceQuotaDefinitionUpserter{ctrl: ctrl}
	mock.recorder = &MockSpaceQuotaDefinitionUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceQuotaDefinitionUpserter) EXPECT() *MockSpaceQuotaDefinitionUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockSpaceQuotaDefinitionUpserter) Upsert(o *models.SpaceQuotaDefinition, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSpaceQuotaDefinitionUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSpaceQuotaDefinitionUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockSpaceQuotaDefinitionDeleter is a mock of SpaceQuotaDefinitionDeleter interface.
type MockSpaceQuotaDefinitionDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceQuotaDefinitionDeleterMockRecorder
}

// MockSpaceQuotaDefinitionDeleterMockRecorder is the mock recorder for MockSpaceQuotaDefinitionDeleter.
type MockSpaceQuotaDefinitionDeleterMockRecorder struct {
	mock *MockSpaceQuotaDefinitionDeleter
}

// NewMockSpaceQuotaDefinitionDeleter creates a new mock instance.
func NewMockSpaceQuotaDefinitionDeleter(ctrl *gomock.Controller) *MockSpaceQuotaDefinitionDeleter {
	mock := &MockSpaceQuotaDefinitionDeleter{ctrl: ctrl}
	mock.recorder = &MockSpaceQuotaDefinitionDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceQuotaDefinitionDeleter) EXPECT() *MockSpaceQuotaDefinitionDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSpaceQuotaDefinitionDeleter) Delete(o *models.SpaceQuotaDefinition, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSpaceQuotaDefinitionDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSpaceQuotaDefinitionDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockSpaceQuotaDefinitionDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockSpaceQuotaDefinitionDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockSpaceQuotaDefinitionDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockSpaceQuotaDefinitionDeleter) DeleteAllSlice(o models.SpaceQuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockSpaceQuotaDefinitionDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockSpaceQuotaDefinitionDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockSpaceQuotaDefinitionReloader is a mock of SpaceQuotaDefinitionReloader interface.
type MockSpaceQuotaDefinitionReloader struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceQuotaDefinitionReloaderMockRecorder
}

// MockSpaceQuotaDefinitionReloaderMockRecorder is the mock recorder for MockSpaceQuotaDefinitionReloader.
type MockSpaceQuotaDefinitionReloaderMockRecorder struct {
	mock *MockSpaceQuotaDefinitionReloader
}

// NewMockSpaceQuotaDefinitionReloader creates a new mock instance.
func NewMockSpaceQuotaDefinitionReloader(ctrl *gomock.Controller) *MockSpaceQuotaDefinitionReloader {
	mock := &MockSpaceQuotaDefinitionReloader{ctrl: ctrl}
	mock.recorder = &MockSpaceQuotaDefinitionReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceQuotaDefinitionReloader) EXPECT() *MockSpaceQuotaDefinitionReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockSpaceQuotaDefinitionReloader) Reload(o *models.SpaceQuotaDefinition, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockSpaceQuotaDefinitionReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSpaceQuotaDefinitionReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockSpaceQuotaDefinitionReloader) ReloadAll(o *models.SpaceQuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockSpaceQuotaDefinitionReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockSpaceQuotaDefinitionReloader)(nil).ReloadAll), o, ctx, exec)
}
