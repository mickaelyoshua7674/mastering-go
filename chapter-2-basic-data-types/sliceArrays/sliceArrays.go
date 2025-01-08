package main
import (
    "fmt"
)

func main() {
    a := [4]string{"Zero", "One", "Two", "Three"}
    fmt.Println("a:", a)

    var S0 = a[0:1]
    fmt.Println(S0)
    S0[0] = "S0"

    var S12 = a[1:3]
    fmt.Println(S12)
    S12[0] = "S12_0"
    S12[1] = "S12_1"

    fmt.Println("a:", a)

    // Changes to slice -> changes to array
    change(S12)
    fmt.Println("a:", a)
}

func change(s []string) {
    s[0] = "Change_function"
}