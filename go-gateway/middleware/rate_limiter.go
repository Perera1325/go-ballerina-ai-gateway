package middleware

import (
	"net/http"
	"sync"
	"time"
)

var (
	lastResetTime  time.Time
	requestCounter int
	requestLimit   = 5
	mutex          sync.Mutex
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		mutex.Lock()
		defer mutex.Unlock()

		now := time.Now()

		if now.Sub(lastResetTime) > time.Second {
			requestCounter = 0
			lastResetTime = now
		}

		requestCounter++

		if requestCounter > requestLimit {
			http.Error(w, "Rate limit exceeded (5 req/sec)", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}