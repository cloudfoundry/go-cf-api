//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_organization_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockOrganizationAnnotationUpserter is a mock of OrganizationAnnotationUpserter interface.
type MockOrganizationAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAnnotationUpserterMockRecorder
}

// MockOrganizationAnnotationUpserterMockRecorder is the mock recorder for MockOrganizationAnnotationUpserter.
type MockOrganizationAnnotationUpserterMockRecorder struct {
	mock *MockOrganizationAnnotationUpserter
}

// NewMockOrganizationAnnotationUpserter creates a new mock instance.
func NewMockOrganizationAnnotationUpserter(ctrl *gomock.Controller) *MockOrganizationAnnotationUpserter {
	mock := &MockOrganizationAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAnnotationUpserter) EXPECT() *MockOrganizationAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockOrganizationAnnotationUpserter) Upsert(o *models.OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOrganizationAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOrganizationAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockOrganizationAnnotationFinisher is a mock of OrganizationAnnotationFinisher interface.
type MockOrganizationAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAnnotationFinisherMockRecorder
}

// MockOrganizationAnnotationFinisherMockRecorder is the mock recorder for MockOrganizationAnnotationFinisher.
type MockOrganizationAnnotationFinisherMockRecorder struct {
	mock *MockOrganizationAnnotationFinisher
}

// NewMockOrganizationAnnotationFinisher creates a new mock instance.
func NewMockOrganizationAnnotationFinisher(ctrl *gomock.Controller) *MockOrganizationAnnotationFinisher {
	mock := &MockOrganizationAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockOrganizationAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAnnotationFinisher) EXPECT() *MockOrganizationAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrganizationAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrganizationAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrganizationAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrganizationAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrganizationAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrganizationAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrganizationAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrganizationAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrganizationAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrganizationAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrganizationAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrganizationAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.OrganizationAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.OrganizationAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrganizationAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrganizationAnnotationFinisher)(nil).One), ctx, exec)
}

// MockOrganizationAnnotationFinder is a mock of OrganizationAnnotationFinder interface.
type MockOrganizationAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAnnotationFinderMockRecorder
}

// MockOrganizationAnnotationFinderMockRecorder is the mock recorder for MockOrganizationAnnotationFinder.
type MockOrganizationAnnotationFinderMockRecorder struct {
	mock *MockOrganizationAnnotationFinder
}

// NewMockOrganizationAnnotationFinder creates a new mock instance.
func NewMockOrganizationAnnotationFinder(ctrl *gomock.Controller) *MockOrganizationAnnotationFinder {
	mock := &MockOrganizationAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockOrganizationAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAnnotationFinder) EXPECT() *MockOrganizationAnnotationFinderMockRecorder {
	return m.recorder
}

// FindOrganizationAnnotation mocks base method.
func (m *MockOrganizationAnnotationFinder) FindOrganizationAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.OrganizationAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganizationAnnotation", varargs...)
	ret0, _ := ret[0].(*models.OrganizationAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationAnnotation indicates an expected call of FindOrganizationAnnotation.
func (mr *MockOrganizationAnnotationFinderMockRecorder) FindOrganizationAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationAnnotation", reflect.TypeOf((*MockOrganizationAnnotationFinder)(nil).FindOrganizationAnnotation), varargs...)
}

// MockOrganizationAnnotationInserter is a mock of OrganizationAnnotationInserter interface.
type MockOrganizationAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAnnotationInserterMockRecorder
}

// MockOrganizationAnnotationInserterMockRecorder is the mock recorder for MockOrganizationAnnotationInserter.
type MockOrganizationAnnotationInserterMockRecorder struct {
	mock *MockOrganizationAnnotationInserter
}

// NewMockOrganizationAnnotationInserter creates a new mock instance.
func NewMockOrganizationAnnotationInserter(ctrl *gomock.Controller) *MockOrganizationAnnotationInserter {
	mock := &MockOrganizationAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAnnotationInserter) EXPECT() *MockOrganizationAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrganizationAnnotationInserter) Insert(o *models.OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrganizationAnnotationUpdater is a mock of OrganizationAnnotationUpdater interface.
type MockOrganizationAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAnnotationUpdaterMockRecorder
}

// MockOrganizationAnnotationUpdaterMockRecorder is the mock recorder for MockOrganizationAnnotationUpdater.
type MockOrganizationAnnotationUpdaterMockRecorder struct {
	mock *MockOrganizationAnnotationUpdater
}

// NewMockOrganizationAnnotationUpdater creates a new mock instance.
func NewMockOrganizationAnnotationUpdater(ctrl *gomock.Controller) *MockOrganizationAnnotationUpdater {
	mock := &MockOrganizationAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAnnotationUpdater) EXPECT() *MockOrganizationAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrganizationAnnotationUpdater) Update(o *models.OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrganizationAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrganizationAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrganizationAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrganizationAnnotationUpdater) UpdateAllSlice(o models.OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrganizationAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrganizationAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrganizationAnnotationDeleter is a mock of OrganizationAnnotationDeleter interface.
type MockOrganizationAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAnnotationDeleterMockRecorder
}

// MockOrganizationAnnotationDeleterMockRecorder is the mock recorder for MockOrganizationAnnotationDeleter.
type MockOrganizationAnnotationDeleterMockRecorder struct {
	mock *MockOrganizationAnnotationDeleter
}

// NewMockOrganizationAnnotationDeleter creates a new mock instance.
func NewMockOrganizationAnnotationDeleter(ctrl *gomock.Controller) *MockOrganizationAnnotationDeleter {
	mock := &MockOrganizationAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAnnotationDeleter) EXPECT() *MockOrganizationAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationAnnotationDeleter) Delete(o *models.OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrganizationAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrganizationAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrganizationAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrganizationAnnotationDeleter) DeleteAllSlice(o models.OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrganizationAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrganizationAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrganizationAnnotationReloader is a mock of OrganizationAnnotationReloader interface.
type MockOrganizationAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAnnotationReloaderMockRecorder
}

// MockOrganizationAnnotationReloaderMockRecorder is the mock recorder for MockOrganizationAnnotationReloader.
type MockOrganizationAnnotationReloaderMockRecorder struct {
	mock *MockOrganizationAnnotationReloader
}

// NewMockOrganizationAnnotationReloader creates a new mock instance.
func NewMockOrganizationAnnotationReloader(ctrl *gomock.Controller) *MockOrganizationAnnotationReloader {
	mock := &MockOrganizationAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockOrganizationAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAnnotationReloader) EXPECT() *MockOrganizationAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrganizationAnnotationReloader) Reload(o *models.OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrganizationAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrganizationAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrganizationAnnotationReloader) ReloadAll(o *models.OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrganizationAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrganizationAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
