package ratelimiter

import (
	"time"
)

type FixedWindowRateLimiter struct {
	requestCount       int
	isAutoReplenishing bool
	window             time.Duration
}

type FixedWindowRateLimiterOptions struct {
	PermitLimit          int
	QueueLimit           int
	QueueProcessingOrder QueueProcessingOrder
	Window               time.Duration
	AutoReplenishment    bool
}

func NewFixedWindowRateLimiter(options FixedWindowRateLimiterOptions) (*FixedWindowRateLimiter, error) {
	if options.PermitLimit < 0 {
		return nil, &ArgumentError{Message: "permit limit must be higher then 0"}
	}
	if options.QueueLimit < 0 {
		return nil, &ArgumentError{Message: "queue limit must be higher then 0"}
	}

	limiter := FixedWindowRateLimiter{
		requestCount:       options.PermitLimit,
		isAutoReplenishing: options.AutoReplenishment,
		window:             options.Window,
	}

	if options.AutoReplenishment {
		// TODO: set timer to automatically replenish
	}

	return &limiter, nil
}

func (f *FixedWindowRateLimiter) GetAvailablePermits() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FixedWindowRateLimiter) Acquire(permitCount int) (RateLimitLease, error) {
	if permitCount < 0 {
		return nil, &ArgumentError{Message: "permit count must be higher than 0"}
	}

	// Amount of permits is so high that they can never be acquired
	if permitCount > f.requestCount {
		return nil, &ArgumentError{Message: "permit count is higher then maximum number of permits available"}
	}

	//TODO implement me
	panic("implement me")
}

func (f *FixedWindowRateLimiter) ReplenishmentPeriod() time.Duration {
	return f.window
}

func (f *FixedWindowRateLimiter) IsAutoReplenishing() bool {
	return f.isAutoReplenishing
}

func (f *FixedWindowRateLimiter) TryReplenish() (bool, error) {
	//TODO implement me
	panic("implement me")
}

// Force FixedWindowRateLimiter to adhere to RateLimiter and ReplenishingRateLimiter interface
var (
	_ RateLimiter             = &FixedWindowRateLimiter{}
	_ ReplenishingRateLimiter = &FixedWindowRateLimiter{}
)
