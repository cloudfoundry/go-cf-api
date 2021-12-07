//go:build unit
// +build unit

//nolint:forcetypeassert // Casting from mock calls doesn't need to be checked
package securitygroups //nolint:testpackage // we have to assign package level vars due to models using static functions

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/auth"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/pagination"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

type ListSecurityGroupTestSuite struct {
	SecurityGroupControllerTestSuite
}

func TestListSecurityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(ListSecurityGroupTestSuite))
}

func (s *ListSecurityGroupTestSuite) SetupTest() {
	s.SetupTestSuite(http.MethodGet, "http://localhost:8080/v3/security_groups")
	s.ctx.Set(auth.Scopes, []string{string(auth.Admin)})
}

func (s *ListSecurityGroupTestSuite) TestStatuOKWithASecurityGroup() {
	securityGroups := models.SecurityGroupSlice{{GUID: "a-security-group"}}

	s.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)

	s.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)
	s.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(securityGroups, nil)

	s.NoError(s.controller.List(s.ctx))
	s.Equal(http.StatusOK, s.rec.Code)

	s.presenter.AssertCalled(s.T(), "ListResponseObject", securityGroups, mock.Anything, mock.Anything, mock.Anything)
}

func (s *ListSecurityGroupTestSuite) TestStatusOKWhenNoSecurityGroups() {
	var securityGroups models.SecurityGroupSlice
	s.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)
	s.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(securityGroups, sql.ErrNoRows)

	s.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)

	s.NoError(s.controller.List(s.ctx))
	s.Equal(http.StatusOK, s.rec.Code)

	s.presenter.AssertCalled(s.T(), "ListResponseObject", securityGroups, int64(0), pagination.Params{Page: 1, PerPage: 50}, mock.Anything)
}

func (s *ListSecurityGroupTestSuite) TestWithPaginationParameters() {
	securityGroups := models.SecurityGroupSlice{{GUID: "a-security-group"}}

	s.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(3), nil)
	s.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)
	s.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(securityGroups, nil)

	s.ctx.Request().URL.RawQuery = "per_page=2&page=3"

	s.NoError(s.controller.List(s.ctx))
	s.querierFunc.AssertNumberOfCalls(s.T(), "Get", 2)
	countQueryMods := s.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
	allQueryMods := s.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)

	s.Contains(allQueryMods, qm.Limit(2))
	s.NotContains(countQueryMods, qm.Limit(2))
	s.Contains(allQueryMods, qm.Offset(4))
	s.NotContains(countQueryMods, qm.Offset(4))
	s.presenter.AssertCalled(s.T(), "ListResponseObject", securityGroups, int64(3), pagination.Params{Page: 3, PerPage: 2}, mock.Anything)
}

