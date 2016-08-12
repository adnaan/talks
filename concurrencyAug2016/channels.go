package main

import "fmt"

func main() {

	//INIT OMIT
	// Declaring and initializing.
	// a channel which allows "int" type of message
	var c chan int
	c = make(chan int)
	// or
	c := make(chan int)
	//ENDINIT OMIT

	//SEND OMIT
	// Sending on a channel.
	c <- 1
	//ENDSEND OMIT

	//REC OMIT
	// Receiving from a channel.
	// The "arrow" indicates the direction of data flow.
	val = <-c
	//ENDREC OMIT

}
