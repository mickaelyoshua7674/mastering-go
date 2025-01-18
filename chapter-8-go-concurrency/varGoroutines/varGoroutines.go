package main
import (
    "fmt"
    "os"
    "strconv"
    "sync"
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

    var waitGroup sync.WaitGroup
    fmt.Printf("%#v\n", waitGroup)
    for i:=0; i<count; i++ {
        waitGroup.Add(1) // avoid race conditions
        go func(x int) {
            defer waitGroup.Done() // call 'waitGroup.Add(-1)'
            fmt.Printf("%d ", x)
        }(i)
    }
    fmt.Printf("%#v\n", waitGroup)
    waitGroup.Wait() // wait for the counter in 'waitGroup' become 0
}