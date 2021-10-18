// +build unit

package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/suite"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/ratelimiter"
)

const (
	guid       = "a-guid"
	guidOffset = 2452 * time.Second
)

type UserHashIntervalSuite struct {
	suite.Suite
	startOfHour   time.Time
	interval      time.Duration
	now           func() time.Time
	resetInterval ResetInterval
}

func TestUserHashIntervalSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserHashIntervalSuite))
}

func (s *UserHashIntervalSuite) SetupTest() {
	s.interval = time.Hour
	s.startOfHour = time.Now().Truncate(time.Hour)
	s.now = func() time.Time { return s.startOfHour }
	s.resetInterval = NewUserHashResetInterval(s.interval, s.now)
}

func (s *UserHashIntervalSuite) TestReturnsNextTimeIntervalIncludingOffset() {
	actual := s.resetInterval.Next(guid)
	s.Equal(guidOffset, actual.Sub(s.startOfHour))
}

func (s *UserHashIntervalSuite) TestReturnsNextTimeIntervalWhenAfterOffset() {
	now := func() time.Time { return s.startOfHour.Add(guidOffset).Add(time.Minute) }
	resetInterval := NewUserHashResetInterval(s.interval, now)
	actual := resetInterval.Next(guid)
	s.Equal(s.interval+guidOffset, actual.Sub(s.startOfHour))
}

func (s *UserHashIntervalSuite) TestProducesOffsetsWithinInterval() {
	for i := 0; i < 1000; i++ {
		guid, err := uuid.NewV4()
		s.NoError(err)
		next := s.resetInterval.Next(guid.String())
		s.WithinDuration(s.startOfHour, next, s.interval)
	}
}

func (s *UserHashIntervalSuite) TestProducesSameIntervalForSameInputIdentifier() {
	first := s.resetInterval.Next(guid)
	second := s.resetInterval.Next(guid)
	s.Equal(first, second)
}
