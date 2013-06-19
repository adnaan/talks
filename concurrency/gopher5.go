package main

import (
	"fmt"
)

//FIRST OMIT
func first(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	//the value sent to c is either <-input1 or <-input2 depending upon who receives first
	return c
}

//ENDFIRST OMIT

//GOPHER OMIT
func gopher(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	return c
}

//ENDGOPHER OMIT

//MAIN OMIT
func main() {
	c := first(gopher("Bat"), gopher("Robin"))
	for i := 0; i < 10; i++ {
		fmt.Printf("I am Gopher%v\n", <-c)
	}
	fmt.Println("You both talk too much. Bye!")
}

//ENDMAIN OMIT
