package ratelimiter

import "time"

type RateLimitLease interface {
	IsAcquired() bool
	Release()
}

type FixedWindowLease struct {
	isAcquired bool
	retryAfter time.Duration
}

func (f *FixedWindowLease) Release() {
	return
}

func (f *FixedWindowLease) IsAcquired() bool {
	return f.isAcquired
}
