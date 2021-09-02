// +build unit

package ratelimiter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	mock_models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/mocks"
)

func (suite *RateLimiterSuite) TestNoRequestCountExists() {
	suite.Finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.RequestCountSlice{}, nil)
	userGUID := "123"
	expected := &models.RequestCount{
		UserGUID:   null.StringFrom(userGUID),
		Count:      null.IntFrom(1),
		ValidUntil: null.TimeFrom(suite.Now.Add(suite.RateLimiter.ResetInterval)),
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
	expected.ValidUntil = null.TimeFrom(suite.Now.Add(suite.RateLimiter.ResetInterval))
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
	queriers = RateLimiterQueriers{
		Finisher: func(mods ...qm.QueryMod) models.RequestCountFinisher { return suite.Finisher },
		Inserter: suite.Inserter,
		Updater:  suite.Updater,
	}

	suite.Now = time.Now().UTC()
	now = func() time.Time { return suite.Now }

	suite.RateLimiter = RateLimiter{
		GeneralLimit:  10,
		ResetInterval: time.Minute * 1,
		DB:            nil,
	}
}

func (suite *RateLimiterSuite) After(suiteName, testName string) {
	suite.ctrl.Finish()
}

type RateLimiterSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	Context     echo.Context
	Finisher    *mock_models.MockRequestCountFinisher
	Inserter    *mock_models.MockRequestCountInserter
	Updater     *mock_models.MockRequestCountUpdater
	Now         time.Time
	RateLimiter RateLimiter
}

func TestRateLimiterSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(RateLimiterSuite))
}
