package ratelimiter

import "time"

type RateLimiter interface {
	GetAvailablePermits() (int, error)
	Acquire(permitCount int) (RateLimitLease, error)
}

type ReplenishingRateLimiter interface {
	ReplenishmentPeriod() time.Duration
	IsAutoReplenishing() bool
	TryReplenish() (bool, error)
}
