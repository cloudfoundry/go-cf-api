//go:build unit
// +build unit

package auth_test

import (
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite

	context        echo.Context
	successHandler middleware.JWTSuccessHandler
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) SetupTest() {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/authed", nil)
	recorder := httptest.NewRecorder()
	s.context = e.NewContext(request, recorder)
	s.successHandler = NewSuccessHandler()
}

func (s *AuthTestSuite) TestSuccessHandler() {
	type expectedFunc func(c echo.Context) bool
	cases := map[string]struct {
		scopes            []string
		expectedToBeTrue  []expectedFunc
		expectedToBeFalse []expectedFunc
	}{
		"no extra scopes": {
			scopes:            []string{},
			expectedToBeFalse: []expectedFunc{CanRead, CanWrite, IsAdmin, IsAdminReadOnly, IsGlobalAuditor},
		},
		"cloud_controller.read scope": {
			scopes:            []string{"cloud_controller.read"},
			expectedToBeTrue:  []expectedFunc{CanRead},
			expectedToBeFalse: []expectedFunc{CanWrite, IsAdmin, IsAdminReadOnly, IsGlobalAuditor},
		},
		"cloud_controller.write scope": {
			scopes:            []string{"cloud_controller.write"},
			expectedToBeTrue:  []expectedFunc{CanWrite},
			expectedToBeFalse: []expectedFunc{CanRead, IsAdmin, IsAdminReadOnly, IsGlobalAuditor},
		},
		"cloud_controller.admin scope": {
			scopes:            []string{"cloud_controller.admin"},
			expectedToBeTrue:  []expectedFunc{IsAdmin},
			expectedToBeFalse: []expectedFunc{CanRead, CanWrite, IsAdminReadOnly, IsGlobalAuditor},
		},
		"cloud_controller.admin_read_only scope": {
			scopes:            []string{"cloud_controller.admin_read_only"},
			expectedToBeTrue:  []expectedFunc{IsAdminReadOnly},
			expectedToBeFalse: []expectedFunc{CanRead, CanWrite, IsAdmin, IsGlobalAuditor},
		},
		"cloud_controller.global_auditor scope": {
			scopes:            []string{"cloud_controller.global_auditor"},
			expectedToBeTrue:  []expectedFunc{IsGlobalAuditor},
			expectedToBeFalse: []expectedFunc{CanRead, CanWrite, IsAdmin, IsAdminReadOnly},
		},
		"all scopes": {
			scopes: []string{
				"cloud_controller.read",
				"cloud_controller.write",
				"cloud_controller.admin",
				"cloud_controller.admin_read_only",
				"cloud_controller.global_auditor",
			},
			expectedToBeTrue: []expectedFunc{CanRead, CanWrite, IsAdmin, IsAdminReadOnly, IsGlobalAuditor},
		},
	}
	for name, tc := range cases {
		s.Run(name, func() {
			token := &jwt.Token{
				Claims: &CFClaims{
					UserID: name,
					Scopes: tc.scopes,
				},
			}
			s.context.Set("user", token)
			s.successHandler(s.context)
			s.Equal(name, s.context.Get(Username).(string))
			for _, expectedToBeTrue := range tc.expectedToBeTrue {
				s.True(expectedToBeTrue(s.context))
			}
			for _, expectedToBeFalse := range tc.expectedToBeFalse {
				s.False(expectedToBeFalse(s.context))
			}
		})
	}
}
