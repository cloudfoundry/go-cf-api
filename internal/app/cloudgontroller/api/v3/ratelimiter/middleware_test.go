// +build unit

package ratelimiter_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/ratelimiter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/auth"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/testutils"
)

const (
	userGUID = "user-guid"
	ip       = "1.2.3.4"
)

type RateLimiterMiddlewareSuite struct {
	suite.Suite
	ctx                        echo.Context
	recorder                   *httptest.ResponseRecorder
	generalRateLimiter         *MockRateLimiter
	unauthenticatedRateLimiter *MockRateLimiter
	handler                    *testutils.HandlerFunc
	middleware                 echo.MiddlewareFunc
}

func TestRateLimiterMiddlewareSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(RateLimiterMiddlewareSuite))
}

func (s *RateLimiterMiddlewareSuite) SetupTest() {
	e := echo.New()
	e.IPExtractor = func(*http.Request) string { return ip }
	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	s.recorder = httptest.NewRecorder()
	s.ctx = e.NewContext(request, s.recorder)

	s.generalRateLimiter = &MockRateLimiter{}
	s.unauthenticatedRateLimiter = &MockRateLimiter{}
	s.handler = &testutils.HandlerFunc{}
	s.middleware = NewRateLimiterMiddleware(s.generalRateLimiter, s.unauthenticatedRateLimiter)
}

func (s *RateLimiterMiddlewareSuite) TestCallsNextMiddlewareAndIncrementsWhenRegularUserIsAllowed() {
	s.ctx.Set(auth.Username, userGUID)
	s.generalRateLimiter.On("Check", userGUID).Return(true, map[string]string{}, nil)
	s.generalRateLimiter.On("Increment", mock.Anything)
	s.handler.On("Next", mock.Anything).Return(nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	s.NoError(err)
	s.handler.AssertCalled(s.T(), "Next", s.ctx)
	s.generalRateLimiter.AssertCalled(s.T(), "Increment", userGUID)
}

func (s *RateLimiterMiddlewareSuite) TestDoesNotCallNextMiddlewareOrIncrementWhenRegularUserIsDenied() {
	s.ctx.Set(auth.Username, userGUID)
	s.generalRateLimiter.On("Check", userGUID).Return(false, map[string]string{}, nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	var ccErr *v3.CloudControllerError
	s.ErrorAs(err, &ccErr)
	s.Equal(http.StatusTooManyRequests, ccErr.HTTPStatus)
	s.handler.AssertNotCalled(s.T(), "Next")
	s.generalRateLimiter.AssertNotCalled(s.T(), "Increment")
}

func (s *RateLimiterMiddlewareSuite) TestAddsHeadersFromGeneralLimiter() {
	headers := map[string]string{"X-Foo": "bar", "X-Bar": "blah"}
	s.ctx.Set(auth.Username, userGUID)
	s.generalRateLimiter.On("Check", userGUID).Return(true, headers, nil)
	s.generalRateLimiter.On("Increment", mock.Anything)
	s.handler.On("Next", mock.Anything).Return(nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	s.NoError(err)
	s.Equal(s.recorder.Header().Get("X-Foo"), "bar")
	s.Equal(s.recorder.Header().Get("X-Bar"), "blah")
}

func (s *RateLimiterMiddlewareSuite) TestCallsNextMiddlewareAndIncremementWhenUnauthenticatedUserIsAllowed() {
	s.unauthenticatedRateLimiter.On("Check", ip).Return(true, map[string]string{}, nil)
	s.unauthenticatedRateLimiter.On("Increment", mock.Anything)
	s.handler.On("Next", mock.Anything).Return(nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	s.NoError(err)
	s.handler.AssertCalled(s.T(), "Next", s.ctx)
	s.unauthenticatedRateLimiter.AssertCalled(s.T(), "Increment", ip)
}

func (s *RateLimiterMiddlewareSuite) TestDoesNotCallNextMiddlewareWhenUnauthenticatedUserIsDenied() {
	s.unauthenticatedRateLimiter.On("Check", ip).Return(false, map[string]string{}, nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	var ccErr *v3.CloudControllerError
	s.ErrorAs(err, &ccErr)
	s.Equal(http.StatusTooManyRequests, ccErr.HTTPStatus)
	s.handler.AssertNotCalled(s.T(), "Next")
	s.unauthenticatedRateLimiter.AssertNotCalled(s.T(), "Increment")
}

func (s *RateLimiterMiddlewareSuite) TestAddsHeadersFromUnauthenticatedLimiter() {
	headers := map[string]string{"X-Foo": "bar", "X-Bar": "blah"}
	s.unauthenticatedRateLimiter.On("Check", ip).Return(true, headers, nil)
	s.unauthenticatedRateLimiter.On("Increment", mock.Anything)
	s.handler.On("Next", mock.Anything).Return(nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	s.NoError(err)
	s.Equal(s.recorder.Header().Get("X-Foo"), "bar")
	s.Equal(s.recorder.Header().Get("X-Bar"), "blah")
}

func (s *RateLimiterMiddlewareSuite) TestAddsRetryAfterHeaderWhenUserIsDenied() {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	headers := map[string]string{"X-RateLimit-Reset": now}
	s.unauthenticatedRateLimiter.On("Check", ip).Return(false, headers, nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	s.Error(err)
	s.Equal(now, s.recorder.Header().Get("Retry-After"))
}

func (s *RateLimiterMiddlewareSuite) TestDoesNotCheckRateLimitersWhenUserIsAdmin() {
	s.ctx.Set(auth.Scopes, []string{string(auth.Admin)})
	s.handler.On("Next", mock.Anything).Return(nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	s.NoError(err)
	s.handler.AssertCalled(s.T(), "Next", s.ctx)
	s.generalRateLimiter.AssertNotCalled(s.T(), "Check", mock.Anything)
	s.generalRateLimiter.AssertNotCalled(s.T(), "Increment", mock.Anything)
	s.unauthenticatedRateLimiter.AssertNotCalled(s.T(), "Check", mock.Anything)
	s.unauthenticatedRateLimiter.AssertNotCalled(s.T(), "Increment", mock.Anything)
}

type MockRateLimiter struct {
	mock.Mock
}

func (m *MockRateLimiter) Check(identifier string) (bool, map[string]string, error) {
	args := m.Called(identifier)
	return args.Bool(0), args.Get(1).(map[string]string), args.Error(2)
}

func (m *MockRateLimiter) Increment(identifier string) {
	m.Called(identifier)
}
