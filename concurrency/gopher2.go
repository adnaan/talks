package main

import (
	"fmt"
	"math/rand"
	"time"
)

//GOPHER OMIT
func gopher(msg string, c chan string) {
	//notice the loop keeps going forever
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // sending a string value on c chan string
		//simulate some work.sleep interval random!
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

//ENDGOPHER OMIT

//MAIN OMIT
func main() {
	c := make(chan string)     //declared and initialzied a channel of type string
	go gopher("I am cool!", c) //launch the goroutine,give it the channel to communicate back on
	// we will listen to only 5 channel messages from gopher since we are busy folk.
	for i := 0; i < 5; i++ {
		fmt.Printf("Gopher says: %q\n", <-c) // receiving a string value from the channel
	}
	//ok enough! bye!
	fmt.Println("Main: You talk too much. Bye!")
}

//ENDMAIN OMIT
