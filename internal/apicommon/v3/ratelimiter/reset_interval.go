package ratelimiter

import (
	"hash/fnv"
	"math"
	"time"
)

type ResetInterval interface {
	Next(identifier string) time.Time
}

type userHashResetInterval struct {
	interval time.Duration
	now      func() time.Time
}

func NewUserHashResetInterval(interval time.Duration, now func() time.Time) ResetInterval {
	return &userHashResetInterval{
		interval: interval,
		now:      now,
	}
}

func (u *userHashResetInterval) Next(identifier string) time.Time {
	hash := fnv.New64()
	hash.Write([]byte(identifier))
	remainder := hash.Sum64() % uint64(u.interval.Seconds())
	offset := time.Duration(remainder * uint64(time.Second))
	numberOfIntervals := math.Floor(float64(u.now().Add(-1*offset).Unix())/u.interval.Seconds()) + 1
	return time.Unix(int64(numberOfIntervals*u.interval.Seconds()), 0).Add(offset)
}
