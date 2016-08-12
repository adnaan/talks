//source: https://gobyexample.com/rate-limiting
package main

import "time"
import "fmt"

func main() {
	//START1 OMIT
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	//END1 OMIT
	//START2 OMIT
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Second * 1) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 6)
	for i := 1; i <= 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
	//END2 OMIT
}
