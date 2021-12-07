//go:build unit
// +build unit

package pagination_test

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/pagination"
)

func TestNewPaginationCustomPerPage(t *testing.T) {
	expected := &pagination.Pagination{
		TotalResults: 23,
		TotalPages:   8,
		First: pagination.Link{
			Href: "/bommel?page=1&per_page=3",
		},
		Last: pagination.Link{
			Href: "/bommel?page=8&per_page=3",
		},
		Next: &pagination.Link{
			Href: "/bommel?page=3&per_page=3",
		},
		Previous: &pagination.Link{
			Href: "/bommel?page=1&per_page=3",
		},
	}
	got := pagination.NewPagination(23, pagination.Params{Page: 2, PerPage: 3}, "/bommel")
	assert.Equal(t, expected, got)
}

func TestNewPaginationZeroResultsDefaultPerPage(t *testing.T) {
	expected := &pagination.Pagination{
		TotalResults: 0,
		TotalPages:   1,
		First: pagination.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Last: pagination.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Next:     nil,
		Previous: nil,
	}
	got := pagination.NewPagination(0, pagination.Default(), "/bommel")
	assert.Equal(t, expected, got)
}

func TestNewPaginationLessResultsThanPerPage(t *testing.T) {
	expected := &pagination.Pagination{
		TotalResults: 49,
		TotalPages:   1,
		First: pagination.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Last: pagination.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Next:     nil,
		Previous: nil,
	}
	got := pagination.NewPagination(49, pagination.Default(), "/bommel")
	assert.Equal(t, expected, got)
}

func TestPaginationStructFieldCount(t *testing.T) {
	assert.Equal(t, 6, reflect.TypeOf(pagination.Pagination{}).NumField())
}

func TestDefaultPaginationValues(t *testing.T) {
	got := pagination.Default()
	assert.Equal(t, 1, got.Page)
	assert.Equal(t, uint16(50), got.PerPage)
}

func TestPaginationParamsValidationInBounds(t *testing.T) {
	t.Parallel()
	const MaxInt = int(^uint(0) >> 1) // Depending on Architecture Uint can be 32 or 64 bit so we need to get MaxUint dynamically
	tests := map[string]pagination.Params{
		"with lowest possible Page and PerPage":  {1, 1},
		"with largest possible Page and PerPage": {MaxInt, 5000},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			got := validator.New().Struct(pagination.Params{Page: testcase.Page, PerPage: testcase.PerPage})
			assert.Nil(t, got, validatorMessage(got))
		})
	}
}

func TestPaginationParamsValidationOutOfBounds(t *testing.T) {
	t.Parallel()
	tests := map[string]pagination.Params{
		"with Page lower than 1":        {0, 1},
		"with PerPage lower than 1":     {1, 0},
		"with PerPage larger than 5000": {1, 5001},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			got := validator.New().Struct(pagination.Params{Page: testcase.Page, PerPage: testcase.PerPage})
			assert.Error(t, got, validatorMessage(got))
		})
	}
}

func validatorMessage(validatorError error) string {
	if validatorError != nil {
		return validatorError.Error()
	}
	return ""
}
