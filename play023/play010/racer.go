package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTIme(a)
	bDuration := measureResponseTIme(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTIme(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
