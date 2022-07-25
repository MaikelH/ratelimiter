package ratelimiter

import "time"

type RateLimiter interface {
	GetAvailablePermits() (int, error)
	Acquire() (RateLimitLease, error)
}

type ReplenishingRateLimiter interface {
	ReplenishmentPeriod() (time.Duration, error)
	IsAutoReplenishing() (bool, error)
	TryReplenish() (bool, error)
}
