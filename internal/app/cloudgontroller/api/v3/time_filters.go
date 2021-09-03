package v3

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
)

var (
	operatorRegex = regexp.MustCompile(`^.*_ats(\[[a-zA-Z]+\]$|$)`)
	Operators     = map[string]string{ //nolint:gochecknoglobals // Cannot define const maps https://github.com/leighmcculloch/gochecknoglobals/issues/19
		"[lt]":  string(qmhelper.LT),
		"[lte]": string(qmhelper.LTE),
		"[gt]":  string(qmhelper.GT),
		"[gte]": string(qmhelper.GTE),
		"":      string(qmhelper.EQ),
	}
)

type TimeFilter struct {
	Timestamp time.Time
	Operator  string
}

func ParseTimeFilters(c echo.Context) ([]TimeFilter, []TimeFilter, error) {
	// Suffix for created_at and updated_at get removed to parse query parameters correctly
	var createdAts, updatedAts []TimeFilter
	for param, values := range c.QueryParams() {
		switch {
		case strings.HasPrefix(param, "created_ats"):
			filter, err := createFilter(param, values)
			if err != nil {
				return nil, nil, err
			}
			createdAts = append(createdAts, filter)
		case strings.HasPrefix(param, "updated_ats"):
			filter, err := createFilter(param, values)
			if err != nil {
				return nil, nil, err
			}
			updatedAts = append(updatedAts, filter)
		}
	}
	return createdAts, updatedAts, nil
}

func createFilter(param string, values []string) (TimeFilter, error) {
	matches := operatorRegex.FindStringSubmatch(param)
	if matches == nil {
		return TimeFilter{}, fmt.Errorf("could not parse 'created_ats' filter %s", param)
	}
	if _, ok := Operators[matches[1]]; !ok {
		return TimeFilter{}, fmt.Errorf("unrecognized comparison operator '%s'", matches[1])
	}
	timestamp, _ := time.Parse(time.RFC3339, values[0])
	return TimeFilter{
		Timestamp: timestamp,
		Operator:  matches[1],
	}, nil
}
