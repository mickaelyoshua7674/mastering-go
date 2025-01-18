package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		gen(0, 2*n, createNumber, end)
		wg.Done()
	}()

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	end <- true
	wg.Wait()
	fmt.Println("Exiting...")
}

func gen(min, max int, createNumber chan int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- random.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			// return
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}