// +build unit

package metadata_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/metadata"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

func TestFilters(t *testing.T) {
	resources := "resources"
	labels := "resource_labels"
	baseSubquery := fmt.Sprintf("SELECT 1 FROM %s WHERE (%s = %s.%s)", Quote(labels, "resource_guid", resources, "guid")...)
	t.Parallel()
	cases := map[string]struct {
		lss             LabelSelectorSlice
		expectedQueries []qm.QueryMod
	}{
		"exists": {
			lss: []LabelSelector{{Key: "key1", Operator: Exists}},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("EXISTS (%s AND (%s = ?))", Quote(baseSubquery, "key_name")...),
				Key("key1"),
			)},
		},
		"exists multiple": {
			lss: []LabelSelector{
				{Key: "key1", Operator: Exists},
				{Key: "key2", Operator: Exists},
				{Key: "key3", Operator: Exists},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("EXISTS (%s AND (%s = ?) AND (%s = ?) AND (%s = ?))", Quote(baseSubquery, "key_name", "key_name", "key_name")...),
				Key("key1"), Key("key2"), Key("key3"),
			)},
		},
		"not exists": {
			lss: []LabelSelector{{Key: "key1", Operator: NotExists}},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("NOT EXISTS (%s AND (%s IN (?)))", Quote(baseSubquery, "key_name")...),
				Key("key1"),
			)},
		},
		"not exists multiple": {
			lss: []LabelSelector{
				{Key: "key1", Operator: NotExists},
				{Key: "key2", Operator: NotExists},
				{Key: "key3", Operator: NotExists},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("NOT EXISTS (%s AND (%s IN (?,?,?)))", Quote(baseSubquery, "key_name")...),
				Key("key1"), Key("key2"), Key("key3"),
			)},
		},
		"equal": {
			lss: []LabelSelector{{Key: "key1", Operator: Eq, Values: []string{"value1"}}},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("EXISTS (%s AND (%s = ? AND %s = ?))", Quote(baseSubquery, "key_name", "value")...),
				Key("key1"), "value1",
			)},
		},
		"equal multiple": {
			lss: []LabelSelector{
				{Key: "key1", Operator: Eq, Values: []string{"value1"}},
				{Key: "key2", Operator: EqAlt, Values: []string{"value2"}},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("EXISTS (%s AND (%s = ? AND %s = ?) AND (%s = ? AND %s = ?))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
				Key("key1"), "value1", Key("key2"), "value2",
			)},
		},
		"not equal": {
			lss: []LabelSelector{{Key: "key1", Operator: NotEq, Values: []string{"value1"}}},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("NOT EXISTS (%s AND ((%s = ? AND %s = ?)))", Quote(baseSubquery, "key_name", "value")...),
				Key("key1"), "value1",
			)},
		},
		"not equal multiple": {
			lss: []LabelSelector{
				{Key: "key1", Operator: NotEq, Values: []string{"value1"}},
				{Key: "key2", Operator: NotEq, Values: []string{"value2"}},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("NOT EXISTS (%s AND ((%s = ? AND %s = ?) OR (%s = ? AND %s = ?)))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
				Key("key1"), "value1", Key("key2"), "value2",
			)},
		},
		"in": {
			lss: []LabelSelector{{Key: "key1", Operator: In, Values: []string{"value1", "value2"}}},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("EXISTS (%s AND (%s = ? AND %s IN (?,?)))", Quote(baseSubquery, "key_name", "value")...),
				Key("key1"), "value1", "value2",
			)},
		},
		"in multiple": {
			lss: []LabelSelector{
				{Key: "key1", Operator: In, Values: []string{"value1", "value2"}},
				{Key: "key2", Operator: In, Values: []string{"value3", "value4"}},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("EXISTS (%s AND (%s = ? AND %s IN (?,?)) AND (%s = ? AND %s IN (?,?)))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
				Key("key1"), "value1", "value2", Key("key2"), "value3", "value4",
			)},
		},
		"notin": {
			lss: []LabelSelector{{Key: "key1", Operator: NotIn, Values: []string{"value1", "value2"}}},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("NOT EXISTS (%s AND ((%s = ? AND %s IN (?,?))))", Quote(baseSubquery, "key_name", "value")...),
				Key("key1"), "value1", "value2",
			)},
		},
		"notin multiple": {
			lss: []LabelSelector{
				{Key: "key1", Operator: NotIn, Values: []string{"value1", "value2"}},
				{Key: "key2", Operator: NotIn, Values: []string{"value3", "value4"}},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("NOT EXISTS (%s AND ((%s = ? AND %s IN (?,?)) OR (%s = ? AND %s IN (?,?))))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
				Key("key1"), "value1", "value2", Key("key2"), "value3", "value4",
			)},
		},
		"in and equal": {
			lss: []LabelSelector{
				{Key: "key1", Operator: Eq, Values: []string{"value1"}},
				{Key: "key2", Operator: In, Values: []string{"value2", "value3"}},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("EXISTS (%s AND (%s = ? AND %s = ?) AND (%s = ? AND %s IN (?,?)))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
				Key("key1"), "value1", Key("key2"), "value2", "value3",
			)},
		},
		"notin and not equal": {
			lss: []LabelSelector{
				{Key: "key1", Operator: NotEq, Values: []string{"value1"}},
				{Key: "key2", Operator: NotIn, Values: []string{"value2", "value3"}},
			},
			expectedQueries: []qm.QueryMod{qm.Where(
				fmt.Sprintf("NOT EXISTS (%s AND ((%s = ? AND %s = ?) OR (%s = ? AND %s IN (?,?))))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
				Key("key1"), "value1", Key("key2"), "value2", "value3",
			)},
		},
		"all together": {
			lss: []LabelSelector{
				{Key: "key1", Operator: Exists},
				{Key: "key2", Operator: NotExists},
				{Key: "key3", Operator: Eq, Values: []string{"value3"}},
				{Key: "key4", Operator: NotEq, Values: []string{"value4"}},
				{Key: "key5", Operator: In, Values: []string{"value5", "value6"}},
				{Key: "key6", Operator: NotIn, Values: []string{"value7", "value8"}},
			},
			expectedQueries: []qm.QueryMod{
				qm.Where(
					fmt.Sprintf("EXISTS (%s AND (%s = ?))", Quote(baseSubquery, "key_name")...),
					Key("key1"),
				),
				qm.Where(
					fmt.Sprintf("NOT EXISTS (%s AND (%s IN (?)))", Quote(baseSubquery, "key_name")...),
					Key("key2"),
				),
				qm.Where(
					fmt.Sprintf("EXISTS (%s AND (%s = ? AND %s = ?) AND (%s = ? AND %s IN (?,?)))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
					Key("key3"), "value3", Key("key5"), "value5", "value6",
				),
				qm.Where(
					fmt.Sprintf("NOT EXISTS (%s AND ((%s = ? AND %s = ?) OR (%s = ? AND %s IN (?,?))))", Quote(baseSubquery, "key_name", "value", "key_name", "value")...),
					Key("key4"), "value4", Key("key6"), "value7", "value8",
				),
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := tc.lss.Filters(resources, labels)
			assert.ElementsMatch(t, tc.expectedQueries, got)
		})
	}
}

func Quote(fields ...string) []interface{} {
	quoted := []interface{}{}
	for _, f := range fields {
		quoted = append(quoted, models.Quote(f))
	}
	return quoted
}
