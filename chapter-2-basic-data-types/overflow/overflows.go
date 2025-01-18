package main
import (
    "fmt"
    "math"
)

func main() {
    i := math.MaxInt - 100
    for {
        if i==math.MaxInt {
            fmt.Println("Max:", i)
            break
        }
        i++
    }

    i = math.MinInt + 100
    for {
        if i==math.MinInt {
            fmt.Println("Min:", i)
            break
        }
        i--
    }
}