package common_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	commoncontroller "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers/common"
	commonpresenter "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/presenter/common"
)

func TestNewPaginationCustomPerPage(t *testing.T) {
	expected := &commonpresenter.Pagination{
		TotalResults: 23,
		TotalPages:   8,
		First: commonpresenter.Link{
			Href: "/bommel?page=1&per_page=3",
		},
		Last: commonpresenter.Link{
			Href: "/bommel?page=8&per_page=3",
		},
		Next: &commonpresenter.Link{
			Href: "/bommel?page=3&per_page=3",
		},
		Previous: &commonpresenter.Link{
			Href: "/bommel?page=1&per_page=3",
		},
	}
	got := commonpresenter.NewPagination(23, commoncontroller.PaginationParams{Page: 2, PerPage: 3}, "/bommel")
	assert.Equal(t, expected, got)
}

func TestNewPaginationZeroResultsDefaultPerPage(t *testing.T) {
	expected := &commonpresenter.Pagination{
		TotalResults: 0,
		TotalPages:   1,
		First: commonpresenter.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Last: commonpresenter.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Next:     nil,
		Previous: nil,
	}
	got := commonpresenter.NewPagination(0, commoncontroller.DefaultPagination(), "/bommel")
	assert.Equal(t, expected, got)
}

func TestNewPaginationLessResultsThanPerPage(t *testing.T) {
	expected := &commonpresenter.Pagination{
		TotalResults: 49,
		TotalPages:   1,
		First: commonpresenter.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Last: commonpresenter.Link{
			Href: "/bommel?page=1&per_page=50",
		},
		Next:     nil,
		Previous: nil,
	}
	got := commonpresenter.NewPagination(49, commoncontroller.DefaultPagination(), "/bommel")
	assert.Equal(t, expected, got)
}

func TestPaginationStructFieldCount(t *testing.T) {
	assert.Equal(t, 6, reflect.TypeOf(commonpresenter.Pagination{}).NumField())
}