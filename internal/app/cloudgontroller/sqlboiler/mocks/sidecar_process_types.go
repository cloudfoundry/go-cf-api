// Code generated by MockGen. DO NOT EDIT.
// Source: psql_sidecar_process_types.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	boil "github.com/volatiletech/sqlboiler/v4/boil"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// MockSidecarProcessTypeFinisher is a mock of SidecarProcessTypeFinisher interface.
type MockSidecarProcessTypeFinisher struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarProcessTypeFinisherMockRecorder
}

// MockSidecarProcessTypeFinisherMockRecorder is the mock recorder for MockSidecarProcessTypeFinisher.
type MockSidecarProcessTypeFinisherMockRecorder struct {
	mock *MockSidecarProcessTypeFinisher
}

// NewMockSidecarProcessTypeFinisher creates a new mock instance.
func NewMockSidecarProcessTypeFinisher(ctrl *gomock.Controller) *MockSidecarProcessTypeFinisher {
	mock := &MockSidecarProcessTypeFinisher{ctrl: ctrl}
	mock.recorder = &MockSidecarProcessTypeFinisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarProcessTypeFinisher) EXPECT() *MockSidecarProcessTypeFinisherMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockSidecarProcessTypeFinisher) All(ctx context.Context, exec boil.ContextExecutor) (models.SidecarProcessTypeSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, exec)
	ret0, _ := ret[0].(models.SidecarProcessTypeSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockSidecarProcessTypeFinisherMockRecorder) All(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockSidecarProcessTypeFinisher)(nil).All), ctx, exec)
}

// Count mocks base method.
func (m *MockSidecarProcessTypeFinisher) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSidecarProcessTypeFinisherMockRecorder) Count(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSidecarProcessTypeFinisher)(nil).Count), ctx, exec)
}

// Exists mocks base method.
func (m *MockSidecarProcessTypeFinisher) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, exec)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockSidecarProcessTypeFinisherMockRecorder) Exists(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSidecarProcessTypeFinisher)(nil).Exists), ctx, exec)
}

// One mocks base method.
func (m *MockSidecarProcessTypeFinisher) One(ctx context.Context, exec boil.ContextExecutor) (*models.SidecarProcessType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, exec)
	ret0, _ := ret[0].(*models.SidecarProcessType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// One indicates an expected call of One.
func (mr *MockSidecarProcessTypeFinisherMockRecorder) One(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockSidecarProcessTypeFinisher)(nil).One), ctx, exec)
}

// MockSidecarProcessTypeFinder is a mock of SidecarProcessTypeFinder interface.
type MockSidecarProcessTypeFinder struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarProcessTypeFinderMockRecorder
}

// MockSidecarProcessTypeFinderMockRecorder is the mock recorder for MockSidecarProcessTypeFinder.
type MockSidecarProcessTypeFinderMockRecorder struct {
	mock *MockSidecarProcessTypeFinder
}

// NewMockSidecarProcessTypeFinder creates a new mock instance.
func NewMockSidecarProcessTypeFinder(ctrl *gomock.Controller) *MockSidecarProcessTypeFinder {
	mock := &MockSidecarProcessTypeFinder{ctrl: ctrl}
	mock.recorder = &MockSidecarProcessTypeFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarProcessTypeFinder) EXPECT() *MockSidecarProcessTypeFinderMockRecorder {
	return m.recorder
}

// FindSidecarProcessType mocks base method.
func (m *MockSidecarProcessTypeFinder) FindSidecarProcessType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*models.SidecarProcessType, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, exec, iD}
	for _, a := range selectCols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSidecarProcessType", varargs...)
	ret0, _ := ret[0].(*models.SidecarProcessType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSidecarProcessType indicates an expected call of FindSidecarProcessType.
func (mr *MockSidecarProcessTypeFinderMockRecorder) FindSidecarProcessType(ctx, exec, iD interface{}, selectCols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, exec, iD}, selectCols...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSidecarProcessType", reflect.TypeOf((*MockSidecarProcessTypeFinder)(nil).FindSidecarProcessType), varargs...)
}

