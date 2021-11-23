package permissions

import (
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlhelpers"
)

//go:generate mockgen -source=$GOFILE -destination=mocks/$GOFILE
type AllowedSpaceIDs interface {
	With() []qm.QueryMod
	Contains(column string) qm.QueryMod
}

type (
	Role string

	allowedSpaceIDs struct {
		userGUID               string
		tableName              string
		tableExpression        *models.Subquery
		orgRolesTableName      string
		orgRowsTableExpression *models.Subquery
	}
)

const (
	SpaceDeveloper    Role = "SpaceDeveloper"
	SpaceSupporter    Role = "SpaceSupporter"
	SpaceManager      Role = "SpaceManager"
	SpaceAuditor      Role = "SpaceAuditor"
	OrgAuditor        Role = "OrgAuditor"
	OrgBillingManager Role = "OrgBillingManager"
	OrgManager        Role = "OrgManager"
)

//nolint:gochecknoglobals
var (
	// The UNION requires that all the tables have the same schema so just pick one to use the column names from
	spaceIDColumn        = models.Quote(models.SpacesDeveloperColumns.SpaceID)
	userIDColumn         = models.Quote(models.SpacesDeveloperColumns.UserID)
	organizationIDColumn = models.Quote(models.OrganizationsManagerColumns.OrganizationID)
	orgRolesColumns      = fmt.Sprintf("%s, %s",
		models.Quote(models.OrganizationsManagerColumns.OrganizationID),
		models.Quote(models.OrganizationsManagerColumns.UserID))
	orgRolesTable        = models.Quote("org_roles")
	allowedSpaceIDsTable = models.Quote("allowed_space_ids")

	AllSpaceRoles = []Role{SpaceDeveloper, SpaceSupporter, SpaceManager, SpaceAuditor}
	AllOrgRoles   = []Role{OrgManager, OrgBillingManager, OrgAuditor}
)

//nolint:funlen // length is not a problem for now
func (q *querier) AllowedSpaceIDsForUser(userGUID string, roles ...Role) (AllowedSpaceIDs, error) {
	if len(userGUID) == 0 {
		return nil, errors.New("no user provided")
	}
	if len(roles) == 0 {
		return nil, errors.New("no roles provided")
	}

	var tables []sqlhelpers.UnionTable
	var orgTables []sqlhelpers.UnionTable

	for _, role := range roles {
		switch role {
		case SpaceAuditor:
			tables = append(tables, sqlhelpers.UnionTable{Table: models.TableNames.SpacesAuditors, Column: spaceIDColumn})
		case SpaceDeveloper:
			tables = append(tables, sqlhelpers.UnionTable{Table: models.TableNames.SpacesDevelopers, Column: spaceIDColumn})
		case SpaceManager:
			tables = append(tables, sqlhelpers.UnionTable{Table: models.TableNames.SpacesManagers, Column: spaceIDColumn})
		case SpaceSupporter:
			tables = append(tables, sqlhelpers.UnionTable{Table: models.TableNames.SpacesApplicationSupporters, Column: spaceIDColumn})
		case OrgAuditor:
			orgTables = append(orgTables, sqlhelpers.UnionTable{Table: models.TableNames.OrganizationsAuditors, Column: orgRolesColumns})
		case OrgBillingManager:
			orgTables = append(orgTables, sqlhelpers.UnionTable{Table: models.TableNames.OrganizationsBillingManagers, Column: orgRolesColumns})
		case OrgManager:
			orgTables = append(orgTables, sqlhelpers.UnionTable{Table: models.TableNames.OrganizationsManagers, Column: orgRolesColumns})
		}
	}

	var qms []qm.QueryMod
	var orgRolesSubquery *models.Subquery
	// Don't both with extra joins if no org roles were required
	if len(orgTables) > 0 {
		orgRolesSubquery = models.NewSubquery(
			sqlhelpers.FromUnion(orgTables),
		)
		tables = append(tables, sqlhelpers.UnionTable{
			Table:  orgRolesTable,
			Column: fmt.Sprintf("%s AS %s", models.Quote(models.SpaceTableColumns.ID), spaceIDColumn),
		})
		qms = append(qms, qm.InnerJoin(
			fmt.Sprintf("%s ON %s = %s",
				models.Quote(models.TableNames.Organizations),
				organizationIDColumn,
				models.Quote(models.OrganizationTableColumns.ID))))
		qms = append(qms, qm.InnerJoin(
			fmt.Sprintf("%s ON %s = %s",
				models.Quote(models.TableNames.Spaces),
				models.Quote(models.SpaceTableColumns.OrganizationID),
				models.Quote(models.OrganizationTableColumns.ID))))
	}

	tableName := allowedSpaceIDsTable
	qms = append(qms,
		sqlhelpers.FromUnion(tables),
		qm.InnerJoin(fmt.Sprintf("%s ON %s = %s", models.Quote(models.TableNames.Users), userIDColumn, models.Quote(models.UserTableColumns.ID))),
		qmhelper.Where(models.Quote(models.UserTableColumns.GUID), qmhelper.EQ, userGUID),
	)
	subquery := models.NewSubquery(qms...)

	return &allowedSpaceIDs{
		userGUID,
		tableName,
		subquery,
		orgRolesTable,
		orgRolesSubquery,
	}, nil
}

func (ws *allowedSpaceIDs) With() []qm.QueryMod {
	var qms []qm.QueryMod

	// Don't both with extra table expression if no org roles were required
	// This has to come first so that org_roles table is defined before allowed_space_ids table
	if ws.orgRowsTableExpression != nil {
		orgRolesSubquery, orgRolesArgs := ws.orgRowsTableExpression.SQL()
		qms = append(qms, qm.With(fmt.Sprintf("%s AS (%s)", ws.orgRolesTableName, orgRolesSubquery), orgRolesArgs...))
	}
	subquery, args := ws.tableExpression.SQL()
	qms = append(qms, qm.With(fmt.Sprintf("%s AS (%s)", ws.tableName, subquery), args...))

	return qms
}

func (ws *allowedSpaceIDs) Contains(column string) qm.QueryMod {
	selectQuery, _ := models.NewSubquery(qm.Select(spaceIDColumn), qm.From(ws.tableName)).SQL()
	return qm.Or(fmt.Sprintf("%s IN (%s)", models.Quote(column), selectQuery))
}
