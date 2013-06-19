package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gopher(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go gopher("I am Gopher")
	fmt.Println("Hello Gopher!")
	time.Sleep(2 * time.Second)
	fmt.Println("Bye Gopher!")
}
