package ratelimiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedWindowInvalidOptions(t *testing.T) {
	options := FixedWindowRateLimiterOptions{}

	limiter, err := NewFixedWindowRateLimiter(options)
	assert.Nil(t, err)
	assert.NotNil(t, limiter)

	options.PermitLimit = -1
	limiter, err = NewFixedWindowRateLimiter(options)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, &ArgumentError{Message: "permit limit must be higher then 0"})

	options = FixedWindowRateLimiterOptions{}
	options.QueueLimit = -1
	limiter, err = NewFixedWindowRateLimiter(options)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, &ArgumentError{Message: "queue limit must be higher then 0"})
}

func TestFixedWindowRateLimiter_FailsWhenAcquiringMoreThanLimit(t *testing.T) {
	options := FixedWindowRateLimiterOptions{PermitLimit: 2}

	limiter, err := NewFixedWindowRateLimiter(options)

	_, err = limiter.Acquire(3)
	if err != nil {
		assert.ErrorIs(t, err, &ArgumentError{Message: "permit count is higher then maximum number of permits available"})
		return
	}
	t.Errorf("acquire should have failed")
}

func TestFixedWindowRateLimiter_Acquire(t *testing.T) {
	options := FixedWindowRateLimiterOptions{PermitLimit: 1}

	limiter, err := NewFixedWindowRateLimiter(options)
	assert.Nil(t, err)
	assert.NotNil(t, limiter)

	// First acquire should succeed
	lease, err := limiter.Acquire(1)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, true, lease.IsAcquired())

	// Second acquire should fail due permit Count of 1
	lease, err = limiter.Acquire(1)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, false, lease.IsAcquired())
	lease.Release()

	// Although lease is released limiter is not replenished, so acquire is false
	lease, err = limiter.Acquire(1)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, false, lease.IsAcquired())

	// Replenish limiter
	replenish, err := limiter.TryReplenish()
	if err != nil {
		t.Error(err)
		return
	}
	assert.True(t, replenish)

	lease, err = limiter.Acquire(1)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, true, lease.IsAcquired())
}
