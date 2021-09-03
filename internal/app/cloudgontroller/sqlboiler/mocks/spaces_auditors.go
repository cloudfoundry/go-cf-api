// Code generated by MockGen. DO NOT EDIT.
// Source: psql_spaces_auditors.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockSpacesAuditorFinisher is a mock of SpacesAuditorFinisher interface.
type MockSpacesAuditorFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesAuditorFinisherMockRecorder
}

// MockSpacesAuditorFinisherMockRecorder is the mock recorder for MockSpacesAuditorFinisher.
type MockSpacesAuditorFinisherMockRecorder struct {
	mock *MockSpacesAuditorFinisher
}

// NewMockSpacesAuditorFinisher creates a new mock instance.
func NewMockSpacesAuditorFinisher(ctrl *gomock.Controller) *MockSpacesAuditorFinisher {
	mock := &MockSpacesAuditorFinisher{ctrl: ctrl}
	mock.recorder = &MockSpacesAuditorFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesAuditorFinisher) EXPECT() *MockSpacesAuditorFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockSpacesAuditorFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.SpacesAuditorSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.SpacesAuditorSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockSpacesAuditorFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockSpacesAuditorFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockSpacesAuditorFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSpacesAuditorFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSpacesAuditorFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockSpacesAuditorFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSpacesAuditorFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSpacesAuditorFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockSpacesAuditorFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.SpacesAuditor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.SpacesAuditor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockSpacesAuditorFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockSpacesAuditorFinisher)(nil).One), ctx, exec)
}

// MockSpacesAuditorFinder is a mock of SpacesAuditorFinder interface.
type MockSpacesAuditorFinder struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesAuditorFinderMockRecorder
}

// MockSpacesAuditorFinderMockRecorder is the mock recorder for MockSpacesAuditorFinder.
type MockSpacesAuditorFinderMockRecorder struct {
	mock *MockSpacesAuditorFinder
}

// NewMockSpacesAuditorFinder creates a new mock instance.
func NewMockSpacesAuditorFinder(ctrl *gomock.Controller) *MockSpacesAuditorFinder {
	mock := &MockSpacesAuditorFinder{ctrl: ctrl}
	mock.recorder = &MockSpacesAuditorFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesAuditorFinder) EXPECT() *MockSpacesAuditorFinderMockRecorder {
	return m.recorder
}

// FindSpacesAuditor mocks base method.
func (m *MockSpacesAuditorFinder) FindSpacesAuditor(ctx context.Context, exec boil.ContextExecutor, spacesAuditorsPK int, selectCols ...string) (*models.SpacesAuditor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, spacesAuditorsPK}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSpacesAuditor", varargs...)
	ret0, _ := ret[0].(*models.SpacesAuditor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSpacesAuditor indicates an expected call of FindSpacesAuditor.
func (mr *MockSpacesAuditorFinderMockRecorder) FindSpacesAuditor(ctx, exec, spacesAuditorsPK interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, spacesAuditorsPK}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSpacesAuditor", reflect.TypeOf((*MockSpacesAuditorFinder)(nil).FindSpacesAuditor), varargs...)
}

// MockSpacesAuditorInserter is a mock of SpacesAuditorInserter interface.
type MockSpacesAuditorInserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesAuditorInserterMockRecorder
}

// MockSpacesAuditorInserterMockRecorder is the mock recorder for MockSpacesAuditorInserter.
type MockSpacesAuditorInserterMockRecorder struct {
	mock *MockSpacesAuditorInserter
}

// NewMockSpacesAuditorInserter creates a new mock instance.
func NewMockSpacesAuditorInserter(ctrl *gomock.Controller) *MockSpacesAuditorInserter {
	mock := &MockSpacesAuditorInserter{ctrl: ctrl}
	mock.recorder = &MockSpacesAuditorInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesAuditorInserter) EXPECT() *MockSpacesAuditorInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockSpacesAuditorInserter) Insert(o *models.SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSpacesAuditorInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSpacesAuditorInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockSpacesAuditorUpdater is a mock of SpacesAuditorUpdater interface.
type MockSpacesAuditorUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesAuditorUpdaterMockRecorder
}

