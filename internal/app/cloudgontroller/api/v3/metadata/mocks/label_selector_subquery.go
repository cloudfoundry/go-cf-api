// Code generated by MockGen. DO NOT EDIT.
// Source: label_selector_subquery.go

// Package mock_metadata is a generated GoMock package.
package mock_metadata

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	qm "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MockLabelSelectorFilters is a mock of LabelSelectorFilters interface.
type MockLabelSelectorFilters struct {
	ctrl     *gomock.Controller
	recorder *MockLabelSelectorFiltersMockRecorder
}

// MockLabelSelectorFiltersMockRecorder is the mock recorder for MockLabelSelectorFilters.
type MockLabelSelectorFiltersMockRecorder struct {
	mock *MockLabelSelectorFilters
}

// NewMockLabelSelectorFilters creates a new mock instance.
func NewMockLabelSelectorFilters(ctrl *gomock.Controller) *MockLabelSelectorFilters {
	mock := &MockLabelSelectorFilters{ctrl: ctrl}
	mock.recorder = &MockLabelSelectorFiltersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLabelSelectorFilters) EXPECT() *MockLabelSelectorFiltersMockRecorder {
	return m.recorder
}

// Filters mocks base method.
func (m *MockLabelSelectorFilters) Filters(resourceTable, labelsTable string) []qm.QueryMod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filters", resourceTable, labelsTable)
	ret0, _ := ret[0].([]qm.QueryMod)
	return ret0
}

// Filters indicates an expected call of Filters.
func (mr *MockLabelSelectorFiltersMockRecorder) Filters(resourceTable, labelsTable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filters", reflect.TypeOf((*MockLabelSelectorFilters)(nil).Filters), resourceTable, labelsTable)
}
