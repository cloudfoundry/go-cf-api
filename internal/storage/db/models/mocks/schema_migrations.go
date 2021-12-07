//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_schema_migrations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockSchemaMigrationUpserter is a mock of SchemaMigrationUpserter interface.
type MockSchemaMigrationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMigrationUpserterMockRecorder
}

// MockSchemaMigrationUpserterMockRecorder is the mock recorder for MockSchemaMigrationUpserter.
type MockSchemaMigrationUpserterMockRecorder struct {
	mock *MockSchemaMigrationUpserter
}

// NewMockSchemaMigrationUpserter creates a new mock instance.
func NewMockSchemaMigrationUpserter(ctrl *gomock.Controller) *MockSchemaMigrationUpserter {
	mock := &MockSchemaMigrationUpserter{ctrl: ctrl}
	mock.recorder = &MockSchemaMigrationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaMigrationUpserter) EXPECT() *MockSchemaMigrationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockSchemaMigrationUpserter) Upsert(o *models.SchemaMigration, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSchemaMigrationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSchemaMigrationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockSchemaMigrationFinisher is a mock of SchemaMigrationFinisher interface.
type MockSchemaMigrationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMigrationFinisherMockRecorder
}

// MockSchemaMigrationFinisherMockRecorder is the mock recorder for MockSchemaMigrationFinisher.
type MockSchemaMigrationFinisherMockRecorder struct {
	mock *MockSchemaMigrationFinisher
}

// NewMockSchemaMigrationFinisher creates a new mock instance.
func NewMockSchemaMigrationFinisher(ctrl *gomock.Controller) *MockSchemaMigrationFinisher {
	mock := &MockSchemaMigrationFinisher{ctrl: ctrl}
	mock.recorder = &MockSchemaMigrationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaMigrationFinisher) EXPECT() *MockSchemaMigrationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockSchemaMigrationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.SchemaMigrationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.SchemaMigrationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockSchemaMigrationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockSchemaMigrationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockSchemaMigrationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSchemaMigrationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSchemaMigrationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockSchemaMigrationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSchemaMigrationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSchemaMigrationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockSchemaMigrationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.SchemaMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.SchemaMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockSchemaMigrationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockSchemaMigrationFinisher)(nil).One), ctx, exec)
}

// MockSchemaMigrationFinder is a mock of SchemaMigrationFinder interface.
type MockSchemaMigrationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMigrationFinderMockRecorder
}

// MockSchemaMigrationFinderMockRecorder is the mock recorder for MockSchemaMigrationFinder.
type MockSchemaMigrationFinderMockRecorder struct {
	mock *MockSchemaMigrationFinder
}

// NewMockSchemaMigrationFinder creates a new mock instance.
func NewMockSchemaMigrationFinder(ctrl *gomock.Controller) *MockSchemaMigrationFinder {
	mock := &MockSchemaMigrationFinder{ctrl: ctrl}
	mock.recorder = &MockSchemaMigrationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaMigrationFinder) EXPECT() *MockSchemaMigrationFinderMockRecorder {
	return m.recorder
}

// FindSchemaMigration mocks base method.
func (m *MockSchemaMigrationFinder) FindSchemaMigration(ctx context.Context, exec boil.ContextExecutor, filename string, selectCols ...string) (*models.SchemaMigration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, filename}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSchemaMigration", varargs...)
	ret0, _ := ret[0].(*models.SchemaMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSchemaMigration indicates an expected call of FindSchemaMigration.
func (mr *MockSchemaMigrationFinderMockRecorder) FindSchemaMigration(ctx, exec, filename interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, filename}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSchemaMigration", reflect.TypeOf((*MockSchemaMigrationFinder)(nil).FindSchemaMigration), varargs...)
}

// MockSchemaMigrationInserter is a mock of SchemaMigrationInserter interface.
type MockSchemaMigrationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMigrationInserterMockRecorder
}

