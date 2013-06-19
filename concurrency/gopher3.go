package main

import (
	"fmt"
	"math/rand"
	"time"
)

//GOPHER OMIT
func gopher(msg string) <-chan string { // returns a receive only channel
	c := make(chan string)
	go func() { //invoke goroutine using a function literal.
		// In Go function literals are closures.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}() //call the function!

	return c
}

//ENDGOPHER OMIT

//MAIN OMIT
func main() {

	c := gopher("I am cool!") //receive the channel
	for i := 0; i < 5; i++ {
		fmt.Printf("Gopher says: %q\n", <-c)
	}
	fmt.Println("Main: You talk too much. Bye!")
}

//ENDMAIN OMIT
