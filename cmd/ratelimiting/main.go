// cmd/ratelimiting/main.go
package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization
// and maintaining quality of service. Go elegantly supports rate limiting with
// goroutines, channels, and tickers.

func main() {
	// First we'll look at basic rate limiting. Suppose we want to limit our
	// handling of incoming requests. We'll serve these requests off a channel
	// of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds. This is
	// the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on receive from the limiter channel before serving each
	// request, we limit ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
