package main

import (
	"log"
	"net/http"

	"github.com/Perera1325/go-ballerina-ai-gateway/go-gateway/middleware"
	"github.com/Perera1325/go-ballerina-ai-gateway/go-gateway/proxy"
)

func main() {

	targetService := "http://localhost:9090"

	reverseProxy := proxy.NewReverseProxy(targetService)

	handler := middleware.RateLimiter(reverseProxy)

	http.Handle("/", handler)

	log.Println("🚀 Go Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}