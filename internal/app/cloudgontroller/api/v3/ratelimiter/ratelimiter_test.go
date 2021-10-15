// +build unit

package ratelimiter //nolint:testpackage // so we can override now function variable

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	userGUID      = "user-guid"
	limit         = 10
	resetInterval = 5 * time.Minute
)

type RateLimiterSuite struct {
	suite.Suite
	frozenTime  time.Time
	rateLimiter RateLimiter
}

func TestRateLimiterSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(RateLimiterSuite))
}

func (s *RateLimiterSuite) SetupTest() {
	s.frozenTime = time.Now()
	now = func() time.Time { return s.frozenTime }
	s.rateLimiter = NewRateLimiter(limit, resetInterval)
}

func (s *RateLimiterSuite) TestNewUserIsAllowed() {
	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.True(allowed)
}

func (s *RateLimiterSuite) TestUserIsAllowedWhenBelowLimit() {
	for i := 0; i < limit-1; i++ {
		s.rateLimiter.Increment(userGUID)
	}

	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.True(allowed)
}

func (s *RateLimiterSuite) TestUserIsDeniedWhenLimitIsReached() {
	for i := 0; i < limit; i++ {
		s.rateLimiter.Increment(userGUID)
	}

	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.False(allowed)
}

func (s *RateLimiterSuite) TestUserIsAllowedWhenCountHasExpired() {
	for i := 0; i < limit; i++ {
		s.rateLimiter.Increment(userGUID)
	}

	newTime := s.frozenTime.Add(resetInterval).Add(time.Minute)
	now = func() time.Time { return newTime }
	allowed, _, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.True(allowed)
}

func (s *RateLimiterSuite) TestCorrectHeadersAreReturnedForNewUser() {
	_, headers, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.Equal(strconv.FormatInt(limit, 10), headers["X-RateLimit-Limit"])
	s.Equal(strconv.FormatInt(limit-1, 10), headers["X-RateLimit-Remaining"])
	s.Equal(strconv.FormatInt(s.frozenTime.Add(resetInterval).Unix(), 10), headers["X-RateLimit-Reset"])
}

func (s *RateLimiterSuite) TestCorrectHeadersAreReturnedForExistingUserWithinValidityPeriod() {
	s.rateLimiter.Increment(userGUID)
	newTime := s.frozenTime.Add(time.Minute)
	now = func() time.Time { return newTime }
	_, headers, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.Equal(strconv.FormatInt(limit, 10), headers["X-RateLimit-Limit"])
	s.Equal(strconv.FormatInt(limit-2, 10), headers["X-RateLimit-Remaining"])
	s.Equal(strconv.FormatInt(s.frozenTime.Add(resetInterval).Unix(), 10), headers["X-RateLimit-Reset"])
}

func (s *RateLimiterSuite) TestCorrectHeadersAreReturnedForExistingUserAfterValidityPeriod() {
	s.rateLimiter.Increment(userGUID)
	newTime := s.frozenTime.Add(resetInterval).Add(time.Minute)
	now = func() time.Time { return newTime }
	_, headers, err := s.rateLimiter.Check(userGUID)
	s.NoError(err)
	s.Equal(strconv.FormatInt(limit, 10), headers["X-RateLimit-Limit"])
	s.Equal(strconv.FormatInt(limit-1, 10), headers["X-RateLimit-Remaining"])
	s.Equal(strconv.FormatInt(newTime.Add(resetInterval).Unix(), 10), headers["X-RateLimit-Reset"])
}
