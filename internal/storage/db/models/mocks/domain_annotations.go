//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_domain_annotations.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

// MockDomainAnnotationUpserter is a mock of DomainAnnotationUpserter interface.
type MockDomainAnnotationUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockDomainAnnotationUpserterMockRecorder
}

// MockDomainAnnotationUpserterMockRecorder is the mock recorder for MockDomainAnnotationUpserter.
type MockDomainAnnotationUpserterMockRecorder struct {
	mock *MockDomainAnnotationUpserter
}

// NewMockDomainAnnotationUpserter creates a new mock instance.
func NewMockDomainAnnotationUpserter(ctrl *gomock.Controller) *MockDomainAnnotationUpserter {
	mock := &MockDomainAnnotationUpserter{ctrl: ctrl}
	mock.recorder = &MockDomainAnnotationUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainAnnotationUpserter) EXPECT() *MockDomainAnnotationUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockDomainAnnotationUpserter) Upsert(o *models.DomainAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockDomainAnnotationUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDomainAnnotationUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockDomainAnnotationFinisher is a mock of DomainAnnotationFinisher interface.
type MockDomainAnnotationFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockDomainAnnotationFinisherMockRecorder
}

// MockDomainAnnotationFinisherMockRecorder is the mock recorder for MockDomainAnnotationFinisher.
type MockDomainAnnotationFinisherMockRecorder struct {
	mock *MockDomainAnnotationFinisher
}

// NewMockDomainAnnotationFinisher creates a new mock instance.
func NewMockDomainAnnotationFinisher(ctrl *gomock.Controller) *MockDomainAnnotationFinisher {
	mock := &MockDomainAnnotationFinisher{ctrl: ctrl}
	mock.recorder = &MockDomainAnnotationFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainAnnotationFinisher) EXPECT() *MockDomainAnnotationFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockDomainAnnotationFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.DomainAnnotationSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.DomainAnnotationSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockDomainAnnotationFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockDomainAnnotationFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockDomainAnnotationFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDomainAnnotationFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDomainAnnotationFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockDomainAnnotationFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDomainAnnotationFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDomainAnnotationFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockDomainAnnotationFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.DomainAnnotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.DomainAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockDomainAnnotationFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockDomainAnnotationFinisher)(nil).One), ctx, exec)
}

// MockDomainAnnotationFinder is a mock of DomainAnnotationFinder interface.
type MockDomainAnnotationFinder struct {
	ctrl     *gomock.Controller
	recorder *MockDomainAnnotationFinderMockRecorder
}

// MockDomainAnnotationFinderMockRecorder is the mock recorder for MockDomainAnnotationFinder.
type MockDomainAnnotationFinderMockRecorder struct {
	mock *MockDomainAnnotationFinder
}

// NewMockDomainAnnotationFinder creates a new mock instance.
func NewMockDomainAnnotationFinder(ctrl *gomock.Controller) *MockDomainAnnotationFinder {
	mock := &MockDomainAnnotationFinder{ctrl: ctrl}
	mock.recorder = &MockDomainAnnotationFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainAnnotationFinder) EXPECT() *MockDomainAnnotationFinderMockRecorder {
	return m.recorder
}

// FindDomainAnnotation mocks base method.
func (m *MockDomainAnnotationFinder) FindDomainAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.DomainAnnotation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindDomainAnnotation", varargs...)
	ret0, _ := ret[0].(*models.DomainAnnotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDomainAnnotation indicates an expected call of FindDomainAnnotation.
func (mr *MockDomainAnnotationFinderMockRecorder) FindDomainAnnotation(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDomainAnnotation", reflect.TypeOf((*MockDomainAnnotationFinder)(nil).FindDomainAnnotation), varargs...)
}

// MockDomainAnnotationInserter is a mock of DomainAnnotationInserter interface.
type MockDomainAnnotationInserter struct {
	ctrl     *gomock.Controller
	recorder *MockDomainAnnotationInserterMockRecorder
}

// MockDomainAnnotationInserterMockRecorder is the mock recorder for MockDomainAnnotationInserter.
type MockDomainAnnotationInserterMockRecorder struct {
	mock *MockDomainAnnotationInserter
}

// NewMockDomainAnnotationInserter creates a new mock instance.
func NewMockDomainAnnotationInserter(ctrl *gomock.Controller) *MockDomainAnnotationInserter {
	mock := &MockDomainAnnotationInserter{ctrl: ctrl}
	mock.recorder = &MockDomainAnnotationInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainAnnotationInserter) EXPECT() *MockDomainAnnotationInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockDomainAnnotationInserter) Insert(o *models.DomainAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockDomainAnnotationInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDomainAnnotationInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockDomainAnnotationUpdater is a mock of DomainAnnotationUpdater interface.
type MockDomainAnnotationUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDomainAnnotationUpdaterMockRecorder
}

