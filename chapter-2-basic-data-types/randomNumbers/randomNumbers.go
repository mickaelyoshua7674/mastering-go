package main
import (
    "fmt"
    "os"
    "strconv"
    "math/rand"
)

func main() {
    min := 0
    max := 100
    total := 100
    var seed int64 = 10

    args := os.Args[1:]

    for i, arg := range args {
        t, err := strconv.Atoi(arg)
        if err != nil {
            continue
        }

        switch i {
        case 0:
            min = t
            max = min + 100
        case 1:
            max = t
        case 2:
            total = t
        case 3:
            seed = int64(t)
        }
    }

    fmt.Println("Usage: ./randomNumbers.go MIN MAX TOTAL SEED")
    switch len(args) {
    case 0:
        fmt.Println("Using all default values!")
    case 1, 2, 3:
        fmt.Println("Using some default values!")
    default:
        fmt.Println("Using given values!")
    }

    r := rand.New(rand.NewSource(seed))
    for i := 0; i<total; i++ {
        fmt.Print(random(min, max, r), " ")
    }
    fmt.Println()

}

func random(min, max int, r *rand.Rand) int {
    return r.Intn(max-min) + min
}