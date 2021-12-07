//go:build unit
// +build unit

//nolint:dupl // Similar tests shouldn't be marked as duplicate
package auth_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/auth"
)

type RequiresScopeMiddlewareSuite struct {
	suite.Suite
	handler HandlerFunc
	ctx     echo.Context
}

func TestRequiresScopeMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(RequiresScopeMiddlewareSuite))
}

func (s *RequiresScopeMiddlewareSuite) SetupTest() {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/authed", nil)
	recorder := httptest.NewRecorder()
	s.handler = HandlerFunc{}
	s.ctx = e.NewContext(request, recorder)
}

func (s *RequiresScopeMiddlewareSuite) TestReadMiddleware() {
	cases := map[Scope]bool{
		Read:            true,
		Admin:           true,
		AdminReadOnly:   true,
		GlobalAuditor:   true,
		Write:           false,
		"unknown.scope": false,
		"":              false,
	}

	readMiddleware := NewRequiresReadMiddleware()
	for scope, shouldBeAbleToRead := range cases {
		s.Run(fmt.Sprintf("user with %s scope", scope), func() {
			s.ctx.Set(Scopes, []string{string(scope)})
			s.handler.On("Next", s.ctx).Return(nil)

			err := readMiddleware(s.handler.Next)(s.ctx)
			if shouldBeAbleToRead {
				s.NoError(err)
				s.handler.AssertCalled(s.T(), "Next", s.ctx)
			} else {
				var ccErr *v3.CfAPIError
				s.ErrorAs(err, &ccErr)
				s.Equal(http.StatusForbidden, ccErr.HTTPStatus)
				s.Equal("user lacks sufficient scope to read resource", ccErr.Err.Error())
				s.handler.AssertNotCalled(s.T(), "Next")
			}
		})
	}
}

func (s *RequiresScopeMiddlewareSuite) TestWriteMiddleware() {
	cases := map[Scope]bool{
		Read:            false,
		Admin:           true,
		AdminReadOnly:   false,
		GlobalAuditor:   false,
		Write:           true,
		"unknown.scope": false,
		"":              false,
	}

	writeMiddleware := NewRequiresWriteMiddleware()
	for scope, shouldBeAbleToWrite := range cases {
		s.Run(fmt.Sprintf("user with %s scope", scope), func() {
			s.ctx.Set(Scopes, []string{string(scope)})
			s.handler.On("Next", s.ctx).Return(nil)

			err := writeMiddleware(s.handler.Next)(s.ctx)
			if shouldBeAbleToWrite {
				s.NoError(err)
				s.handler.AssertCalled(s.T(), "Next", s.ctx)
			} else {
				var ccErr *v3.CfAPIError
				s.ErrorAs(err, &ccErr)
				s.Equal(http.StatusForbidden, ccErr.HTTPStatus)
				s.Equal("user lacks sufficient scope to write resource", ccErr.Err.Error())
				s.handler.AssertNotCalled(s.T(), "Next")
			}
		})
	}
}

func (s *RequiresScopeMiddlewareSuite) TestMiddlewaresReturnErrorsFromNextMiddleware() {
	nextErr := errors.New("something went wrong in future middleware")
	readMiddleware := NewRequiresReadMiddleware()
	writeMiddleware := NewRequiresWriteMiddleware()
	s.ctx.Set(Scopes, []string{string(Admin)})
	s.handler.On("Next", s.ctx).Return(nextErr)

	readErr := readMiddleware(s.handler.Next)(s.ctx)
	s.Equal(nextErr, readErr)
	writeErr := writeMiddleware(s.handler.Next)(s.ctx)
	s.Equal(nextErr, writeErr)
}

type HandlerFunc struct {
	mock.Mock
}

func (m *HandlerFunc) Next(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
