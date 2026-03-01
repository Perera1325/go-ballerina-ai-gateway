package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(targetURL string) http.Handler {

	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal("Invalid target URL:", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-API-Gateway", "Go-Gateway-v2")

		log.Printf("Proxying request: %s %s", r.Method, r.URL.Path)

		proxy.ServeHTTP(w, r)
	})
}