func (s *ListSecurityGroupTestSuite) TestFilters() {
	baseQueryMods := []qm.QueryMod{
		qm.Limit(50),
		qm.Offset(0),
		qm.Load(qm.Rels(models.SecurityGroupRels.SecurityGroupsSpaces, models.SecurityGroupsSpaceRels.Space)),
		qm.Load(qm.Rels(models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces, models.StagingSecurityGroupsSpaceRels.StagingSpace)),
	}
	now := time.Now().UTC().Format(time.RFC3339)
	cases := map[string]struct {
		query                string
		expectedCountFilters []qm.QueryMod
		expectedAllFilters   []qm.QueryMod
		expectedErr          *v3.CfAPIError
	}{
		"no name": {
			query:                "names=",
			expectedCountFilters: []qm.QueryMod{qmhelper.WhereIsNull("security_groups.name")},
			expectedAllFilters:   []qm.QueryMod{qmhelper.WhereIsNull("security_groups.name")},
		},
		"single name": {
			query:                "names=test-sg",
			expectedCountFilters: []qm.QueryMod{qm.WhereIn("security_groups.name IN ?", "test-sg")},
			expectedAllFilters:   []qm.QueryMod{qm.WhereIn("security_groups.name IN ?", "test-sg")},
		},
		"multiple names": {
			query:                "names=test-sg-1,test-sg-2,test-sg-3",
			expectedCountFilters: []qm.QueryMod{qm.WhereIn("security_groups.name IN ?", "test-sg-1", "test-sg-2", "test-sg-3")},
			expectedAllFilters:   []qm.QueryMod{qm.WhereIn("security_groups.name IN ?", "test-sg-1", "test-sg-2", "test-sg-3")},
		},
		"order by created_at": {
			query:                "order_by=created_at",
			expectedCountFilters: []qm.QueryMod{qm.OrderBy("created_at ASC")},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("created_at ASC")},
		},
		"order by -created_at": {
			query:                "order_by=-created_at",
			expectedCountFilters: []qm.QueryMod{qm.OrderBy("created_at DESC")},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("created_at DESC")},
		},
		"order by updated_at": {
			query:                "order_by=updated_at",
			expectedCountFilters: []qm.QueryMod{qm.OrderBy("updated_at ASC")},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("updated_at ASC")},
		},
		"order by -updated_at": {
			query:                "order_by=-updated_at",
			expectedCountFilters: []qm.QueryMod{qm.OrderBy("updated_at DESC")},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("updated_at DESC")},
		},
		"order by unknown filter": {
			query:       "order_by=foo",
			expectedErr: v3.BadQueryParameter(errors.New("validation error")),
		},
		"filter by time": {
			query: fmt.Sprintf("created_ats=%s&updated_ats=%s", now, now),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at = ?", now),
				qm.Where("security_groups.updated_at = ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at = ?", now),
				qm.Where("security_groups.updated_at = ?", now),
			},
		},
		"filter by time with lt/gt suffix": {
			query: fmt.Sprintf("created_ats[lt]=%s&updated_ats[gt]=%s", now, now),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at < ?", now),
				qm.Where("security_groups.updated_at > ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at < ?", now),
				qm.Where("security_groups.updated_at > ?", now),
			},
		},
		"filter by time with lte/gte suffix": {
			query: fmt.Sprintf("created_ats[gte]=%s&updated_ats[lte]=%s", now, now),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at >= ?", now),
				qm.Where("security_groups.updated_at <= ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at >= ?", now),
				qm.Where("security_groups.updated_at <= ?", now),
			},
		},
		"filter by time between timestamps": {
			query: fmt.Sprintf(
				"created_ats[gte]=%s&created_ats[lte]=%s",
				time.Now().UTC().Add(-1*time.Hour).Format(time.RFC3339),
				time.Now().UTC().Add(1*time.Hour).Format(time.RFC3339),
			),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at >= ?", time.Now().UTC().Add(-1*time.Hour).Format(time.RFC3339)),
				qm.Where("security_groups.created_at <= ?", time.Now().UTC().Add(1*time.Hour).Format(time.RFC3339)),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.Where("security_groups.created_at >= ?", time.Now().UTC().Add(-1*time.Hour).Format(time.RFC3339)),
				qm.Where("security_groups.created_at <= ?", time.Now().UTC().Add(1*time.Hour).Format(time.RFC3339)),
			},
		},
		"filter by time with invalid param": {
			query:       fmt.Sprintf("created_ats(lte)=%s", now),
			expectedErr: v3.BadQueryParameter(errors.New("validation error")),
		},
		"filter by time with invalid comparison operator": {
			query:       fmt.Sprintf("created_ats[foo]=%s", now),
			expectedErr: v3.BadQueryParameter(errors.New("validation error")),
		},
		"filter by everything": {
			query: fmt.Sprintf(
				"names=test-sg-1,test-sg-2&order_by=-updated_at&created_ats[gt]=%s&updated_ats[lte]=%s",
				now, now,
			),
			expectedCountFilters: []qm.QueryMod{
				qm.WhereIn("security_groups.name IN ?", "test-sg-1", "test-sg-2"),
				qm.OrderBy("updated_at DESC"),
				qm.Where("security_groups.created_at > ?", now),
				qm.Where("security_groups.updated_at <= ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.WhereIn("security_groups.name IN ?", "test-sg-1", "test-sg-2"),
				qm.OrderBy("updated_at DESC"),
				qm.Where("security_groups.created_at > ?", now),
				qm.Where("security_groups.updated_at <= ?", now),
			},
		},
	}
	for testCaseName, testCase := range cases {
		s.Run(testCaseName, func() {
			s.SetupTest() // needed to ensure mocks are fresh for every test case
			s.req.URL.RawQuery = testCase.query
			s.querier.EXPECT().Count(gomock.Any(), gomock.Any()).AnyTimes().Return(int64(50), nil)
			s.querier.EXPECT().All(gomock.Any(), gomock.Any()).AnyTimes().Return(models.SecurityGroupSlice{{Name: "test-security-group"}}, nil)
			s.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)

			err := s.controller.List(s.ctx)
			if testCase.expectedErr != nil {
				var ccErr *v3.CfAPIError
				s.ErrorAs(err, &ccErr)
				s.Equal(testCase.expectedErr.HTTPStatus, ccErr.HTTPStatus)
				return
			}
			s.NoError(err)
			s.querierFunc.AssertNumberOfCalls(s.T(), "Get", 2)

			countQueryMods := s.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
			s.ElementsMatch(testCase.expectedCountFilters, countQueryMods)

			allQueryMods := s.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)
			expectedAllQueryMods := append(baseQueryMods, testCase.expectedAllFilters...) //nolint:gocritic // Deliberately appending to a different slice
			s.ElementsMatch(expectedAllQueryMods, allQueryMods)
		})
	}
}

func (s *ListSecurityGroupTestSuite) TestInternalServerError() {
	s.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("something went wrong"))
	s.Error(v3.UnknownError(nil), s.controller.List(s.ctx))
}
