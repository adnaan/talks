//source: https://blog.cloudflare.com/recycling-memory-buffers-in-go/
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

//START1 OMIT
func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	// our big pool
	pool := make([][]byte, 20)
	// pool of buffers
	buffer := make(chan []byte, 5)

	var m runtime.MemStats
	makes := 0
	for {
		var b []byte
		select {
		// retrieve a buffer if pool has previously released buffers
		case b = <-buffer:
		default:
			// or create a new buffer
			makes++
			b = makeBuffer()
		}
		//END1 OMIT
		//START2 OMIT
		i := rand.Intn(len(pool))
		if pool[i] != nil {
			select {
			// release buffers to the buffer pool
			case buffer <- pool[i]:
				pool[i] = nil
			default:
			}
		}

		// store created buffer in the big pool
		pool[i] = b

		time.Sleep(time.Second)
		//END2 OMIT
		//START3 OMIT

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, makes)
	}
}

//END3 OMIT
