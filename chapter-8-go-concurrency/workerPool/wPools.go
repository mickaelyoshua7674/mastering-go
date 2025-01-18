package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)


var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)

	var wg sync.WaitGroup
	for i:=0; i<nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
	}()

	wg.Wait()
	close(data)

	fmt.Println("Finished")
}

type Client struct {
	id int
	integer int
}

type Result struct {
	job Client
	square int
}

func worker(wg *sync.WaitGroup) {
	for c := range clients { // with clients closed it will stop after the last element
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	for i:=0; i<n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients) // can read but no longer can insert.
					// With this channel closed the for range will
					// stop after go through the last element
}

