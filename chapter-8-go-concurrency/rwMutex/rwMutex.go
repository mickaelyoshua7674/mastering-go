package main

import (
	"fmt"
	"sync"
	"time"
)

type secret struct {
	RWM sync.RWMutex
	password string
}

var Password *secret
var wg sync.WaitGroup

func main() {
	Password = &secret{password: "myPass"}
	for i:=0; i<3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("123456")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("54321")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}

func Change(pass string) {
	if Password == nil {
		fmt.Println("Password is nil!")
		return
	}
	fmt.Println("Change() function")
	Password.RWM.Lock()
	fmt.Println("Change() Locked")
	time.Sleep(4*time.Second)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() Unlocked")
}

func show() {
	defer wg.Done()
	defer Password.RWM.RUnlock()
	Password.RWM.RLock()

	fmt.Println("Show function locked!")
	time.Sleep(2*time.Second)
	fmt.Println("Pass value:", Password.password)
}