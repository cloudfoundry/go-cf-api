package pagination

import (
	"fmt"
	"math"
)

type Link struct {
	Href   string `json:"href"`
	Method string `json:"method,omitempty"`
}

type Pagination struct {
	TotalResults int   `json:"total_results"`
	TotalPages   int   `json:"total_pages"`
	First        Link  `json:"first"`
	Last         Link  `json:"last"`
	Next         *Link `json:"next"`
	Previous     *Link `json:"previous"`
}

type Params struct {
	Page    int    `query:"page" validate:"gte=1"`
	PerPage uint16 `query:"per_page" validate:"gte=1,lte=5000"`
}

func Default() Params {
	return Params{Page: 1, PerPage: 50} //nolint:gomnd // Default values
}

func NewPagination(totalResults int, paginationParams Params, resourcePath string) *Pagination {
	totalPageCount := totalPages(totalResults, paginationParams.PerPage)
	return &Pagination{
		TotalResults: totalResults,
		TotalPages:   totalPageCount,
		First: Link{
			Href: fmt.Sprintf("%s?page=1&per_page=%d", resourcePath, paginationParams.PerPage),
		},
		Last: Link{
			Href: fmt.Sprintf("%s?page=%d&per_page=%d", resourcePath, totalPageCount, paginationParams.PerPage),
		},
		Next:     nextLink(totalResults, paginationParams, resourcePath),
		Previous: previousLink(totalResults, paginationParams, resourcePath),
	}
}

func totalPages(totalResults int, perPage uint16) int {
	pages := int(math.Ceil(float64(totalResults) / float64(perPage)))
	if pages < 1 {
		return 1
	}
	return pages
}

func nextLink(totalResults int, paginationParams Params, resourcePath string) *Link {
	var next *Link
	if paginationParams.Page < totalPages(totalResults, paginationParams.PerPage) {
		next = &Link{
			Href: fmt.Sprintf("%s?page=%d&per_page=%d", resourcePath, paginationParams.Page+1, paginationParams.PerPage),
		}
	}
	return next
}

func previousLink(totalResults int, paginationParams Params, resourcePath string) *Link {
	var previous *Link
	if paginationParams.Page > 1 && paginationParams.Page <= totalPages(totalResults, paginationParams.PerPage)+1 {
		previous = &Link{
			Href: fmt.Sprintf("%s?page=%d&per_page=%d", resourcePath, paginationParams.Page-1, paginationParams.PerPage),
		}
	}
	return previous
}
