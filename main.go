package main

import (
	"fmt"
	"github.com/ZaharBorisenko/go-rate-limiting/lib"
	"github.com/ZaharBorisenko/go-rate-limiting/limiter"
	"github.com/ZaharBorisenko/go-rate-limiting/middleware"
	"log"
	"net/http"
)

const (
	addr = "127.0.0.1:8080"
)

type User struct {
	Name     string
	UserName string
}

func main() {
	mux := http.NewServeMux()

	lib.RegisterRoutes(mux)
	handler := middleware.RateLimiterMiddleware(mux, limiter.NewSlidingWindowLimiter(1, 2))

	fmt.Println("Server starting on address:", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal("Error ListenAndServe", err)
	}
}