// MockDomainAnnotationUpdaterMockRecorder is the mock recorder for MockDomainAnnotationUpdater.
type MockDomainAnnotationUpdaterMockRecorder struct {
	mock *MockDomainAnnotationUpdater
}

// NewMockDomainAnnotationUpdater creates a new mock instance.
func NewMockDomainAnnotationUpdater(ctrl *gomock.Controller) *MockDomainAnnotationUpdater {
	mock := &MockDomainAnnotationUpdater{ctrl: ctrl}
	mock.recorder = &MockDomainAnnotationUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainAnnotationUpdater) EXPECT() *MockDomainAnnotationUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockDomainAnnotationUpdater) Update(o *models.DomainAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDomainAnnotationUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDomainAnnotationUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockDomainAnnotationUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockDomainAnnotationUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockDomainAnnotationUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockDomainAnnotationUpdater) UpdateAllSlice(o models.DomainAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockDomainAnnotationUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockDomainAnnotationUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockDomainAnnotationDeleter is a mock of DomainAnnotationDeleter interface.
type MockDomainAnnotationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDomainAnnotationDeleterMockRecorder
}

// MockDomainAnnotationDeleterMockRecorder is the mock recorder for MockDomainAnnotationDeleter.
type MockDomainAnnotationDeleterMockRecorder struct {
	mock *MockDomainAnnotationDeleter
}

// NewMockDomainAnnotationDeleter creates a new mock instance.
func NewMockDomainAnnotationDeleter(ctrl *gomock.Controller) *MockDomainAnnotationDeleter {
	mock := &MockDomainAnnotationDeleter{ctrl: ctrl}
	mock.recorder = &MockDomainAnnotationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainAnnotationDeleter) EXPECT() *MockDomainAnnotationDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDomainAnnotationDeleter) Delete(o *models.DomainAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDomainAnnotationDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDomainAnnotationDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockDomainAnnotationDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockDomainAnnotationDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockDomainAnnotationDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockDomainAnnotationDeleter) DeleteAllSlice(o models.DomainAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockDomainAnnotationDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockDomainAnnotationDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockDomainAnnotationReloader is a mock of DomainAnnotationReloader interface.
type MockDomainAnnotationReloader struct {
	ctrl     *gomock.Controller
	recorder *MockDomainAnnotationReloaderMockRecorder
}

// MockDomainAnnotationReloaderMockRecorder is the mock recorder for MockDomainAnnotationReloader.
type MockDomainAnnotationReloaderMockRecorder struct {
	mock *MockDomainAnnotationReloader
}

// NewMockDomainAnnotationReloader creates a new mock instance.
func NewMockDomainAnnotationReloader(ctrl *gomock.Controller) *MockDomainAnnotationReloader {
	mock := &MockDomainAnnotationReloader{ctrl: ctrl}
	mock.recorder = &MockDomainAnnotationReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainAnnotationReloader) EXPECT() *MockDomainAnnotationReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockDomainAnnotationReloader) Reload(o *models.DomainAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockDomainAnnotationReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockDomainAnnotationReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockDomainAnnotationReloader) ReloadAll(o *models.DomainAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockDomainAnnotationReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockDomainAnnotationReloader)(nil).ReloadAll), o, ctx, exec)
}
