package common

import (
	"github.com/stretchr/testify/assert"
	commoncontroller "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers/common"
	"reflect"
	"testing"
)


func TestNewPagination(t *testing.T) {
	expected := &Pagination{
		TotalResults: 23,
		TotalPages: 8,
		First: Link{
			Href: "/bommel?page=1&per_page=3",
		},
		Last: Link{
			Href: "/bommel?page=8&per_page=3",
		},
		Next: &Link{
			Href: "/bommel?page=3&per_page=3",
		},
		Previous: &Link{
			Href: "/bommel?page=1&per_page=3",
		}}
	got := NewPagination(23,commoncontroller.PaginationParams{Page: 2,PerPage: 3},"/bommel")
	assert.Equal(t, expected, got)
}

func TestPaginationStruct(t *testing.T){
	assert.Equal(t, 6, reflect.TypeOf(Pagination{}).NumField())
}

func TestTotalPages(t *testing.T) {
	got := totalPages(5,2)
	assert.Equal(t, 3,got)
	got = totalPages(0,2)
	assert.Equal(t, 0,got)
	got = totalPages(5,6)
	assert.Equal(t, 1,got)
}

func TestNextLink(t *testing.T){
	got := nextLink(15, commoncontroller.PaginationParams{Page: 1, PerPage: 14}, "/bommel")
	assert.Equal(t, "/bommel?page=2&per_page=14", got.Href)
	got = nextLink(15, commoncontroller.PaginationParams{Page: 2, PerPage: 14}, "/bommel")
	assert.Nil(t, got)
}

func TestPreviousLink(t *testing.T){
	got := previousLink(15, commoncontroller.PaginationParams{Page: 1, PerPage: 14}, "/bommel")
	assert.Nil(t, got)
	got = previousLink(15, commoncontroller.PaginationParams{Page: 4, PerPage: 4}, "/bommel")
	assert.Equal(t, "/bommel?page=3&per_page=4", got.Href)
	got = previousLink(15, commoncontroller.PaginationParams{Page: 5, PerPage: 4}, "/bommel")
	assert.Equal(t, "/bommel?page=4&per_page=4", got.Href)
}