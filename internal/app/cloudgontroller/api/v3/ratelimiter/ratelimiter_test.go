// +build unit

package ratelimiter //nolint:testpackage // so we can override now function variable

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	userGUID        = "user-guid"
	perProcessLimit = 10
	globalLimit     = 100
)

type RateLimiterSuite struct {
	suite.Suite
	frozenTime    time.Time
	rateLimiter   RateLimiter
	resetInterval *MockResetInterval
}

func TestRateLimiterSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(RateLimiterSuite))
}

type MockResetInterval struct {
	mock.Mock
}

func (m *MockResetInterval) Next(identifier string) time.Time {
	args := m.Called(identifier)
	return args.Get(0).(time.Time)
}

func (s *RateLimiterSuite) SetupTest() {
	s.frozenTime = time.Now()
	now = func() time.Time { return s.frozenTime }
	s.resetInterval = &MockResetInterval{}
	s.rateLimiter = NewRateLimiter(globalLimit, perProcessLimit, s.resetInterval)
}

func (s *RateLimiterSuite) TestNewUserIsAllowed() {
	s.resetInterval.On("Next", mock.Anything).Return(s.frozenTime.Add(time.Hour))
	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.True(allowed)
}

func (s *RateLimiterSuite) TestUserIsAllowedWhenBelowLimit() {
	s.resetInterval.On("Next", mock.Anything).Return(s.frozenTime.Add(time.Hour))
	for i := 0; i < perProcessLimit-1; i++ {
		s.rateLimiter.Increment(userGUID)
	}

	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.True(allowed)
}

func (s *RateLimiterSuite) TestUserIsDeniedWhenLimitIsReached() {
	s.resetInterval.On("Next", mock.Anything).Return(s.frozenTime.Add(time.Hour))
	for i := 0; i < perProcessLimit; i++ {
		s.rateLimiter.Increment(userGUID)
	}

	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.False(allowed)
}

func (s *RateLimiterSuite) TestUserIsAllowedWhenCountHasExpired() {
	interval := 5 * time.Minute
	s.resetInterval.On("Next", mock.Anything).Return(s.frozenTime.Add(interval))
	for i := 0; i < perProcessLimit; i++ {
		s.rateLimiter.Increment(userGUID)
	}

	newTime := s.frozenTime.Add(interval).Add(time.Minute)
	now = func() time.Time { return newTime }
	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.True(allowed)
}

func (s *RateLimiterSuite) TestCorrectHeadersAreReturnedForNewUser() {
	expectedTime := s.frozenTime.Add(5 * time.Minute)
	s.resetInterval.On("Next", mock.Anything).Return(expectedTime)
	_, headers, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.Equal(strconv.Itoa(globalLimit), headers["X-RateLimit-Limit"])
	s.Equal(strconv.Itoa(90), headers["X-RateLimit-Remaining"])
	s.Equal(strconv.FormatInt(expectedTime.Unix(), 10), headers["X-RateLimit-Reset"])
}

func (s *RateLimiterSuite) TestCorrectHeadersAreReturnedForExistingUserWithinValidityPeriod() {
	expectedTime := s.frozenTime.Add(5 * time.Minute)
	s.resetInterval.On("Next", mock.Anything).Return(expectedTime)
	s.rateLimiter.Increment(userGUID)

	newTime := expectedTime.Add(-1 * time.Minute)
	now = func() time.Time { return newTime }

	_, headers, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.Equal(strconv.Itoa(globalLimit), headers["X-RateLimit-Limit"])
	s.Equal(strconv.Itoa(80), headers["X-RateLimit-Remaining"])
	s.Equal(strconv.FormatInt(expectedTime.Unix(), 10), headers["X-RateLimit-Reset"])
}

func (s *RateLimiterSuite) TestCorrectHeadersAreReturnedForExistingUserAfterValidityPeriod() {
	firstInterval := s.frozenTime.Add(5 * time.Minute)
	secondInterval := s.frozenTime.Add(5 * time.Minute)
	s.resetInterval.On("Next", mock.Anything).Return(firstInterval)
	s.rateLimiter.Increment(userGUID)

	s.resetInterval.On("Next", mock.Anything).Return(secondInterval)
	newTime := firstInterval.Add(time.Minute)
	now = func() time.Time { return newTime }

	_, headers, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.Equal(strconv.Itoa(globalLimit), headers["X-RateLimit-Limit"])
	s.Equal(strconv.Itoa(90), headers["X-RateLimit-Remaining"])
	s.Equal(strconv.FormatInt(secondInterval.Unix(), 10), headers["X-RateLimit-Reset"])
}
