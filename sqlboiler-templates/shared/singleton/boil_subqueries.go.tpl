var subqueryDialect = drivers.Dialect{
	LQ: 0x{{printf "%x" .Dialect.LQ}},
	RQ: 0x{{printf "%x" .Dialect.RQ}},

	UseIndexPlaceholders:    false,
	UseLastInsertID:         {{.Dialect.UseLastInsertID}},
	UseSchema:               {{.Dialect.UseSchema}},
	UseDefaultKeyword:       {{.Dialect.UseDefaultKeyword}},
	UseAutoColumns:          {{.Dialect.UseAutoColumns}},
	UseTopClause:            {{.Dialect.UseTopClause}},
	UseOutputClause:         {{.Dialect.UseOutputClause}},
	UseCaseWhenExistsClause: {{.Dialect.UseCaseWhenExistsClause}},
}

type Subquery struct {
	*queries.Query
}

// NewSubquery initializes a new Subquery using the passed in QueryMods to be used as part of a larger Query
func NewSubquery(mods ...qm.QueryMod) *Subquery {
	sq := &Subquery{
		Query: NewQuery(mods...),
	}
	queries.SetDialect(sq.Query, &subqueryDialect)

	return sq
}

func (s *Subquery) SQL() (string, []interface{}) {
	sq, args := queries.BuildQuery(s.Query)
	return strings.TrimSuffix(sq, ";"), args
}
