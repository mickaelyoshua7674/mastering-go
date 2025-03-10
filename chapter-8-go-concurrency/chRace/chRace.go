package main
import (
    "fmt"
)

func main() {
    var ch chan bool = make(chan bool)
    // Write 5 values to channel with a single goroutine
    go printer(ch, 5)
    // IMPORTANT: As the channel ch is closed,
    // the range loop is foinf to exit on its own.
    for val := range ch {
        fmt.Print(val, " ")
    }
    fmt.Println()
    for i:= 0; i<15; i++ {
        fmt.Print(<-ch, " ")
    }
    fmt.Println()
}

func printer(ch chan bool, times int) {
    for i:=0; i<times; i++ {
        ch <- true
    }
    close(ch)
}