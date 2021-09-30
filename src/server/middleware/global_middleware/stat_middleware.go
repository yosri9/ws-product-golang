package global_middleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

// limit the re
var (
	w       http.ResponseWriter
	rt      = rate.Every(1 * time.Hour / 1)
	limiter = rate.NewLimiter(rt, 2)
)

func GlobalMiddleware() {
	if limiter.Allow() == false {
		http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
		return
	}
}
