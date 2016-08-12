//source: http://www.slideshare.net/jgrahamc/go-oncurrency
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type job struct {
	url  string
	resp chan *http.Response
}

type worker struct {
	jobs  chan *job
	count int
}

func (w *worker) getter(done chan *worker) {
	for {
		// listen for job
		j := <-w.jobs
		resp, _ := http.Get(j.url)
		j.resp <- resp
		// job is done
		done <- w
	}
}

func get(jobs chan *job, url string, answer chan string) {

	resp := make(chan *http.Response)
	//queue job
	jobs <- &job{url, resp}

	//wait and receive response
	r := <-resp

	if r != nil {
		body, _ := ioutil.ReadAll(r.Body)
		answer <- string(body)
		return
	}

	// reply back
	answer <- "Error!!"
}

func balancer(count int, depth int) chan *job {

	jobs := make(chan *job)
	done := make(chan *worker)
	workers := make([]*worker, count)

	// initialize workers
	for i := 0; i < count; i++ {
		workers[i] = &worker{make(chan *job, depth), 0}
		go workers[i].getter(done)
	}

	go func() {
		for {

			var free *worker
			min := depth

			// get a free worker with maximum slots open
			for _, w := range workers {
				if w.count < min {
					free = w
					min = w.count
				}
			}

			// job sourcer
			var jobsource chan *job

			if free != nil {
				jobsource = jobs
			}

			select {
			case j := <-jobsource:
				// send receiver job to worker with free slots
				free.jobs <- j
				free.count++
			case w := <-done:
				w.count--
			}

		}
	}()

	return jobs
}

func main() {
	// a balancer with 10 workers with 10 jobs each
	jobs := balancer(10, 10)
	answer := make(chan string)

	for i := 1; i < 100; i++ {
		go get(jobs, "https://httpbin.org/delay/"+strconv.Itoa(i), answer)
	}

	for u := range answer {
		fmt.Printf("%s\n", u)
	}
}
