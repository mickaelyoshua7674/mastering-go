package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var random = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	c := make(chan int)
	waitGroup.Add(1)
	go add(c)
	go send(c)
	waitGroup.Wait()
}

func send(c chan int) {
	for {
		c <- random.Intn(10)
	}
}

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)
	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C:
			c = nil
			fmt.Println(sum)
			waitGroup.Done()
		}
	}
}