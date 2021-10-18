package ratelimiter

import (
	"hash/fnv"
	"time"
)

func UserHashOffset(identifier string, interval time.Duration) time.Duration {
	hash := fnv.New64()
	hash.Write([]byte(identifier))
	remainder := hash.Sum64() % uint64(interval.Seconds())
	return time.Duration(remainder * uint64(time.Second))
}