// MockSidecarProcessTypeInserter is a mock of SidecarProcessTypeInserter interface.
type MockSidecarProcessTypeInserter struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarProcessTypeInserterMockRecorder
}

// MockSidecarProcessTypeInserterMockRecorder is the mock recorder for MockSidecarProcessTypeInserter.
type MockSidecarProcessTypeInserterMockRecorder struct {
	mock *MockSidecarProcessTypeInserter
}

// NewMockSidecarProcessTypeInserter creates a new mock instance.
func NewMockSidecarProcessTypeInserter(ctrl *gomock.Controller) *MockSidecarProcessTypeInserter {
	mock := &MockSidecarProcessTypeInserter{ctrl: ctrl}
	mock.recorder = &MockSidecarProcessTypeInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarProcessTypeInserter) EXPECT() *MockSidecarProcessTypeInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockSidecarProcessTypeInserter) Insert(o *models.SidecarProcessType, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", o, ctx, exec, columns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSidecarProcessTypeInserterMockRecorder) Insert(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSidecarProcessTypeInserter)(nil).Insert), o, ctx, exec, columns)
}

// MockSidecarProcessTypeUpdater is a mock of SidecarProcessTypeUpdater interface.
type MockSidecarProcessTypeUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarProcessTypeUpdaterMockRecorder
}

// MockSidecarProcessTypeUpdaterMockRecorder is the mock recorder for MockSidecarProcessTypeUpdater.
type MockSidecarProcessTypeUpdaterMockRecorder struct {
	mock *MockSidecarProcessTypeUpdater
}

// NewMockSidecarProcessTypeUpdater creates a new mock instance.
func NewMockSidecarProcessTypeUpdater(ctrl *gomock.Controller) *MockSidecarProcessTypeUpdater {
	mock := &MockSidecarProcessTypeUpdater{ctrl: ctrl}
	mock.recorder = &MockSidecarProcessTypeUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarProcessTypeUpdater) EXPECT() *MockSidecarProcessTypeUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockSidecarProcessTypeUpdater) Update(o *models.SidecarProcessType, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", o, ctx, exec, columns)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSidecarProcessTypeUpdaterMockRecorder) Update(o, ctx, exec, columns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSidecarProcessTypeUpdater)(nil).Update), o, ctx, exec, columns)
}

// UpdateAll mocks base method.
func (m *MockSidecarProcessTypeUpdater) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll.
func (mr *MockSidecarProcessTypeUpdaterMockRecorder) UpdateAll(ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockSidecarProcessTypeUpdater)(nil).UpdateAll), ctx, exec, cols)
}

// UpdateAllSlice mocks base method.
func (m *MockSidecarProcessTypeUpdater) UpdateAllSlice(o models.SidecarProcessTypeSlice, ctx context.Context, exec boil.ContextExecutor, cols models.M) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSlice", o, ctx, exec, cols)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllSlice indicates an expected call of UpdateAllSlice.
func (mr *MockSidecarProcessTypeUpdaterMockRecorder) UpdateAllSlice(o, ctx, exec, cols interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSlice", reflect.TypeOf((*MockSidecarProcessTypeUpdater)(nil).UpdateAllSlice), o, ctx, exec, cols)
}

// MockSidecarProcessTypeUpserter is a mock of SidecarProcessTypeUpserter interface.
type MockSidecarProcessTypeUpserter struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarProcessTypeUpserterMockRecorder
}

// MockSidecarProcessTypeUpserterMockRecorder is the mock recorder for MockSidecarProcessTypeUpserter.
type MockSidecarProcessTypeUpserterMockRecorder struct {
	mock *MockSidecarProcessTypeUpserter
}

