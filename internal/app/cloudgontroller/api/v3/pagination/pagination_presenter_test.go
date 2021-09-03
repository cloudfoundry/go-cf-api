package pagination_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
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
	got := pagination.NewPagination(0, pagination.DefaultPagination(), "/bommel")
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
	got := pagination.NewPagination(49, pagination.DefaultPagination(), "/bommel")
	assert.Equal(t, expected, got)
}

func TestPaginationStructFieldCount(t *testing.T) {
	assert.Equal(t, 6, reflect.TypeOf(pagination.Pagination{}).NumField())
}
