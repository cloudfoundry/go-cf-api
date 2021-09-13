package permissions

//go:generate mockgen -source=$GOFILE -destination=mocks/$GOFILE
type Querier interface {
	AllowedSpaceIDsForUser(userGUID string, roles ...Role) (AllowedSpaceIDs, error)
}

type querier struct{}

//nolint:revive
func NewQuerier() *querier {
	return &querier{}
}
