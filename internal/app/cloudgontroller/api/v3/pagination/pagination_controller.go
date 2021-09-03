package pagination

type Params struct {
	Page    int    `query:"page" validate:"gte=1"`
	PerPage uint16 `query:"per_page" validate:"gte=1,lte=5000"`
}

func DefaultPagination() Params {
	return Params{Page: 1, PerPage: 50} //nolint:gomnd // Default values
}
