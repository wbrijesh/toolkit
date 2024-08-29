package middleware

import (
	"net/http"
	"sync"
	"time"
)

type Middleware func(http.Handler) http.Handler

type rateLimiter struct {
	resetTime     time.Time
	requests      int
	maxRequests   int
	resetDuration time.Duration
	mu            sync.Mutex
}

func newRateLimiter(maxRequests int, duration time.Duration) *rateLimiter {
	return &rateLimiter{
		maxRequests:   maxRequests,
		resetDuration: duration,
		resetTime:     time.Now().Add(duration),
	}
}

func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if time.Now().After(rl.resetTime) {
		rl.requests = 0
		rl.resetTime = time.Now().Add(rl.resetDuration)
	}

	if rl.requests >= rl.maxRequests {
		return false
	}

	rl.requests++
	return true
}

func RateLimit(maxRequests int, duration time.Duration) Middleware {
	limiter := newRateLimiter(maxRequests, duration)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.allow() {
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
