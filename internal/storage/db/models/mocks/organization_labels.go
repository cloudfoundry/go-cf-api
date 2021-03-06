//go:build unit
// +build unit

//

// Code generated by MockGen. DO NOT EDIT.
// Source: psql_organization_labels.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
)

// MockOrganizationLabelUpserter is a mock of OrganizationLabelUpserter interface.
type MockOrganizationLabelUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationLabelUpserterMockRecorder
}

// MockOrganizationLabelUpserterMockRecorder is the mock recorder for MockOrganizationLabelUpserter.
type MockOrganizationLabelUpserterMockRecorder struct {
	mock *MockOrganizationLabelUpserter
}

// NewMockOrganizationLabelUpserter creates a new mock instance.
func NewMockOrganizationLabelUpserter(ctrl *gomock.Controller) *MockOrganizationLabelUpserter {
	mock := &MockOrganizationLabelUpserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationLabelUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationLabelUpserter) EXPECT() *MockOrganizationLabelUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockOrganizationLabelUpserter) Upsert(o *models.OrganizationLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOrganizationLabelUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOrganizationLabelUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockOrganizationLabelFinisher is a mock of OrganizationLabelFinisher interface.
type MockOrganizationLabelFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationLabelFinisherMockRecorder
}

// MockOrganizationLabelFinisherMockRecorder is the mock recorder for MockOrganizationLabelFinisher.
type MockOrganizationLabelFinisherMockRecorder struct {
	mock *MockOrganizationLabelFinisher
}

// NewMockOrganizationLabelFinisher creates a new mock instance.
func NewMockOrganizationLabelFinisher(ctrl *gomock.Controller) *MockOrganizationLabelFinisher {
	mock := &MockOrganizationLabelFinisher{ctrl: ctrl}
	mock.recorder = &MockOrganizationLabelFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationLabelFinisher) EXPECT() *MockOrganizationLabelFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockOrganizationLabelFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.OrganizationLabelSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.OrganizationLabelSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockOrganizationLabelFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockOrganizationLabelFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockOrganizationLabelFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrganizationLabelFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrganizationLabelFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockOrganizationLabelFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockOrganizationLabelFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockOrganizationLabelFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockOrganizationLabelFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.OrganizationLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.OrganizationLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockOrganizationLabelFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOrganizationLabelFinisher)(nil).One), ctx, exec)
}

// MockOrganizationLabelFinder is a mock of OrganizationLabelFinder interface.
type MockOrganizationLabelFinder struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationLabelFinderMockRecorder
}

// MockOrganizationLabelFinderMockRecorder is the mock recorder for MockOrganizationLabelFinder.
type MockOrganizationLabelFinderMockRecorder struct {
	mock *MockOrganizationLabelFinder
}

// NewMockOrganizationLabelFinder creates a new mock instance.
func NewMockOrganizationLabelFinder(ctrl *gomock.Controller) *MockOrganizationLabelFinder {
	mock := &MockOrganizationLabelFinder{ctrl: ctrl}
	mock.recorder = &MockOrganizationLabelFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationLabelFinder) EXPECT() *MockOrganizationLabelFinderMockRecorder {
	return m.recorder
}

// FindOrganizationLabel mocks base method.
func (m *MockOrganizationLabelFinder) FindOrganizationLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.OrganizationLabel, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganizationLabel", varargs...)
	ret0, _ := ret[0].(*models.OrganizationLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationLabel indicates an expected call of FindOrganizationLabel.
func (mr *MockOrganizationLabelFinderMockRecorder) FindOrganizationLabel(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationLabel", reflect.TypeOf((*MockOrganizationLabelFinder)(nil).FindOrganizationLabel), varargs...)
}

// MockOrganizationLabelInserter is a mock of OrganizationLabelInserter interface.
type MockOrganizationLabelInserter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationLabelInserterMockRecorder
}

