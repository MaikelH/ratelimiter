package ratelimiter

import "fmt"

type ArgumentError struct {
	Message string
	Err     error
}

func (e *ArgumentError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s caused by %s", e.Message, e.Err.Error())
	}
	return e.Message
}

func (e *ArgumentError) Unwrap() error { return e.Err }

func (e *ArgumentError) Is(target error) bool {
	_, ok := target.(*ArgumentError)

	return ok
}
