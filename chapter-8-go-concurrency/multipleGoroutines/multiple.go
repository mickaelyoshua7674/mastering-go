package main
import (
    "fmt"
    "time"
    "os"
    "strconv"
)

func main() {

    if len(os.Args) == 1 {
        fmt.Println("Give one argument!")
        return
    }
    count, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Going to create %d goroutines.\n", count)

    for i:=0; i<count; i++ {
        go func(x int) {
            fmt.Printf("%d ", x)
        }(i)
    }
    time.Sleep(time.Second)
    fmt.Println("\nExiting...")
}