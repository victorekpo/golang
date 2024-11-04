package utils

import (
	"batch-v1/internal/config"
	"math"
	"time"
)

// ExponentialBackoff calculates the backoff duration based on the retry attempt.
func ExponentialBackoff(attempt int) time.Duration {
	initialBackoff := float64(config.InitialBackoff)
	maxBackoff := float64(config.MaxBackoff)
	backoff := initialBackoff * math.Pow(2, float64(attempt))
	if backoff > maxBackoff {
		backoff = maxBackoff
	}
	return time.Duration(backoff)
}
