package common

type PaginationParams struct {
	Page    int `query:"page" validate:"gte=1"`
	PerPage uint16 `query:"per_page" validate:"gte=1,lte=5000"`
}

func DefaultPagination() PaginationParams {
	return PaginationParams{Page: 1, PerPage: 50} //nolint:gomnd // Default values
}
