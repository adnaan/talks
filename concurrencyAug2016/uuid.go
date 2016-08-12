//START OMIT
package main

import "fmt"

func main() {
	id := make(chan string)
	go func() {
		var counter int64 = 0
		for {
			//blocked on receive
			id <- fmt.Sprintf("%x", counter)
			counter++
		}
	}()

	fmt.Println("id 0: ", <-id)
	fmt.Println("id 1: ", <-id)
}

//END OMIT
