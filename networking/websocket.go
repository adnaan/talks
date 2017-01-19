//S1 OMIT
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", websocket.Handler(WSHandler))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//S2 OMIT
