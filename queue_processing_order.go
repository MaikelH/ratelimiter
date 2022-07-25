package ratelimiter

type QueueProcessingOrder int

const (
	OldestFirst QueueProcessingOrder = iota
	NewestFirst
)
