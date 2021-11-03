// +build unit

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_quota_definitions.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockQuotaDefinitionUpserter is a mock of QuotaDefinitionUpserter interface.
type MockQuotaDefinitionUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaDefinitionUpserterMockRecorder
}

// MockQuotaDefinitionUpserterMockRecorder is the mock recorder for MockQuotaDefinitionUpserter.
type MockQuotaDefinitionUpserterMockRecorder struct {
	mock *MockQuotaDefinitionUpserter
}

// NewMockQuotaDefinitionUpserter creates a new mock instance.
func NewMockQuotaDefinitionUpserter(ctrl *gomock.Controller) *MockQuotaDefinitionUpserter {
	mock := &MockQuotaDefinitionUpserter{ctrl: ctrl}
	mock.recorder = &MockQuotaDefinitionUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaDefinitionUpserter) EXPECT() *MockQuotaDefinitionUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockQuotaDefinitionUpserter) Upsert(o *models.QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockQuotaDefinitionUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockQuotaDefinitionUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockQuotaDefinitionFinisher is a mock of QuotaDefinitionFinisher interface.
type MockQuotaDefinitionFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaDefinitionFinisherMockRecorder
}

// MockQuotaDefinitionFinisherMockRecorder is the mock recorder for MockQuotaDefinitionFinisher.
type MockQuotaDefinitionFinisherMockRecorder struct {
	mock *MockQuotaDefinitionFinisher
}

// NewMockQuotaDefinitionFinisher creates a new mock instance.
func NewMockQuotaDefinitionFinisher(ctrl *gomock.Controller) *MockQuotaDefinitionFinisher {
	mock := &MockQuotaDefinitionFinisher{ctrl: ctrl}
	mock.recorder = &MockQuotaDefinitionFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaDefinitionFinisher) EXPECT() *MockQuotaDefinitionFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockQuotaDefinitionFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.QuotaDefinitionSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.QuotaDefinitionSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockQuotaDefinitionFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockQuotaDefinitionFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockQuotaDefinitionFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockQuotaDefinitionFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockQuotaDefinitionFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockQuotaDefinitionFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockQuotaDefinitionFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockQuotaDefinitionFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockQuotaDefinitionFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.QuotaDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.QuotaDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockQuotaDefinitionFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockQuotaDefinitionFinisher)(nil).One), ctx, exec)
}

// MockQuotaDefinitionFinder is a mock of QuotaDefinitionFinder interface.
type MockQuotaDefinitionFinder struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaDefinitionFinderMockRecorder
}

// MockQuotaDefinitionFinderMockRecorder is the mock recorder for MockQuotaDefinitionFinder.
type MockQuotaDefinitionFinderMockRecorder struct {
	mock *MockQuotaDefinitionFinder
}

// NewMockQuotaDefinitionFinder creates a new mock instance.
func NewMockQuotaDefinitionFinder(ctrl *gomock.Controller) *MockQuotaDefinitionFinder {
	mock := &MockQuotaDefinitionFinder{ctrl: ctrl}
	mock.recorder = &MockQuotaDefinitionFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaDefinitionFinder) EXPECT() *MockQuotaDefinitionFinderMockRecorder {
	return m.recorder
}

// FindQuotaDefinition mocks base method.
func (m *MockQuotaDefinitionFinder) FindQuotaDefinition(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.QuotaDefinition, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindQuotaDefinition", varargs...)
	ret0, _ := ret[0].(*models.QuotaDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindQuotaDefinition indicates an expected call of FindQuotaDefinition.
func (mr *MockQuotaDefinitionFinderMockRecorder) FindQuotaDefinition(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindQuotaDefinition", reflect.TypeOf((*MockQuotaDefinitionFinder)(nil).FindQuotaDefinition), varargs...)
}

// MockQuotaDefinitionInserter is a mock of QuotaDefinitionInserter interface.
type MockQuotaDefinitionInserter struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaDefinitionInserterMockRecorder
}

// MockQuotaDefinitionInserterMockRecorder is the mock recorder for MockQuotaDefinitionInserter.
type MockQuotaDefinitionInserterMockRecorder struct {
	mock *MockQuotaDefinitionInserter
}

// NewMockQuotaDefinitionInserter creates a new mock instance.
func NewMockQuotaDefinitionInserter(ctrl *gomock.Controller) *MockQuotaDefinitionInserter {
	mock := &MockQuotaDefinitionInserter{ctrl: ctrl}
	mock.recorder = &MockQuotaDefinitionInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaDefinitionInserter) EXPECT() *MockQuotaDefinitionInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockQuotaDefinitionInserter) Insert(o *models.QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockQuotaDefinitionInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockQuotaDefinitionInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockQuotaDefinitionUpdater is a mock of QuotaDefinitionUpdater interface.
type MockQuotaDefinitionUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaDefinitionUpdaterMockRecorder
}

