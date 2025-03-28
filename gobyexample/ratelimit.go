package main

import (
	"fmt"
	"time"
	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(2, 5) // 2 requests/sec with a burst of 5

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d allowed at %s\n", i, time.Now().Format("15:04:05"))
		} else {
			fmt.Printf("Request %d denied at %s\n", i, time.Now().Format("15:04:05"))
		}
		time.Sleep(300 * time.Millisecond)
	}

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter2 := time.Tick(1000 * time.Millisecond)
	for req := range requests {
		<-limiter2
		fmt.Printf("Request %d allowed at %s\n", req, time.Now().Format("15:04:05"))
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }
	go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

	burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
// Output:
// Request 1 allowed at 15:04:05
// Request 2 allowed at 15:04:05
// Request 3 denied at 15:04:05
// Request 4 denied at 15:04:05
// Request 5 denied at 15:04:05
// Request 6 allowed at 15:04:05
// Request 7 allowed at 15:04:05
// Request 8 denied at 15:04:05
// Request 9 denied at 15:04:05
// Request 10 denied at 15:04:05