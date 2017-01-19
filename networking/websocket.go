//S1 OMIT
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func WSHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
func main() {
	http.Handle("/", websocket.Handler(WSHandler))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

//S2 OMIT

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
