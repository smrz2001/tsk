package ratelimiter

import (
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	failIfNoPanic := func(f func()) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("")
			}
		}()

		f()
	}

	opts := RateLimiterOpts{Limit: -1, Burst: 1}
	failIfNoPanic(opts.validate)

	opts = RateLimiterOpts{Limit: Every(10 * time.Millisecond), Burst: 0}
	failIfNoPanic(opts.validate)

	opts = RateLimiterOpts{Limit: Every(10 * time.Millisecond), Burst: 1, MaxQueueDepth: -1}
	failIfNoPanic(opts.validate)
}
