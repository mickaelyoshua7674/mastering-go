package main
import (
    "fmt"
    "math/rand"
)

func main() {
    min := 0
    max := 100
    numRand := 100
    var seed int64 = 10
    rand.New(rand.NewSource(seed))

    for i := 0; i<numRand; i++ {
        fmt.Print(random(min, max), " ")
    }

}

func random(min, max int) int {
    return rand.Intn(max-min) + min
}