// MockQuotaDefinitionUpdaterMockRecorder is the mock recorder for MockQuotaDefinitionUpdater.
type MockQuotaDefinitionUpdaterMockRecorder struct {
	mock *MockQuotaDefinitionUpdater
}

// NewMockQuotaDefinitionUpdater creates a new mock instance.
func NewMockQuotaDefinitionUpdater(ctrl *gomock.Controller) *MockQuotaDefinitionUpdater {
	mock := &MockQuotaDefinitionUpdater{ctrl: ctrl}
	mock.recorder = &MockQuotaDefinitionUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaDefinitionUpdater) EXPECT() *MockQuotaDefinitionUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockQuotaDefinitionUpdater) Update(o *models.QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockQuotaDefinitionUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockQuotaDefinitionUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockQuotaDefinitionUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockQuotaDefinitionUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockQuotaDefinitionUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockQuotaDefinitionUpdater) UpdateAllSlice(o models.QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockQuotaDefinitionUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockQuotaDefinitionUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockQuotaDefinitionDeleter is a mock of QuotaDefinitionDeleter interface.
type MockQuotaDefinitionDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaDefinitionDeleterMockRecorder
}

// MockQuotaDefinitionDeleterMockRecorder is the mock recorder for MockQuotaDefinitionDeleter.
type MockQuotaDefinitionDeleterMockRecorder struct {
	mock *MockQuotaDefinitionDeleter
}

// NewMockQuotaDefinitionDeleter creates a new mock instance.
func NewMockQuotaDefinitionDeleter(ctrl *gomock.Controller) *MockQuotaDefinitionDeleter {
	mock := &MockQuotaDefinitionDeleter{ctrl: ctrl}
	mock.recorder = &MockQuotaDefinitionDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaDefinitionDeleter) EXPECT() *MockQuotaDefinitionDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockQuotaDefinitionDeleter) Delete(o *models.QuotaDefinition, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockQuotaDefinitionDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockQuotaDefinitionDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockQuotaDefinitionDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockQuotaDefinitionDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockQuotaDefinitionDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockQuotaDefinitionDeleter) DeleteAllSlice(o models.QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockQuotaDefinitionDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockQuotaDefinitionDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockQuotaDefinitionReloader is a mock of QuotaDefinitionReloader interface.
type MockQuotaDefinitionReloader struct {
	ctrl     *gomock.Controller
	recorder *MockQuotaDefinitionReloaderMockRecorder
}

// MockQuotaDefinitionReloaderMockRecorder is the mock recorder for MockQuotaDefinitionReloader.
type MockQuotaDefinitionReloaderMockRecorder struct {
	mock *MockQuotaDefinitionReloader
}

// NewMockQuotaDefinitionReloader creates a new mock instance.
func NewMockQuotaDefinitionReloader(ctrl *gomock.Controller) *MockQuotaDefinitionReloader {
	mock := &MockQuotaDefinitionReloader{ctrl: ctrl}
	mock.recorder = &MockQuotaDefinitionReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuotaDefinitionReloader) EXPECT() *MockQuotaDefinitionReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockQuotaDefinitionReloader) Reload(o *models.QuotaDefinition, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockQuotaDefinitionReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockQuotaDefinitionReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockQuotaDefinitionReloader) ReloadAll(o *models.QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockQuotaDefinitionReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockQuotaDefinitionReloader)(nil).ReloadAll), o, ctx, exec)
}
