package middleware

import (
	"github.com/ZaharBorisenko/go-rate-limiting/lib"
	"github.com/ZaharBorisenko/go-rate-limiting/limiter"
	"log"
	"net"
	"net/http"
	"sync"
)

func RateLimiterMiddleware(next http.Handler, RateLimiter limiter.RateLimiter) http.Handler {
	ipLimiterMap := map[string]limiter.RateLimiter{}
	mutex := sync.Mutex{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//user ip
		ip, port, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err, port)
		}

		//limit
		mutex.Lock()
		ipLimiter, exists := ipLimiterMap[ip]
		if !exists {
			ipLimiter = RateLimiter
		}
		mutex.Unlock()

		if !ipLimiter.Allow() {
			lib.WriteError(w, http.StatusTooManyRequests, "too many requests... wait")
			return
		}

		next.ServeHTTP(w, r)
	})
}