// MockSchemaMigrationInserterMockRecorder is the mock recorder for MockSchemaMigrationInserter.
type MockSchemaMigrationInserterMockRecorder struct {
	mock *MockSchemaMigrationInserter
}

// NewMockSchemaMigrationInserter creates a new mock instance.
func NewMockSchemaMigrationInserter(ctrl *gomock.Controller) *MockSchemaMigrationInserter {
	mock := &MockSchemaMigrationInserter{ctrl: ctrl}
	mock.recorder = &MockSchemaMigrationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaMigrationInserter) EXPECT() *MockSchemaMigrationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockSchemaMigrationInserter) Insert(o *models.SchemaMigration, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSchemaMigrationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSchemaMigrationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockSchemaMigrationUpdater is a mock of SchemaMigrationUpdater interface.
type MockSchemaMigrationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMigrationUpdaterMockRecorder
}

// MockSchemaMigrationUpdaterMockRecorder is the mock recorder for MockSchemaMigrationUpdater.
type MockSchemaMigrationUpdaterMockRecorder struct {
	mock *MockSchemaMigrationUpdater
}

// NewMockSchemaMigrationUpdater creates a new mock instance.
func NewMockSchemaMigrationUpdater(ctrl *gomock.Controller) *MockSchemaMigrationUpdater {
	mock := &MockSchemaMigrationUpdater{ctrl: ctrl}
	mock.recorder = &MockSchemaMigrationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaMigrationUpdater) EXPECT() *MockSchemaMigrationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockSchemaMigrationUpdater) Update(o *models.SchemaMigration, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSchemaMigrationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSchemaMigrationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockSchemaMigrationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockSchemaMigrationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockSchemaMigrationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockSchemaMigrationUpdater) UpdateAllSlice(o models.SchemaMigrationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockSchemaMigrationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockSchemaMigrationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockSchemaMigrationDeleter is a mock of SchemaMigrationDeleter interface.
type MockSchemaMigrationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMigrationDeleterMockRecorder
}

// MockSchemaMigrationDeleterMockRecorder is the mock recorder for MockSchemaMigrationDeleter.
type MockSchemaMigrationDeleterMockRecorder struct {
	mock *MockSchemaMigrationDeleter
}

// NewMockSchemaMigrationDeleter creates a new mock instance.
func NewMockSchemaMigrationDeleter(ctrl *gomock.Controller) *MockSchemaMigrationDeleter {
	mock := &MockSchemaMigrationDeleter{ctrl: ctrl}
	mock.recorder = &MockSchemaMigrationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaMigrationDeleter) EXPECT() *MockSchemaMigrationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSchemaMigrationDeleter) Delete(o *models.SchemaMigration, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSchemaMigrationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSchemaMigrationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockSchemaMigrationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockSchemaMigrationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockSchemaMigrationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockSchemaMigrationDeleter) DeleteAllSlice(o models.SchemaMigrationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockSchemaMigrationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockSchemaMigrationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockSchemaMigrationReloader is a mock of SchemaMigrationReloader interface.
type MockSchemaMigrationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaMigrationReloaderMockRecorder
}

// MockSchemaMigrationReloaderMockRecorder is the mock recorder for MockSchemaMigrationReloader.
type MockSchemaMigrationReloaderMockRecorder struct {
	mock *MockSchemaMigrationReloader
}

// NewMockSchemaMigrationReloader creates a new mock instance.
func NewMockSchemaMigrationReloader(ctrl *gomock.Controller) *MockSchemaMigrationReloader {
	mock := &MockSchemaMigrationReloader{ctrl: ctrl}
	mock.recorder = &MockSchemaMigrationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchemaMigrationReloader) EXPECT() *MockSchemaMigrationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockSchemaMigrationReloader) Reload(o *models.SchemaMigration, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockSchemaMigrationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSchemaMigrationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockSchemaMigrationReloader) ReloadAll(o *models.SchemaMigrationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockSchemaMigrationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockSchemaMigrationReloader)(nil).ReloadAll), o, ctx, exec)
}
