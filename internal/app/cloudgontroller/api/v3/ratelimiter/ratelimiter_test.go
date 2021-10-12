// +build unit

package ratelimiter //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/auth"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	mock_models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/mocks"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/testutils"
)

func (suite *RateLimiterSuite) TestNoRequestCountExists() {
	suite.Finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.RequestCountSlice{}, nil)
	userGUID := "123"
	expected := &models.RequestCount{
		UserGUID:   null.StringFrom(userGUID),
		Count:      null.IntFrom(1),
		ValidUntil: null.TimeFrom(suite.Now.Add(suite.ResetInterval)),
	}
	suite.Inserter.EXPECT().Insert(gomock.Eq(expected), gomock.Any(), gomock.Any(), boil.Infer()).Return(nil)

	result, err := suite.RateLimiter.Allow(userGUID, suite.Context)
	suite.NoError(err)
	suite.True(result)
}

//nolint:dupl // test wrongly recognised as a duplicate
func (suite *RateLimiterSuite) TestRequestCountExceedsRateLimit() {
	expiryTime := suite.Now.Add(1 * time.Hour)

	userGUID := "123"
	rc := &models.RequestCount{
		ID:         1,
		UserGUID:   null.StringFrom(userGUID),
		Count:      null.IntFrom(10),
		ValidUntil: null.TimeFrom(expiryTime),
	}
	suite.Finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.RequestCountSlice{rc}, nil)
	expected := rc
	expected.Count = null.IntFrom(11)
	suite.Updater.EXPECT().Update(gomock.Eq(expected), gomock.Any(), gomock.Any(), boil.Infer()).Return(int64(1), nil)

	result, err := suite.RateLimiter.Allow(userGUID, suite.Context)
	suite.NoError(err)
	suite.False(result)
}

//nolint:dupl // test wrongly recognised as a duplicate
func (suite *RateLimiterSuite) TestRequestCountWithinRateLimit() {
	expiryTime := suite.Now.Add(1 * time.Hour)

	userGUID := "123"
	rc := &models.RequestCount{
		ID:         1,
		UserGUID:   null.StringFrom(userGUID),
		Count:      null.IntFrom(3),
		ValidUntil: null.TimeFrom(expiryTime),
	}
	suite.Finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.RequestCountSlice{rc}, nil)
	expected := rc
	expected.Count = null.IntFrom(4)
	suite.Updater.EXPECT().Update(gomock.Eq(expected), gomock.Any(), gomock.Any(), boil.Infer()).Return(int64(1), nil)

	result, err := suite.RateLimiter.Allow(userGUID, suite.Context)
	suite.NoError(err)
	suite.True(result)
}

// When we're above the rate limit but its reset time is in the past
func (suite *RateLimiterSuite) TestRequestCountExpired() {
	expiryTime := suite.Now.Add(-1 * time.Hour)

	userGUID := "123"
	rc := &models.RequestCount{
		ID:         1,
		UserGUID:   null.StringFrom(userGUID),
		Count:      null.IntFrom(3),
		ValidUntil: null.TimeFrom(expiryTime),
	}
	suite.Finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.RequestCountSlice{rc}, nil)
	expected := rc
	expected.Count = null.IntFrom(1)
	expected.ValidUntil = null.TimeFrom(suite.Now.Add(suite.ResetInterval))
	suite.Updater.EXPECT().Update(gomock.Eq(expected), gomock.Any(), gomock.Any(), boil.Infer()).Return(int64(1), nil)

	result, err := suite.RateLimiter.Allow(userGUID, suite.Context)
	suite.NoError(err)
	suite.True(result)
}

func (suite *RateLimiterSuite) SetupSuite() {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	suite.Context = e.NewContext(request, recorder)

	suite.ctrl = gomock.NewController(suite.T())
	suite.Finisher = mock_models.NewMockRequestCountFinisher(suite.ctrl)
	suite.Inserter = mock_models.NewMockRequestCountInserter(suite.ctrl)
	suite.Updater = mock_models.NewMockRequestCountUpdater(suite.ctrl)
	queriers = Queriers{
		Finisher: func(mods ...qm.QueryMod) models.RequestCountFinisher { return suite.Finisher },
		Inserter: suite.Inserter,
		Updater:  suite.Updater,
	}

	suite.Now = time.Now().UTC()
	now = func() time.Time { return suite.Now }

	suite.ResetInterval = time.Minute
	suite.RateLimiter = NewRateLimiter(nil, 10, suite.ResetInterval)
}

func (suite *RateLimiterSuite) After(suiteName, testName string) {
	suite.ctrl.Finish()
}

type RateLimiterSuite struct {
	suite.Suite
	ctrl          *gomock.Controller
	Context       echo.Context
	Finisher      *mock_models.MockRequestCountFinisher
	Inserter      *mock_models.MockRequestCountInserter
	Updater       *mock_models.MockRequestCountUpdater
	Now           time.Time
	RateLimiter   RateLimiter
	ResetInterval time.Duration
}

func TestRateLimiterSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(RateLimiterSuite))
}

type RateLimiterMiddlewareSuite struct {
	suite.Suite
	ctx         echo.Context
	rateLimiter *MockRateLimiter
	handler     *testutils.HandlerFunc
	middleware  echo.MiddlewareFunc
}

func TestRateLimiterMiddlewareSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(RateLimiterMiddlewareSuite))
}

func (s *RateLimiterMiddlewareSuite) SetupTest() {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	s.ctx = e.NewContext(request, recorder)
	s.ctx.Set(auth.Username, "user-guid")

	s.rateLimiter = &MockRateLimiter{}
	s.handler = &testutils.HandlerFunc{}
	s.middleware = CustomRateLimiter(s.rateLimiter)
}

func (s *RateLimiterMiddlewareSuite) TestAllowsRegularUser() {
	s.rateLimiter.On("Allow", mock.Anything, mock.Anything).Return(true, nil)
	s.handler.On("Next", mock.Anything).Return(nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	s.NoError(err)
	s.handler.AssertCalled(s.T(), "Next", s.ctx)
}

func (s *RateLimiterMiddlewareSuite) TestDeniesRegularUser() {
	s.rateLimiter.On("Allow", mock.Anything, mock.Anything).Return(false, nil)

	err := s.middleware(s.handler.Next)(s.ctx)
	var ccErr *v3.CloudControllerError
	s.ErrorAs(err, &ccErr)
	s.Equal(http.StatusTooManyRequests, ccErr.HTTPStatus)
	s.handler.AssertNotCalled(s.T(), "Next", s.ctx)
}

type MockRateLimiter struct {
	mock.Mock
}

func (m *MockRateLimiter) Allow(identifier string, ctx echo.Context) (bool, error) {
	args := m.Called(identifier, ctx)
	return args.Bool(0), args.Error(1)
}
