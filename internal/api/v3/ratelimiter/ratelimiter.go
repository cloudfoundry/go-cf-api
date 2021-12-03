package ratelimiter

import (
	"math"
	"strconv"
	"sync"
	"time"
)

type RateLimiter interface {
	Check(identifier string) (bool, map[string]string, error)
	Increment(identifier string)
}

type rateLimiter struct {
	mutex           sync.Mutex
	store           map[string]*record
	globalLimit     int
	perProcessLimit int
	resetInterval   ResetInterval
}

func NewRateLimiter(globalLimit, perProcessLimit int, resetInterval ResetInterval) RateLimiter {
	return &rateLimiter{
		mutex:           sync.Mutex{},
		store:           map[string]*record{},
		globalLimit:     globalLimit,
		perProcessLimit: perProcessLimit,
		resetInterval:   resetInterval,
	}
}

//nolint:gochecknoglobals // here to be overridden in tests
var now = time.Now

type record struct {
	count      int
	validUntil time.Time
}

func (r *rateLimiter) Check(identifier string) (bool, map[string]string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	storedRecord, ok := r.store[identifier]
	if !ok || now().After(storedRecord.validUntil) {
		storedRecord = &record{0, r.resetInterval.Next(identifier)}
		r.store[identifier] = storedRecord
	}

	headers := map[string]string{
		"X-RateLimit-Limit":     strconv.Itoa(r.globalLimit),
		"X-RateLimit-Reset":     strconv.FormatInt(storedRecord.validUntil.Unix(), 10), //nolint:gomnd
		"X-RateLimit-Remaining": strconv.Itoa(r.estimateRemaining(storedRecord.count)),
	}

	return (r.perProcessLimit - storedRecord.count) > 0, headers, nil
}

func (r *rateLimiter) estimateRemaining(count int) int {
	fraction := math.Round(float64(r.perProcessLimit-count-1)/float64(r.perProcessLimit)*10.0) / 10 //nolint:gomnd
	return int(math.Max(0.0, fraction*float64(r.globalLimit)))
}

func (r *rateLimiter) Increment(identifier string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	stored, ok := r.store[identifier]
	if !ok {
		r.store[identifier] = &record{1, r.resetInterval.Next(identifier)}
	} else {
		stored.count++
	}
}