// MockSpacesAuditorUpdaterMockRecorder is the mock recorder for MockSpacesAuditorUpdater.
type MockSpacesAuditorUpdaterMockRecorder struct {
	mock *MockSpacesAuditorUpdater
}

// NewMockSpacesAuditorUpdater creates a new mock instance.
func NewMockSpacesAuditorUpdater(ctrl *gomock.Controller) *MockSpacesAuditorUpdater {
	mock := &MockSpacesAuditorUpdater{ctrl: ctrl}
	mock.recorder = &MockSpacesAuditorUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesAuditorUpdater) EXPECT() *MockSpacesAuditorUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockSpacesAuditorUpdater) Update(o *models.SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSpacesAuditorUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSpacesAuditorUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockSpacesAuditorUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockSpacesAuditorUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockSpacesAuditorUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockSpacesAuditorUpdater) UpdateAllSlice(o models.SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockSpacesAuditorUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockSpacesAuditorUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockSpacesAuditorUpserter is a mock of SpacesAuditorUpserter interface.
type MockSpacesAuditorUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesAuditorUpserterMockRecorder
}

// MockSpacesAuditorUpserterMockRecorder is the mock recorder for MockSpacesAuditorUpserter.
type MockSpacesAuditorUpserterMockRecorder struct {
	mock *MockSpacesAuditorUpserter
}

// NewMockSpacesAuditorUpserter creates a new mock instance.
func NewMockSpacesAuditorUpserter(ctrl *gomock.Controller) *MockSpacesAuditorUpserter {
	mock := &MockSpacesAuditorUpserter{ctrl: ctrl}
	mock.recorder = &MockSpacesAuditorUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesAuditorUpserter) EXPECT() *MockSpacesAuditorUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockSpacesAuditorUpserter) Upsert(o *models.SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSpacesAuditorUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSpacesAuditorUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockSpacesAuditorDeleter is a mock of SpacesAuditorDeleter interface.
type MockSpacesAuditorDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesAuditorDeleterMockRecorder
}

// MockSpacesAuditorDeleterMockRecorder is the mock recorder for MockSpacesAuditorDeleter.
type MockSpacesAuditorDeleterMockRecorder struct {
	mock *MockSpacesAuditorDeleter
}

// NewMockSpacesAuditorDeleter creates a new mock instance.
func NewMockSpacesAuditorDeleter(ctrl *gomock.Controller) *MockSpacesAuditorDeleter {
	mock := &MockSpacesAuditorDeleter{ctrl: ctrl}
	mock.recorder = &MockSpacesAuditorDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesAuditorDeleter) EXPECT() *MockSpacesAuditorDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSpacesAuditorDeleter) Delete(o *models.SpacesAuditor, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSpacesAuditorDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSpacesAuditorDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockSpacesAuditorDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockSpacesAuditorDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockSpacesAuditorDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockSpacesAuditorDeleter) DeleteAllSlice(o models.SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockSpacesAuditorDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockSpacesAuditorDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockSpacesAuditorReloader is a mock of SpacesAuditorReloader interface.
type MockSpacesAuditorReloader struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesAuditorReloaderMockRecorder
}

// MockSpacesAuditorReloaderMockRecorder is the mock recorder for MockSpacesAuditorReloader.
type MockSpacesAuditorReloaderMockRecorder struct {
	mock *MockSpacesAuditorReloader
}

// NewMockSpacesAuditorReloader creates a new mock instance.
func NewMockSpacesAuditorReloader(ctrl *gomock.Controller) *MockSpacesAuditorReloader {
	mock := &MockSpacesAuditorReloader{ctrl: ctrl}
	mock.recorder = &MockSpacesAuditorReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesAuditorReloader) EXPECT() *MockSpacesAuditorReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockSpacesAuditorReloader) Reload(o *models.SpacesAuditor, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockSpacesAuditorReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSpacesAuditorReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockSpacesAuditorReloader) ReloadAll(o *models.SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockSpacesAuditorReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockSpacesAuditorReloader)(nil).ReloadAll), o, ctx, exec)
}
