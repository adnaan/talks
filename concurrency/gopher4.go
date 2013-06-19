package main

import (
	"fmt"
	"math/rand"
	"time"
)

//GOPHER OMIT
func gopher(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

//ENDGOPHER OMIT

//MAIN OMIT
func main() {
	bat := gopher("Bat")
	robin := gopher("Robin")
	for i := 0; i < 5; i++ {
		fmt.Printf("I am Gopher%v\n", <-bat)
		fmt.Printf("I am Gopher%v\n", <-robin)
	}
	fmt.Println("You both talk too much. Bye!")
}

//ENDMAIN OMIT
