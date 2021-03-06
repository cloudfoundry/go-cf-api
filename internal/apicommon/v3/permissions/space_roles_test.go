//go:build unit
// +build unit

package permissions_test

import (
	"fmt"
	"testing"

	. "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/permissions"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestAllowedSpaceIDsForUser(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		userGUID         string
		roles            []Role
		expectedWith     []qm.QueryMod
		columnName       string
		expectedContains qm.QueryMod
		expectedErr      string
	}{
		"spaces single role ": {
			userGUID: "user-guid1",
			roles:    []Role{SpaceAuditor},
			expectedWith: []qm.QueryMod{qm.With(fmt.Sprintf("%s AS (SELECT %s FROM %s INNER JOIN %s ON %s = %s.%s WHERE (%s.%s = ?))",
				models.Quote("allowed_space_ids"),
				models.Quote("space_id"),
				models.Quote("spaces_auditors"),
				models.Quote("users"),
				models.Quote("user_id"),
				models.Quote("users"),
				models.Quote("id"),
				models.Quote("users"),
				models.Quote("guid"),
			), "user-guid1")},
			columnName:       "foo",
			expectedContains: qm.Or(fmt.Sprintf("%s IN (SELECT %s FROM %s)", models.Quote("foo"), models.Quote("space_id"), models.Quote("allowed_space_ids"))),
		},
		"spaces multiple roles": {
			userGUID: "user-guid1",
			roles:    []Role{SpaceAuditor, SpaceDeveloper},
			expectedWith: []qm.QueryMod{qm.With(fmt.Sprintf("%s AS (SELECT %s FROM %s UNION SELECT %s FROM %s INNER JOIN %s ON %s = %s.%s WHERE (%s.%s = ?))",
				models.Quote("allowed_space_ids"),
				models.Quote("space_id"),
				models.Quote("spaces_auditors"),
				models.Quote("space_id"),
				models.Quote("spaces_developers"),
				models.Quote("users"),
				models.Quote("user_id"),
				models.Quote("users"),
				models.Quote("id"),
				models.Quote("users"),
				models.Quote("guid"),
			), "user-guid1")},
			columnName:       "foo",
			expectedContains: qm.Or(fmt.Sprintf("%s IN (SELECT %s FROM %s)", models.Quote("foo"), models.Quote("space_id"), models.Quote("allowed_space_ids"))),
		},
		"orgs multiple roles": {
			userGUID: "user-guid1",
			roles:    []Role{OrgAuditor, OrgBillingManager},
			expectedWith: []qm.QueryMod{
				qm.With(
					fmt.Sprintf(
						"%s AS (SELECT %s, %s FROM %s UNION SELECT %s, %s FROM %s)",
						models.Quote("org_roles"),
						models.Quote("organization_id"),
						models.Quote("user_id"),
						models.Quote("organizations_auditors"),
						models.Quote("organization_id"),
						models.Quote("user_id"),
						models.Quote("organizations_billing_managers"),
					),
				),
				qm.With(
					fmt.Sprintf(
						"%s AS (SELECT %s.%s AS %s FROM %s INNER JOIN %s ON %s = %s.%s INNER JOIN %s ON %s.%s = %s.%s INNER JOIN %s ON %s = %s.%s WHERE (%s.%s = ?))",
						models.Quote("allowed_space_ids"),
						models.Quote("spaces"),
						models.Quote("id"),
						models.Quote("space_id"),
						models.Quote("org_roles"),
						models.Quote("organizations"),
						models.Quote("organization_id"),
						models.Quote("organizations"),
						models.Quote("id"),
						models.Quote("spaces"),
						models.Quote("spaces"),
						models.Quote("organization_id"),
						models.Quote("organizations"),
						models.Quote("id"),
						models.Quote("users"),
						models.Quote("user_id"),
						models.Quote("users"),
						models.Quote("id"),
						models.Quote("users"),
						models.Quote("guid"),
					),
					"user-guid1",
				),
			},
			columnName:       "foo",
			expectedContains: qm.Or(fmt.Sprintf("%s IN (SELECT %s FROM %s)", models.Quote("foo"), models.Quote("space_id"), models.Quote("allowed_space_ids"))),
		},
		"roles not provided": {
			userGUID:    "user-guid1",
			roles:       []Role{},
			columnName:  "foo",
			expectedErr: "no roles provided",
		},
		"userGUID not provided": {
			userGUID:    "",
			roles:       []Role{SpaceAuditor},
			columnName:  "foo",
			expectedErr: "no user provided",
		},
	}

	qur := NewQuerier()
	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			asid, err := qur.AllowedSpaceIDsForUser(testCase.userGUID, testCase.roles...)
			if testCase.expectedErr != "" {
				assert.EqualError(t, err, testCase.expectedErr)
				assert.Nil(t, asid)
				return
			}
			with := asid.With()
			contains := asid.Contains(testCase.columnName)
			assert.Equal(t, testCase.expectedWith, with)
			assert.Equal(t, testCase.expectedContains, contains)
		})
	}
}