// NewMockSidecarProcessTypeUpserter creates a new mock instance.
func NewMockSidecarProcessTypeUpserter(ctrl *gomock.Controller) *MockSidecarProcessTypeUpserter {
	mock := &MockSidecarProcessTypeUpserter{ctrl: ctrl}
	mock.recorder = &MockSidecarProcessTypeUpserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarProcessTypeUpserter) EXPECT() *MockSidecarProcessTypeUpserterMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockSidecarProcessTypeUpserter) Upsert(o *models.SidecarProcessType, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", o, ctx, exec, updateColumns, insertColumns)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockSidecarProcessTypeUpserterMockRecorder) Upsert(o, ctx, exec, updateColumns, insertColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockSidecarProcessTypeUpserter)(nil).Upsert), o, ctx, exec, updateColumns, insertColumns)
}

// MockSidecarProcessTypeDeleter is a mock of SidecarProcessTypeDeleter interface.
type MockSidecarProcessTypeDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarProcessTypeDeleterMockRecorder
}

// MockSidecarProcessTypeDeleterMockRecorder is the mock recorder for MockSidecarProcessTypeDeleter.
type MockSidecarProcessTypeDeleterMockRecorder struct {
	mock *MockSidecarProcessTypeDeleter
}

// NewMockSidecarProcessTypeDeleter creates a new mock instance.
func NewMockSidecarProcessTypeDeleter(ctrl *gomock.Controller) *MockSidecarProcessTypeDeleter {
	mock := &MockSidecarProcessTypeDeleter{ctrl: ctrl}
	mock.recorder = &MockSidecarProcessTypeDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarProcessTypeDeleter) EXPECT() *MockSidecarProcessTypeDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSidecarProcessTypeDeleter) Delete(o *models.SidecarProcessType, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSidecarProcessTypeDeleterMockRecorder) Delete(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSidecarProcessTypeDeleter)(nil).Delete), o, ctx, exec)
}

// DeleteAll mocks base method.
func (m *MockSidecarProcessTypeDeleter) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockSidecarProcessTypeDeleterMockRecorder) DeleteAll(ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockSidecarProcessTypeDeleter)(nil).DeleteAll), ctx, exec)
}

// DeleteAllSlice mocks base method.
func (m *MockSidecarProcessTypeDeleter) DeleteAllSlice(o models.SidecarProcessTypeSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllSlice", o, ctx, exec)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllSlice indicates an expected call of DeleteAllSlice.
func (mr *MockSidecarProcessTypeDeleterMockRecorder) DeleteAllSlice(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllSlice", reflect.TypeOf((*MockSidecarProcessTypeDeleter)(nil).DeleteAllSlice), o, ctx, exec)
}

// MockSidecarProcessTypeReloader is a mock of SidecarProcessTypeReloader interface.
type MockSidecarProcessTypeReloader struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarProcessTypeReloaderMockRecorder
}

// MockSidecarProcessTypeReloaderMockRecorder is the mock recorder for MockSidecarProcessTypeReloader.
type MockSidecarProcessTypeReloaderMockRecorder struct {
	mock *MockSidecarProcessTypeReloader
}

// NewMockSidecarProcessTypeReloader creates a new mock instance.
func NewMockSidecarProcessTypeReloader(ctrl *gomock.Controller) *MockSidecarProcessTypeReloader {
	mock := &MockSidecarProcessTypeReloader{ctrl: ctrl}
	mock.recorder = &MockSidecarProcessTypeReloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSidecarProcessTypeReloader) EXPECT() *MockSidecarProcessTypeReloaderMockRecorder {
	return m.recorder
}

// Reload mocks base method.
func (m *MockSidecarProcessTypeReloader) Reload(o *models.SidecarProcessType, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload.
func (mr *MockSidecarProcessTypeReloaderMockRecorder) Reload(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSidecarProcessTypeReloader)(nil).Reload), o, ctx, exec)
}

// ReloadAll mocks base method.
func (m *MockSidecarProcessTypeReloader) ReloadAll(o *models.SidecarProcessTypeSlice, ctx context.Context, exec boil.ContextExecutor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadAll", o, ctx, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadAll indicates an expected call of ReloadAll.
func (mr *MockSidecarProcessTypeReloaderMockRecorder) ReloadAll(o, ctx, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadAll", reflect.TypeOf((*MockSidecarProcessTypeReloader)(nil).ReloadAll), o, ctx, exec)
}