// MockOrganizationLabelInserterMockRecorder is the mock recorder for MockOrganizationLabelInserter.
type MockOrganizationLabelInserterMockRecorder struct {
	mock *MockOrganizationLabelInserter
}

// NewMockOrganizationLabelInserter creates a new mock instance.
func NewMockOrganizationLabelInserter(ctrl *gomock.Controller) *MockOrganizationLabelInserter {
	mock := &MockOrganizationLabelInserter{ctrl: ctrl}
	mock.recorder = &MockOrganizationLabelInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationLabelInserter) EXPECT() *MockOrganizationLabelInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockOrganizationLabelInserter) Insert(o *models.OrganizationLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationLabelInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationLabelInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockOrganizationLabelUpdater is a mock of OrganizationLabelUpdater interface.
type MockOrganizationLabelUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationLabelUpdaterMockRecorder
}

// MockOrganizationLabelUpdaterMockRecorder is the mock recorder for MockOrganizationLabelUpdater.
type MockOrganizationLabelUpdaterMockRecorder struct {
	mock *MockOrganizationLabelUpdater
}

// NewMockOrganizationLabelUpdater creates a new mock instance.
func NewMockOrganizationLabelUpdater(ctrl *gomock.Controller) *MockOrganizationLabelUpdater {
	mock := &MockOrganizationLabelUpdater{ctrl: ctrl}
	mock.recorder = &MockOrganizationLabelUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationLabelUpdater) EXPECT() *MockOrganizationLabelUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockOrganizationLabelUpdater) Update(o *models.OrganizationLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationLabelUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationLabelUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockOrganizationLabelUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockOrganizationLabelUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockOrganizationLabelUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockOrganizationLabelUpdater) UpdateAllSlice(o models.OrganizationLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockOrganizationLabelUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockOrganizationLabelUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockOrganizationLabelDeleter is a mock of OrganizationLabelDeleter interface.
type MockOrganizationLabelDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationLabelDeleterMockRecorder
}

// MockOrganizationLabelDeleterMockRecorder is the mock recorder for MockOrganizationLabelDeleter.
type MockOrganizationLabelDeleterMockRecorder struct {
	mock *MockOrganizationLabelDeleter
}

// NewMockOrganizationLabelDeleter creates a new mock instance.
func NewMockOrganizationLabelDeleter(ctrl *gomock.Controller) *MockOrganizationLabelDeleter {
	mock := &MockOrganizationLabelDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationLabelDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationLabelDeleter) EXPECT() *MockOrganizationLabelDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationLabelDeleter) Delete(o *models.OrganizationLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationLabelDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationLabelDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockOrganizationLabelDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockOrganizationLabelDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockOrganizationLabelDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockOrganizationLabelDeleter) DeleteAllSlice(o models.OrganizationLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockOrganizationLabelDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockOrganizationLabelDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockOrganizationLabelReloader is a mock of OrganizationLabelReloader interface.
type MockOrganizationLabelReloader struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationLabelReloaderMockRecorder
}

// MockOrganizationLabelReloaderMockRecorder is the mock recorder for MockOrganizationLabelReloader.
type MockOrganizationLabelReloaderMockRecorder struct {
	mock *MockOrganizationLabelReloader
}

// NewMockOrganizationLabelReloader creates a new mock instance.
func NewMockOrganizationLabelReloader(ctrl *gomock.Controller) *MockOrganizationLabelReloader {
	mock := &MockOrganizationLabelReloader{ctrl: ctrl}
	mock.recorder = &MockOrganizationLabelReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationLabelReloader) EXPECT() *MockOrganizationLabelReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockOrganizationLabelReloader) Reload(o *models.OrganizationLabel, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockOrganizationLabelReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockOrganizationLabelReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockOrganizationLabelReloader) ReloadAll(o *models.OrganizationLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockOrganizationLabelReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockOrganizationLabelReloader)(nil).ReloadAll), o, ctx, exec)
}
