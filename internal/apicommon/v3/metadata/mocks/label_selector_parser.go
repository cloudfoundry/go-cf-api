// Code generated by MockGen. DO NOT EDIT.
// Source: label_selector_parser.go

// Package mock_metadata is a generated GoMock package.
package mock_metadata

import (
	reflect "reflect"

	metadata "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/metadata"
	gomock "github.com/golang/mock/gomock"
)

// MockLabelSelectorParser is a mock of LabelSelectorParser interface.
type MockLabelSelectorParser struct {
	ctrl     *gomock.Controller
	recorder *MockLabelSelectorParserMockRecorder
}

// MockLabelSelectorParserMockRecorder is the mock recorder for MockLabelSelectorParser.
type MockLabelSelectorParserMockRecorder struct {
	mock *MockLabelSelectorParser
}

// NewMockLabelSelectorParser creates a new mock instance.
func NewMockLabelSelectorParser(ctrl *gomock.Controller) *MockLabelSelectorParser {
	mock := &MockLabelSelectorParser{ctrl: ctrl}
	mock.recorder = &MockLabelSelectorParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLabelSelectorParser) EXPECT() *MockLabelSelectorParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockLabelSelectorParser) Parse(labelSelector string) (metadata.LabelSelectorFilters, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", labelSelector)
	ret0, _ := ret[0].(metadata.LabelSelectorFilters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockLabelSelectorParserMockRecorder) Parse(labelSelector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockLabelSelectorParser)(nil).Parse), labelSelector)
}
