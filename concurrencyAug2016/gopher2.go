package main

import (
	"fmt"
	"math/rand"
	"time"
)

//GOPHER OMIT
func gopher(msg string, c chan string) {
	for i := 0; ; i++ { //notice the loop keeps going forever
		c <- fmt.Sprintf("%s %d", msg, i)                            // sending a string value on c chan string
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond) //simulate some work.sleep interval random!
	}
}

//ENDGOPHER OMIT

//MAIN OMIT
func main() {
	c := make(chan string)     //declared and initialzied a channel of type string
	go gopher("I am cool!", c) //launch the goroutine,pass channel to communicate back on
	// we will listen to only 5 channel messages from gopher since we are busy folk.
	for i := 0; i < 5; i++ {
		fmt.Printf("Gopher says: %q\n", <-c) // receiving a string value from the channel
	}
	//ok enough! bye!
	fmt.Println("Main: You talk too much. Bye!")
}

//ENDMAIN OMIT
