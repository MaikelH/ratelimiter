package ratelimiter

import (
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	requestCount       int
	permitLimit        int
	isAutoReplenishing bool
	window             time.Duration
	mutex              sync.Mutex
	idleSince          time.Duration
	queue              Queue[RequestRegistration]
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
		permitLimit:        options.PermitLimit,
		isAutoReplenishing: options.AutoReplenishment,
		window:             options.Window,
		queue:              NewQueue[RequestRegistration](),
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
	if permitCount > f.permitLimit {
		return nil, &ArgumentError{Message: "permit count is higher then maximum number of permits available"}
	}

	if permitCount == 0 {

	}

	f.mutex.Lock()
	defer f.mutex.Unlock()

	success, lease, err := f.tryLeaseUnsynchronized(permitCount)
	if err != nil {
		return nil, err
	}
	if success {
		return lease, nil
	}

	return &FixedWindowLease{isAcquired: false}, nil
}

func (f *FixedWindowRateLimiter) ReplenishmentPeriod() time.Duration {
	return f.window
}

func (f *FixedWindowRateLimiter) IsAutoReplenishing() bool {
	return f.isAutoReplenishing
}

func (f *FixedWindowRateLimiter) TryReplenish() (bool, error) {
	if f.isAutoReplenishing {
		return false, nil
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()

	return false, nil
}

func (f *FixedWindowRateLimiter) tryLeaseUnsynchronized(requestCount int) (bool, *FixedWindowLease, error) {
	if f.requestCount >= requestCount && f.requestCount != 0 {
		// Edge case where the lock show 0 availabe permits
		if requestCount == 0 {
			return true, &FixedWindowLease{
				isAcquired: true,
				retryAfter: 0,
			}, nil
		}

		if f.queue.GetSize() == 0 {
			f.idleSince = 0
			f.requestCount -= requestCount
			return true, &FixedWindowLease{
				isAcquired: true,
				retryAfter: 0,
			}, nil
		}
	}

	return false, nil, nil
}

type RequestRegistration struct {
	Count int
}

func NewRequestRegistration(requestCount int) RequestRegistration {
	return RequestRegistration{Count: requestCount}
}

// Force FixedWindowRateLimiter to adhere to RateLimiter and ReplenishingRateLimiter interface
var (
	_ RateLimiter             = &FixedWindowRateLimiter{}
	_ ReplenishingRateLimiter = &FixedWindowRateLimiter{}
)
