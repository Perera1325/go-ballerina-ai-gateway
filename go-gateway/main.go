package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

var (
	requestCount   int
	mutex          sync.Mutex
	lastResetTime  time.Time
	requestCounter int
	requestLimit   = 5 // 5 requests per second
)

func main() {

	target, err := url.Parse("http://localhost:9090")
	if err != nil {
		log.Fatal("Invalid target URL:", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	lastResetTime = time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Rate limiting
		if !allowRequest() {
			http.Error(w, "Rate limit exceeded (5 req/sec)", http.StatusTooManyRequests)
			return
		}

		// Add Gateway Header
		w.Header().Set("X-API-Gateway", "Go-Gateway-v1")

		// Log Request
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

		// Simulate AI Usage Counter
		mutex.Lock()
		requestCount++
		log.Printf("Total API Requests Processed: %d", requestCount)
		mutex.Unlock()

		// Forward request to Ballerina service
		proxy.ServeHTTP(w, r)
	})

	log.Println("🚀 Go Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func allowRequest() bool {
	now := time.Now()

	if now.Sub(lastResetTime) > time.Second {
		requestCounter = 0
		lastResetTime = now
	}

	requestCounter++

	return requestCounter <= requestLimit
}